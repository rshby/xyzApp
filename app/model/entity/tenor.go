package entity

import "time"

type Tenor struct {
	Id        int       `json:"id,omitempty"`
	Nik       string    `json:"nik,omitempty"`
	Bulan     string    `json:"bulan,omitempty"`
	StartDate time.Time `json:"start_date,omitempty"`
	EndDate   time.Time `json:"end_date,omitempty"`
	Tenor     float64   `json:"tenor,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
