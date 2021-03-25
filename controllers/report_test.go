package controllers_test

import (
	"fmt"
	"net/http"

	"go-crawler-challenge/models"
	. "go-crawler-challenge/tests"
	. "go-crawler-challenge/tests/fixtures"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ReportController", func() {
	BeforeEach(func() {
		SeedPositionTable()
	})

	AfterEach(func() {
		TruncateTable("link")
		TruncateTable("page")
		TruncateTable("position")
		TruncateTable("keyword")
		TruncateTable("user")
	})

	Describe("GET /report/:key_id", func() {
		Context("when the user has already signed in", func() {
			Context("given a valid report", func() {
				It("renders with status 200", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					keyword := FabricateKeyword(faker.Word(), faker.URL(), models.GetKeywordStatus("completed"), user)
					position := FabricatePosition(faker.Word(), faker.Word(), "normal")
					_ = FabricatePage(faker.Paragraph(), keyword)
					_ = FabricateLink(faker.URL(), keyword, position)

					response := MakeAuthenticatedRequest("GET", fmt.Sprintf("/report/%v", keyword.Id), nil, nil, user)

					Expect(response.StatusCode).To(Equal(http.StatusOK))
				})
			})

			Context("given an INVALID report", func() {
				It("redirects to dashboard page", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					keyword := FabricateKeyword(faker.Word(), faker.URL(), models.GetKeywordStatus("pending"), user)

					response := MakeAuthenticatedRequest("GET", fmt.Sprintf("/report/%v", keyword.Id), nil, nil, user)
					currentPath := GetCurrentPath(response)

					Expect(response.StatusCode).To(Equal(http.StatusFound))
					Expect(currentPath).To(Equal("/dashboard"))
				})
			})
		})

		Context("when the user has NOT signed in yet", func() {
			It("redirects to sign-in page", func() {
				response := MakeRequest("GET", "/report/1", nil)
				currentPath := GetCurrentPath(response)

				Expect(response.StatusCode).To(Equal(http.StatusFound))
				Expect(currentPath).To(Equal("/user/sign_in"))
			})
		})
	})
})
