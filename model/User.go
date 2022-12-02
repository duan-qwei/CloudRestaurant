package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        int64  `gorm:"primarykey"`
	Phone     string `gorm:"column:phone"`
	Password  string `gorm:"column:password"`
	Username  string `gorm:"column:username"`
	Email     string `gorm:"column:email"`
	Picture   string `gorm:"column:picture"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
