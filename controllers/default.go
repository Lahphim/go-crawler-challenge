package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Layout = "layouts/application.html"
	c.TplName = "main/index.html"
}
