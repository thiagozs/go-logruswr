package logruswr

import (
	"bytes"
	"context"
	"errors"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/thiagozs/go-logruswr/hooks"
)

func TestLogWrapper_WithField(t *testing.T) {
	logger, err := New()
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}

	entry := logger.WithField("key", "value")
	if entry.Data["key"] != "value" {
		t.Errorf("Expected field 'key' to be 'value', got '%v'", entry.Data["key"])
	}
}

func TestLogWrapper_WithFields(t *testing.T) {
	logger, err := New()
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}

	fields := Fields{"key1": "value1", "key2": "value2"}
	entry := logger.WithFields(fields)
	for k, v := range fields {
		if entry.Data[k] != v {
			t.Errorf("Expected field '%s' to be '%v', got '%v'", k, v, entry.Data[k])
		}
	}
}

func TestLogWrapper_WithError(t *testing.T) {
	logger, err := New()
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}

	testError := errors.New("test error")
	entry := logger.WithError(testError)
	if entry.Data["error"] != testError {
		t.Errorf("Expected error to be '%v', got '%v'", testError, entry.Data["error"])
	}
}

func TestLogWrapper_LogLevels(t *testing.T) {
	var buf bytes.Buffer

	logger, err := New(WithOutput(Stdout))
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}

	// Redirect log output to buffer
	logger.log.SetOutput(&buf)

	// Set log level to Info, should not log Debug
	logger.SetLevel(Info)
	logger.Debug("This is a debug message that should not appear")
	if buf.Len() > 0 {
		t.Error("Debug message was logged when log level was set to Info")
	}
	buf.Reset() // Clear buffer

	// Log at Info level, should appear
	logger.Info("Test1 message")
	if !strings.Contains(buf.String(), "Test1 message") {
		t.Error("The message are not the same on Info")
	}
	buf.Reset() // Clear buffer

	// Log at Error level, should appear
	logger.Error("Test2 message")
	if !strings.Contains(buf.String(), "Test2 message") {
		t.Error("The message are not the same on Info")
	}
	buf.Reset() // Clear buffer

	logger.Warn("Test3 message")
	if !strings.Contains(buf.String(), "Test3 message") {
		t.Error("The message are not the same on Warn")
	}
	buf.Reset() // Clear buffer

	logger.Trace("Test4 message")
	if buf.Len() > 0 {
		t.Error("Trace message was logged when log level was set to Info")
	}
	buf.Reset() // Clear buffer

}

func TestLogWrapper_LogLevelsPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	logger, err := New(WithOutput(Stdout))
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}

	// Set log level to Panic, should panic
	logger.SetLevel(Panic)
	logger.Panic("This is a panic message")
}

func TestLogWrapper_AddHooks(t *testing.T) {
	logger, err := New(WithOutput(Stdout))
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}

	hook := &hooks.TestHook{}

	logger.AddHook(hook)

	logger.Info("Test message")

	if !hook.IsFired() {
		t.Error("The hook was not called")
	}
}

func TestLogWrapper_WithContext(t *testing.T) {
	logger, err := New(WithOutput(Stdout))
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}

	ctx := context.Background()

	logger.WithContext(ctx).Info("Test message")
}

func TestLogWrapper_WithFormatter(t *testing.T) {
	logger, err := New(WithOutput(Stdout), WithFormatter(FormatterJSON))
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}

	logger.Info("Test message")
}

func TestLogWrapper_LogFilePath(t *testing.T) {
	defer func() {
		if err := os.Remove("/tmp/myapp.log"); err != nil {
			log.Fatalf("Failed to remove log file: %v", err)
		}
	}()
	logger, err := New(
		WithFormatter(FormatterJSON),
		WithOutput(File),
		WithLevel(Info),
		WithLogFilePath("/tmp/myapp.log"),
		WithMaxLogSize(10),     // Rotate after 10 MB
		WithMaxBackups(3),      // Keep only 3 log files
		WithMaxAge(28),         // 28 days
		WithCompressLogs(true), // Compress old log files
	)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	logger.Info("Test1 message")

	// open file and check if the message is there
	f, err := os.Open("/tmp/myapp.log")
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	buf := make([]byte, 1024)
	n, err := f.Read(buf)
	if err != nil {
		log.Fatalf("Failed to read log file: %v", err)
	}

	if !strings.Contains(string(buf[:n]), "Test1 message") {
		log.Fatalf("Log message not found in file")
	}

}

func TestLogWrapper_LogFilePathNoFile(t *testing.T) {
	logger, err := New(
		WithFormatter(FormatterJSON),
		WithOutput(File),
		WithLevel(Info),
		WithLogFilePath(""),
	)
	if err != nil {
		t.Fatalf("Failed to initialize logger: %v", err)
	}

	logger.Info("Test1 message")
}

func TestLogWrapper_Constants(t *testing.T) {
	if FormatterText.String() != "text" {
		t.Errorf("Expected FormatterText.String() to be 'text', got '%s'", FormatterText.String())
	}

	if FormatterJSON.String() != "json" {
		t.Errorf("Expected FormatterJSON.String() to be 'json', got '%s'", FormatterJSON.String())
	}

	if Panic.String() != "panic" {
		t.Errorf("Expected LevelPanic.String() to be 'panic', got '%s'", Panic.String())
	}

	if Fatal.String() != "fatal" {
		t.Errorf("Expected LevelFatal.String() to be 'fatal', got '%s'", Fatal.String())
	}

	if Error.String() != "error" {
		t.Errorf("Expected LevelError.String() to be 'error', got '%s'", Error.String())
	}

	if Warn.String() != "warn" {
		t.Errorf("Expected LevelWarn.String() to be 'warn', got '%s'", Warn.String())
	}

	if Info.String() != "info" {
		t.Errorf("Expected LevelInfo.String() to be 'info', got '%s'", Info.String())
	}

	if Debug.String() != "debug" {
		t.Errorf("Expected LevelDebug.String() to be 'debug', got '%s'", Debug.String())
	}

	if Trace.String() != "trace" {
		t.Errorf("Expected LevelTrace.String() to be 'trace', got '%s'", Trace.String())
	}
}

func TestLogWrapper_MarshalText(t *testing.T) {
	l := Info
	b, err := l.MarshalText()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	t.Logf("baaa: %s", b)

	if string(b) != Info.String() {
		t.Errorf("Expected 'info', got %s", string(b))
	}
}

func TestLogWrapper_UnmarshalText(t *testing.T) {
	var l Level
	err := l.UnmarshalText([]byte("info"))
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if l != Info {
		t.Errorf("Expected 'info', got %v", l)
	}
}
