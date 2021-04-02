package apicontrollers

import (
	"net/http"
	"strconv"

	"go-crawler-challenge/models"
	"go-crawler-challenge/services/oauth"

	"github.com/beego/beego/v2/server/web"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/google/jsonapi"
)

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

func (c *BaseController) GetPageSize() (pageSize int) {
	return DefaultPageSize
}

func (c *BaseController) GetOrderBy() (orderBy []string) {
	return DefaultOrderBy
}

func (c *BaseController) MappingPolicy(method string, policy Policy) {
	c.actionPolicy[method] = policy
}

func (c *BaseController) RenderJSONMany(data interface{}, meta *jsonapi.Meta, links *jsonapi.Links, status int) {
	response, err := jsonapi.Marshal(data)
	if err != nil {
		c.RenderGenericError(err)

		return
	}

	payload, ok := response.(*jsonapi.ManyPayload)
	if !ok {
		c.RenderGenericError(ErrorInvalidPayloaderType)
	}

	if meta != nil {
		payload.Meta = meta
	}

	if links != nil {
		payload.Links = links
	}

	c.renderJSON(payload, status)
}

func (c *BaseController) RenderJSON(data interface{}, status int) {
	response, err := jsonapi.Marshal(data)
	if err != nil {
		c.RenderGenericError(err)

		return
	}

	c.renderJSON(response, status)
}

func (c *BaseController) renderJSON(payloader interface{}, status int) {
	c.Ctx.Output.Header("Content-Type", ContentType)
	c.Ctx.ResponseWriter.WriteHeader(status)

	c.Data["json"] = payloader
	err := c.ServeJSON()
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

func (c *BaseController) handleAuthorizeRequest() {
	_, actionName := c.GetControllerAndAction()
	actionPolicy := c.actionPolicy[actionName]

	if actionPolicy.RequireAuthenticatedUser {
		err := c.validateBearerToken()
		if err != nil {
			c.RenderUnauthorizedError(err)

			return
		}

		err = c.validateExistingUser()
		if err != nil {
			c.RenderUnauthorizedError(err)

			return
		}
	}
}

func (c *BaseController) validateBearerToken() error {
	tokenInfo, err := oauth.ServerOauth.ValidationBearerToken(c.Ctx.Request)
	if err != nil {
		return err
	}

	c.CurrentTokenInfo = tokenInfo

	return nil
}

func (c *BaseController) validateExistingUser() error {
	if c.CurrentTokenInfo == nil {
		return ErrorMissingAccessToken
	}

	userId, err := strconv.Atoi(c.CurrentTokenInfo.GetUserID())
	if err != nil {
		return ErrorInvalidUser
	}

	currentUser, err := models.GetUserById(int64(userId))
	if err != nil {
		return ErrorNotFoundUser
	}

	c.CurrentUser = currentUser

	return nil
}

func (c *BaseController) disableXSRF() {
	c.EnableXSRF = false
}

func (c *BaseController) initActionPolicy() {
	c.actionPolicy = make(map[string]Policy)
}
