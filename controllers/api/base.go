package apicontrollers

import (
	"github.com/beego/beego/v2/server/web"
)

const ContentType = "application/vnd.api+json; charset=utf-8"

type BaseController struct {
	web.Controller
}

func (c *BaseController) Prepare() {
	c.disableXSRF()
}

func (c *BaseController) disableXSRF() {
	c.EnableXSRF = false
}
