package main

import (
	"github.com/gin-gonic/gin"
	"github.com/marcosvliras/no-name/internal/controllers"
)

func main() {
	server := gin.Default()

	server.GET("/health", controllers.HealthCheck)

	server.Run(":8000")
}
