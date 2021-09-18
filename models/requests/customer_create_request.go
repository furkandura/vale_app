package requests

// Endpoint => api/customer/create
// Method => POST

type CustomerCreateRequest struct {
	FullName string  `json:"full_name" validate:"required"`
	Phone    *string `json:"phone"`
	Vehicles string  `json:"vehicles" validate:"required"`
	Note     *string `json:"note"`
}
