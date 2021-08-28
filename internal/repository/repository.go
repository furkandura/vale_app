package repository

import (
	"vale_app/configs/db"

	"gorm.io/gorm"
)

type Repositories struct {
	DB *gorm.DB
}

func Get() *Repositories {
	return &Repositories{
		DB: db.DB,
	}
}
