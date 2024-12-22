package logging

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type LogConfig struct {
	Format string // "json" or "text"
	Level  string // "debug", "info", "warn", "error", "fatal", "panic"
}

var Logger *logrus.Logger

func InitLogger(config LogConfig) {
	Logger = logrus.New()
	Logger.SetFormatter(&logrus.JSONFormatter{})

	if config.Format == "json" {
		Logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.RFC3339,
		})
	} else {
		Logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: time.RFC3339,
		})
	}

	level, err := logrus.ParseLevel(config.Level)
	if err != nil {
		Logger.Warn("Invalid log level, defaulting to 'info'")
		level = logrus.InfoLevel
	}
	Logger.SetLevel(level)

	Logger.SetOutput(os.Stdout)

	Logger.Info("Logger initialized")
}
