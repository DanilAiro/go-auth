package main

import (
	"os"

	"github.com/DanilAiro/go-auth/controllers"
	"github.com/DanilAiro/go-auth/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}
 
func main() {
	router := gin.Default()

	router.POST("/tokens", controllers.GetTokens)
	router.POST("/tokens/refresh", controllers.RefreshTokens)

	router.Run(":" + os.Getenv("PORT"))
}