package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["go-crawler-challenge/controllers:UserController"] = append(beego.GlobalControllerRouter["go-crawler-challenge/controllers:UserController"],
        beego.ControllerComments{
            Method: "Signup",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
