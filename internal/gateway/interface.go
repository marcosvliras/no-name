package gateway

import "github.com/marcosvliras/no-name/stock"

type IGateway interface {
	GetData(symbol string) (stock.Stock, error)
}
