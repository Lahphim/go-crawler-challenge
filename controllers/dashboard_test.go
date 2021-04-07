package controllers_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	. "go-crawler-challenge/tests"
	. "go-crawler-challenge/tests/fixtures"

	"github.com/PuerkitoBio/goquery"
	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DashboardController", func() {
	AfterEach(func() {
		TruncateTable("keyword")
		TruncateTable("user")
	})

	Describe("GET /dashboard", func() {
		Context("when the user has already signed in", func() {
			It("renders with status 200", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				response := MakeAuthenticatedRequest("GET", "/dashboard", nil, nil, user)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})

			Context("given 12 records of the keyword in the database", func() {
				It("lists 2 pages in pagination", func() {
					pageClass := "data-page-number"

					totalRecords := 12
					user := FabricateUser(faker.Email(), faker.Password())
					for i := 1; i <= totalRecords; i++ {
						_ = FabricateKeyword(fmt.Sprintf("key%02dword", i), "https://www.sample.com", 0, user)
					}

					response := MakeAuthenticatedRequest("GET", "/dashboard", nil, nil, user)
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

					Expect(len(matches)).To(BeNumerically("==", 2))
				})

				It("shows the latest search of the keyword at the top of the table", func() {
					totalRecords := 12
					user := FabricateUser(faker.Email(), faker.Password())
					for i := 1; i <= totalRecords; i++ {
						_ = FabricateKeyword(fmt.Sprintf("key%02dword", i), "https://www.sample.com", 0, user)
					}

					response := MakeAuthenticatedRequest("GET", "/dashboard", nil, nil, user)
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
						_ = FabricateKeyword(fmt.Sprintf("key%02dword", i), "https://www.sample.com", 0, user)
					}

					response := MakeAuthenticatedRequest("GET", "/dashboard", nil, nil, user)
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

					response := MakeAuthenticatedRequest("GET", "/dashboard", nil, nil, user)
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

					response := MakeAuthenticatedRequest("GET", "/dashboard", nil, nil, user)
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

			Context("given `keyword` param is set", func() {
				It("shows only `keyword` filtered records", func() {
					keywordFilter := "expected_keyword"
					user := FabricateUser(faker.Email(), faker.Password())
					// Exact match keyword
					_ = FabricateKeyword(keywordFilter, faker.URL(), 0, user)
					// Fuzzy match keyword
					_ = FabricateKeyword(fmt.Sprintf("%v %v %v", faker.Word(), keywordFilter, faker.Word()), faker.URL(), 0, user)
					// Non-match keyword
					_ = FabricateKeyword("nonexpected__keyword", faker.URL(), 0, user)

					response := MakeAuthenticatedRequest("GET", fmt.Sprintf("/dashboard?keyword=%v", keywordFilter), nil, nil, user)
					err := response.Body.Close()
					if err != nil {
						Fail("Close body content failed: " + err.Error())
					}

					// Load the HTML document
					document, err := goquery.NewDocumentFromReader(response.Body)
					if err != nil {
						Fail("New document from reader failed: " + err.Error())
					}

					var matches []string
					document.Find(".list-keyword .table__row .table__cell .link").Each(func(_ int, selector *goquery.Selection) {
						matches = append(matches, selector.Text())
					})

					Expect(len(matches)).To(BeNumerically("==", 2))
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
				It("shows a success message", func() {
					user := FabricateUser("dev@nimblehq.co", "password")
					body := GenerateRequestBody(map[string]string{
						"keyword": "keyword",
					})
					response := MakeAuthenticatedRequest("POST", "/dashboard/search", nil, body, user)
					flash := GetFlashMessage(response.Cookies())

					Expect(flash.Data["success"]).NotTo(BeEmpty())
					Expect(flash.Data["error"]).To(BeEmpty())
				})

				It("redirects to dashboard page", func() {
					user := FabricateUser("dev@nimblehq.co", "password")
					body := GenerateRequestBody(map[string]string{
						"keyword": "keyword",
					})
					response := MakeAuthenticatedRequest("POST", "/dashboard/search", nil, body, user)
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
					response := MakeAuthenticatedRequest("POST", "/dashboard/search", nil, body, user)
					flash := GetFlashMessage(response.Cookies())

					Expect(flash.Data["success"]).To(BeEmpty())
					Expect(flash.Data["error"]).NotTo(BeEmpty())
				})

				It("redirects to dashboard page", func() {
					user := FabricateUser("dev@nimblehq.co", "password")
					body := GenerateRequestBody(map[string]string{
						"keyword": "",
					})
					response := MakeAuthenticatedRequest("POST", "/dashboard/search", nil, body, user)
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

	Describe("POST /dashboard/upload", func() {
		Context("when the user has already signed in", func() {
			Context("given valid params", func() {
				It("shows a success message", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					headers, body := CreateMultipartRequestInfo("tests/fixtures/files/valid.csv", "text/csv")
					response := MakeAuthenticatedRequest("POST", "/dashboard/upload", headers, body, user)
					flash := GetFlashMessage(response.Cookies())

					Expect(flash.Data["success"]).NotTo(BeEmpty())
					Expect(flash.Data["error"]).To(BeEmpty())
				})

				It("redirects to dashboard page", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					headers, body := CreateMultipartRequestInfo("tests/fixtures/files/valid.csv", "text/csv")
					response := MakeAuthenticatedRequest("POST", "/dashboard/upload", headers, body, user)
					currentPath := GetCurrentPath(response)

					Expect(response.StatusCode).To(Equal(http.StatusFound))
					Expect(currentPath).To(Equal("/dashboard"))
				})
			})

			Context("given INVALID params", func() {
				It("shows an error message", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					headers, body := CreateMultipartRequestInfo("tests/fixtures/files/text.txt", "text/plain")
					response := MakeAuthenticatedRequest("POST", "/dashboard/upload", headers, body, user)
					flash := GetFlashMessage(response.Cookies())

					Expect(flash.Data["success"]).To(BeEmpty())
					Expect(flash.Data["error"]).NotTo(BeEmpty())
				})

				It("redirects to dashboard page", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					headers, body := CreateMultipartRequestInfo("tests/fixtures/files/text.txt", "text/plain")
					response := MakeAuthenticatedRequest("POST", "/dashboard/upload", headers, body, user)
					currentPath := GetCurrentPath(response)

					Expect(response.StatusCode).To(Equal(http.StatusFound))
					Expect(currentPath).To(Equal("/dashboard"))
				})
			})
		})
	})
})
