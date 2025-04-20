package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcosvliras/sophie/internal/service"
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
	symbolList := c.QueryArray("symbol")
	if len(symbolList) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No symbols provided"})
		return
	}

	stocks := ctrl.svc.GetStockData(symbolList)

	c.JSON(http.StatusOK, stocks)
}
