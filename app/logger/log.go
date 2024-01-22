package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

// function log to console
func NewLoggerConsole() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.DebugLevel)

	return log
}

// function log to file
func NewLoggerFile() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.DebugLevel)
	log.SetOutput(os.Stdout)

	// open file
	file, err := os.OpenFile("./log/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Failed to log to file, using default stdout")
	}

	return log
}
