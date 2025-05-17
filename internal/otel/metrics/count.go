// reference to implementaion: https://github.com/open-telemetry/opentelemetry-go-contrib/blob/main/examples/otel-collector/main.go

package metrics

import (
	"context"
	"fmt"

	"github.com/marcosvliras/sophie/internal/otel/config"
	"go.opentelemetry.io/otel"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

var Meter = otel.Meter(config.GetServiceName())

func InitMetrics(ctx context.Context, exporter sdkmetric.Exporter) func() {

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
