package server

import (
	"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"xyzApp/app/config"
	"xyzApp/app/handler"
	"xyzApp/app/repository"
	"xyzApp/app/service"
	"xyzApp/routes"
)

type ServerApp struct {
	Router *fiber.App
	Config *config.AppConfig
}

func NewServerApp(cfg *config.AppConfig, validate *validator.Validate, db *sql.DB) IServer {
	// register repository
	konsumerRepo := repository.NewKonsumerRepository(db)
	tenorRepo := repository.NewTenorRepository(db)

	// register service
	konsumerService := service.NewKonsumerService(validate, konsumerRepo)
	tenorService := service.NewTenorService(validate, tenorRepo, konsumerRepo)

	// register handler
	konsumerHandler := handler.NewKonsumerHandler(konsumerService)
	tenorHandler := handler.NewTenorHandler(tenorService)

	app := fiber.New(fiber.Config{
		Prefork: false,
	})

	v1 := app.Group("/v1")
	v1.Use(logger.New())

	routes.SetKonsumerRoutes(v1, konsumerHandler)
	routes.SetTenorRoutes(v1, tenorHandler)
	return &ServerApp{
		Router: app,
		Config: cfg,
	}
}

// method implementasi run server
func (s *ServerApp) RunServer() error {
	addr := fmt.Sprintf(":%v", s.Config.App.Port)
	return s.Router.Listen(addr)
}
