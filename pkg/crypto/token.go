package crypto

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// GenerateToken generates new token using bscrypt * sha256
// THOUGHTS: this may need a proper hash
func GenerateToken(secret string, email string) (string, error) {
	if email == "" {
		return "", fmt.Errorf("empty email")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(email+secret), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("bcrypt gen password: %v", err)
	}

	return Sha256(string(hash)), nil
}
