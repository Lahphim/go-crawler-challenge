package oauth

import (
	"fmt"
	"os"

	"go-crawler-challenge/helpers"

	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/google/uuid"
)

type ClientGenerator struct{}

func (service *ClientGenerator) Generate() (id string, err error) {
	clientId := uuid.New().String()
	clientDomain := fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("PORT"))
	clientSecret, err := helpers.GenerateSecretToken()
	if err != nil {
		return "", err
	}

	client := &models.Client{
		ID:     clientId,
		Secret: clientSecret,
		Domain: clientDomain,
	}
	err = ClientStore.Create(client)
	if err != nil {
		return "", err
	}

	return client.GetID(), nil
}
