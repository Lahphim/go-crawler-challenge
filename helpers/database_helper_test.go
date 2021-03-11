package helpers_test

import (
	"go-crawler-challenge/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DatabaseHelper", func() {
	Describe("#FormatOrderByFor", func() {
		Context("given an item in the order param is ['email desc']", func() {
			It("returns ['-email']", func() {
				orderParam := []string{"email desc"}
				formattedOrderParam := helpers.FormatOrderByFor(orderParam)

				Expect(formattedOrderParam[0]).To(Equal("-email"))
			})
		})

		Context("given an item in the order param is ['email asc']", func() {
			It("returns ['email']", func() {
				orderParam := []string{"email asc"}
				formattedOrderParam := helpers.FormatOrderByFor(orderParam)

				Expect(formattedOrderParam[0]).To(Equal("email"))
			})
		})

		Context("given an item in the order param is ['user.email desc']", func() {
			It("returns ['-user__email']", func() {
				orderParam := []string{"user.email desc"}
				formattedOrderParam := helpers.FormatOrderByFor(orderParam)

				Expect(formattedOrderParam[0]).To(Equal("-user__email"))
			})
		})

		Context("given an item in the order param is ['user.email asc']", func() {
			It("returns ['user__email']", func() {
				orderParam := []string{"user.email asc"}
				formattedOrderParam := helpers.FormatOrderByFor(orderParam)

				Expect(formattedOrderParam[0]).To(Equal("user__email"))
			})
		})
	})
})
