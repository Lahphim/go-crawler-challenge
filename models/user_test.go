package models_test

import (
	"fmt"

	"go-crawler-challenge/models"
	. "go-crawler-challenge/tests"
	. "go-crawler-challenge/tests/fixtures"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {
	AfterEach(func() {
		TruncateTable("user")
	})

	Describe("#AddUser", func() {
		Context("given a valid user", func() {
			It("returns the user ID", func() {
				user := models.User{
					Email:          "dev@nimblehq.co",
					HashedPassword: "password",
				}

				id, err := models.AddUser(&user)
				if err != nil {
					Fail(fmt.Sprintf("Add a new user failed: %v", err))
				}

				Expect(id).To(BeNumerically(">", 0))
			})

			It("does NOT return any errors", func() {
				user := models.User{
					Email:          "dev@nimblehq.co",
					HashedPassword: "password",
				}

				_, err := models.AddUser(&user)

				Expect(err).To(BeNil())
			})
		})

		Context("given an INVALID user", func() {
			Context("given an existing user", func() {
				It("returns 0", func() {
					_ = FabricateUser("dev@nimblehq.co", "password")
					user := models.User{
						Email:          "dev@nimblehq.co",
						HashedPassword: "password",
					}

					id, _ := models.AddUser(&user)

					Expect(id).To(BeZero())
				})

				It("returns an error", func() {
					_ = FabricateUser("dev@nimblehq.co", "password")
					user := models.User{
						Email:          "dev@nimblehq.co",
						HashedPassword: "password",
					}

					_, err := models.AddUser(&user)

					Expect(err.Error()).To(Equal(`pq: duplicate key value violates unique constraint "user_email_key"`))
				})
			})
		})
	})

	Describe("#GetUserById", func() {
		Context("given an existing user", func() {
			It("returns a user", func() {
				existUser := FabricateUser("dev@nimblehq.co", "password")

				user, err := models.GetUserById(existUser.Id)
				if err != nil {
					Fail(fmt.Sprintf("Get user with ID failed: %v", err))
				}

				Expect(user.Email).To(Equal(existUser.Email))
				Expect(user.HashedPassword).To(Equal(existUser.HashedPassword))
			})

			It("does NOT return any errors", func() {
				existUser := FabricateUser("dev@nimblehq.co", "password")

				_, err := models.GetUserById(existUser.Id)
				if err != nil {
					Fail(fmt.Sprintf("Get user with ID failed: %v", err))
				}

				Expect(err).To(BeNil())
			})
		})

		Context("given NO user exist", func() {
			It("does NOT return a user", func() {
				user, _ := models.GetUserById(42)

				Expect(user).To(BeNil())
			})

			It("returns an error", func() {
				_, err := models.GetUserById(42)

				Expect(err).NotTo(BeNil())
			})
		})
	})

	Describe("#GetUserByEmail", func() {
		Context("given an existing user", func() {
			It("returns a user", func() {
				existUser := FabricateUser("dev@nimblehq.co", "password")

				user, err := models.GetUserByEmail(existUser.Email)
				if err != nil {
					Fail(fmt.Sprintf("Get user with email failed: %v", err))
				}

				Expect(user.Email).To(Equal(existUser.Email))
				Expect(user.HashedPassword).To(Equal(existUser.HashedPassword))
			})

			It("does NOT return any errors", func() {
				existUser := FabricateUser("dev@nimblehq.co", "password")

				_, err := models.GetUserByEmail(existUser.Email)
				if err != nil {
					Fail(fmt.Sprintf("Get user with email failed: %v", err))
				}

				Expect(err).To(BeNil())
			})
		})

		Context("given NO user exist", func() {
			It("does NOT return a user", func() {
				user, _ := models.GetUserByEmail("dev@nimblehq.co")

				Expect(user).To(BeNil())
			})

			It("returns an error", func() {
				_, err := models.GetUserByEmail("dev@nimblehq.co")

				Expect(err).NotTo(BeNil())
			})
		})
	})
})
