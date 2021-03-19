package fixtures

import (
	"go-crawler-challenge/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/onsi/ginkgo"
)

func FabricatePage(rawHtml string, keyword *models.Keyword) (page *models.Page) {
	page = &models.Page{
		RawHtml: rawHtml,
		Keyword: keyword,
	}

	ormer := orm.NewOrm()
	_, err := ormer.Insert(page)
	if err != nil {
		ginkgo.Fail("Add page failed: " + err.Error())
	}

	return page
}
