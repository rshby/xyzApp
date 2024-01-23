package mock

import (
	"context"
	"github.com/stretchr/testify/mock"
	"time"
	modelEntity "xyzApp/app/model/entity"
)

type TenorRepoMock struct {
	Mock mock.Mock
}

// function provider
func NewTenorRepoMock() *TenorRepoMock {
	return &TenorRepoMock{Mock: mock.Mock{}}
}

func (t *TenorRepoMock) Insert(ctx context.Context, input *modelEntity.Tenor) (*modelEntity.Tenor, error) {
	args := t.Mock.Called(ctx, input)

	value := args.Get(0)
	if value == nil {
		return nil, args.Error(1)
	}

	return value.(*modelEntity.Tenor), nil
}

func (t *TenorRepoMock) GetByNikAndDate(ctx context.Context, nik string, dateTransaction time.Time) (*modelEntity.Tenor, error) {
	args := t.Mock.Called(ctx, nik, dateTransaction)

	value := args.Get(0)
	if value == nil {
		return nil, args.Error(1)
	}

	return value.(*modelEntity.Tenor), nil
}

func (t *TenorRepoMock) UpdateLimit(ctx context.Context, limit float64, nik, bulan string) error {
	args := t.Mock.Called(ctx, limit, nik, bulan)
	return args.Error(0)
}
