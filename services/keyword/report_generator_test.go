package keyword_test

import (
	"fmt"

	"go-crawler-challenge/models"
	service "go-crawler-challenge/services/keyword"
	. "go-crawler-challenge/tests"
	. "go-crawler-challenge/tests/fixtures"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Keyword/ReportGenerator", func() {
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

	Describe("#Generate", func() {
		Context("given a valid report", func() {
			It("returns report details", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				keyword := FabricateKeyword(faker.Word(), faker.URL(), 0, user)
				position := FabricatePosition(faker.Word(), faker.Word(), "normal")
				page := FabricatePage(faker.Paragraph(), keyword)
				link := FabricateLink(faker.URL(), keyword, position)

				reportGeneratorService := service.ReportGenerator{Keyword: keyword}
				reportResult, err := reportGeneratorService.Generate()
				if err != nil {
					Fail(fmt.Sprintf("Generate report failed: %v", err.Error()))
				} else {
					reportInterface := reportResult.(models.Report)

					Expect(reportInterface.Keyword).To(Equal(keyword.Keyword))
					Expect(reportInterface.Url).To(Equal(keyword.Url))
					Expect(reportInterface.RawHtml).To(Equal(page.RawHtml))
					Expect(reportInterface.LinkList["normal"][0]).To(Equal(link.Url))
				}

				Expect(reportResult).NotTo(BeNil())
			})
		})

		Context("given an INVALID report", func() {
			It("returns `nil` with an error message", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				keyword := FabricateKeyword(faker.Word(), faker.URL(), 0, user)

				reportGeneratorService := service.ReportGenerator{Keyword: keyword}
				reportResult, err := reportGeneratorService.Generate()

				Expect(reportResult).To(BeNil())
				Expect(err.Error()).NotTo(BeNil())
			})
		})
	})
})
