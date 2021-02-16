package test_helpers

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

// GetFlashMessage gets Beego flash message from http cookie
func GetFlashMessage(cookies []*http.Cookie) *web.FlashData {
	flash := web.NewFlash()

	for _, cookie := range cookies {
		if cookie.Name == "BEEGO_FLASH" {
			decodedCookie := decodeQueryString(cookie.Value)
			// Trim null character out of the decoded cookie value
			trimmedCookie := strings.Trim(decodedCookie, "\x00")
			cookieParts := strings.Split(trimmedCookie, "#BEEGOFLASH#")
			if len(cookieParts) == 2 {
				flash.Data[cookieParts[0]] = cookieParts[1]
			}
		}
	}

	return flash
}

// decodeQueryString decodes query string to normal string
func decodeQueryString(encodedString string) string {
	decodedString, err := url.QueryUnescape(encodedString)
	if err != nil {
		logs.Critical(fmt.Sprintf("Decode query string failed: %v", err))

		return ""
	}

	return decodedString
}
