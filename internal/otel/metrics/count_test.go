package metrics

import (
	"context"
	"testing"

	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
)

func TestInitMetrics(t *testing.T) {

	exporter, err := stdoutmetric.New(
		stdoutmetric.WithPrettyPrint(),
	)
	if err != nil {
		t.Fatalf("failed to create stdout exporter: %v", err)
	}

	// Initialize the metrics
	cleanup := InitMetrics(context.Background(), exporter)
	defer cleanup()

	// Check if the Meter is initialized
	if Meter == nil {
		t.Error("Meter should not be nil after initialization")
	}

}
