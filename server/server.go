package server

import (
	"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"xyzApp/app/config"
	"xyzApp/app/handler"
	"xyzApp/app/middleware"
	"xyzApp/app/repository"
	"xyzApp/app/service"
	"xyzApp/routes"
)

type ServerApp struct {
	Router *fiber.App
	Config config.IConfig
}

func NewServerApp(cfg config.IConfig, validate *validator.Validate, db *sql.DB) IServer {
	// register repository
	konsumerRepo := repository.NewKonsumerRepository(db)
	tenorRepo := repository.NewTenorRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)
	accountRepo := repository.NewAccountRepository(db)

	// register service
	konsumerService := service.NewKonsumerService(validate, konsumerRepo)
	tenorService := service.NewTenorService(validate, tenorRepo, konsumerRepo)
	transactionService := service.NewTransactionService(validate, konsumerRepo, tenorRepo, transactionRepo)
	accountService := service.NewAccountService(cfg, validate, accountRepo, konsumerRepo)

	// register handler
	konsumerHandler := handler.NewKonsumerHandler(konsumerService)
	tenorHandler := handler.NewTenorHandler(tenorService)
	transactionHandler := handler.NewTrasactionHandler(transactionService)
	accountHandler := handler.NewAccountHandler(accountService)

	// middleware
	authMiddleware := middleware.AuthMiddleware(cfg.GetConfig())
	loggerMiddleware := middleware.LoggerMiddleware(cfg)

	app := fiber.New(fiber.Config{
		Prefork: false,
	})

	v1 := app.Group("/v1").Use(loggerMiddleware)
	v1.Use(logger.New())

	routes.SetKonsumerRoutes(v1, konsumerHandler)
	routes.SetTenorRoutes(v1, authMiddleware, tenorHandler)
	routes.SetTransactionRoutes(v1, authMiddleware, transactionHandler)
	routes.SetAccountRoutes(v1, accountHandler)

	return &ServerApp{
		Router: app,
		Config: cfg,
	}
}

// method implementasi run server
func (s *ServerApp) RunServer() error {
	addr := fmt.Sprintf(":%v", s.Config.GetConfig().App.Port)
	return s.Router.Listen(addr)
}
