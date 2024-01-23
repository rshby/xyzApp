package repository

import (
	"context"
	"xyzApp/app/model/entity"
)

type ITransactionRepository interface {
	Insert(ctx context.Context, input *entity.Transaction) (*entity.Transaction, error)
}
