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

func TestTenorHandlerInsertLimit(t *testing.T) {
	t.Run("test insert limit error validasi", func(t *testing.T) {
		tenorService := mockService.NewTenorServiceMock()
		tenorHandler := handler.NewTenorHandler(tenorService)
		app := fiber.New()
		app.Post("/", tenorHandler.InsertLimit)

		// create request
		request := httptest.NewRequest(http.MethodPost, "/", nil)
		request.Header.Add("Content-Type", "application/json")

		// receive response
		response, err := app.Test(request)
		assert.NotNil(t, response)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	})
	t.Run("test insert limit error not found", func(t *testing.T) {
		tenorService := mockService.NewTenorServiceMock()
		tenorHandler := handler.NewTenorHandler(tenorService)
		app := fiber.New()
		app.Post("/", tenorHandler.InsertLimit)

		// mock
		tenorService.Mock.On("InsertLimit", mock.Anything, mock.Anything).Return(nil, customError.NewNotFoundError("record not found"))

		// test
		requestBody := dto.InsertLimitRequest{
			Nik:   "1234567890123456",
			Bulan: 2,
			Tahun: 2023,
			Limit: 1000000,
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
	t.Run("test insert limit error bad request", func(t *testing.T) {
		tenorService := mockService.NewTenorServiceMock()
		tenorHandler := handler.NewTenorHandler(tenorService)
		app := fiber.New()
		app.Post("/", tenorHandler.InsertLimit)

		// mock
		tenorService.Mock.On("InsertLimit", mock.Anything, mock.Anything).
			Return(nil, customError.NewBadRequestError("error bad request"))

		// test
		requestBody := dto.InsertLimitRequest{
			Nik:   "1234567890123456",
			Bulan: 2,
			Tahun: 2020,
			Limit: 1000000,
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
	t.Run("test insert limit error internal server", func(t *testing.T) {
		tenorService := mockService.NewTenorServiceMock()
		tenorHandler := handler.NewTenorHandler(tenorService)
		app := fiber.New()
		app.Post("/", tenorHandler.InsertLimit)

		// mock
		tenorService.Mock.On("InsertLimit", mock.Anything, mock.Anything).
			Return(nil, customError.NewInternalSeverError("error internal server"))

		// test
		requestBody := dto.InsertLimitRequest{
			Nik:   "1234567890123456",
			Bulan: 2,
			Tahun: 2020,
			Limit: 2000000,
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
	t.Run("test insert lmit success", func(t *testing.T) {
		tenorService := mockService.NewTenorServiceMock()
		tenorHandler := handler.NewTenorHandler(tenorService)
		app := fiber.New()
		app.Post("/", tenorHandler.InsertLimit)

		// mock
		tenorService.Mock.On("InsertLimit", mock.Anything, mock.Anything).
			Return(&dto.InsertLimitResponse{
				Nik:       "1234567890123456",
				Bulan:     "Januari",
				StartDate: "2020-01-01",
				EndDate:   "2020-01-31",
				Limit:     12000000,
			}, nil)

		// test
		requestBody := dto.InsertLimitRequest{
			Nik:   "1234567890123456",
			Bulan: 1,
			Tahun: 2020,
			Limit: 12000000,
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
