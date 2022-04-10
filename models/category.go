package models

import (
	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
}

type Category struct {
	BaseModel
	Name     string `gorm:"column:name" json:"username"`
	ParentId int    `gorm:"column:parent_id,index:idx_parent_id" json:"parent_id"`
	IsParent bool   `gorm:"column:is_parent" json:"is_parent"`
	Sort     int    `gorm:"column:sort" json:"sort"`
}
