package forms

import (
	"github.com/beego/beego/v2/core/validation"
)

type SearchKeywordForm struct {
	Keyword string `form:"keyword" valid:"Required"`
}

func (form *SearchKeywordForm) Validate() (errors []error) {
	validator := validation.Validation{}

	valid, err := validator.Valid(form)
	if err != nil {
		return []error{err}
	}

	if !valid {
		for _, err := range validator.Errors {
			errors = append(errors, err)
		}

		return errors
	}

	return errors
}
