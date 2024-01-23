package routes

import (
	"github.com/gofiber/fiber/v2"
	"xyzApp/app/handler"
)

func SetKonsumerRoutes(r fiber.Router, handler *handler.KonsumerHandler) {
	r.Post("/register", handler.Register)
}
