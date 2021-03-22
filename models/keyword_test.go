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

	Describe("#AddKeyword", func() {
		Context("given a valid keyword", func() {
			It("returns the keyword ID", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				keyword := models.Keyword{
					Keyword: faker.Word(),
					Url:     faker.URL(),
					User:    user,
				}

				id, err := models.AddKeyword(&keyword)
				if err != nil {
					Fail(fmt.Sprintf("Add a new keyword failed: %v", err))
				}

				Expect(id).To(BeNumerically(">", 0))
			})

			It("does NOT return any errors", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				keyword := models.Keyword{
					Keyword: faker.Word(),
					Url:     faker.URL(),
					User:    user,
				}

				_, err := models.AddKeyword(&keyword)

				Expect(err).To(BeNil())
			})
		})

		Context("given an INVALID keyword", func() {
			Context("given NO user assigns", func() {
				It("returns 0", func() {
					keyword := models.Keyword{
						Keyword: faker.Word(),
						Url:     faker.URL(),
						User:    nil,
					}

					id, _ := models.AddKeyword(&keyword)

					Expect(id).To(BeZero())
				})

				It("returns an error", func() {
					keyword := models.Keyword{
						Keyword: faker.Word(),
						Url:     faker.URL(),
						User:    nil,
					}

					_, err := models.AddKeyword(&keyword)

					Expect(err.Error()).NotTo(BeNil())
				})
			})
		})
	})

	Describe("#GetAllKeyword", func() {
		Context("given an existing keyword", func() {
			It("returns a keyword record", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				keyword := FabricateKeyword(faker.Word(), "https://www.google.com/search?lr=lang_en", 0, user)
				orderBy := []string{"created_at desc"}
				offset := 0
				limit := 1

				var queryList map[string]interface{}

				keywords, err := models.GetAllKeyword(queryList, orderBy, int64(offset), int64(limit))
				if err != nil {
					Fail(fmt.Sprintf("Get all keywords failed: %v", err.Error()))
				}

				Expect(len(keywords)).To(BeNumerically(">", 0))
				Expect(keywords[0].Id).To(Equal(keyword.Id))
			})

			Context("given an operator with case insensitive contain `keyword`", func() {
				It("returns only `keyword` filter records", func() {
					containKeyword := "keyword_keyword"
					var queryList = map[string]interface{}{"keyword__icontains": containKeyword}

					user := FabricateUser(faker.Email(), faker.Password())
					firstKeyword := FabricateKeyword(faker.Word(), faker.URL(), 0, user)
					secondKeyword := FabricateKeyword(faker.Word(), faker.URL(), 0, user)
					thirdKeyword := FabricateKeyword(fmt.Sprintf("%v %v %v", faker.Word(), containKeyword, faker.Word()), faker.URL(), 0, user)

					keywords, err := models.GetAllKeyword(queryList, []string{}, 0, 10)
					if err != nil {
						Fail(fmt.Sprintf("Get all keywords failed: %v", err.Error()))
					}

					Expect(keywords[0].Id).NotTo(Equal(firstKeyword.Id))
					Expect(keywords[0].Id).NotTo(Equal(secondKeyword.Id))
					Expect(keywords[0].Id).To(Equal(thirdKeyword.Id))
				})
			})

			Context("given order by `id` in ascending order", func() {
				It("orders the first record to the top of the list", func() {
					orderBy := []string{"id asc"}

					user := FabricateUser(faker.Email(), faker.Password())
					firstKeyword := FabricateKeyword(faker.Word(), "https://www.google.com/search?lr=lang_en", 0, user)
					secondKeyword := FabricateKeyword(faker.Word(), "https://www.google.com/search?lr=lang_en", 0, user)
					offset := 0
					limit := 2

					var queryList map[string]interface{}

					keywords, err := models.GetAllKeyword(queryList, orderBy, int64(offset), int64(limit))
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
					firstKeyword := FabricateKeyword(faker.Word(), "https://www.google.com/search?lr=lang_en", 0, user)
					secondKeyword := FabricateKeyword(faker.Word(), "https://www.google.com/search?lr=lang_en", 0, user)
					offset := 0
					limit := 2

					var queryList map[string]interface{}

					keywords, err := models.GetAllKeyword(queryList, orderBy, int64(offset), int64(limit))
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
					firstKeyword := FabricateKeyword(faker.Word(), "https://www.google.com/search?lr=lang_en", 0, user)
					secondKeyword := FabricateKeyword(faker.Word(), "https://www.google.com/search?lr=lang_en", 0, user)
					orderBy := []string{"created_at asc"}
					limit := 1

					var queryList map[string]interface{}

					keywords, err := models.GetAllKeyword(queryList, orderBy, int64(offset), int64(limit))
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
					firstKeyword := FabricateKeyword(faker.Word(), "https://www.google.com/search?lr=lang_en", 0, user)
					secondKeyword := FabricateKeyword(faker.Word(), "https://www.google.com/search?lr=lang_en", 0, user)
					orderBy := []string{"created_at asc"}
					limit := 1

					var queryList map[string]interface{}

					keywords, err := models.GetAllKeyword(queryList, orderBy, int64(offset), int64(limit))
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
					_ = FabricateKeyword(faker.Word(), "https://www.google.com/search?lr=lang_en", 0, user)
					_ = FabricateKeyword(faker.Word(), "https://www.google.com/search?lr=lang_en", 0, user)
					_ = FabricateKeyword(faker.Word(), "https://www.google.com/search?lr=lang_en", 0, user)
					orderBy := []string{"id asc"}
					offset := 0

					var queryList map[string]interface{}

					keywords, err := models.GetAllKeyword(queryList, orderBy, int64(offset), int64(limit))
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

				var queryList map[string]interface{}

				keywords, err := models.GetAllKeyword(queryList, orderBy, int64(offset), int64(limit))
				if err != nil {
					Fail(fmt.Sprintf("Get all keywords failed: %v", err.Error()))
				}

				Expect(len(keywords)).To(BeZero())
			})
		})
	})

	Describe("#CountAllKeyword", func() {
		Context("given a keyword record in the database", func() {
			It("returns a keyword record", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				_ = FabricateKeyword(faker.Word(), "https://www.google.com/search?lr=lang_en", 0, user)

				var queryList map[string]interface{}

				totalRows, err := models.CountAllKeyword(queryList)
				if err != nil {
					Fail(fmt.Sprintf("Count all keywords failed: %v", err.Error()))
				}

				Expect(totalRows).To(BeNumerically(">", 0))
			})
		})

		Context("given NO keyword exists", func() {
			It("returns an empty array", func() {
				var queryList map[string]interface{}

				totalRows, err := models.CountAllKeyword(queryList)
				if err != nil {
					Fail(fmt.Sprintf("Count all keyword failed: %v", err.Error()))
				}

				Expect(totalRows).To(BeZero())
			})
		})
	})

	Describe("#GetKeyword", func() {
		Context("given a keyword record in the database", func() {
			It("returns a keyword record", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				keyword := FabricateKeyword(faker.Word(), "https://www.google.com/search?lr=lang_en", 0, user)
				query := map[string]interface{}{
					"id": keyword.Id,
				}

				keywordResult, err := models.GetKeywordBy(query, []string{})
				if err != nil {
					Fail(fmt.Sprintf("Get keyword failed: %v", err.Error()))
				}

				Expect(keyword.Id).To(Equal(keywordResult.Id))
			})

			Context("given the keyword belongs to the user", func() {
				It("returns a keyword record", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					keyword := FabricateKeyword(faker.Word(), "https://www.google.com/search?lr=lang_en", 0, user)
					query := map[string]interface{}{
						"id":      keyword.Id,
						"user_id": user.Id,
					}

					keywordResult, err := models.GetKeywordBy(query, []string{})
					if err != nil {
						Fail(fmt.Sprintf("Get keyword failed: %v", err.Error()))
					}

					Expect(keyword.Id).To(Equal(keywordResult.Id))
				})
			})

			Context("given the keyword does NOT belong to the user", func() {
				It("returns `nil` with an error message", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					anotherUser := FabricateUser(faker.Email(), faker.Password())
					keyword := FabricateKeyword(faker.Word(), "https://www.google.com/search?lr=lang_en", 0, anotherUser)
					query := map[string]interface{}{
						"id":      keyword.Id,
						"user_id": user.Id,
					}

					keywordResult, err := models.GetKeywordBy(query, []string{})

					Expect(keywordResult).To(BeNil())
					Expect(err.Error()).To(ContainSubstring("no row found"))
				})
			})

			Context("given the keyword order by `id` in ascending order", func() {
				It("returns the first keyword", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					keyword := FabricateKeyword(faker.Word(), faker.URL(), 0, user)
					_ = FabricateKeyword(faker.Word(), faker.URL(), 0, user)
					query := map[string]interface{}{
						"user_id": user.Id,
					}
					order := []string{"id asc"}

					keywordResult, err := models.GetKeywordBy(query, order)
					if err != nil {
						Fail(fmt.Sprintf("Get keyword failed: %v", err.Error()))
					}

					Expect(keyword.Id).To(Equal(keywordResult.Id))
				})
			})

			Context("given the keyword order by `id` in descending order", func() {
				It("returns the first keyword", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					_ = FabricateKeyword(faker.Word(), faker.URL(), 0, user)
					keyword := FabricateKeyword(faker.Word(), faker.URL(), 0, user)
					query := map[string]interface{}{
						"user_id": user.Id,
					}
					order := []string{"id desc"}

					keywordResult, err := models.GetKeywordBy(query, order)
					if err != nil {
						Fail(fmt.Sprintf("Get keyword failed: %v", err.Error()))
					}

					Expect(keyword.Id).To(Equal(keywordResult.Id))
				})
			})
		})

		Context("given NO keyword exists", func() {
			It("returns `nil` with an error message", func() {
				query := map[string]interface{}{
					"id": 1,
				}

				keyword, err := models.GetKeywordBy(query, []string{})

				Expect(keyword).To(BeNil())
				Expect(err.Error()).To(ContainSubstring("no row found"))
			})
		})
	})

	Describe("#GetKeywordStatus", func() {
		Context("given `failed` status", func() {
			It("returns -1", func() {
				status := models.GetKeywordStatus("failed")

				Expect(status).To(Equal(-1))
			})
		})

		Context("given `pending` status", func() {
			It("returns 0", func() {
				status := models.GetKeywordStatus("pending")

				Expect(status).To(Equal(0))
			})
		})

		Context("given `completed` status", func() {
			It("returns 1", func() {
				status := models.GetKeywordStatus("completed")

				Expect(status).To(Equal(1))
			})
		})
	})

	Describe("#UpdateKeyword", func() {
		Context("given update keyword status from `pending` to `completed", func() {
			user := FabricateUser(faker.Email(), faker.Password())
			keyword := FabricateKeyword(faker.Word(), faker.URL(), 0, user)

			keyword.Status = models.GetKeywordStatus("completed")
			err := models.UpdateKeyword(keyword)
			if err != nil {
				Fail(fmt.Sprintf("Update keyword failed: %v", err))
			}

			refreshKeyword, err := models.GetKeywordBy(map[string]interface{}{"id": keyword.Id}, []string{})
			if err != nil {
				Fail(fmt.Sprintf("Get keyword failed: %v", err))
			}

			Expect(keyword.Id).To(Equal(refreshKeyword.Id))
			Expect(keyword.Status).To(Equal(refreshKeyword.Status))
		})
	})
})
