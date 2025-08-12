##
##编译solidity，并产生bin文件，abi文件，和相应的go文件 solc版本0.8.20

currentDir := `pwd`

SRC_CONTRACTJu := ../contracts
GO_OUT4Ju := ${SRC_CONTRACTJu}/generated
OUT := build
ALLOW_PATH := $(currentDir)

PACKAGE := generated

# 项目信息
PROJECT_NAME := congress-cli
VERSION := 1.1.0
BUILD_DATE := $(shell date -u +'%Y-%m-%d_%H:%M:%S')
GIT_COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")

# 构建标志
LDFLAGS := -X 'juchain.org/chain/congress-cli/cmd.Version=$(VERSION)' \
           -X 'juchain.org/chain/congress-cli/cmd.BuildDate=$(BUILD_DATE)' \
           -X 'juchain.org/chain/congress-cli/cmd.GitCommit=$(GIT_COMMIT)'

.PHONY: default build clean proposal test lint fmt help

default: build

help: ## 显示帮助信息
	@echo "Congress CLI Makefile Commands:"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-15s %s\n", $$1, $$2}'

proposal: ## 生成 Proposal 合约的 Go 绑定
	@echo "Generating Proposal contract bindings..."
	@solc --allow-paths $(ALLOW_PATH) --optimize --combined-json abi,bin,userdoc,devdoc  $(SRC_CONTRACTJu)/Proposal.sol -o $(SRC_CONTRACTJu)/ --overwrite
	@abigen --combined-json $(SRC_CONTRACTJu)/combined.json --pkg $(PACKAGE) --out $(GO_OUT4Ju)/proposal.go
	@rm $(SRC_CONTRACTJu)/combined.json
	@echo "✅ Contract bindings generated successfully"

build: ## 编译项目
	@echo "Building $(PROJECT_NAME) v$(VERSION)..."
	@go build -ldflags "$(LDFLAGS)" -o ${OUT}/$(PROJECT_NAME)
	@echo "✅ Build completed: ${OUT}/$(PROJECT_NAME)"

build_linux: ## 编译 Linux 版本
	@echo "Building $(PROJECT_NAME) for Linux..."
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o ${OUT}/$(PROJECT_NAME)-linux-amd64
	@echo "✅ Linux build completed: ${OUT}/$(PROJECT_NAME)-linux-amd64"

# mac遇到依赖包用cgo的，需要开启cgo，先安装c交叉编译器 brew install FiloSottile/musl-cross/musl-cross
build_linux_cgo: ## 编译带 CGO 的 Linux 版本
	@echo "Building $(PROJECT_NAME) for Linux with CGO..."
	@CC=x86_64-linux-musl-gcc CXX=x86_64-linux-musl-g++ CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS) -linkmode external -extldflags -static" -o ${OUT}/$(PROJECT_NAME)-linux-static
	@echo "✅ Static Linux build completed: ${OUT}/$(PROJECT_NAME)-linux-static"

test: ## 运行测试
	@echo "Running tests..."
	@go test -v ./cmd
	@echo "✅ Tests completed"

lint: ## 运行代码检查
	@echo "Running linter..."
	@which golangci-lint > /dev/null || (echo "Installing golangci-lint..." && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
	@export PATH=$$PATH:~/go/bin && golangci-lint run --timeout=5m || echo "⚠️  Linting completed with warnings"
	@echo "✅ Linting completed"

fmt: ## 格式化代码
	@echo "Formatting code..."
	@go fmt ./...
	@echo "✅ Code formatted"

tidy: ## 整理依赖
	@echo "Tidying go modules..."
	@go mod tidy
	@echo "✅ Dependencies tidied"

install: build ## 安装到系统
	@echo "Installing $(PROJECT_NAME)..."
	@sudo cp ${OUT}/$(PROJECT_NAME) /usr/local/bin/
	@echo "✅ $(PROJECT_NAME) installed to /usr/local/bin/"

cleanContract: ## 清理生成的合约文件
	@echo "Cleaning contract files..."
	@rm -fr $(GO_OUT4Ju)/*
	@echo "✅ Contract files cleaned"

clean: ## 清理构建文件
	@echo "Cleaning build files..."
	@rm -fr build/*
	@echo "✅ Build files cleaned"

all: fmt lint test build ## 运行所有检查和构建

release: clean all build_linux ## 构建发布版本
	@echo "Creating release package..."
	@mkdir -p ${OUT}/release
	@cp ${OUT}/$(PROJECT_NAME) ${OUT}/release/$(PROJECT_NAME)-darwin-amd64
	@cp ${OUT}/$(PROJECT_NAME)-linux-amd64 ${OUT}/release/
	@cp README.md ${OUT}/release/
	@echo "✅ Release package created in ${OUT}/release/"

asmExample: ## 生成汇编示例
	@solc --asm --bin -o tmp/ valset-bin=./tmp/valset-bin Valset.sol

