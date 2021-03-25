package helpers_test

import (
	"fmt"

	"go-crawler-challenge/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OauthClientHelper", func() {
	Describe("#GenerateSecretToken", func() {
		It("returns a secret token", func() {
			secretToken, err := helpers.GenerateSecretToken()
			if err != nil {
				Fail(fmt.Sprintf("Generate secret token failed: %v", err.Error()))
			}

			Expect(len(secretToken)).To(BeNumerically(">", 0))
			Expect(secretToken).NotTo(BeNil())
		})
	})
})
