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

type NestPreparer interface {
	NestPrepare()
}

//  BaseController operations for all controller
type BaseController struct {
	web.Controller

	CurrentUser              *models.User
	requireAuthenticatedUser bool
	requireGuestUser         bool
}

func (c *BaseController) Prepare() {
	helpers.SetControllerAttributes(&c.Controller)

	c.handleAuthorizeRequest()
	c.setCurrentUserData()
	c.setCustomLayout()
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
	if c.requireGuestUser && !c.ensureGuestUser() {
		c.Redirect("/", http.StatusFound)
	}

	if c.requireAuthenticatedUser && !c.ensureAuthenticatedUser() {
		c.SetSessionCurrentUser(nil)

		c.Redirect("/user/sign_in", http.StatusFound)
	}
}

func (c *BaseController) ensureAuthenticatedUser() bool {
	currentUser := c.GetSessionCurrentUser()

	return currentUser != nil
}

func (c *BaseController) ensureGuestUser() bool {
	currentUser := c.GetSessionCurrentUser()

	return currentUser == nil
}

func (c *BaseController) setCurrentUserData() {
	currentUser := c.GetSessionCurrentUser()

	c.Data["CurrentUser"] = currentUser
	c.CurrentUser = currentUser
}

func (c *BaseController) setCustomLayout() {
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["FlashMessage"] = "shared/alert.html"
	c.LayoutSections["HeaderContent"] = "shared/header.html"

	app, ok := c.AppController.(NestPreparer)
	if ok {
		app.NestPrepare()
	}
}
