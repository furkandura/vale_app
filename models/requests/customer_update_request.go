package requests

type CustomerUpdateRequest struct {
	CustomerId int     `json:"customer_id" validate:"required"`
	FullName   string  `json:"full_name" validate:"required"`
	Phone      *string `json:"phone"`
	Vehicles   string  `json:"vehicles" validate:"required"`
	Note       *string `json:"note"`
}
