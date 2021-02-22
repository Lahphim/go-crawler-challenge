package scraper

import (
	"fmt"
	"net/url"

	"github.com/beego/beego/v2/core/logs"
	"github.com/gocolly/colly/v2"
)

type SearchKeywordService struct {
	Keyword string
}

type searchKeywordResult struct {
	Keyword  string
	VisitURL string
	NonAds   []string
	OtherAds []string
	TopAds   []string
}

var selectorList = map[string]string{
	"nonAds":      "#search .g .yuRUbf > a",
	"otherAds":    "#rhs .pla-unit a.pla-unit-title-link",
	"topImageAds": "#tvcap .pla-unit a.pla-unit-title-link",
	"topLinkAds":  "#tads .d5oMvf > a",
}

const collectingURLPattern = "https://www.google.com/search?q=%s"

func (service *SearchKeywordService) Call() (err error) {
	collector := colly.NewCollector()
	visitURL := service.getVisitURL()
	searchResult := searchKeywordResult{Keyword: service.Keyword, VisitURL: visitURL}

	collector.OnResponse(onResponseHandler)
	collector.OnRequest(onRequestHandler)
	collector.OnError(onResponseErrorHandler)

	collector.OnHTML(selectorList["nonAds"], func(element *colly.HTMLElement) {
		searchResult.NonAds = append(searchResult.NonAds, element.Attr("href"))
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
	})

	err = collector.Visit(visitURL)
	if err != nil {
		logs.Critical(fmt.Sprintf("Collector visit failed: %v", err))
	}

	return err
}

func (service *SearchKeywordService) getVisitURL() (visitURL string) {
	return fmt.Sprintf(collectingURLPattern, url.QueryEscape(service.Keyword))
}