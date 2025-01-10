package logging

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

type LogConfig struct {
	Level        string // "debug", "info", "warn", "error", "fatal", "panic"
	EnableCaller bool
}

var Logger *logrus.Logger

func BaseInitLogger(config LogConfig) {
	Logger = logrus.New()

	Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000Z07:00",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyFunc:  "caller",
		},
	})
	if config.EnableCaller {
		Logger.SetReportCaller(true)
	}

	level, err := logrus.ParseLevel(config.Level)
	if err != nil {
		Logger.Warn("Invalid log level, defaulting to 'info'")
		level = logrus.InfoLevel
	}

	Logger.SetLevel(level)
}

func InitLogger(config LogConfig) {
	BaseInitLogger(config)
	file, err := os.OpenFile("./logs/log.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		Logger.SetOutput(io.MultiWriter(os.Stdout, file))
	} else {
		Logger.Info("Failed to log to file, using default stderr")
	}

	Logger.Info("Logger initialized")
}
