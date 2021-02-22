package scraper

import (
	"fmt"

	"github.com/beego/beego/v2/core/logs"
	"github.com/gocolly/colly/v2"
)

const currentBrowser = "Chrome/88.0.4324.182"
const currentOs = "Macintosh; Intel Mac OS X 10_15_5"

func onRequestHandler(request *colly.Request) {
	userAgent := fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) %s Safari/537.36", currentOs, currentBrowser)
	request.Headers.Set("User-Agent", userAgent)

	logs.Info(fmt.Sprintf("Visiting: %v", request.URL))
}

func onResponseErrorHandler(response *colly.Response, err error) {
	logs.Critical(fmt.Sprintf("Response failed: [%v][%v] %v", response.StatusCode, response.Request.URL, err))
}

func onResponseHandler(response *colly.Response) {
	logs.Info("HTML status code: %v", response.StatusCode)
}
