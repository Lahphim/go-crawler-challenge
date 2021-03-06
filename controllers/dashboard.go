package controllers

import (
	"fmt"
	"html/template"

	"go-crawler-challenge/models"

	"github.com/beego/beego/v2/adapter/context"
	"github.com/beego/beego/v2/adapter/utils/pagination"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

// DashboardController operations for Dashboard
type DashboardController struct {
	BaseController
}

// NestPrepare prepares some configurations to the controller
func (c *DashboardController) NestPrepare() {
	c.actionPolicyMapping()
}

// URLMapping maps dashboard controller actions to functions
func (c *DashboardController) URLMapping() {
	c.Mapping("Index", c.Index)
}

// actionPolicyMapping maps dashboard controller actions to policies
func (c *DashboardController) actionPolicyMapping() {
	c.MappingPolicy("Index", Policy{requireAuthenticatedUser: true})
}

// Index handles dashboard with handy widgets
// @Title Index
// @Description show some widgets such as search, listing and summary detail
// @Success 200
// @router / [get]
func (c *DashboardController) Index() {
	web.ReadFromRequest(&c.Controller)

	var keywords []*models.Keyword
	var orderByList = []string{"created_at desc"}

	pageSize := 10
	totalRows, err := models.CountAllKeyword()
	if err != nil {
		logs.Critical(fmt.Sprintf("Get total rows failed: %v", err.Error()))
	} else {
		paginator := pagination.SetPaginator((*context.Context)(c.Ctx), pageSize, totalRows)

		keywords, err = models.GetAllKeyword(orderByList, int64(paginator.Offset()), int64(pageSize))
		if err != nil {
			logs.Critical(fmt.Sprintf("Get all keyword failed: %v", err.Error()))
		}
	}

	c.Data["Keywords"] = keywords
	c.Data["XSRFData"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layouts/application.html"
	c.TplName = "dashboard/index.html"
}
