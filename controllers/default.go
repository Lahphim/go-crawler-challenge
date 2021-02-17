package controllers

import "github.com/beego/beego/v2/server/web"

type MainController struct {
	BaseController
}

// URLMapping maps main controller actions to functions
func (c *MainController) URLMapping() {
	c.Mapping("Index", c.Index)
}

// Index handles public landing page
// @Title Index
// @Description show website objective
// @Success 200
// @router / [get]
func (c *MainController) Index() {
	web.ReadFromRequest(&c.Controller)

	c.Layout = "layouts/application.html"
	c.TplName = "main/index.html"
}
