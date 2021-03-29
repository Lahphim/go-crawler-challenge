package apicontrollers

import (
	"net/http"
	"strconv"

	"github.com/beego/beego/v2/server/web"
	. "github.com/google/jsonapi"
)

const ContentType = "application/vnd.api+json; charset=utf-8"

type BaseController struct {
	web.Controller
}

func (c *BaseController) Prepare() {
	c.disableXSRF()
}

func (c *BaseController) RenderJSON(data interface{}) error {
	c.Ctx.Output.Header("Content-Type", ContentType)

	err := MarshalPayload(c.Ctx.ResponseWriter, data)
	if err != nil {
		http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)

		return err
	}

	return nil
}

func (c *BaseController) RenderGenericError(err error) error {
	return c.renderError("Generic Error", err.Error(), http.StatusInternalServerError, "generic_error", nil)
}

func (c *BaseController) RenderUnauthorizedError(err error) error {
	return c.renderError("Unauthorized Error", err.Error(), http.StatusUnauthorized, "unauthorized_error", nil)
}

func (c *BaseController) disableXSRF() {
	c.EnableXSRF = false
}

func (c *BaseController) renderError(title string, detail string, status int, code string, meta *map[string]interface{}) (err error) {
	c.Ctx.Output.Header("Content-Type", ContentType)
	c.Ctx.Output.SetStatus(status)

	writer := c.Ctx.ResponseWriter

	err = MarshalErrors(writer, []*ErrorObject{{
		Title:  title,
		Detail: detail,
		Status: strconv.Itoa(status),
		Code:   code,
		Meta:   meta,
	}})
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)

		return err
	}

	return nil
}
