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

type TenorHandler struct {
	TenorService service.ITenorService
}

// function provider
func NewTenorHandler(tenorService service.ITenorService) *TenorHandler {
	return &TenorHandler{TenorService: tenorService}
}

// handler limit
func (t *TenorHandler) InsertLimit(ctx *fiber.Ctx) error {
	span, ctxTracing := opentracing.StartSpanFromContext(ctx.Context(), "TenorHandler InsertLimit")
	defer span.Finish()

	// parsing request body
	var request dto.InsertLimitRequest
	if err := ctx.BodyParser(&request); err != nil {
		statusCode := http.StatusBadRequest
		ctx.Status(statusCode)
		return ctx.JSON(&dto.ApiResponse{
			StatusCode: statusCode,
			Status:     helper.CodeToStatus(statusCode),
			Message:    err.Error(),
		})
	}

	// call proceduer service
	tenor, err := t.TenorService.InsertLimit(ctxTracing, &request)
	if err != nil {
		// error bad request
		if errorValidations, ok := err.(validator.ValidationErrors); ok {
			var errorMessages []string
			for _, errorField := range errorValidations {
				msg := fmt.Sprintf("error in field [%v] with tag [%v]",
					errorField.Field(), errorField.Tag())
				errorMessages = append(errorMessages, msg)
			}

			statusCode := http.StatusBadRequest
			ctx.Status(statusCode)
			return ctx.JSON(&dto.ApiResponse{
				StatusCode: statusCode,
				Status:     helper.CodeToStatus(statusCode),
				Message:    strings.Join(errorMessages, ", "),
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
		Message:    "success insert tenor",
		Data:       tenor,
	})
}
