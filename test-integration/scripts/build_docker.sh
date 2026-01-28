#!/bin/bash
set -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PROJECT_ROOT="$SCRIPT_DIR/../../.."
CHAIN_ROOT="$PROJECT_ROOT/chain"
DOCKER_DIR="$SCRIPT_DIR/../docker"

echo "=== Building geth binary from $CHAIN_ROOT ==="

# Compile geth using make
pushd "$CHAIN_ROOT"
make geth
popd

# Copy binary (renaming to juchain to match Dockerfile/start.sh expectations)
echo "=== Copying binary to Docker context ==="
if [ -f "$CHAIN_ROOT/build/bin/geth" ]; then
    cp "$CHAIN_ROOT/build/bin/geth" "$DOCKER_DIR/juchain"
else
    echo "❌ Error: $CHAIN_ROOT/build/bin/geth not found. Build failed?"
    exit 1
fi

echo "=== Building Docker image ==="
pushd "$DOCKER_DIR"
docker build -t juchain-node:latest .
rm juchain  # Cleanup binary after build
popd

echo "✅ Docker image 'juchain-node:latest' built successfully."