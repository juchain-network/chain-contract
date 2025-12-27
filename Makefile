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

help:
# Available targets:
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
	@echo "$(YELLOW)Anvil Test Environment:$(NC)"
	@echo "  $(GREEN)anvil-start$(NC)    - Start Anvil test node"
	@echo "  $(GREEN)anvil-stop$(NC)     - Stop Anvil test node"
	@echo "  $(GREEN)anvil-status$(NC)   - Check Anvil node status"
	@echo "  $(GREEN)anvil-clean$(NC)    - Clean Anvil logs and temporary files"
	@echo ""
	@echo "$(YELLOW)PoSA Integration Tests:$(NC)"
	@echo "  $(GREEN)test-all$(NC)       - Run comprehensive PoSA test suite (includes all test scenarios, report generation, and environment management)"
	@echo ""
	@echo "$(YELLOW)Test Utilities:$(NC)"
	@echo "  $(GREEN)load-env$(NC)       - Load environment variables from .env file"
	@echo ""

# Clean build artifacts
clean:
	@echo "$(YELLOW)Cleaning build artifacts...$(NC)"
	forge clean
	@make anvil-stop
	@make anvil-clean
	@make test-env-clean

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

# =========================
# Anvil Test Environment
# =========================

# Start Anvil test node
anvil-start:
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

# =========================
# PoSA Integration Tests
# =========================

# Generate mock contracts before running tests
test-env:
	@echo "$(YELLOW)Setting up test environment with mock contracts...$(NC)"
	@make generate-contracts-mock
	@echo "$(GREEN)✅ Test environment set up successfully$(NC)"

# Clean test environment and regenerate production contracts
test-env-clean:
	@echo "$(YELLOW)Cleaning test environment...$(NC)"
	@make generate-contracts
	@echo "$(GREEN)✅ Test environment cleaned$(NC)"


# Run all tests in one comprehensive task
# This task includes: test-env setup, anvil start/stop for each test, test-env-clean
test-debug:
	@make clean
	@echo "$(YELLOW)Starting comprehensive PoSA test suite...$(NC)"
	
	# Step 1: Set up test environment with mock contracts
	@make test-env
	@make load-env
	
	# Run Validator Management test with fresh Anvil instance
	@echo "\n$(YELLOW)Running Validator Management tests...$(NC)"
	@make anvil-start
	@sleep 3
	@forge script script/integration/ValidatorManagement.s.sol:ValidatorManagementScript --fork-url http://localhost:8545 --broadcast --skip-simulation -vvv
	@make anvil-stop
	@make anvil-clean
	
	# Step 6: Clean test environment
	@make test-env-clean
	
	@echo "\n$(GREEN)✅ All tests completed successfully!$(NC)"

# Run all tests in one comprehensive task
# This task includes: test-env setup, anvil start/stop for each test, test-env-clean
test-all:
	@echo "$(YELLOW)Starting comprehensive PoSA test suite...$(NC)"
	
	# Step 1: Set up test environment with mock contracts
	@make test-env
	
	# Run Deployment test with fresh Anvil instance
	@echo "\n$(YELLOW)Running Deployment tests...$(NC)"
	@make anvil-start
	@forge script script/tools/Deployment.s.sol:DeploymentScript --fork-url http://localhost:8545 --broadcast --sender 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 --gas-limit 100000000 --gas-price 1
	@make anvil-stop
	@make anvil-clean
	
	# Run Validator Management test with fresh Anvil instance
	@echo "\n$(YELLOW)Running Validator Management tests...$(NC)"
	@make anvil-start
	@forge script script/integration/ValidatorManagement.s.sol:ValidatorManagementScript --fork-url http://localhost:8545 --broadcast --sender 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 --gas-limit 100000000 --gas-price 1
	@make anvil-stop
	@make anvil-clean
	
	# Run Validator Lifecycle test with fresh Anvil instance
	@echo "\n$(YELLOW)Running Validator Lifecycle tests...$(NC)"
	@make anvil-start
	@forge script script/integration/ValidatorLifecycleTest.s.sol:ValidatorLifecycleTest --fork-url http://localhost:8545 --broadcast --sender 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 --gas-limit 100000000 --gas-price 1
	@make anvil-stop
	@make anvil-clean
	
	# Run Delegator Lifecycle test with fresh Anvil instance
	@echo "\n$(YELLOW)Running Delegator Lifecycle tests...$(NC)"
	@make anvil-start
	@forge script script/integration/DelegatorLifecycleTest.s.sol:DelegatorLifecycleTest --fork-url http://localhost:8545 --broadcast --sender 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 --gas-limit 100000000 --gas-price 1
	@make anvil-stop
	@make anvil-clean
	
	# Run Staking Mechanism test with fresh Anvil instance
	@echo "\n$(YELLOW)Running Staking Mechanism tests...$(NC)"
	@make anvil-start
	@forge script script/integration/StakingMechanism.s.sol:StakingMechanismScript --fork-url http://localhost:8545 --broadcast --sender 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 --gas-limit 100000000 --gas-price 1
	@make anvil-stop
	@make anvil-clean
	
	# Run Proposal System test with fresh Anvil instance
	@echo "\n$(YELLOW)Running Proposal System tests...$(NC)"
	@make anvil-start
	@forge script script/tools/ProposalSystem.s.sol:ProposalSystemScript --fork-url http://localhost:8545 --broadcast --sender 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 --gas-limit 100000000 --gas-price 1
	@make anvil-stop
	@make anvil-clean
	
	# Run Punishment Mechanism test with fresh Anvil instance
	@echo "\n$(YELLOW)Running Punishment Mechanism tests...$(NC)"
	@make anvil-start
	@forge script script/integration/PunishmentMechanism.s.sol:PunishmentMechanismScript --fork-url http://localhost:8545 --broadcast --sender 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 --gas-limit 100000000 --gas-price 1
	@make anvil-stop
	@make anvil-clean
	
	# Run PoSA Integration test with fresh Anvil instance
	@echo "\n$(YELLOW)Running PoSA Integration tests...$(NC)"
	@make anvil-start
	@forge script script/integration/PoSAIntegrationTest.s.sol:PoSAIntegrationTest --fork-url http://localhost:8545 --broadcast --sender 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 --gas-limit 100000000 --gas-price 1
	@make anvil-stop
	@make anvil-clean
	
	# Step 6: Clean test environment
	@make test-env-clean
	
	@echo "\n$(GREEN)✅ All tests completed successfully!$(NC)"

test-report:
	@echo "$(YELLOW)Starting comprehensive PoSA test suite with report generation...$(NC)"
	
	# Step 1: Set up test environment with mock contracts
	@make test-env
	
	# Step 2: Start Anvil test node
	@make anvil-start

	# Step 3: Run test-report to generate comprehensive test results
	@echo "$(YELLOW)Running tests and generating reports...$(NC)"
	@python3 test-report.py --broadcast
	
	# Step 4: Stop Anvil test node
	@make anvil-stop
	
	# Step 5: Clean Anvil logs and temporary files
	@make anvil-clean
	
	# Step 6: Clean test environment
	@make test-env-clean
	
	@echo "$(GREEN)✅ All tests completed successfully!$(NC)"
	@echo "$(GREEN)Test reports generated in ./test-results directory$(NC)"
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
