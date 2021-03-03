package forms

import (
	"fmt"
	"net/url"

	. "go-crawler-challenge/forms"
	"go-crawler-challenge/models"
	"go-crawler-challenge/transactions"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
)

type KeywordResultForm struct {
	Keyword  string `valid:"Required; MaxSize(128)"`
	Url      string `valid:"Required; MaxSize(128)"`
	LinkList []models.Link
	RawHtml  string       `valid:"Required;"`
	User     *models.User `valid:"Required;"`
}

// Valid handles some custom form validations about checking an existing user, validating valid visit url
// and validating valid url from the list of link
func (form *KeywordResultForm) Valid(validation *validation.Validation) {
	// Validate current existing user
	existedUser, _ := models.GetUserById(form.User.Id)
	if existedUser == nil {
		err := validation.SetError("User", ValidationMessages["InvalidUser"])
		if err == nil {
			logs.Warning(fmt.Sprintf("Set validation error failed: %v", err))
		}
	}

	// Validate visit url pattern
	if !validateUrl(form.Url) {
		err := validation.SetError("Url", ValidationMessages["InvalidUrl"])
		if err == nil {
			logs.Warning(fmt.Sprintf("Set validation error failed: %v", err))
		}
	}

	// Validate list of link
	for _, link := range form.LinkList {
		if !validateUrl(link.Url) {
			err := validation.SetError("LinkList", ValidationMessages["InvalidLinkList"])
			if err == nil {
				logs.Warning(fmt.Sprintf("Set validation error failed: %v", err))
			}
			break
		}
	}
}

// Save handles transaction for `keyword`, `page` and `link` table.
// If there are some errors in the transaction, it rollbacks them all and returns with errors.
func (form *KeywordResultForm) Save() (keyword *models.Keyword, errors []error) {
	validator := validation.Validation{}

	valid, err := validator.Valid(form)
	if err != nil {
		return nil, []error{err}
	}

	if !valid {
		for _, err := range validator.Errors {
			errors = append(errors, err)
		}

		return nil, errors
	}

	keyword, err = form.processTransaction()
	if err != nil {
		return nil, []error{err}
	}

	return keyword, errors
}

// processTransaction will save keyword result into `keyword`, `page` and `link` table at the same time.
func (form *KeywordResultForm) processTransaction() (keyword *models.Keyword, err error) {
	keywordResult := &transactions.KeywordResult{
		Keyword:  form.Keyword,
		Url:      form.Url,
		LinkList: form.LinkList,
		RawHtml:  form.RawHtml,
		User:     form.User,
	}

	return transactions.AddKeywordResult(keywordResult)
}

func validateUrl(plainUrl string) (valid bool) {
	_, err := url.ParseRequestURI(plainUrl)

	return err == nil
}
