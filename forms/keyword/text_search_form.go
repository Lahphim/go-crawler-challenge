package forms

import (
	"go-crawler-challenge/models"

	"github.com/beego/beego/v2/core/validation"
)

type TextSearchForm struct {
	Keyword string       `form:"keyword" valid:"Required"`
	User    *models.User `valid:"Required"`
}

// Create handles validation and creating a new keyword from the search keyword form.
// If the form is invalid, it will returns an error to the controller.
func (form *TextSearchForm) Create() (err error) {
	validator := validation.Validation{}

	valid, err := validator.Valid(form)
	if err != nil {
		return err
	}

	if !valid {
		return validator.Errors[0]
	}

	keyword := &models.Keyword{
		Keyword: form.Keyword,
		User:    form.User,
	}

	_, err = models.AddKeyword(keyword)
	if err != nil {
		return err
	}

	return nil
}
