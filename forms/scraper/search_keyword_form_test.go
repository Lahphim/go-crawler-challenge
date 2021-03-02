package forms_test

import (
	form "go-crawler-challenge/forms/scraper"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Scraper/SearchKeywordForm", func() {
	Describe("#Validate", func() {
		Context("given valid params", func() {
			It("does NOT produce any errors", func() {
				form := form.SearchKeywordForm{
					Keyword: "keyword",
				}

				errors := form.Validate()

				Expect(len(errors)).To(BeZero())
			})
		})

		Context("given INVALID params", func() {
			It("produces an error", func() {
				form := form.SearchKeywordForm{
					Keyword: "",
				}

				errors := form.Validate()

				Expect(errors[0].Error()).To(Equal("Keyword can not be empty"))
			})
		})
	})
})
