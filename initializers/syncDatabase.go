package initializers

import "github.com/DanilAiro/go-auth/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}