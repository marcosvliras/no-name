package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/marcosvliras/no-name/internal/controllers"
	"github.com/marcosvliras/no-name/internal/service"
)

func main() {
	server := gin.Default()

	svc := service.NewAlphavantageSVC()
	stockCtrl := controllers.NewStocksCtrl(svc)

	server.GET("/health", controllers.HealthCheck)
	server.GET("/stocks", stockCtrl.Handle)

	if err := server.Run(":8000"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
