package controllers

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	service "go-crawler-challenge/services/oauth"

	"github.com/beego/beego/v2/server/web"
)

// OauthClientController operations for oauth client generator
type OauthClientController struct {
	BaseController
}

// NestPrepare prepares some configurations to the controller
func (c *OauthClientController) NestPrepare() {
	c.actionPolicyMapping()
}

// URLMapping maps oauth client controller actions to functions
func (c *OauthClientController) URLMapping() {
	c.Mapping("New", c.New)
	c.Mapping("Create", c.Create)
}

// actionPolicyMapping maps oauth client controller actions to policies
func (c *OauthClientController) actionPolicyMapping() {
	c.MappingPolicy("New", Policy{requireAuthenticatedUser: true})
	c.MappingPolicy("Create", Policy{requireAuthenticatedUser: true})
}

// New handles a form for generating a new oauth client credential and only display into the same form
// @Title New
// @Description show an oauth client credential form with generated values
// @Success 200
// @router /oauth_client [get]
func (c *OauthClientController) New() {
	web.ReadFromRequest(&c.Controller)
	flash := web.NewFlash()

	clientId := c.GetString("client_id")
	if clientId != "" {
		oauthClient, err := service.ClientStore.GetByID(context.TODO(), clientId)
		if err != nil {
			flash.Error("OAuth client credential not found!")
		} else {
			c.Data["ClientId"] = oauthClient.GetID()
			c.Data["ClientSecret"] = oauthClient.GetSecret()
		}
	}

	flash.Store(&c.Controller)
	c.Data["XSRFForm"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layouts/application.html"
	c.TplName = "oauth_client/new.html"
}

// Create handles create an oauth client credential
// @Title Create
// @Description create an oauth client credential
// @Success 302 redirect to the oauth client form path with success message
// @Success 302 redirect to the oauth client form path with error message
// @router /oauth_client [post]
func (c *OauthClientController) Create() {
	flash := web.NewFlash()
	redirectPath := "/oauth_client"

	domain := c.Ctx.Request.Host
	serviceOauth := service.ClientGenerator{Domain: domain}
	clientId, err := serviceOauth.Generate()
	if err != nil {
		flash.Error(err.Error())
	} else {
		flash.Success("You have successfully created a new client")
		redirectPath = fmt.Sprintf("%v?client_id=%v", redirectPath, clientId)
	}
	flash.Success("You have successfully created a new client")

	flash.Store(&c.Controller)
	c.Redirect(redirectPath, http.StatusFound)
}
