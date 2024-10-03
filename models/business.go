package models

import "github.com/jinzhu/gorm"

type Business struct {
	gorm.Model
	Name        string `gorm:"type:varchar(200)"`
	Branchlimit int64  `gorm:"default:1"`
	Tipe        string `gorm:"column:tipe;type:enum('Retail','F&B')"`
}
