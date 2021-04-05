package apiv1controllers

import (
	"net/http"

	. "go-crawler-challenge/controllers/api"
	form "go-crawler-challenge/forms/keyword"
	"go-crawler-challenge/models"
	v1serializers "go-crawler-challenge/serializers/v1"

	"github.com/beego/beego/v2/adapter/context"
	"github.com/beego/beego/v2/adapter/utils/pagination"
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
	c.Mapping("FileSearch", c.FileSearch)
}

// actionPolicyMapping maps keyword controller actions to policies
func (c *KeywordController) actionPolicyMapping() {
	c.MappingPolicy("Index", Policy{RequireAuthenticatedUser: true})
	c.MappingPolicy("TextSearch", Policy{RequireAuthenticatedUser: true})
	c.MappingPolicy("FileSearch", Policy{RequireAuthenticatedUser: true})
}

// Index handles keyword list
// @Title Index
// @Description response with keyword list filterable with keyword and page number
// @Success 200 {object} v1serializers.KeywordList
// @Param keyword	query string	false
// @Param p			query integer	false
// @Failure 500 Internal Server Error
// @Accept json
// @router /api/v1/keywords [get]
func (c *KeywordController) Index() {
	keyword := c.GetString("keyword")
	queryList := map[string]interface{}{
		"user_id":            c.CurrentUser.Id,
		"keyword__icontains": keyword,
	}

	totalRows, err := models.CountAllKeyword(queryList)
	if err != nil {
		c.RenderGenericError(ErrorRetrieveKeywordFailed)
	}

	orderByList := c.GetOrderBy()
	pageSize := c.GetPageSize()
	paginator := pagination.SetPaginator((*context.Context)(c.Ctx), pageSize, totalRows)

	keywords, err := models.GetAllKeyword(queryList, orderByList, int64(paginator.Offset()), int64(pageSize))
	if err != nil {
		c.RenderGenericError(ErrorRetrieveKeywordFailed)
	}

	serializer := v1serializers.KeywordList{
		KeywordList: keywords,
		Paginator:   paginator,
	}

	c.RenderJSONList(serializer.Data(), serializer.Meta(), serializer.Links(), http.StatusOK)
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
	}

	serializer := v1serializers.KeywordScraper{Message: "Scraping a keyword :)"}

	c.RenderJSON(serializer.Data(), http.StatusCreated)
}

// FileSearch handles keyword for scrapping
// @Title FileSearch
// @Description create new scrapping result by CSV file
// @Success 201
// @Param file formData file true
// @Failure 500 Internal Server Error
// @router /api/v1/keyword/upload [post]
func (c *KeywordController) FileSearch() {
	file, fileHeader, err := c.GetFile("file")
	if err != nil {
		c.RenderGenericError(ErrorUploadFileFailed)
	}

	fileForm := form.FileSearchForm{File: file, FileHeader: fileHeader, User: c.CurrentUser}
	err = fileForm.Save()
	if err != nil {
		c.RenderGenericError(err)
	}

	c.RenderJSON(nil, http.StatusNoContent)
}
