package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        int64  `gorm:"primarykey"`
	Phone     string `gorm:"column:phone;default:null"`
	Password  string `gorm:"column:password"`
	Username  string `gorm:"column:username"`
	Email     string `gorm:"column:email;default:null"`
	Picture   string `gorm:"column:picture;default:null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
