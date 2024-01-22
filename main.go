package main

import (
	"fmt"
	"xyzApp/app/config"
	"xyzApp/app/logger"
)

func main() {
	// set config App
	cfg := config.LoadConfig()
	fmt.Println(cfg)

	// set logger
	log := logger.NewLoggerConsole()
	log.Info("start app")

	// set tracing
}
