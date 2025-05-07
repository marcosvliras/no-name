package controllers

import (
	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.String(200, "Ok")
}
