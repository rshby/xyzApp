package repository

import (
	"context"
	"time"
	"xyzApp/app/model/entity"
)

type ITenorRepository interface {
	Insert(ctx context.Context, input *entity.Tenor) (*entity.Tenor, error)
	GetByNikAndDate(ctx context.Context, nik string, dateTransaction time.Time) (*entity.Tenor, error)
}
