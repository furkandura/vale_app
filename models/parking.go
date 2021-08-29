package models

import (
	"time"

	"gorm.io/gorm"
)

type Parking struct {
	gorm.Model
	CompanyId      int        `json:"company_id"`
	CustomerId     int        `json:"customer_id"`
	Type           int8       `json:"type" gorm:"default:1"` // 1 => Müşteriden teslim alındı. 2 => Müşteriye teslim edildi.
	Plate          string     `json:"plate"`
	DateOfReceipt  time.Time  `json:"date_of_receipt"`
	DateOfDelivery *time.Time `json:"date_of_delivery"`
	Amount         *float64   `json:"amount"`
	Note           *string    `json:"note"`
}
