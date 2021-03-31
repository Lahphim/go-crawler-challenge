package apiv1controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	. "go-crawler-challenge/controllers/api"
	v1serializers "go-crawler-challenge/serializers/v1"
	"go-crawler-challenge/services/oauth"
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
	writer := httptest.NewRecorder()

	err := oauth.ServerOauth.HandleTokenRequest(writer, c.Ctx.Request)
	if err != nil {
		c.RenderUnauthorizedError(err)

		return
	}

	jsonResponse := writer.Body.Bytes()

	if writer.Code != 200 {
		c.handleResponseError(writer)

		return
	}

	var tokenResponseObject v1serializers.TokenInformation

	err = json.Unmarshal(jsonResponse, &tokenResponseObject)
	if err != nil {
		c.RenderGenericError(err)

		return
	}

	c.RenderJSON(tokenResponseObject.Data())
}

func (c *TokenController) handleResponseError(writer *httptest.ResponseRecorder) {
	var tokenErrorObject v1serializers.TokenError
	jsonResponse := writer.Body.Bytes()

	err := json.Unmarshal(jsonResponse, &tokenErrorObject)
	if err != nil {
		c.RenderGenericError(err)
	} else {
		c.RenderError(http.StatusText(writer.Code), tokenErrorObject.Detail, writer.Code, tokenErrorObject.Code)
	}
}
