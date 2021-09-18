package requests

// Endpoint => api/customer/update
// Method => POST

type CustomerUpdateRequest struct {
	CustomerId int     `json:"customer_id" validate:"required"`
	FullName   string  `json:"full_name" validate:"required"`
	Phone      *string `json:"phone"`
	Vehicles   string  `json:"vehicles" validate:"required"`
	Note       *string `json:"note"`
}
