package controllers_test

import (
	"net/http"

	. "go-crawler-challenge/tests"
	. "go-crawler-challenge/tests/fixtures"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DashboardController", func() {
	AfterEach(func() {
		TruncateTable("user")
	})

	Describe("GET /", func() {
		Context("when the user has NOT signed in yet", func() {
			It("redirects to sign-in page", func() {
				response := MakeRequest("GET", "/dashboard", nil)
				currentPath := GetCurrentPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath).To(Equal("/user/sign_in"))
			})
		})

		Context("when the user has already signed in", func() {
			It("renders with status 200", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				response := MakeAuthenticatedRequest("GET", "/dashboard", nil, user)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})
		})
	})
})
