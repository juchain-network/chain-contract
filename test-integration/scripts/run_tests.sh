#!/bin/bash
set -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
TEST_INT_DIR="$SCRIPT_DIR/.."
DOCKER_DIR="$TEST_INT_DIR/docker"

# Help message
if [[ "$1" == "--help" ]]; then
    echo "Usage: ./run_tests.sh [options]"
    echo "Options:"
    echo "  --build       Force rebuild of docker image"
    echo "  --keep        Keep environment running after tests"
    exit 0
fi

# 1. Build Docker if requested or not exists
if [[ "$1" == "--build" ]] || [[ "$(docker images -q juchain-node:latest 2> /dev/null)" == "" ]]; then
    echo "🚀 Building Docker environment..."
    "$SCRIPT_DIR/build_docker.sh"
fi

# 2. Cleanup old environment
echo "🧹 Cleaning up old environment..."
pushd "$DOCKER_DIR" > /dev/null
docker-compose down -v
popd > /dev/null

# 3. Generate Configuration
echo "⚙️  Generating network configuration..."
"$SCRIPT_DIR/gen_network_config.sh"

# 4. Start Network
echo "🚀 Starting network cluster..."
pushd "$DOCKER_DIR" > /dev/null
docker-compose up -d
popd > /dev/null

# 5. Wait for Node 0
echo "⏳ Waiting for Node 0 to be ready..."
RETRIES=30
while [ $RETRIES -gt 0 ]; do
    if curl -s -X POST -H "Content-Type: application/json" --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' http://localhost:8545 > /dev/null; then
        echo "✅ Node 0 is up!"
        break
    fi
    sleep 1
    RETRIES=$((RETRIES-1))
    echo -n "."
done

if [ $RETRIES -eq 0 ]; then
    echo "❌ Timeout waiting for Node 0"
    docker-compose -f "$DOCKER_DIR/docker-compose.yml" logs
    exit 1
fi

# 6. Run Tests
echo "🧪 Running Integration Tests..."
pushd "$TEST_INT_DIR" > /dev/null
# Pass the generated config file path
go test ./tests/... -v -config "$TEST_INT_DIR/data/test_config.yaml"
TEST_EXIT_CODE=$?
popd > /dev/null

# 7. Cleanup (unless --keep)
if [[ "$1" != "--keep" ]] && [[ "$2" != "--keep" ]]; then
    echo "🧹 Shutting down network..."
    pushd "$DOCKER_DIR" > /dev/null
    docker-compose down
    popd > /dev/null
else
    echo "ℹ️  Environment kept running. To stop: cd test-integration/docker && docker-compose down"
fi

exit $TEST_EXIT_CODE
