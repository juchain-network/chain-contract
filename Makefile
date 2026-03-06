# Makefile for sys-contract project
# Usage: make [target]

# Variables
ABIGEN = ../chain/build/bin/abigen
GO_CLIENT_DIR = tools/contracts
STORAGE_DIR = storage
STORAGE_LAYOUT_TARGETS = contracts/Proposal.sol:Proposal contracts/Validators.sol:Validators contracts/Punish.sol:Punish contracts/Staking.sol:Staking

# Colors for output
GREEN = \033[0;32m
YELLOW = \033[1;33m
RED = \033[0;31m
NC = \033[0m # No Color

# Define all phony targets
.PHONY: help clean build test fmt security coverage coverage-html gas-test all update version addresses generate-contracts generate-go-client storage-layout test-by-forge test-by-shell anvil-start anvil-stop anvil-status anvil-clean get-system-params load-env

help:
	@echo "$(YELLOW)Available targets:$(NC)"
	@echo ""
	@echo "$(YELLOW)Core Tasks:$(NC)"
	@echo "  $(GREEN)build$(NC)               - Compile contracts and generate artifacts"
	@echo "  $(GREEN)bytecode$(NC)            - Build contracts and extract bytecode"
	@echo "  $(GREEN)test$(NC)                - Run all tests using forge"
	@echo "  $(GREEN)clean$(NC)               - Clean build artifacts and test environment"
	@echo "  $(GREEN)fmt$(NC)                 - Format Solidity code"
	@echo "  $(GREEN)all$(NC)                 - Run clean + build + test"
	@echo ""
	@echo "$(YELLOW)Testing & Analysis:$(NC)"
	@echo "  $(GREEN)security$(NC)            - Run security analysis with Slither (if installed)"
	@echo "  $(GREEN)coverage$(NC)            - Generate test coverage report"
	@echo "  $(GREEN)coverage-html$(NC)       - Generate HTML test coverage report"
	@echo "  $(GREEN)gas-test$(NC)            - Run gas optimization tests with report"
	@echo ""
	@echo "$(YELLOW)Development Utilities:$(NC)"
	@echo "  $(GREEN)update$(NC)              - Update dependencies"
	@echo "  $(GREEN)version$(NC)             - Show forge version and dependencies"
	@echo "  $(GREEN)addresses$(NC)           - Show system contract addresses"
	@echo "  $(GREEN)generate-contracts$(NC)   - Generate production contracts from templates"
	@echo "  $(GREEN)generate-go-client$(NC)   - Generate Go client code using abigen"
	@echo "  $(GREEN)storage-layout$(NC)        - Export storage layout JSON to ./storage"
	@echo "  $(GREEN)get-system-params$(NC)    - Get system configuration parameters from deployed contracts"
	@echo ""
	@echo "$(YELLOW)Anvil Test Environment:$(NC)"
	@echo "  $(GREEN)anvil-start$(NC)          - Start Anvil test node"
	@echo "  $(GREEN)anvil-stop$(NC)           - Stop Anvil test node"
	@echo "  $(GREEN)anvil-status$(NC)         - Check Anvil node status"
	@echo "  $(GREEN)anvil-clean$(NC)          - Clean Anvil logs and temporary files"
	@echo ""
	@echo "$(YELLOW)Environment Management:$(NC)"
	@echo "  $(GREEN)load-env$(NC)             - Load environment variables from .env file"
	@echo ""
	@echo "Usage examples:"
	@echo "  make build       - Build all contracts"
	@echo "  make bytecode    - Build contracts and extract bytecode"
	@echo "  make clean build - Clean and then build"
	@echo ""

# Clean build artifacts
clean:
	@echo "$(YELLOW)Cleaning build artifacts...$(NC)"
	@forge clean
	@make anvil-stop
	@make anvil-clean
	@rm -fr state

# Build contracts
build:generate-contracts
	@echo "$(YELLOW)Building contracts...$(NC)"
	forge build

# Build contracts and extract bytecode
bytecode: clean generate-contracts build
	@echo "$(YELLOW)Building contracts and extracting bytecode...$(NC)"
	npm run build-and-extract
# Run tests
test: build
	@echo "$(YELLOW)Running tests...$(NC)"
	@forge test

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
	@forge coverage

# HTML Coverage report
coverage-html:
	@echo "$(YELLOW)Generating HTML coverage report...$(NC)"
	@echo "$(YELLOW)Step 1: Installing lcov if not present...$(NC)"
# 	@brew install lcov >/dev/null 2>&1 || true
	@echo "$(YELLOW)Step 2: Generating lcov report...$(NC)"
	@forge coverage --report lcov
	@echo "$(YELLOW)Step 3: Converting lcov to HTML...$(NC)"
	@genhtml lcov.info -o coverage-report
	@echo "$(YELLOW)Step 4: Cleaning up temporary files...$(NC)"
	@rm -f lcov.info
	@echo "$(GREEN)✅ HTML coverage report generated successfully at ./coverage-report$(NC)"
	@echo "$(YELLOW)To view the report: open coverage-report/index.html$(NC)"

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
	@echo "$(YELLOW)Generating contracts from templates...$(NC)"
	node generate-contracts.js
	@echo "$(GREEN)✅ Contracts generated successfully$(NC)"

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
generate-go-client: clean generate-contracts build
	@echo "$(YELLOW)Generating Go client code...$(NC)"
	@node generate-go-client.js

# Export storage layout for system contracts
storage-layout: build
	@echo "$(YELLOW)Exporting storage layouts...$(NC)"
	@mkdir -p $(STORAGE_DIR)
	@for target in $(STORAGE_LAYOUT_TARGETS); do \
		name=$${target##*:}; \
		echo "  - $$name"; \
		forge inspect $$target storage-layout > $(STORAGE_DIR)/$$name.storage.json; \
	done
	@echo "$(GREEN)✅ Storage layouts exported to ./$(STORAGE_DIR)$(NC)"

# =========================
# Anvil Test Environment
# =========================

# Start Anvil test node
anvil-start:
	@mkdir -p state
	@echo "$(YELLOW)Starting Anvil test node...$(NC)"
	@./anvil-setup.sh --start

# Stop Anvil test node
anvil-stop:
	@echo "$(YELLOW)Stopping Anvil test node...$(NC)"
	@./anvil-setup.sh --stop

# Check Anvil node status
anvil-status:
	@./anvil-setup.sh --status

# Clean Anvil logs and temporary files
anvil-clean:
	@echo "$(YELLOW)Cleaning Anvil logs and temporary files...$(NC)"
	@./anvil-setup.sh --clean

# Get system configuration parameters
get-system-params:
	@echo "$(YELLOW)System Configuration Parameters$(NC)"
	@echo -n "Epoch Duration: "; cast call 0x000000000000000000000000000000000000f010 "getEpochDuration()(uint256)" | grep -oE '[0-9]+'; echo -n " blocks"
	@echo -n "Unbonding Period: "; cast call 0x000000000000000000000000000000000000F013 "getUnbondingPeriod()(uint256)" | grep -oE '[0-9]+'; echo -n " blocks"
	@echo -n "Validator Unjail Period: "; cast call 0x000000000000000000000000000000000000F013 "getValidatorUnjailPeriod()(uint256)" | grep -oE '[0-9]+'; echo -n " blocks"
	@echo -n "Proposal Lasting Period: "; cast call 0x000000000000000000000000000000000000F012 "getProposalLastingPeriod()(uint256)" | grep -oE '[0-9]+'; echo -n " blocks"

# =========================
# Test Utilities
# =========================

# Load environment variables from .env file
load-env:
	@echo "$(YELLOW)Loading environment variables from .env file...$(NC)"
	@if [ -f .env ]; then \
		export $$(grep -v '^#' .env | xargs); \
		echo "$(GREEN)✅ Environment variables loaded successfully!$(NC)"; \
	else \
		echo "$(YELLOW)Warning: .env file not found. Using default values.$(NC)"; \
	fi
