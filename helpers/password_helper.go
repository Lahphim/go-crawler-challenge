package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes plain password with bcrypt adaptive hashing algorithm.
func HashPassword(plainPassword string) (hashedPassword string, err error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	hashedPassword = string(hashPassword)

	return hashedPassword, err
}

// CheckMatchPassword checks match password between plain password and hashed password
func CheckMatchPassword(hashedPassword string, plainPassword string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}
