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
	"xyzApp/app/service"
	mck "xyzApp/test/mock/repository"
)

func TestInserLimit(t *testing.T) {
	t.Run("test insert limit success", func(t *testing.T) {
		validate := validator.New()
		tenorRepo := mck.NewTenorRepoMock()
		konsumerRepo := mck.NewKonsumerRepoMock()
		tenorService := service.NewTenorService(validate, tenorRepo, konsumerRepo)

		// mock
		konsumerRepo.Mock.On("GetByNik", mock.Anything, mock.Anything).
			Return(&entity.Konsumer{}, nil)
		tenorRepo.Mock.On("Insert", mock.Anything, mock.Anything).
			Return(&entity.Tenor{
				Id:        1,
				Nik:       "1234567890123456",
				Bulan:     "Januari",
				StartDate: time.Now(),
				EndDate:   time.Now(),
				Tenor:     1000000,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}, nil)

		// test
		request := dto.InsertLimitRequest{
			Nik:   "1234567890123456",
			Bulan: 1,
			Tahun: 2024,
			Limit: 1000000,
		}

		result, err := tenorService.InsertLimit(context.Background(), &request)
		assert.NotNil(t, result)
		assert.Nil(t, err)
		assert.Equal(t, "1234567890123456", result.Nik)
		assert.Equal(t, float64(1000000), result.Limit)
	})
	t.Run("test insert limit failed validation", func(t *testing.T) {
		validate := validator.New()
		konsumerRepo := mck.NewKonsumerRepoMock()
		tenorRepo := mck.NewTenorRepoMock()
		tenorService := service.NewTenorService(validate, tenorRepo, konsumerRepo)

		// test
		request := dto.InsertLimitRequest{
			Nik:   "123",
			Bulan: 0,
			Tahun: 0,
			Limit: 0,
		}

		result, err := tenorService.InsertLimit(context.Background(), &request)
		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Nil(t, result)
	})
	t.Run("test insert limit failed konsumer not found", func(t *testing.T) {
		validate := validator.New()
		konsumerRepo := mck.NewKonsumerRepoMock()
		tenorRepo := mck.NewTenorRepoMock()
		tenorService := service.NewTenorService(validate, tenorRepo, konsumerRepo)

		// mock
		errorMessage := "record konsumer tidak ditemukan"
		konsumerRepo.Mock.On("GetByNik", mock.Anything, mock.Anything).
			Return(nil, customError.NewNotFoundError(errorMessage))

		// test
		result, err := tenorService.InsertLimit(context.Background(), &dto.InsertLimitRequest{
			Nik:   "1234567890123456",
			Bulan: 1,
			Tahun: 2024,
			Limit: 1000000,
		})
		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Equal(t, errorMessage, err.Error())
	})
	t.Run("test insert limit failed bulan lebih dari 12", func(t *testing.T) {
		validate := validator.New()
		konsumerRepo := mck.NewKonsumerRepoMock()
		tenorRepo := mck.NewTenorRepoMock()
		tenorService := service.NewTenorService(validate, tenorRepo, konsumerRepo)

		// mock
		konsumerRepo.Mock.On("GetByNik", mock.Anything, mock.Anything).
			Return(&entity.Konsumer{}, nil)

		// test
		request := dto.InsertLimitRequest{
			Nik:   "1234567890123456",
			Bulan: 15,
			Tahun: 2024,
			Limit: 1000,
		}
		result, err := tenorService.InsertLimit(context.Background(), &request)

		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Equal(t, "bulan tidak boleh lebih dari 12", err.Error())
	})
	t.Run("test insert limit error when inserted to tenor", func(t *testing.T) {
		validate := validator.New()
		konsumerRepo := mck.NewKonsumerRepoMock()
		tenorRepo := mck.NewTenorRepoMock()
		tenorService := service.NewTenorService(validate, tenorRepo, konsumerRepo)

		// mock
		konsumerRepo.Mock.On("GetByNik", mock.Anything, mock.Anything).
			Return(&entity.Konsumer{}, nil)

		errorMessage := "failed to insert tenor"
		tenorRepo.Mock.On("Insert", mock.Anything, mock.Anything).
			Return(nil, customError.NewInternalSeverError(errorMessage))

		// test
		result, err := tenorService.InsertLimit(context.Background(), &dto.InsertLimitRequest{
			Nik:   "1234567890123456",
			Bulan: 1,
			Tahun: 2024,
			Limit: 1000000,
		})

		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Equal(t, errorMessage, err.Error())
	})
}
