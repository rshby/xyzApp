package routes

import (
	"github.com/gofiber/fiber/v2"
	"xyzApp/app/handler"
)

func SetTransactionRoutes(r fiber.Router, middleware fiber.Handler, handler *handler.TransactionHandler) {
	r.Post("/buy", middleware, handler.Buy)
}
