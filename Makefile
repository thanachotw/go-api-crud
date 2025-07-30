.PHONY: build run swag test

build:
	go build -o app ./cmd

run:
	go run ./cmd/main.go serve-rest

swag:
	swag init -g cmd/main.go

test:
	go test -v -cover ./...

test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html