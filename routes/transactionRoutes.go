package routes

import (
	"github.com/gofiber/fiber/v2"
	"xyzApp/app/handler"
)

func SetTransactionRoutes(r fiber.Router, handler *handler.TransactionHandler) {
	r.Post("/buy", handler.Buy)
}
