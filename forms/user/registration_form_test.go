package forms_test

import (
	"strings"

	form "go-crawler-challenge/forms/user"
	. "go-crawler-challenge/tests"
	. "go-crawler-challenge/tests/fixtures"

	"github.com/beego/beego/v2/core/validation"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User/RegistrationForm", func() {
	AfterEach(func() {
		TruncateTable("user")
	})

	Describe("#Valid", func() {
		Context("given valid params", func() {
			It("does NOT produces any errors", func() {
				form := form.RegistrationForm{
					Email:           "dev@nimblehq.co",
					Password:        "password",
					ConfirmPassword: "password",
				}

				formValidation := validation.Validation{}
				form.Valid(&formValidation)

				Expect(len(formValidation.Errors)).To(BeZero())
			})
		})

		Context("given INVALID params", func() {
			Context("given an existing email", func() {
				It("produces an error", func() {
					email := "dev@nimblehq.co"
					_ = FabricateUser(email, "password")
					form := form.RegistrationForm{
						Email:           email,
						Password:        "password",
						ConfirmPassword: "password",
					}

					formValidation := validation.Validation{}
					form.Valid(&formValidation)

					Expect(len(formValidation.Errors)).To(Equal(1))
					Expect(formValidation.Errors[0].Key).To(Equal("Email"))
					Expect(formValidation.Errors[0].Message).To(Equal("Email is already in use"))
				})
			})

			Context("given a mismatch confirm password", func() {
				It("produces an error", func() {
					form := form.RegistrationForm{
						Email:           "dev@nimblehq.co",
						Password:        "password",
						ConfirmPassword: "pas____d",
					}

					formValidation := validation.Validation{}
					form.Valid(&formValidation)

					Expect(len(formValidation.Errors)).To(Equal(1))
					Expect(formValidation.Errors[0].Key).To(Equal("ConfirmPassword"))
					Expect(formValidation.Errors[0].Message).To(Equal("Confirm password confirmation must match"))
				})
			})
		})
	})

	Describe("#Create", func() {
		Context("given valid params", func() {
			It("returns user object", func() {
				form := form.RegistrationForm{
					Email:           "dev@nimblehq.co",
					Password:        "password",
					ConfirmPassword: "password",
				}

				user, errors := form.Create()
				if len(errors) > 0 {
					Fail("Save a new user with registration form failed")
				}

				Expect(user).NotTo(BeNil())
			})

			It("does NOT produce any errors", func() {
				form := form.RegistrationForm{
					Email:           "dev@nimblehq.co",
					Password:        "password",
					ConfirmPassword: "password",
				}

				_, errors := form.Create()

				Expect(len(errors)).To(BeZero())
			})
		})

		Context("given INVALID params", func() {
			Context("given an existing email", func() {
				It("does NOT return a user object", func() {
					email := "dev@nimblehq.co"
					_ = FabricateUser(email, "password")
					form := form.RegistrationForm{
						Email:           email,
						Password:        "password",
						ConfirmPassword: "password",
					}

					user, _ := form.Create()

					Expect(user).To(BeNil())
				})

				It("produces an error", func() {
					email := "dev@nimblehq.co"
					_ = FabricateUser(email, "password")
					form := form.RegistrationForm{
						Email:           "dev@nimblehq.co",
						Password:        "password",
						ConfirmPassword: "password",
					}

					_, errors := form.Create()

					Expect(errors[0].Error()).To(Equal("Email is already in use"))
				})
			})

			Context("given NO email", func() {
				It("does NOT return a user object", func() {
					form := form.RegistrationForm{
						Email:           "",
						Password:        "password",
						ConfirmPassword: "password",
					}

					user, _ := form.Create()

					Expect(user).To(BeNil())
				})

				It("produces an error", func() {
					form := form.RegistrationForm{
						Email:           "",
						Password:        "password",
						ConfirmPassword: "password",
					}

					_, errors := form.Create()

					Expect(errors[0].Error()).To(Equal("Email can not be empty"))
				})
			})

			Context("given an INVALID email", func() {
				It("does NOT return a user object", func() {
					form := form.RegistrationForm{
						Email:           "INVALID_EMAIL",
						Password:        "password",
						ConfirmPassword: "password",
					}

					user, _ := form.Create()

					Expect(user).To(BeNil())
				})

				It("produces an error", func() {
					form := form.RegistrationForm{
						Email:           "INVALID_EMAIL",
						Password:        "password",
						ConfirmPassword: "password",
					}

					_, errors := form.Create()

					Expect(errors[0].Error()).To(Equal("Email must be a valid email address"))
				})
			})

			Context("given email length is over than 100", func() {
				It("does NOT return a user object", func() {
					form := form.RegistrationForm{
						Email:           "dev" + strings.Repeat("*", 100) + "@nimblehq.co",
						Password:        "password",
						ConfirmPassword: "password",
					}

					user, _ := form.Create()

					Expect(user).To(BeNil())
				})

				It("produces an error", func() {
					form := form.RegistrationForm{
						Email:           "dev" + strings.Repeat("*", 100) + "@nimblehq.co",
						Password:        "password",
						ConfirmPassword: "password",
					}

					_, errors := form.Create()

					Expect(errors[0].Error()).To(Equal("Email maximum size is 100"))
				})
			})

			Context("given a mismatch confirm password", func() {
				It("does NOT return a user object", func() {
					form := form.RegistrationForm{
						Email:           "dev@nimblehq.co",
						Password:        "password",
						ConfirmPassword: "pas____d",
					}

					user, _ := form.Create()

					Expect(user).To(BeNil())
				})

				It("produces an error", func() {
					form := form.RegistrationForm{
						Email:           "dev@nimblehq.co",
						Password:        "password",
						ConfirmPassword: "pas____d",
					}

					_, errors := form.Create()

					Expect(errors[0].Error()).To(Equal("Confirm password confirmation must match"))
				})
			})

			Context("given NO password", func() {
				It("does NOT return a user object", func() {
					form := form.RegistrationForm{
						Email:           "dev@nimblehq.co",
						Password:        "",
						ConfirmPassword: "password",
					}

					user, _ := form.Create()

					Expect(user).To(BeNil())
				})

				It("produces an error", func() {
					form := form.RegistrationForm{
						Email:           "dev@nimblehq.co",
						Password:        "",
						ConfirmPassword: "password",
					}

					_, errors := form.Create()

					Expect(errors[0].Error()).To(Equal("Password can not be empty"))
				})
			})

			Context("given password length is less than 6", func() {
				It("does NOT return a user object", func() {
					form := form.RegistrationForm{
						Email:           "dev@nimblehq.co",
						Password:        "pass",
						ConfirmPassword: "password",
					}

					user, _ := form.Create()

					Expect(user).To(BeNil())
				})

				It("produces an error", func() {
					form := form.RegistrationForm{
						Email:           "dev@nimblehq.co",
						Password:        "pass",
						ConfirmPassword: "password",
					}

					_, errors := form.Create()

					Expect(errors[0].Error()).To(Equal("Password minimum size is 6"))
				})
			})

			Context("given password length is over than 12", func() {
				It("does NOT return a user object", func() {
					form := form.RegistrationForm{
						Email:           "dev@nimblehq.co",
						Password:        "password" + strings.Repeat("*", 12),
						ConfirmPassword: "password",
					}

					user, _ := form.Create()

					Expect(user).To(BeNil())
				})

				It("produces an error", func() {
					form := form.RegistrationForm{
						Email:           "dev@nimblehq.co",
						Password:        "password" + strings.Repeat("*", 12),
						ConfirmPassword: "password",
					}

					_, errors := form.Create()

					Expect(errors[0].Error()).To(Equal("Password maximum size is 12"))
				})
			})

			Context("given NO confirm password", func() {
				It("does NOT return a user object", func() {
					form := form.RegistrationForm{
						Email:           "dev@nimblehq.co",
						Password:        "password",
						ConfirmPassword: "",
					}

					user, _ := form.Create()

					Expect(user).To(BeNil())
				})

				It("produces an error", func() {
					form := form.RegistrationForm{
						Email:           "dev@nimblehq.co",
						Password:        "password",
						ConfirmPassword: "",
					}

					_, errors := form.Create()

					Expect(errors[0].Error()).To(Equal("ConfirmPassword can not be empty"))
				})
			})

			Context("given confirm password length is less than 6", func() {
				It("does NOT return a user object", func() {
					form := form.RegistrationForm{
						Email:           "dev@nimblehq.co",
						Password:        "password",
						ConfirmPassword: "pass",
					}

					user, _ := form.Create()

					Expect(user).To(BeNil())
				})

				It("produces an error", func() {
					form := form.RegistrationForm{
						Email:           "dev@nimblehq.co",
						Password:        "password",
						ConfirmPassword: "pass",
					}

					_, errors := form.Create()

					Expect(errors[0].Error()).To(Equal("ConfirmPassword minimum size is 6"))
				})
			})

			Context("given confirm password length is over than 12", func() {
				It("does NOT return a user object", func() {
					form := form.RegistrationForm{
						Email:           "dev@nimblehq.co",
						Password:        "password",
						ConfirmPassword: "password" + strings.Repeat("*", 12),
					}

					user, _ := form.Create()

					Expect(user).To(BeNil())
				})

				It("produces an error", func() {
					form := form.RegistrationForm{
						Email:           "dev@nimblehq.co",
						Password:        "password",
						ConfirmPassword: "password" + strings.Repeat("*", 12),
					}

					_, errors := form.Create()

					Expect(errors[0].Error()).To(Equal("ConfirmPassword maximum size is 12"))
				})
			})
		})
	})
})
