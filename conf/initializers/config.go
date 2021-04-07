package initializers

import (
	"fmt"
	"log"

	"go-crawler-challenge/helpers"

	"github.com/beego/beego/v2/server/web"
)

func LoadAppConfig() {
	configPath := fmt.Sprintf("%s/conf/app.conf", helpers.RootDir())
	err := web.LoadAppConfig("ini", configPath)

	if err != nil {
		log.Fatal(fmt.Sprintf("Load configuration failed: %v", err))
	}
}
