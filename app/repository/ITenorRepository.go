package repository

import (
	"context"
	"xyzApp/app/model/entity"
)

type ITenorRepository interface {
	Insert(ctx context.Context, input *entity.Tenor) (*entity.Tenor, error)
}
