.PHONY: build run test clean lint help

# Build the binary
build:
	go build -o bin/godash cmd/godash/main.go

# Run the application
run:
	go run cmd/godash/main.go

# Run tests
test:
	go test -v ./...

# Run tests with coverage
coverage:
	go test -cover ./...

# Clean build artifacts
clean:
	rm -rf bin/
	go clean

# Run linter (requires golangci-lint)
lint:
	golangci-lint run

# Install dependencies
deps:
	go mod download
	go mod tidy

# Development build with race detector
dev:
	go build -race -o bin/godash-dev cmd/godash/main.go

# Production build (optimized, stripped)
prod:
	go build -ldflags="-s -w" -o bin/godash cmd/godash/main.go

# Show help
help:
	@echo "Available targets:"
	@echo "  build    - Build the binary"
	@echo "  run      - Run the application"
	@echo "  test     - Run tests"
	@echo "  coverage - Run tests with coverage"
	@echo "  clean    - Clean build artifacts"
	@echo "  lint     - Run linter"
	@echo "  deps     - Install/update dependencies"
	@echo "  dev      - Build with race detector"
	@echo "  prod     - Build optimized binary"
	@echo "  help     - Show this help message"
