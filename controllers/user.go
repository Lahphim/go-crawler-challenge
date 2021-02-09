package controllers

import (
	"fmt"
	"net/http"

	forms "go-crawler-challenge/forms/user"

	log "github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

//  UserController operations for User
type UserController struct {
	web.Controller
}

// URLMapping ...
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
	c.Layout = "layouts/application.html"
	c.TplName = "user/new.html"
}

// Create handles validation and adding a new unique user
// @Title Create
// @Description create a new unique user
// @Success 302 redirect to the signup page
// @Failure 302 edirect to the signup page and print some error messages
// @router / [post]
func (c *UserController) Create() {
	form := forms.RegistrationForm{}

	err := c.ParseForm(&form)
	if err != nil {
		log.Info(fmt.Sprintf("%v", err.Error()))
	}

	_, errors := form.Create()
	if len(errors) > 0 {
		for _, err := range errors {
			log.Info(fmt.Sprintf("%v", err))
		}
	}

	c.Redirect("/user/signup", http.StatusFound)
}
