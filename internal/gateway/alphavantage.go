package gateway

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/marcosvliras/sophie/stock"
)

type AlphavantageGateway struct {
	Client *http.Client
}

func NewAlphavantageGateway() AlphavantageGateway {
	return AlphavantageGateway{
		Client: &http.Client{},
	}
}

func (a *AlphavantageGateway) GetData(symbol string) (stock.Stock, error) {

	stockStruct := stock.Stock{}

	apiKey := os.Getenv("API_KEY")
	url := fmt.Sprintf(
		"https://www.alphavantage.co/query?function=TIME_SERIES_MONTHLY_ADJUSTED&symbol=%s.SAO&apikey=%s",
		symbol,
		apiKey,
	)

	res, err := a.Client.Get(url)
	if err != nil {
		return stockStruct, err
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return stockStruct, err
		}

		err = json.Unmarshal(body, &stockStruct)
		if err != nil {
			return stockStruct, err
		}

		return stockStruct, nil
	} else {
		return stockStruct, fmt.Errorf("error: %s", res.Status)
	}

}
