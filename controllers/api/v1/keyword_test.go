package apiv1controllers_test

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"go-crawler-challenge/models"
	. "go-crawler-challenge/tests"
	. "go-crawler-challenge/tests/custom_matchers"
	. "go-crawler-challenge/tests/fixtures"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("KeywordController", func() {
	AfterEach(func() {
		TruncateTable("oauth2_tokens")
		TruncateTable("oauth2_clients")
		TruncateTable("keyword")
		TruncateTable("user")
	})

	Describe("POST /api/v1/keyword/search", func() {
		Context("given a valid access token", func() {
			Context("given a valid param", func() {
				It("returns status 201", func() {
					oauthClient := FabricateOauthClient(faker.UUIDHyphenated(), faker.Password())
					user := FabricateUser(faker.Email(), faker.Password())
					accessToken := FabricatorAccessToken(oauthClient.ID, user.Id)
					keyword := faker.Word()
					form := url.Values{"keyword": {keyword}}
					headers := http.Header{"Authorization": {fmt.Sprintf("Bearer %v", accessToken)}}
					body := strings.NewReader(form.Encode())

					response := MakeAuthenticatedRequest("POST", "/api/v1/keyword/search", headers, body, nil)

					Expect(response.StatusCode).To(Equal(http.StatusCreated))
				})

				It("creates a new keyword", func() {
					oauthClient := FabricateOauthClient(faker.UUIDHyphenated(), faker.Password())
					user := FabricateUser(faker.Email(), faker.Password())
					accessToken := FabricatorAccessToken(oauthClient.ID, user.Id)
					keyword := faker.Word()
					form := url.Values{"keyword": {keyword}}
					headers := http.Header{"Authorization": {fmt.Sprintf("Bearer %v", accessToken)}}
					body := strings.NewReader(form.Encode())
					queryList := map[string]interface{}{"user_id": user.Id}

					totalRowsBeforeRequest, err := models.CountAllKeyword(queryList)
					if err != nil {
						Fail(fmt.Sprintf("Count all keyword failed: %v", err.Error()))
					}

					_ = MakeAuthenticatedRequest("POST", "/api/v1/keyword/search", headers, body, nil)

					totalRowsAfterRequest, err := models.CountAllKeyword(queryList)
					if err != nil {
						Fail(fmt.Sprintf("Count all keyword failed: %v", err.Error()))
					}

					Expect(totalRowsAfterRequest - totalRowsBeforeRequest).To(Equal(int64(1)))
				})

				It("matches with valid schema", func() {
					oauthClient := FabricateOauthClient(faker.UUIDHyphenated(), faker.Password())
					user := FabricateUser(faker.Email(), faker.Password())
					accessToken := FabricatorAccessToken(oauthClient.ID, user.Id)
					keyword := faker.Word()
					form := url.Values{"keyword": {keyword}}
					headers := http.Header{"Authorization": {fmt.Sprintf("Bearer %v", accessToken)}}
					body := strings.NewReader(form.Encode())

					response := MakeAuthenticatedRequest("POST", "/api/v1/keyword/search", headers, body, nil)

					Expect(response).To(MatchJSONSchema("keyword/search/valid"))
				})
			})

			Context("given an INVALID param", func() {
				It("returns status 500", func() {
					oauthClient := FabricateOauthClient(faker.UUIDHyphenated(), faker.Password())
					user := FabricateUser(faker.Email(), faker.Password())
					accessToken := FabricatorAccessToken(oauthClient.ID, user.Id)
					form := url.Values{"keyword": {""}}
					headers := http.Header{"Authorization": {fmt.Sprintf("Bearer %v", accessToken)}}
					body := strings.NewReader(form.Encode())

					response := MakeAuthenticatedRequest("POST", "/api/v1/keyword/search", headers, body, nil)

					Expect(response.StatusCode).To(Equal(http.StatusInternalServerError))
				})

				It("does NOT create any new keywords", func() {
					oauthClient := FabricateOauthClient(faker.UUIDHyphenated(), faker.Password())
					user := FabricateUser(faker.Email(), faker.Password())
					accessToken := FabricatorAccessToken(oauthClient.ID, user.Id)
					form := url.Values{"keyword": {""}}
					headers := http.Header{"Authorization": {fmt.Sprintf("Bearer %v", accessToken)}}
					body := strings.NewReader(form.Encode())
					queryList := map[string]interface{}{"user_id": user.Id}

					totalRowsBeforeRequest, err := models.CountAllKeyword(queryList)
					if err != nil {
						Fail(fmt.Sprintf("Count all keyword failed: %v", err.Error()))
					}

					_ = MakeAuthenticatedRequest("POST", "/api/v1/keyword/search", headers, body, nil)

					totalRowsAfterRequest, err := models.CountAllKeyword(queryList)
					if err != nil {
						Fail(fmt.Sprintf("Count all keyword failed: %v", err.Error()))
					}

					Expect(totalRowsAfterRequest - totalRowsBeforeRequest).To(Equal(int64(0)))
				})

				It("matches with INVALID schema", func() {
					oauthClient := FabricateOauthClient(faker.UUIDHyphenated(), faker.Password())
					user := FabricateUser(faker.Email(), faker.Password())
					accessToken := FabricatorAccessToken(oauthClient.ID, user.Id)
					form := url.Values{"keyword": {""}}
					headers := http.Header{"Authorization": {fmt.Sprintf("Bearer %v", accessToken)}}
					body := strings.NewReader(form.Encode())

					response := MakeAuthenticatedRequest("POST", "/api/v1/keyword/search", headers, body, nil)

					Expect(response).To(MatchJSONSchema("keyword/search/invalid"))
				})
			})
		})

		Context("given an INVALID access token", func() {
			It("returns status 401", func() {
				accessToken := "INVALID"
				keyword := faker.Word()
				form := url.Values{"keyword": {keyword}}
				headers := http.Header{"Authorization": {fmt.Sprintf("Bearer %v", accessToken)}}
				body := strings.NewReader(form.Encode())

				response := MakeAuthenticatedRequest("POST", "/api/v1/keyword/search", headers, body, nil)

				Expect(response.StatusCode).To(Equal(http.StatusUnauthorized))
			})

			It("matches with INVALID schema", func() {
				accessToken := "INVALID"
				keyword := faker.Word()
				form := url.Values{"keyword": {keyword}}
				headers := http.Header{"Authorization": {fmt.Sprintf("Bearer %v", accessToken)}}
				body := strings.NewReader(form.Encode())

				response := MakeAuthenticatedRequest("POST", "/api/v1/keyword/search", headers, body, nil)

				Expect(response).To(MatchJSONSchema("keyword/search/invalid"))
			})
		})
	})
})
