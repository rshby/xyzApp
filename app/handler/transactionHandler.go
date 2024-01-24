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

type TransactionHandler struct {
	TransactionService service.ITransactionService
}

// function provider
func NewTrasactionHandler(transactionService service.ITransactionService) *TransactionHandler {
	return &TransactionHandler{TransactionService: transactionService}
}

// handler buy
func (t *TransactionHandler) Buy(ctx *fiber.Ctx) error {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx.Context(), "TransactionHandler Buy")
	defer span.Finish()

	// parsing request body
	var request dto.BuyRequest
	if err := ctx.BodyParser(&request); err != nil {
		statusCode := http.StatusBadRequest
		ctx.Status(statusCode)
		return ctx.JSON(&dto.ApiResponse{
			StatusCode: statusCode,
			Status:     helper.CodeToStatus(statusCode),
			Message:    err.Error(),
		})
	}

	// call procedure buy in service
	buyRespoonse, err := t.TransactionService.Buy(ctxTracing, &request)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errorMessages []string
			for _, fieldError := range validationErrors {
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

	// success
	statusCode := http.StatusOK
	ctx.Status(statusCode)
	return ctx.JSON(&dto.ApiResponse{
		StatusCode: statusCode,
		Status:     helper.CodeToStatus(statusCode),
		Message:    "success buy",
		Data:       buyRespoonse,
	})
}
