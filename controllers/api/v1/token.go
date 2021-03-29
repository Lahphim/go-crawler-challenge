package apiv1controllers

import (
	. "go-crawler-challenge/controllers/api"
	apiforms "go-crawler-challenge/forms/api/token"
	"go-crawler-challenge/models"
	v1serializers "go-crawler-challenge/serializers/v1"

	"github.com/beego/beego/v2/core/logs"
	"github.com/google/jsonapi"
)

type TokenController struct {
	BaseController
}

func (c *TokenController) URLMapping() {
	c.Mapping("Create", c.Create)
}

func (c *TokenController) Create() {
	credential := &models.Credential{}
	err := jsonapi.UnmarshalPayload(c.Ctx.Request.Body, credential)
	if err != nil {
		err = c.RenderGenericError(err)
		if err != nil {
			logs.Error("Generic error: ", err.Error())
		}
	} else {
		tokenGeneratorForm := apiforms.GeneratorForm{
			ClientId:     credential.ClientId,
			ClientSecret: credential.ClientSecret,
			GrantType:    credential.GrantType,
			Email:        credential.Email,
			Password:     credential.Password,
		}
		authToken, err := tokenGeneratorForm.Generate()
		if err != nil {
			err = c.RenderUnauthorizedError(err)
			if err != nil {
				logs.Error("Generic error: ", err.Error())
			}
		} else {
			tokenInformationSerializer := v1serializers.TokenInformation{
				AccessToken:  authToken.GetAccess(),
				RefreshToken: authToken.GetRefresh(),
				Expiry:       authToken.GetAccessExpiresIn(),
			}
			err = c.RenderJSON(tokenInformationSerializer.Data())
			if err != nil {
				err = c.RenderGenericError(err)
				if err != nil {
					logs.Error("Generic error: ", err.Error())
				}
			}
		}
	}
}
