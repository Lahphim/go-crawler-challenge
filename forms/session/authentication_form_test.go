package forms_test

import (
	form "go-crawler-challenge/forms/session"
	. "go-crawler-challenge/tests/test_helpers"
	. "go-crawler-challenge/tests/test_helpers/fabricators"

	"github.com/beego/beego/v2/core/validation"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Session/AuthenticationForm", func() {
	AfterEach(func() {
		TruncateTable("user")
	})

	Describe("#Valid", func() {
		Context("given valid params", func() {
			It("does NOT produces any errors", func() {
				_ = FabricateUser("dev@nimblehq.co", "password")
				form := form.AuthenticationForm{
					Email:    "dev@nimblehq.co",
					Password: "password",
				}

				formValidation := validation.Validation{}
				form.Valid(&formValidation)

				Expect(len(formValidation.Errors)).To(BeZero())
			})
		})

		Context("given INVALID params", func() {
			Context("given the email does NOT exist", func() {
				It("produces an error", func() {
					form := form.AuthenticationForm{
						Email:    "dev@nimblehq.co",
						Password: "password",
					}

					formValidation := validation.Validation{}
					form.Valid(&formValidation)

					Expect(len(formValidation.Errors)).To(Equal(1))
					Expect(formValidation.Errors[0].Key).To(Equal("Email"))
					Expect(formValidation.Errors[0].Message).To(Equal("Your email or password is incorrect"))
				})
			})

			Context("given an mismatch password", func() {
				It("produces an error", func() {
					_ = FabricateUser("dev@nimblehq.co", "password")
					form := form.AuthenticationForm{
						Email:    "dev@nimblehq.co",
						Password: "INVALID_PASSWORD",
					}

					formValidation := validation.Validation{}
					form.Valid(&formValidation)

					Expect(len(formValidation.Errors)).To(Equal(1))
					Expect(formValidation.Errors[0].Key).To(Equal("Email"))
					Expect(formValidation.Errors[0].Message).To(Equal("Your email or password is incorrect"))
				})
			})
		})
	})

	Describe("#Authenticate", func() {
		Context("given valid params", func() {
			It("returns user object", func() {
				user := FabricateUser("dev@nimblehq.co", "password")
				form := form.AuthenticationForm{
					Email:    "dev@nimblehq.co",
					Password: "password",
				}

				currentUser, _ := form.Authenticate()

				Expect(currentUser).NotTo(BeNil())
				Expect(currentUser.Id).To(Equal(user.Id))
			})

			It("does NOT produce any errors", func() {
				_ = FabricateUser("dev@nimblehq.co", "password")
				form := form.AuthenticationForm{
					Email:    "dev@nimblehq.co",
					Password: "password",
				}

				_, errors := form.Authenticate()

				Expect(len(errors)).To(BeZero())
			})
		})

		Context("given INVALID params", func() {
			Context("given the email does NOT exist", func() {
				It("does NOT return a user object", func() {
					form := form.AuthenticationForm{
						Email:    "dev@nimblehq.co",
						Password: "password",
					}

					currentUser, _ := form.Authenticate()

					Expect(currentUser).To(BeNil())
				})

				It("produces an error", func() {
					form := form.AuthenticationForm{
						Email:    "dev@nimblehq.co",
						Password: "password",
					}

					_, errors := form.Authenticate()

					Expect(errors[0].Error()).To(Equal("Your email or password is incorrect"))
				})
			})

			Context("given an mismatch password", func() {
				It("does NOT return a user object", func() {
					_ = FabricateUser("dev@nimblehq.co", "password")
					form := form.AuthenticationForm{
						Email:    "dev@nimblehq.co",
						Password: "INVALID_PASSWORD",
					}

					currentUser, _ := form.Authenticate()

					Expect(currentUser).To(BeNil())
				})

				It("produces an error", func() {
					_ = FabricateUser("dev@nimblehq.co", "password")
					form := form.AuthenticationForm{
						Email:    "dev@nimblehq.co",
						Password: "INVALID_PASSWORD",
					}

					_, errors := form.Authenticate()

					Expect(errors[0].Error()).To(Equal("Your email or password is incorrect"))
				})
			})
		})
	})
})
