package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/marcosvliras/sophie/internal/controllers"
	"github.com/marcosvliras/sophie/internal/service"
	"github.com/marcosvliras/sophie/internal/tracing"

	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func main() {
	cleanup := tracing.InitTracer()
	defer cleanup()

	server := gin.Default()
	server.Use(otelgin.Middleware(tracing.ServiceName))

	svc := service.NewAlphavantageSVC()
	stockCtrl := controllers.NewStocksCtrl(svc)

	server.GET("/health", controllers.HealthCheck)
	server.GET("/stocks", stockCtrl.Handle)

	if err := server.Run(":8000"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
