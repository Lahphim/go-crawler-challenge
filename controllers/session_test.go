package controllers_test

import (
	"go-crawler-challenge/controllers"
	"net/http"

	. "go-crawler-challenge/tests/test_helpers"
	. "go-crawler-challenge/tests/test_helpers/fabricators"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SessionController", func() {
	AfterEach(func() {
		TruncateTable("user")
	})

	Describe("GET /user/sign_in", func() {
		Context("when the user has NOT signed in yet", func() {
			It("renders with status 200", func() {
				response := MakeRequest("GET", "/user/sign_in", nil)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})
		})

		Context("when the user has already signed in", func() {
			It("redirects to root path", func() {
				user := FabricateUser("dev@nimblehq.co", "password")

				response := MakeAuthenticatedRequest("GET", "/user/sign_in", nil, user)
				currentPath := GetCurrentPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath).To(Equal("/"))
			})
		})
	})

	Describe("POST /session/create", func() {
		Context("when the user has NOT signed in yet", func() {
			Context("given a valid params", func() {
				It("redirects to root page", func() {
					_ = FabricateUser("dev@nimblehq.co", "password")
					body := GenerateRequestBody(map[string]string{
						"email":    "dev@nimblehq.co",
						"password": "password",
					})

					response := MakeRequest("POST", "/session/create", body)
					currentPath := GetCurrentPath(response)

					Expect(response.StatusCode).To(Equal(http.StatusFound))
					Expect(currentPath).To(Equal("/"))
				})

				It("shows a success message", func() {
					_ = FabricateUser("dev@nimblehq.co", "password")
					body := GenerateRequestBody(map[string]string{
						"email":    "dev@nimblehq.co",
						"password": "password",
					})

					response := MakeRequest("POST", "/session/create", body)
					flash := GetFlashMessage(response.Cookies())

					Expect(flash.Data["success"]).To(Equal("You have successfully signed in"))
					Expect(flash.Data["error"]).To(BeEmpty())
				})

				It("sets a session for the current user", func() {
					user := FabricateUser("dev@nimblehq.co", "password")
					body := GenerateRequestBody(map[string]string{
						"email":    "dev@nimblehq.co",
						"password": "password",
					})

					response := MakeRequest("POST", "/session/create", body)
					currentUserId := GetSession(response.Cookies(), controllers.CurrentUserKey)

					Expect(currentUserId).To(Equal(user.Id))
				})
			})

			Context("given an INVALID params", func() {
				Context("when the user does NOT exist", func() {
					It("redirects to sign-in page", func() {
						body := GenerateRequestBody(map[string]string{
							"email":    "dev@nimblehq.co",
							"password": "password",
						})

						response := MakeRequest("POST", "/session/create", body)
						currentPath := GetCurrentPath(response)

						Expect(response.StatusCode).To(Equal(http.StatusFound))
						Expect(currentPath).To(Equal("/user/sign_in"))
					})

					It("shows a error message", func() {
						body := GenerateRequestBody(map[string]string{
							"email":    "dev@nimblehq.co",
							"password": "password",
						})

						response := MakeRequest("POST", "/session/create", body)
						flash := GetFlashMessage(response.Cookies())

						Expect(flash.Data["success"]).To(BeEmpty())
						Expect(flash.Data["error"]).NotTo(BeEmpty())
					})

					It("does NOT set any sessions", func() {
						body := GenerateRequestBody(map[string]string{
							"email":    "dev@nimblehq.co",
							"password": "password",
						})

						response := MakeRequest("POST", "/session/create", body)
						currentUserId := GetSession(response.Cookies(), controllers.CurrentUserKey)

						Expect(currentUserId).To(BeNil())
					})
				})

				Context("when sign in with INVALID credential", func() {
					It("redirects to sign-in page", func() {
						_ = FabricateUser("dev@nimblehq.co", "password")
						body := GenerateRequestBody(map[string]string{
							"email":    "dev@nimblehq.co",
							"password": "INVALID_PASSWORD",
						})

						response := MakeRequest("POST", "/session/create", body)
						currentPath := GetCurrentPath(response)

						Expect(response.StatusCode).To(Equal(http.StatusFound))
						Expect(currentPath).To(Equal("/user/sign_in"))
					})

					It("shows a error message", func() {
						_ = FabricateUser("dev@nimblehq.co", "password")
						body := GenerateRequestBody(map[string]string{
							"email":    "dev@nimblehq.co",
							"password": "INVALID_PASSWORD",
						})

						response := MakeRequest("POST", "/session/create", body)
						flash := GetFlashMessage(response.Cookies())

						Expect(flash.Data["success"]).To(BeEmpty())
						Expect(flash.Data["error"]).NotTo(BeEmpty())
					})

					It("does NOT set any sessions", func() {
						_ = FabricateUser("dev@nimblehq.co", "password")
						body := GenerateRequestBody(map[string]string{
							"email":    "dev@nimblehq.co",
							"password": "INVALID_PASSWORD",
						})

						response := MakeRequest("POST", "/session/create", body)
						currentUserId := GetSession(response.Cookies(), controllers.CurrentUserKey)

						Expect(currentUserId).To(BeNil())
					})
				})
			})
		})

		Context("when the user has already signed in", func() {
			It("redirects to root page", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				body := GenerateRequestBody(map[string]string{
					"email":    "dev@nimblehq.co",
					"password": "password",
				})

				response := MakeAuthenticatedRequest("POST", "/session/create", body, user)
				currentPath := GetCurrentPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath).To(Equal("/"))
			})
		})
	})
})
