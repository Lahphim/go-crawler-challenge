package helpers_test

import (
	"fmt"

	"go-crawler-challenge/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/crypto/bcrypt"
)

var _ = Describe("PasswordHelper", func() {
	Describe("#HashPassword", func() {
		Context("given a plain password", func() {
			It("returns a hashed password", func() {
				plainPassword := "password"
				hashedPassword, err := helpers.HashPassword(plainPassword)
				if err != nil {
					Fail(fmt.Sprintf("Hash password failed: %v", err))
				}

				err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))

				Expect(err).To(BeNil())
			})
		})
	})

	Describe("#CheckMatchPassword", func() {
		Context("given valid hashed and plain password", func() {
			It("returns an error", func() {
				plainPassword := "password"
				hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
				if err != nil {
					Fail(fmt.Sprintf("Hash password failed: %v", err))
				}

				err = helpers.CheckMatchPassword(string(hashedPassword), plainPassword)

				Expect(err).To(BeNil())
			})
		})

		Context("given INVALID hashed and plain password", func() {
			It("returns false", func() {
				plainPassword := "password"
				mismatchPassword := "mismatch-password"
				hashedPassword, err := bcrypt.GenerateFromPassword([]byte(mismatchPassword), bcrypt.DefaultCost)
				if err != nil {
					Fail(fmt.Sprintf("Hash password failed: %v", err))
				}

				err = helpers.CheckMatchPassword(string(hashedPassword), plainPassword)

				Expect(err).NotTo(BeNil())
			})
		})
	})
})
