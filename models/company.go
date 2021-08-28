package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	FullName    string  `json:"full_name"`
	Phone       string  `json:"phone"`
	CompanyName string  `json:"company_name"`
	Password    string  `json:"password"`
	Email       *string `json:"email"`
}
