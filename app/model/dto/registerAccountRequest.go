package dto

type RegisterAccount struct {
	Nik      string `validate:"required,min=16" json:"nik,omitempty"`
	Email    string `validate:"required,email" json:"email,omitempty"`
	Password string `validate:"required,min=6" json:"password,omitempty"`
}
