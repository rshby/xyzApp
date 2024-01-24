package mock

import (
	"context"
	"github.com/stretchr/testify/mock"
	"xyzApp/app/model/dto"
)

type AccountServiceMock struct {
	Mock mock.Mock
}

// function provider
func NewAccountServiceMock() *AccountServiceMock {
	return &AccountServiceMock{Mock: mock.Mock{}}
}

func (a *AccountServiceMock) Register(ctx context.Context, request *dto.RegisterAccount) error {
	args := a.Mock.Called(ctx, request)
	value := args.Get(0)
	if value == nil {
		return nil
	}

	return args.Error(0)
}

func (a *AccountServiceMock) Login(ctx context.Context, request *dto.LoginRequest) (*dto.LoginResponse, error) {
	args := a.Mock.Called(ctx, request)
	value := args.Get(0)
	if value == nil {
		return nil, args.Error(1)
	}

	return value.(*dto.LoginResponse), nil
}
