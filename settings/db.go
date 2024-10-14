package settings

import (
	"pos-webapp/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open("mysql", "root:@(localhost)/posweb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Koneksi ke database gagal!")
	}

	DB.AutoMigrate(&models.Role{})
	DB.AutoMigrate(&models.Subscription{})
	DB.AutoMigrate(&models.Business{})
	DB.AutoMigrate(&models.Warehouse{}).AddForeignKey("business_id", "businesses(id)", "NO ACTION", "NO ACTION").AddForeignKey("subscription_id", "subscriptions(id)", "NO ACTION", "NO ACTION")
	DB.AutoMigrate(&models.User{}).AddForeignKey("business_id", "businesses(id)", "NO ACTION", "NO ACTION").AddForeignKey("warehouse_id", "warehouses(id)", "NO ACTION", "NO ACTION").AddForeignKey("role_id", "roles(id)", "NO ACTION", "NO ACTION")
	DB.AutoMigrate(&models.Client{}).AddForeignKey("business_id", "businesses(id)", "NO ACTION", "NO ACTION").AddForeignKey("warehouse_id", "warehouses(id)", "NO ACTION", "NO ACTION")

	DB.Model(&models.Warehouse{}).Related(&models.Business{}).Related(&models.Subscription{})
	DB.Model(&models.User{}).Related(&models.Business{}).Related(&models.Warehouse{}).Related(&models.Role{})
	DB.Model(&models.Client{}).Related(&models.Business{}).Related(&models.Warehouse{})
}
