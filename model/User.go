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
	Id       int64  `json:"id"`
	Phone    string `json:"phone"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Picture  string `json:"picture"`
	RoleId   int64  `json:"roleId"`
	RoleName string `json:"roleName"`
	Points   int32  `json:"points"`
	Token    string `json:"token"`
}
