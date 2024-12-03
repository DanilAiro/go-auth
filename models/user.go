package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Ip 				string
	Email 			string `gorm:"unique"`
	RefreshToken 	string
}