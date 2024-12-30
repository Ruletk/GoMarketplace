package logging

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestInitLogger(t *testing.T) {
	// Buffer for checking output
	var buf bytes.Buffer

	// Logger configuration
	config := LogConfig{
		Format: "json",
		Level:  "info",
	}

	// Initialize logger
	logger := InitLogger(config)

	// Set buffer as output
	logger.SetOutput(&buf)

	// Log a message
	logger.Info("Test message")

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
	// Buffer for checking output
	var buf bytes.Buffer

	// Configuration with invalid log level
	config := LogConfig{
		Format: "json",
		Level:  "invalid",
	}

	// Initialize logger
	logger := InitLogger(config)
	logger.SetOutput(&buf)

	// Check if log output is not empty
	if logger.GetLevel() != logrus.InfoLevel {
		t.Errorf("Expected default level 'info', got '%v'", logger.GetLevel())
	}
}
