package models_test

import (
	"fmt"

	"go-crawler-challenge/models"
	. "go-crawler-challenge/tests"
	. "go-crawler-challenge/tests/fixtures"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Position", func() {
	AfterEach(func() {
		TruncateTable("position")
	})

	Describe("#GetAllPosition", func() {
		Context("given an existing position", func() {
			It("returns a position", func() {
				position := FabricatePosition("nonAds", "#search .g .yuRUbf > a", "normal")

				positions, err := models.GetAllPosition()
				if err != nil {
					Fail(fmt.Sprintf("Get all positions fails: %v", err))
				}

				Expect(len(positions)).To(BeNumerically(">", 0))
				Expect(positions[0].Id).To(Equal(position.Id))
			})

			It("does NOT return any errors", func() {
				_ = FabricatePosition("nonAds", "#search .g .yuRUbf > a", "normal")

				_, err := models.GetAllPosition()
				if err != nil {
					Fail(fmt.Sprintf("Get all positions fails: %v", err))
				}

				Expect(err).To(BeNil())
			})
		})

		Context("given NO position exist", func() {
			It("returns an empty array", func() {
				positions, err := models.GetAllPosition()
				if err != nil {
					Fail(fmt.Sprintf("Get all positions fails: %v", err))
				}

				Expect(len(positions)).To(BeZero())
			})

			It("does NOT return any errors", func() {
				_, err := models.GetAllPosition()
				if err != nil {
					Fail(fmt.Sprintf("Get all positions fails: %v", err))
				}

				Expect(err).To(BeNil())
			})
		})
	})
})
