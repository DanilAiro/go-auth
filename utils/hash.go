package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashToken(token string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(token)[:72], bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func VerifyToken(token, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(token))
	
	return err == nil
}