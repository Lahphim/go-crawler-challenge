package forms

import (
	"fmt"

	. "go-crawler-challenge/forms"
	"go-crawler-challenge/helpers"
	"go-crawler-challenge/models"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
)

type RegistrationForm struct {
	Email           string `form:"email" valid:"Required; Email; MaxSize(100)"`
	Password        string `form:"password" valid:"Required; MinSize(6); MaxSize(12)"`
	ConfirmPassword string `form:"confirm_password" valid:"Required; MinSize(6); MaxSize(12)"`
}

// Valid handles some custom form validations and sets some errors for the invalid case
func (form *RegistrationForm) Valid(validation *validation.Validation) {
	existingUser, _ := models.GetUserByEmail(form.Email)
	if existingUser != nil {
		err := validation.SetError("Email", ValidationMessages["ExistingEmail"])
		if err == nil {
			logs.Warning(fmt.Sprintf("Set validation error failed: %v", err))
		}
	}

	if form.Password != form.ConfirmPassword {
		err := validation.SetError("ConfirmPassword", ValidationMessages["ConfirmPassword"])
		if err == nil {
			logs.Warning(fmt.Sprintf("Set validation error failed: %v", err))
		}
	}
}

// Create handles validation and creating a new user from the registration form.
// If there are some invalid cases, it will returns with set of errors to the controller.
func (form *RegistrationForm) Create() (user *models.User, errors []error) {
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

	hashedPassword, err := helpers.HashPassword(form.Password)
	if err != nil {
		return nil, []error{err}
	}

	user = &models.User{
		Email:          form.Email,
		HashedPassword: hashedPassword,
	}
	userId, err := models.AddUser(user)
	if err != nil {
		return nil, []error{err}
	}

	user, err = models.GetUserById(userId)
	if err != nil {
		return nil, []error{err}
	}

	return user, errors
}
