package forms_test

import (
	form "go-crawler-challenge/forms/keyword"
	. "go-crawler-challenge/tests"
	. "go-crawler-challenge/tests/fixtures"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Keyword/FileSearchForm", func() {
	Describe("#Save", func() {
		Context("given valid params", func() {
			It("does NOT produce any errors", func() {
				file, fileHeader, err := GetMultipartAttributesFromFile("tests/fixtures/files/valid.csv", "text/csv")
				if err != nil {
					Fail(err.Error())
				}

				user := FabricateUser(faker.Email(), faker.Password())
				fileSearchForm := form.FileSearchForm{
					File:       file,
					FileHeader: fileHeader,
					User:       user,
				}

				err = fileSearchForm.Save()

				Expect(err).To(BeNil())
			})
		})

		Context("given INVALID params", func() {
			Context("given NO file exists", func() {
				It("produces an error", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					fileSearchForm := form.FileSearchForm{
						File:       nil,
						FileHeader: nil,
						User:       user,
					}

					err := fileSearchForm.Save()

					Expect(err).To(Equal("File cannot be empty"))
				})
			})

			Context("given INVALID file type", func() {
				It("produces an error", func() {
					file, fileHeader, err := GetMultipartAttributesFromFile("tests/fixtures/files/text.txt", "text/plain")
					if err != nil {
						Fail(err.Error())
					}

					user := FabricateUser(faker.Email(), faker.Password())
					fileSearchForm := form.FileSearchForm{
						File:       file,
						FileHeader: fileHeader,
						User:       user,
					}

					err = fileSearchForm.Save()

					Expect(err).To(Equal("File type is not allowed"))
				})
			})

			Context("given INVALID keyword size", func() {
				It("produces an error", func() {
					file, fileHeader, err := GetMultipartAttributesFromFile("tests/fixtures/files/invalid.csv", "text/csv")
					if err != nil {
						Fail(err.Error())
					}

					user := FabricateUser(faker.Email(), faker.Password())
					fileSearchForm := form.FileSearchForm{
						File:       file,
						FileHeader: fileHeader,
						User:       user,
					}

					err = fileSearchForm.Save()

					Expect(err).To(Equal("Acceptance keyword size from 1 to 1,000"))
				})
			})

			Context("given NO user object", func() {
				It("produces an error", func() {
					file, fileHeader, err := GetMultipartAttributesFromFile("tests/fixtures/files/valid.csv", "text/csv")
					if err != nil {
						Fail(err.Error())
					}

					fileSearchForm := form.FileSearchForm{
						File:       file,
						FileHeader: fileHeader,
						User:       nil,
					}

					err = fileSearchForm.Save()

					Expect(err).To(Equal("User cannot be empty"))
				})
			})
		})
	})
})
