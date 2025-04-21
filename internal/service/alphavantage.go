package service

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/marcosvliras/sophie/internal/gateway"
	"github.com/marcosvliras/sophie/stock"
)

type AlphavantageSVC struct {
	Gtw gateway.IGateway
}

func NewAlphavantageSVC() *AlphavantageSVC {
	gtw := gateway.NewAlphavantageGateway()
	return &AlphavantageSVC{
		Gtw: &gtw,
	}
}

func (svc *AlphavantageSVC) GetStockData(data []string) []stock.AggStockData {
	var wg sync.WaitGroup
	stockData := make([]stock.AggStockData, len(data))

	wg.Add(len(data))
	for index, symbol := range data {
		go func(symbol string, index int) {

			defer wg.Done()
			aggData, err := svc.getSingleStockData(symbol)
			if err != nil {
				fmt.Println("Error fetching data for symbol:", symbol, "Error:", err)
			}
			stockData[index] = aggData
		}(symbol, index)
	}
	wg.Wait()

	return stockData
}
func (svc *AlphavantageSVC) getSingleStockData(symbol string) (stock.AggStockData, error) {

	data, err := svc.Gtw.GetData(symbol)
	if err != nil {
		return stock.AggStockData{}, err
	}
	aggData, err := svc.agregateDividendPerYear(data)
	if err != nil {
		return stock.AggStockData{}, err
	}
	return aggData, nil
}

func (svc *AlphavantageSVC) agregateDividendPerYear(data stock.Stock) (stock.AggStockData, error) {

	stockAgg := stock.AggStockData{}
	currentYear := time.Now().Year()
	endIntervalYear := currentYear - 6

	yearDividends := map[int]float64{}

	var latestDate string

	//TODO: Do not need iterate over all the years, just the last 6
	for key, rawData := range data.MonthTimeSeries {
		yearStr := strings.Split(key, "-")[0]
		year, err := strconv.Atoi(yearStr)
		if err != nil {
			return stockAgg, err
		}

		if year < currentYear && year >= endIntervalYear {
			div, _ := strconv.ParseFloat(rawData.DividendAmount, 64)
			yearDividends[year] += div
		}

		if key > latestDate {
			latestDate = key
		}

	}

	var total float64
	for _, sum := range yearDividends {
		total += sum
	}

	meanAnnualDividend := total / float64(len(yearDividends))
	maxStockPrice := meanAnnualDividend / 0.06
	actualPrice := data.MonthTimeSeries[latestDate].Close
	actualPriceFloat, err := strconv.ParseFloat(actualPrice, 64)
	if err != nil {
		return stockAgg, err
	}

	return stock.AggStockData{
		Stock:         data.MetaData.Symbol,
		MaxStockPrice: &maxStockPrice,
		ActualPrice:   &actualPriceFloat,
	}, nil

}
