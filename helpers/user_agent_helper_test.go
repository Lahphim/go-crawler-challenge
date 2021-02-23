package helpers_test

import (
	"go-crawler-challenge/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserAgentHelper", func() {
	Describe("#RandomUserAgent", func() {
		userAgent := helpers.RandomUserAgent()

		Expect(len(userAgent)).NotTo(BeZero())
	})
})
