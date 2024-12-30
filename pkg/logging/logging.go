package logging

import (
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type LogConfig struct {
	Format string // "json" or "text"
	Level  string // "debug", "info", "warn", "error", "fatal", "panic"
}

func InitLogger(config LogConfig) *logrus.Logger {
	logger := logrus.New()

	if config.Format == "json" {
		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02T15:04:05.000Z07:00",
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "timestamp",
				logrus.FieldKeyLevel: "level",
				logrus.FieldKeyMsg:   "message",
				logrus.FieldKeyFunc:  "caller",
			},
		})
	} else {
		logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: time.RFC3339,
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "timestamp",
				logrus.FieldKeyLevel: "level",
				logrus.FieldKeyFunc:  "caller",
				logrus.FieldKeyMsg:   "message",
			},
		})
	}

	level, err := logrus.ParseLevel(config.Level)
	if err != nil {
		logger.Warn("Invalid log level, defaulting to 'info'")
		level = logrus.InfoLevel
	}

	_, file, _, _ := runtime.Caller(0) // Текущий файл
	projectRoot := filepath.Dir(filepath.Dir(file))

	logger.SetLevel(level)
	logger.SetOutput(os.Stdout)
	logger.AddHook(&CallerHook{
		ProjectRoot: projectRoot,
	})

	logger.Info("Logger initialized")
	return logger
}

type CallerHook struct {
	ProjectRoot string
}

func (hook *CallerHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook *CallerHook) Fire(entry *logrus.Entry) error {
	if _, file, line, ok := runtime.Caller(8); ok {
		relativeFile, err := filepath.Rel(hook.ProjectRoot, file)
		if err != nil {
			relativeFile = file
		}
		entry.Data["caller"] = relativeFile
		entry.Data["line"] = line
	}
	return nil
}
