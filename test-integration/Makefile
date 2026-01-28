.PHONY: all ci clean image init run ready test stop help logs

# Current directory
PWD := $(shell pwd)
# Scripts directory
SCRIPTS_DIR := scripts
# Data directory
DATA_DIR := data
# Docker directory
DOCKER_DIR := docker

# Detect docker compose command (v2 plugin or legacy standalone)
DOCKER_COMPOSE := $(shell if docker compose version >/dev/null 2>&1; then echo "docker compose"; else echo "docker-compose"; fi)

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  image   - Build the Docker image (juchain-node:latest)"
	@echo "  init    - Generate genesis, keys, and network config"
	@echo "  run     - Start the network cluster (background)"
	@echo "  ready   - Wait for Node 0 to be ready (RPC check)"
	@echo "  test    - Run Go integration tests"
	@echo "  stop    - Stop the network cluster"
	@echo "  clean   - Stop network and remove ALL data"
	@echo "  logs    - View cluster logs (Use NODE=node0 for specific node)"
	@echo "  test-all- Run all test groups sequentially with isolation"
	@echo "  ci      - Run full CI flow (build + test-all)"

# Build Docker Image
image:
	@echo "🚀 Building Docker image..."
	@bash $(SCRIPTS_DIR)/build_docker.sh

# Initialize Configuration (Genesis, Keys, Configs)
init:
	@echo "⚙️  Initializing network configuration..."
	@bash $(SCRIPTS_DIR)/gen_network_config.sh

# Start Network
run:
	@echo "🚀 Starting network cluster..."
	@cd $(DOCKER_DIR) && $(DOCKER_COMPOSE) up -d

# Wait for Node Readiness
ready:
	@bash $(SCRIPTS_DIR)/wait_for_node.sh

# Helper to reset environment
reset: clean init run ready

# Individual Test Groups
test-config: reset
	@echo "🧪 Running Config Tests..."
	@go test ./tests/... -v -run "TestA_.*|TestB_Config.*" -config $(PWD)/$(DATA_DIR)/test_config.yaml
	@$(MAKE) stop

test-governance: reset
	@echo "🧪 Running Governance Tests..."
	@go test ./tests/... -v -run "TestB_Governance.*" -config $(PWD)/$(DATA_DIR)/test_config.yaml
	@$(MAKE) stop

test-staking: reset
	@echo "🧪 Running Staking Tests..."
	@go test ./tests/... -v -run "TestC_Staking.*|TestD_Staking.*" -config $(PWD)/$(DATA_DIR)/test_config.yaml
	@$(MAKE) stop

test-delegation: reset
	@echo "🧪 Running Delegation Tests..."
	@go test ./tests/... -v -run "TestE_Delegation.*" -config $(PWD)/$(DATA_DIR)/test_config.yaml
	@$(MAKE) stop

test-punish: reset
	@echo "🧪 Running Punishment Tests..."
	@go test ./tests/... -v -run "TestF_Punish.*|TestG_.*" -config $(PWD)/$(DATA_DIR)/test_config.yaml
	@$(MAKE) stop

test-rewards: reset
	@echo "🧪 Running Rewards & Robustness Tests..."
	@go test ./tests/... -v -run "TestH_.*|TestI_.*" -config $(PWD)/$(DATA_DIR)/test_config.yaml
	@$(MAKE) stop

test-epoch: reset
	@echo "🧪 Running Epoch & Upgrade Tests..."
	@go test ./tests/... -v -run "TestY_.*|TestZ_.*" -config $(PWD)/$(DATA_DIR)/test_config.yaml
	@$(MAKE) stop

# Run ALL Groups
test-all: test-config test-governance test-staking test-delegation test-punish test-rewards test-epoch

# Legacy single-pass test (runs everything in one go without reset)
test: ready
	@echo "🧪 Running Integration Tests (Single Pass)..."
	@go test ./tests/... -v -config $(PWD)/$(DATA_DIR)/test_config.yaml

# View Logs
# Usage: make logs
# Usage: make logs NODE=node0
logs:
	@cd $(DOCKER_DIR) && $(DOCKER_COMPOSE) logs -f $(NODE)

# Stop Network (Keep Data)
stop:
	@echo "🛑 Stopping network..."
	@cd $(DOCKER_DIR) && $(DOCKER_COMPOSE) down

# Clean Everything
clean: stop
	@echo "🧹 Cleaning data directory..."
	@sudo rm -rf $(DATA_DIR)
	@echo "✅ Clean complete."

# CI / All-in-one flow
ci: image test-all
