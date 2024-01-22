package entity

import "time"

type Konsumer struct {
	Nik          string    `json:"nik,omitempty"`
	FullName     string    `json:"full_name,omitempty"`
	LegalName    string    `json:"legal_name,omitempty"`
	TempatLahir  string    `json:"tempat_lahir,omitempty"`
	TanggalLahir time.Time `json:"tanggal_lahir,omitempty"`
	Gaji         float64   `json:"gaji,omitempty"`
	FotoKtp      string    `json:"foto_ktp,omitempty"`
	FotoSelfie   string    `json:"foto_selfie,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}
