package apiv1controllers

import (
	"net/http"

	. "go-crawler-challenge/controllers/api"
	form "go-crawler-challenge/forms/keyword"
	v1serializers "go-crawler-challenge/serializers/v1"
)

// KeywordController operations for Keyword
type KeywordController struct {
	BaseController
}

// NestPrepare prepares some configurations to the controller
func (c *KeywordController) NestPrepare() {
	c.actionPolicyMapping()
}

// URLMapping maps keyword controller actions to functions
func (c *KeywordController) URLMapping() {
	c.Mapping("TextSearch", c.TextSearch)
}

// actionPolicyMapping maps keyword controller actions to policies
func (c *KeywordController) actionPolicyMapping() {
	c.MappingPolicy("TextSearch", Policy{RequireAuthenticatedUser: true})
}

// TextSearch handles keyword for scrapping
// @Title TextSearch
// @Description create a new scrapping result by plain text
// @Success 201 {object} v1serializers.KeywordScraper
// @Param keyword formData string true
// @Failure 500 Internal Server Error
// @Accept json
// @router /api/v1/keyword/search [post]
func (c *KeywordController) TextSearch() {
	textSearchForm := form.TextSearchForm{}

	err := c.ParseForm(&textSearchForm)
	if err != nil {
		return
	}

	textSearchForm.User = c.CurrentUser
	err = textSearchForm.Create()
	if err != nil {
		c.RenderGenericError(err)

		return
	}

	serializer := v1serializers.KeywordScraper{Message: "Scraping a keyword :)"}

	c.RenderJSON(serializer.Data(), http.StatusCreated)
}
