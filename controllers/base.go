package controllers

import (
	"fmt"
	"net/http"

	"go-crawler-challenge/helpers"
	"go-crawler-challenge/models"

	log "github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

const currentUserKey = "CURRENT_USER_ID"

//  BaseController operations for all controller
type BaseController struct {
	web.Controller

	CurrentUser              *models.User
	requireAuthenticatedUser bool
	requireGuestUser         bool
}

func (c *BaseController) Prepare() {
	helpers.SetControllerAttributes(&c.Controller)
	helpers.SetFlashMessageLayout(&c.Controller)

	c.handleAuthorizeRequest()
}

func (c *BaseController) SetSessionCurrentUser(user *models.User) {
	if user != nil {
		err := c.SetSession(currentUserKey, user.Id)
		if err != nil {
			log.Critical(fmt.Sprintf("Set session failed: %v", err))
		}
	} else {
		err := c.DelSession(currentUserKey)
		if err != nil {
			log.Critical(fmt.Sprintf("Delete session failed: %v", err))
		}
	}

	c.Data["CurrentUser"] = user
	c.CurrentUser = user
}

func (c *BaseController) GetSessionCurrentUser() (user *models.User) {
	userId := c.GetSession(currentUserKey)
	if userId == nil {
		return nil
	}

	user, err := models.GetUserById(userId.(int64))
	if err != nil {
		return nil
	}

	return user
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
