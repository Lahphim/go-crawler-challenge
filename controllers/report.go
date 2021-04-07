package controllers

import (
	"net/http"

	"go-crawler-challenge/models"
	service "go-crawler-challenge/services/keyword"

	"github.com/beego/beego/v2/server/web"
)

// ReportController operations for Report
type ReportController struct {
	BaseController
}

// NestPrepare prepares some configurations to the controller
func (c *ReportController) NestPrepare() {
	c.actionPolicyMapping()
}

// URLMapping maps report controller actions to functions
func (c *ReportController) URLMapping() {
	c.Mapping("Show", c.Show)
}

// actionPolicyMapping maps report controller actions to policies
func (c *ReportController) actionPolicyMapping() {
	c.MappingPolicy("Show", Policy{requireAuthenticatedUser: true})
}

// Show handles report page
// @Title Show
// @Description show the search result of the given keyword stored in the database
// @Success 200
// @Failure 302 redirect to the dashboard page and show an error message
// @router /report/:keyword_id [get]
func (c *ReportController) Show() {
	flash := web.NewFlash()
	hasReport := false
	redirectPath := "/dashboard"

	keywordId := c.GetString(":keyword_id")
	query := map[string]interface{}{
		"id":      keywordId,
		"user_id": c.CurrentUser.Id,
		"status":  models.GetKeywordStatus("completed"),
	}

	keyword, err := models.GetKeywordBy(query, []string{})
	if err == nil {
		reportGeneratorService := service.ReportGenerator{Keyword: keyword}
		reportResult, err := reportGeneratorService.Generate()
		if err == nil {
			hasReport = true
			c.Data["ReportResult"] = reportResult
		}
	}

	if hasReport {
		c.TplName = "report/index.html"
	} else {
		flash.Error("Report not found!")
		flash.Store(&c.Controller)
		c.Redirect(redirectPath, http.StatusFound)
	}
}
