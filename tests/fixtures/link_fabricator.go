package fixtures

import (
	"go-crawler-challenge/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/onsi/ginkgo"
)

func FabricateLink(url string, keyword *models.Keyword, position *models.Position) (link *models.Link) {
	link = &models.Link{
		Url:      url,
		Keyword:  keyword,
		Position: position,
	}

	ormer := orm.NewOrm()
	_, err := ormer.Insert(link)
	if err != nil {
		ginkgo.Fail("Add link failed: " + err.Error())
	}

	return link
}
