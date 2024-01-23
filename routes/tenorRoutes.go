package routes

import (
	"github.com/gofiber/fiber/v2"
	"xyzApp/app/handler"
)

func SetTenorRoutes(r fiber.Router, handler *handler.TenorHandler) {
	r.Post("/tenor", handler.InsertLimit)
}
