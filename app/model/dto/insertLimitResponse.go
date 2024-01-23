package dto

type InsertLimitResponse struct {
	Nik       string  `json:"nik,omitempty"`
	Bulan     string  `json:"bulan,omitempty"`
	StartDate string  `json:"start_date,omitempty"`
	EndDate   string  `json:"end_date,omitempty"`
	Limit     float64 `json:"limit,omitempty"`
}
