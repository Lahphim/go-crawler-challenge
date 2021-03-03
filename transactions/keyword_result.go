package transactions

import (
	"go-crawler-challenge/models"

	"github.com/beego/beego/v2/client/orm"
)

type KeywordResult struct {
	Keyword  string
	Url      string
	LinkList []models.Link
	RawHtml  string
	User     *models.User
}

// AddKeywordResult creates multiple records at the same time within transaction.
// Related table are `keyword`, `page` and `link`.
// If there are some errors in the middle of the process, it will rollbacks to the beginning.
func AddKeywordResult(keywordResult *KeywordResult) (keyword *models.Keyword, err error) {
	ormer := orm.NewOrm()

	// Transaction : Begin
	txnOrmer, err := ormer.Begin()
	if err != nil {
		return nil, err
	}

	// Transaction : Keyword
	keyword = &models.Keyword{
		User: keywordResult.User,

		Keyword: keywordResult.Keyword,
		Url:     keywordResult.Url,
	}
	_, err = txnOrmer.Insert(keyword)
	if err != nil {
		errRollback := txnOrmer.Rollback()
		if errRollback != nil {
			return nil, errRollback
		}

		return nil, err
	}

	// Transaction : Page
	page := &models.Page{
		Keyword: keyword,
		RawHtml: keywordResult.RawHtml,
	}
	_, err = txnOrmer.Insert(page)
	if err != nil {
		errRollback := txnOrmer.Rollback()
		if errRollback != nil {
			return nil, errRollback
		}

		return nil, err
	}

	// Transaction : Link
	if len(keywordResult.LinkList) > 0 {
		for index := range keywordResult.LinkList {
			keywordResult.LinkList[index].Keyword = keyword
		}
		_, err = txnOrmer.InsertMulti(50, keywordResult.LinkList)
		if err != nil {
			errRollback := txnOrmer.Rollback()
			if errRollback != nil {
				return nil, errRollback
			}

			return nil, err
		}
	}

	// Transaction : Commit
	err = txnOrmer.Commit()
	if err != nil {
		errRollback := txnOrmer.Rollback()
		if errRollback != nil {
			return nil, errRollback
		}

		return nil, err
	}

	return keyword, nil
}
