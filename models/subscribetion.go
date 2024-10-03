package models

import "github.com/jinzhu/gorm"

type Subscription struct {
	gorm.Model
	Name  string `gorm:"type:varchar(150)"`
	Days  int64  `gorm:"default:0"`
	Price int64  `gorm:"default:0"`
	Tipe  string `gorm:"column:tipe;type:enum('Retail','F&B','All')"`
	Desc  string `gorm:"type:text"`
}

type ResSub struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Days  int64  `json:"days"`
	Price int64  `json:"price"`
	Tipe  string `json:"tipe"`
	Desc  string `json:"desc"`
}
