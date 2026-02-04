.PHONY: all ci clean image init run ready reset test stop help logs test-all test-all-legacy ci-tool ci-groups ci-tests

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

CI_TOOL := go run ./ci.go
CI_COMMON_FLAGS := $(if $(DEBUG),-debug,) $(if $(GOCACHE),-gocache $(GOCACHE),) $(if $(CONFIG),-config $(CONFIG),) $(if $(REPORT_DIR),-report-dir $(REPORT_DIR),)

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
	@echo "  ci-tool - Run ci.go with custom ARGS (see ci.go flags)"
	@echo "  ci-groups - Grouped runs via ci.go (set GROUPS, CI_LOG=1)"
	@echo "  ci-tests - Per-test runs via ci.go (set TESTS or RUN)"

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
	@$(CI_TOOL) -mode tests $(CI_COMMON_FLAGS) -tests TestF1_ExitFlow,TestF2_QuickReEntry,TestF3_WithdrawProfits,TestF4_MiscExit,TestF5_RoleChange,TestF6_DoubleSignWindow,TestF7_PunishedRedemption,TestG_DoubleSign,TestG_PunishPaths

# Individual Test Groups
test-config:
	@$(CI_TOOL) -mode tests $(CI_COMMON_FLAGS) -run "TestA_SystemConfigSetup|TestB_ConfigBoundaryChecks"

test-governance:
	@$(CI_TOOL) -mode tests $(CI_COMMON_FLAGS) -run "TestB_Governance.*"

test-staking:
	@$(CI_TOOL) -mode tests $(CI_COMMON_FLAGS) -run "TestC_Staking.*|TestD_Staking.*"

test-delegation:
	@$(CI_TOOL) -mode tests $(CI_COMMON_FLAGS) -run "TestE_Delegation.*"

test-rewards:
	@$(CI_TOOL) -mode tests $(CI_COMMON_FLAGS) -run "TestH_Robustness"
	@$(CI_TOOL) -mode tests $(CI_COMMON_FLAGS) -run "TestI_ConsensusRewards|TestI_PublicQueryCoverage|TestI_ValidatorExtras"

test-epoch:
	@$(CI_TOOL) -mode tests $(CI_COMMON_FLAGS) -run "TestY_UpdateActiveValidatorSet"
	@$(CI_TOOL) -mode tests $(CI_COMMON_FLAGS) -run "TestZ_LastManStanding"
	@$(CI_TOOL) -mode tests $(CI_COMMON_FLAGS) -run "TestZ_UpgradesAndInitGuards"

# Run ALL tests serially with isolated reset per test
test-all:
	@$(CI_TOOL) -mode all $(CI_COMMON_FLAGS)

# Legacy grouped runs
test-all-legacy:
	@$(CI_TOOL) -mode groups $(CI_COMMON_FLAGS) -groups config,governance,staking,delegation,punish,rewards,epoch

# Legacy single-pass test (runs everything in one go without reset)
test: ready
	@echo "🧪 Running Integration Tests (Single Pass)..."
	@$(CI_TOOL) -mode tests $(CI_COMMON_FLAGS) -run "." -skip-setup

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

# ci.go wrapper targets (does not remove legacy targets)
ci-tool:
	@$(CI_TOOL) $(CI_COMMON_FLAGS) $(ARGS)

ci-groups:
	@$(CI_TOOL) -mode groups $(CI_COMMON_FLAGS) $(if $(GROUPS),-groups $(GROUPS),) $(if $(CI_LOG),-ci-log,)

ci-tests:
	@if [ -z "$(TESTS)" ] && [ -z "$(RUN)" ]; then echo "Set TESTS or RUN"; exit 1; fi
	@$(CI_TOOL) -mode tests $(CI_COMMON_FLAGS) $(if $(TESTS),-tests $(TESTS),) $(if $(RUN),-run $(RUN),) $(if $(TIMEOUT),-timeout $(TIMEOUT),)
