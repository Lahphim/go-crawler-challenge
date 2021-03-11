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
		TruncateTable("keyword")
		TruncateTable("user")
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

			Context("given order by `id` in ascending order", func() {
				It("orders the first record to the top of the list", func() {
					orderBy := []string{"id asc"}

					user := FabricateUser(faker.Email(), faker.Password())
					firstKeyword := FabricateKeyword(faker.Word(), "https://www.google.com/search?lr=lang_en", user)
					secondKeyword := FabricateKeyword(faker.Word(), "https://www.google.com/search?lr=lang_en", user)
					offset := 0
					limit := 2

					keywords, err := models.GetAllKeyword(orderBy, int64(offset), int64(limit))
					if err != nil {
						Fail(fmt.Sprintf("Get all keywords failed: %v", err.Error()))
					}

					Expect(keywords[0].Id).To(Equal(firstKeyword.Id))
					Expect(keywords[1].Id).To(Equal(secondKeyword.Id))
				})
			})

			Context("given order by `id` in descending order", func() {
				It("orders the first record to the end of the list", func() {
					orderBy := []string{"id desc"}

					user := FabricateUser(faker.Email(), faker.Password())
					firstKeyword := FabricateKeyword(faker.Word(), "https://www.google.com/search?lr=lang_en", user)
					secondKeyword := FabricateKeyword(faker.Word(), "https://www.google.com/search?lr=lang_en", user)
					offset := 0
					limit := 2

					keywords, err := models.GetAllKeyword(orderBy, int64(offset), int64(limit))
					if err != nil {
						Fail(fmt.Sprintf("Get all keywords failed: %v", err.Error()))
					}

					Expect(keywords[0].Id).NotTo(Equal(firstKeyword.Id))
					Expect(keywords[0].Id).To(Equal(secondKeyword.Id))
				})
			})

			Context("given number of offset is 0", func() {
				It("shows the first insert keyword", func() {
					offset := 0

					user := FabricateUser(faker.Email(), faker.Password())
					firstKeyword := FabricateKeyword(faker.Word(), "https://www.google.com/search?lr=lang_en", user)
					secondKeyword := FabricateKeyword(faker.Word(), "https://www.google.com/search?lr=lang_en", user)
					orderBy := []string{"created_at asc"}
					limit := 1

					keywords, err := models.GetAllKeyword(orderBy, int64(offset), int64(limit))
					if err != nil {
						Fail(fmt.Sprintf("Get all keywords failed: %v", err.Error()))
					}

					Expect(keywords[0].Id).To(Equal(firstKeyword.Id))
					Expect(keywords[0].Id).NotTo(Equal(secondKeyword.Id))
				})
			})

			Context("given number of offset is 1", func() {
				It("does NOT show the first insert keyword", func() {
					offset := 1

					user := FabricateUser(faker.Email(), faker.Password())
					firstKeyword := FabricateKeyword(faker.Word(), "https://www.google.com/search?lr=lang_en", user)
					secondKeyword := FabricateKeyword(faker.Word(), "https://www.google.com/search?lr=lang_en", user)
					orderBy := []string{"created_at asc"}
					limit := 1

					keywords, err := models.GetAllKeyword(orderBy, int64(offset), int64(limit))
					if err != nil {
						Fail(fmt.Sprintf("Get all keywords failed: %v", err.Error()))
					}

					Expect(keywords[0].Id).NotTo(Equal(firstKeyword.Id))
					Expect(keywords[0].Id).To(Equal(secondKeyword.Id))
				})
			})

			Context("given number of limit data selection is 2", func() {
				It("returns only 2 records", func() {
					limit := 2

					user := FabricateUser(faker.Email(), faker.Password())
					_ = FabricateKeyword(faker.Word(), "https://www.google.com/search?lr=lang_en", user)
					_ = FabricateKeyword(faker.Word(), "https://www.google.com/search?lr=lang_en", user)
					_ = FabricateKeyword(faker.Word(), "https://www.google.com/search?lr=lang_en", user)
					orderBy := []string{"id asc"}
					offset := 0

					keywords, err := models.GetAllKeyword(orderBy, int64(offset), int64(limit))
					if err != nil {
						Fail(fmt.Sprintf("Get all keywords failed: %v", err.Error()))
					}

					Expect(len(keywords)).To(Equal(limit))
				})
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

				Expect(len(keywords)).To(BeZero())
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
