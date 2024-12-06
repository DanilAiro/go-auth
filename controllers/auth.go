package controllers

import (
	"net/http"
	"strings"

	"github.com/DanilAiro/go-auth/initializers"
	"github.com/DanilAiro/go-auth/models"
	"github.com/DanilAiro/go-auth/utils"
	"github.com/gin-gonic/gin"
)

func GetTokens(c *gin.Context) {
	_, err := c.Cookie("JwtAccess")
	if err == nil {
		return
	}

	user := models.User{}

	err = c.Bind(&user)
	if err != nil || user.GUID == "" || user.Email == "" {
		utils.ReadBodyError(c)
		return
	}

	user.Ip = c.ClientIP()

	accessToken, refreshToken, err := utils.GenerateTokenPair(user.GUID, user.Ip)
	if utils.DefaultError("Failed to generate token: ", err, c) {
		return
	}

	refreshHash, err := utils.HashToken(refreshToken)
	if utils.DefaultError("Failed to generate token: ", err, c) {
		return
	}

	user.RefreshToken = refreshHash
	
	err = initializers.DB.Create(&user).Error
	if utils.DefaultError("Failed to create user: ", err, c) {
		return
	}

	c.SetCookie("JwtAccess", accessToken + ":" + refreshToken, 3600, "/", "", false, true)

	c.JSON(http.StatusOK, "'ACCESS GRANTED'©DeusEx")
}

func RefreshTokens(c *gin.Context) {
	cookie, err := c.Cookie("JwtAccess")
	if utils.DefaultError("Refresh token not found: ", err, c) {
		return
	}

	var user models.User
	c.Bind(&user)
	err = initializers.DB.First(&user, user.GUID).Error
	if utils.DefaultError("Failed to find user: ", err, c) {
		return
	}
	
	tokenCookie := strings.Split(cookie, ":")
	if len(tokenCookie) != 2 {
		c.JSON(http.StatusBadRequest, "Cookie is not valid")
		return
	}

	if utils.ValidTokenError(tokenCookie[1], user.RefreshToken, c) {
		return
	}
	
	if user.Ip != c.ClientIP() {
		user.Ip = c.ClientIP()

		utils.SendMail(user.Email, "Someone try to log in", "User data: " + user.Ip)
	}
	
	accessToken, refreshToken, err := utils.GenerateTokenPair(user.GUID, user.Ip)
	if utils.DefaultError("Failed to generate token: ", err, c) {
		return
	}
	
	refreshHash, err := utils.HashToken(refreshToken)
	if utils.DefaultError("Failed to generate token: ", err, c) {
		return
	}
	
	user.RefreshToken = refreshHash
	
	err = initializers.DB.Updates(&user).Error
	if utils.DefaultError("Failed to update user: ", err, c) {
		return
	}
	
	c.SetCookie("JwtAccess", accessToken + ":" + refreshToken, 3600, "/", "", false, true)

	c.JSON(http.StatusOK, "Access refreshed")
}