package service

import (
	"context"
	"xyzApp/app/model/dto"
)

type ITenorService interface {
	InsertLimit(ctx context.Context, request *dto.InsertLimitRequest) (*dto.InsertLimitResponse, error)
}
