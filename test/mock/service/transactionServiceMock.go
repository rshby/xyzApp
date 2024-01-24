package mock

import (
	"context"
	"github.com/stretchr/testify/mock"
	"xyzApp/app/model/dto"
)

type TransactionServiceMock struct {
	Mock mock.Mock
}

// function provider
func NewTransactionServiceMock() *TransactionServiceMock {
	return &TransactionServiceMock{Mock: mock.Mock{}}
}

func (t *TransactionServiceMock) Buy(ctx context.Context, request *dto.BuyRequest) (*dto.BuyResponse, error) {
	args := t.Mock.Called(ctx, request)

	value := args.Get(0)
	if value == nil {
		return nil, args.Error(1)
	}

	return value.(*dto.BuyResponse), nil
}
