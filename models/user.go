package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	RoleId      uint      `gorm:"default:0"`
	BusinessId  uint      `gorm:"default:0"`
	WarehouseId uint      `gorm:"default:0"`
	Role        Role      `gorm:"foreignKey:RoleId"`
	Business    Business  `gorm:"foreignKey:BusinessId"`
	Warehouse   Warehouse `gorm:"foreignKey:WarehouseId"`
	Name        string    `gorm:"type:varchar(200)"`
	Email       string    `gorm:"unique;not null"`
	Password    string    `gorm:"type:varchar(200)"`
	Phone       string    `gorm:"type:varchar(20)"`
	IsVerified  bool      `gorm:"type:bool"`
}
