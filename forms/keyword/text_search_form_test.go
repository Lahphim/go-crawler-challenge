package forms_test

import (
	form "go-crawler-challenge/forms/keyword"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Keyword/TextSearchForm", func() {
	Describe("#Create", func() {
		Context("given valid params", func() {
			It("does NOT produce any errors", func() {
				textSearchForm := form.TextSearchForm{
					Keyword: "keyword",
				}

				err := textSearchForm.Create()

				Expect(err).To(BeNil())
			})
		})

		Context("given INVALID params", func() {
			It("produces an error", func() {
				textSearchForm := form.TextSearchForm{
					Keyword: "",
				}

				err := textSearchForm.Create()

				Expect(err).To(Equal("Keyword cannot be empty"))
			})
		})
	})
})
