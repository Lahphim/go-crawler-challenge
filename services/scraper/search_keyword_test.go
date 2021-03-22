package scraper_test

import (
	"fmt"
	"net/url"

	"go-crawler-challenge/services/scraper"
	. "go-crawler-challenge/tests"
	. "go-crawler-challenge/tests/fixtures"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Scraper/SearchKeyword", func() {
	BeforeEach(func() {
		keyword := "keyword"
		visitURL := fmt.Sprintf("https://www.google.com/search?q=%s&lr=lang_en", url.QueryEscape(keyword))
		cassetteName := "scraper/success_valid_params"

		RecordCassette(cassetteName, visitURL)
		SeedPositionTable()
	})

	AfterEach(func() {
		TruncateTable("link")
		TruncateTable("page")
		TruncateTable("position")
		TruncateTable("keyword")
		TruncateTable("user")
	})

	Describe("#Run", func() {
		Context("given valid params", func() {
			It("collects keywords", func() {
				currentUser := FabricateUser(faker.Email(), faker.Password())
				keyword := FabricateKeyword("keyword", "https://www.google.com/search?q=keyword&lr=lang_en", 0, currentUser)
				service := scraper.SearchKeywordService{Keyword: keyword}
				service.EnableSynchronous()
				service.Run()

				searchResult := service.GetSearchResult()

				Expect(searchResult.Keyword).NotTo(BeNil())
				Expect(searchResult.Keyword).To(Equal(keyword))
			})

			It("collects some links based on selector list", func() {
				currentUser := FabricateUser(faker.Email(), faker.Password())
				keyword := FabricateKeyword("keyword", "https://www.google.com/search?q=keyword&lr=lang_en", 0, currentUser)
				service := scraper.SearchKeywordService{Keyword: keyword}
				service.EnableSynchronous()
				service.Run()

				Expect(len(service.GetSearchResult().LinkList)).NotTo(BeZero())
			})

			It("collects the raw HTML", func() {
				currentUser := FabricateUser(faker.Email(), faker.Password())
				keyword := FabricateKeyword("keyword", "https://www.google.com/search?q=keyword&lr=lang_en", 0, currentUser)
				service := scraper.SearchKeywordService{Keyword: keyword}
				service.EnableSynchronous()
				service.Run()

				searchResult := service.GetSearchResult()

				Expect(searchResult.RawHtml).NotTo(BeNil())
				Expect(searchResult.RawHtml).To(MatchRegexp(`<\/body>`))
			})
		})
	})
})
