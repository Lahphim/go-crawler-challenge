package helpers

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateSecretToken generates a slice of 64 random bytes then encode them to a secret string
func GenerateSecretToken() (secret string, err error) {
	byteSecret := make([]byte, 64)
	_, err = rand.Read(byteSecret)
	if err != nil {
		return "", err
	}

	secret = hex.EncodeToString(byteSecret)

	return secret, nil
}
