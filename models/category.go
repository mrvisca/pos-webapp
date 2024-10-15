package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	BusinessId  uint      `gorm:"default:0"`
	WarehouseId uint      `gorm:"default:0"`
	Business    Business  `gorm:"foreignKey:BusinessId"`
	Warehouse   Warehouse `gorm:"foreignKey:WarehouseId"`
	Name        string    `gorm:"type:varchar(200)"`
	Desc        string    `gorm:"type:text;null"`
	Item        int64     `gorm:"default:0"`
}

type DataCategory struct {
	ID          uint   `json:"id"`
	BusinessId  uint   `json:"business_id"`
	WarehouseId uint   `json:"warehouse_id"`
	Name        string `json:"name"`
	Desc        string `json:"desc"`
	Item        int64  `json:"item"`
}
