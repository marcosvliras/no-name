package main

import (
	"context"
	"fmt"
	"os"

	"github.com/marcosvliras/sophie/internal/service"
	"github.com/marcosvliras/sophie/printer"
	"github.com/spf13/cobra"
)

func main() {
	svc := service.NewAlphavantageSVC()

	var symbols []string

	var rootCmd = &cobra.Command{
		Use:  "sophie",
		Long: "cli for get stock information",
	}

	var getStocksCmd = &cobra.Command{
		Use:   "get",
		Short: "Get stock data",
		Run: func(cmd *cobra.Command, args []string) {
			if len(symbols) == 0 {
				fmt.Println("You need to pass at least one argument to `--symbols`")
				return
			}

			result := svc.GetStockData(context.TODO(), symbols)
			printer.PrintAggStockDataTable(result)
		},
	}

	getStocksCmd.Flags().StringSliceVarP(
		&symbols,
		"symbols",
		"s",
		[]string{},
		"List of symbols (ex: BBAS3,ITSA4,PETR4)",
	)

	rootCmd.AddCommand(getStocksCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
