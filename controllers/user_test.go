package controllers_test

import (
	"net/http"

	. "go-crawler-challenge/tests/test_helpers"
	. "go-crawler-challenge/tests/test_helpers/fabricators"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserController", func() {
	AfterEach(func() {
		TruncateTable("user")
	})

	Describe("GET /user/sign_up", func() {
		Context("when the user has NOT signed in yet", func() {
			It("renders with status 200", func() {
				response := MakeRequest("GET", "/user/sign_up", nil)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})
		})

		Context("when the user has already signed in", func() {
			It("redirects to root path", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				response := MakeAuthenticatedRequest("GET", "/user/sign_up", nil, user)
				currentPath := GetCurrentPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath).To(Equal("/"))
			})
		})
	})
})
