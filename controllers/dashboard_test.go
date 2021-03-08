package controllers_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"

	"github.com/bxcodec/faker/v3"

	. "go-crawler-challenge/tests"
	. "go-crawler-challenge/tests/fixtures"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DashboardController", func() {
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
		TruncateTable("keyword")
		TruncateTable("user")
	})

	Describe("GET /dashboard", func() {
		Context("when the user has already signed in", func() {
			It("renders with status 200", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				response := MakeAuthenticatedRequest("GET", "/dashboard", nil, user)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})

			Context("given 12 records of the keyword in the database", func() {
				It("shows 2 pages in pagination", func() {
					pageClass := "pagination__page"

					totalRecords := 12
					user := FabricateUser(faker.Email(), faker.Password())
					for i := 1; i <= totalRecords; i++ {
						_ = FabricateKeyword(fmt.Sprintf("key%02dword", i), "https://www.sample.com", user)
					}

					response := MakeAuthenticatedRequest("GET", "/dashboard", nil, user)
					body, err := ioutil.ReadAll(response.Body)
					if err != nil {
						Fail("Read body content failed: " + err.Error())
					}
					err = response.Body.Close()
					if err != nil {
						Fail("Close body content failed: " + err.Error())
					}

					r, err := regexp.Compile(pageClass)
					if err != nil {
						Fail("Regexp failed: " + err.Error())
					}

					matches := r.FindAllString(string(body), -1)

					Expect(len(matches)).To(BeNumerically("==", 3))
				})

				It("shows the latest search of the keyword at the top of the table", func() {
					totalRecords := 12
					user := FabricateUser(faker.Email(), faker.Password())
					for i := 1; i <= totalRecords; i++ {
						_ = FabricateKeyword(fmt.Sprintf("key%02dword", i), "https://www.sample.com", user)
					}

					response := MakeAuthenticatedRequest("GET", "/dashboard", nil, user)
					body, err := ioutil.ReadAll(response.Body)
					if err != nil {
						Fail("Read body content failed: " + err.Error())
					}
					err = response.Body.Close()
					if err != nil {
						Fail("Close body content failed: " + err.Error())
					}

					r, err := regexp.Compile("key[0-9]{2}word")
					if err != nil {
						Fail("Regexp failed: " + err.Error())
					}

					matches := r.FindAllString(string(body), -1)

					Expect(matches[0]).To(Equal("key12word"))
				})

				It("shows 10 records", func() {
					totalRecords := 12
					user := FabricateUser(faker.Email(), faker.Password())
					for i := 1; i <= totalRecords; i++ {
						_ = FabricateKeyword(fmt.Sprintf("key%02dword", i), "https://www.sample.com", user)
					}

					response := MakeAuthenticatedRequest("GET", "/dashboard", nil, user)
					body, err := ioutil.ReadAll(response.Body)
					if err != nil {
						Fail("Read body content failed: " + err.Error())
					}
					err = response.Body.Close()
					if err != nil {
						Fail("Close body content failed: " + err.Error())
					}

					r, err := regexp.Compile("key[0-9]{2}word")
					if err != nil {
						Fail("Regexp failed: " + err.Error())
					}

					matches := r.FindAllString(string(body), -1)

					Expect(len(matches)).To(BeNumerically("==", 10))
				})
			})

			Context("given 0 record of the keyword", func() {
				It("shows a placeholder", func() {
					placeholderClass := "table__row--placeholder"
					user := FabricateUser(faker.Email(), faker.Password())

					response := MakeAuthenticatedRequest("GET", "/dashboard", nil, user)
					body, err := ioutil.ReadAll(response.Body)
					if err != nil {
						Fail("Read body content failed: " + err.Error())
					}
					err = response.Body.Close()
					if err != nil {
						Fail("Close body content failed: " + err.Error())
					}

					r, err := regexp.Compile(placeholderClass)
					if err != nil {
						Fail("Regexp failed: " + err.Error())
					}

					matches := r.FindAllString(string(body), -1)

					Expect(len(matches)).To(BeNumerically("==", 1))
				})

				It("does NOT show a pagination", func() {
					paginationClass := "pagination"
					user := FabricateUser(faker.Email(), faker.Password())

					response := MakeAuthenticatedRequest("GET", "/dashboard", nil, user)
					body, err := ioutil.ReadAll(response.Body)
					if err != nil {
						Fail("Read body content failed: " + err.Error())
					}
					err = response.Body.Close()
					if err != nil {
						Fail("Close body content failed: " + err.Error())
					}

					r, err := regexp.Compile(paginationClass)
					if err != nil {
						Fail("Regexp failed: " + err.Error())
					}

					matches := r.FindAllString(string(body), -1)

					Expect(len(matches)).To(BeZero())
				})
			})
		})

		Context("when the user has NOT signed in yet", func() {
			It("redirects to sign-in page", func() {
				response := MakeRequest("GET", "/dashboard", nil)
				currentPath := GetCurrentPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath).To(Equal("/user/sign_in"))
			})
		})
	})

	Describe("POST /dashboard/search", func() {
		Context("when the user has already signed in", func() {
			Context("given a valid param", func() {
				XIt("shows a success message", func() {
					user := FabricateUser("dev@nimblehq.co", "password")
					body := GenerateRequestBody(map[string]string{
						"keyword": "keyword",
					})
					response := MakeAuthenticatedRequest("POST", "/dashboard/search", body, user)
					flash := GetFlashMessage(response.Cookies())

					Expect(flash.Data["success"]).NotTo(BeEmpty())
					Expect(flash.Data["error"]).To(BeEmpty())
				})

				XIt("redirects to dashboard page", func() {
					user := FabricateUser("dev@nimblehq.co", "password")
					body := GenerateRequestBody(map[string]string{
						"keyword": "keyword",
					})
					response := MakeAuthenticatedRequest("POST", "/dashboard/search", body, user)
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
					response := MakeAuthenticatedRequest("POST", "/dashboard/search", body, user)
					flash := GetFlashMessage(response.Cookies())

					Expect(flash.Data["success"]).To(BeEmpty())
					Expect(flash.Data["error"]).NotTo(BeEmpty())
				})

				It("redirects to dashboard page", func() {
					user := FabricateUser("dev@nimblehq.co", "password")
					body := GenerateRequestBody(map[string]string{
						"keyword": "",
					})
					response := MakeAuthenticatedRequest("POST", "/dashboard/search", body, user)
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
				response := MakeRequest("POST", "/dashboard/search", body)
				currentPath := GetCurrentPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath).To(Equal("/user/sign_in"))
			})
		})
	})
})
