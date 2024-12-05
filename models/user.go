package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	GUID 			string `gorm:"not null;unique"`
	Email 			string `gorm:"not null;unique"`
	Ip				string `gorm:"not null"`
	RefreshToken	string `gorm:"not null;unique"`
}