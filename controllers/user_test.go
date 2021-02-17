package controllers_test

import (
	"net/http"

	. "go-crawler-challenge/tests"
	. "go-crawler-challenge/tests/fixtures"

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

	Describe("POST /user/create", func() {
		Context("when the user has NOT signed in yet", func() {
			Context("given a valid params", func() {
				It("redirects to sign-in page", func() {
					body := GenerateRequestBody(map[string]string{
						"email":            "dev@nimblehq.co",
						"password":         "password",
						"confirm_password": "password",
					})

					response := MakeRequest("POST", "/user/create", body)
					currentPath := GetCurrentPath(response)

					Expect(response.StatusCode).To(Equal(http.StatusFound))
					Expect(currentPath).To(Equal("/user/sign_in"))
				})

				It("shows a success message", func() {
					body := GenerateRequestBody(map[string]string{
						"email":            "dev@nimblehq.co",
						"password":         "password",
						"confirm_password": "password",
					})

					response := MakeRequest("POST", "/user/create", body)
					flash := GetFlashMessage(response.Cookies())

					Expect(flash.Data["success"]).To(Equal("Congrats on creating a new account"))
					Expect(flash.Data["error"]).To(BeEmpty())
				})
			})

			Context("given an INVALID params", func() {
				It("redirects to sign-up page", func() {
					body := GenerateRequestBody(map[string]string{
						"email":            "",
						"password":         "",
						"confirm_password": "",
					})

					response := MakeRequest("POST", "/user/create", body)
					currentPath := GetCurrentPath(response)

					Expect(response.StatusCode).To(Equal(http.StatusFound))
					Expect(currentPath).To(Equal("/user/sign_up"))
				})

				It("shows an error message", func() {
					body := GenerateRequestBody(map[string]string{
						"email":            "",
						"password":         "",
						"confirm_password": "",
					})

					response := MakeRequest("POST", "/user/create", body)
					flash := GetFlashMessage(response.Cookies())

					Expect(flash.Data["success"]).To(BeEmpty())
					Expect(flash.Data["error"]).NotTo(BeEmpty())
				})
			})
		})

		Context("when the user has already signed in", func() {
			It("redirects to root path", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				body := GenerateRequestBody(map[string]string{
					"email":            "dev@nimblehq.co",
					"password":         "password",
					"confirm_password": "password",
				})

				response := MakeAuthenticatedRequest("POST", "/user/create", body, user)
				currentPath := GetCurrentPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath).To(Equal("/"))
			})
		})
	})
})
