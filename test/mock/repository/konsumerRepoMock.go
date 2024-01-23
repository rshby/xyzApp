package mock

import (
	"context"
	"github.com/stretchr/testify/mock"
	modelEntity "xyzApp/app/model/entity"
)

type KonsumerRepoMock struct {
	Mock mock.Mock
}

// function provider
func NewKonsumerRepoMock() *KonsumerRepoMock {
	return &KonsumerRepoMock{Mock: mock.Mock{}}
}

func (k *KonsumerRepoMock) Insert(ctx context.Context, entity *modelEntity.Konsumer) (*modelEntity.Konsumer, error) {
	args := k.Mock.Called(ctx, entity)

	value := args.Get(0)
	if value == nil {
		return nil, args.Error(1)
	}

	return value.(*modelEntity.Konsumer), nil
}

func (k *KonsumerRepoMock) GetAll(ctx context.Context) ([]modelEntity.Konsumer, error) {
	args := k.Mock.Called(ctx)

	value := args.Get(0)
	if value == nil {
		return nil, args.Error(1)
	}

	return value.([]modelEntity.Konsumer), nil
}

func (k *KonsumerRepoMock) GetByNik(ctx context.Context, nik string) (*modelEntity.Konsumer, error) {
	args := k.Mock.Called(ctx, nik)

	value := args.Get(0)
	if value == nil {
		return nil, args.Error(1)
	}

	return value.(*modelEntity.Konsumer), nil
}
