package initializers

import "github.com/beego/beego/v2/server/web"

func init() {
	web.SetStaticPath("/assets", "assets")
}
