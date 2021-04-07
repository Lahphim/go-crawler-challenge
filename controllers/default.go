package controllers

import "github.com/beego/beego/v2/server/web"

type MainController struct {
	BaseController
}

// NestPrepare prepares some configurations to the controller
func (c *MainController) NestPrepare() {
	c.actionPolicyMapping()
}

// URLMapping maps main controller actions to functions
func (c *MainController) URLMapping() {
	c.Mapping("Index", c.Index)
}

// actionPolicyMapping maps main controller actions to policies
func (c *MainController) actionPolicyMapping() {
	c.MappingPolicy("Index", Policy{})
}

// Index handles public landing page
// @Title Index
// @Description show website objective
// @Success 200
// @router / [get]
func (c *MainController) Index() {
	web.ReadFromRequest(&c.Controller)

	c.TplName = "main/index.html"
}
