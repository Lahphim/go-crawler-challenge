package controllers_test

import (
	"net/http"

	. "go-crawler-challenge/tests"
	. "go-crawler-challenge/tests/fixtures"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MainController", func() {
	AfterEach(func() {
		TruncateTable("user")
	})

	Describe("GET /", func() {
		Context("when the user has NOT signed in yet", func() {
			It("renders with status 200", func() {
				response := MakeRequest("GET", "/", nil)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})
		})

		Context("when the user has already signed in", func() {
			It("renders with status 200", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				response := MakeAuthenticatedRequest("GET", "/", nil, user)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})
		})
	})
})
