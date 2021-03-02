package forms_test

import (
	"fmt"
	"strings"

	form "go-crawler-challenge/forms/scraper"
	. "go-crawler-challenge/models"
	. "go-crawler-challenge/tests"
	. "go-crawler-challenge/tests/fixtures"

	"github.com/beego/beego/v2/core/validation"
	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Scraper/KeywordResultForm", func() {
	AfterEach(func() {
		TruncateTable("link")
		TruncateTable("page")
		TruncateTable("keyword")
		TruncateTable("user")
	})

	Describe("#Valid", func() {
		Context("given valid params", func() {
			It("does NOT produce any errors", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				keyword := faker.Word()
				visitUrl := fmt.Sprintf("https://www.google.com/search?q=%s&lr=lang_en", keyword)
				rawHtml := faker.Paragraph()
				form := form.KeywordResultForm{
					Keyword:  keyword,
					VisitUrl: visitUrl,
					RawHtml:  rawHtml,
					User:     user,
				}

				formValidation := validation.Validation{}
				form.Valid(&formValidation)

				Expect(len(formValidation.Errors)).To(BeZero())
			})
		})

		Context("given INVALID params", func() {
			Context("given the user does NOT exist", func() {
				It("produces an error", func() {
					notExistingUser := &User{Base: Base{Id: 1}}

					keyword := faker.Word()
					visitUrl := fmt.Sprintf("https://www.google.com/search?q=%s&lr=lang_en", keyword)
					rawHtml := faker.Paragraph()
					form := form.KeywordResultForm{
						Keyword:  keyword,
						VisitUrl: visitUrl,
						RawHtml:  rawHtml,
						User:     notExistingUser,
					}

					formValidation := validation.Validation{}
					form.Valid(&formValidation)

					Expect(len(formValidation.Errors)).To(Equal(1))
					Expect(formValidation.Errors[0].Key).To(Equal("User"))
					Expect(formValidation.Errors[0].Message).To(Equal("User does not exist"))
				})
			})

			Context("given an INVALID visit URL", func() {
				It("produces an error", func() {
					invalidVisitUrl := "INVALID_VISIT_URL"

					user := FabricateUser(faker.Email(), faker.Password())
					keyword := faker.Word()
					rawHtml := faker.Paragraph()
					form := form.KeywordResultForm{
						Keyword:  keyword,
						VisitUrl: invalidVisitUrl,
						RawHtml:  rawHtml,
						User:     user,
					}

					formValidation := validation.Validation{}
					form.Valid(&formValidation)

					Expect(len(formValidation.Errors)).To(Equal(1))
					Expect(formValidation.Errors[0].Key).To(Equal("VisitUrl"))
					Expect(formValidation.Errors[0].Message).To(Equal("Visit URL must be valid URL"))
				})
			})

			Context("given an INVALID URL in the list of link", func() {
				It("produces an error", func() {
					invalidLinkList := []Link{
						{Url: "INVALID_URL"},
					}

					user := FabricateUser(faker.Email(), faker.Password())
					keyword := faker.Word()
					visitUrl := fmt.Sprintf("https://www.google.com/search?q=%s&lr=lang_en", keyword)
					rawHtml := faker.Paragraph()
					form := form.KeywordResultForm{
						Keyword:  keyword,
						VisitUrl: visitUrl,
						RawHtml:  rawHtml,
						LinkList: invalidLinkList,
						User:     user,
					}

					formValidation := validation.Validation{}
					form.Valid(&formValidation)

					Expect(len(formValidation.Errors)).To(Equal(1))
					Expect(formValidation.Errors[0].Key).To(Equal("LinkList"))
					Expect(formValidation.Errors[0].Message).To(Equal("All Link list must be valid URL"))
				})
			})
		})
	})

	Describe("#Create", func() {
		Context("given valid params", func() {
			It("returns a keyword record", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				keyword := faker.Word()
				visitUrl := fmt.Sprintf("https://www.google.com/search?q=%s&lr=lang_en", keyword)
				rawHtml := faker.Paragraph()
				form := form.KeywordResultForm{
					Keyword:  keyword,
					VisitUrl: visitUrl,
					RawHtml:  rawHtml,
					User:     user,
				}

				keywordRecord, _ := form.Save()

				Expect(keywordRecord).NotTo(BeNil())
				Expect(keywordRecord.Id).To(BeNumerically(">", 0))
			})

			It("does NOT produce any errors", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				keyword := faker.Word()
				visitUrl := fmt.Sprintf("https://www.google.com/search?q=%s&lr=lang_en", keyword)
				rawHtml := faker.Paragraph()
				form := form.KeywordResultForm{
					Keyword:  keyword,
					VisitUrl: visitUrl,
					RawHtml:  rawHtml,
					User:     user,
				}

				_, errors := form.Save()

				Expect(len(errors)).To(BeZero())
			})
		})

		Context("given INVALID params", func() {
			Context("given NO keyword", func() {
				It("does NOT return a keyword record", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					visitUrl := "https://www.google.com/search?q=&lr=lang_en"
					rawHtml := faker.Paragraph()
					form := form.KeywordResultForm{
						VisitUrl: visitUrl,
						RawHtml:  rawHtml,
						User:     user,
					}

					keywordRecord, errors := form.Save()

					Expect(keywordRecord).To(BeNil())
					Expect(errors[0].Error()).To(Equal("Keyword can not be empty"))
				})
			})

			Context("given keyword length is over than 128", func() {
				It("does NOT return a keyword record", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					keyword := faker.Word()
					visitUrl := fmt.Sprintf("https://www.google.com/search?q=%s&lr=lang_en", keyword)
					rawHtml := faker.Paragraph()
					form := form.KeywordResultForm{
						Keyword:  keyword + strings.Repeat("*", 128),
						VisitUrl: visitUrl,
						RawHtml:  rawHtml,
						User:     user,
					}

					keywordRecord, errors := form.Save()

					Expect(keywordRecord).To(BeNil())
					Expect(errors[0].Error()).To(Equal("Keyword maximum size is 128"))
				})
			})

			Context("given NO visit URL", func() {
				It("does NOT return a keyword record", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					keyword := faker.Word()
					rawHtml := faker.Paragraph()
					form := form.KeywordResultForm{
						Keyword: keyword,
						RawHtml: rawHtml,
						User:    user,
					}

					keywordRecord, errors := form.Save()

					Expect(keywordRecord).To(BeNil())
					Expect(errors[0].Error()).To(Equal("VisitUrl can not be empty"))
				})
			})

			Context("given keyword length is over than 128", func() {
				It("does NOT return a keyword record", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					keyword := faker.Word()
					visitUrl := fmt.Sprintf("https://www.google.com/search?q=%s%s&lr=lang_en", keyword, strings.Repeat("*", 128))
					rawHtml := faker.Paragraph()
					form := form.KeywordResultForm{
						Keyword:  keyword,
						VisitUrl: visitUrl,
						RawHtml:  rawHtml,
						User:     user,
					}

					keywordRecord, errors := form.Save()

					Expect(keywordRecord).To(BeNil())
					Expect(errors[0].Error()).To(Equal("VisitUrl maximum size is 128"))
				})
			})

			Context("given an INVALID URL in the list of link", func() {
				It("does NOT return a keyword record", func() {
					invalidLinkList := []Link{
						{Url: "INVALID_URL"},
					}

					user := FabricateUser(faker.Email(), faker.Password())
					keyword := faker.Word()
					visitUrl := fmt.Sprintf("https://www.google.com/search?q=%s&lr=lang_en", keyword)
					rawHtml := faker.Paragraph()
					form := form.KeywordResultForm{
						Keyword:  keyword,
						VisitUrl: visitUrl,
						RawHtml:  rawHtml,
						LinkList: invalidLinkList,
						User:     user,
					}

					keywordRecord, errors := form.Save()

					Expect(keywordRecord).To(BeNil())
					Expect(errors[0].Error()).To(Equal("All Link list must be valid URL"))
				})
			})

			Context("given NO raw HTML", func() {
				It("does NOT return a keyword record", func() {
					user := FabricateUser(faker.Email(), faker.Password())
					keyword := faker.Word()
					visitUrl := fmt.Sprintf("https://www.google.com/search?q=%s&lr=lang_en", keyword)
					form := form.KeywordResultForm{
						Keyword:  keyword,
						VisitUrl: visitUrl,
						User:     user,
					}

					keywordRecord, errors := form.Save()

					Expect(keywordRecord).To(BeNil())
					Expect(errors[0].Error()).To(Equal("RawHtml can not be empty"))
				})
			})

			Context("given NOT existing user", func() {
				It("does NOT return a keyword record", func() {
					notExistingUser := &User{Base: Base{Id: 1}}

					keyword := faker.Word()
					visitUrl := fmt.Sprintf("https://www.google.com/search?q=%s&lr=lang_en", keyword)
					rawHtml := faker.Paragraph()
					form := form.KeywordResultForm{
						Keyword:  keyword,
						VisitUrl: visitUrl,
						RawHtml:  rawHtml,
						User:     notExistingUser,
					}

					keywordRecord, errors := form.Save()

					Expect(keywordRecord).To(BeNil())
					Expect(errors[0].Error()).To(Equal("User does not exist"))
				})
			})

			Context("given NO user", func() {
				It("does NOT return a keyword record", func() {
					keyword := faker.Word()
					visitUrl := fmt.Sprintf("https://www.google.com/search?q=%s&lr=lang_en", keyword)
					rawHtml := faker.Paragraph()
					form := form.KeywordResultForm{
						Keyword:  keyword,
						VisitUrl: visitUrl,
						RawHtml:  rawHtml,
					}

					keywordRecord, errors := form.Save()

					Expect(keywordRecord).To(BeNil())
					Expect(errors[0].Error()).To(Equal("User can not be empty"))
				})
			})
		})
	})
})
