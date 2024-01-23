package dto

type LoginRequest struct {
	Email    string `validate:"required,email" json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
