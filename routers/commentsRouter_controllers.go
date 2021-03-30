package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["go-crawler-challenge/controllers/api/v1:TokenController"] = append(beego.GlobalControllerRouter["go-crawler-challenge/controllers/api/v1:TokenController"],
		beego.ControllerComments{
			Method:           "Create",
			Router:           "/api/v1/oauth/token",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["go-crawler-challenge/controllers:DashboardController"] = append(beego.GlobalControllerRouter["go-crawler-challenge/controllers:DashboardController"],
		beego.ControllerComments{
			Method:           "Index",
			Router:           "/dashboard",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["go-crawler-challenge/controllers:DashboardController"] = append(beego.GlobalControllerRouter["go-crawler-challenge/controllers:DashboardController"],
		beego.ControllerComments{
			Method:           "TextSearch",
			Router:           "/dashboard/search",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["go-crawler-challenge/controllers:DashboardController"] = append(beego.GlobalControllerRouter["go-crawler-challenge/controllers:DashboardController"],
		beego.ControllerComments{
			Method:           "FileSearch",
			Router:           "/dashboard/upload",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["go-crawler-challenge/controllers:MainController"] = append(beego.GlobalControllerRouter["go-crawler-challenge/controllers:MainController"],
		beego.ControllerComments{
			Method:           "Index",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["go-crawler-challenge/controllers:OauthClientController"] = append(beego.GlobalControllerRouter["go-crawler-challenge/controllers:OauthClientController"],
		beego.ControllerComments{
			Method:           "New",
			Router:           "/oauth_client",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["go-crawler-challenge/controllers:OauthClientController"] = append(beego.GlobalControllerRouter["go-crawler-challenge/controllers:OauthClientController"],
		beego.ControllerComments{
			Method:           "Create",
			Router:           "/oauth_client",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["go-crawler-challenge/controllers:ReportController"] = append(beego.GlobalControllerRouter["go-crawler-challenge/controllers:ReportController"],
		beego.ControllerComments{
			Method:           "Show",
			Router:           "/report/:keyword_id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["go-crawler-challenge/controllers:SessionController"] = append(beego.GlobalControllerRouter["go-crawler-challenge/controllers:SessionController"],
		beego.ControllerComments{
			Method:           "Create",
			Router:           "/user/session",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["go-crawler-challenge/controllers:SessionController"] = append(beego.GlobalControllerRouter["go-crawler-challenge/controllers:SessionController"],
		beego.ControllerComments{
			Method:           "New",
			Router:           "/user/sign_in",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["go-crawler-challenge/controllers:SessionController"] = append(beego.GlobalControllerRouter["go-crawler-challenge/controllers:SessionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/user/sign_out",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["go-crawler-challenge/controllers:UserController"] = append(beego.GlobalControllerRouter["go-crawler-challenge/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Create",
			Router:           "/user/create",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["go-crawler-challenge/controllers:UserController"] = append(beego.GlobalControllerRouter["go-crawler-challenge/controllers:UserController"],
		beego.ControllerComments{
			Method:           "New",
			Router:           "/user/sign_up",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
