package mock

import (
	"context"
	"github.com/stretchr/testify/mock"
	modelEntity "xyzApp/app/model/entity"
)

type TransactionRepoMock struct {
	Mock mock.Mock
}

// function provider
func NewTransactionRepoMock() *TransactionRepoMock {
	return &TransactionRepoMock{Mock: mock.Mock{}}
}

func (t *TransactionRepoMock) Insert(ctx context.Context, input *modelEntity.Transaction) (*modelEntity.Transaction, error) {
	args := t.Mock.Called(ctx, input)

	value := args.Get(0)

	if value == nil {
		return nil, args.Error(1)
	}

	return value.(*modelEntity.Transaction), nil
}
