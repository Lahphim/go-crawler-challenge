package controllers

import "github.com/beego/beego/v2/server/web"

type MainController struct {
	BaseController
}

func (c *MainController) NestPrepare() {
	c.requireAuthenticatedUser = true
}

func (c *MainController) Get() {
	web.ReadFromRequest(&c.Controller)

	c.Layout = "layouts/application.html"
	c.TplName = "main/index.html"
}
