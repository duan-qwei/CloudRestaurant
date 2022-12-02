package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        uint   `gorm:"primarykey"`
	Password  string `gorm:"column:password"`
	Username  string `gorm:"column:username"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
