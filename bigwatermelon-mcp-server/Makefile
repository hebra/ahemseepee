# Variables
BINARY_NAME=bigwatermelon-mcp
GO=go
GOFLAGS=-v
GOCMD=$(GO)
GOBUILD=$(GOCMD) build $(GOFLAGS)
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=gofmt
GOVET=$(GOCMD) vet
GOLINT=golangci-lint

# Build variables
VERSION?=1.0.0
BUILD_DIR=build
MAIN_FILE=main.go
LDFLAGS=-ldflags "-X main.Version=${VERSION}"

# Docker variables
DOCKER_IMAGE=watermelon-mcp
DOCKER_TAG?=latest

.PHONY: all build clean test coverage fmt lint vet tidy deps docker-build docker-run help

all: clean lint test build

# Build the application
build:
	@echo "Building..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) $(LDFLAGS) $(MAIN_FILE)

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	$(GOCLEAN)

# Run tests
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

# Generate test coverage report
coverage:
	@echo "Generating coverage report..."
	$(GOTEST) -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out

# Format code
fmt:
	@echo "Formatting code..."
	$(GOFMT) -w .

# Run linter
lint:
	@echo "Running linter..."
	$(GOLINT) run

# Run go vet
vet:
	@echo "Running go vet..."
	$(GOVET) ./...

# Tidy and verify dependencies
tidy:
	@echo "Tidying dependencies..."
	$(GOMOD) tidy
	$(GOMOD) verify

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	$(GOMOD) download

# Build Docker image
docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

# Run Docker container
docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 $(DOCKER_IMAGE):$(DOCKER_TAG)

# Run the application
run: build
	@echo "Running application..."
	./$(BUILD_DIR)/$(BINARY_NAME)

# Development mode with hot reload (requires air: https://github.com/cosmtrek/air)
dev:
	@echo "Starting development server..."
	air

# Show help
help:
	@echo "Available targets:"
	@echo "  all          - Clean, lint, test, and build"
	@echo "  build        - Build the application"
	@echo "  clean        - Clean build artifacts"
	@echo "  test         - Run tests"
	@echo "  coverage     - Generate test coverage report"
	@echo "  fmt          - Format code"
	@echo "  lint         - Run linter"
	@echo "  vet          - Run go vet"
	@echo "  tidy         - Tidy and verify dependencies"
	@echo "  deps         - Download dependencies"
	@echo "  docker-build - Build Docker image"
	@echo "  docker-run   - Run Docker container"
	@echo "  run          - Run the application"
	@echo "  dev          - Run in development mode with hot reload"
	@echo "  help         - Show this help message"

# Default target
.DEFAULT_GOAL := help