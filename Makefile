# Makefile for sys-contract project
# Usage: make [target]

# Variables
ANVIL_PORT = 8545
ANVIL_CHAIN_ID = 31337
PRIVATE_KEY = 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
RPC_URL = http://localhost:$(ANVIL_PORT)
ABIGEN = ../chain/build/bin/abigen
GO_CLIENT_DIR = tools/contracts

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

.PHONY: help clean build test anvil deploy scripts test-scripts all check-anvil check-contracts stop-anvil reset-anvil deploy-chain deploy-chain-local deploy-mainnet deploy-sepolia deploy-bsc deploy-bsc-testnet deploy-polygon deploy-mumbai verify-contract generate-go-client

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
	@echo "  $(GREEN)generate-go-client$(NC) - Generate Go client code using abigen"
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
	@echo "  $(GREEN)generate-contracts$(NC) - Generate production contracts from templates"
	@echo "  $(GREEN)generate-contracts-mock$(NC) - Generate mock contracts for testing"
	@echo "  $(GREEN)fmt$(NC)              - Format Solidity code"
	@echo "  $(GREEN)security$(NC)         - Run security analysis (requires slither)"
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

# ======================================
# Chain Deployment Targets
# ======================================

# Deploy to real network (auto-staking will fail gracefully, validators register manually)
deploy-chain:
	@echo "$(YELLOW)Deploying to REAL NETWORK...$(NC)"
	@echo "$(GREEN)RPC URL: $(CHAIN_RPC_URL)$(NC)"
	@echo "$(YELLOW)Note: Auto-staking will fail on real networks - validators must register manually$(NC)"
	@if [ "$(CHAIN_PRIVATE_KEY)" = "$(PRIVATE_KEY)" ]; then \
		echo "$(RED)Warning: Using default private key. Set CHAIN_PRIVATE_KEY environment variable.$(NC)"; \
	fi
	@forge script script/DeployToChain.s.sol:DeployToChainScript \
		--rpc-url $(CHAIN_RPC_URL) \
		--broadcast \
		--private-key $(CHAIN_PRIVATE_KEY) \
		--gas-price 2000000000 \
		--gas-limit 15000000 \
		--legacy \
		-vv

# Single validator registration (for individual validators)
validator-register:
	@echo "$(YELLOW)Registering individual validator...$(NC)"
	@echo "$(GREEN)Staking Contract: $(STAKING_CONTRACT)$(NC)"
	@echo "$(GREEN)Validator: $(shell forge wallet address --private-key $(VALIDATOR_PRIVATE_KEY))$(NC)"
	@VALIDATOR_PRIVATE_KEY=$(VALIDATOR_PRIVATE_KEY) \
	 STAKING_CONTRACT=$(STAKING_CONTRACT) \
	 COMMISSION_RATE=$(COMMISSION_RATE) \
	 forge script script/ValidatorStake.s.sol:ValidatorStakeScript \
		--rpc-url $(CHAIN_RPC_URL) \
		--broadcast \
		--gas-price 2000000000 \
		--gas-limit 15000000 \
		--legacy \
		-vv

# Batch validator registration (if you control multiple validator keys)
validators-register-batch:
	@echo "$(YELLOW)Batch registering validators...$(NC)"
	@echo "$(GREEN)Staking Contract: $(STAKING_CONTRACT)$(NC)"
	@echo "$(RED)Warning: Only use if you control all validator private keys$(NC)"
	@STAKING_CONTRACT=$(STAKING_CONTRACT) \
	 forge script script/BatchValidatorStake.s.sol:BatchValidatorStakeScript \
		--rpc-url $(CHAIN_RPC_URL) \
		--broadcast \
		--gas-price 2000000000 \
		--gas-limit 15000000 \
		--legacy \
		-vv


# Deploy to JuChain mainnet
deploy-juchain:
	@echo "$(YELLOW)Deploying to JuChain Mainnet...$(NC)"
	@echo "$(RED)WARNING: This will deploy to JUCHAIN MAINNET. Are you sure? [y/N]$(NC)" && read ans && [ $${ans:-N} = y ]
	@make deploy-chain CHAIN_RPC_URL=$(JUCHAIN_MAINNET_RPC_URL) CHAIN_ID=210000

# Deploy to JuChain testnet
deploy-juchain-testnet:
	@echo "$(YELLOW)Deploying to JuChain Testnet...$(NC)"
	@make deploy-chain CHAIN_RPC_URL=$(JUCHAIN_TESTNET_RPC_URL) CHAIN_ID=202599

# Deploy to local test chain (JuChain 202599) with auto-staking for testing
deploy-chain-local:
	@echo "$(YELLOW)Deploying contracts to JuChain (202599) with auto-staking...$(NC)"
	@echo "$(GREEN)RPC URL: http://localhost:8545$(NC)"
	@echo "$(GREEN)Chain ID: 202599$(NC)"
	@PRIVATE_KEY=$(PRIVATE_KEY) forge script script/DeployToChain.s.sol:DeployToChainScript \
		--rpc-url http://localhost:8545 \
		--broadcast \
		--gas-price 2000000000 \
		--gas-limit 15000000 \
		--legacy \
		-vv

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
