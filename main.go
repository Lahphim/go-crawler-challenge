package main

import (
	_ "go-crawler-challenge/conf/initializers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}
