package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type MainController struct {
	web.Controller
}

func (c *MainController) Get() {
	c.Layout = "layouts/application.html"
	c.TplName = "main/index.html"
}
