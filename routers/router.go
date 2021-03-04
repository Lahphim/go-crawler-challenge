package routers

import (
	"go-crawler-challenge/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/", &controllers.MainController{}, "get:Index")

	// Dashboard
	web.Router("/dashboard", &controllers.DashboardController{}, "get:Index")

	// Scraper
	web.Router("/dashboard/scraper/keyword", &controllers.ScraperController{}, "post:TextSearch")
	web.Router("/dashboard/scraper/keyword_csv", &controllers.ScraperController{}, "post:FileSearch")

	// User management
	web.Router("/user/sign_up", &controllers.UserController{}, "get:New")
	web.Router("/user/create", &controllers.UserController{}, "post:Create")

	// Session management
	web.Router("/user/sign_in", &controllers.SessionController{}, "get:New")
	web.Router("/user/sign_out", &controllers.SessionController{}, "get:Delete")
	web.Router("/user/session_create", &controllers.SessionController{}, "post:Create")
}
