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
)

func main() {
	config.InitConfig()

	cleanupTrace := tracing.InitTracer()
	defer cleanupTrace()

	cleanupMeter := metrics.InitMetrics()
	defer cleanupMeter()

	err := logging.InitLogger()
	if err != nil {
		panic(err)
	}
	defer logging.SLogger.Close(context.Background())

	server := gin.Default()
	server.Use(otelgin.Middleware(config.ServiceName))
	server.Use(middleware.LoggerMiddleware())
	server.Use(middleware.RequestCounterMiddleware())

	svc := service.NewAlphavantageSVC()
	stockCtrl := controllers.NewStocksCtrl(svc)

	server.GET("/health", controllers.HealthCheck)
	server.GET("/stocks", stockCtrl.Handle)

	if err := server.Run(":8000"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
