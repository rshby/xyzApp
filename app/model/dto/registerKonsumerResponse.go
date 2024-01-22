package dto

type RegisterKonsumerResponse struct {
	Nik          string  `json:"nik,omitempty"`
	FullName     string  `json:"full_name,omitempty"`
	LegalName    string  `json:"legal_name,omitempty"`
	TempatLahir  string  `json:"tempat_lahir,omitempty"`
	TanggalLahir string  `json:"tanggal_lahir,omitempty"`
	Gaji         float64 `json:"gaji,omitempty"`
	FotoKtp      string  `json:"foto_ktp,omitempty"`
	FotoSelfie   string  `json:"foto_selfie,omitempty"`
	CreatedAt    string  `json:"created_at,omitempty"`
	UpdatedAt    string  `json:"updated_at,omitempty"`
}
