package apiv1controllers_test

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"go-crawler-challenge/models"
	v1serializers "go-crawler-challenge/serializers/v1"
	. "go-crawler-challenge/tests"
	. "go-crawler-challenge/tests/custom_matchers"
	. "go-crawler-challenge/tests/fixtures"

	"github.com/bxcodec/faker/v3"
	"github.com/google/jsonapi"
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

	Describe("GET /api/v1/keywords", func() {
		Context("given a valid access token", func() {
			It("returns status 200", func() {
				oauthClient := FabricateOauthClient(faker.UUIDHyphenated(), faker.Password())
				user := FabricateUser(faker.Email(), faker.Password())
				_ = FabricateKeyword(faker.Word(), faker.URL(), 0, user)
				_ = FabricateKeyword(faker.Word(), faker.URL(), 0, user)

				anotherUser := FabricateUser(faker.Email(), faker.Password())
				_ = FabricateKeyword(faker.Word(), faker.URL(), 0, anotherUser)

				accessToken := FabricatorAccessToken(oauthClient.ID, user.Id)
				headers := http.Header{"Authorization": {fmt.Sprintf("Bearer %v", accessToken)}}

				response := MakeAuthenticatedRequest("GET", "/api/v1/keywords", headers, nil, nil)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})

			It("matches with valid schema", func() {
				oauthClient := FabricateOauthClient(faker.UUIDHyphenated(), faker.Password())
				user := FabricateUser(faker.Email(), faker.Password())
				_ = FabricateKeyword(faker.Word(), faker.URL(), 0, user)
				_ = FabricateKeyword(faker.Word(), faker.URL(), 0, user)

				anotherUser := FabricateUser(faker.Email(), faker.Password())
				_ = FabricateKeyword(faker.Word(), faker.URL(), 0, anotherUser)

				accessToken := FabricatorAccessToken(oauthClient.ID, user.Id)
				headers := http.Header{"Authorization": {fmt.Sprintf("Bearer %v", accessToken)}}

				response := MakeAuthenticatedRequest("GET", "/api/v1/keywords", headers, nil, nil)

				Expect(response).To(MatchJSONSchema("keyword/list/valid"))
			})

			Context("given `p`(page) param is set", func() {
				It("shows only records in page number 2", func() {
					pageAt := 2
					totalRecords := 12
					oauthClient := FabricateOauthClient(faker.UUIDHyphenated(), faker.Password())
					user := FabricateUser(faker.Email(), faker.Password())
					for i := 1; i <= totalRecords; i++ {
						_ = FabricateKeyword(fmt.Sprintf("key%02dword", i), "https://www.sample.com", 0, user)
					}

					accessToken := FabricatorAccessToken(oauthClient.ID, user.Id)
					headers := http.Header{"Authorization": {fmt.Sprintf("Bearer %v", accessToken)}}

					response := MakeAuthenticatedRequest("GET", fmt.Sprintf("/api/v1/keywords?p=%v", pageAt), headers, nil, nil)

					responseKeywords, err := jsonapi.UnmarshalManyPayload(response.Body, reflect.TypeOf(new(v1serializers.KeywordItemResponse)))
					if err != nil {
						Fail(fmt.Sprintf("Unmarshal many payload `KeywordItemResponse` list failed: %v", err.Error()))
					}

					Expect(len(responseKeywords)).To(Equal(2))
					Expect(responseKeywords[0].(*v1serializers.KeywordItemResponse).Keyword).To(Equal("key02word"))
					Expect(responseKeywords[1].(*v1serializers.KeywordItemResponse).Keyword).To(Equal("key01word"))
				})
			})

			Context("given `keyword` param is set", func() {
				It("shows only `keyword` filtered records", func() {
					keywordFilter := "expected_keyword"
					oauthClient := FabricateOauthClient(faker.UUIDHyphenated(), faker.Password())
					user := FabricateUser(faker.Email(), faker.Password())
					// Exact match keyword
					_ = FabricateKeyword(keywordFilter, faker.URL(), 0, user)
					// Fuzzy match keyword
					_ = FabricateKeyword(fmt.Sprintf("%v %v %v", faker.Word(), keywordFilter, faker.Word()), faker.URL(), 0, user)
					// Non-match keyword
					_ = FabricateKeyword("nonexpected__keyword", faker.URL(), 0, user)

					accessToken := FabricatorAccessToken(oauthClient.ID, user.Id)
					headers := http.Header{"Authorization": {fmt.Sprintf("Bearer %v", accessToken)}}

					response := MakeAuthenticatedRequest("GET", fmt.Sprintf("/api/v1/keywords?keyword=%v", keywordFilter), headers, nil, nil)

					responseKeywords, err := jsonapi.UnmarshalManyPayload(response.Body, reflect.TypeOf(new(v1serializers.KeywordItemResponse)))
					if err != nil {
						Fail(fmt.Sprintf("Unmarshal many payload `KeywordItemResponse` list failed: %v", err.Error()))
					}

					Expect(len(responseKeywords)).To(Equal(2))
				})
			})
		})

		Context("given an INVALID access token", func() {
			It("returns status 401", func() {
				accessToken := "INVALID"
				headers := http.Header{"Authorization": {fmt.Sprintf("Bearer %v", accessToken)}}

				response := MakeAuthenticatedRequest("GET", "/api/v1/keywords", headers, nil, nil)

				Expect(response.StatusCode).To(Equal(http.StatusUnauthorized))
			})

			It("matches with INVALID schema", func() {
				accessToken := "INVALID"
				keyword := faker.Word()
				form := url.Values{"keyword": {keyword}}
				headers := http.Header{"Authorization": {fmt.Sprintf("Bearer %v", accessToken)}}
				body := strings.NewReader(form.Encode())

				response := MakeAuthenticatedRequest("GET", "/api/v1/keywords", headers, body, nil)

				Expect(response).To(MatchJSONSchema("keyword/list/invalid"))
			})
		})
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
