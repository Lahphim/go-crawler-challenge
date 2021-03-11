package routers

import (
	"go-crawler-challenge/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/", &controllers.MainController{}, "get:Index")

	// Dashboard
	web.Router("/dashboard", &controllers.DashboardController{}, "get:Index")
	web.Router("/dashboard/search", &controllers.DashboardController{}, "post:TextSearch")
	web.Router("/dashboard/upload", &controllers.DashboardController{}, "post:FileSearch")

	// User management
	web.Router("/user/sign_up", &controllers.UserController{}, "get:New")
	web.Router("/user/create", &controllers.UserController{}, "post:Create")

	// Session management
	web.Router("/user/sign_in", &controllers.SessionController{}, "get:New")
	web.Router("/user/sign_out", &controllers.SessionController{}, "get:Delete")
	web.Router("/user/session", &controllers.SessionController{}, "post:Create")
}
