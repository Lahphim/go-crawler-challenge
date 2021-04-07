package v1serializers_test

import (
	"time"

	"go-crawler-challenge/models"
	v1serializers "go-crawler-challenge/serializers/v1"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("V1/ReportDetail", func() {
	Describe("#Data", func() {
		Context("given valid data", func() {
			It("returns serialize data", func() {
				adsTopLink := []string{faker.URL(), faker.URL(), faker.URL(), faker.URL()}
				adsOtherLink := []string{faker.URL()}
				normalLink := []string{faker.URL(), faker.URL(), faker.URL()}

				report := &models.Report{
					Id:      0,
					Keyword: faker.Word(),
					Url:     faker.URL(),
					RawHtml: faker.Sentence(),

					LinkList: map[string][]string{"normal": normalLink, "other": adsOtherLink, "top": adsTopLink},

					CreatedAt: time.Now().Local(),
					UpdatedAt: time.Now().Local(),

					TotalAdsTop: len(adsTopLink),
					TotalAds:    len(adsTopLink) + len(adsOtherLink),
					TotalNonAds: len(normalLink),
					TotalLink:   len(adsTopLink) + len(adsOtherLink) + len(normalLink),
				}

				serializer := v1serializers.ReportDetail{
					Report: report,
				}

				var data = serializer.Data()

				Expect(data).NotTo(BeNil())
				Expect(data.Id).To(Equal("0"))
				Expect(len(data.Keyword)).To(BeNumerically(">", 0))
				Expect(len(data.Url)).To(BeNumerically(">", 0))
				Expect(len(data.RawHtml)).To(BeNumerically(">", 0))
				Expect(len(data.LinkList["normal"])).To(BeNumerically(">=", 0))
				Expect(len(data.LinkList["other"])).To(BeNumerically(">=", 0))
				Expect(len(data.LinkList["top"])).To(BeNumerically(">=", 0))
				Expect(data.TotalAdsTop).To(Equal(len(adsTopLink)))
				Expect(data.TotalAds).To(Equal(len(adsTopLink) + len(adsOtherLink)))
				Expect(data.TotalNonAds).To(Equal(len(normalLink)))
				Expect(data.TotalLink).To(Equal(len(adsTopLink) + len(adsOtherLink) + len(normalLink)))
			})
		})
	})
})
