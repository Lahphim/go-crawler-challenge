package keyword_test

import (
	"fmt"
	"strings"

	. "go-crawler-challenge/models"
	service "go-crawler-challenge/services/keyword"
	. "go-crawler-challenge/tests"
	. "go-crawler-challenge/tests/fixtures"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/validation"
	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Keyword/CreateKeywordResult", func() {
	AfterEach(func() {
		TruncateTable("link")
		TruncateTable("page")
		TruncateTable("position")
		TruncateTable("keyword")
		TruncateTable("user")
	})

	Describe("#Run", func() {
		Context("given valid params", func() {
			It("returns a keyword record", func() {
				position := FabricatePosition(faker.Word(), faker.Word(), faker.Word())
				user := FabricateUser(faker.Email(), faker.Password())
				keyword := FabricateKeyword(faker.Word(), "", 0, user)
				url := fmt.Sprintf("https://www.google.com/search?q=%s&lr=lang_en", keyword.Keyword)
				rawHtml := faker.Paragraph()
				linkList := []Link{
					{
						Url:      fmt.Sprintf("https://www.google.com/search?q=%s", faker.Word()),
						Position: position,
					},
				}

				createKeywordResultService := service.CreateKeywordResult{
					Keyword:  keyword,
					Url:      url,
					LinkList: linkList,
					RawHtml:  rawHtml,
				}

				keywordRecord, _ := createKeywordResultService.Run()

				Expect(keywordRecord).NotTo(BeNil())
				Expect(keywordRecord.Id).To(BeNumerically(">", 0))
			})

			It("does NOT produce any errors", func() {
				position := FabricatePosition(faker.Word(), faker.Word(), faker.Word())
				user := FabricateUser(faker.Email(), faker.Password())
				keyword := FabricateKeyword(faker.Word(), "", 0, user)
				url := fmt.Sprintf("https://www.google.com/search?q=%s&lr=lang_en", keyword.Keyword)
				rawHtml := faker.Paragraph()
				linkList := []Link{
					{
						Url:      fmt.Sprintf("https://www.google.com/search?q=%s", faker.Word()),
						Position: position,
					},
				}

				createKeywordResultService := service.CreateKeywordResult{
					Keyword:  keyword,
					Url:      url,
					LinkList: linkList,
					RawHtml:  rawHtml,
				}

				_, err := createKeywordResultService.Run()

				Expect(err).To(BeNil())
			})

			It("saves keyword to keyword table", func() {
				position := FabricatePosition(faker.Word(), faker.Word(), faker.Word())
				user := FabricateUser(faker.Email(), faker.Password())
				keyword := FabricateKeyword(faker.Word(), "", 0, user)
				url := fmt.Sprintf("https://www.google.com/search?q=%s&lr=lang_en", keyword.Keyword)
				rawHtml := faker.Paragraph()
				linkList := []Link{
					{
						Url:      fmt.Sprintf("https://www.google.com/search?q=%s", faker.Word()),
						Position: position,
					},
				}

				createKeywordResultService := service.CreateKeywordResult{
					Keyword:  keyword,
					Url:      url,
					LinkList: linkList,
					RawHtml:  rawHtml,
				}

				_, _ = createKeywordResultService.Run()

				ormer := orm.NewOrm()
				keywordRecord := &Keyword{}
				err := ormer.QueryTable(Keyword{}).RelatedSel().One(keywordRecord)
				if err != nil {
					Fail(fmt.Sprintf("Get first record failed: %v", err.Error()))
				}

				Expect(keywordRecord).NotTo(BeNil())
				Expect(keywordRecord.Id).To(BeNumerically(">", 0))
				Expect(keywordRecord.Keyword).To(Equal(keyword.Keyword))

				Expect(keywordRecord.User.Id).To(Equal(user.Id))
			})

			It("saves raw HTML to page table", func() {
				position := FabricatePosition(faker.Word(), faker.Word(), faker.Word())
				user := FabricateUser(faker.Email(), faker.Password())
				keyword := FabricateKeyword(faker.Word(), "", 0, user)
				url := fmt.Sprintf("https://www.google.com/search?q=%s&lr=lang_en", keyword.Keyword)
				rawHtml := faker.Paragraph()
				linkList := []Link{
					{
						Url:      fmt.Sprintf("https://www.google.com/search?q=%s", faker.Word()),
						Position: position,
					},
				}

				createKeywordResultService := service.CreateKeywordResult{
					Keyword:  keyword,
					Url:      url,
					LinkList: linkList,
					RawHtml:  rawHtml,
				}

				keywordRecord, _ := createKeywordResultService.Run()

				ormer := orm.NewOrm()
				pageRecord := &Page{}
				err := ormer.QueryTable(Page{}).RelatedSel().One(pageRecord)
				if err != nil {
					Fail(fmt.Sprintf("Get first record failed: %v", err.Error()))
				}

				Expect(pageRecord).NotTo(BeNil())
				Expect(pageRecord.Id).To(BeNumerically(">", 0))
				Expect(pageRecord.RawHtml).To(Equal(rawHtml))

				Expect(pageRecord.Keyword.Id).To(Equal(keywordRecord.Id))
			})

			It("saves all links to link table", func() {
				position := FabricatePosition(faker.Word(), faker.Word(), faker.Word())
				user := FabricateUser(faker.Email(), faker.Password())
				keyword := FabricateKeyword(faker.Word(), "", 0, user)
				url := fmt.Sprintf("https://www.google.com/search?q=%s&lr=lang_en", keyword.Keyword)
				rawHtml := faker.Paragraph()
				linkList := []Link{
					{
						Url:      fmt.Sprintf("https://www.google.com/search?q=%s", faker.Word()),
						Position: position,
					},
				}

				createKeywordResultService := service.CreateKeywordResult{
					Keyword:  keyword,
					Url:      url,
					LinkList: linkList,
					RawHtml:  rawHtml,
				}

				keywordRecord, _ := createKeywordResultService.Run()

				ormer := orm.NewOrm()
				linkRecord := &Link{}
				err := ormer.QueryTable(Link{}).RelatedSel().One(linkRecord)
				if err != nil {
					Fail(fmt.Sprintf("Get first record failed: %v", err.Error()))
				}

				Expect(linkRecord).NotTo(BeNil())
				Expect(linkRecord.Id).To(BeNumerically(">", 0))
				Expect(linkRecord.Url).To(Equal(linkList[0].Url))

				Expect(linkRecord.Keyword.Id).To(Equal(keywordRecord.Id))
				Expect(linkRecord.Position.Id).To(Equal(position.Id))
			})
		})

		Context("given INVALID url", func() {
			It("does NOT return a keyword record", func() {
				position := FabricatePosition(faker.Word(), faker.Word(), faker.Word())
				user := FabricateUser(faker.Email(), faker.Password())
				keyword := FabricateKeyword("keyword"+strings.Repeat("*", 120), "", 0, user)
				overLengthUrl := fmt.Sprintf("https://www.google.com/search?q=%s&lr=lang_en", keyword.Keyword)
				rawHtml := faker.Paragraph()
				linkList := []Link{
					{
						Url:      fmt.Sprintf("https://www.google.com/search?q=%s", faker.Word()),
						Position: position,
					},
				}

				createKeywordResultService := service.CreateKeywordResult{
					Keyword:  keyword,
					Url:      overLengthUrl,
					LinkList: linkList,
					RawHtml:  rawHtml,
				}

				keywordRecord, _ := createKeywordResultService.Run()

				Expect(keywordRecord).To(BeNil())
			})

			It("returns an error", func() {
				position := FabricatePosition(faker.Word(), faker.Word(), faker.Word())
				user := FabricateUser(faker.Email(), faker.Password())
				keyword := FabricateKeyword("keyword"+strings.Repeat("*", 120), "", 0, user)
				overLengthUrl := fmt.Sprintf("https://www.google.com/search?q=%s&lr=lang_en", keyword.Keyword)
				rawHtml := faker.Paragraph()
				linkList := []Link{
					{
						Url:      fmt.Sprintf("https://www.google.com/search?q=%s", faker.Word()),
						Position: position,
					},
				}

				createKeywordResultService := service.CreateKeywordResult{
					Keyword:  keyword,
					Url:      overLengthUrl,
					LinkList: linkList,
					RawHtml:  rawHtml,
				}

				_, err := createKeywordResultService.Run()

				Expect(err).NotTo(BeNil())
			})

			It("does NOT save to any tables", func() {
				position := FabricatePosition(faker.Word(), faker.Word(), faker.Word())
				user := FabricateUser(faker.Email(), faker.Password())
				keyword := FabricateKeyword("keyword"+strings.Repeat("*", 120), "", 0, user)
				overLengthUrl := fmt.Sprintf("https://www.google.com/search?q=%s&lr=lang_en", keyword.Keyword)
				rawHtml := faker.Paragraph()
				linkList := []Link{
					{
						Url:      fmt.Sprintf("https://www.google.com/search?q=%s", faker.Word()),
						Position: position,
					},
				}

				createKeywordResultService := service.CreateKeywordResult{
					Keyword:  keyword,
					Url:      overLengthUrl,
					LinkList: linkList,
					RawHtml:  rawHtml,
				}

				_, _ = createKeywordResultService.Run()

				ormer := orm.NewOrm()

				pageRecord := &Page{}
				pageErr := ormer.QueryTable(Page{}).RelatedSel().One(pageRecord)

				linkRecord := &Link{}
				linkErr := ormer.QueryTable(Link{}).RelatedSel().One(linkRecord)

				Expect(pageErr).NotTo(BeNil())
				Expect(pageErr.Error()).To(ContainSubstring("no row found"))
				Expect(linkErr).NotTo(BeNil())
				Expect(linkErr.Error()).To(ContainSubstring("no row found"))
			})
		})
	})

	Describe("#Valid", func() {
		Context("given valid params", func() {
			It("does NOT produce any errors", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				keyword := FabricateKeyword(faker.Word(), "", 0, user)
				url := fmt.Sprintf("https://www.google.com/search?q=%s&lr=lang_en", keyword.Keyword)
				rawHtml := faker.Paragraph()
				createKeywordResultService := service.CreateKeywordResult{
					Keyword: keyword,
					Url:     url,
					RawHtml: rawHtml,
				}

				formValidation := validation.Validation{}
				createKeywordResultService.Valid(&formValidation)

				Expect(len(formValidation.Errors)).To(BeZero())
			})
		})

		Context("given INVALID params", func() {
			Context("given an INVALID URL", func() {
				It("produces an error", func() {
					invalidUrl := "INVALID_URL"

					user := FabricateUser(faker.Email(), faker.Password())
					keyword := FabricateKeyword(faker.Word(), "", 0, user)
					rawHtml := faker.Paragraph()
					createKeywordResultService := service.CreateKeywordResult{
						Keyword: keyword,
						Url:     invalidUrl,
						RawHtml: rawHtml,
					}

					formValidation := validation.Validation{}
					createKeywordResultService.Valid(&formValidation)

					Expect(len(formValidation.Errors)).To(Equal(1))
					Expect(formValidation.Errors[0].Key).To(Equal("Url"))
					Expect(formValidation.Errors[0].Message).To(Equal("URL must be valid"))
				})
			})

			Context("given an INVALID URL in the list of link", func() {
				It("produces an error", func() {
					invalidLinkList := []Link{
						{Url: "INVALID_URL"},
					}

					user := FabricateUser(faker.Email(), faker.Password())
					keyword := FabricateKeyword(faker.Word(), "", 0, user)
					url := fmt.Sprintf("https://www.google.com/search?q=%s&lr=lang_en", keyword.Keyword)
					rawHtml := faker.Paragraph()
					createKeywordResultService := service.CreateKeywordResult{
						Keyword:  keyword,
						Url:      url,
						RawHtml:  rawHtml,
						LinkList: invalidLinkList,
					}

					formValidation := validation.Validation{}
					createKeywordResultService.Valid(&formValidation)

					Expect(len(formValidation.Errors)).To(Equal(1))
					Expect(formValidation.Errors[0].Key).To(Equal("LinkList"))
					Expect(formValidation.Errors[0].Message).To(Equal("All Link list must be valid URL"))
				})
			})
		})
	})
})
