package v1serializers_test

import (
	"fmt"

	"go-crawler-challenge/models"
	v1serializers "go-crawler-challenge/serializers/v1"
	. "go-crawler-challenge/tests"
	. "go-crawler-challenge/tests/fixtures"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("V1/KeywordList", func() {
	AfterEach(func() {
		TruncateTable("keyword")
		TruncateTable("user")
	})

	Describe("#Data", func() {
		Context("given a valid data", func() {
			It("returns serialize data", func() {
				user := FabricateUser(faker.Email(), faker.Password())
				firstKeyword := FabricateKeyword(faker.Word(), faker.URL(), 0, user)
				secondKeyword := FabricateKeyword(faker.Word(), faker.URL(), 0, user)

				anotherUser := FabricateUser(faker.Email(), faker.Password())
				_ = FabricateKeyword(faker.Word(), faker.URL(), 0, anotherUser)

				keywords, err := models.GetAllKeyword(map[string]interface{}{"user_id": user.Id}, []string{"created_at desc"}, 0, 10)
				if err != nil {
					Fail(fmt.Sprintf("Get all keyword failed: %v", err.Error()))
				}

				serializer := v1serializers.KeywordList{
					KeywordList: keywords,
				}

				data := serializer.Data()

				Expect(len(data)).To(Equal(2))
				Expect(data[0].Id).To(Equal(fmt.Sprint(secondKeyword.Id)))
				Expect(data[1].Id).To(Equal(fmt.Sprint(firstKeyword.Id)))
			})
		})
	})
})
