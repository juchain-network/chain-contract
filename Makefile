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

.PHONY: help clean build test fmt security coverage gas-test all update version addresses generate-contracts  generate-go-client

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
	@forge clean
	@make anvil-stop
	@make anvil-clean

# Build contracts
build:generate-contracts
	@echo "$(YELLOW)Building contracts...$(NC)"
	forge build
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

# Generate contracts from templates (both production and integration test versions)
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
generate-go-client: build
	@echo "$(YELLOW)Generating Go client code...$(NC)"
	@node generate-go-client.js

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

# Run all tests in one comprehensive task
# This task includes:  setup, anvil start/stop for each test
test-debug:
	@make clean
	@make build
	@echo "$(YELLOW)Starting comprehensive PoSA test suite...$(NC)"
	
	# Step 1: Set up test environment with mock contracts
	@make load-env
	
	# Run Validator Management test with fresh Anvil instance
	@echo "\n$(YELLOW)Running Validator Management tests...$(NC)"
	@make anvil-start
	@sleep 3
	@forge script script/integration/ValidatorLifecycleTest.s.sol:ValidatorLifecycleTest --rpc-url http://localhost:8545 --broadcast --skip-simulation -vvv
	@make anvil-stop
	@make anvil-clean
	
	@echo "\n$(GREEN)✅ All tests completed successfully!$(NC)"

# Run all tests in one comprehensive task
# This task includes:  setup, anvil start/stop for each test
test-all:
	@echo "$(YELLOW)Starting comprehensive PoSA test suite...$(NC)"
	# Step 1: Set up test environment with mock contracts
	@make clean
	@make build
	@make load-env
	
	# Run each test case
	@for test_script in \
		"Deployment tests;script/tools/Deployment.s.sol:DeploymentScript" \
		"Validator Management tests;script/integration/ValidatorManagement.s.sol:ValidatorManagementScript" \
		"Validator Lifecycle tests;script/integration/ValidatorLifecycleTest.s.sol:ValidatorLifecycleTest" \
		"Delegator Lifecycle tests;script/integration/DelegatorLifecycleTest.s.sol:DelegatorLifecycleTest" \
		"Staking Mechanism tests;script/integration/StakingMechanism.s.sol:StakingMechanismScript" \
		"Proposal System tests;script/tools/ProposalSystem.s.sol:ProposalSystemScript" \
		"Punishment Mechanism tests;script/integration/PunishmentMechanism.s.sol:PunishmentMechanismScript" \
		"PoSA Integration tests;script/integration/PoSAIntegrationTest.s.sol:PoSAIntegrationTest"; do \
		TEST_NAME="$$(echo $$test_script | cut -d ';' -f 1)"; \
		SCRIPT="$$(echo $$test_script | cut -d ';' -f 2)"; \
		echo "\n$(YELLOW)Running $$TEST_NAME...$(NC)"; \
		make anvil-start; \
		forge script $$SCRIPT --rpc-url http://localhost:8545 --broadcast --skip-simulation -vvv; \
		make anvil-stop; \
		make anvil-clean; \
	done
	
	@echo "\n$(GREEN)✅ All tests completed successfully!$(NC)"

test-report:
	@echo "$(YELLOW)Starting comprehensive PoSA test suite with report generation...$(NC)"
	@make clean
	@make build
	@make load-env
	# Step 2: Start Anvil test node
	@make anvil-start

	# Step 3: Run test-report to generate comprehensive test results
	@echo "$(YELLOW)Running tests and generating reports...$(NC)"
	@python3 test-report.py --broadcast
	
	# Step 4: Stop Anvil test node
	@make anvil-stop
	
	# Step 5: Clean Anvil logs and temporary files
	@make anvil-clean
	
	@echo "$(GREEN)✅ All tests completed successfully!$(NC)"
	@echo "$(GREEN)Test reports generated in ./test-results directory$(NC)"

# =========================
# New Test Framework
# =========================

# 运行特定测试场景
test-scenario:
	@if [ -z "$(SCENARIO)" ]; then \
		echo "$(RED)Error: Please specify test scenario with SCENARIO=script/ci/your-test.sh$(NC)"; \
		exit 1; \
	fi
	@make test-env-stop
	@echo "$(YELLOW)Running test scenario: $(SCENARIO)$(NC)"
	@make test-env-start
	@bash $(SCENARIO)
	@make test-env-stop

# 运行所有测试场景
test-all-scenarios:
	@echo "$(YELLOW)Running all test scenarios...$(NC)"
	@make test-env-start
	@for scenario in script/ci/*.sh; do \
		echo "\n$(YELLOW)=== Running: $$scenario ===$(NC)"; \
		bash $$scenario; \
		echo "$(YELLOW)=== Completed: $$scenario ===$(NC)"; \
	done
	@make test-env-stop

# 快速运行单个原子脚本
test-atomic:
	@if [ -z "$(SCRIPT)" ]; then \
		echo "$(RED)Error: Please specify atomic script with SCRIPT=script/path/YourScript.s.sol:YourScript$(NC)"; \
		exit 1; \
	fi
	@echo "$(YELLOW)Running atomic script: $(SCRIPT)$(NC)"
	@make test-env-start
	@forge script $(SCRIPT) --rpc-url http://localhost:8545 --broadcast --skip-simulation -vvv
	@make test-env-stop

# 获取系统配置参数
get-system-params:
	@echo "$(YELLOW)System Configuration Parameters$(NC)"
	@echo -n "Epoch Duration: "; cast call 0x000000000000000000000000000000000000f010 "getEpochDuration()(uint256)" | grep -oE '[0-9]+'; echo -n " seconds"
	@echo -n "Unbonding Period: "; cast call 0x000000000000000000000000000000000000F013 "getUnbondingPeriod()(uint256)" | grep -oE '[0-9]+'; echo -n " seconds"
	@echo -n "Validator Unjail Period: "; cast call 0x000000000000000000000000000000000000F013 "getValidatorUnjailPeriod()(uint256)" | grep -oE '[0-9]+'; echo -n " seconds"
	@echo -n "Proposal Lasting Period: "; cast call 0x000000000000000000000000000000000000F012 "getProposalLastingPeriod()(uint256)" | grep -oE '[0-9]+'; echo -n " seconds"

# 启动单个Anvil实例用于所有测试
test-env-start:
	@echo "$(YELLOW)Starting Anvil test node for test suite...$(NC)"
	@./script/ci/anvil-setup.sh --start
	@sleep 3

# 停止测试环境
test-env-stop:
	@echo "$(YELLOW)Stopping Anvil test node...$(NC)"
	@./script/ci/anvil-setup.sh --stop
	@./script/ci/anvil-setup.sh --clean

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
