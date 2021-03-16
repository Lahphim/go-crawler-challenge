package fixtures

import (
	"go-crawler-challenge/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/onsi/ginkgo"
)

func FabricateKeyword(keyword string, url string, user *models.User) (keywordRecord *models.Keyword) {
	keywordRecord = &models.Keyword{
		User:    user,
		Keyword: keyword,
		Url:     url,
	}

	ormer := orm.NewOrm()
	_, err := ormer.Insert(keywordRecord)
	if err != nil {
		ginkgo.Fail("Add keyword failed: " + err.Error())
	}

	return keywordRecord
}
