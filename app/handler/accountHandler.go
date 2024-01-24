package handler

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/opentracing/opentracing-go"
	"net/http"
	"strings"
	"xyzApp/app/customError"
	"xyzApp/app/helper"
	"xyzApp/app/model/dto"
	"xyzApp/app/service"
)

type AccountHandler struct {
	AccountService service.IAccountService
}

// function Provider
func NewAccountHandler(accService service.IAccountService) *AccountHandler {
	return &AccountHandler{AccountService: accService}
}

// handler register
func (a *AccountHandler) Register(ctx *fiber.Ctx) error {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx.Context(), "AccountHandler Register")
	defer span.Finish()

	// parsing request body
	var request dto.RegisterAccount
	if err := ctx.BodyParser(&request); err != nil {
		statusCode := http.StatusBadRequest
		ctx.Status(statusCode)
		return ctx.JSON(&dto.ApiResponse{
			StatusCode: statusCode,
			Status:     helper.CodeToStatus(statusCode),
			Message:    err.Error(),
		})
	}

	// call procedure in service
	if err := a.AccountService.Register(ctxTracing, &request); err != nil {
		if validationErros, ok := err.(validator.ValidationErrors); ok {
			var errorMessages []string
			for _, fieldError := range validationErros {
				msg := fmt.Sprintf("error on fiel [%v] with tag [%v]", fieldError.Field(), fieldError.Tag())
				errorMessages = append(errorMessages, msg)
			}

			statusCode := http.StatusBadRequest
			ctx.Status(statusCode)
			return ctx.JSON(&dto.ApiResponse{
				StatusCode: statusCode,
				Status:     helper.CodeToStatus(statusCode),
				Message:    strings.Join(errorMessages, ". "),
			})
		}

		var statusCode int
		switch err.(type) {
		case *customError.NotFoundError:
			statusCode = http.StatusNotFound
		case *customError.BadRequestError:
			statusCode = http.StatusBadRequest
		default:
			statusCode = http.StatusInternalServerError
		}

		ctx.Status(statusCode)
		return ctx.JSON(&dto.ApiResponse{
			StatusCode: statusCode,
			Status:     helper.CodeToStatus(statusCode),
			Message:    err.Error(),
		})
	}

	// success register
	statusCode := http.StatusOK
	ctx.Status(statusCode)
	return ctx.JSON(&dto.ApiResponse{
		StatusCode: statusCode,
		Status:     helper.CodeToStatus(statusCode),
		Message:    "success register account",
	})
}

// handler login
func (a *AccountHandler) Login(ctx *fiber.Ctx) error {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx.Context(), "AccountHandler Login")
	defer span.Finish()

	// parsing request_body
	var request dto.LoginRequest
	if err := ctx.BodyParser(&request); err != nil {
		statusCode := http.StatusBadRequest
		ctx.Status(statusCode)
		return ctx.JSON(&dto.ApiResponse{
			StatusCode: statusCode,
			Status:     helper.CodeToStatus(statusCode),
			Message:    err.Error(),
		})
	}

	// call procedure in service
	login, err := a.AccountService.Login(ctxTracing, &request)
	if err != nil {
		// jika error validasi
		if validationError, ok := err.(validator.ValidationErrors); ok {
			var errorMessages []string
			for _, fieldError := range validationError {
				msg := fmt.Sprintf("error on field [%v] with tag [%v]", fieldError.Field(), fieldError.Tag())
				errorMessages = append(errorMessages, msg)
			}

			statusCode := http.StatusBadRequest
			ctx.Status(statusCode)
			return ctx.JSON(&dto.ApiResponse{
				StatusCode: statusCode,
				Status:     helper.CodeToStatus(statusCode),
				Message:    strings.Join(errorMessages, ". "),
			})
		}

		var statusCode int
		switch err.(type) {
		case *customError.NotFoundError:
			statusCode = http.StatusNotFound
		case *customError.BadRequestError:
			statusCode = http.StatusBadRequest
		default:
			statusCode = http.StatusInternalServerError
		}

		ctx.Status(statusCode)
		return ctx.JSON(&dto.ApiResponse{
			StatusCode: statusCode,
			Status:     helper.CodeToStatus(statusCode),
			Message:    err.Error(),
		})
	}

	// success login
	statusCode := http.StatusOK
	ctx.Status(statusCode)
	return ctx.JSON(&dto.ApiResponse{
		StatusCode: statusCode,
		Status:     helper.CodeToStatus(statusCode),
		Message:    "success login",
		Data:       login,
	})

}
