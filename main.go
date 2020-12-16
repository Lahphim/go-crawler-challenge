package main

import (
	_ "go-crawler-challenge/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

