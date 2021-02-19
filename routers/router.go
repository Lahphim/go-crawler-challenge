package routers

import (
	"go-crawler-challenge/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/", &controllers.MainController{}, "get:Index")

	// Dashboard
	web.Router("/dashboard", &controllers.DashboardController{}, "get:Index")

	// User management
	web.Router("/user/sign_up", &controllers.UserController{}, "get:New")
	web.Router("/user/create", &controllers.UserController{}, "post:Create")

	// Session management
	web.Router("/user/sign_in", &controllers.SessionController{}, "get:New")
	web.Router("/user/sign_out", &controllers.SessionController{}, "get:Delete")
	web.Router("/session/create", &controllers.SessionController{}, "post:Create")
}
