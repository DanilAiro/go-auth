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

func GenerateTokenError(err error, c *gin.Context) bool {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to generate token: " + err.Error(),
		})

		return true
	}

	return false
}

func HashTokenError(err error, c *gin.Context) bool {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash token: " + err.Error(),
		})

		return true
	}

	return false
}

func CookieNotFoundError(err error, c *gin.Context) bool {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Refresh token not found: " + err.Error(),
		})

		return true
	}

	return false
}

func CreateUserError(err error, c *gin.Context) bool {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user: " + err.Error(),
		})

		return true
	}

	return false
}