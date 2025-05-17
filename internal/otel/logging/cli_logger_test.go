package logging

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"testing"
)

func TestCliLoggerMethods(t *testing.T) {
	ctx := context.Background()

	var buf bytes.Buffer
	handlerOptions := slog.HandlerOptions{Level: slog.LevelDebug}
	logger := slog.New(slog.NewTextHandler(&buf, &handlerOptions))

	cliLogger := NewCliLogger()
	cliLogger.logger = logger

	t.Run("Debug", func(t *testing.T) {
		cliLogger.Debug(ctx, "debug message")
	})

	t.Run("Info", func(t *testing.T) {
		cliLogger.Info(ctx, "info message")
	})

	t.Run("Warning", func(t *testing.T) {
		cliLogger.Warning(ctx, "warning message")
	})

	t.Run("Error", func(t *testing.T) {
		cliLogger.Error(ctx, "error message")
	})

	t.Run("Fatal", func(t *testing.T) {
		cliLogger.Fatal(ctx, "fatal message")
	})

	t.Run("Close", func(t *testing.T) {
		if err := cliLogger.Close(ctx); err != nil {
			t.Errorf("Close returned error: %v", err)
		}
	})

	logLines := bytes.Count(buf.Bytes(), []byte("\n"))
	if logLines != 5 {
		t.Errorf("Expected 5 log entries, got %d", logLines)
	}
	fmt.Println(buf.String())
}
