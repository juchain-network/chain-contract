# Makefile for sys-contract project
# Usage: make [target]

# Variables
ABIGEN = ../chain/build/bin/abigen
GO_CLIENT_DIR = tools/contracts

# Colors for output
GREEN = \033[0;32m
YELLOW = \033[1;33m
RED = \033[0;31m
NC = \033[0m # No Color

.PHONY: help clean build test fmt security coverage gas-test all update version addresses generate-contracts generate-contracts-mock generate-go-client

# Default target
help:
	@echo "$(YELLOW)Available targets:$(NC)"
	@echo "  $(GREEN)build$(NC)         - Compile contracts"
	@echo "  $(GREEN)test$(NC)          - Run all tests"
	@echo "  $(GREEN)clean$(NC)         - Clean build artifacts"
	@echo "  $(GREEN)fmt$(NC)           - Format Solidity code"
	@echo "  $(GREEN)security$(NC)      - Run security analysis (requires slither)"
	@echo "  $(GREEN)coverage$(NC)      - Generate coverage report"
	@echo "  $(GREEN)gas-test$(NC)      - Run gas optimization tests"
	@echo "  $(GREEN)all$(NC)           - Clean + build + test"
	@echo "  $(GREEN)update$(NC)        - Update dependencies"
	@echo "  $(GREEN)version$(NC)       - Show forge version and dependencies"
	@echo "  $(GREEN)addresses$(NC)     - Show system contract addresses"
	@echo "  $(GREEN)generate-contracts$(NC) - Generate production contracts from templates"
	@echo "  $(GREEN)generate-contracts-mock$(NC) - Generate mock contracts for testing"
	@echo "  $(GREEN)generate-go-client$(NC) - Generate Go client code using abigen"
	@echo ""

# Clean build artifacts
clean:
	@echo "$(YELLOW)Cleaning build artifacts...$(NC)"
	forge clean

# Build contracts
build:
	@echo "$(YELLOW)Building contracts...$(NC)"
	forge build
# Run tests
test:
	@echo "$(YELLOW)Running tests...$(NC)"
	forge test

# Full build and test
all: clean build test
	@echo "$(GREEN)All tasks completed successfully!$(NC)"

# Show contract addresses (after deployment)
addresses:
	@echo "$(YELLOW)System Contract Addresses:$(NC)"
	@echo "  Validators: 0x000000000000000000000000000000000000f010"
	@echo "  Punish:     0x000000000000000000000000000000000000F011"
	@echo "  Proposal:   0x000000000000000000000000000000000000F012"
	@echo "  Staking:    0x000000000000000000000000000000000000F013"

# Gas optimization tests
gas-test:
	@echo "$(YELLOW)Running gas optimization tests...$(NC)"
	forge test --gas-report

# Coverage report
coverage:
	@echo "$(YELLOW)Generating coverage report...$(NC)"
	forge coverage

# Security check with slither (if installed)
security:
	@echo "$(YELLOW)Running security checks...$(NC)"
	@if command -v slither >/dev/null 2>&1; then \
		slither . --print human-summary; \
	else \
		echo "$(RED)Slither not installed. Install with: pip install slither-analyzer$(NC)"; \
	fi

# Format code
fmt:
	@echo "$(YELLOW)Formatting code...$(NC)"
	forge fmt

# Generate contracts from templates
generate-contracts:
	@echo "$(YELLOW)Generating production contracts from templates...$(NC)"
	node generate-contracts.js
	@echo "$(GREEN)✅ Production contracts generated$(NC)"

# Generate mock contracts for testing
generate-contracts-mock:
	@echo "$(YELLOW)Generating mock contracts from templates...$(NC)"
	node generate-contracts.js --mock
	@echo "$(GREEN)✅ Mock contracts generated$(NC)"

# Update dependencies
update:
	@echo "$(YELLOW)Updating dependencies...$(NC)"
	forge update

# Show forge version and dependencies
version:
	@echo "$(YELLOW)Forge version:$(NC)"
	@forge --version
	@echo ""
	@echo "$(YELLOW)Dependencies:$(NC)"
	@forge tree --no-dedupe

# Generate Go client code using abigen
generate-go-client: build
	@echo "$(YELLOW)Generating Go client code...$(NC)"
	@mkdir -p $(GO_CLIENT_DIR)
	@echo "$(GREEN)Using abigen: $(ABIGEN)$(NC)"
	@echo "$(GREEN)Output directory: $(GO_CLIENT_DIR)$(NC)"
	
	# Check if jq is installed
	@which jq > /dev/null 2>&1 || { echo "$(RED)Error: jq is not installed. Please install jq first.$(NC)"; exit 1; }
	
	# Create temporary directory for ABI and bytecode files
	@mkdir -p .tmp
	
	# Generate Go client for Validators contract
	@echo "$(YELLOW)Generating Validators Go client...$(NC)"
	@jq '.abi' out/Validators.sol/Validators.json > .tmp/Validators.abi 2>/dev/null
	@jq -r '.bytecode.object' out/Validators.sol/Validators.json > .tmp/Validators.bin 2>/dev/null
	@if [ -s .tmp/Validators.abi ] && [ -s .tmp/Validators.bin ]; then \
		$(ABIGEN) \
			--abi=.tmp/Validators.abi \
			--bin=.tmp/Validators.bin \
			--pkg=contracts \
			--type=Validators \
			--out=$(GO_CLIENT_DIR)/validators.go 2>/dev/null \
		&& echo "$(GREEN)✅ Validators Go client generated successfully!$(NC)" \
		|| echo "$(RED)Failed to generate Validators Go client$(NC)"; \
	else \
		@echo "$(RED)Failed to extract Validators ABI or bytecode$(NC)"; \
	fi
	
	# Generate Go client for Staking contract
	@echo "$(YELLOW)Generating Staking Go client...$(NC)"
	@jq '.abi' out/Staking.sol/Staking.json > .tmp/Staking.abi 2>/dev/null
	@jq -r '.bytecode.object' out/Staking.sol/Staking.json > .tmp/Staking.bin 2>/dev/null
	@if [ -s .tmp/Staking.abi ] && [ -s .tmp/Staking.bin ]; then \
		$(ABIGEN) \
			--abi=.tmp/Staking.abi \
			--bin=.tmp/Staking.bin \
			--pkg=contracts \
			--type=Staking \
			--out=$(GO_CLIENT_DIR)/staking.go 2>/dev/null \
		&& echo "$(GREEN)✅ Staking Go client generated successfully!$(NC)" \
		|| echo "$(RED)Failed to generate Staking Go client$(NC)"; \
	else \
		@echo "$(RED)Failed to extract Staking ABI or bytecode$(NC)"; \
	fi
	
	# Generate Go client for Proposal contract
	@echo "$(YELLOW)Generating Proposal Go client...$(NC)"
	@jq '.abi' out/Proposal.sol/Proposal.json > .tmp/Proposal.abi 2>/dev/null
	@jq -r '.bytecode.object' out/Proposal.sol/Proposal.json > .tmp/Proposal.bin 2>/dev/null
	@if [ -s .tmp/Proposal.abi ] && [ -s .tmp/Proposal.bin ]; then \
		$(ABIGEN) \
			--abi=.tmp/Proposal.abi \
			--bin=.tmp/Proposal.bin \
			--pkg=contracts \
			--type=Proposal \
			--out=$(GO_CLIENT_DIR)/proposal.go 2>/dev/null \
		&& echo "$(GREEN)✅ Proposal Go client generated successfully!$(NC)" \
		|| echo "$(RED)Failed to generate Proposal Go client$(NC)"; \
	else \
		@echo "$(RED)Failed to extract Proposal ABI or bytecode$(NC)"; \
	fi
	
	# Generate Go client for Punish contract
	@echo "$(YELLOW)Generating Punish Go client...$(NC)"
	@jq '.abi' out/Punish.sol/Punish.json > .tmp/Punish.abi 2>/dev/null
	@jq -r '.bytecode.object' out/Punish.sol/Punish.json > .tmp/Punish.bin 2>/dev/null
	@if [ -s .tmp/Punish.abi ] && [ -s .tmp/Punish.bin ]; then \
		$(ABIGEN) \
			--abi=.tmp/Punish.abi \
			--bin=.tmp/Punish.bin \
			--pkg=contracts \
			--type=Punish \
			--out=$(GO_CLIENT_DIR)/punish.go 2>/dev/null \
		&& echo "$(GREEN)✅ Punish Go client generated successfully!$(NC)" \
		|| echo "$(RED)Failed to generate Punish Go client$(NC)"; \
	else \
		@echo "$(RED)Failed to extract Punish ABI or bytecode$(NC)"; \
	fi
	
	# Clean up temporary files
	@rm -rf .tmp
	
	# Show generated files
	@echo "$(YELLOW)Files generated:$(NC)"
	@ls -la $(GO_CLIENT_DIR)/ 2>/dev/null || echo "$(RED)No files generated$(NC)"
	
	@echo "$(GREEN)✅ Go client code generation completed!$(NC)"
