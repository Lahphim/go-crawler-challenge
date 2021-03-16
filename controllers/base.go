package controllers

import (
	"fmt"
	"net/http"

	"go-crawler-challenge/helpers"
	"go-crawler-challenge/models"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

const CurrentUserKey = "CURRENT_USER_ID"
const defaultPageSize = 10

var defaultOrderBy = []string{"created_at desc"}

type NestPreparer interface {
	NestPrepare()
}

//  BaseController operations for all controller
type BaseController struct {
	web.Controller

	CurrentUser  *models.User
	actionPolicy map[string]Policy
}

type Policy struct {
	requireAuthenticatedUser bool
	requireGuestUser         bool
}

func (c *BaseController) Prepare() {
	helpers.SetControllerAttributes(&c.Controller)
	c.applyCustomLayout()
	c.initActionPolicy()

	app, ok := c.AppController.(NestPreparer)
	if ok {
		app.NestPrepare()
	}

	c.handleAuthorizeRequest()
}

func (c *BaseController) GetPageSize() (pageSize int) {
	return defaultPageSize
}

func (c *BaseController) GetOrderBy() (orderBy []string) {
	return defaultOrderBy
}

func (c *BaseController) MappingPolicy(method string, policy Policy) {
	c.actionPolicy[method] = policy
}

func (c *BaseController) SetSessionCurrentUser(user *models.User) {
	if user != nil {
		err := c.SetSession(CurrentUserKey, user.Id)
		if err != nil {
			logs.Critical(fmt.Sprintf("Set session failed: %v", err))
		}
	} else {
		err := c.DelSession(CurrentUserKey)
		if err != nil {
			logs.Critical(fmt.Sprintf("Delete session failed: %v", err))
		}
	}
}

func (c *BaseController) GetSessionCurrentUser() (user *models.User) {
	userId := c.GetSession(CurrentUserKey)
	if userId == nil {
		return nil
	}

	user, err := models.GetUserById(userId.(int64))
	if err != nil {
		return nil
	}

	return user
}

func (c *BaseController) RevokeSessionCurrentUser() error {
	return c.DelSession(CurrentUserKey)
}

func (c *BaseController) handleAuthorizeRequest() {
	_, actionName := c.GetControllerAndAction()
	actionPolicy := c.actionPolicy[actionName]

	if actionPolicy.requireGuestUser && !c.ensureGuestUser() {
		c.Redirect("/dashboard", http.StatusFound)
	}

	if actionPolicy.requireAuthenticatedUser && !c.ensureAuthenticatedUser() {
		c.SetSessionCurrentUser(nil)

		c.Redirect("/user/sign_in", http.StatusFound)
	}

	c.assignCurrentUser()
}

func (c *BaseController) ensureAuthenticatedUser() bool {
	currentUser := c.GetSessionCurrentUser()

	return currentUser != nil
}

func (c *BaseController) ensureGuestUser() bool {
	currentUser := c.GetSessionCurrentUser()

	return currentUser == nil
}

func (c *BaseController) applyCustomLayout() {
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["FlashMessage"] = "shared/_alert.html"
	c.LayoutSections["HeaderContent"] = "shared/_header.html"
}

func (c *BaseController) initActionPolicy() {
	c.actionPolicy = make(map[string]Policy)
}

func (c *BaseController) assignCurrentUser() {
	currentUser := c.GetSessionCurrentUser()

	c.Data["CurrentUser"] = currentUser
	c.CurrentUser = currentUser
}
