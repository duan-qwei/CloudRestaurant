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
	RoleId    int64  `gorm:"column:role_id"`
	RoleName  string `json:"roleName"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserLogin struct {
	Id       int64
	Phone    string
	Username string
	Email    string
	Picture  string
	RoleId   int64
	RoleName string
}
