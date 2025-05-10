package config

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

var Resource *resource.Resource

func newResource() error {

	var err error

	Resource, err = resource.New(
		context.Background(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String(ServiceName),
			semconv.ServiceVersionKey.String("1.0.0"),
			semconv.TelemetrySDKLanguageGo,
		),
	)
	if err != nil {
		return fmt.Errorf("failed to create resource: %w", err)
	}

	return nil
}
