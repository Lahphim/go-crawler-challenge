package tests

import (
	"fmt"
	"net/http"

	"github.com/dnaeon/go-vcr/recorder"
	. "github.com/onsi/ginkgo"
)

func RecordCassette(cassetteName string, visitURL string) {
	rec, err := recorder.New(fmt.Sprintf("fixtures/vcr/%s", cassetteName))
	if err != nil {
		Fail(err.Error())
	}
	defer rec.Stop()

	// Create an HTTP client and inject our transport
	client := &http.Client{Transport: rec}
	_, err = client.Get(visitURL)
	if err != nil {
		Fail(err.Error())
	}
}
