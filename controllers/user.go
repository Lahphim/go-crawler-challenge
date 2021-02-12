package controllers

import (
	"net/http"

	form "go-crawler-challenge/forms/user"

	"github.com/beego/beego/v2/server/web"
)

// UserController operations for User
type UserController struct {
	BaseController
}

// NestPrepare prepares some configurations to the controller
func (c *UserController) NestPrepare() {
	c.requireGuestUser = true
}

// URLMapping maps user controller actions to functions
func (c *UserController) URLMapping() {
	c.Mapping("New", c.New)
	c.Mapping("Create", c.Create)
}

// New handles a form for creating a new user
// @Title New
// @Description show a new user form
// @Success 200
// @router / [get]
func (c *UserController) New() {
	web.ReadFromRequest(&c.Controller)

	c.Layout = "layouts/authentication.html"
	c.TplName = "user/new.html"
}

// Create handles validation and adding a new unique user
// @Title Create
// @Description create a new unique user
// @Success 302 redirect to the sign-up page
// @Failure 302 redirect to the sign-up page and print some error messages
// @router / [post]
func (c *UserController) Create() {
	flash := web.NewFlash()
	registrationForm := form.RegistrationForm{}
	redirectPath := "/user/sign_up"

	err := c.ParseForm(&registrationForm)
	if err != nil {
		flash.Error(err.Error())
	}

	_, errors := registrationForm.Create()
	if len(errors) > 0 {
		flash.Error(errors[0].Error())
	} else {
		flash.Success("Congrats on creating a new account")
		redirectPath = "/user/sign_in"
	}

	flash.Store(&c.Controller)
	c.Redirect(redirectPath, http.StatusFound)
}
