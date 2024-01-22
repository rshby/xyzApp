package service

import (
	"context"
	"xyzApp/app/model/dto"
)

type IKonsumerService interface {
	Register(ctx context.Context, request *dto.RegisterKonsumerRequest) (*dto.RegisterKonsumerResponse, error)
}
