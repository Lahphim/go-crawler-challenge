package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

// DashboardController operations for Dashboard
type DashboardController struct {
	BaseController
}

// NestPrepare prepares some configurations to the controller
func (c *DashboardController) NestPrepare() {
	c.actionPolicyMapping()
}

// URLMapping maps dashboard controller actions to functions
func (c *DashboardController) URLMapping() {
	c.Mapping("Index", c.Index)
}

// actionPolicyMapping maps dashboard controller actions to policies
func (c *DashboardController) actionPolicyMapping() {
	c.MappingPolicy("Index", Policy{requireAuthenticatedUser: true})
}

// Index handles dashboard with handy widgets
// @Title Index
// @Description show some widgets such as search, listing and summary detail
// @Success 200
// @router / [get]
func (c *DashboardController) Index() {
	web.ReadFromRequest(&c.Controller)

	c.Layout = "layouts/application.html"
	c.TplName = "dashboard/index.html"
}
