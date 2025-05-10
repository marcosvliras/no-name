// reference to implementaion: https://github.com/open-telemetry/opentelemetry-go-contrib/blob/main/examples/otel-collector/main.go

package metrics

import (
	"context"
	"fmt"

	"github.com/marcosvliras/sophie/internal/otel/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

var Meter = otel.Meter(config.ServiceName)

func metricExporter(ctx context.Context) (*otlpmetricgrpc.Exporter, error) {
	exporter, err := otlpmetricgrpc.New(
		ctx,
		otlpmetricgrpc.WithGRPCConn(config.Conn),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create OTLP trace exporter: %w", err)
	}
	return exporter, nil
}

func InitMetrics() func() {
	ctx := context.Background()
	exporter, err := metricExporter(ctx)
	if err != nil {
		return nil
	}

	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(sdkmetric.NewPeriodicReader(exporter)),
		sdkmetric.WithResource(config.Resource),
	)
	otel.SetMeterProvider(meterProvider)

	return func() {
		err := meterProvider.Shutdown(ctx)
		if err != nil {
			fmt.Printf("failed to shut down meter provider: %v", err)
		}
	}

}
