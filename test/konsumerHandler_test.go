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

func TestRegisterKonsumerHandler(t *testing.T) {
	t.Run("test register konsumer error validasi", func(t *testing.T) {
		konsumerService := mockService.NewKonsumerServiceMock()
		konsumerHandler := handler.NewKonsumerHandler(konsumerService)

		app := fiber.New()
		app.Post("/", konsumerHandler.Register)

		// create request
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Add("Content-Type", "application/json")

		// hit and receive response
		response, err := app.Test(req)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	})
	t.Run("test register error not bad request", func(t *testing.T) {
		konsumerService := mockService.NewKonsumerServiceMock()
		konsumerHandler := handler.NewKonsumerHandler(konsumerService)
		app := fiber.New()
		app.Post("/", konsumerHandler.Register)

		// mock
		konsumerService.Mock.On("Register", mock.Anything, mock.Anything).
			Return(nil, customError.NewBadRequestError("error"))

		// test
		requestBody := &dto.RegisterKonsumerRequest{
			Nik:          "1234567890123456",
			FullName:     "hello world",
			LegalName:    "hello world",
			TempatLahir:  "Jakarta",
			TanggalLahir: "2020-10-10",
			Gaji:         20000000,
			FotoKtp:      "www.google.com",
			FotoSelfie:   "www.google.com",
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
	t.Run("test register error not found", func(t *testing.T) {
		konsumerService := mockService.NewKonsumerServiceMock()
		konsumerHandler := handler.NewKonsumerHandler(konsumerService)
		app := fiber.New()
		app.Post("/", konsumerHandler.Register)

		// mock
		konsumerService.Mock.On("Register", mock.Anything, mock.Anything).
			Return(nil, customError.NewNotFoundError("error data not found"))

		// test
		requestBody := dto.RegisterKonsumerRequest{
			Nik:          "1234567890123456",
			FullName:     "hello world",
			LegalName:    "hello world",
			TempatLahir:  "Jakarta",
			TanggalLahir: "1999-10-10",
			Gaji:         12000000,
			FotoKtp:      "www.google.com",
			FotoSelfie:   "www.google.com",
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
	t.Run("test reigster error internal server error", func(t *testing.T) {
		konsumerService := mockService.NewKonsumerServiceMock()
		konsumerHandler := handler.NewKonsumerHandler(konsumerService)
		app := fiber.New()
		app.Post("/", konsumerHandler.Register)

		// mock
		konsumerService.Mock.On("Register", mock.Anything, mock.Anything).
			Return(nil, customError.NewInternalSeverError("internal server error"))

		// test
		requestBody := dto.RegisterKonsumerRequest{
			Nik:          "1234567890123456",
			FullName:     "hello world",
			LegalName:    "hello world",
			TempatLahir:  "jakarta",
			TanggalLahir: "2010-10-10",
			Gaji:         12000000,
			FotoKtp:      "www.google.com",
			FotoSelfie:   "www.google.com",
		}
		reqJson, _ := json.Marshal(&requestBody)

		// create request
		request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(reqJson)))
		request.Header.Add("Content-Type", "application/json")

		// get response
		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, http.StatusInternalServerError, response.StatusCode)
	})
	t.Run("test register success", func(t *testing.T) {
		konsumerService := mockService.NewKonsumerServiceMock()
		konsumerHandler := handler.NewKonsumerHandler(konsumerService)
		app := fiber.New()
		app.Post("/", konsumerHandler.Register)

		// mock
		konsumerService.Mock.On("Register", mock.Anything, mock.Anything).
			Return(&dto.RegisterKonsumerResponse{
				Nik:          "1234567890123456",
				FullName:     "hello world",
				LegalName:    "hello world",
				TempatLahir:  "jakarta",
				TanggalLahir: "2010-10-10",
				Gaji:         12000000,
				FotoKtp:      "www.google.com",
				FotoSelfie:   "www.google.com",
				CreatedAt:    "2020-01-01 00:00:00",
				UpdatedAt:    "2020-01-01 00:00:00",
			}, nil)

		// test
		requestBody := dto.RegisterKonsumerRequest{
			Nik:          "1234567890123456",
			FullName:     "hello world",
			LegalName:    "hello world",
			TempatLahir:  "jakarta",
			TanggalLahir: "2020-01-01",
			Gaji:         12000000,
			FotoKtp:      "www.google.com",
			FotoSelfie:   "www.google.com",
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
