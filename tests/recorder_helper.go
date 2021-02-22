package tests

import (
	"fmt"
	"net/http"

	"go-crawler-challenge/helpers"

	"github.com/dnaeon/go-vcr/recorder"
	. "github.com/onsi/ginkgo"
)

func RecordCassette(cassetteName string, visitURL string) {
	rec, err := recorder.New(fmt.Sprintf("%s/tests/fixtures/vcr/%s", helpers.RootDir(), cassetteName))
	if err != nil {
		Fail(err.Error())
	}
	defer func() {
		err := rec.Stop()
		if err != nil {
			Fail(err.Error())
		}
	}()

	// Create an HTTP client and inject our transport
	client := &http.Client{Transport: rec}
	_, err = client.Get(visitURL)
	if err != nil {
		Fail(err.Error())
	}
}
