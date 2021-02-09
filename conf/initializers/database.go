package initializers

import (
	"fmt"

	orm "github.com/beego/beego/v2/client/orm"
	log "github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
)

// SetUpDatabase : Set up database connection with Postgres driver
func SetUpDatabase() {
	dbURL, err := web.AppConfig.String("dburl")
	if err != nil {
		log.Critical(fmt.Sprintf("Database URL not found: %v", err))
	}

	err = orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		log.Critical(fmt.Sprintf("Postgres Driver registration failed: %v", err))
	}

	err = orm.RegisterDataBase("default", "postgres", dbURL)
	if err != nil {
		log.Critical(fmt.Sprintf("Database Registration failed: %v", err))
	}

	if web.AppConfig.DefaultString("runmode", "dev") == "prod" {
		orm.Debug = false
	} else {
		orm.Debug = true
	}
}
