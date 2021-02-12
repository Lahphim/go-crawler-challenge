package controllers

import (
	"fmt"

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
}

func (c *BaseController) SetCurrentUser(user *models.User) {
	err := c.SetSession(currentUserKey, user.Id)
	if err != nil {
		log.Critical(fmt.Sprintf("Set session failed: %v", err))
	}
}

func (c *BaseController) GetCurrentUser() (user *models.User) {
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
