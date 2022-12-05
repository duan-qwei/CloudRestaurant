package model

import (
	"gorm.io/gorm"
	"time"
)

type Role struct {
	Id        int64  `gorm:"primarykey"`
	Name      string `gorm:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
