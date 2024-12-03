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
	authG := router.Group("/api")

	authG.POST("/token", controllers.GetTokens)
	authG.POST("/token/refresh", controllers.RefreshTokens)

	router.Run(":" + os.Getenv("PORT"))
}