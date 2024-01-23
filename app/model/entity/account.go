package entity

type Account struct {
	Nik      string `json:"nik,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
