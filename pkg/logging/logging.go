package logging

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"math"
	"net/http"
	"os"
	"time"
)

type LogConfig struct {
	Level        string // "debug", "info", "warn", "error", "fatal", "panic"
	EnableCaller bool
	LoggerName   string
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

	Logger.AddHook(&LoggerNameHook{LoggerName: config.LoggerName})
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

func GinLogger(logger logrus.FieldLogger) gin.HandlerFunc {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	return func(c *gin.Context) {
		path := c.Request.URL.Path
		start := time.Now()
		c.Next()
		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		clientUserAgent := c.Request.UserAgent()
		referer := c.Request.Referer()
		dataLength := c.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}

		entry := logger.WithFields(logrus.Fields{
			"hostname":   hostname,
			"statusCode": statusCode,
			"latency":    latency, // time to process
			"clientIP":   clientIP,
			"method":     c.Request.Method,
			"path":       path,
			"referer":    referer,
			"dataLength": dataLength,
			"userAgent":  clientUserAgent,
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			if statusCode >= http.StatusInternalServerError {
				entry.Error("Internal server error")
			} else if statusCode >= http.StatusBadRequest {
				entry.Warn("Client error")
			} else {
				entry.Info("Request processed successfully")
			}
		}
	}
}
