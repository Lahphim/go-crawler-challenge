package routers

import (
	"go-crawler-challenge/controllers"
	apiv1controllers "go-crawler-challenge/controllers/api/v1"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/", &controllers.MainController{}, "get:Index")

	// OAuth client
	web.Router("/oauth_client", &controllers.OauthClientController{}, "get:New;post:Create")

	// Dashboard
	web.Router("/dashboard", &controllers.DashboardController{}, "get:Index")
	web.Router("/dashboard/search", &controllers.DashboardController{}, "post:TextSearch")
	web.Router("/dashboard/upload", &controllers.DashboardController{}, "post:FileSearch")

	// Report
	web.Router("/report/:keyword_id", &controllers.ReportController{}, "get:Show")

	// User management
	web.Router("/user/sign_up", &controllers.UserController{}, "get:New")
	web.Router("/user/create", &controllers.UserController{}, "post:Create")

	// Session management
	web.Router("/user/sign_in", &controllers.SessionController{}, "get:New")
	web.Router("/user/sign_out", &controllers.SessionController{}, "get:Delete")
	web.Router("/user/session", &controllers.SessionController{}, "post:Create")

	// API
	// V1
	namespaceV1 := web.NewNamespace("/api/v1",
		// OAuth
		web.NSRouter("/oauth/token", &apiv1controllers.TokenController{}, "post:Create"),

		// Keyword
		web.NSRouter("/keywords", &apiv1controllers.KeywordController{}, "get:Index"),
		web.NSRouter("/keyword/search", &apiv1controllers.KeywordController{}, "post:TextSearch"),
		web.NSRouter("/keyword/upload", &apiv1controllers.KeywordController{}, "post:FileSearch"),

		// Report
		web.NSRouter("/report/:keyword_id", &apiv1controllers.ReportController{}, "get:Show"),
	)

	web.AddNamespace(namespaceV1)
}
