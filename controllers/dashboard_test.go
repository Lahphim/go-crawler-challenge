package controllers_test

import (
	"net/http"

	. "go-crawler-challenge/tests"
	. "go-crawler-challenge/tests/fixtures"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DashboardController", func() {
	BeforeEach(func() {
		prepareFabricator()
	})

	AfterEach(func() {
		TruncateTable("user")
		TruncateTable("position")
	})

	Describe("GET /dashboard", func() {
		Context("when the user has already signed in", func() {
			It("renders with status 200", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				response := MakeAuthenticatedRequest("GET", "/dashboard", nil, user)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
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
})

func prepareFabricator() {
	FabricatePosition("nonAds", "#search .g .yuRUbf > a", "normal")
	FabricatePosition("bottomLinkAds", "#tadsb .d5oMvf > a", "other")
	FabricatePosition("otherAds", "#rhs .pla-unit a.pla-unit-title-link", "other")
	FabricatePosition("topImageAds", "#tvcap .pla-unit a.pla-unit-title-link", "top")
	FabricatePosition("topLinkAds", "#tads .d5oMvf > a", "top")
}
