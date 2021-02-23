package helpers

import (
	"fmt"
	"math/rand"
)

var uaGens = []func() string{
	genFirefoxUA,
	genChromeUA,
}

var ffVersions = []float32{
	52.0,
	53.0,
	54.0,
	56.0,
	57.0,
	58.0,
	59.0,
	6.0,
	60.0,
	61.0,
	63.0,
}

var chromeVersions = []string{
	"49.0.2623.112",
	"55.0.2883.87",
	"56.0.2924.87",
	"57.0.2987.133",
	"61.0.3163.100",
	"63.0.3239.132",
	"64.0.3282.0",
	"65.0.3325.146",
	"68.0.3440.106",
	"69.0.3497.100",
	"70.0.3538.102",
	"74.0.3729.169",
	"88.0.4324.182",
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
	return uaGens[rand.Intn(len(uaGens))]()
}

// Generates Firefox Browser User-Agent (Desktop)
//	-> "Mozilla/5.0 (Windows NT 10.0; rv:63.0) Gecko/20100101 Firefox/63.0"
func genFirefoxUA() string {
	version := ffVersions[rand.Intn(len(ffVersions))]
	os := osStrings[rand.Intn(len(osStrings))]
	return fmt.Sprintf("Mozilla/5.0 (%s; rv:%.1f) Gecko/20100101 Firefox/%.1f", os, version, version)
}

// Generates Chrome Browser User-Agent (Desktop)
//	-> "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36"
func genChromeUA() string {
	version := chromeVersions[rand.Intn(len(chromeVersions))]
	os := osStrings[rand.Intn(len(osStrings))]
	return fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36", os, version)
}
