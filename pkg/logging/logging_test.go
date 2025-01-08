package logging

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/sirupsen/logrus"
)

func setupLoggerTest(level string) *bytes.Buffer {
	var buf bytes.Buffer

	config := LogConfig{
		Level:        level,
		EnableCaller: true,
	}

	BaseInitLogger(config)
	Logger.SetOutput(&buf)

	return &buf
}

func TestInitLogger(t *testing.T) {
	buf := setupLoggerTest("info")

	// Log a message
	Logger.Info("Test message")

	// Check if log output is not empty
	output := buf.String()
	if output == "" {
		t.Fatalf("Expected log output, got empty string")
	}

	// Unmarshal log output
	var logEntry map[string]interface{}
	if err := json.Unmarshal([]byte(output), &logEntry); err != nil {
		t.Fatalf("Failed to unmarshal log output: %v", err)
	}

	// Check log output fields
	if logEntry["message"] != "Test message" {
		t.Errorf("Expected message 'Test message', got '%v'", logEntry["message"])
	}

	if logEntry["level"] != "info" {
		t.Errorf("Expected level 'info', got '%v'", logEntry["level"])
	}

	// Check if 'timestamp' field is present
	if _, ok := logEntry["caller"]; !ok {
		t.Errorf("Expected 'caller' field in log output")
	}
}

func TestInitLoggerWithInvalidLevel(t *testing.T) {
	_ = setupLoggerTest("invalid")

	// Check if log output is not empty
	if Logger.GetLevel() != logrus.InfoLevel {
		t.Errorf("Expected default level 'info', got '%v'", Logger.GetLevel())
	}
}

func TestInitiatorFieldExists(t *testing.T) {
	buf := setupLoggerTest("info")

	// Log a message
	Logger.Info("Test message")

	// Check if log output is not empty
	output := buf.String()
	if output == "" {
		t.Fatalf("Expected log output, got empty string")
	}

	// Unmarshal log output
	var logEntry map[string]interface{}
	if err := json.Unmarshal([]byte(output), &logEntry); err != nil {
		t.Fatalf("Failed to unmarshal log output: %v", err)
	}

	// Check if 'caller' field is present
	if _, ok := logEntry["caller"]; !ok {
		t.Errorf("Expected 'caller' field in log output")
	}
}
