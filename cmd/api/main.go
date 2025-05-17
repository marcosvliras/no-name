package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/marcosvliras/sophie/internal/api/controllers"
	"github.com/marcosvliras/sophie/internal/api/middleware"
	"github.com/marcosvliras/sophie/internal/otel/logging"
	"github.com/marcosvliras/sophie/internal/otel/metrics"
	"github.com/marcosvliras/sophie/internal/otel/tracing"
	"github.com/marcosvliras/sophie/internal/service"

	"github.com/marcosvliras/sophie/internal/otel/config"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/metric"
)

func main() {
	config.InitConfig()

	cleanupTrace := tracing.InitTracer()
	defer cleanupTrace()

	cleanupMeter := metrics.InitMetrics()
	defer cleanupMeter()

	exporter, err := otlploggrpc.New(
		context.TODO(),
		otlploggrpc.WithGRPCConn(config.Conn),
		otlploggrpc.WithInsecure(),
	)
	if err != nil {
		panic(err)
	}
	err = logging.InitLogger("api", exporter)
	if err != nil {
		panic(err)
	}
	defer logging.SLogger.Close(context.Background())

	counter, err := metrics.Meter.Int64Counter(
		"api_request_count",
		metric.WithDescription("Counts the number of HTTP requests"),
	)
	if err != nil {
		panic("failed to create request counter metric")
	}

	server := gin.Default()
	server.Use(otelgin.Middleware(config.ServiceName))
	server.Use(middleware.LoggerMiddleware())
	server.Use(middleware.RequestCounterMiddleware(counter))

	svc := service.NewAlphavantageSVC()
	stockCtrl := controllers.NewStocksCtrl(svc)

	server.GET("/health", controllers.HealthCheck)
	server.GET("/stocks", stockCtrl.Handle)

	if err := server.Run(":8000"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
