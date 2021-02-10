package controllers

import (
	"go-crawler-challenge/helpers"

	"github.com/beego/beego/v2/server/web"
)

//  BaseController operations for all controller
type BaseController struct {
	web.Controller
}

func (this *BaseController) Prepare() {
	helpers.SetControllerAttributes(&this.Controller)
}
