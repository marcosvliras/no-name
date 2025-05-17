package logging

import (
	"bytes"
	"testing"

	"go.opentelemetry.io/otel/exporters/stdout/stdoutlog"
)

func TestInitLogger(t *testing.T) {
	t.Run("cli logger", func(t *testing.T) {
		err := InitLogger("cli", nil)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if SLogger == nil {
			t.Errorf("expected SLogger to be non-nil")
		}
	})

	t.Run("api logger", func(t *testing.T) {

		var buffer bytes.Buffer

		exporter, err := stdoutlog.New(
			stdoutlog.WithWriter(&buffer),
		)
		if err != nil {
			t.Fatalf("failed to create stdout exporter: %v", err)
		}

		err = InitLogger("api", exporter)

		if err != nil {
			t.Errorf("expected error, got nil")
		}
		if SLogger == nil {
			t.Errorf("expected SLogger to be nil")
		}
	})

	t.Run("invalid log type", func(t *testing.T) {
		SLogger = nil

		err := InitLogger("invalid", nil)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
		if SLogger != nil {
			t.Errorf("expected SLogger to be nil")
		}
	})

}
