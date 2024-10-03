package models

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100)"`
	Desc     string `gorm:"type:text"`
	IsActive bool   `gorm:"type:boolean"`
}

type Resrole struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	IsActive bool   `json:"is_active"`
}
