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
		// Record for valid case
		keyword := "keyword"
		visitURL := fmt.Sprintf("https://www.google.com/search?q=%s&lr=lang_en", url.QueryEscape(keyword))
		cassetteName := "scraper/success_valid_params"

		RecordCassette(cassetteName, visitURL)

		// Record for invalid case
		keyword = ""
		visitURL = fmt.Sprintf("https://www.google.com/search?q=%s&lr=lang_en", url.QueryEscape(keyword))
		cassetteName = "scraper/success_invalid_params"

		RecordCassette(cassetteName, visitURL)
	})

	AfterEach(func() {
		TruncateTable("user")
	})

	Describe("POST /dashboard/search_keyword", func() {
		Context("when the user has already signed in", func() {
			Context("given a valid param", func() {
				XIt("shows a success message", func() {
					user := FabricateUser("dev@nimblehq.co", "password")
					body := GenerateRequestBody(map[string]string{
						"keyword": "keyword",
					})
					response := MakeAuthenticatedRequest("POST", "/dashboard/search_keyword", body, user)
					flash := GetFlashMessage(response.Cookies())

					Expect(flash.Data["success"]).NotTo(BeEmpty())
					Expect(flash.Data["error"]).To(BeEmpty())
				})

				XIt("redirects to dashboard page", func() {
					user := FabricateUser("dev@nimblehq.co", "password")
					body := GenerateRequestBody(map[string]string{
						"keyword": "keyword",
					})
					response := MakeAuthenticatedRequest("POST", "/dashboard/search_keyword", body, user)
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
					response := MakeAuthenticatedRequest("POST", "/dashboard/search_keyword", body, user)
					flash := GetFlashMessage(response.Cookies())

					Expect(flash.Data["success"]).To(BeEmpty())
					Expect(flash.Data["error"]).NotTo(BeEmpty())
				})

				It("redirects to dashboard page", func() {
					user := FabricateUser("dev@nimblehq.co", "password")
					body := GenerateRequestBody(map[string]string{
						"keyword": "",
					})
					response := MakeAuthenticatedRequest("POST", "/dashboard/search_keyword", body, user)
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
				response := MakeRequest("POST", "/dashboard/search_keyword", body)
				currentPath := GetCurrentPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath).To(Equal("/user/sign_in"))
			})
		})
	})
})
