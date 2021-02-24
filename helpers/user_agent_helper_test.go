package helpers_test

import (
	"go-crawler-challenge/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserAgentHelper", func() {
	Describe("#RandomUserAgent", func() {
		It("returns a random DESKTOP browser user-agent", func() {
			userAgent := helpers.RandomUserAgent()

			Expect(len(userAgent)).NotTo(BeZero())
		})

		It("returns user-agent format", func() {
			userAgent := helpers.RandomUserAgent()

			Expect(userAgent).To(MatchRegexp(`(Firefox\/\d{2}.\d|Chrome\/\d{2}.\d.\d{4}.\d{1,3})`))
		})
	})
})
