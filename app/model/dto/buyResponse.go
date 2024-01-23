package dto

type BuyResponse struct {
	ReffNumber    string  `json:"reff_number,omitempty"`
	Nik           string  `json:"nik,omitempty"`
	Otr           float64 `json:"otr,omitempty"`
	AdminFee      float64 `json:"admin_fee,omitempty"`
	JumlahCicilan int     `json:"jumlah_cicilan,omitempty"`
	Bunga         float64 `json:"bunga,omitempty"`
	Aset          string  `json:"aset,omitempty"`
	TotalDebet    float64 `json:"total_debet,omitempty"`
	SisaLimit     float64 `json:"sisa_limit,omitempty"`
}
