package controllers

import (
	"net/http"

	form "go-crawler-challenge/forms/scrapper"
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
	c.Mapping("Create", c.Create)
}

// actionPolicyMapping maps search controller actions to policies
func (c *ScraperController) actionPolicyMapping() {
	c.MappingPolicy("Create", Policy{requireAuthenticatedUser: true})
}

// Create handles keyword for scrapping
// @Title Create
// @Description create a new scrapping result
// @Success 302 redirect to the dashboard page
// @Failure 302 redirect to the dashboard page and show an error message
// @router / [post]
func (c *ScraperController) Create() {
	flash := web.NewFlash()
	searchKeywordForm := form.SearchKeywordForm{}
	redirectPath := "/dashboard"

	err := c.ParseForm(&searchKeywordForm)
	if err != nil {
		flash.Error(err.Error())
	}

	errors := searchKeywordForm.Validate()
	if len(errors) > 0 {
		flash.Error(errors[0].Error())
	} else {
		searchKeyword := scraper.SearchKeywordService{Keyword: searchKeywordForm.Keyword}
		searchKeyword.Run()

		flash.Success("Scraping a keyword :)")
	}

	flash.Store(&c.Controller)
	c.Redirect(redirectPath, http.StatusFound)
}
