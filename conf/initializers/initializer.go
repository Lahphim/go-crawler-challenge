package initializers

import (
	_ "go-crawler-challenge/models"
	_ "go-crawler-challenge/routers"

	_ "github.com/beego/beego/v2/server/web/session/postgres"
)

func init() {
	LoadAppConfig()
	SetUpDatabase()
	SetUpTask()
	SetUpTemplate()
}
