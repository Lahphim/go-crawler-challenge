package helpers_test

import (
	"go-crawler-challenge/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UtilityMapHelper", func() {
	Describe("#GetMapKeyByNumber", func() {
		Context("given a valid map list with integer", func() {
			It("returns the first key", func() {
				mapList := map[string]int{"one": 1, "two": 2}
				key := helpers.GetMapKeyByNumber(mapList, 1)

				Expect(key).To(Equal("one"))
				Expect(key).NotTo(Equal("two"))
			})
		})
	})
})
