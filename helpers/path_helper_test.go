package helpers_test

import (
	"go-crawler-challenge/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PathHelper", func() {
	Describe("#RootDir", func() {
		Context("given the current file", func() {
			It("returns root directory of this project", func() {
				Expect(helpers.RootDir()).To(ContainSubstring("go-crawler-challenge"))
			})
		})
	})
})
