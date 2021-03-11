package forms_test

import (
	form "go-crawler-challenge/forms/scraper"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Scraper/TextKeywordForm", func() {
	Describe("#Validate", func() {
		Context("given valid params", func() {
			It("does NOT produce any errors", func() {
				form := form.TextKeywordForm{
					Keyword: "keyword",
				}

				errors := form.Validate()

				Expect(len(errors)).To(BeZero())
			})
		})

		Context("given INVALID params", func() {
			It("produces an error", func() {
				form := form.TextKeywordForm{
					Keyword: "",
				}

				errors := form.Validate()

				Expect(errors[0].Error()).To(Equal("Keyword cannot be empty"))
			})
		})
	})
})
