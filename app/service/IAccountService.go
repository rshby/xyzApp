package service

import (
	"context"
	"xyzApp/app/model/dto"
)

type IAccountService interface {
	Register(ctx context.Context, request *dto.RegisterAccount) error
	Login(ctx context.Context, request *dto.LoginRequest) (*dto.LoginResponse, error)
}
