package controllers

import (
	"net/http"

	form "go-crawler-challenge/forms/scraper"
	"go-crawler-challenge/models"
	"go-crawler-challenge/services/scraper"

	"github.com/beego/beego/v2/server/web"
)

// ScraperController operations for Dashboard
type ScraperController struct {
	BaseController
}

// NestPrepare prepares some configurations to the controller
func (c *ScraperController) NestPrepare() {
	c.actionPolicyMapping()
}

// URLMapping maps search controller actions to functions
func (c *ScraperController) URLMapping() {
	c.Mapping("TextSearch", c.TextSearch)
	c.Mapping("FileSearch", c.FileSearch)
}

// actionPolicyMapping maps search controller actions to policies
func (c *ScraperController) actionPolicyMapping() {
	c.MappingPolicy("TextSearch", Policy{requireAuthenticatedUser: true})
	c.MappingPolicy("FileSearch", Policy{requireAuthenticatedUser: true})
}

// TextSearch handles keyword for scrapping
// @Title TextSearch
// @Description create a new scrapping result by plain text
// @Success 302 redirect to the dashboard page
// @Failure 302 redirect to the dashboard page and show an error message
// @router / [post]
func (c *ScraperController) TextSearch() {
	flash := web.NewFlash()
	textKeywordForm := form.TextKeywordForm{}
	redirectPath := "/dashboard"

	err := c.ParseForm(&textKeywordForm)
	if err != nil {
		flash.Error(err.Error())
	}

	errors := textKeywordForm.Validate()
	if len(errors) > 0 {
		flash.Error(errors[0].Error())
	} else {
		positionList, err := models.GetAllPosition()
		if err != nil {
			flash.Error(err.Error())
		} else {
			searchKeyword := scraper.SearchKeywordService{User: c.CurrentUser, Keyword: textKeywordForm.Keyword}
			searchKeyword.SetPositionList(positionList)
			searchKeyword.Run()

			flash.Success("Scraping a keyword :)")
		}
	}

	flash.Store(&c.Controller)
	c.Redirect(redirectPath, http.StatusFound)
}

// FileSearch handles keyword for scrapping
// @Title FileSearch
// @Description create a new scrapping result by CSV file
// @Success 302 redirect to the dashboard page
// @Failure 302 redirect to the dashboard page and show an error message
// @router / [post]
func (c *ScraperController) FileSearch() {
	flash := web.NewFlash()
	redirectPath := "/dashboard"

	flash.Success("Scraping all keywords :)")

	flash.Store(&c.Controller)
	c.Redirect(redirectPath, http.StatusFound)
}
