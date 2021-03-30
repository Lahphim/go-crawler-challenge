package controllers_test

import (
	"fmt"
	"net/http"

	. "go-crawler-challenge/tests"
	. "go-crawler-challenge/tests/fixtures"

	"github.com/PuerkitoBio/goquery"
	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OauthClientController", func() {
	AfterEach(func() {
		TruncateTable("oauth2_clients")
		TruncateTable("user")
	})

	Describe("GET /oauth_client", func() {
		Context("when the user has already signed in", func() {
			It("renders with status 200", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				response := MakeAuthenticatedRequest("GET", "/oauth_client", nil, nil, user)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})

			Context("given `client_id` param is set", func() {
				Context("given a valid param", func() {
					It("renders with status 200", func() {
						user := FabricateUser(faker.Email(), faker.Password())
						oauthClient := FabricateOauthClient(faker.UUIDHyphenated(), faker.Password())
						response := MakeAuthenticatedRequest("GET", fmt.Sprintf("/oauth_client?client_id=%v", oauthClient.GetID()), nil, nil, user)

						Expect(response.StatusCode).To(Equal(http.StatusOK))
					})

					It("shows the `client_id`", func() {
						user := FabricateUser(faker.Email(), faker.Password())
						clientId := faker.UUIDHyphenated()
						oauthClient := FabricateOauthClient(clientId, faker.Password())
						response := MakeAuthenticatedRequest("GET", fmt.Sprintf("/oauth_client?client_id=%v", oauthClient.GetID()), nil, nil, user)
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
						document.Find("#client_id").Each(func(_ int, selector *goquery.Selection) {
							value, _ := selector.Attr("value")
							matches = append(matches, value)
						})

						Expect(len(matches)).To(BeNumerically("==", 1))
						Expect(matches[0]).To(Equal(clientId))
					})

					It("shows the `client_secret`", func() {
						user := FabricateUser(faker.Email(), faker.Password())
						clientSecret := faker.Password()
						oauthClient := FabricateOauthClient(faker.UUIDHyphenated(), clientSecret)
						response := MakeAuthenticatedRequest("GET", fmt.Sprintf("/oauth_client?client_id=%v", oauthClient.GetID()), nil, nil, user)
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
						document.Find("#client_secret").Each(func(_ int, selector *goquery.Selection) {
							value, _ := selector.Attr("value")
							matches = append(matches, value)
						})

						Expect(len(matches)).To(BeNumerically("==", 1))
						Expect(matches[0]).To(Equal(clientSecret))
					})
				})

				Context("given an INVALID param", func() {
					It("renders with status 200", func() {
						clientID := "INVALID_CLIENT_ID"
						user := FabricateUser(faker.Email(), faker.Password())
						response := MakeAuthenticatedRequest("GET", fmt.Sprintf("/oauth_client?client_id=%v", clientID), nil, nil, user)

						Expect(response.StatusCode).To(Equal(http.StatusOK))
					})

					It("renders an error message", func() {
						clientID := "INVALID_CLIENT_ID"
						user := FabricateUser(faker.Email(), faker.Password())
						response := MakeAuthenticatedRequest("GET", fmt.Sprintf("/oauth_client?client_id=%v", clientID), nil, nil, user)
						flash := GetFlashMessage(response.Cookies())

						Expect(flash.Data["success"]).To(BeEmpty())
						Expect(flash.Data["error"]).NotTo(BeEmpty())
					})

					It("does NOT show the `client_id", func() {
						clientId := "INVALID_CLIENT_ID"
						user := FabricateUser(faker.Email(), faker.Password())
						response := MakeAuthenticatedRequest("GET", fmt.Sprintf("/oauth_client?client_id=%v", clientId), nil, nil, user)
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
						document.Find("#client_id").Each(func(_ int, selector *goquery.Selection) {
							value, _ := selector.Attr("value")
							matches = append(matches, value)
						})

						Expect(len(matches)).To(BeZero())
					})

					It("does NOT show the `client_secret", func() {
						clientId := "INVALID_CLIENT_ID"
						user := FabricateUser(faker.Email(), faker.Password())
						response := MakeAuthenticatedRequest("GET", fmt.Sprintf("/oauth_client?client_id=%v", clientId), nil, nil, user)
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
						document.Find("#client_secret").Each(func(_ int, selector *goquery.Selection) {
							value, _ := selector.Attr("value")
							matches = append(matches, value)
						})

						Expect(len(matches)).To(BeZero())
					})
				})
			})
		})

		Context("when the user has NOT signed in yet", func() {
			It("redirects to sign-in page", func() {
				response := MakeRequest("GET", "/oauth_client", nil)
				currentPath := GetCurrentPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath).To(Equal("/user/sign_in"))
			})
		})
	})

	Describe("POST /oauth_client", func() {
		Context("when the user has already signed in", func() {
			It("shows a success message", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				body := GenerateRequestBody(map[string]string{})
				response := MakeAuthenticatedRequest("POST", "/oauth_client", nil, body, user)
				flash := GetFlashMessage(response.Cookies())

				Expect(flash.Data["success"]).NotTo(BeEmpty())
				Expect(flash.Data["error"]).To(BeEmpty())
			})

			It("redirects to oauth client form page", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				body := GenerateRequestBody(map[string]string{})
				response := MakeAuthenticatedRequest("POST", "/oauth_client", nil, body, user)
				currentPath := GetCurrentPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath).To(Equal("/oauth_client"))
			})
		})

		Context("when the user has NOT signed in yet", func() {
			It("redirects to sign-in page", func() {
				body := GenerateRequestBody(map[string]string{})
				response := MakeRequest("POST", "/oauth_client", body)
				currentPath := GetCurrentPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath).To(Equal("/user/sign_in"))
			})
		})
	})
})
