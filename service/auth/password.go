package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(p string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}
