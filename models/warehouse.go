package models

import "github.com/jinzhu/gorm"

type Warehouse struct {
	gorm.Model
	BusinessId     uint         `gorm:"default:0"`
	SubscriptionId uint         `gorm:"default:0"`
	Business       Business     `gorm:"foreignKey:BusinessId"`
	Subscription   Subscription `gorm:"foreignKey:SubscriptionId"`
	Name           string       `gorm:"type:varchar(200)"`
	Address        string       `gorm:"type:text"`
	Phone          string       `gorm:"type:varchar(20)"`
	Join           string
	EndDate        string
	IsDefault      bool `gorm:"type:boolean"`
}

type SupportCabang struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
