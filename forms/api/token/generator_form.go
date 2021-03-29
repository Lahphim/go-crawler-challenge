package apiforms

import (
	"context"
	"strconv"

	. "go-crawler-challenge/forms"
	"go-crawler-challenge/helpers"
	"go-crawler-challenge/models"
	. "go-crawler-challenge/services/oauth"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	"github.com/go-oauth2/oauth2/v4"
)

type GeneratorForm struct {
	ClientId     string `valid:"Required;"`
	ClientSecret string `valid:"Required;"`
	GrantType    string `valid:"Required;"`
	Email        string `valid:"Required; Email; MaxSize(100)"`
	Password     string `valid:"Required;"`
}

var currentUser *models.User

// Valid handles some custom form validation about validating existing client, grant type and user
func (form *GeneratorForm) Valid(validation *validation.Validation) {
	_, err := form.validateClient()
	if err != nil {
		err = validation.SetError("ClientId", ValidationMessages["InvalidClient"])
		if err == nil {
			logs.Warning(ValidationMessages["ValidationFailed"])
		}

		return
	}

	isValidGrantType := ServerOauth.CheckGrantType(oauth2.GrantType(form.GrantType))
	if !isValidGrantType {
		err = validation.SetError("GrantType", ValidationMessages["InvalidGrantType"])
		if err == nil {
			logs.Warning(ValidationMessages["ValidationFailed"])
		}

		return
	}

	user, err := form.validateUser()
	if err != nil {
		err = validation.SetError("Email", ValidationMessages["InvalidCredential"])
		if err == nil {
			logs.Warning(ValidationMessages["ValidationFailed"])
		}

		return
	}

	currentUser = user
}

// Generate handles generating a new token.
// If there are some invalid cases, it will returns a first error to the controller.
func (form *GeneratorForm) Generate() (authToken oauth2.TokenInfo, err error) {
	validator := validation.Validation{}

	valid, err := validator.Valid(form)
	if err != nil {
		return nil, err
	}

	if !valid {
		return nil, validator.Errors[0]
	}

	tokenGenerator := oauth2.TokenGenerateRequest{
		ClientID:     form.ClientId,
		ClientSecret: form.ClientSecret,
		UserID:       strconv.FormatInt(currentUser.Id, 10),
	}

	authToken, err = ServerOauth.GetAccessToken(
		context.Background(),
		oauth2.GrantType(form.GrantType),
		&tokenGenerator,
	)
	if err != nil {
		return nil, err
	}

	return authToken, nil
}

func (form *GeneratorForm) validateClient() (client oauth2.ClientInfo, err error) {
	client, err = ClientStore.GetByID(context.TODO(), form.ClientId)
	if err != nil {
		return nil, err
	}

	if client.GetSecret() != form.ClientSecret {
		return nil, err
	}

	return client, nil
}

func (form *GeneratorForm) validateUser() (user *models.User, err error) {
	user, err = models.GetUserByEmail(form.Email)
	if err != nil {
		return nil, err
	}

	err = helpers.CheckMatchPassword(user.HashedPassword, form.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
