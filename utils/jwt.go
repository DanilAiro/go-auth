package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(guid, ip string) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub": guid,
		"user_ip": ip,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	
	tokenString, err = token.SignedString([]byte(os.Getenv("SECRET")))
	return
}

func GenerateRefreshToken(guid, ip string) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"user_ip": ip,
	})
	
	tokenString, err = token.SignedString([]byte(os.Getenv("SECRET")))
	return
}