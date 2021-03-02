package tests

import (
	"fmt"

	. "go-crawler-challenge/tests/fixtures"

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

func PreparePositionTable() {
	FabricatePosition("nonAds", "#search .g .yuRUbf > a", "normal")
	FabricatePosition("bottomLinkAds", "#tadsb .d5oMvf > a", "other")
	FabricatePosition("otherAds", "#rhs .pla-unit a.pla-unit-title-link", "other")
	FabricatePosition("topImageAds", "#tvcap .pla-unit a.pla-unit-title-link", "top")
	FabricatePosition("topLinkAds", "#tads .d5oMvf > a", "top")
}
