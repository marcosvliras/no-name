package tracing

import (
	"context"
	"testing"

	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
)

func TestInitTracer(t *testing.T) {
	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		t.Fatalf("failed to initialize stdouttrace exporter: %v", err)
	}

	ctx := context.Background()
	shutdown := InitTracer(ctx, exporter)
	defer shutdown()

	// Check if the tracer is initialized
	if Tracer == nil {
		t.Error("Tracer is not initialized")
	}

}
