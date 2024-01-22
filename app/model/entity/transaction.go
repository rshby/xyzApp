package entity

import "time"

type Transaction struct {
	Id              int       `json:"id,omitempty"`
	ReffNumber      string    `json:"reff_number,omitempty"`
	Nik             string    `json:"nik,omitempty"`
	DateTransaction time.Time `json:"date_transaction,omitempty"`
	Otr             float64   `json:"otr,omitempty"`
	AdminFee        float64   `json:"admin_fee,omitempty"`
	JumlahCicilan   int       `json:"jumlah_cicilan,omitempty"`
	JumlahBunga     float64   `json:"jumlah_bunga,omitempty"`
	Aset            string    `json:"aset,omitempty"`
	TotalDebet      float64   `json:"total_debet,omitempty"`
}
