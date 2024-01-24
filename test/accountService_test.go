package test

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"xyzApp/app/config"
	"xyzApp/app/customError"
	"xyzApp/app/model/dto"
	"xyzApp/app/model/entity"
	"xyzApp/app/service"
	mck "xyzApp/test/mock/config"
	mockRepo "xyzApp/test/mock/repository"
)

func TestRegisterAccount(t *testing.T) {
	t.Run("test register account error validasi", func(t *testing.T) {
		cfg := mck.NewConfigMock()
		validate := validator.New()
		accountRepo := mockRepo.NewAccountRepoMock()
		konsumerRepo := mockRepo.NewKonsumerRepoMock()
		accountService := service.NewAccountService(cfg, validate, accountRepo, konsumerRepo)

		// test
		request := dto.RegisterAccount{
			Nik:      "123456",
			Email:    "reoreo",
			Password: "123456",
		}
		err := accountService.Register(context.Background(), &request)
		assert.NotNil(t, err)
		assert.Error(t, err)
	})
	t.Run("test register account error konsumer not found", func(t *testing.T) {
		cfg := mck.NewConfigMock()
		validate := validator.New()
		accountRepo := mockRepo.NewAccountRepoMock()
		konsumerRepo := mockRepo.NewKonsumerRepoMock()
		accountService := service.NewAccountService(cfg, validate, accountRepo, konsumerRepo)

		// mock
		errMessage := "record konsumer with this nik not found"
		konsumerRepo.Mock.On("GetByNik", mock.Anything, mock.Anything).
			Return(nil, customError.NewNotFoundError(errMessage))

		accountRepo.Mock.On("GetByEmail", mock.Anything, mock.Anything).
			Return(&entity.Account{}, nil)

		// test
		request := dto.RegisterAccount{
			Nik:      "1234567890123456",
			Email:    "hello@gmail.com",
			Password: "123456",
		}

		err := accountService.Register(context.Background(), &request)
		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Equal(t, errMessage, err.Error())
	})
	t.Run("test register account error account not found", func(t *testing.T) {
		cfg := mck.NewConfigMock()
		validate := validator.New()
		accountRepo := mockRepo.NewAccountRepoMock()
		konsumerRepo := mockRepo.NewKonsumerRepoMock()
		accountService := service.NewAccountService(cfg, validate, accountRepo, konsumerRepo)

		// mock
		konsumerRepo.Mock.On("GetByNik", mock.Anything, mock.Anything).
			Return(&entity.Konsumer{}, nil)

		errorMessage := "data with same email already exist"
		accountRepo.Mock.On("GetByEmail", mock.Anything, mock.Anything).
			Return(&entity.Account{}, nil)

		// test
		reqquest := dto.RegisterAccount{
			Nik:      "1234567890123456",
			Email:    "reo@gmail.com",
			Password: "123456",
		}

		err := accountService.Register(context.Background(), &reqquest)
		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Equal(t, errorMessage, err.Error())
	})
	t.Run("test register account error insert", func(t *testing.T) {
		cfg := mck.NewConfigMock()
		validate := validator.New()
		accountRepo := mockRepo.NewAccountRepoMock()
		konsumerRepo := mockRepo.NewKonsumerRepoMock()
		accountService := service.NewAccountService(cfg, validate, accountRepo, konsumerRepo)

		// mock
		konsumerRepo.Mock.On("GetByNik", mock.Anything, mock.Anything).
			Return(&entity.Konsumer{}, nil)

		accountRepo.Mock.On("GetByEmail", mock.Anything, mock.Anything).
			Return(nil, customError.NewNotFoundError("record not found"))

		errorMessage := "failed to insert"
		accountRepo.Mock.On("Insert", mock.Anything, mock.Anything).
			Return(nil, customError.NewInternalSeverError(errorMessage))

		// test
		request := dto.RegisterAccount{
			Nik:      "1234567890123456",
			Email:    "reo@gmail.com",
			Password: "123456",
		}
		err := accountService.Register(context.Background(), &request)
		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Equal(t, errorMessage, err.Error())
	})
	t.Run("test register success", func(t *testing.T) {
		cfg := mck.NewConfigMock()
		validate := validator.New()
		accountRepo := mockRepo.NewAccountRepoMock()
		konsumerRepo := mockRepo.NewKonsumerRepoMock()
		accountService := service.NewAccountService(cfg, validate, accountRepo, konsumerRepo)

		// mock
		konsumerRepo.Mock.On("GetByNik", mock.Anything, mock.Anything).
			Return(&entity.Konsumer{}, nil)

		accountRepo.Mock.On("GetByEmail", mock.Anything, mock.Anything).
			Return(nil, customError.NewNotFoundError("record not found"))

		accountRepo.Mock.On("Insert", mock.Anything, mock.Anything).
			Return(&entity.Account{}, nil)

		// test
		request := dto.RegisterAccount{
			Nik:      "1234567890123456",
			Email:    "hello@gmail.com",
			Password: "123456",
		}
		err := accountService.Register(context.Background(), &request)
		assert.Nil(t, err)
	})
}

func TestLogin(t *testing.T) {
	t.Run("test login error validate", func(t *testing.T) {
		cfg := mck.NewConfigMock()
		validate := validator.New()
		accountRepo := mockRepo.NewAccountRepoMock()
		konsumerRepo := mockRepo.NewKonsumerRepoMock()
		accountService := service.NewAccountService(cfg, validate, accountRepo, konsumerRepo)

		// test
		request := dto.LoginRequest{
			Email:    "hello",
			Password: "123",
		}

		resultLogin, err := accountService.Login(context.Background(), &request)
		assert.Nil(t, resultLogin)
		assert.NotNil(t, err)
		assert.Error(t, err)
	})
	t.Run("test login error email tidak ada di database", func(t *testing.T) {
		cfg := mck.NewConfigMock()
		validate := validator.New()
		accountRepo := mockRepo.NewAccountRepoMock()
		konsumerRepo := mockRepo.NewKonsumerRepoMock()
		accountService := service.NewAccountService(cfg, validate, accountRepo, konsumerRepo)

		// mock
		errorMessage := "record not found"
		accountRepo.Mock.On("GetByEmail", mock.Anything, mock.Anything).
			Return(nil, customError.NewNotFoundError(errorMessage))

		// test
		request := dto.LoginRequest{
			Email:    "hello@gmail.com",
			Password: "123456",
		}

		resultLogin, err := accountService.Login(context.Background(), &request)
		assert.Nil(t, resultLogin)
		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Equal(t, errorMessage, err.Error())
	})
	t.Run("test login error password not match", func(t *testing.T) {
		cfg := mck.NewConfigMock()
		validate := validator.New()
		accountRepo := mockRepo.NewAccountRepoMock()
		konsumerRepo := mockRepo.NewKonsumerRepoMock()
		accountService := service.NewAccountService(cfg, validate, accountRepo, konsumerRepo)

		// mock
		accountRepo.Mock.On("GetByEmail", mock.Anything, mock.Anything).
			Return(&entity.Account{
				Nik:      "1234567890123456",
				Email:    "hello@gmail.com",
				Password: "123456",
			}, nil)

		// test
		request := dto.LoginRequest{
			Email:    "hello@gmail.com",
			Password: "654321",
		}

		resultLogin, err := accountService.Login(context.Background(), &request)
		assert.Nil(t, resultLogin)
		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Equal(t, "password not match", err.Error())
	})
	t.Run("test login success", func(t *testing.T) {
		cfg := mck.NewConfigMock()
		validate := validator.New()
		accountRepo := mockRepo.NewAccountRepoMock()
		konsumerRepo := mockRepo.NewKonsumerRepoMock()
		accountService := service.NewAccountService(cfg, validate, accountRepo, konsumerRepo)

		// mock
		cfg.Mock.On("GetConfig").Return(&config.AppConfig{
			Jwt: &config.Jwt{SecretKey: "sangatRahasia123"},
		})
		accountRepo.Mock.On("GetByEmail", mock.Anything, mock.Anything).
			Return(&entity.Account{
				Nik:      "1234567890123456",
				Email:    "hello@gmail.com",
				Password: "$2a$10$1TYRPP0PpzcDh3/l9E5L2OQaqfnmvuinpEciquD3qgd34udl4odSC",
			})

		// test
		request := dto.LoginRequest{
			Email:    "hello@gmail.com",
			Password: "123456",
		}

		resultLogin, err := accountService.Login(context.Background(), &request)
		assert.Nil(t, err)
		assert.NotNil(t, resultLogin)
	})
}
