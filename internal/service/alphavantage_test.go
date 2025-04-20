package service

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	m "github.com/marcosvliras/sophie/internal/gateway/mock_gateway"
	"github.com/marcosvliras/sophie/stock"
	"github.com/stretchr/testify/assert"
)

func jsonToStockData(name string) (stock.Stock, error) {
	file, err := os.ReadFile(name)
	if err != nil {
		return stock.Stock{}, err
	}
	var stockData stock.Stock
	err = json.Unmarshal(file, &stockData)
	if err != nil {
		return stock.Stock{}, err
	}
	return stockData, nil
}

func roundToTwoDecimals(f float64) float64 {
	return math.Round(f*100) / 100
}

func TestAlphavantageSVC(t *testing.T) {

	t.Run("GetStockData when gateway work as expected", func(t *testing.T) {
		bbase, _ := jsonToStockData("mock_service/bbas3.json")
		itsa4, _ := jsonToStockData("mock_service/itsa4.json")

		gomockCtrl := gomock.NewController(t)
		defer gomockCtrl.Finish()

		gtwMock := m.NewMockIGateway(gomockCtrl)

		gtwMock.EXPECT().GetData("ITSA4").Return(itsa4, nil).Times(1)
		gtwMock.EXPECT().GetData("BBAS3").Return(bbase, nil).Times(1)

		svc := AlphavantageSVC{
			Gtw: gtwMock,
		}
		symbols := []string{"ITSA4", "BBAS3"}
		r := svc.GetStockData(symbols)

		assert.Equal(t, len(r), 2)
		assert.Equal(t, r[0].Stock, "ITSA4.SAO")
		assert.Equal(t, roundToTwoDecimals(*r[0].MaxStockPrice), roundToTwoDecimals(7.2747222222222225))
		assert.Equal(t, *r[0].ActualPrice, 9.95)

		assert.Equal(t, r[1].Stock, "BBAS3.SAO")
		assert.Equal(t, roundToTwoDecimals(*r[1].MaxStockPrice), roundToTwoDecimals(28.151388888888892))
		assert.Equal(t, *r[1].ActualPrice, 27.43)
	})

	t.Run("GetStockData when gateway return an error", func(t *testing.T) {
		gomockCtrl := gomock.NewController(t)
		defer gomockCtrl.Finish()

		gtwMock := m.NewMockIGateway(gomockCtrl)

		gtwMock.EXPECT().GetData("ITSA4").Return(stock.Stock{}, fmt.Errorf("fake error")).Times(1)

		svc := AlphavantageSVC{
			Gtw: gtwMock,
		}
		symbols := []string{"ITSA4"}
		r := svc.GetStockData(symbols)

		assert.Equal(t, len(r), 1)
		assert.Equal(t, r[0].Stock, "")
		assert.Nil(t, r[0].MaxStockPrice)
		assert.Nil(t, r[0].ActualPrice)
	})

	t.Run("GetStockData when agregateDividendPerYear return an error", func(t *testing.T) {
		gomockCtrl := gomock.NewController(t)
		defer gomockCtrl.Finish()

		gtwMock := m.NewMockIGateway(gomockCtrl)

		gtwMock.EXPECT().GetData("ITSA4").Return(stock.Stock{}, nil).Times(1)

		svc := AlphavantageSVC{
			Gtw: gtwMock,
		}
		symbols := []string{"ITSA4"}
		r := svc.GetStockData(symbols)

		assert.Equal(t, len(r), 1)
		assert.Equal(t, r[0].Stock, "")
		assert.Nil(t, r[0].MaxStockPrice)
		assert.Nil(t, r[0].ActualPrice)
	})
}
