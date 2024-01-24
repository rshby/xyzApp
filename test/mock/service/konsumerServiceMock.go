package mock

import (
	"context"
	"github.com/stretchr/testify/mock"
	"xyzApp/app/model/dto"
)

type KonsumerServiceMock struct {
	Mock mock.Mock
}

// function provider
func NewKonsumerServiceMock() *KonsumerServiceMock {
	return &KonsumerServiceMock{Mock: mock.Mock{}}
}

func (k *KonsumerServiceMock) Register(ctx context.Context, request *dto.RegisterKonsumerRequest) (*dto.RegisterKonsumerResponse, error) {
	args := k.Mock.Called(ctx, request)
	value := args.Get(0)
	if value == nil {
		return nil, args.Error(1)
	}

	return value.(*dto.RegisterKonsumerResponse), nil
}
