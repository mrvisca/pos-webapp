package models

import "github.com/jinzhu/gorm"

type Client struct {
	gorm.Model
	BusinessId  uint      `gorm:"default:0"`
	WarehouseId uint      `gorm:"default:0"`
	Business    Business  `gorm:"foreignKey:BusinessId"`
	Warehouse   Warehouse `gorm:"foreignKey:WarehouseId"`
	Name        string    `gorm:"type:varchar(200)"`
	Email       string    `gorm:"unique;null"`
	Phone       string    `gorm:"type:varchar(20);null"`
	Address     string    `gorm:"type:text;null"`
	Join        string
}

type Pelanggan struct {
	ID          uint   `json:"id"`
	BusinessId  uint   `json:"business_id"`
	WarehouseId uint   `json:"warehouse_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	Join        string `json:"join"`
}
