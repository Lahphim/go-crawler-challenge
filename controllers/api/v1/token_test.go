package apiv1controllers_test

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	v1serializers "go-crawler-challenge/serializers/v1"
	. "go-crawler-challenge/tests"
	. "go-crawler-challenge/tests/custom_matchers"
	. "go-crawler-challenge/tests/fixtures"

	"github.com/bxcodec/faker/v3"
	"github.com/google/jsonapi"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TokenController", func() {
	AfterEach(func() {
		TruncateTable("oauth2_tokens")
		TruncateTable("oauth2_clients")
		TruncateTable("user")
	})

	Describe("POST /api/v1/oauth/token", func() {
		Context("given valid params", func() {
			It("returns status 200", func() {
				oauthClient := FabricateOauthClient(faker.UUIDHyphenated(), faker.Password())
				password := faker.Password()
				user := FabricateUser(faker.Email(), password)
				form := url.Values{
					"client_id":     {oauthClient.ID},
					"client_secret": {oauthClient.Secret},
					"grant_type":    {"password"},
					"username":      {user.Email},
					"password":      {password},
				}
				body := strings.NewReader(form.Encode())

				response := MakeRequest("POST", "/api/v1/oauth/token", body)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})

			It("returns token information", func() {
				oauthClient := FabricateOauthClient(faker.UUIDHyphenated(), faker.Password())
				password := faker.Password()
				user := FabricateUser(faker.Email(), password)
				form := url.Values{
					"client_id":     {oauthClient.ID},
					"client_secret": {oauthClient.Secret},
					"grant_type":    {"password"},
					"username":      {user.Email},
					"password":      {password},
				}
				body := strings.NewReader(form.Encode())

				response := MakeRequest("POST", "/api/v1/oauth/token", body)

				var responseToken v1serializers.TokenInformationResponse
				err := jsonapi.UnmarshalPayload(response.Body, &responseToken)
				if err != nil {
					Fail(fmt.Sprintf("Unmarshal payload `TokenInfo` failed: %v", err.Error()))
				}

				Expect(len(responseToken.AccessToken)).To(BeNumerically(">", 0))
				Expect(len(responseToken.RefreshToken)).To(BeNumerically(">", 0))
				Expect(len(responseToken.AccessToken)).To(BeNumerically(">", 0))
				Expect(responseToken.TokenType).To(Equal("Bearer"))
			})

			It("matches with valid schema", func() {
				oauthClient := FabricateOauthClient(faker.UUIDHyphenated(), faker.Password())
				password := faker.Password()
				user := FabricateUser(faker.Email(), password)
				form := url.Values{
					"client_id":     {oauthClient.ID},
					"client_secret": {oauthClient.Secret},
					"grant_type":    {"password"},
					"username":      {user.Email},
					"password":      {password},
				}
				body := strings.NewReader(form.Encode())

				response := MakeRequest("POST", "/api/v1/oauth/token", body)

				Expect(response).To(MatchJSONSchema("oauth/token/valid"))
			})
		})

		Context("given INVALID params", func() {
			Context("given NO client credentials exist", func() {
				It("returns status 401", func() {
					password := faker.Password()
					user := FabricateUser(faker.Email(), password)
					form := url.Values{
						"client_id":     {""},
						"client_secret": {""},
						"grant_type":    {"password"},
						"username":      {user.Email},
						"password":      {password},
					}
					body := strings.NewReader(form.Encode())

					response := MakeRequest("POST", "/api/v1/oauth/token", body)

					Expect(response.StatusCode).To(Equal(http.StatusUnauthorized))
				})

				It("returns error information", func() {
					password := faker.Password()
					user := FabricateUser(faker.Email(), password)
					form := url.Values{
						"client_id":     {""},
						"client_secret": {""},
						"grant_type":    {"password"},
						"username":      {user.Email},
						"password":      {password},
					}
					body := strings.NewReader(form.Encode())

					response := MakeRequest("POST", "/api/v1/oauth/token", body)

					Expect(response).To(MatchJSONSchema("oauth/token/invalid"))
				})
			})

			Context("given INVALID client credentials", func() {
				It("returns status 401", func() {
					oauthClient := FabricateOauthClient(faker.UUIDHyphenated(), faker.Password())
					password := faker.Password()
					user := FabricateUser(faker.Email(), password)
					form := url.Values{
						"client_id":     {oauthClient.ID},
						"client_secret": {"INVALID"},
						"grant_type":    {"password"},
						"username":      {user.Email},
						"password":      {password},
					}
					body := strings.NewReader(form.Encode())

					response := MakeRequest("POST", "/api/v1/oauth/token", body)

					Expect(response.StatusCode).To(Equal(http.StatusUnauthorized))
				})

				It("returns error information", func() {
					oauthClient := FabricateOauthClient(faker.UUIDHyphenated(), faker.Password())
					password := faker.Password()
					user := FabricateUser(faker.Email(), password)
					form := url.Values{
						"client_id":     {oauthClient.ID},
						"client_secret": {"INVALID"},
						"grant_type":    {"password"},
						"username":      {user.Email},
						"password":      {password},
					}
					body := strings.NewReader(form.Encode())

					response := MakeRequest("POST", "/api/v1/oauth/token", body)

					Expect(response).To(MatchJSONSchema("oauth/token/invalid"))
				})
			})

			Context("given NO grant type exists", func() {
				It("returns status 401", func() {
					oauthClient := FabricateOauthClient(faker.UUIDHyphenated(), faker.Password())
					password := faker.Password()
					user := FabricateUser(faker.Email(), password)
					form := url.Values{
						"client_id":     {oauthClient.ID},
						"client_secret": {oauthClient.Secret},
						"grant_type":    {""},
						"username":      {user.Email},
						"password":      {password},
					}
					body := strings.NewReader(form.Encode())

					response := MakeRequest("POST", "/api/v1/oauth/token", body)

					Expect(response.StatusCode).To(Equal(http.StatusUnauthorized))
				})

				It("returns error information", func() {
					oauthClient := FabricateOauthClient(faker.UUIDHyphenated(), faker.Password())
					password := faker.Password()
					user := FabricateUser(faker.Email(), password)
					form := url.Values{
						"client_id":     {oauthClient.ID},
						"client_secret": {oauthClient.Secret},
						"grant_type":    {""},
						"username":      {user.Email},
						"password":      {password},
					}
					body := strings.NewReader(form.Encode())

					response := MakeRequest("POST", "/api/v1/oauth/token", body)

					Expect(response).To(MatchJSONSchema("oauth/token/invalid"))
				})
			})

			Context("given INVALID user credentials", func() {
				It("returns status 401", func() {
					oauthClient := FabricateOauthClient(faker.UUIDHyphenated(), faker.Password())
					user := FabricateUser(faker.Email(), faker.Password())
					form := url.Values{
						"client_id":     {oauthClient.ID},
						"client_secret": {oauthClient.Secret},
						"grant_type":    {"password"},
						"username":      {user.Email},
						"password":      {"INVALID"},
					}
					body := strings.NewReader(form.Encode())

					response := MakeRequest("POST", "/api/v1/oauth/token", body)

					Expect(response.StatusCode).To(Equal(http.StatusUnauthorized))
				})

				It("returns error information", func() {
					oauthClient := FabricateOauthClient(faker.UUIDHyphenated(), faker.Password())
					user := FabricateUser(faker.Email(), faker.Password())
					form := url.Values{
						"client_id":     {oauthClient.ID},
						"client_secret": {oauthClient.Secret},
						"grant_type":    {"password"},
						"username":      {user.Email},
						"password":      {"INVALID"},
					}
					body := strings.NewReader(form.Encode())

					response := MakeRequest("POST", "/api/v1/oauth/token", body)

					Expect(response).To(MatchJSONSchema("oauth/token/invalid"))
				})
			})
		})
	})
})
