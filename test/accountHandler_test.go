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

func TestAccountHandlerRegister(t *testing.T) {
	t.Run("test register error validasi", func(t *testing.T) {
		accountService := mockService.NewAccountServiceMock()
		accountHandler := handler.NewAccountHandler(accountService)
		app := fiber.New()
		app.Post("/", accountHandler.Register)

		// create request
		request := httptest.NewRequest(http.MethodPost, "/", nil)
		request.Header.Add("Content-Type", "application/json")

		// receive response
		response, err := app.Test(request)
		assert.Nil(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	})
	t.Run("test register error not found", func(t *testing.T) {
		accountService := mockService.NewAccountServiceMock()
		accountHandler := handler.NewAccountHandler(accountService)
		app := fiber.New()
		app.Post("/", accountHandler.Register)

		// mock
		accountService.Mock.On("Register", mock.Anything, mock.Anything).
			Return(customError.NewNotFoundError("error"))

		// test
		requestBody := dto.RegisterAccount{
			Nik:      "1234567890123456",
			Email:    "hello@gmail.com",
			Password: "123456",
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
	t.Run("test register error bad request", func(t *testing.T) {
		accountService := mockService.NewAccountServiceMock()
		accountHandler := handler.NewAccountHandler(accountService)
		app := fiber.New()
		app.Post("/", accountHandler.Register)

		// mock
		accountService.Mock.On("Register", mock.Anything, mock.Anything).
			Return(customError.NewBadRequestError("error"))

		// test
		requestBody := dto.RegisterAccount{
			Nik:      "1234567890123456",
			Email:    "hello@gmail.com",
			Password: "123456",
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
	t.Run("test register error internal server", func(t *testing.T) {
		accountService := mockService.NewAccountServiceMock()
		accountHandler := handler.NewAccountHandler(accountService)
		app := fiber.New()
		app.Post("/", accountHandler.Register)

		// mock
		accountService.Mock.On("Register", mock.Anything, mock.Anything).
			Return(customError.NewInternalSeverError("error"))

		// request
		requestBody := dto.RegisterAccount{
			Nik:      "1234567890123456",
			Email:    "hello@gmail.com",
			Password: "123456",
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
	t.Run("test success register", func(t *testing.T) {
		accountService := mockService.NewAccountServiceMock()
		accountHandler := handler.NewAccountHandler(accountService)
		app := fiber.New()
		app.Post("/", accountHandler.Register)

		// mock
		accountService.Mock.On("Register", mock.Anything, mock.Anything).
			Return(nil)

		// test
		requestBody := dto.RegisterAccount{
			Nik:      "1234567890",
			Email:    "hello@gmail.com",
			Password: "123456",
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
