package v1serializers_test

import (
	"time"

	v1serializers "go-crawler-challenge/serializers/v1"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("V1/TokenInformation", func() {
	Describe("#Data", func() {
		Context("given a valid data", func() {
			It("returns serialize data", func() {
				accessToken := faker.Password()
				refreshToken := faker.Password()
				expiry := time.Hour * 2

				serializer := v1serializers.TokenInformation{
					AccessToken:  accessToken,
					RefreshToken: refreshToken,
					Expiry:       expiry,
				}

				data := serializer.Data()

				Expect(data.Id).To(BeZero())
				Expect(data.AccessToken).To(Equal(accessToken))
				Expect(data.TokenType).To(Equal(v1serializers.TokenType))
				Expect(data.RefreshToken).To(Equal(refreshToken))
				Expect(data.Expiry).To(Equal(expiry))
			})
		})
	})
})
