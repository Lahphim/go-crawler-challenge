package main

import (
	_ "go-crawler-challenge/conf/initializers"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	web.Run()
}
