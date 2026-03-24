.PHONY: build run test clean docker-build docker-up docker-down swagger proto mod fmt lint dev help

build:
	go build -o bin/server cmd/main.go

run:
	go run cmd/main.go

test:
	go test -v ./...

clean:
	rm -rf bin/

docker-build:
	docker build -t qb-accounting -f deployment/Dockerfile .

docker-up:
	docker-compose -f deployment/docker-compose.yml up -d

docker-down:
	docker-compose -f deployment/docker-compose.yml down

db-setup:
	psql -h localhost -U postgres -d qb_accounting -f schema/schema.sql

dev:
	docker-compose -f deployment/docker-compose.yml up -d postgres redis
	sleep 5
	go run cmd/main.go

fmt:
	go fmt ./...

lint:
	golangci-lint run

mod:
	go mod tidy

swagger:
	swag init -g cmd/main.go -o docs

build-prod:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/server cmd/main.go

deps:
	go list -m -u all

help:
	@echo "Available commands:"
	@echo "  build       - Build the application"
	@echo "  run         - Run the application"
	@echo "  test        - Run tests"
	@echo "  clean       - Clean build artifacts"
	@echo "  docker-*    - Docker commands"
	@echo "  db-setup    - Setup database schema"
	@echo "  dev         - Run in development mode"
	@echo "  fmt         - Format code"
	@echo "  lint        - Run linter"
	@echo "  mod         - Tidy go modules"
	@echo "  swagger     - Generate Swagger docs"
