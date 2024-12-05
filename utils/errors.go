package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadBodyError(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": "Failed to read body",
	})
}

func DefaultError(text string, err error, c *gin.Context) bool {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": text + err.Error(),
		})

		return true
	}

	return false
}

func ValidTokenError(token, hash string, c *gin.Context) bool {
	if !VerifyToken(token, hash) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Refresh token is not valid",
		})

		return true
	}

	return false
}