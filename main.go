package main

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"xyzApp/app/config"
	"xyzApp/app/logger"
	"xyzApp/app/tracing"
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

	// connect db
}
