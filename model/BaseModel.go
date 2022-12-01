package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type BaseModel struct {
	Id         int64                 `gorm:"primary_key" json:"id"`
	UpdateTime time.Time             `json:"updateTime" gorm:"autoUpdateTime"`
	CreateTime time.Time             `json:"createTime" gorm:"autoCreateTime"`
	isDelete   soft_delete.DeletedAt `json:"isDelete" gorm:"softDelete:flag"`
}
