package scraper

import (
	"fmt"
	"net/url"

	form "go-crawler-challenge/forms/scrapper"
	"go-crawler-challenge/models"

	"github.com/beego/beego/v2/core/logs"
	"github.com/gocolly/colly/v2"
)

type SearchKeywordService struct {
	User    *models.User
	Keyword string

	isSynchronous     bool
	positionList      []*models.Position
	keywordResultForm *form.KeywordResultForm
}

const searchEngineUrl = "https://www.google.com/search?q=%s&lr=lang_en"

// Scrape data from the results page of a Google search for a given keyword.
// It will return an error when the collector cannot visit the URL.
func (service *SearchKeywordService) Run() {
	collector := colly.NewCollector(colly.Async(true))
	visitUrl := fmt.Sprintf(searchEngineUrl, url.QueryEscape(service.Keyword))
	keywordResultForm := form.KeywordResultForm{Keyword: service.Keyword}

	collector.OnRequest(onRequestHandler)
	collector.OnError(onResponseErrorHandler)

	for _, position := range service.positionList {
		collector.OnHTML(position.Selector, func(element *colly.HTMLElement) {
			keywordResultForm.LinkList = append(keywordResultForm.LinkList, []string{position.Category, element.Attr("href")})
		})
	}

	collector.OnResponse(func(response *colly.Response) {
		logs.Info(fmt.Sprintf("HTML status code: %v", response.StatusCode))

		keywordResultForm.RawHtml = string(response.Body[:])
	})

	collector.OnScraped(func(response *colly.Response) {
		logs.Info(fmt.Sprintf("Search keyword result: %+v", keywordResultForm))

		service.keywordResultForm = &keywordResultForm
	})

	err := collector.Visit(visitUrl)
	if err != nil {
		logs.Critical(fmt.Sprintf("Collector visit failed: %v", err))
	} else {
		keywordResultForm.VisitUrl = visitUrl
	}

	// Disable asynchronous when synchronous flag is enabled
	if service.isSynchronous {
		collector.Wait()
	}
}

func (service *SearchKeywordService) SetPositionList(positionList []*models.Position) {
	service.positionList = positionList
}

func (service *SearchKeywordService) EnableSynchronous() {
	service.isSynchronous = true
}

func (service *SearchKeywordService) GetSearchResult() *form.KeywordResultForm {
	return service.keywordResultForm
}
