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
	Kode        string    `gorm:"type:varchar(10)"`
}

type Staff struct {
	ID          uint   `json:"id"`
	RoleId      uint   `json:"role_id"`
	RoleName    string `json:"role_name"`
	BusinessId  uint   `json:"business_id"`
	WarehouseId uint   `json:"warehouse_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	IsVerified  bool   `json:"is_verified"`
	Kode        string `json:"kode"`
}
