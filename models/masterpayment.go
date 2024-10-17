package models

import "github.com/jinzhu/gorm"

type MasterPayment struct {
	gorm.Model
	Tipe  string `gorm:"type:varchar(150)"`
	Name  string `gorm:"type:varchar(150)"`
	Photo string `gorm:"type:varchar(200);null"`
}

type ResMasterPayment struct {
	ID   uint   `json:"id"`
	Tipe string `json:"tipe"`
	Name string `json:"name"`
}
