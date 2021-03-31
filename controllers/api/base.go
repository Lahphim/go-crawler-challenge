package apicontrollers

import (
	"fmt"
	"go-crawler-challenge/models"
	"go-crawler-challenge/services/oauth"
	"net/http"
	"strconv"

	"github.com/go-oauth2/oauth2/v4/errors"

	"github.com/go-oauth2/oauth2/v4"

	"github.com/beego/beego/v2/server/web"
	"github.com/google/jsonapi"
)

const ContentType = "application/vnd.api+json; charset=utf-8"

type NestPreparer interface {
	NestPrepare()
}

type BaseController struct {
	web.Controller

	CurrentUser      *models.User
	CurrentTokenInfo oauth2.TokenInfo
	actionPolicy     map[string]Policy
}

type Policy struct {
	RequireAuthenticatedUser bool
}

func (c *BaseController) Prepare() {
	c.disableXSRF()
	c.initActionPolicy()

	app, ok := c.AppController.(NestPreparer)
	if ok {
		app.NestPrepare()
	}

	c.handleAuthorizeRequest()
}

func (c *BaseController) MappingPolicy(method string, policy Policy) {
	c.actionPolicy[method] = policy
}

func (c *BaseController) handleAuthorizeRequest() {
	_, actionName := c.GetControllerAndAction()
	actionPolicy := c.actionPolicy[actionName]

	if actionPolicy.RequireAuthenticatedUser {
		if c.ensureBearerToken() {
			currentUser, err := c.getTokenUser()
			if err != nil {
				// return internal
			}

			c.CurrentUser = currentUser
		} else {
			// return unauthenticated
		}
	}
}

func (c *BaseController) ensureBearerToken() bool {
	tokenInfo, err := oauth.ServerOauth.ValidationBearerToken(c.Ctx.Request)
	if err != nil {
		return false
	}

	c.CurrentTokenInfo = tokenInfo

	return true
}

func (c *BaseController) getTokenUser() (user *models.User, err error) {
	if c.CurrentTokenInfo == nil {
		return nil, errors.New(fmt.Sprintf("No current token exists"))
	}

	userId, err := strconv.Atoi(c.CurrentTokenInfo.GetUserID())
	if err != nil {
		return nil, err
	}

	currentUser, err := models.GetUserById(int64(userId))
	if err != nil {
		return nil, err
	}

	return currentUser, nil
}

func (c *BaseController) RenderJSON(data interface{}) {
	response, err := jsonapi.Marshal(data)
	if err != nil {
		c.RenderGenericError(err)

		return
	}

	c.Data["json"] = response
	err = c.ServeJSON()
	if err != nil {
		c.RenderGenericError(err)
	}
}

func (c *BaseController) RenderGenericError(err error) {
	statusCode := http.StatusInternalServerError

	c.RenderError(http.StatusText(statusCode), err.Error(), statusCode, "generic_error")
}

func (c *BaseController) RenderUnauthorizedError(err error) {
	statusCode := http.StatusUnauthorized

	c.RenderError(http.StatusText(statusCode), err.Error(), statusCode, "unauthorized_error")
}

func (c *BaseController) RenderError(title string, detail string, status int, code string) {
	c.Ctx.Output.Header("Content-Type", ContentType)
	c.Ctx.ResponseWriter.WriteHeader(status)

	writer := c.Ctx.ResponseWriter

	err := jsonapi.MarshalErrors(writer, []*jsonapi.ErrorObject{{
		Title:  title,
		Detail: detail,
		Code:   code,
	}})
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func (c *BaseController) disableXSRF() {
	c.EnableXSRF = false
}

func (c *BaseController) initActionPolicy() {
	c.actionPolicy = make(map[string]Policy)
}
