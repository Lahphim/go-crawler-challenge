package apiforms_test

import (
	"fmt"
	"go-crawler-challenge/forms"

	apiform "go-crawler-challenge/forms/api/token"
	. "go-crawler-challenge/tests/fixtures"

	"github.com/bxcodec/faker/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Token/GeneratorForm", func() {
	Describe("#Generate", func() {
		Context("given valid params", func() {
			It("returns a token information", func() {
				password := faker.Password()
				user := FabricateUser(faker.Email(), password)
				oauthClient := FabricateOauthClient(faker.UUIDHyphenated(), faker.Password(), "localhost:8080")
				generatorForm := apiform.GeneratorForm{
					ClientId:     oauthClient.ID,
					ClientSecret: oauthClient.Secret,
					GrantType:    "password",
					Email:        user.Email,
					Password:     password,
				}

				tokenInfo, err := generatorForm.Generate()
				if err != nil {
					Fail(fmt.Sprintf("Generate token failed: %v", err.Error()))
				}

				Expect(tokenInfo).NotTo(BeNil())
				Expect(len(tokenInfo.GetAccess())).To(BeNumerically(">", 0))
				Expect(len(tokenInfo.GetRefresh())).To(BeNumerically(">", 0))
				Expect(tokenInfo.GetAccessExpiresIn()).To(BeNumerically(">", 0))
			})

			It("does NOT produce any errors", func() {
				password := faker.Password()
				user := FabricateUser(faker.Email(), password)
				oauthClient := FabricateOauthClient(faker.UUIDHyphenated(), faker.Password(), "localhost:8080")
				generatorForm := apiform.GeneratorForm{
					ClientId:     oauthClient.ID,
					ClientSecret: oauthClient.Secret,
					GrantType:    "password",
					Email:        user.Email,
					Password:     password,
				}

				_, err := generatorForm.Generate()

				Expect(err).To(BeNil())
			})
		})

		Context("given INVALID params", func() {
			Context("given NO `ClientId` exists", func() {
				It("produces an error", func() {
					password := faker.Password()
					user := FabricateUser(faker.Email(), password)
					oauthClient := FabricateOauthClient(faker.UUIDHyphenated(), faker.Password(), "localhost:8080")
					generatorForm := apiform.GeneratorForm{
						ClientId:     "",
						ClientSecret: oauthClient.Secret,
						GrantType:    "password",
						Email:        user.Email,
						Password:     password,
					}

					_, err := generatorForm.Generate()

					Expect(err).NotTo(BeNil())
					Expect(err.Error()).To(Equal("ClientId cannot be empty"))
				})
			})

			Context("given NO `ClientSecret` exists", func() {
				It("produces an error", func() {
					password := faker.Password()
					user := FabricateUser(faker.Email(), password)
					oauthClient := FabricateOauthClient(faker.UUIDHyphenated(), faker.Password(), "localhost:8080")
					generatorForm := apiform.GeneratorForm{
						ClientId:     oauthClient.ID,
						ClientSecret: "",
						GrantType:    "password",
						Email:        user.Email,
						Password:     password,
					}

					_, err := generatorForm.Generate()

					Expect(err).NotTo(BeNil())
					Expect(err.Error()).To(Equal("ClientSecret cannot be empty"))
				})
			})

			Context("given INVALID client credentials", func() {
				It("produces an error", func() {
					password := faker.Password()
					user := FabricateUser(faker.Email(), password)
					_ = FabricateOauthClient(faker.UUIDHyphenated(), faker.Password(), "localhost:8080")
					generatorForm := apiform.GeneratorForm{
						ClientId:     "INVALID_CLIENT_ID",
						ClientSecret: "INVALID_CLIENT_SECRET",
						GrantType:    "password",
						Email:        user.Email,
						Password:     password,
					}

					_, err := generatorForm.Generate()

					Expect(err).NotTo(BeNil())
					Expect(err.Error()).To(Equal(forms.ValidationMessages["InvalidClient"]))
				})
			})

			Context("given NO `GrantType` exists", func() {
				It("produces an error", func() {
					password := faker.Password()
					user := FabricateUser(faker.Email(), password)
					oauthClient := FabricateOauthClient(faker.UUIDHyphenated(), faker.Password(), "localhost:8080")
					generatorForm := apiform.GeneratorForm{
						ClientId:     oauthClient.ID,
						ClientSecret: oauthClient.Secret,
						GrantType:    "",
						Email:        user.Email,
						Password:     password,
					}

					_, err := generatorForm.Generate()

					Expect(err).NotTo(BeNil())
					Expect(err.Error()).To(Equal("GrantType cannot be empty"))
				})
			})

			Context("given INVALID grant type", func() {
				It("produces an error", func() {
					password := faker.Password()
					user := FabricateUser(faker.Email(), password)
					oauthClient := FabricateOauthClient(faker.UUIDHyphenated(), faker.Password(), "localhost:8080")
					generatorForm := apiform.GeneratorForm{
						ClientId:     oauthClient.ID,
						ClientSecret: oauthClient.Secret,
						GrantType:    "INVALID_GRANT_TYPE",
						Email:        user.Email,
						Password:     password,
					}

					_, err := generatorForm.Generate()

					Expect(err).NotTo(BeNil())
					Expect(err.Error()).To(Equal(forms.ValidationMessages["InvalidGrantType"]))
				})
			})

			Context("given NO `Email` exists", func() {
				It("produces an error", func() {
					password := faker.Password()
					_ = FabricateUser(faker.Email(), password)
					oauthClient := FabricateOauthClient(faker.UUIDHyphenated(), faker.Password(), "localhost:8080")
					generatorForm := apiform.GeneratorForm{
						ClientId:     oauthClient.ID,
						ClientSecret: oauthClient.Secret,
						GrantType:    "password",
						Email:        "",
						Password:     password,
					}

					_, err := generatorForm.Generate()

					Expect(err).NotTo(BeNil())
					Expect(err.Error()).To(Equal("Email cannot be empty"))
				})
			})

			Context("given NO `Password` exists", func() {
				It("produces an error", func() {
					password := faker.Password()
					user := FabricateUser(faker.Email(), password)
					oauthClient := FabricateOauthClient(faker.UUIDHyphenated(), faker.Password(), "localhost:8080")
					generatorForm := apiform.GeneratorForm{
						ClientId:     oauthClient.ID,
						ClientSecret: oauthClient.Secret,
						GrantType:    "password",
						Email:        user.Email,
						Password:     "",
					}

					_, err := generatorForm.Generate()

					Expect(err).NotTo(BeNil())
					Expect(err.Error()).To(Equal("Password cannot be empty"))
				})
			})

			Context("given INVALID user credentials", func() {
				It("produces an error", func() {
					password := faker.Password()
					_ = FabricateUser(faker.Email(), password)
					oauthClient := FabricateOauthClient(faker.UUIDHyphenated(), faker.Password(), "localhost:8080")
					generatorForm := apiform.GeneratorForm{
						ClientId:     oauthClient.ID,
						ClientSecret: oauthClient.Secret,
						GrantType:    "password",
						Email:        "INVALID@MAIL.COM",
						Password:     "INVALID_PASSWORD",
					}

					_, err := generatorForm.Generate()

					Expect(err).NotTo(BeNil())
					Expect(err.Error()).To(Equal(forms.ValidationMessages["InvalidCredential"]))
				})
			})
		})
	})
})
