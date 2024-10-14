package models

import "github.com/jinzhu/gorm"

type PajakLayanan struct {
	gorm.Model
	BusinessId  uint      `gorm:"default:0"`
	WarehouseId uint      `gorm:"default:0"`
	Business    Business  `gorm:"foreignKey:BusinessId"`
	Warehouse   Warehouse `gorm:"foreignKey:WarehouseId"`
	IsTax       bool      `gorm:"type:boolean"`
	Taxval      int64     `gorm:"default:0"`
	IsService   bool      `gorm:"type:bool"`
	Serviceval  int64     `gorm:"default:0"`
}
