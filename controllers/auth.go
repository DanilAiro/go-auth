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

	var body struct {
		GUID	string `json:"guid"`
		Email	string `json:"e-mail"`
		Ip		string
	}

	err = c.Bind(&body)
	if err != nil || body.GUID == "" || body.Email == "" {
		utils.ReadBodyError(c)
		return
	}

	body.Ip = c.ClientIP()

	accessToken, err := utils.GenerateAccessToken(body.GUID, body.Ip)
	if utils.GenerateTokenError(err, c) {
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(body.GUID, body.Ip)
	if utils.GenerateTokenError(err, c) {
		return
	}

	refreshHash, err := utils.HashToken(refreshToken)
	if utils.GenerateTokenError(err, c) {
		return
	}
	
	user := models.User{
		GUID: body.GUID, 
		Email: body.Email, 
		Ip: body.Ip, 
		RefreshToken: refreshHash,
	}
	err = initializers.DB.Create(&user).Error
	if utils.CreateUserError(err, c) {
		return
	}

	c.SetCookie("JwtAccess", accessToken + ":" + refreshToken, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, "\"ACCESS GRANTED\"©DeusEx")
}

func RefreshTokens(c *gin.Context) {
	cookie, err := c.Cookie("JwtAccess")
	if utils.CookieNotFoundError(err, c) {
		return
	}

	s := strings.Split(cookie, ":")[1]

	c.JSON(200, s)
}