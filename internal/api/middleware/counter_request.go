package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/marcosvliras/sophie/internal/otel/metrics"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
)

func RequestCounterMiddleware() gin.HandlerFunc {
	counter, err := metrics.Meter.Int64Counter(
		"api_request_count",
		metric.WithDescription("Counts the number of HTTP requests"),
	)
	if err != nil {
		panic("failed to create request counter metric")
	}

	return func(c *gin.Context) {

		ctx := otel.GetTextMapPropagator().Extract(
			c.Request.Context(),
			propagation.HeaderCarrier(c.Request.Header),
		)

		c.Next()

		attr := []attribute.KeyValue{
			attribute.String("httpMethod", c.Request.Method),
			attribute.String("httpRoute", c.FullPath()),
			attribute.String("httpStatusCode", fmt.Sprintf("%d", c.Writer.Status())),
		}

		counter.Add(ctx, 1, metric.WithAttributes(attr...))

	}
}
