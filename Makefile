.PHONY: help build run clean lint vet fmt practice list add deps test test-short test-verbose test-coverage

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


lint: ## Run linter and vet
	go vet ./...
	golint ./...

vet: ## Run go vet
	go vet ./...

fmt: ## Format code
	go fmt ./...

# Run with specific commands
practice: build ## Start practice session (N=<num> CATEGORY=<category> TAGS=<tags>)
	./$(BINARY_NAME) -practice $(if $(N),-n $(N),) $(if $(CATEGORY),-category $(CATEGORY),) $(if $(TAGS),-tags $(TAGS),)

list: build ## List all questions
	./$(BINARY_NAME) -list

add: build ## Add a new question
	./$(BINARY_NAME) -add

deps: ## Install dependencies
	go get -u golang.org/x/lint/golint

test: ## Run all tests
	go test -v ./tests/...

test-short: ## Run only unit tests (skip integration tests)
	go test -v -short ./tests/...

test-verbose: ## Run tests with verbose output
	go test -v -count=1 ./tests/...

test-coverage: ## Run tests with coverage report
	go test -coverprofile=coverage.out ./tests/...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated at coverage.html"

.DEFAULT_GOAL := help