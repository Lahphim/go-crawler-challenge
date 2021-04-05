package apicontrollers

import (
	"net/http"

	"github.com/beego/beego/v2/server/web"
	"github.com/google/jsonapi"
)

const ContentType = "application/vnd.api+json; charset=utf-8"

type BaseController struct {
	web.Controller
}

func (c *BaseController) Prepare() {
	c.disableXSRF()
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
