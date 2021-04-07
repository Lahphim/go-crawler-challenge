package apiv1controllers_test

import (
	"fmt"
	"net/http"

	"go-crawler-challenge/models"
	. "go-crawler-challenge/tests"
	. "go-crawler-challenge/tests/custom_matchers"
	. "go-crawler-challenge/tests/fixtures"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ReportController", func() {
	AfterEach(func() {
		TruncateTable("oauth2_tokens")
		TruncateTable("oauth2_clients")
		TruncateTable("link")
		TruncateTable("page")
		TruncateTable("position")
		TruncateTable("keyword")
		TruncateTable("user")
	})

	Describe("GET /api/v1/report/:keyword_id", func() {
		Context("given a valid access token", func() {
			Context("given a valid keyword ID", func() {
				It("returns status 200", func() {
					oauthClient := FabricateOauthClient(faker.UUIDHyphenated(), faker.Password())
					user := FabricateUser(faker.Email(), faker.Password())
					keyword := FabricateKeyword(faker.Word(), faker.URL(), models.GetKeywordStatus("completed"), user)
					position := FabricatePosition(faker.Word(), faker.Word(), "normal")
					_ = FabricatePage(faker.Paragraph(), keyword)
					_ = FabricateLink(faker.URL(), keyword, position)

					accessToken := FabricatorAccessToken(oauthClient.ID, user.Id)
					headers := http.Header{"Authorization": {fmt.Sprintf("Bearer %v", accessToken)}}

					response := MakeAuthenticatedRequest("GET", fmt.Sprintf("/api/v1/report/%v", keyword.Id), headers, nil, nil)

					Expect(response.StatusCode).To(Equal(http.StatusOK))
				})

				It("matches with valid schema", func() {
					oauthClient := FabricateOauthClient(faker.UUIDHyphenated(), faker.Password())
					user := FabricateUser(faker.Email(), faker.Password())
					keyword := FabricateKeyword(faker.Word(), faker.URL(), models.GetKeywordStatus("completed"), user)
					position := FabricatePosition(faker.Word(), faker.Word(), "normal")
					_ = FabricatePage(faker.Paragraph(), keyword)
					_ = FabricateLink(faker.URL(), keyword, position)

					accessToken := FabricatorAccessToken(oauthClient.ID, user.Id)
					headers := http.Header{"Authorization": {fmt.Sprintf("Bearer %v", accessToken)}}

					response := MakeAuthenticatedRequest("GET", fmt.Sprintf("/api/v1/report/%v", keyword.Id), headers, nil, nil)

					Expect(response).To(MatchJSONSchema("report/show/valid"))
				})
			})

			Context("given an INVALID keyword ID", func() {
				It("returns status 404", func() {
					oauthClient := FabricateOauthClient(faker.UUIDHyphenated(), faker.Password())
					user := FabricateUser(faker.Email(), faker.Password())
					keyword := FabricateKeyword(faker.Word(), faker.URL(), models.GetKeywordStatus("pending"), user)

					accessToken := FabricatorAccessToken(oauthClient.ID, user.Id)
					headers := http.Header{"Authorization": {fmt.Sprintf("Bearer %v", accessToken)}}

					response := MakeAuthenticatedRequest("GET", fmt.Sprintf("/api/v1/report/%v", keyword.Id), headers, nil, nil)

					Expect(response.StatusCode).To(Equal(http.StatusNotFound))
				})

				It("matches with INVALID schema", func() {
					oauthClient := FabricateOauthClient(faker.UUIDHyphenated(), faker.Password())
					user := FabricateUser(faker.Email(), faker.Password())
					keyword := FabricateKeyword(faker.Word(), faker.URL(), models.GetKeywordStatus("pending"), user)

					accessToken := FabricatorAccessToken(oauthClient.ID, user.Id)
					headers := http.Header{"Authorization": {fmt.Sprintf("Bearer %v", accessToken)}}

					response := MakeAuthenticatedRequest("GET", fmt.Sprintf("/api/v1/report/%v", keyword.Id), headers, nil, nil)

					Expect(response).To(MatchJSONSchema("report/show/invalid"))
				})
			})
		})

		Context("given an INVALID access token", func() {
			It("returns status 401", func() {
				accessToken := "INVALID"
				headers := http.Header{"Authorization": {fmt.Sprintf("Bearer %v", accessToken)}}

				response := MakeAuthenticatedRequest("GET", "/api/v1/report/0", headers, nil, nil)

				Expect(response.StatusCode).To(Equal(http.StatusUnauthorized))
			})

			It("matches with INVALID schema", func() {
				accessToken := "INVALID"
				headers := http.Header{"Authorization": {fmt.Sprintf("Bearer %v", accessToken)}}

				response := MakeAuthenticatedRequest("GET", "/api/v1/report/0", headers, nil, nil)

				Expect(response).To(MatchJSONSchema("report/show/invalid"))
			})
		})
	})
})
