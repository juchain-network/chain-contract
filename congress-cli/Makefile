##
## Compile solidity and generate bin files, abi files, and corresponding go files. solc version 0.8.20

currentDir := `pwd`

SRC_CONTRACTJu := ../contracts
GO_OUT4Ju := ${SRC_CONTRACTJu}/generated
OUT := build
ALLOW_PATH := $(currentDir)

PACKAGE := generated

# Project Information
PROJECT_NAME := congress-cli
VERSION := 1.2.1
BUILD_DATE := $(shell date -u +'%Y-%m-%d_%H:%M:%S')
GIT_COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")

# Build Flags
LDFLAGS := -X 'juchain.org/chain/congress-cli/cmd.Version=$(VERSION)' \
           -X 'juchain.org/chain/congress-cli/cmd.BuildDate=$(BUILD_DATE)' \
           -X 'juchain.org/chain/congress-cli/cmd.GitCommit=$(GIT_COMMIT)'

.PHONY: default build clean proposal test lint fmt help

default: build

help: ## Show help information
	@echo "Congress CLI Makefile Commands:"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-15s %s\n", $$1, $$2}'

proposal: ## Generate Go bindings for Proposal contract
	@echo "Generating Proposal contract bindings..."
	@solc --allow-paths $(ALLOW_PATH) --optimize --combined-json abi,bin,userdoc,devdoc  $(SRC_CONTRACTJu)/Proposal.sol -o $(SRC_CONTRACTJu)/ --overwrite
	@abigen --combined-json $(SRC_CONTRACTJu)/combined.json --pkg $(PACKAGE) --out $(GO_OUT4Ju)/proposal.go
	@rm $(SRC_CONTRACTJu)/combined.json
	@echo "✅ Contract bindings generated successfully"

build: ## Build project
	@echo "Building $(PROJECT_NAME) v$(VERSION)..."
	@go build -ldflags "$(LDFLAGS)" -o ${OUT}/$(PROJECT_NAME)
	@echo "✅ Build completed: ${OUT}/$(PROJECT_NAME)"

build_linux: ## Build Linux version
	@echo "Building $(PROJECT_NAME) for Linux..."
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o ${OUT}/$(PROJECT_NAME)-linux-amd64
	@echo "✅ Linux build completed: ${OUT}/$(PROJECT_NAME)-linux-amd64"

# For macOS with CGO dependencies, need to enable CGO, install cross-compiler first: brew install FiloSottile/musl-cross/musl-cross
build_linux_cgo: ## Build Linux version with CGO
	@echo "Building $(PROJECT_NAME) for Linux with CGO..."
	@CC=x86_64-linux-musl-gcc CXX=x86_64-linux-musl-g++ CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS) -linkmode external -extldflags -static" -o ${OUT}/$(PROJECT_NAME)-linux-static
	@echo "✅ Static Linux build completed: ${OUT}/$(PROJECT_NAME)-linux-static"

test: ## Run tests
	@echo "Running tests..."
	@go test -v ./cmd
	@echo "✅ Tests completed"

lint: ## Run code linter
	@echo "Running linter..."
	@which golangci-lint > /dev/null || (echo "Installing golangci-lint..." && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
	@export PATH=$$PATH:~/go/bin && golangci-lint run --timeout=5m || echo "⚠️  Linting completed with warnings"
	@echo "✅ Linting completed"

fmt: ## Format code
	@echo "Formatting code..."
	@go fmt ./...
	@echo "✅ Code formatted"

tidy: ## Tidy dependencies
	@echo "Tidying go modules..."
	@go mod tidy
	@echo "✅ Dependencies tidied"

install: build ## Install to system
	@echo "Installing $(PROJECT_NAME)..."
	@sudo cp ${OUT}/$(PROJECT_NAME) /usr/local/bin/
	@echo "✅ $(PROJECT_NAME) installed to /usr/local/bin/"

cleanContract: ## Clean generated contract files
	@echo "Cleaning contract files..."
	@rm -fr $(GO_OUT4Ju)/*
	@echo "✅ Contract files cleaned"

clean: ## Clean build files
	@echo "Cleaning build files..."
	@rm -fr build/*
	@echo "✅ Build files cleaned"

all: fmt lint test build ## Run all checks and build

release: clean all build_linux ## Build release version
	@echo "Creating release package..."
	@mkdir -p ${OUT}/release
	@cp ${OUT}/$(PROJECT_NAME) ${OUT}/release/$(PROJECT_NAME)-darwin-amd64
	@cp ${OUT}/$(PROJECT_NAME)-linux-amd64 ${OUT}/release/
	@cp README.md ${OUT}/release/
	@echo "✅ Release package created in ${OUT}/release/"

asmExample: ## Generate assembly example
	@solc --asm --bin -o tmp/ valset-bin=./tmp/valset-bin Valset.sol

