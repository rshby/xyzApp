package routes

import (
	"github.com/gofiber/fiber/v2"
	"xyzApp/app/handler"
)

func SetAccountRoutes(r fiber.Router, handler *handler.AccountHandler) {
	r.Post("/account", handler.Register)
	r.Post("/login", handler.Login)
}
