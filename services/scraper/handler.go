package scraper

import (
	"fmt"

	"go-crawler-challenge/helpers"

	"github.com/beego/beego/v2/core/logs"
	"github.com/gocolly/colly/v2"
)

func onRequestHandler(request *colly.Request) {
	userAgent := helpers.RandomUserAgent()
	request.Headers.Set("User-Agent", userAgent)

	logs.Info(fmt.Sprintf("Visiting: %v", request.URL))
}

func onResponseErrorHandler(response *colly.Response, err error) {
	logs.Critical(fmt.Sprintf("Response failed: [%v][%v] %v", response.StatusCode, response.Request.URL, err))
}

func onResponseHandler(response *colly.Response) {
	logs.Info("HTML status code: %v", response.StatusCode)
}
