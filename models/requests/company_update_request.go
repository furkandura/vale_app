package requests

// Endpoint => api/company/update
// Method => POST

type CompanyUpdateRequest struct {
	FullName string  `json:"full_name"  validate:"required"`
	Phone    string  `json:"phone" validate:"required"`
	Password string  `json:"password" validate:"required,min=6,max=16"`
	Email    *string `json:"email" validate:"omitempty,email"`
}
