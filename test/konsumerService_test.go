package test

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
	"xyzApp/app/customError"
	"xyzApp/app/model/dto"
	"xyzApp/app/model/entity"
	service "xyzApp/app/service"
	mck "xyzApp/test/mock/repository"
)

func TestRegisterKonsumer(t *testing.T) {
	t.Run("test insert konsumer success", func(t *testing.T) {
		validate := validator.New()
		konsumerRepo := mck.NewKonsumerRepoMock()
		service := service.NewKonsumerService(validate, konsumerRepo)

		// mock
		request := dto.RegisterKonsumerRequest{
			Nik:          "1234567890123456",
			FullName:     "Hello World",
			LegalName:    "Hello World",
			TempatLahir:  "Jakarta",
			TanggalLahir: "1999-10-10",
			Gaji:         12000000,
			FotoKtp:      "www.google.com",
			FotoSelfie:   "www.google.com",
		}
		konsumerRepo.Mock.On("Insert", mock.Anything, mock.Anything).
			Return(&entity.Konsumer{
				Nik:          "",
				FullName:     "",
				LegalName:    "",
				TempatLahir:  "",
				TanggalLahir: time.Now(),
				Gaji:         0,
				FotoKtp:      "",
				FotoSelfie:   "",
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			}, nil).Times(1)

		// test
		register, err := service.Register(context.Background(), &request)
		assert.NotNil(t, register)
		assert.Nil(t, err)
		konsumerRepo.Mock.AssertExpectations(t)
	})
	t.Run("test register konsumer error validate", func(t *testing.T) {
		validate := validator.New()
		repo := mck.NewKonsumerRepoMock()
		service := service.NewKonsumerService(validate, repo)

		// test
		register, err := service.Register(context.Background(), &dto.RegisterKonsumerRequest{})
		assert.Error(t, err)
		assert.Nil(t, register)
	})
	t.Run("test register konsumer error parse datetime", func(t *testing.T) {
		validate := validator.New()
		konsumerRepo := mck.NewKonsumerRepoMock()
		service := service.NewKonsumerService(validate, konsumerRepo)

		// test
		register, err := service.Register(context.Background(), &dto.RegisterKonsumerRequest{
			Nik:          "1234567890123456",
			FullName:     "Hello World",
			LegalName:    "Hello World",
			TempatLahir:  "Jakarta",
			TanggalLahir: "123456789",
			Gaji:         10000000,
			FotoKtp:      "www.google.com",
			FotoSelfie:   "www.google.com",
		})

		assert.Nil(t, register)
		assert.NotNil(t, err)
		assert.Error(t, err)
	})
	t.Run("test register konsumer error insert db", func(t *testing.T) {
		validate := validator.New()
		konsumerRepo := mck.NewKonsumerRepoMock()
		service := service.NewKonsumerService(validate, konsumerRepo)

		// mock
		errorMessage := "failed to insert"
		konsumerRepo.Mock.On("Insert", mock.Anything, mock.Anything).
			Return(nil, customError.NewInternalSeverError(errorMessage))

		// test
		register, err := service.Register(context.Background(), &dto.RegisterKonsumerRequest{
			Nik:          "1234567890123456",
			FullName:     "Hello World",
			LegalName:    "Hello World",
			TempatLahir:  "Bandung",
			TanggalLahir: "1999-10-10",
			Gaji:         12000000,
			FotoKtp:      "www.google.com",
			FotoSelfie:   "www.google.com",
		})

		assert.Nil(t, register)
		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Equal(t, errorMessage, err.Error())
	})
}
