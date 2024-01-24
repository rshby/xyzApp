package mock

import (
	"context"
	"github.com/stretchr/testify/mock"
	modelEntity "xyzApp/app/model/entity"
)

type AccountRepoMock struct {
	Mock mock.Mock
}

// function provider
func NewAccountRepoMock() *AccountRepoMock {
	return &AccountRepoMock{
		Mock: mock.Mock{},
	}
}

// method insert
func (a *AccountRepoMock) Insert(ctx context.Context, entity *modelEntity.Account) (*modelEntity.Account, error) {
	args := a.Mock.Called(ctx, entity)

	value := args.Get(0)

	if value == nil {
		return nil, args.Error(1)
	}

	return value.(*modelEntity.Account), nil
}

// method get by email
func (a *AccountRepoMock) GetByEmail(ctx context.Context, email string) (*modelEntity.Account, error) {
	args := a.Mock.Called(ctx, email)

	value := args.Get(0)
	if value == nil {
		return nil, args.Error(1)
	}

	return value.(*modelEntity.Account), nil
}
