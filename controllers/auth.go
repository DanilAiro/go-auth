package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTokens(c *gin.Context) {
	c.JSON(http.StatusOK, "Ура, авторизация прошла успешно!")
}

func RefreshTokens(c *gin.Context) {
	c.JSON(http.StatusOK, "Ура, обновление данных прошло успешно!")
}