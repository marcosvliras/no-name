package service

import "github.com/marcosvliras/sophie/stock"

type ISVC interface {
	GetStockData(data []string) []stock.AggStockData
}
