package model

import "gorm.io/gorm"

type Point struct {
	gorm.Model
	Points int32 `gorm:"column:points;not null default 0"`
	UserId int64 `gorm:"column:user_id;not null"`
}
