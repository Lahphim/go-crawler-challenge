package oauth_test

import (
	"context"
	"fmt"

	"go-crawler-challenge/services/oauth"
	. "go-crawler-challenge/tests"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Oauth/ClientGenerator", func() {
	AfterEach(func() {
		TruncateTable("oauth2_clients")
		TruncateTable("user")
	})

	Describe("#Generate", func() {
		It("returns an oauth client id", func() {
			service := oauth.ClientGenerator{}
			oauthClientId, err := service.Generate()
			if err != nil {
				Fail(fmt.Sprintf("Generate an oauth client failed: %v", err.Error()))
			}

			Expect(len(oauthClientId)).To(BeNumerically(">", 0))
			Expect(oauthClientId).NotTo(BeNil())
		})

		It("creates a new oauth client record", func() {
			service := oauth.ClientGenerator{}
			oauthClientId, err := service.Generate()
			if err != nil {
				Fail(fmt.Sprintf("Generate an oauth client failed: %v", err.Error()))
			}

			oauthClient, err := oauth.ClientStore.GetByID(context.TODO(), oauthClientId)
			if err != nil {
				Fail(fmt.Sprintf("Get an oauth client by id failed: %v", err.Error()))
			}

			id := oauthClient.GetID()
			secret := oauthClient.GetSecret()

			Expect(id).To(Equal(oauthClientId))

			Expect(len(secret)).To(BeNumerically(">", 0))
			Expect(secret).NotTo(BeNil())
		})
	})
})
