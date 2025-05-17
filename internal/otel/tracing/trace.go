package tracing

import (
	"log"

	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/trace"

	"github.com/marcosvliras/sophie/internal/otel/config"
)

var (
	Tracer = otel.Tracer(config.GetServiceName())
)

func InitTracer(ctx context.Context, exporter trace.SpanExporter) func() {

	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(config.Resource),
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
