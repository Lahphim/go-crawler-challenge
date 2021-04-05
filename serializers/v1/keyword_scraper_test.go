package v1serializers_test

import (
	v1serializers "go-crawler-challenge/serializers/v1"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("V1/KeywordScraper", func() {
	Describe("#Data", func() {
		Context("given a valid data", func() {
			It("returns serialize data", func() {
				message := faker.Sentence()

				serializer := v1serializers.KeywordScraper{
					Message: message,
				}

				data := serializer.Data()

				Expect(data.Id).To(BeZero())
				Expect(data.Message).To(Equal(message))
			})
		})
	})
})
