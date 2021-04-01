package apiv1controllers

import (
	"go-crawler-challenge/models"
	"net/http"

	"github.com/beego/beego/v2/adapter/utils/pagination"

	. "go-crawler-challenge/controllers/api"
	form "go-crawler-challenge/forms/keyword"
	v1serializers "go-crawler-challenge/serializers/v1"

	"github.com/beego/beego/v2/adapter/context"
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
	c.Mapping("Index", c.Index)
	c.Mapping("TextSearch", c.TextSearch)
}

// actionPolicyMapping maps keyword controller actions to policies
func (c *KeywordController) actionPolicyMapping() {
	c.MappingPolicy("Index", Policy{RequireAuthenticatedUser: true})
	c.MappingPolicy("TextSearch", Policy{RequireAuthenticatedUser: true})
}

func (c *KeywordController) Index() {
	keyword := c.GetString("keyword")
	queryList := map[string]interface{}{
		"user_id":            c.CurrentUser.Id,
		"keyword__icontains": keyword,
	}

	totalRows, err := models.CountAllKeyword(queryList)
	if err != nil {
		c.RenderGenericError(ErrorRetrieveKeywordFailed)

		return
	}

	orderByList := c.GetOrderBy()
	pageSize := c.GetPageSize()
	paginator := pagination.SetPaginator((*context.Context)(c.Ctx), pageSize, totalRows)

	keywords, err := models.GetAllKeyword(queryList, orderByList, int64(paginator.Offset()), int64(pageSize))
	if err != nil {
		c.RenderGenericError(ErrorRetrieveKeywordFailed)

		return
	}

	serializer := v1serializers.KeywordList{
		KeywordList: keywords,
		TotalRows:   int(totalRows),
		PageSize:    pageSize,
	}

	c.RenderJSON(serializer.Data(), http.StatusOK)
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
