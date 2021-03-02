package forms_test

import (
	"strings"

	form "go-crawler-challenge/forms/session"
	. "go-crawler-challenge/tests"
	. "go-crawler-challenge/tests/fixtures"

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
			It("does NOT produce any errors", func() {
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

			Context("given a mismatch password", func() {
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

			Context("given a mismatch password", func() {
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

			Context("given NO email", func() {
				It("does NOT return a user object", func() {
					form := form.AuthenticationForm{
						Email:    "",
						Password: "password",
					}

					currentUser, _ := form.Authenticate()

					Expect(currentUser).To(BeNil())
				})

				It("produces an error", func() {
					form := form.AuthenticationForm{
						Email:    "",
						Password: "password",
					}

					_, errors := form.Authenticate()

					Expect(errors[0].Error()).To(Equal("Email can not be empty"))
				})
			})

			Context("given INVALID email", func() {
				It("does NOT return a user object", func() {
					form := form.AuthenticationForm{
						Email:    "INVALID_EMAIL",
						Password: "password",
					}

					currentUser, _ := form.Authenticate()

					Expect(currentUser).To(BeNil())
				})

				It("produces an error", func() {
					form := form.AuthenticationForm{
						Email:    "INVALID_EMAIL",
						Password: "password",
					}

					_, errors := form.Authenticate()

					Expect(errors[0].Error()).To(Equal("Email must be a valid email address"))
				})
			})

			Context("given email length is over than 100", func() {
				It("does NOT return a user object", func() {
					form := form.AuthenticationForm{
						Email:    "dev" + strings.Repeat("*", 100) + "@nimblehq.co",
						Password: "password",
					}

					currentUser, _ := form.Authenticate()

					Expect(currentUser).To(BeNil())
				})

				It("produces an error", func() {
					form := form.AuthenticationForm{
						Email:    "dev" + strings.Repeat("*", 100) + "@nimblehq.co",
						Password: "password",
					}

					_, errors := form.Authenticate()

					Expect(errors[0].Error()).To(Equal("Email maximum size is 100"))
				})
			})

			Context("given NO password", func() {
				It("does NOT return a user object", func() {
					form := form.AuthenticationForm{
						Email:    "dev@nimblehq.co",
						Password: "",
					}

					currentUser, _ := form.Authenticate()

					Expect(currentUser).To(BeNil())
				})

				It("produces an error", func() {
					form := form.AuthenticationForm{
						Email:    "dev@nimblehq.co",
						Password: "",
					}

					_, errors := form.Authenticate()

					Expect(errors[0].Error()).To(Equal("Password can not be empty"))
				})
			})
		})
	})
})
