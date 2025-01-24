# Makefile for my-tab-manager

# Variables
APP_NAME := tabman
BUILD_DIR := ./bin
SOURCE_DIR := ./
GO_FILES := $(shell find . -type f -name '*.go')
GO_VERSION := 1.21
LDFLAGS := -s -w

# Default target
all: build

# Build the application
build: $(BUILD_DIR)/$(APP_NAME)

$(BUILD_DIR)/$(APP_NAME): $(GO_FILES)
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -ldflags="$(LDFLAGS)" -o $(BUILD_DIR)/$(APP_NAME) $(SOURCE_DIR)

# Run the application
run: build
	@echo "Running $(APP_NAME)..."
	@$(BUILD_DIR)/$(APP_NAME)

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod download

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)/*

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Format Go code
fmt:
	@echo "Formatting Go code..."
	@go fmt ./...

# Lint Go code
lint:
	@echo "Linting Go code..."
	@golangci-lint run

# Check for outdated dependencies
tidy:
	@echo "Tidying up dependencies..."
	@go mod tidy

# Check Go version
check-go-version:
	@echo "Checking Go version..."
	@if ! go version | grep -q "go$(GO_VERSION)"; then \
		echo "Go version $(GO_VERSION) is required"; \
		exit 1; \
	fi

# Install golangci-lint (if not already installed)
install-lint:
	@echo "Installing golangci-lint..."
	@if ! command -v golangci-lint &> /dev/null; then \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.55.2; \
	fi

# Install tools (dependencies, linter, etc.)
install-tools: install-lint

# Default target
.PHONY: all build run deps clean test fmt lint tidy check-go-version install-lint install-tools