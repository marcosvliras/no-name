package gateway

import "github.com/marcosvliras/sophie/stock"

type IGateway interface {
	GetData(symbol string) (stock.Stock, error)
}
