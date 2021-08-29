package db

import (
	"fmt"
	"vale_app/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	Connect()
	AutoMigrate()
}

func Connect() {

	connectString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", Username, Password, Host, Port, Database)

	var err error

	DB, err = gorm.Open(mysql.Open(connectString), &gorm.Config{})

	if err != nil {
		panic("Veritabanı bağlantısı yapılamadı, Hata mesajı : " + err.Error())
	}

}

func AutoMigrate() {
	DB.AutoMigrate(
		&models.Company{},
		&models.Parking{},
		&models.Customer{},
	)
}
