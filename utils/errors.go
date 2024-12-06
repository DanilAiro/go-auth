package utils

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadBodyError(c *gin.Context) {
	text := "Failed to read body"
	c.JSON(http.StatusBadRequest, gin.H{
		"error": text,
	})

	log.Println(text)
}

func DefaultError(text string, err error, c *gin.Context) bool {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": text + err.Error(),
		})

		log.Println(text)
		return true
	}

	return false
}

func ValidTokenError(token, hash string, c *gin.Context) bool {
	if !VerifyToken(token, hash) {
		text := "Refresh token is not valid"
		
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Refresh token is not valid",
		})

		log.Println(text)
		return true
	}

	return false
}