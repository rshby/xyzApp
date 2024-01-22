package dto

type RegisterKonsumerRequest struct {
	Nik          string  `validate:"required,min=16" json:"nik,omitempty"`
	FullName     string  `validate:"required" json:"full_name,omitempty"`
	LegalName    string  `validate:"required" json:"legal_name,omitempty"`
	TempatLahir  string  `validate:"required" json:"tempat_lahir,omitempty"`
	TanggalLahir string  `validate:"required" json:"tanggal_lahir,omitempty"`
	Gaji         float64 `validate:"required" json:"gaji,omitempty"`
	FotoKtp      string  `validate:"required" json:"foto_ktp,omitempty"`
	FotoSelfie   string  `validate:"required" json:"foto_selfie,omitempty"`
}
