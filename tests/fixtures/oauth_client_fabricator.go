package fixtures

import (
	. "go-crawler-challenge/services/oauth"

	"github.com/go-oauth2/oauth2/v4/models"
	. "github.com/onsi/ginkgo"
)

func FabricateOauthClient(id string, secret string) (oauthClient *models.Client) {
	oauthClient = &models.Client{
		ID:     id,
		Secret: secret,
	}

	err := ClientStore.Create(oauthClient)
	if err != nil {
		Fail("Add OauthClient failed: " + err.Error())
	}

	return oauthClient
}
