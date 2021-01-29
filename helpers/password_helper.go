package helpers

import "golang.org/x/crypto/bcrypt"

// EncryptPassword encrypts plain password with bcrypt adaptive hashing algorithm.
func EncryptPassword(plainPassword string) (encryptedPassword string, err error) {
	bytePassword := []byte(plainPassword)
	hashPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	encryptedPassword = string(hashPassword)

	return encryptedPassword, err
}
