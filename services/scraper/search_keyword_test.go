package scraper_test

import (
	"fmt"
	"go-crawler-challenge/models"
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
		PreparePositionTable()
	})

	AfterEach(func() {
		TruncateTable("link")
		TruncateTable("page")
		TruncateTable("position")
		TruncateTable("keyword")
		TruncateTable("user")
	})

	Describe("#Call", func() {
		Context("given valid params", func() {
			It("assigns all links", func() {
				positionList, err := models.GetAllPosition()
				if err != nil {
					Fail(fmt.Sprintf("Get all position failed: %v", err.Error()))
				}

				service := scraper.SearchKeywordService{Keyword: "keyword"}
				service.SetPositionList(positionList)
				service.EnableSynchronous()
				service.Run()

				Expect(len(service.GetSearchResult().LinkList)).NotTo(BeZero())
			})
		})
	})
})
