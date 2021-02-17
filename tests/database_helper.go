package tests

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

func TruncateTable(tableName string) {
	ormer := orm.NewOrm()
	rawSql := fmt.Sprintf("TRUNCATE TABLE \"%s\";", tableName)

	_, err := ormer.Raw(rawSql).Exec()
	if err != nil {
		err := orm.RunSyncdb("default", true, false)
		if err != nil {
			logs.Critical(fmt.Sprintf("Sync the database failed: %v", err))
		}
	}
}
