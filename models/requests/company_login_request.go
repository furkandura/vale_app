package requests

// Endpoint => api/company/login

type CompanyLoginRequest struct {
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}
