lint:
	golangci-lint run ./...

test:
	go test -coverprofile=coverage.out ./...

coverage: test
	go tool cover -html=coverage.out -o coverage.html