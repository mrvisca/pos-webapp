package models

import "github.com/jinzhu/gorm"

type PaymentMethod struct {
	gorm.Model
	BusinessId  uint      `gorm:"default:0"`
	WarehouseId uint      `gorm:"default:0"`
	Business    Business  `gorm:"foreignKey:BusinessId"`
	Warehouse   Warehouse `gorm:"foreignKey:WarehouseId"`
	Tipe        string    `gorm:"type:varchar(150)"`
	Name        string    `gorm:"type:varchar(150)"`
	Norek       int64     `gorm:"default:0"`
	Admin       int64     `gorm:"default:0"`
}

type ResponsePaymentMethod struct {
	ID          uint   `json:"id"`
	BusinessId  uint   `json:"business_id"`
	WarehouseId uint   `json:"warehouse_id"`
	Tipe        string `json:"tipe"`
	Name        string `json:"name"`
	Norek       int64  `json:"norek"`
	Admin       int64  `json:"admin"`
}
