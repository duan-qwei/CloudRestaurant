package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type BaseModel struct {
	Id         int64                 `gorm:"primary_key" json:"id"`
	UpdateTime time.Time             `json:"-"`
	CreateTime time.Time             `json:"-"`
	isDelete   soft_delete.DeletedAt `json:"isDelete" gorm:"softDelete:flag"`
}
