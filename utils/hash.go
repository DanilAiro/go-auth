package utils

import (
	"golang.org/x/crypto/bcrypt"
)

var lim = 72

func HashToken(token string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(token)[:lim], bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func VerifyToken(token, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(token)[:lim])
	
	return err == nil
}