.PHONY: help build run clean test lint vet fmt practice list add deps

BINARY_NAME=interview-cli
BUILD_DIR=.

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the binary
	go build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/interview-cli

run: build ## Build and run the binary
	./$(BINARY_NAME)

clean: ## Remove the binary
	rm -f $(BINARY_NAME)

test: ## Run tests
	go test ./...

lint: ## Run linter and vet
	go vet ./...
	golint ./...

vet: ## Run go vet
	go vet ./...

fmt: ## Format code
	go fmt ./...

# Run with specific commands
practice: build ## Start practice session
	./$(BINARY_NAME) -practice

list: build ## List all questions
	./$(BINARY_NAME) -list

add: build ## Add a new question
	./$(BINARY_NAME) -add

deps: ## Install dependencies
	go get -u golang.org/x/lint/golint

.DEFAULT_GOAL := help