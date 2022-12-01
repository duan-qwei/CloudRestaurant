package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Password string `gorm:"column:password"`
	Username string `gorm:"column:username"`
}
