package security

import (
	"clickcash_backend/logs"
	"golang.org/x/crypto/bcrypt"
)

func NewEncryptPassword(password string) (string, error) {
	generateFromPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logs.Error(err)
		return "", err
	}
	return string(generateFromPassword), nil
}
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
