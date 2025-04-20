# Sophie

Sophie is a Go library that provides a simple and efficient way to retrieve brazilian stock prices to help you decide which stocks to buy or sell. It uses the [Alpha Vantage](https://www.alphavantage.co/) API to fetch stock data and provides a command-line interface (CLI) for easy access.

## Why Sophie?

The name Sophie was chosen because, in most cultures, it stands for wisdom. And when it comes to investments, let's face it: you need wisdom to make good choices.

When I mentioned the name to a few friends, one of them immediately thought of the book "Sophie's Choice", which is about tough decisions and their consequences. I knew right then that it was the perfect name haha.

## CLI

1. How to install stock cli
```
go install github.com/marcosvliras/sophie/cmd/cli/stock@latest
```

2. Add /go/bin to your PATH
```
export PATH=$PATH:$(go env GOPATH)/bin
```

3. How to use stock cli
```
stock get --symbols=BBAS3,ITUB4
```
retreives the last price of the stocks BBAS3 and ITUB4

```
[
  {
    "Stock": "BBAS3.SAO",
    "MaxStockPrice": 28.151388888888892,
    "ActualPrice": 27.43
  },
  {
    "Stock": "ITUB4.SAO",
    "MaxStockPrice": 31.30833333333333,
    "ActualPrice": 32.75
  }
]
```

4. Install local cli
- 4.a Clone the repository
- 4.b Build the project
```
go build -o stock-cli cmd/cli/stock/main.go
```

