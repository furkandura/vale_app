package requests

import "time"

type ParkingUpdateRequest struct {
	ParkingId      int        `json:"parking_id" validate:"required"`
	CustomerId     int        `json:"customer_id" validate:"required"`
	Type           int8       `json:"type" validate:"required"`
	Plate          string     `json:"plate" validate:"required"`
	DateOfReceipt  time.Time  `json:"date_of_receipt" validate:"required"`
	DateOfDelivery *time.Time `json:"date_of_delivery"`
	Amount         *float64   `json:"amount"`
	Note           *string    `json:"note"`
}
