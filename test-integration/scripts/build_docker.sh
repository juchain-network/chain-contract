#!/bin/bash
set -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PROJECT_ROOT="$SCRIPT_DIR/../../.."
CHAIN_ROOT="$PROJECT_ROOT/chain"
DOCKER_DIR="$SCRIPT_DIR/../docker"

TARGET_GOOS="${TARGET_GOOS:-linux}"
TARGET_GOARCH="${TARGET_GOARCH:-}"
if [ -z "$TARGET_GOARCH" ]; then
    DOCKER_ARCH="$(docker info --format '{{.Architecture}}' 2>/dev/null || true)"
    case "$DOCKER_ARCH" in
        amd64|x86_64) TARGET_GOARCH="amd64" ;;
        arm64|aarch64) TARGET_GOARCH="arm64" ;;
        arm|armv7l) TARGET_GOARCH="arm" ;;
        *) TARGET_GOARCH="$(go env GOARCH)" ;;
    esac
fi

binary_matches_target() {
    local info="$1"
    local arch="$2"
    case "$arch" in
        amd64) echo "$info" | grep -Eiq 'x86-64|amd64' ;;
        arm64) echo "$info" | grep -Eiq 'aarch64|arm64' ;;
        arm) echo "$info" | grep -Eiq ' ARM ' ;;
        *) echo "$info" | grep -Eiq "$arch" ;;
    esac
}

echo "=== Building geth binary from $CHAIN_ROOT ==="
echo "=== Target binary: ${TARGET_GOOS}/${TARGET_GOARCH} ==="

# Fast path: reuse existing binary unless FORCE_BUILD is set.
if [ -x "$DOCKER_DIR/juchain" ] && [ -z "${FORCE_BUILD:-}" ]; then
    if command -v file >/dev/null 2>&1; then
        BIN_INFO="$(file "$DOCKER_DIR/juchain" 2>/dev/null || true)"
        if echo "$BIN_INFO" | grep -q "ELF" && binary_matches_target "$BIN_INFO" "$TARGET_GOARCH"; then
            echo "✅ Binary already exists at $DOCKER_DIR/juchain (${TARGET_GOOS}/${TARGET_GOARCH})"
            exit 0
        fi
        echo "⚠️  Existing binary incompatible, rebuilding: $BIN_INFO"
    else
        echo "✅ Binary already exists at $DOCKER_DIR/juchain (file tool unavailable, set FORCE_BUILD=1 to rebuild)"
        exit 0
    fi
fi

# Compile geth into a writable temp path to avoid permission issues in chain/build.
TMP_BUILD_OUT="${TMPDIR:-/tmp}/juchain-geth"
GOCACHE_DIR="${GOCACHE:-/tmp/go-build}"
CGO_MODE="${CGO_ENABLED:-0}"

pushd "$CHAIN_ROOT" >/dev/null
echo "⚠️  Building geth to $TMP_BUILD_OUT (chain/build is not writable in this environment)"
GOCACHE="$GOCACHE_DIR" GOOS="$TARGET_GOOS" GOARCH="$TARGET_GOARCH" CGO_ENABLED="$CGO_MODE" \
    go build -trimpath -tags=urfave_cli_no_docs,ckzg -o "$TMP_BUILD_OUT" ./cmd/geth
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
