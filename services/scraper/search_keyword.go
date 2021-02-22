package scraper

import (
	"fmt"
	"net/url"

	"github.com/beego/beego/v2/core/logs"
	"github.com/gocolly/colly/v2"
)

type SearchKeywordService struct {
	Keyword string

	synchronous  bool
	searchResult *searchKeywordResult
}

type searchKeywordResult struct {
	Keyword  string
	VisitURL string
	NonAds   []string
	OtherAds []string
	TopAds   []string
}

var selectorList = map[string]string{
	"nonAds":        "#search .g .yuRUbf > a",
	"bottomLinkAds": "#tadsb .d5oMvf > a",
	"otherAds":      "#rhs .pla-unit a.pla-unit-title-link",
	"topImageAds":   "#tvcap .pla-unit a.pla-unit-title-link",
	"topLinkAds":    "#tads .d5oMvf > a",
}

const collectingURLPattern = "https://www.google.com/search?q=%s&lr=lang_en"

// Call handles making a service call to scrape some data from google search engine with a given keyword.
// It will return an error when the collector can not visit the url.
func (service *SearchKeywordService) Call() {
	collector := colly.NewCollector(colly.Async(true))
	visitURL := fmt.Sprintf(collectingURLPattern, url.QueryEscape(service.Keyword))
	searchResult := searchKeywordResult{Keyword: service.Keyword, VisitURL: visitURL}

	collector.OnResponse(onResponseHandler)
	collector.OnRequest(onRequestHandler)
	collector.OnError(onResponseErrorHandler)

	collector.OnHTML(selectorList["nonAds"], func(element *colly.HTMLElement) {
		searchResult.NonAds = append(searchResult.NonAds, element.Attr("href"))
	})
	collector.OnHTML(selectorList["bottomLinkAds"], func(element *colly.HTMLElement) {
		searchResult.OtherAds = append(searchResult.OtherAds, element.Attr("href"))
	})
	collector.OnHTML(selectorList["otherAds"], func(element *colly.HTMLElement) {
		searchResult.OtherAds = append(searchResult.OtherAds, element.Attr("href"))
	})
	collector.OnHTML(selectorList["topImageAds"], func(element *colly.HTMLElement) {
		searchResult.TopAds = append(searchResult.TopAds, element.Attr("href"))
	})
	collector.OnHTML(selectorList["topLinkAds"], func(element *colly.HTMLElement) {
		searchResult.TopAds = append(searchResult.TopAds, element.Attr("href"))
	})

	collector.OnScraped(func(response *colly.Response) {
		logs.Info(fmt.Sprintf("Search keyword result: %+v", searchResult))

		service.searchResult = &searchResult
	})

	err := collector.Visit(visitURL)
	if err != nil {
		logs.Critical(fmt.Sprintf("Collector visit failed: %v", err))
	}

	// Disable asynchronous when synchronous flag is enabled
	if service.synchronous {
		collector.Wait()
	}
}

func (service *SearchKeywordService) EnableSynchronous() {
	service.synchronous = true
}

func (service *SearchKeywordService) GetSearchResult() *searchKeywordResult {
	return service.searchResult
}
