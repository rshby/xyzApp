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

func TestBuyTransaction(t *testing.T) {
	t.Run("test buy error validasi", func(t *testing.T) {
		validate := validator.New()
		konsumerRepo := mck.NewKonsumerRepoMock()
		tenorRepo := mck.NewTenorRepoMock()
		transactionRepo := mck.NewTransactionRepoMock()
		transactionService := service.NewTransactionService(validate, konsumerRepo, tenorRepo, transactionRepo)

		// test
		request := dto.BuyRequest{
			Nik:             "123",
			DateTransaction: "",
			Otr:             0,
			AdminFee:        0,
			JumlahCicilan:   0,
			Bunga:           0,
			Aset:            "",
		}
		result, err := transactionService.Buy(context.Background(), &request)
		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Nil(t, result)
	})
	t.Run("test buy error nik not found", func(t *testing.T) {
		validate := validator.New()
		konsumerRepo := mck.NewKonsumerRepoMock()
		tenorRepo := mck.NewTenorRepoMock()
		transactionRepo := mck.NewTransactionRepoMock()
		transactionService := service.NewTransactionService(validate, konsumerRepo, tenorRepo, transactionRepo)

		// mock
		errorMessage := "record konsumer not found"
		konsumerRepo.Mock.On("GetByNik", mock.Anything, mock.Anything).
			Return(nil, customError.NewNotFoundError(errorMessage))

		tenorRepo.Mock.On("GetByNikAndDate", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, customError.NewNotFoundError("record tenor not found"))

		// test
		request := dto.BuyRequest{
			Nik:             "1234567890123456",
			DateTransaction: "2024-01-20 10:00:00",
			Otr:             1000,
			AdminFee:        100,
			JumlahCicilan:   4,
			Bunga:           2,
			Aset:            "sepeda",
		}

		result, err := transactionService.Buy(context.Background(), &request)
		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Equal(t, errorMessage, err.Error())
	})
	t.Run("test but error tenor not found", func(t *testing.T) {
		validate := validator.New()
		konsumerRepo := mck.NewKonsumerRepoMock()
		tenorRepo := mck.NewTenorRepoMock()
		transactionRepo := mck.NewTransactionRepoMock()
		transactionService := service.NewTransactionService(validate, konsumerRepo, tenorRepo, transactionRepo)

		// mock
		konsumerRepo.Mock.On("GetByNik", mock.Anything, mock.Anything).
			Return(&entity.Konsumer{}, nil)

		errorMessage := "record tenor not found"
		tenorRepo.Mock.On("GetByNikAndDate", mock.Anything, mock.Anything, mock.Anything).
			Return(nil, customError.NewNotFoundError(errorMessage))

		// test
		request := dto.BuyRequest{
			Nik:             "12345678901234567",
			DateTransaction: "2024-10-10 12:00:00",
			Otr:             1000000,
			AdminFee:        2500,
			JumlahCicilan:   4,
			Bunga:           5,
			Aset:            "laptop",
		}

		result, err := transactionService.Buy(context.Background(), &request)
		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Equal(t, "no record limit was found by the specified date_transaction", err.Error())
	})
	t.Run("test limit not enough", func(t *testing.T) {
		validate := validator.New()
		konsumerRepo := mck.NewKonsumerRepoMock()
		tenorRepo := mck.NewTenorRepoMock()
		transactionRepo := mck.NewTransactionRepoMock()
		transactionService := service.NewTransactionService(validate, konsumerRepo, tenorRepo, transactionRepo)

		// mock
		konsumerRepo.Mock.On("GetByNik", mock.Anything, mock.Anything).
			Return(&entity.Konsumer{}, nil)

		tenorRepo.Mock.On("GetByNikAndDate", mock.Anything, mock.Anything, mock.Anything).
			Return(&entity.Tenor{
				Id:        1,
				Nik:       "1234567890123456",
				Bulan:     "Januari",
				StartDate: time.Now(),
				EndDate:   time.Now(),
				Tenor:     100,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}, nil)

		// test
		request := dto.BuyRequest{
			Nik:             "1234567890123456",
			DateTransaction: "2024-01-20 10:00:00",
			Otr:             10000000,
			AdminFee:        5000,
			JumlahCicilan:   2,
			Bunga:           50,
			Aset:            "jam tangan",
		}

		result, err := transactionService.Buy(context.Background(), &request)
		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Equal(t, "limit not enough", err.Error())
	})
	t.Run("test buy error update limit", func(t *testing.T) {
		validate := validator.New()
		konsumerRepo := mck.NewKonsumerRepoMock()
		tenorRepo := mck.NewTenorRepoMock()
		transactionRepo := mck.NewTransactionRepoMock()
		transactionService := service.NewTransactionService(validate, konsumerRepo, tenorRepo, transactionRepo)

		// mock
		konsumerRepo.Mock.On("GetByNik", mock.Anything, mock.Anything).
			Return(&entity.Konsumer{}, nil)

		tenorRepo.Mock.On("GetByNikAndDate", mock.Anything, mock.Anything, mock.Anything).
			Return(&entity.Tenor{
				Id:        1,
				Nik:       "1234567890123456",
				Bulan:     "Januari",
				StartDate: time.Now(),
				EndDate:   time.Now(),
				Tenor:     50000000,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}, nil)

		errorMessage := "error when updating data"
		tenorRepo.Mock.On("UpdateLimit", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(customError.NewInternalSeverError(errorMessage))

		// test
		request := dto.BuyRequest{
			Nik:             "1234567890123456",
			DateTransaction: "2024-01-20 19:00:00",
			Otr:             1000,
			AdminFee:        100,
			JumlahCicilan:   4,
			Bunga:           2.5,
			Aset:            "pakaian",
		}

		result, err := transactionService.Buy(context.Background(), &request)
		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Equal(t, errorMessage, err.Error())
	})
	t.Run("test buy error insert transaction", func(t *testing.T) {
		validate := validator.New()
		konsumerRepo := mck.NewKonsumerRepoMock()
		tenorRepo := mck.NewTenorRepoMock()
		transactionRepo := mck.NewTransactionRepoMock()
		transactionService := service.NewTransactionService(validate, konsumerRepo, tenorRepo, transactionRepo)

		// mock
		konsumerRepo.Mock.On("GetByNik", mock.Anything, mock.Anything).
			Return(&entity.Konsumer{}, nil)

		tenorRepo.Mock.On("GetByNikAndDate", mock.Anything, mock.Anything, mock.Anything).
			Return(&entity.Tenor{
				Id:        1,
				Nik:       "1234567890123456",
				Bulan:     "Januari",
				StartDate: time.Now(),
				EndDate:   time.Now(),
				Tenor:     50000000,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}, nil)

		tenorRepo.Mock.On("UpdateLimit", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil)

		errorMessage := "cant insert data transaction"
		transactionRepo.Mock.On("Insert", mock.Anything, mock.Anything).
			Return(nil, customError.NewInternalSeverError(errorMessage))

		// test
		request := dto.BuyRequest{
			Nik:             "1234567890123456",
			DateTransaction: "2024-01-20 19:00:00",
			Otr:             2000000,
			AdminFee:        5000,
			JumlahCicilan:   2,
			Bunga:           2,
			Aset:            "laptop",
		}

		result, err := transactionService.Buy(context.Background(), &request)
		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Equal(t, errorMessage, err.Error())
	})
	t.Run("test buy success", func(t *testing.T) {
		validate := validator.New()
		konsumerRepo := mck.NewKonsumerRepoMock()
		tenorRepo := mck.NewTenorRepoMock()
		transactionRepo := mck.NewTransactionRepoMock()
		transactionService := service.NewTransactionService(validate, konsumerRepo, tenorRepo, transactionRepo)

		// mock
		konsumerRepo.Mock.On("GetByNik", mock.Anything, mock.Anything).
			Return(&entity.Konsumer{}, nil)

		tenorRepo.Mock.On("GetByNikAndDate", mock.Anything, mock.Anything, mock.Anything).
			Return(&entity.Tenor{
				Id:        1,
				Nik:       "1234567890123456",
				Bulan:     "Januari",
				StartDate: time.Now(),
				EndDate:   time.Now(),
				Tenor:     60000000,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}, nil)

		tenorRepo.Mock.On("UpdateLimit", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil)

		transactionRepo.Mock.On("Insert", mock.Anything, mock.Anything).
			Return(&entity.Transaction{}, nil)

		// test
		request := dto.BuyRequest{
			Nik:             "1234567890123456",
			DateTransaction: "2024-01-20 19:00:00",
			Otr:             1000000,
			AdminFee:        2000,
			JumlahCicilan:   4,
			Bunga:           2,
			Aset:            "sepeda",
		}

		result, err := transactionService.Buy(context.Background(), &request)
		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
}
