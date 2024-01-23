package service

import (
	"context"
	"xyzApp/app/model/dto"
)

type ITransactionService interface {
	Buy(ctx context.Context, request *dto.BuyRequest) (*dto.BuyResponse, error)
}
