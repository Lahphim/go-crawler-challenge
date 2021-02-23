package controllers_test

import (
	"fmt"
	"net/http"
	"net/url"

	. "go-crawler-challenge/tests"
	. "go-crawler-challenge/tests/fixtures"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ScraperController", func() {
	BeforeEach(func() {
		keyword := "keyword"
		visitURL := fmt.Sprintf("https://www.google.com/search?q=%s&lr=lang_en", url.QueryEscape(keyword))
		cassetteName := "scraper/success_valid_params"

		RecordCassette(cassetteName, visitURL)
	})

	AfterEach(func() {
		TruncateTable("user")
	})

	Describe("POST /scraper/keyword", func() {
		Context("when the user has already signed in", func() {
			Context("given a valid param", func() {
				It("shows a success message", func() {
					user := FabricateUser("dev@nimblehq.co", "password")
					body := GenerateRequestBody(map[string]string{
						"keyword": "keyword",
					})
					response := MakeAuthenticatedRequest("POST", "/scraper/keyword", body, user)
					flash := GetFlashMessage(response.Cookies())

					Expect(flash.Data["success"]).NotTo(BeEmpty())
					Expect(flash.Data["error"]).To(BeEmpty())
				})

				It("redirects to dashboard page", func() {
					user := FabricateUser("dev@nimblehq.co", "password")
					body := GenerateRequestBody(map[string]string{
						"keyword": "keyword",
					})
					response := MakeAuthenticatedRequest("POST", "/scraper/keyword", body, user)
					currentPath := GetCurrentPath(response)

					Expect(response.StatusCode).To(Equal(http.StatusFound))
					Expect(currentPath).To(Equal("/dashboard"))
				})
			})

			Context("given an INVALID param", func() {
				It("shows an error message", func() {
					user := FabricateUser("dev@nimblehq.co", "password")
					body := GenerateRequestBody(map[string]string{
						"keyword": "",
					})
					response := MakeAuthenticatedRequest("POST", "/scraper/keyword", body, user)
					flash := GetFlashMessage(response.Cookies())

					Expect(flash.Data["success"]).To(BeEmpty())
					Expect(flash.Data["error"]).NotTo(BeEmpty())
				})

				It("redirects to dashboard page", func() {
					user := FabricateUser("dev@nimblehq.co", "password")
					body := GenerateRequestBody(map[string]string{
						"keyword": "",
					})
					response := MakeAuthenticatedRequest("POST", "/scraper/keyword", body, user)
					currentPath := GetCurrentPath(response)

					Expect(response.StatusCode).To(Equal(http.StatusFound))
					Expect(currentPath).To(Equal("/dashboard"))
				})
			})
		})

		Context("when the user has NOT signed in yet", func() {
			It("redirects to sign-in page", func() {
				body := GenerateRequestBody(map[string]string{
					"keyword": "",
				})
				response := MakeRequest("POST", "/scraper/keyword", body)
				currentPath := GetCurrentPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath).To(Equal("/user/sign_in"))
			})
		})
	})
})
