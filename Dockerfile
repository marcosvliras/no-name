FROM golang:1.24.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOARCH=amd64 GOOS=linux go build -o /app/bin/main ./cmd/api/main.go

FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/bin/main .

RUN apt-get update && apt-get install -y ca-certificates

CMD ["./main"]
