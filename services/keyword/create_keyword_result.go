package keyword

import (
	"fmt"
	"net/url"

	"go-crawler-challenge/models"
	. "go-crawler-challenge/services"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
)

type CreateKeywordResult struct {
	Keyword  *models.Keyword `valid:"Required;"`
	Url      string          `valid:"Required; MaxSize(128)"`
	LinkList []models.Link
	RawHtml  string `valid:"Required;"`
}

// Run handles transaction for `keyword`, `page` and `link` table.
// If there are some errors in the transaction, it rollbacks them all and returns with errors.
func (service *CreateKeywordResult) Run() (keyword *models.Keyword, err error) {
	validator := validation.Validation{}

	valid, err := validator.Valid(service)
	if err != nil {
		return nil, err
	}

	if !valid {
		return nil, validator.Errors[0]
	}

	keyword, err = createKeywordResult(service)
	if err != nil {
		return nil, err
	}

	return keyword, err
}

// Valid handles some custom form validations about validating valid visit url
// and validating valid url from the list of link
func (service *CreateKeywordResult) Valid(validation *validation.Validation) {
	// Validate visit url pattern
	if !validateUrl(service.Url) {
		err := validation.SetError("Url", ValidationMessages["InvalidUrl"])
		if err == nil {
			logs.Warning(fmt.Sprintf("Set validation error failed: %v", err))
		}
	}

	// Validate list of link
	for _, link := range service.LinkList {
		if !validateUrl(link.Url) {
			err := validation.SetError("LinkList", ValidationMessages["InvalidLinkList"])
			if err == nil {
				logs.Warning(fmt.Sprintf("Set validation error failed: %v", err))
			}
			break
		}
	}
}

// validateUrl validates valid url
func validateUrl(plainUrl string) (valid bool) {
	_, err := url.ParseRequestURI(plainUrl)

	return err == nil
}

// createKeywordResult will update `keyword` and save key result `page` and `link` table at the same time.
func createKeywordResult(keywordResult *CreateKeywordResult) (keyword *models.Keyword, err error) {
	ormer := orm.NewOrm()

	// Transaction : Begin
	txnOrmer, err := ormer.Begin()
	if err != nil {
		return nil, err
	}

	// Transaction : Keyword
	keyword = keywordResult.Keyword
	keyword.Url = keywordResult.Url
	keyword.Status = models.GetKeywordStatus("completed")
	_, err = txnOrmer.Update(keyword)
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
