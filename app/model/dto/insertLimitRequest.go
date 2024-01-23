package dto

type InsertLimitRequest struct {
	Nik   string  `validate:"required,min=16" json:"nik,omitempty"`
	Bulan int     `validate:"required,gt=0" json:"bulan,omitempty"`
	Tahun int     `validate:"required,gt=0" json:"tahun,omitempty"`
	Limit float64 `validate:"required" json:"limit,omitempty"`
}
