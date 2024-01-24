package test

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"xyzApp/app/customError"
	"xyzApp/app/handler"
	"xyzApp/app/model/dto"
	mockService "xyzApp/test/mock/service"
)

func TestTransactionHandlerBuy(t *testing.T) {
	t.Run("test buy error validasi", func(t *testing.T) {
		trService := mockService.NewTransactionServiceMock()
		trHandler := handler.NewTrasactionHandler(trService)
		app := fiber.New()
		app.Post("/", trHandler.Buy)

		// create request
		request := httptest.NewRequest(http.MethodPost, "/", nil)
		request.Header.Add("Content-Type", "application/json")

		// receive response
		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	})
	t.Run("test buy error not found", func(t *testing.T) {
		trService := mockService.NewTransactionServiceMock()
		trHandler := handler.NewTrasactionHandler(trService)
		app := fiber.New()
		app.Post("/", trHandler.Buy)

		// mock
		trService.Mock.On("Buy", mock.Anything, mock.Anything).
			Return(nil, customError.NewNotFoundError("record not found"))

		// test
		requestBody := dto.BuyRequest{
			Nik:             "1234567890123456",
			DateTransaction: "2020-10-10 10:00:00",
			Otr:             10000000,
			AdminFee:        2000,
			JumlahCicilan:   4,
			Bunga:           2.5,
			Aset:            "laptop",
		}
		reqJson, _ := json.Marshal(&requestBody)

		// create request
		request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(reqJson)))
		request.Header.Add("Content-Type", "application/json")

		// receive response
		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, http.StatusNotFound, response.StatusCode)
	})
	t.Run("test buy error bad request", func(t *testing.T) {
		trService := mockService.NewTransactionServiceMock()
		trHandler := handler.NewTrasactionHandler(trService)
		app := fiber.New()
		app.Post("/", trHandler.Buy)

		// mock
		trService.Mock.On("Buy", mock.Anything, mock.Anything).
			Return(nil, customError.NewBadRequestError("error"))

		// test
		requestBody := dto.BuyRequest{
			Nik:             "1234567890123456",
			DateTransaction: "2020-10-10 10:00:00",
			Otr:             1000000,
			AdminFee:        2500,
			JumlahCicilan:   3,
			Bunga:           10,
			Aset:            "laptop",
		}
		reqJson, _ := json.Marshal(&requestBody)

		// create request
		request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(reqJson)))
		request.Header.Add("Content-Type", "application/json")

		// receive response
		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	})
	t.Run("test buy error internal server", func(t *testing.T) {
		trService := mockService.NewTransactionServiceMock()
		trHandler := handler.NewTrasactionHandler(trService)
		app := fiber.New()
		app.Post("/", trHandler.Buy)

		// mock
		trService.Mock.On("Buy", mock.Anything, mock.Anything).
			Return(nil, customError.NewInternalSeverError("error"))

		// test
		requestBody := dto.BuyRequest{
			Nik:             "1234567890123456",
			DateTransaction: "2020-10-10 10:00:00",
			Otr:             1000000,
			AdminFee:        2000,
			JumlahCicilan:   2,
			Bunga:           10,
			Aset:            "laptop",
		}
		reqJson, _ := json.Marshal(&requestBody)

		// create request
		request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(reqJson)))
		request.Header.Add("Content-Type", "application/json")

		// receive response
		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, http.StatusInternalServerError, response.StatusCode)
	})
	t.Run("test buy success", func(t *testing.T) {
		trService := mockService.NewTransactionServiceMock()
		trHandler := handler.NewTrasactionHandler(trService)
		app := fiber.New()
		app.Post("/", trHandler.Buy)

		// mock
		trService.Mock.On("Buy", mock.Anything, mock.Anything).Return(&dto.BuyResponse{
			ReffNumber:    "1234567",
			Nik:           "1234567890123456",
			Otr:           10000000,
			AdminFee:      5000,
			JumlahCicilan: 4,
			Bunga:         2.5,
			Aset:          "laptop",
			TotalDebet:    500000,
			SisaLimit:     2000000,
		}, nil)

		// test
		requestBody := dto.BuyRequest{
			Nik:             "1234567890123456",
			DateTransaction: "2020-10-10 10:00:00",
			Otr:             10000000,
			AdminFee:        2000,
			JumlahCicilan:   4,
			Bunga:           2.5,
			Aset:            "laptop",
		}
		reqJson, _ := json.Marshal(&requestBody)

		// create request
		request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(reqJson)))
		request.Header.Add("Content-Type", "application/json")

		// receive response
		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, http.StatusOK, response.StatusCode)
	})
}
