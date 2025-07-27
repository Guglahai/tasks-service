run:
	go run cmd/server/main.go

lint:
	golangci-lint run --out-format=colored-line-number