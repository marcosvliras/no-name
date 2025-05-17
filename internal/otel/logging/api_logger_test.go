package logging

import (
	"bytes"
	"strings"
	"testing"

	"go.opentelemetry.io/otel/exporters/stdout/stdoutlog"
)

func TestSophieLogger(t *testing.T) {

	var buffer bytes.Buffer
	exporter, err := stdoutlog.New(
		stdoutlog.WithWriter(&buffer),
		stdoutlog.WithPrettyPrint(),
	)

	if err != nil {
		t.Fatalf("failed to create stdout exporter: %v", err)
	}

	logger := newSophieLogger(exporter)

	if logger == nil {
		t.Fatalf("expected logger to be non-nil")
	}
	defer logger.Close(t.Context())

	logger.Debug(t.Context(), "test debug message")
	logger.Info(t.Context(), "test info message")
	logger.Warning(t.Context(), "test warning message")
	logger.Error(t.Context(), "test error message")
	logger.Fatal(t.Context(), "test fatal message")

	// close the logger to flush any buffered logs
	logger.Close(t.Context())

	// divide strings in Scope name inside the buffer
	amount := strings.Count(buffer.String(), "Scope")
	if amount != 5 {
		t.Fatalf("expected 1 log message, got %d", amount)
	}

}
