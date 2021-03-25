package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	form "go-crawler-challenge/forms/keyword"
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
	c.Mapping("TextSearch", c.TextSearch)
	c.Mapping("FileSearch", c.FileSearch)
}

// actionPolicyMapping maps dashboard controller actions to policies
func (c *DashboardController) actionPolicyMapping() {
	c.MappingPolicy("Index", Policy{requireAuthenticatedUser: true})
	c.MappingPolicy("TextSearch", Policy{requireAuthenticatedUser: true})
	c.MappingPolicy("FileSearch", Policy{requireAuthenticatedUser: true})
}

// Index handles dashboard with handy widgets
// @Title Index
// @Description show some widgets such as search, listing and summary detail
// @Success 200
// @router /dashboard [get]
func (c *DashboardController) Index() {
	web.ReadFromRequest(&c.Controller)

	keyword := c.GetString("keyword")
	queryList := map[string]interface{}{
		"user_id":            c.CurrentUser.Id,
		"keyword__icontains": keyword,
	}

	totalRows, err := models.CountAllKeyword(queryList)
	if err != nil {
		logs.Critical(fmt.Sprintf("Get total rows failed: %v", err.Error()))
		c.Data["RetrieveKeywordFailed"] = "There was a problem retrieving all keywords :("
	} else {
		orderByList := c.GetOrderBy()
		pageSize := c.GetPageSize()
		paginator := pagination.SetPaginator((*context.Context)(c.Ctx), pageSize, totalRows)

		keywords, err := models.GetAllKeyword(queryList, orderByList, int64(paginator.Offset()), int64(pageSize))
		if err != nil {
			logs.Critical(fmt.Sprintf("Get all keyword failed: %v", err.Error()))
			c.Data["RetrieveKeywordFailed"] = "There was a problem retrieving all keywords :("
		} else {
			c.Data["Keywords"] = keywords
		}
	}

	c.Data["Keyword"] = keyword
	c.Data["XSRFForm"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layouts/application.html"
	c.TplName = "dashboard/index.html"
}

// TextSearch handles keyword for scrapping
// @Title TextSearch
// @Description create a new scrapping result by plain text
// @Success 302 redirect to the dashboard page
// @Failure 302 redirect to the dashboard page and show an error message
// @router /dashboard/search [post]
func (c *DashboardController) TextSearch() {
	flash := web.NewFlash()
	textSearchForm := form.TextSearchForm{}
	redirectPath := "/dashboard"

	err := c.ParseForm(&textSearchForm)
	if err != nil {
		flash.Error(err.Error())
	}

	textSearchForm.User = c.CurrentUser
	err = textSearchForm.Create()
	if err != nil {
		flash.Error(err.Error())
	} else {
		flash.Success("Scraping a keyword :)")
	}

	flash.Store(&c.Controller)
	c.Redirect(redirectPath, http.StatusFound)
}

// FileSearch handles keyword for scrapping
// @Title FileSearch
// @Description create a new scrapping result by CSV file
// @Success 302 redirect to the dashboard page
// @Failure 302 redirect to the dashboard page and show an error message
// @router /dashboard/upload [post]
func (c *DashboardController) FileSearch() {
	flash := web.NewFlash()
	redirectPath := "/dashboard"

	file, fileHeader, err := c.GetFile("file")
	if err != nil {
		flash.Error("The specified file could not be uploaded :(")
	} else {
		fileForm := form.FileSearchForm{File: file, FileHeader: fileHeader, User: c.CurrentUser}
		err = fileForm.Save()
		if err != nil {
			flash.Error(err.Error())
		} else {
			flash.Success("Scraping all keywords :)")
		}
	}

	flash.Store(&c.Controller)
	c.Redirect(redirectPath, http.StatusFound)
}
