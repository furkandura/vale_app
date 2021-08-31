package models

import "gorm.io/gorm"

type Customer struct {
	Base
	FullName  string  `json:"full_name"`
	CompanyId int     `json:"company_id"`
	Phone     *string `json:"phone"`
	Vehicles  string  `json:"vehicles"`
	Note      *string `json:"note"`
}
