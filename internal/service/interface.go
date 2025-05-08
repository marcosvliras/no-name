package service

import (
	"github.com/marcosvliras/sophie/stock"
	"golang.org/x/net/context"
)

type ISVC interface {
	GetStockData(ctx context.Context, data []string) []stock.AggStockData
}
