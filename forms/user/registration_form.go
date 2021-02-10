package forms

import (
	"go-crawler-challenge/helpers"
	"go-crawler-challenge/models"

	"github.com/beego/beego/v2/core/validation"
)

type RegistrationForm struct {
	Email           string `valid:"Required; MaxSize(100)"`
	Password        string `valid:"Required; MinSize(6); MaxSize(12)"`
	ConfirmPassword string `valid:"Required; MinSize(6); MaxSize(12)"`
}

var ValidationMessages = map[string]string{
	"ExistingEmail":   "Email is already in use",
	"ConfirmPassword": "Confirm password confirmation must match",
}

// Valid handles some custom form validations and sets some errors for the invalid case
func (form *RegistrationForm) Valid(validation *validation.Validation) {
	existedUser, _ := models.GetUserByEmail(form.Email)
	if existedUser != nil {
		_ = validation.SetError("Email", ValidationMessages["ExistingEmail"])
	}

	if form.Password != form.ConfirmPassword {
		_ = validation.SetError("ConfirmPassword", ValidationMessages["ConfirmPassword"])
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