package stock

type MetaData struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	TimeZone      string `json:"4. Time Zone"`
}

type MonthlyAdjustedTimeSeries struct {
	Open           string `json:"1. open"`
	High           string `json:"2. high"`
	Low            string `json:"3. low"`
	Close          string `json:"4. close"`
	AdjustedClose  string `json:"5. adjusted close"`
	Volume         string `json:"6. volume"`
	DividendAmount string `json:"7. dividend amount"`
}

type Stock struct {
	MetaData        MetaData                             `json:"Meta Data"`
	MonthTimeSeries map[string]MonthlyAdjustedTimeSeries `json:"Monthly Adjusted Time Series"`
}

type AggStockData struct {
	Stock         string   `json:"Stock"`
	MaxStockPrice *float64 `json:"MaxStockPrice"`
	ActualPrice   *float64 `json:"ActualPrice"`
}
