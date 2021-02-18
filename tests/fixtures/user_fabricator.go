package fixtures

import (
	"go-crawler-challenge/models"

	"github.com/onsi/ginkgo"
	"golang.org/x/crypto/bcrypt"
)

func FabricateUser(email string, plainPassword string) (user *models.User) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		ginkgo.Fail("Generate hashed password failed: " + err.Error())
	}

	user = &models.User{
		Email:          email,
		HashedPassword: string(hashPassword),
	}
	_, err = models.AddUser(user)
	if err != nil {
		ginkgo.Fail("Add user failed: " + err.Error())
	}

	return user
}
