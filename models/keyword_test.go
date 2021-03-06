package models_test

import (
	"fmt"
	"go-crawler-challenge/models"
	. "go-crawler-challenge/tests"
	. "go-crawler-challenge/tests/fixtures"

	"github.com/bxcodec/faker/v3"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Keyword", func() {
	AfterEach(func() {
		TruncateTable("position")
	})

	Describe("#GetAllKeyword", func() {
		Context("given an existing keyword", func() {
			It("returns a keyword record", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				keyword := FabricateKeyword(faker.Word(), "https://www.google.com/search?lr=lang_en", user)
				orderBy := []string{"created_at desc"}
				offset := 0
				limit := 1

				keywords, err := models.GetAllKeyword(orderBy, int64(offset), int64(limit))
				if err != nil {
					Fail(fmt.Sprintf("Get all keywords failed: %v", err.Error()))
				}

				Expect(len(keywords)).To(BeNumerically(">", 0))
				Expect(keywords[0].Id).To(Equal(keyword.Id))
			})
		})

		Context("given NO keyword exists", func() {
			It("does NOT return any keyword record", func() {
				orderBy := []string{"created_at desc"}
				offset := 0
				limit := 1

				keywords, err := models.GetAllKeyword(orderBy, int64(offset), int64(limit))
				if err != nil {
					Fail(fmt.Sprintf("Get all keywords failed: %v", err.Error()))
				}

				Expect(len(keywords)).To(BeNumerically(">", 0))
				Expect(keywords[0].Id).To(BeNil())
			})
		})
	})

	Describe("#CountAllKeyword", func() {
		Context("given a keyword record in the database", func() {
			It("returns a position record", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				_ = FabricateKeyword(faker.Word(), "https://www.google.com/search?lr=lang_en", user)

				totalRows, err := models.CountAllKeyword()
				if err != nil {
					Fail(fmt.Sprintf("Count all keywords failed: %v", err.Error()))
				}

				Expect(totalRows).To(BeNumerically(">", 0))
			})
		})

		Context("given NO keyword exists", func() {
			It("returns an empty array", func() {
				totalRows, err := models.CountAllKeyword()
				if err != nil {
					Fail(fmt.Sprintf("Count all keyword failed: %v", err.Error()))
				}

				Expect(totalRows).To(BeZero())
			})
		})
	})
})
