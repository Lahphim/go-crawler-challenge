package apiv1controllers

import (
	"net/http"

	. "go-crawler-challenge/controllers/api"
	"go-crawler-challenge/models"
	v1serializers "go-crawler-challenge/serializers/v1"
	service "go-crawler-challenge/services/keyword"
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
	c.MappingPolicy("Show", Policy{RequireAuthenticatedUser: true})
}

// Show handles report detail
// @Title Show
// @Description show the search result of the given keyword stored in the database
// @Success 202 {object} v1serializers.ReportDetail
// @Param :keyword_id path string true
// @Failure 500 Internal Server Error
// @Accept json
// @router /api/v1/report/:keyword_id [post]
func (c *ReportController) Show() {
	keywordId := c.GetString(":keyword_id")
	query := map[string]interface{}{
		"id":      keywordId,
		"user_id": c.CurrentUser.Id,
		"status":  models.GetKeywordStatus("completed"),
	}

	keyword, err := models.GetKeywordBy(query, []string{})
	if err != nil {
		c.RenderGenericError(ErrorNotFoundReport)

		return
	}

	reportGeneratorService := service.ReportGenerator{Keyword: keyword}
	reportResult, err := reportGeneratorService.Generate()
	if err != nil {
		c.RenderGenericError(ErrorGenerateReportFailed)

		return
	}

	serializer := v1serializers.ReportDetail{
		Report: reportResult.(*models.Report),
	}

	c.RenderJSON(serializer.Data(), http.StatusOK)
}
