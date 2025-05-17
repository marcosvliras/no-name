package middleware

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
)

type Counter interface {
	Add(ctx context.Context, value int64, options ...metric.AddOption)
}

func RequestCounterMiddleware(counter Counter) gin.HandlerFunc {

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
