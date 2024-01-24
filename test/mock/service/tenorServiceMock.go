package mock

import (
	"context"
	"github.com/stretchr/testify/mock"
	"xyzApp/app/model/dto"
)

type TenorServiceMock struct {
	Mock mock.Mock
}

// function provider
func NewTenorServiceMock() *TenorServiceMock {
	return &TenorServiceMock{Mock: mock.Mock{}}
}

func (t *TenorServiceMock) InsertLimit(ctx context.Context, request *dto.InsertLimitRequest) (*dto.InsertLimitResponse, error) {
	args := t.Mock.Called(ctx, request)

	value := args.Get(0)
	if value == nil {
		return nil, args.Error(1)
	}

	return value.(*dto.InsertLimitResponse), nil
}
