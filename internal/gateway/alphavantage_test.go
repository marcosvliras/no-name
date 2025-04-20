package gateway

import (
	"testing"

	m "github.com/marcosvliras/sophie/internal/gateway/mock_gateway"
	"github.com/stretchr/testify/assert"
)

func TestGetData(t *testing.T) {

	t.Run("TestGetData when correct payload", func(t *testing.T) {

		clientMock := m.GetHttpClientMock(0)

		gtw := AlphavantageGateway{
			Client: clientMock,
		}
		stock, err := gtw.GetData("IBM")
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		assert.Equal(t, stock.MetaData.Information, "Monthly Adjusted Prices and Volumes")
		assert.Equal(t, stock.MetaData.Symbol, "IBM")
		assert.Equal(t, stock.MetaData.LastRefreshed, "2025-04-17")
		assert.Equal(t, stock.MetaData.TimeZone, "US/Eastern")
	})

	t.Run("TestGetData when network error", func(t *testing.T) {
		clientMock := m.GetHttpClientMock(1)

		gtw := AlphavantageGateway{
			Client: clientMock,
		}
		_, err := gtw.GetData("IBM")
		assert.Contains(t, err.Error(), "fake error")
	})

	t.Run("TestGetData when http error", func(t *testing.T) {
		clientMock := m.GetHttpClientMock(2)

		gtw := AlphavantageGateway{
			Client: clientMock,
		}
		_, err := gtw.GetData("IBM")
		assert.Equal(t, err.Error(), "error: 500 Internal Server Error")
	})
}
