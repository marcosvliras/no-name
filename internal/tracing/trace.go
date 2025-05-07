package tracing

import (
	"fmt"
	"log"
	"os"

	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

var (
	ServiceName  = os.Getenv("SERVICE_NAME")
	collectorURL = os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	Tracer       = otel.Tracer(ServiceName)
)

func newExporter(ctx context.Context) (*otlptrace.Exporter, error) {
	exporter, err := otlptracegrpc.New(
		ctx,
		otlptracegrpc.WithEndpoint(collectorURL),
		otlptracegrpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create OTLP trace exporter: %w", err)
	}
	return exporter, nil
}

func InitTracer() func() {
	ctx := context.Background()
	exporter, err := newExporter(ctx)
	if err != nil {
		log.Fatal(err)
	}

	resources, err := resource.New(
		ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String(ServiceName),
		),
	)
	if err != nil {
		log.Fatalf("failed to create resource: %v\n", err)
	}

	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resources),
	)
	otel.SetTracerProvider(traceProvider)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return func() {
		err := traceProvider.Shutdown(ctx)
		if err != nil {
			log.Printf("failed to shut down trace provider: %v", err)
		}
	}
}

//func newExporter() (*stdouttrace.Exporter, error) {
//	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
//	if err != nil {
//		return nil, fmt.Errorf("failed to initialize stdouttrace exporter %w", err)
//	}
//
//	return exporter, nil
//
//}
