package forms_test

import (
	form "go-crawler-challenge/forms/keyword"
	. "go-crawler-challenge/tests/fixtures"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Keyword/TextSearchForm", func() {
	Describe("#Create", func() {
		Context("given valid params", func() {
			It("does NOT produce any errors", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				textSearchForm := form.TextSearchForm{
					Keyword: "keyword",
					User:    user,
				}

				err := textSearchForm.Create()

				Expect(err).To(BeNil())
			})
		})

		Context("given INVALID params", func() {
			Context("given INVALID keyword", func() {
				It("produces an error", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					textSearchForm := form.TextSearchForm{
						Keyword: "",
						User:    user,
					}

					err := textSearchForm.Create()

					Expect(err.Error()).To(Equal("Keyword cannot be empty"))
				})
			})

			Context("given NO user object", func() {
				It("produces an error", func() {
					textSearchForm := form.TextSearchForm{
						Keyword: "keyword",
						User:    nil,
					}

					err := textSearchForm.Create()

					Expect(err.Error()).To(Equal("User cannot be empty"))
				})
			})
		})
	})
})
