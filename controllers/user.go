package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

//  UserController operations for User
type UserController struct {
	beego.Controller
}

// URLMapping ...
func (c *UserController) URLMapping() {
	c.Mapping("New", c.New)
	c.Mapping("Create", c.Create)
}

func (c *UserController) New() {
	c.Layout = "layouts/application.html"
	c.TplName = "user/new.html"
}

func (c *UserController) Create() {

}
