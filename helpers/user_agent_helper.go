package helpers

import (
	"fmt"
	"math/rand"
	"strings"
)

var platformVersionList = []string{
	"ff_52.0",
	"ff_53.0",
	"ff_54.0",
	"ff_56.0",
	"ff_57.0",
	"ff_58.0",
	"ff_59.0",
	"ff_60.0",
	"ff_61.0",
	"ff_63.0",
	"gc_49.0.2623.112",
	"gc_55.0.2883.87",
	"gc_56.0.2924.87",
	"gc_57.0.2987.133",
	"gc_61.0.3163.100",
	"gc_63.0.3239.132",
	"gc_64.0.3282.0",
	"gc_65.0.3325.146",
	"gc_68.0.3440.106",
	"gc_69.0.3497.100",
	"gc_70.0.3538.102",
	"gc_74.0.3729.169",
	"gc_88.0.4324.182",
}

var osStrings = []string{
	"Macintosh; Intel Mac OS X 10_15_5",
	"Macintosh; Intel Mac OS X 10_10",
	"Windows NT 10.0",
	"Windows NT 5.1",
	"Windows NT 6.1; WOW64",
	"Windows NT 6.1; Win64; x64",
	"X11; Linux x86_64",
}

// RandomUserAgent generates a random DESKTOP browser user-agent on every requests
func RandomUserAgent() string {
	os := osStrings[rand.Intn(len(osStrings))]
	version := strings.Split(platformVersionList[rand.Intn(len(platformVersionList))], "_")

	switch version[0] {
	case "ff":
		// Firefox Browser User-Agent (Desktop)
		//	-> "Mozilla/5.0 (Windows NT 10.0; rv:63.0) Gecko/20100101 Firefox/63.0"
		return fmt.Sprintf("Mozilla/5.0 (%s; rv:%s) Gecko/20100101 Firefox/%s", os, version[1], version[1])
	case "gc":
		// Google Chrome Browser User-Agent (Desktop)
		//	-> "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36"
		return fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36", os, version[1])
	default:
		return ""
	}
}
