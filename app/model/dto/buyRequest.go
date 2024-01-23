package dto

type BuyRequest struct {
	Nik             string  `validate:"required,min=16" json:"nik,omitempty"`
	DateTransaction string  `validate:"required" json:"date_transaction,omitempty"`
	Otr             float64 `validate:"required,gt=0" json:"otr,omitempty"`
	AdminFee        float64 `validate:"required,gt=0" json:"admin_fee,omitempty"`
	JumlahCicilan   int     `validate:"required" json:"jumlah_cicilan,omitempty"`
	Bunga           float64 `validate:"required" json:"bunga,omitempty"`
	Aset            string  `validate:"required" json:"aset,omitempty"`
}
