package helpers_test

import (
	"go-crawler-challenge/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UtilityMapHelper", func() {
	Describe("#GetMapKeyByNumberValue", func() {
		Context("given a valid map list with integer", func() {
			It("returns the key of the first entry matched", func() {
				mapList := map[string]int{"one": 1, "two": 2}
				key := helpers.GetMapKeyByNumberValue(mapList, 1)

				Expect(key).To(Equal("one"))
				Expect(key).NotTo(Equal("two"))
			})
		})
	})
})
