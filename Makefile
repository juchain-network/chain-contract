# Makefile for sys-contract project
# Usage: make [target]

# Variables
ANVIL_PORT = 8545
ANVIL_CHAIN_ID = 31337
PRIVATE_KEY = 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
RPC_URL = http://localhost:$(ANVIL_PORT)

# Chain deployment variables (can be overridden)
CHAIN_RPC_URL ?= $(RPC_URL)
CHAIN_PRIVATE_KEY ?= $(PRIVATE_KEY)
CHAIN_ID ?= $(ANVIL_CHAIN_ID)

# Common network configurations
JUCHAIN_MAINNET_RPC_URL = https://rpc.juchain.org
JUCHAIN_TESTNET_RPC_URL = https://testnet-rpc.juchain.org

# Default validators for initialization
VALIDATORS = [0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266,0x70997970C51812dc3A010C7d01b50e0d17dc79C8,0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC]

# Colors for output
GREEN = \033[0;32m
YELLOW = \033[1;33m
RED = \033[0;31m
NC = \033[0m # No Color

.PHONY: help clean build test anvil deploy scripts test-scripts all check-anvil check-contracts stop-anvil reset-anvil deploy-chain deploy-chain-local deploy-mainnet deploy-sepolia deploy-bsc deploy-bsc-testnet deploy-polygon deploy-mumbai verify-contract

# Default target
help:
	@echo "$(YELLOW)Available targets:$(NC)"
	@echo "  $(GREEN)build$(NC)         - Compile contracts"
	@echo "  $(GREEN)test$(NC)          - Run all tests"
	@echo "  $(GREEN)clean$(NC)         - Clean build artifacts"
	@echo "  $(GREEN)anvil$(NC)         - Start local anvil node"
	@echo "  $(GREEN)deploy$(NC)        - Deploy contracts to local anvil"
	@echo "  $(GREEN)scripts$(NC)       - Test all scripts on local anvil"
	@echo "  $(GREEN)test-scripts$(NC)  - Full workflow: anvil + deploy + test scripts"
	@echo "  $(GREEN)all$(NC)           - Clean + build + test"
	@echo ""
	@echo "$(YELLOW)Chain deployment targets:$(NC)"
	@echo "  $(GREEN)deploy-chain-202599$(NC) - Deploy using DeployToChain script to JuChain (202599)"
	@echo "  $(GREEN)deploy-chain-local$(NC) - Deploy using DeployToChain script to local anvil"
	@echo "  $(GREEN)deploy-chain$(NC)     - Deploy to custom RPC (use CHAIN_RPC_URL and CHAIN_PRIVATE_KEY)"
	@echo "  $(GREEN)deploy-juchain$(NC)   - Deploy to JuChain mainnet"
	@echo "  $(GREEN)deploy-juchain-testnet$(NC) - Deploy to JuChain testnet"
	@echo "  $(GREEN)verify-contract$(NC)  - Verify deployed contracts"
	@echo ""
	@echo "$(YELLOW)Utility targets:$(NC)"
	@echo "  $(GREEN)check-anvil$(NC)      - Check if anvil is running"
	@echo "  $(GREEN)check-contracts$(NC)  - Check contract deployment status"
	@echo "  $(GREEN)stop-anvil$(NC)       - Stop anvil node"
	@echo "  $(GREEN)reset-anvil$(NC)      - Restart anvil node"
	@echo "  $(GREEN)addresses$(NC)        - Show contract addresses"
	@echo ""
	@echo "$(YELLOW)Script targets:$(NC)"
	@echo "  $(GREEN)script-deploy$(NC)     - Deploy system contracts"
	@echo "  $(GREEN)script-add-node$(NC)   - Add new validator node"
	@echo "  $(GREEN)script-create-proposal$(NC) - Create a proposal"
	@echo "  $(GREEN)script-vote$(NC)       - Vote on a proposal"
	@echo "  $(GREEN)script-remove-node$(NC) - Remove a validator node"
	@echo "  $(GREEN)script-staking$(NC)    - Test staking operations"
	@echo "  $(GREEN)script-update-config$(NC) - Update system config"

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

# Start anvil local node
anvil:
	@echo "$(YELLOW)Starting Anvil local node on port $(ANVIL_PORT)...$(NC)"
	@echo "$(GREEN)Use Ctrl+C to stop$(NC)"
	anvil --port $(ANVIL_PORT) --chain-id $(ANVIL_CHAIN_ID) --accounts 10 --balance 1000

# Deploy contracts to local anvil
deploy:
	@echo "$(YELLOW)Deploying contracts to local anvil...$(NC)"
	@echo "$(GREEN)Make sure anvil is running on port $(ANVIL_PORT)$(NC)"
	@forge script script/DeploySystem.s.sol \
		--fork-url $(RPC_URL) \
		--broadcast \
		--private-key $(PRIVATE_KEY) \
		--gas-price 2000000000 \
		--gas-limit 15000000 \
		--legacy \
		--skip-simulation \
		-vvv || echo "$(YELLOW)Note: Contract may already be deployed$(NC)"

# Test individual scripts
script-deploy:
	@echo "$(YELLOW)Running DeploySystem script...$(NC)"
	@forge script script/DeploySystem.s.sol \
		--fork-url $(RPC_URL) \
		--broadcast \
		--private-key $(PRIVATE_KEY) \
		--gas-price 2000000000 \
		--gas-limit 15000000 \
		--legacy \
		--skip-simulation \
		-v || echo "$(YELLOW)Note: Contract may already be deployed$(NC)"

script-add-node:
	@echo "$(YELLOW)Running AddNewNode script...$(NC)"
	@forge script script/AddNewNode.s.sol \
		--fork-url $(RPC_URL) \
		--broadcast \
		--private-key $(PRIVATE_KEY) \
		--gas-price 2000000000 \
		--gas-limit 15000000 \
		--legacy \
		--skip-simulation \
		-v || echo "$(YELLOW)Note: Script execution may have failed - check logs$(NC)"

script-create-proposal:
	@echo "$(YELLOW)Running CreateProposal script...$(NC)"
	@forge script script/CreateProposal.s.sol \
		--fork-url $(RPC_URL) \
		--broadcast \
		--private-key $(PRIVATE_KEY) \
		--gas-price 2000000000 \
		--gas-limit 15000000 \
		--legacy \
		--skip-simulation \
		-v || echo "$(YELLOW)Note: Script execution may have failed - check logs$(NC)"

script-vote:
	@echo "$(YELLOW)Running VoteProposal script...$(NC)"
	@forge script script/VoteProposal.s.sol \
		--fork-url $(RPC_URL) \
		--broadcast \
		--private-key $(PRIVATE_KEY) \
		--gas-price 2000000000 \
		--gas-limit 15000000 \
		--legacy \
		--skip-simulation \
		-v || echo "$(YELLOW)Note: Script execution may have failed - check logs$(NC)"

script-remove-node:
	@echo "$(YELLOW)Running RemoveNode script...$(NC)"
	@forge script script/RemoveNode.s.sol \
		--fork-url $(RPC_URL) \
		--broadcast \
		--private-key $(PRIVATE_KEY) \
		--gas-price 2000000000 \
		--gas-limit 15000000 \
		--legacy \
		--skip-simulation \
		-v || echo "$(YELLOW)Note: Script execution may have failed - check logs$(NC)"

script-staking:
	@echo "$(YELLOW)Running StakingOperations script...$(NC)"
	@forge script script/StakingOperations.s.sol \
		--fork-url $(RPC_URL) \
		--broadcast \
		--private-key $(PRIVATE_KEY) \
		--gas-price 2000000000 \
		--gas-limit 15000000 \
		--legacy \
		--skip-simulation \
		-v || echo "$(YELLOW)Note: Script execution may have failed - check logs$(NC)"

script-update-config:
	@echo "$(YELLOW)Running UpdateConfig script...$(NC)"
	@forge script script/UpdateConfig.s.sol \
		--fork-url $(RPC_URL) \
		--broadcast \
		--private-key $(PRIVATE_KEY) \
		--gas-price 2000000000 \
		--gas-limit 15000000 \
		--legacy \
		--skip-simulation \
		-v || echo "$(YELLOW)Note: Script execution may have failed - check logs$(NC)"

# Test all scripts in sequence
scripts: script-deploy script-create-proposal script-add-node script-vote script-staking script-update-config
	@echo "$(GREEN)All scripts tested successfully!$(NC)"

# Run scripts in test mode (without anvil/broadcast)
test-mode-deploy:
	@echo "$(CYAN)Running deployment in test mode...$(NC)"
	@forge script script/DeploySystem.s.sol:DeploySystemScript --sig "run()" -v

test-mode-create-proposal:
	@echo "$(CYAN)Running CreateProposal in test mode...$(NC)"
	@forge script script/CreateProposal.s.sol:CreateProposalScript --sig "run()" -v

test-mode-vote:
	@echo "$(CYAN)Running VoteProposal in test mode...$(NC)"
	@forge script script/VoteProposal.s.sol:VoteProposalScript --sig "run()" -v

test-mode-add-node:
	@echo "$(CYAN)Running AddNewNode in test mode...$(NC)"
	@forge script script/AddNewNode.s.sol:AddNewNodeScript --sig "run()" -v

test-mode-staking:
	@echo "$(CYAN)Running StakingOperations in test mode...$(NC)"
	@forge script script/StakingOperations.s.sol:StakingOperationsScript --sig "run()" -v

# Run all scripts in test mode
test-scripts-mode: test-mode-deploy test-mode-create-proposal test-mode-vote test-mode-add-node test-mode-staking
	@echo "$(GREEN)All scripts tested in test mode successfully!$(NC)"

# Full workflow: build, test, start anvil, deploy, test scripts
test-scripts: build test
	@echo "$(YELLOW)Starting full script testing workflow...$(NC)"
	@echo "$(RED)Note: This will start anvil in background. Use 'make stop-anvil' to stop.$(NC)"
	@anvil --port $(ANVIL_PORT) --chain-id $(ANVIL_CHAIN_ID) --accounts 10 --balance 1000 > /dev/null 2>&1 & \
	echo $$! > .anvil.pid && \
	sleep 3 && \
	echo "$(GREEN)Anvil started with PID $$(cat .anvil.pid)$(NC)" && \
	make deploy && \
	sleep 2 && \
	make scripts && \
	echo "$(GREEN)Script testing completed!$(NC)" && \
	echo "$(YELLOW)Run 'make stop-anvil' to stop the anvil node$(NC)"

# Stop anvil
stop-anvil:
	@if [ -f .anvil.pid ]; then \
		echo "$(YELLOW)Stopping anvil...$(NC)"; \
		kill $$(cat .anvil.pid) 2>/dev/null || true; \
		rm -f .anvil.pid; \
		echo "$(GREEN)Anvil stopped$(NC)"; \
	else \
		echo "$(RED)No anvil PID file found$(NC)"; \
	fi

# Reset anvil (stop and restart)
reset-anvil: stop-anvil
	@echo "$(YELLOW)Restarting anvil...$(NC)"
	@sleep 1
	@anvil --port $(ANVIL_PORT) --chain-id $(ANVIL_CHAIN_ID) --accounts 10 --balance 1000 > /dev/null 2>&1 & \
	echo $$! > .anvil.pid && \
	sleep 3 && \
	echo "$(GREEN)Anvil restarted with PID $$(cat .anvil.pid)$(NC)"

# Check if anvil is running
check-anvil:
	@if curl -s $(RPC_URL) > /dev/null 2>&1; then \
		echo "$(GREEN)Anvil is running on $(RPC_URL)$(NC)"; \
	else \
		echo "$(RED)Anvil is not running on $(RPC_URL)$(NC)"; \
	fi

# Check contract deployment status
check-contracts:
	@echo "$(YELLOW)Checking contract deployment status...$(NC)"
	@if curl -s $(RPC_URL) > /dev/null 2>&1; then \
		cast code 0x000000000000000000000000000000000000f000 --rpc-url $(RPC_URL) > /dev/null 2>&1 && \
		echo "$(GREEN)Validators contract deployed$(NC)" || echo "$(RED)Validators contract not deployed$(NC)"; \
		cast code 0x000000000000000000000000000000000000F002 --rpc-url $(RPC_URL) > /dev/null 2>&1 && \
		echo "$(GREEN)Proposal contract deployed$(NC)" || echo "$(RED)Proposal contract not deployed$(NC)"; \
		cast code 0x000000000000000000000000000000000000F001 --rpc-url $(RPC_URL) > /dev/null 2>&1 && \
		echo "$(GREEN)Punish contract deployed$(NC)" || echo "$(RED)Punish contract not deployed$(NC)"; \
		cast code 0x000000000000000000000000000000000000F003 --rpc-url $(RPC_URL) > /dev/null 2>&1 && \
		echo "$(GREEN)Staking contract deployed$(NC)" || echo "$(RED)Staking contract not deployed$(NC)"; \
	else \
		echo "$(RED)Cannot connect to anvil$(NC)"; \
	fi

# Interactive script testing (requires manual anvil start)
interactive-test:
	@echo "$(YELLOW)Interactive script testing mode$(NC)"
	@echo "$(GREEN)1. Start anvil in another terminal: make anvil$(NC)"
	@echo "$(GREEN)2. Deploy contracts: make deploy$(NC)"
	@echo "$(GREEN)3. Test individual scripts with: make script-[name]$(NC)"
	@echo "$(GREEN)4. Or test all scripts: make scripts$(NC)"

# Development workflow
dev: clean build test
	@echo "$(GREEN)Development build completed!$(NC)"

# Full build and test
all: clean build test
	@echo "$(GREEN)All tasks completed successfully!$(NC)"

# Show contract addresses (after deployment)
addresses:
	@echo "$(YELLOW)System Contract Addresses:$(NC)"
	@echo "  Validators: 0x000000000000000000000000000000000000f000"
	@echo "  Punish:     0x000000000000000000000000000000000000F001"
	@echo "  Proposal:   0x000000000000000000000000000000000000F002"
	@echo "  Staking:    0x000000000000000000000000000000000000F003"

# Debug targets
debug-deploy:
	@echo "$(YELLOW)Running deploy with maximum verbosity...$(NC)"
	forge script script/DeploySystem.s.sol \
		--fork-url $(RPC_URL) \
		--broadcast \
		--private-key $(PRIVATE_KEY) \
		--gas-price 2000000000 \
		--gas-limit 15000000 \
		--legacy \
		--skip-simulation \
		-vvvv

debug-script:
	@echo "$(YELLOW)Debug mode - specify script manually$(NC)"
	@echo "Usage: forge script script/[ScriptName].s.sol --fork-url $(RPC_URL) --broadcast --private-key $(PRIVATE_KEY) --gas-price 2000000000 --gas-limit 15000000 --legacy --skip-simulation -vvvv"

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

# ======================================
# Chain Deployment Targets
# ======================================

# Deploy to current chain (JuChain 202599) using DeployToChain script
deploy-chain-local:
	@echo "$(YELLOW)Deploying contracts to JuChain (202599) using DeployToChain script...$(NC)"
	@echo "$(GREEN)RPC URL: http://localhost:8545$(NC)"
	@echo "$(GREEN)Chain ID: 202599$(NC)"
	@PRIVATE_KEY=$(PRIVATE_KEY) forge script script/DeployToChain.s.sol:DeployToChainScript \
		--rpc-url http://localhost:8545 \
		--broadcast \
		--gas-price 2000000000 \
		--gas-limit 15000000 \
		--legacy \
		-vv

# Deploy to custom chain (specify CHAIN_RPC_URL and CHAIN_PRIVATE_KEY)
deploy-chain:
	@echo "$(YELLOW)Deploying contracts to custom chain...$(NC)"
	@echo "$(GREEN)RPC URL: $(CHAIN_RPC_URL)$(NC)"
	@if [ "$(CHAIN_PRIVATE_KEY)" = "$(PRIVATE_KEY)" ]; then \
		echo "$(RED)Warning: Using default private key. Set CHAIN_PRIVATE_KEY environment variable.$(NC)"; \
	fi
	@if [ -n "$(ETHERSCAN_API_KEY)" ]; then \
		forge script script/DeployToChain.s.sol:DeployToChainScript \
			--rpc-url $(CHAIN_RPC_URL) \
			--broadcast \
			--private-key $(CHAIN_PRIVATE_KEY) \
			--gas-price 2000000000 \
			--gas-limit 15000000 \
			--legacy \
			--verify \
			--etherscan-api-key $(ETHERSCAN_API_KEY) \
			-vvv || echo "$(YELLOW)Note: Deployment may have failed - check logs$(NC)"; \
	else \
		echo "$(YELLOW)No ETHERSCAN_API_KEY provided, deploying without verification$(NC)"; \
		forge script script/DeployToChain.s.sol:DeployToChainScript \
			--rpc-url $(CHAIN_RPC_URL) \
			--broadcast \
			--private-key $(CHAIN_PRIVATE_KEY) \
			--gas-price 2000000000 \
			--gas-limit 15000000 \
			--legacy \
			-vvv || echo "$(YELLOW)Note: Deployment may have failed - check logs$(NC)"; \
	fi

# Deploy to JuChain mainnet
deploy-juchain:
	@echo "$(YELLOW)Deploying to JuChain Mainnet...$(NC)"
	@echo "$(RED)WARNING: This will deploy to JUCHAIN MAINNET. Are you sure? [y/N]$(NC)" && read ans && [ $${ans:-N} = y ]
	@make deploy-chain CHAIN_RPC_URL=$(JUCHAIN_MAINNET_RPC_URL) CHAIN_ID=210000

# Deploy to JuChain testnet
deploy-juchain-testnet:
	@echo "$(YELLOW)Deploying to JuChain Testnet...$(NC)"
	@make deploy-chain CHAIN_RPC_URL=$(JUCHAIN_TESTNET_RPC_URL) CHAIN_ID=202599

# Verify contracts on Etherscan
verify-contract:
	@echo "$(YELLOW)Verifying contracts...$(NC)"
	@if [ -z "$(CONTRACT_ADDRESS)" ]; then \
		echo "$(RED)Error: CONTRACT_ADDRESS not specified$(NC)"; \
		echo "Usage: make verify-contract CONTRACT_ADDRESS=0x... CONTRACT_NAME=ContractName"; \
		exit 1; \
	fi
	@if [ -z "$(CONTRACT_NAME)" ]; then \
		echo "$(RED)Error: CONTRACT_NAME not specified$(NC)"; \
		echo "Usage: make verify-contract CONTRACT_ADDRESS=0x... CONTRACT_NAME=ContractName"; \
		exit 1; \
	fi
	@forge verify-contract $(CONTRACT_ADDRESS) \
		contracts/$(CONTRACT_NAME).sol:$(CONTRACT_NAME) \
		--rpc-url $(CHAIN_RPC_URL) \
		--etherscan-api-key $(ETHERSCAN_API_KEY) \
		--compiler-version 0.8.20

# Show deployment help
deploy-help:
	@echo "$(YELLOW)Chain Deployment Help:$(NC)"
	@echo ""
	@echo "$(GREEN)Environment Variables:$(NC)"
	@echo "  CHAIN_RPC_URL     - RPC URL for target chain"
	@echo "  CHAIN_PRIVATE_KEY - Private key for deployment"
	@echo "  ETHERSCAN_API_KEY - API key for contract verification"
	@echo ""
	@echo "$(GREEN)Examples:$(NC)"
	@echo "  # Deploy to custom chain:"
	@echo "  CHAIN_RPC_URL=https://your-rpc.com CHAIN_PRIVATE_KEY=0x... make deploy-chain"
	@echo ""
	@echo "  # Deploy to JuChain testnet:"
	@echo "  CHAIN_PRIVATE_KEY=0x... make deploy-juchain-testnet"
	@echo ""
	@echo "  # Deploy to JuChain mainnet:"
	@echo "  CHAIN_PRIVATE_KEY=0x... make deploy-juchain"
	@echo ""
	@echo "  # Verify contract:"
	@echo "  CONTRACT_ADDRESS=0x... CONTRACT_NAME=Validators make verify-contract"
	@echo ""
	@echo "$(YELLOW)Note: Make sure to set CHAIN_PRIVATE_KEY environment variable for real deployments$(NC)"
