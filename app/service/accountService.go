package service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/opentracing/opentracing-go"
	"sync"
	"xyzApp/app/customError"
	"xyzApp/app/helper"
	"xyzApp/app/model/dto"
	"xyzApp/app/model/entity"
	"xyzApp/app/repository"
)

type AccountService struct {
	Validate     *validator.Validate
	AccountRepo  repository.IAccountRepository
	KonsumerRepo repository.IKonsumerRepository
}

// function provider
func NewAccountService(validate *validator.Validate, accRepo repository.IAccountRepository,
	konsumerRepo repository.IKonsumerRepository) IAccountService {
	return &AccountService{
		Validate:     validate,
		AccountRepo:  accRepo,
		KonsumerRepo: konsumerRepo,
	}
}

// method register
func (a *AccountService) Register(ctx context.Context, request *dto.RegisterAccount) error {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx, "AccountService Register")
	defer span.Finish()

	// validate
	if err := a.Validate.StructCtx(ctxTracing, *request); err != nil {
		return err
	}

	wg := &sync.WaitGroup{}

	errKonsumerChan := make(chan error, 1)
	errAccountChan := make(chan error, 1)
	defer func() {
		close(errAccountChan)
		close(errKonsumerChan)
	}()

	// cek apakah nik sudah ada di database
	wg.Add(1)
	go func(ctx context.Context, wg *sync.WaitGroup, nik string, errChan chan error) {
		defer wg.Done()
		_, err := a.KonsumerRepo.GetByNik(ctx, nik)
		errChan <- err
	}(ctxTracing, wg, request.Nik, errKonsumerChan)

	// cek apakah email sudah ada di database
	wg.Add(1)
	go func(ctx context.Context, wg *sync.WaitGroup, email string, errChan chan error) {
		defer wg.Done()
		_, err := a.AccountRepo.GetByEmail(ctxTracing, request.Email)
		errChan <- err
	}(ctxTracing, wg, request.Email, errAccountChan)

	// hash password
	hashedPassword, err := helper.HashPassword(request.Password)
	if err != nil {
		return customError.NewInternalSeverError("failed to hash password")
	}

	// wait all task
	wg.Wait()

	// cek
	if err := <-errKonsumerChan; err != nil {
		return customError.NewNotFoundError("record konsumer with this nik not found")
	}

	if err := <-errAccountChan; err == nil {
		return customError.NewBadRequestError("data with same email already exist")
	}

	// create entity
	input := entity.Account{
		Nik:      request.Nik,
		Email:    request.Email,
		Password: hashedPassword,
	}

	// insert
	_, err = a.AccountRepo.Insert(ctxTracing, &input)
	if err != nil {
		return err
	}

	// success insert
	return nil
}

// method login
func (a *AccountService) Login(ctx context.Context, request *dto.LoginRequest) (*dto.LoginResponse, error) {
	//TODO implement me
	panic("implement me")
}
