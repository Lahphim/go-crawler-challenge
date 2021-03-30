package apiv1controllers

import (
	"fmt"
	"net/http"

	. "go-crawler-challenge/controllers/api"
	form "go-crawler-challenge/forms/session"
	"go-crawler-challenge/services/oauth"

	"github.com/go-oauth2/oauth2/v4/errors"
)

// TokenController operations for Token
type TokenController struct {
	BaseController
}

// URLMapping maps token controller actions to functions
func (c *TokenController) URLMapping() {
	c.Mapping("Create", c.Create)
}

// Create handles token generator by authenticate some client credentials and user credentials
// @Title Create
// @Description generate a token information
// @Success 200
// @router /api/v1/oauth/token [post]
func (c *TokenController) Create() {
	oauth.ServerOauth.SetPasswordAuthorizationHandler(passwordAuthorizationHandler)

	err := oauth.ServerOauth.HandleTokenRequest(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusForbidden)
	}
}

func passwordAuthorizationHandler(email string, password string) (string, error) {
	authenticationForm := form.AuthenticationForm{
		Email:    email,
		Password: password,
	}

	user, errorList := authenticationForm.Authenticate()
	if len(errorList) > 0 {
		return "", errors.ErrInvalidClient
	}

	return fmt.Sprint(user.Id), nil
}
