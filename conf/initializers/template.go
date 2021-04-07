package initializers

import (
	"fmt"

	. "go-crawler-challenge/helpers"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

func SetUpTemplate() {
	err := web.AddFuncMap("hashEmail", HashEmail)
	if err != nil {
		logs.Error(fmt.Sprintf("Map hashEmail function failed: %v", err))
	}

	err = web.AddFuncMap("toTimeAgo", ToTimeAgo)
	if err != nil {
		logs.Error(fmt.Sprintf("Map toTimeAgo function failed: %v", err))
	}

	err = web.AddFuncMap("toTimeStamp", ToTimeStamp)
	if err != nil {
		logs.Error(fmt.Sprintf("Map toTimeStamp function failed: %v", err))
	}

	err = web.AddFuncMap("unescape", Unescape)
	if err != nil {
		logs.Error(fmt.Sprintf("Map unescape function failed: %v", err))
	}
}
