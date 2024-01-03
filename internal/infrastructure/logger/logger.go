package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Debug(args ...interface{})
	Fatal(args ...interface{})
	Println(args ...interface{})
	Printf(format string, args ...interface{})
}

func NewLogger() Logger {
	logger := logrus.New()

	logger.Out = os.Stdout
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Level = logrus.InfoLevel

	return logger
}
