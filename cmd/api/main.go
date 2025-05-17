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
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/metric"
)

func main() {
	config.InitConfig()

	// Initialize the OpenTelemetry Tracer
	traceExporter, err := otlptracegrpc.New(
		context.TODO(),
		otlptracegrpc.WithGRPCConn(config.Conn),
	)
	if err != nil {
		panic("failed to create OTLP trace exporter")
	}

	cleanupTrace := tracing.InitTracer(context.TODO(), traceExporter)
	defer cleanupTrace()

	// Initialize the metrics
	metricExporter, err := otlpmetricgrpc.New(
		context.TODO(),
		otlpmetricgrpc.WithGRPCConn(config.Conn),
	)
	if err != nil {
		panic("failed to create OTLP trace exporter")
	}
	cleanupMeter := metrics.InitMetrics(context.TODO(), metricExporter)
	defer cleanupMeter()

	counter, err := metrics.Meter.Int64Counter(
		"api_request_count",
		metric.WithDescription("Counts the number of HTTP requests"),
	)
	if err != nil {
		panic("failed to create request counter metric")
	}

	// Initialize the logger
	logExporter, err := otlploggrpc.New(
		context.TODO(),
		otlploggrpc.WithGRPCConn(config.Conn),
		otlploggrpc.WithInsecure(),
	)
	if err != nil {
		panic(err)
	}
	err = logging.InitLogger("api", logExporter)
	if err != nil {
		panic(err)
	}
	defer logging.SLogger.Close(context.Background())

	server := gin.Default()
	server.Use(otelgin.Middleware(config.GetServiceName()))
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
