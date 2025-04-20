package service

import "github.com/marcosvliras/no-name/stock"

type ISVC interface {
	GetStockData(data []string) []stock.AggStockData
}
