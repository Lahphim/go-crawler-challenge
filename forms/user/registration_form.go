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
	"ExistedEmail":      "Email is already in use",
	"ConfirmationMatch": "Confirm password confirmation must match",
}

// Valid handles some custom form validations and sets some errors for the invalid case
func (form *RegistrationForm) Valid(validation *validation.Validation) {
	existedUser, _ := models.GetUserByEmail(form.Email)
	if existedUser != nil {
		_ = validation.SetError("Email", ValidationMessages["ExistedEmail"])
	}

	if form.Password != form.ConfirmPassword {
		_ = validation.SetError("ConfirmPassword", ValidationMessages["ConfirmationMatch"])
	}
}

// Create handles validation and creating a new user from the registration form.
// If there are some invalid cases, it will returns with set of errors to the controller.
func (form RegistrationForm) Create() (id int64, errors []error) {
	validation := validation.Validation{}

	valid, err := validation.Valid(&form)
	if err != nil {
		errors = append(errors, err)

		return 0, errors
	}

	if !valid {
		for _, err := range validation.Errors {
			errors = append(errors, err)
		}

		return 0, errors
	}

	encryptedPassword, err := helpers.EncryptPassword(form.Password)
	if err != nil {
		errors = append(errors, err)

		return 0, errors
	}

	user := models.User{
		Email:             form.Email,
		EncryptedPassword: encryptedPassword,
	}
	id, err = models.AddUser(&user)
	if err != nil {
		errors = append(errors, err)

		return 0, errors
	}

	return id, errors
}
