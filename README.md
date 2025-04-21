# Sophie

Sophie is a Go library that provides a simple and efficient way to retrieve brazilian stock prices to help you decide which stocks to buy or sell. It uses the [Alpha Vantage](https://www.alphavantage.co/) API to fetch stock data and provides a command-line interface (CLI) for easy access.

## Why Sophie?

The name Sophie was chosen because, in most cultures, it stands for wisdom. And when it comes to investments, let's face it: you need wisdom to make good choices.

When I mentioned the name to a few friends, one of them immediately thought of the book "Sophie's Choice", which is about tough decisions and their consequences. I knew right then that it was the perfect name. By the way, much better than "no-name-project" haha.

## CLI

1. set up your environment with alphavante api key: https://www.alphavantage.co/documentation/
```
API_KEY=
```

2. Install stock cli
```
go install github.com/marcosvliras/sophie/cmd/cli/stock@latest
```

3. Add /go/bin to your PATH
```
export PATH=$PATH:$(go env GOPATH)/bin
```

4. How to use stock cli
```
stock get --symbols=BBAS3,ITUB4
```
retreives the last price of the stocks BBAS3 and ITUB4

```
STOCK      ACTUAL PRICE  MAX PRICE
BBAS3.SAO  27.43         28.15
ITSA4.SAO  9.95          7.27
```

5. Local

- 5.1 install `make`
```
brew install make
```

- 5.2 install pre-commit
```
brew install pre-commit
pre-commit install
```

- 5.3 Clone the repository

- 5.4 go mod tidy

- 5.5 Build the project
```
go build -o stock-cli cmd/cli/stock/main.go
```
