#!/bin/bash
set -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PROJECT_ROOT="$SCRIPT_DIR/../../.."
CHAIN_ROOT="$PROJECT_ROOT/chain"
DOCKER_DIR="$SCRIPT_DIR/../docker"

echo "=== Building geth binary from $CHAIN_ROOT ==="

# Compile geth into a writable temp path to avoid permission issues in chain/build.
TMP_BUILD_OUT="${TMPDIR:-/tmp}/juchain-geth"
GOCACHE_DIR="${GOCACHE:-/tmp/go-build}"

pushd "$CHAIN_ROOT" >/dev/null
echo "⚠️  Building geth to $TMP_BUILD_OUT (chain/build is not writable in this environment)"
GOCACHE="$GOCACHE_DIR" go build -trimpath -tags=urfave_cli_no_docs,ckzg -o "$TMP_BUILD_OUT" ./cmd/geth
BIN_SRC="$TMP_BUILD_OUT"
popd >/dev/null

# Copy binary (renaming to juchain to match Dockerfile/start.sh expectations)
echo "=== Copying binary to Docker context ==="
if [ -f "$BIN_SRC" ]; then
    cp "$BIN_SRC" "$DOCKER_DIR/juchain"
else
    echo "❌ Error: $BIN_SRC not found. Build failed?"
    exit 1
fi

echo "=== Building Docker image ==="
pushd "$DOCKER_DIR"
docker build -t juchain-node:latest .
rm juchain  # Cleanup binary after build
popd

echo "✅ Docker image 'juchain-node:latest' built successfully."
