package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	GUID 			string `gorm:"not null;unique" json:"guid"`
	Email 			string `gorm:"not null;unique" json:"e-mail"`
	Ip				string `gorm:"not null"`
	RefreshToken	string `gorm:"not null;unique"`
}