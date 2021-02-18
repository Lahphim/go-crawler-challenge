package initializers

import (
	"fmt"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/beego/beego/v2/server/web/session/postgres"
)

// SetUpSession : Set up session with Postgres
func SetUpSession() {
	runMode := web.AppConfig.DefaultString("runmode", "dev")

	dbURL, err := web.AppConfig.String("dburl")
	if err != nil {
		logs.Critical(fmt.Sprintf("Database URL not found: %v", err))
	}

	sessionOn, err := web.AppConfig.Bool("SessionOn")
	if err != nil {
		logs.Critical(fmt.Sprintf("Session flag not found: %v", err))
	}

	if runMode != "test" && sessionOn {
		web.BConfig.WebConfig.Session.SessionProvider = "postgresql"
		web.BConfig.WebConfig.Session.SessionProviderConfig = dbURL
	}
}
