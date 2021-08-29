package models

type Company struct {
	Base
	FullName    string  `json:"full_name"`
	Phone       string  `json:"phone"`
	CompanyName string  `json:"company_name"`
	Password    string  `json:"password"`
	Email       *string `json:"email"`
}
