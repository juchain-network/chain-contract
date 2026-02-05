#!/bin/bash
set -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PROJECT_ROOT="$SCRIPT_DIR/../../.."
CHAIN_ROOT="$PROJECT_ROOT/chain"
DOCKER_DIR="$SCRIPT_DIR/../docker"

echo "=== Building geth binary from $CHAIN_ROOT ==="

# Fast path: reuse existing binary unless FORCE_BUILD is set.
if [ -x "$DOCKER_DIR/juchain" ] && [ -z "${FORCE_BUILD:-}" ]; then
    echo "✅ Binary already exists at $DOCKER_DIR/juchain (set FORCE_BUILD=1 to rebuild)"
    exit 0
fi

# Compile geth into a writable temp path to avoid permission issues in chain/build.
TMP_BUILD_OUT="${TMPDIR:-/tmp}/juchain-geth"
GOCACHE_DIR="${GOCACHE:-/tmp/go-build}"

pushd "$CHAIN_ROOT" >/dev/null
echo "⚠️  Building geth to $TMP_BUILD_OUT (chain/build is not writable in this environment)"
GOCACHE="$GOCACHE_DIR" go build -trimpath -tags=urfave_cli_no_docs,ckzg -o "$TMP_BUILD_OUT" ./cmd/geth
BIN_SRC="$TMP_BUILD_OUT"
popd >/dev/null

# Copy binary (renaming to juchain to match Dockerfile/start.sh expectations)
echo "=== Copying binary to docker/juchain ==="
if [ -f "$BIN_SRC" ]; then
    cp "$BIN_SRC" "$DOCKER_DIR/juchain"
    chmod +x "$DOCKER_DIR/juchain"
else
    echo "❌ Error: $BIN_SRC not found. Build failed?"
    exit 1
fi

echo "✅ Binary ready at $DOCKER_DIR/juchain"
