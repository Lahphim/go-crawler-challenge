package scraper_test

import (
	"fmt"
	"net/url"

	"go-crawler-challenge/services/scraper"
	. "go-crawler-challenge/tests"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Scraper/SearchKeyword", func() {
	BeforeEach(func() {
		keyword := "keyword"
		visitURL := fmt.Sprintf("https://www.google.com/search?q=%s&lr=lang_en", url.QueryEscape(keyword))
		cassetteName := "scraper/success_valid_params"

		RecordCassette(cassetteName, visitURL)
	})

	Describe("#Call", func() {
		Context("given valid params", func() {
			It("assigns non-AdWords", func() {
				service := scraper.SearchKeywordService{Keyword: "keyword"}
				service.EnableSynchronous()
				service.Run()

				Expect(len(service.GetSearchResult().NonAds)).NotTo(BeZero())
			})
		})
	})
})
