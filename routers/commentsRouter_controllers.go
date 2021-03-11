package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["go-crawler-challenge/controllers:DashboardController"] = append(beego.GlobalControllerRouter["go-crawler-challenge/controllers:DashboardController"],
		beego.ControllerComments{
			Method:           "Index",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["go-crawler-challenge/controllers:DashboardController"] = append(beego.GlobalControllerRouter["go-crawler-challenge/controllers:DashboardController"],
		beego.ControllerComments{
			Method:           "TextSearch",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["go-crawler-challenge/controllers:DashboardController"] = append(beego.GlobalControllerRouter["go-crawler-challenge/controllers:DashboardController"],
		beego.ControllerComments{
			Method:           "FileSearch",
			Router:           "/",
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

	beego.GlobalControllerRouter["go-crawler-challenge/controllers:SessionController"] = append(beego.GlobalControllerRouter["go-crawler-challenge/controllers:SessionController"],
		beego.ControllerComments{
			Method:           "New",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["go-crawler-challenge/controllers:SessionController"] = append(beego.GlobalControllerRouter["go-crawler-challenge/controllers:SessionController"],
		beego.ControllerComments{
			Method:           "Create",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["go-crawler-challenge/controllers:SessionController"] = append(beego.GlobalControllerRouter["go-crawler-challenge/controllers:SessionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["go-crawler-challenge/controllers:UserController"] = append(beego.GlobalControllerRouter["go-crawler-challenge/controllers:UserController"],
		beego.ControllerComments{
			Method:           "New",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["go-crawler-challenge/controllers:UserController"] = append(beego.GlobalControllerRouter["go-crawler-challenge/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Create",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
