package models

import "gorm.io/gorm"

type Token struct {
	gorm.Model
	UserID       uint
	AccessToken  string
	RefreshToken string
}
