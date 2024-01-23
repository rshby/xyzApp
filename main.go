package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/opentracing/opentracing-go"
	"xyzApp/app/config"
	"xyzApp/app/logger"
	"xyzApp/app/tracing"
	"xyzApp/database"
	server "xyzApp/server"
)

func main() {
	// set config App
	cfg := config.LoadConfig()
	fmt.Println(cfg)

	// set logger console
	log := logger.NewLoggerConsole()
	log.Info("start app")

	// set tracing
	tracer, closer := tracing.ConnectJaeger(cfg, log, "xyzApp")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	validate := validator.New()

	// connect db
	db := database.ConnectDB(cfg, log)

	// run server
	server := server.NewServerApp(cfg, validate, db)
	if err := server.RunServer(); err != nil {
		log.Fatalf("cant run server : %v", err)
	}
}
