package forms

import (
	"fmt"

	"go-crawler-challenge/helpers"
	"go-crawler-challenge/models"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
)

type AuthenticationForm struct {
	Email    string `form:"email" valid:"Required; Email; MaxSize(100)"`
	Password string `form:"password" valid:"Required;"`
}

var ValidationMessages = map[string]string{
	"InvalidCredential": "Your email or password is incorrect",
}

var currentUser *models.User

// Valid handles some custom form validation about validating existing user and checking match password
func (form *AuthenticationForm) Valid(validation *validation.Validation) {
	user, _ := models.GetUserByEmail(form.Email)
	if user == nil {
		err := validation.SetError("Email", ValidationMessages["InvalidCredential"])
		if err == nil {
			logs.Warning(fmt.Sprintf("Set validation error failed: %v", err))
		}
	} else {
		err := helpers.CheckMatchPassword(user.HashedPassword, form.Password)
		if err != nil {
			err = validation.SetError("Email", ValidationMessages["InvalidCredential"])
			if err == nil {
				logs.Warning(fmt.Sprintf("Set validation error failed: %v", err))
			}
		} else {
			currentUser = user
		}
	}
}

// Authenticate handles validation and authenticate the login form.
// If there are some invalid cases, it will returns with set of errors to the controller.
func (form *AuthenticationForm) Authenticate() (user *models.User, errors []error) {
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

	return currentUser, errors
}
