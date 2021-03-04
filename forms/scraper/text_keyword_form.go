package forms

import (
	"github.com/beego/beego/v2/core/validation"
)

type TextKeywordForm struct {
	Keyword string `form:"keyword" valid:"Required"`
}

// Validate handles validation the search keyword form.
// If the form is invalid, it will returns with set of errors to the controller.
func (form *TextKeywordForm) Validate() (errors []error) {
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
