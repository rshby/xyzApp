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

type KonsumerHandler struct {
	KonsumerService service.IKonsumerService
}

func NewKonsumerHandler(konsumerService service.IKonsumerService) *KonsumerHandler {
	return &KonsumerHandler{KonsumerService: konsumerService}
}

// handler register
func (k *KonsumerHandler) Register(ctx *fiber.Ctx) error {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx.Context(), "KonsumerHandler Register")
	defer span.Finish()

	// parsing body
	var request dto.RegisterKonsumerRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		statusCode := http.StatusBadRequest
		ctx.Status(statusCode)
		return ctx.JSON(&dto.ApiResponse{
			StatusCode: statusCode,
			Status:     helper.CodeToStatus(statusCode),
			Message:    err.Error(),
		})
	}

	// call service
	response, err := k.KonsumerService.Register(ctxTracing, &request)
	if err != nil {
		// if error validation
		if validationError, ok := err.(validator.ValidationErrors); ok {
			var errorMessages []string
			for _, fieldError := range validationError {
				message := fmt.Sprintf("error on field : [%v], with error : [%v]", fieldError.Field(), helper.GetErrorByTag(fieldError.Tag()))
				errorMessages = append(errorMessages, message)
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
			// internal server error
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
		Message:    "success register konsumer",
		Data:       response,
	})
}
