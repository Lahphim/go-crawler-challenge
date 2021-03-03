package transactions_test

import (
	"fmt"
	"strings"

	. "go-crawler-challenge/models"
	. "go-crawler-challenge/tests"
	. "go-crawler-challenge/tests/fixtures"
	"go-crawler-challenge/transactions"

	"github.com/beego/beego/v2/client/orm"
	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("KeywordResult", func() {
	AfterEach(func() {
		TruncateTable("link")
		TruncateTable("page")
		TruncateTable("position")
		TruncateTable("keyword")
		TruncateTable("user")
	})

	Describe("#AddKeywordResult", func() {
		Context("given valid params", func() {
			It("returns keyword record", func() {
				position := FabricatePosition(faker.Word(), faker.Word(), faker.Word())
				user := FabricateUser(faker.Email(), faker.Password())
				keyword := faker.Word()
				url := fmt.Sprintf("https://www.google.com/search?q=%s&lr=lang_en", keyword)
				rawHtml := faker.Paragraph()
				linkList := []Link{
					{
						Url:      fmt.Sprintf("https://www.google.com/search?q=%s", faker.Word()),
						Position: position,
					},
				}

				keywordResult := transactions.KeywordResult{
					Keyword:  keyword,
					Url:      url,
					LinkList: linkList,
					RawHtml:  rawHtml,
					User:     user,
				}

				keywordRecord, _ := transactions.AddKeywordResult(&keywordResult)

				Expect(keywordRecord).NotTo(BeNil())
				Expect(keywordRecord.Id).To(BeNumerically(">", 0))
			})

			It("does NOT produces any errors", func() {
				position := FabricatePosition(faker.Word(), faker.Word(), faker.Word())
				user := FabricateUser(faker.Email(), faker.Password())
				keyword := faker.Word()
				url := fmt.Sprintf("https://www.google.com/search?q=%s&lr=lang_en", keyword)
				rawHtml := faker.Paragraph()
				linkList := []Link{
					{
						Url:      fmt.Sprintf("https://www.google.com/search?q=%s", faker.Word()),
						Position: position,
					},
				}

				keywordResult := transactions.KeywordResult{
					Keyword:  keyword,
					Url:      url,
					LinkList: linkList,
					RawHtml:  rawHtml,
					User:     user,
				}

				_, err := transactions.AddKeywordResult(&keywordResult)

				Expect(err).To(BeNil())
			})

			It("saves keyword to keyword table", func() {
				position := FabricatePosition(faker.Word(), faker.Word(), faker.Word())
				user := FabricateUser(faker.Email(), faker.Password())
				keyword := faker.Word()
				url := fmt.Sprintf("https://www.google.com/search?q=%s&lr=lang_en", keyword)
				rawHtml := faker.Paragraph()
				linkList := []Link{
					{
						Url:      fmt.Sprintf("https://www.google.com/search?q=%s", faker.Word()),
						Position: position,
					},
				}

				keywordResult := transactions.KeywordResult{
					Keyword:  keyword,
					Url:      url,
					LinkList: linkList,
					RawHtml:  rawHtml,
					User:     user,
				}

				_, _ = transactions.AddKeywordResult(&keywordResult)

				ormer := orm.NewOrm()
				keywordRecord := &Keyword{}
				err := ormer.QueryTable(Keyword{}).RelatedSel().One(keywordRecord)
				if err != nil {
					Fail(fmt.Sprintf("Get first record failed: %v", err.Error()))
				}

				Expect(keywordRecord).NotTo(BeNil())
				Expect(keywordRecord.Id).To(BeNumerically(">", 0))
				Expect(keywordRecord.Keyword).To(Equal(keyword))

				Expect(keywordRecord.User.Id).To(Equal(user.Id))
			})

			It("saves raw HTML to page table", func() {
				position := FabricatePosition(faker.Word(), faker.Word(), faker.Word())
				user := FabricateUser(faker.Email(), faker.Password())
				keyword := faker.Word()
				url := fmt.Sprintf("https://www.google.com/search?q=%s&lr=lang_en", keyword)
				rawHtml := faker.Paragraph()
				linkList := []Link{
					{
						Url:      fmt.Sprintf("https://www.google.com/search?q=%s", faker.Word()),
						Position: position,
					},
				}

				keywordResult := transactions.KeywordResult{
					Keyword:  keyword,
					Url:      url,
					LinkList: linkList,
					RawHtml:  rawHtml,
					User:     user,
				}

				keywordRecord, _ := transactions.AddKeywordResult(&keywordResult)

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
				keyword := faker.Word()
				url := fmt.Sprintf("https://www.google.com/search?q=%s&lr=lang_en", keyword)
				rawHtml := faker.Paragraph()
				linkList := []Link{
					{
						Url:      fmt.Sprintf("https://www.google.com/search?q=%s", faker.Word()),
						Position: position,
					},
				}

				keywordResult := transactions.KeywordResult{
					Keyword:  keyword,
					Url:      url,
					LinkList: linkList,
					RawHtml:  rawHtml,
					User:     user,
				}

				keywordRecord, _ := transactions.AddKeywordResult(&keywordResult)

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

		Context("given INVALID params", func() {
			It("does NOT return a keyword record", func() {
				overLengthKeyword := faker.Word() + strings.Repeat("*", 500)

				position := FabricatePosition(faker.Word(), faker.Word(), faker.Word())
				user := FabricateUser(faker.Email(), faker.Password())
				url := fmt.Sprintf("https://www.google.com/search?q=%s&lr=lang_en", overLengthKeyword)
				rawHtml := faker.Paragraph()
				linkList := []Link{
					{
						Url:      fmt.Sprintf("https://www.google.com/search?q=%s", faker.Word()),
						Position: position,
					},
				}

				keywordResult := transactions.KeywordResult{
					Keyword:  overLengthKeyword,
					Url:      url,
					LinkList: linkList,
					RawHtml:  rawHtml,
					User:     user,
				}

				keywordRecord, _ := transactions.AddKeywordResult(&keywordResult)

				Expect(keywordRecord).To(BeNil())
			})

			It("returns an error", func() {
				overLengthKeyword := faker.Word() + strings.Repeat("*", 500)

				position := FabricatePosition(faker.Word(), faker.Word(), faker.Word())
				user := FabricateUser(faker.Email(), faker.Password())
				url := fmt.Sprintf("https://www.google.com/search?q=%s&lr=lang_en", overLengthKeyword)
				rawHtml := faker.Paragraph()
				linkList := []Link{
					{
						Url:      fmt.Sprintf("https://www.google.com/search?q=%s", faker.Word()),
						Position: position,
					},
				}

				keywordResult := transactions.KeywordResult{
					Keyword:  overLengthKeyword,
					Url:      url,
					LinkList: linkList,
					RawHtml:  rawHtml,
					User:     user,
				}

				_, err := transactions.AddKeywordResult(&keywordResult)

				Expect(err).NotTo(BeNil())
			})

			It("does NOT save to any tables", func() {
				overLengthKeyword := faker.Word() + strings.Repeat("*", 500)

				position := FabricatePosition(faker.Word(), faker.Word(), faker.Word())
				user := FabricateUser(faker.Email(), faker.Password())
				url := fmt.Sprintf("https://www.google.com/search?q=%s&lr=lang_en", overLengthKeyword)
				rawHtml := faker.Paragraph()
				linkList := []Link{
					{
						Url:      fmt.Sprintf("https://www.google.com/search?q=%s", faker.Word()),
						Position: position,
					},
				}

				keywordResult := transactions.KeywordResult{
					Keyword:  overLengthKeyword,
					Url:      url,
					LinkList: linkList,
					RawHtml:  rawHtml,
					User:     user,
				}

				_, _ = transactions.AddKeywordResult(&keywordResult)

				ormer := orm.NewOrm()
				keywordRecord := &Keyword{}
				keywordErr := ormer.QueryTable(Keyword{}).RelatedSel().One(keywordRecord)

				pageRecord := &Page{}
				pageErr := ormer.QueryTable(Page{}).RelatedSel().One(pageRecord)

				linkRecord := &Link{}
				linkErr := ormer.QueryTable(Link{}).RelatedSel().One(linkRecord)

				Expect(keywordErr).NotTo(BeNil())
				Expect(keywordErr.Error()).To(ContainSubstring("no row found"))
				Expect(pageErr).NotTo(BeNil())
				Expect(pageErr.Error()).To(ContainSubstring("no row found"))
				Expect(linkErr).NotTo(BeNil())
				Expect(linkErr.Error()).To(ContainSubstring("no row found"))
			})
		})
	})
})
