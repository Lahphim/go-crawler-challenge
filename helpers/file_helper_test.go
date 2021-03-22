package helpers_test

import (
	"fmt"
	"go-crawler-challenge/helpers"
	. "go-crawler-challenge/tests"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FileHelper", func() {
	Describe("#CheckMatchFileType", func() {
		Context("given matched file type", func() {
			It("returns true", func() {
				_, fileHeader, err := GetMultipartAttributesFromFile("tests/fixtures/files/valid.csv", "text/csv")
				if err != nil {
					Fail(fmt.Sprintf("Get multipart atrributes from file failed: %v", err))
				}

				expectFileTypes := []string{"text/csv"}
				matchResult := helpers.CheckMatchFileType(fileHeader, expectFileTypes)

				Expect(matchResult).To(Equal(true))
			})
		})

		Context("given unmatched file type", func() {
			It("returns false", func() {
				_, fileHeader, err := GetMultipartAttributesFromFile("tests/fixtures/files/text.txt", "text/txt")
				if err != nil {
					Fail(fmt.Sprintf("Get multipart atrributes from file failed: %v", err))
				}

				expectFileTypes := []string{"text/csv"}
				matchResult := helpers.CheckMatchFileType(fileHeader, expectFileTypes)

				Expect(matchResult).To(Equal(false))
			})
		})
	})

	Describe("#ReadFileContent", func() {
		It("returns list of content", func() {
			file, _, err := GetMultipartAttributesFromFile("tests/fixtures/files/valid.csv", "text/csv")
			if err != nil {
				Fail(fmt.Sprintf("Get multipart atrributes from file failed: %v", err))
			}

			contentList, err := helpers.ReadFileContent(file)
			if err != nil {
				Fail(fmt.Sprintf("Read file content failed: %v", err))
			}

			Expect(len(contentList)).To(BeNumerically(">", 0))
		})
	})
})
