package printer

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/marcosvliras/sophie/stock"
)

func StdoutPrint(v interface{}) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println(string(data))
}

func PrintAggStockDataTable(data []stock.AggStockData) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	fmt.Fprintln(w, "STOCK\tACTUAL PRICE\tMAX PRICE")

	for _, item := range data {
		max := "N/A"
		actual := "N/A"
		if item.MaxStockPrice != nil {
			max = fmt.Sprintf("%.2f", *item.MaxStockPrice)
		}
		if item.ActualPrice != nil {
			actual = fmt.Sprintf("%.2f", *item.ActualPrice)
		}
		fmt.Fprintf(w, "%s\t%s\t%s\n", item.Stock, actual, max)
	}

	w.Flush()
}
