package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcosvliras/sophie/internal/service"
	"github.com/marcosvliras/sophie/internal/tracing"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type ICtrl interface {
	Handle(c *gin.Context)
}

type StocksCtrl struct {
	svc service.ISVC
}

func NewStocksCtrl(svc service.ISVC) *StocksCtrl {
	return &StocksCtrl{
		svc: svc,
	}
}

func (ctrl *StocksCtrl) Handle(c *gin.Context) {
	ctx := otel.GetTextMapPropagator().Extract(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))

	ctx, span := tracing.Tracer.Start(ctx, "StockController")
	defer span.End()

	symbolList := c.QueryArray("symbol")
	if len(symbolList) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No symbols provided"})
		return
	}

	stocks := ctrl.svc.GetStockData(ctx, symbolList)

	c.JSON(http.StatusOK, stocks)
}
