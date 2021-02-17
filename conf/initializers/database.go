package initializers

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
)

// SetUpDatabase : Set up database connection with Postgres driver
func SetUpDatabase() {
	runMode := web.AppConfig.DefaultString("runmode", "dev")
	if runMode == "dev" {
		orm.Debug = true
	} else {
		orm.Debug = false
	}

	dbURL, err := web.AppConfig.String("dburl")
	if err != nil {
		logs.Critical(fmt.Sprintf("Database URL not found: %v", err))
	}

	err = orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		logs.Critical(fmt.Sprintf("Postgres Driver registration failed: %v", err))
	}

	err = orm.RegisterDataBase("default", "postgres", dbURL)
	if err != nil {
		logs.Critical(fmt.Sprintf("Database Registration failed: %v", err))
	} else {
		if runMode != "test" {
			SetUpSession(dbURL)
		}
	}

	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		logs.Critical(fmt.Sprintf("Sync the database failed: %v", err))
	}
}
