package scraper

import (
	"fmt"
	"net/url"

	"go-crawler-challenge/models"
	"go-crawler-challenge/services/keyword"

	"github.com/beego/beego/v2/core/logs"
	"github.com/gocolly/colly/v2"
)

type SearchKeywordService struct {
	Keyword *models.Keyword

	isSynchronous        bool
	keywordResultService *keyword.CreateKeywordResult
}

const searchEngineUrl = "https://www.google.com/search?q=%s&lr=lang_en"

// Scrape data from the results page of a Google search for a given keyword.
// It will return an error when the collector cannot visit the URL.
func (service *SearchKeywordService) Run() (err error) {
	collector := colly.NewCollector(colly.Async(true))
	searchUrl := fmt.Sprintf(searchEngineUrl, url.QueryEscape(service.Keyword.Keyword))
	keywordResultService := keyword.CreateKeywordResult{Keyword: service.Keyword}

	collector.OnRequest(onRequestHandler)
	collector.OnError(onResponseErrorHandler)

	positionList, err := service.GetPositionList()
	if err != nil {
		logs.Critical(fmt.Sprintf("Get position list failed: %v", err))
		return err
	}
	for _, position := range positionList {
		positionClone := position
		collector.OnHTML(position.Selector, func(element *colly.HTMLElement) {
			keywordResultService.LinkList = append(keywordResultService.LinkList, models.Link{Position: positionClone, Url: element.Attr("href")})
		})
	}

	collector.OnResponse(func(response *colly.Response) {
		keywordResultService.RawHtml = string(response.Body[:])
	})

	collector.OnScraped(func(response *colly.Response) {
		service.keywordResultService = &keywordResultService

		_, err := service.keywordResultService.Run()
		if err != nil {
			logs.Critical(fmt.Sprintf("Save keyword result failed: %v", err.Error()))
		}
	})

	err = collector.Visit(searchUrl)
	if err != nil {
		logs.Critical(fmt.Sprintf("Collector visit failed: %v", err))
		return err
	} else {
		keywordResultService.Url = searchUrl
	}

	// Disable asynchronous when synchronous flag is enabled
	if service.isSynchronous {
		collector.Wait()
	}

	return nil
}

func (service *SearchKeywordService) GetPositionList() (positions []*models.Position, err error) {
	positionList, err := models.GetAllPosition()
	if err != nil {
		return []*models.Position{}, err
	}

	return positionList, nil
}

func (service *SearchKeywordService) EnableSynchronous() {
	service.isSynchronous = true
}

func (service *SearchKeywordService) GetSearchResult() *keyword.CreateKeywordResult {
	return service.keywordResultService
}
