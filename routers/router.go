package routers

import (
	"go-crawler-challenge/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/", &controllers.MainController{})

	web.Router("/user/signup", &controllers.UserController{}, "get:New")
	web.Router("/user/create", &controllers.UserController{}, "post:Create")
}
