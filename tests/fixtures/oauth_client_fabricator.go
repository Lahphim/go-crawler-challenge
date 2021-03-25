package fixtures

import (
	. "go-crawler-challenge/services/oauth"

	"github.com/go-oauth2/oauth2/v4/models"
	. "github.com/onsi/ginkgo"
)

func FabricateOauthClient(id string, secret string, domain string) (oauth_client *models.Client) {
	oauth_client = &models.Client{
		ID:     id,
		Secret: secret,
		Domain: domain,
	}

	err := ClientStore.Create(oauth_client)
	if err != nil {
		Fail("Add page failed: " + err.Error())
	}

	return oauth_client
}
