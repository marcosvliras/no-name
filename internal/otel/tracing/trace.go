package tracing

import (
	"fmt"
	"log"

	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/trace"

	"github.com/marcosvliras/sophie/internal/otel/config"
)

var (
	Tracer = otel.Tracer(config.ServiceName)
)

func traceExporter(ctx context.Context) (*otlptrace.Exporter, error) {
	exporter, err := otlptracegrpc.New(
		ctx,
		otlptracegrpc.WithGRPCConn(config.Conn),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create OTLP trace exporter: %w", err)
	}
	return exporter, nil
}

func InitTracer() func() {
	log.Println(config.ServiceName)
	log.Println(config.Resource)

	ctx := context.Background()
	exporter, err := traceExporter(ctx)
	if err != nil {
		log.Fatal(err)
	}

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

//func newExporter() (*stdouttrace.Exporter, error) {
//	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
//	if err != nil {
//		return nil, fmt.Errorf("failed to initialize stdouttrace exporter %w", err)
//	}
//	return exporter, nil
//}
