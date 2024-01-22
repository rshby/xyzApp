package repository

import (
	"context"
	"xyzApp/app/model/entity"
)

type IKonsumerRepository interface {
	// insert
	Insert(ctx context.Context, entity *entity.Konsumer) (*entity.Konsumer, error)

	// get all data
	GetAll(ctx context.Context) ([]entity.Konsumer, error)

	// get by nik
	GetByNik(ctx context.Context, nik string) (*entity.Konsumer, error)
}
