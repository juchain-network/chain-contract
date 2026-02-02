.PHONY: all ci clean image init run ready reset test stop help logs test-all test-all-legacy

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
	@compose_file="$(PWD)/$(DATA_DIR)/docker-compose.runtime.yml"; \
	if [ ! -f "$$compose_file" ]; then \
		compose_file="$(PWD)/$(DOCKER_DIR)/docker-compose.yml"; \
	fi; \
	cd $(DOCKER_DIR) && COMPOSE_FILE="$$compose_file" $(DOCKER_COMPOSE) up -d
	@bash $(SCRIPTS_DIR)/ensure_miners.sh
	@INCREMENTS_REQUIRED=0 bash $(SCRIPTS_DIR)/wait_for_node.sh http://localhost:18545

# Wait for Node Readiness
ready:
	@INCREMENTS_REQUIRED=0 bash $(SCRIPTS_DIR)/wait_for_node.sh

# Helper to reset environment
reset: clean init run ready

# Punish tests are split with resets to avoid chain stall
test-punish:
	@echo "🧪 Running Punishment Test Group..."
	@$(MAKE) reset && go test ./tests/... -v -run "TestF1_ExitFlow" -config $(PWD)/$(DATA_DIR)/test_config.yaml; $(MAKE) stop
	@$(MAKE) reset && go test ./tests/... -v -run "TestF2_QuickReEntry" -config $(PWD)/$(DATA_DIR)/test_config.yaml; $(MAKE) stop
	@$(MAKE) reset && go test ./tests/... -v -run "TestF3_WithdrawProfits" -config $(PWD)/$(DATA_DIR)/test_config.yaml; $(MAKE) stop
	@$(MAKE) reset && go test ./tests/... -v -run "TestF4_MiscExit" -config $(PWD)/$(DATA_DIR)/test_config.yaml; $(MAKE) stop
	@$(MAKE) reset && go test ./tests/... -v -run "TestF5_RoleChange" -config $(PWD)/$(DATA_DIR)/test_config.yaml; $(MAKE) stop
	@$(MAKE) reset && go test ./tests/... -v -run "TestF6_DoubleSignWindow" -config $(PWD)/$(DATA_DIR)/test_config.yaml; $(MAKE) stop
	@$(MAKE) reset && go test ./tests/... -v -run "TestF7_PunishedRedemption" -config $(PWD)/$(DATA_DIR)/test_config.yaml; $(MAKE) stop
	@$(MAKE) reset && go test ./tests/... -v -run "TestG_DoubleSign" -config $(PWD)/$(DATA_DIR)/test_config.yaml; $(MAKE) stop
	@$(MAKE) reset && go test ./tests/... -v -run "TestG_PunishPaths" -config $(PWD)/$(DATA_DIR)/test_config.yaml; $(MAKE) stop

# Individual Test Groups
test-config:
	@$(MAKE) reset && go test ./tests/... -v -run "TestA_SystemConfigSetup|TestB_ConfigBoundaryChecks" -config $(PWD)/$(DATA_DIR)/test_config.yaml; $(MAKE) stop

test-governance:
	@$(MAKE) reset && go test ./tests/... -v -run "TestB_Governance.*" -config $(PWD)/$(DATA_DIR)/test_config.yaml; $(MAKE) stop

test-staking:
	@$(MAKE) reset && go test ./tests/... -v -run "TestC_Staking.*|TestD_Staking.*" -config $(PWD)/$(DATA_DIR)/test_config.yaml; $(MAKE) stop

test-delegation:
	@$(MAKE) reset && go test ./tests/... -v -run "TestE_Delegation.*" -config $(PWD)/$(DATA_DIR)/test_config.yaml; $(MAKE) stop

test-rewards:
	@$(MAKE) reset && go test ./tests/... -v -run "TestH_Robustness" -config $(PWD)/$(DATA_DIR)/test_config.yaml; $(MAKE) stop
	@$(MAKE) reset && go test ./tests/... -v -run "TestI_ConsensusRewards|TestI_PublicQueryCoverage|TestI_ValidatorExtras" -config $(PWD)/$(DATA_DIR)/test_config.yaml; $(MAKE) stop

test-epoch:
	@$(MAKE) reset && go test ./tests/... -v -run "TestY_UpdateActiveValidatorSet" -config $(PWD)/$(DATA_DIR)/test_config.yaml; $(MAKE) stop
	@$(MAKE) reset && go test ./tests/... -v -run "TestZ_LastManStanding" -config $(PWD)/$(DATA_DIR)/test_config.yaml; $(MAKE) stop
	@$(MAKE) reset && go test ./tests/... -v -run "TestZ_UpgradesAndInitGuards" -config $(PWD)/$(DATA_DIR)/test_config.yaml; $(MAKE) stop

# Run ALL tests serially with isolated reset per test
test-all:
	@bash $(SCRIPTS_DIR)/run_tests_serial.sh

# Legacy grouped runs
test-all-legacy: test-config test-governance test-staking test-delegation test-punish test-rewards test-epoch

# Legacy single-pass test (runs everything in one go without reset)
test: ready
	@echo "🧪 Running Integration Tests (Single Pass)..."
	@go test ./tests/... -v -config $(PWD)/$(DATA_DIR)/test_config.yaml

# View Logs
logs:
	@compose_file="$(PWD)/$(DATA_DIR)/docker-compose.runtime.yml"; \
	if [ ! -f "$$compose_file" ]; then \
		compose_file="$(PWD)/$(DOCKER_DIR)/docker-compose.yml"; \
	fi; \
	cd $(DOCKER_DIR) && COMPOSE_FILE="$$compose_file" $(DOCKER_COMPOSE) logs -f $(NODE)

# Stop Network
stop:
	@echo "🛑 Stopping network..."
	@compose_file="$(PWD)/$(DATA_DIR)/docker-compose.runtime.yml"; \
	if [ ! -f "$$compose_file" ]; then \
		compose_file="$(PWD)/$(DOCKER_DIR)/docker-compose.yml"; \
	fi; \
	cd $(DOCKER_DIR) && COMPOSE_FILE="$$compose_file" $(DOCKER_COMPOSE) down --remove-orphans

# Clean Everything
clean: stop
	@echo "🧹 Cleaning data directory..."
	@if [ -d "$(DATA_DIR)" ]; then \
		docker run --rm -v $(PWD):/work alpine sh -c "rm -rf /work/$(DATA_DIR)"; \
	fi
	@docker network prune -f
	@docker image prune -f
	@docker volume prune -f
	@echo "✅ Clean complete."

# CI
ci: image test-all
