package helpers_test

import (
	"go-crawler-challenge/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("QuerySetterHelper", func() {
	Describe("#BuildOrderByFor", func() {
		Context("given an item in the order list is ['email desc']", func() {
			It("returns ['-email']", func() {
				oldOrderList := []string{"email desc"}
				newOrderList := helpers.BuildOrderByFor(oldOrderList)

				Expect(newOrderList[0]).To(Equal("-email"))
			})
		})

		Context("given an item in the order list is ['email asc']", func() {
			It("returns ['email']", func() {
				oldOrderList := []string{"email asc"}
				newOrderList := helpers.BuildOrderByFor(oldOrderList)

				Expect(newOrderList[0]).To(Equal("email"))
			})
		})

		Context("given an item in the order list is ['user.email desc']", func() {
			It("returns ['-user__email']", func() {
				oldOrderList := []string{"user.email desc"}
				newOrderList := helpers.BuildOrderByFor(oldOrderList)

				Expect(newOrderList[0]).To(Equal("-user__email"))
			})
		})

		Context("given an item in the order list is ['user.email asc']", func() {
			It("returns ['user__email']", func() {
				oldOrderList := []string{"user.email asc"}
				newOrderList := helpers.BuildOrderByFor(oldOrderList)

				Expect(newOrderList[0]).To(Equal("user__email"))
			})
		})
	})
})
