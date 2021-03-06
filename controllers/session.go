package controllers

import (
	"html/template"
	"net/http"
	"net/url"

	form "go-crawler-challenge/forms/session"

	"github.com/beego/beego/v2/server/web"
)

// SessionController operations for User's session
type SessionController struct {
	BaseController
}

// NestPrepare prepares some configurations to the controller
func (c *SessionController) NestPrepare() {
	c.actionPolicyMapping()
}

// URLMapping maps session controller actions to functions
func (c *SessionController) URLMapping() {
	c.Mapping("New", c.New)
	c.Mapping("Create", c.Create)
	c.Mapping("Delete", c.Delete)
}

// actionPolicyMapping maps session controller actions to policies
func (c *SessionController) actionPolicyMapping() {
	c.MappingPolicy("New", Policy{requireGuestUser: true})
	c.MappingPolicy("Create", Policy{requireGuestUser: true})
	c.MappingPolicy("Delete", Policy{requireAuthenticatedUser: true})
}

// New handles a form for signing in
// @Title New
// @Description show a sign-in form
// @Success 200
// @router /user/sign_in [get]
func (c *SessionController) New() {
	web.ReadFromRequest(&c.Controller)

	c.Data["XSRFForm"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layouts/authentication.html"
	c.TplName = "session/new.html"
}

// Create handles create a session for an authenticated user
// @Title Create
// @Description create a session
// @Success 302 redirect to root path with success message
// @Failure 302 redirect to sign-in path with error message
// @router /user/session [post]
func (c *SessionController) Create() {
	flash := web.NewFlash()
	authenticationForm := form.AuthenticationForm{}
	redirectPath := "/user/sign_in"

	err := c.ParseForm(&authenticationForm)
	if err != nil {
		flash.Error(err.Error())
	}

	user, errors := authenticationForm.Authenticate()
	if len(errors) > 0 {
		flash.Error(errors[0].Error())
	} else {
		c.SetSessionCurrentUser(user)

		flash.Success("You have successfully signed in")
		redirectPath = "/dashboard"
	}

	flash.Store(&c.Controller)
	c.Redirect(redirectPath, http.StatusFound)
}

// Delete handles revoke a session for an authenticated user
// @Title Delete
// @Description revoke a session
// @Success 302 redirect to root path with success message
// @Failure 302 redirect to current path with error message
// @router /user/sign_out [get]
func (c *SessionController) Delete() {
	flash := web.NewFlash()
	redirectPath := "/"

	u, err := url.Parse(c.Ctx.Request.Header.Get("Referer"))
	if err == nil {
		redirectPath = u.Path
	}

	err = c.RevokeSessionCurrentUser()
	if err != nil {
		flash.Error("Sign out failed")
	} else {
		flash.Success("You have successfully signed out")
		redirectPath = "/"
	}

	flash.Store(&c.Controller)
	c.Redirect(redirectPath, http.StatusFound)
}
