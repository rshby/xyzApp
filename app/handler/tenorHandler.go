package handler

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/opentracing/opentracing-go"
	"net/http"
	"strings"
	"xyzApp/app/helper"
	"xyzApp/app/model/dto"
	"xyzApp/app/service"
)

type TenorHandler struct {
	TenorService service.ITenorService
}

// function provider
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
	}

}
