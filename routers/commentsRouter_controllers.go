package routers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	web.GlobalControllerRouter["go-crawler-challenge/controllers:UserController"] = append(web.GlobalControllerRouter["go-crawler-challenge/controllers:UserController"],
		web.ControllerComments{
			Method:           "New",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	web.GlobalControllerRouter["go-crawler-challenge/controllers:UserController"] = append(web.GlobalControllerRouter["go-crawler-challenge/controllers:UserController"],
		web.ControllerComments{
			Method:           "Create",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
