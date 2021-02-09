package helpers

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes plain password with bcrypt adaptive hashing algorithm.
func HashPassword(plainPassword string) (hashedPassword string, err error) {
	bytePassword := []byte(plainPassword)
	hashPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	hashedPassword = string(hashPassword)

	return hashedPassword, err
}
