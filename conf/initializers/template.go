package initializers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/justincampbell/timeago"
)

func SetUpTemplate() {
	err := web.AddFuncMap("hashEmail", hashEmail)
	if err != nil {
		logs.Error(fmt.Sprintf("Map hashEmail function failed: %v", err))
	}

	err = web.AddFuncMap("toTimeAgo", toTimeAgo)
	if err != nil {
		logs.Error(fmt.Sprintf("TimeAgo function failed: %v", err))
	}
}

func hashEmail(plainEmail string) string {
	byteEmail := md5.Sum([]byte(plainEmail))

	return hex.EncodeToString(byteEmail[:])
}

func toTimeAgo(timestamp time.Time) string {
	return timeago.FromTime(timestamp)
}
