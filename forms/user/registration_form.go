package forms

import (
	log "github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
)

type RegistrationForm struct {
	Email           string `valid:"Required; MaxSize(100)"`
	Password        string `valid:"Required; MinSize(6); MaxSize(12)"`
	ConfirmPassword string `valid:"Required; MinSize(6); MaxSize(12)"`
}

func (form *RegistrationForm) Valid(validation *validation.Validation) {}

func (form RegistrationForm) Create() (id *int64, errors []error) {
	validation := validation.Validation{}

	valid, err := validation.Valid(&form)
	if err != nil {
		return nil, errors
	}

	if !valid {
		for _, err := range validation.Errors {
			log.Info(err)
		}
	}

	return id, errors
}
