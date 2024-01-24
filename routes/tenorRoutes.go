package routes

import (
	"github.com/gofiber/fiber/v2"
	"xyzApp/app/handler"
)

func SetTenorRoutes(r fiber.Router, middleware fiber.Handler, handler *handler.TenorHandler) {
	r.Post("/tenor", middleware, handler.InsertLimit)
}
