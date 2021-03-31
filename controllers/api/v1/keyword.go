package apiv1controllers

import (
	. "go-crawler-challenge/controllers/api"
	form "go-crawler-challenge/forms/keyword"

	"github.com/beego/beego/v2/core/logs"

	"github.com/go-oauth2/oauth2/v4/errors"
)

type KeywordController struct {
	BaseController
}

func (c *KeywordController) NestPrepare() {
	c.actionPolicyMapping()
}

func (c *KeywordController) URLMapping() {
	c.Mapping("TextSearch", c.TextSearch)
}

func (c *KeywordController) actionPolicyMapping() {
	//c.MappingPolicy("TextSearch", Policy{RequireAuthenticatedUser: true})
}

func (c *KeywordController) TextSearch() {
	textSearchForm := form.TextSearchForm{}

	err := c.ParseForm(&textSearchForm)
	if err != nil {
		return
	}

	//err = textSearchForm.Create()
	//if err != nil {
	//	return
	//} else {
	//	return
	//}

	type mystruct struct {
		FieldOne string `json:"field_one"`
	}

	_ = mystruct{FieldOne: textSearchForm.Keyword}

	err = c.RenderGenericError(errors.New("esjklfjsdklfjdksl"))
	if err != nil {
		logs.Error("Generic error: ", err.Error())
	}

	//response := mystruct{FieldOne: textSearchForm.Keyword}
	//
	//
	//
	//c.Ctx.Output.Header("Content-Type", ContentType)
	//c.Ctx.Output.Status = 500
	//c.Data["json"] = &response
	//c.ServeJSON()
}
