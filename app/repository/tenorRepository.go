package repository

import (
	"context"
	"database/sql"
	"xyzApp/app/model/entity"
)

type TonerRepository struct {
	db *sql.DB
}

// func provider
func NewTenorRepository(db *sql.DB) ITenorRepository {
	return &TonerRepository{db}
}

// method insert data tenor
func (t *TonerRepository) Insert(ctx context.Context, input *entity.Tenor) (*entity.Tenor, error) {
	//TODO implement me
	panic("implement me")
}
