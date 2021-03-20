package helpers_test

import (
	"html/template"
	"time"

	"go-crawler-challenge/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ViewHelper", func() {
	Describe("#HashEmail", func() {
		It("returns hash", func() {
			plainText := "dev@nimblehq.co"
			hashText := helpers.HashEmail(plainText)

			Expect(hashText).To(Equal("6733d09432e89459dba795de8312ac2d"))
		})
	})

	Describe("#ToTimeAgo", func() {
		Describe("given timestamp is less than 30 seconds", func() {
			It("returns `less than a minute ago`", func() {
				currentTime := time.Now()
				timeAgo := helpers.ToTimeAgo(currentTime)

				Expect(timeAgo).To(Equal("less than a minute ago"))
			})
		})

		Describe("given timestamp is over than 30 seconds", func() {
			It("returns `1 minute ago`", func() {
				historyTime := time.Now().Local().Add(-time.Second * 42)
				timeAgo := helpers.ToTimeAgo(historyTime)

				Expect(timeAgo).To(Equal("1 minute ago"))
			})
		})
	})

	Describe("#ToTimeStamp", func() {
		It("returns custom format timestamp", func() {
			currentTime := time.Now()
			formattedTimestamp := helpers.ToTimeStamp(currentTime)

			Expect(formattedTimestamp).To(Equal(currentTime.Local().Format("01/02/2006 | 3:04PM")))
		})
	})

	Describe("#Unescape", func() {
		It("returns unescape HTML", func() {
			rawHtml := "<b>HTML</b>"
			unescapedHtml := helpers.Unescape(rawHtml)

			Expect(unescapedHtml).To(Equal(template.HTML(rawHtml)))
		})
	})
})
