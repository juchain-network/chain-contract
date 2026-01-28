#!/bin/bash

# Default to localhost:8545 if not set
RPC_URL="${1:-http://localhost:8545}"
RETRIES=30

echo "⏳ Waiting for node at $RPC_URL to be ready..."

while [ $RETRIES -gt 0 ]; do
    if curl -s -X POST -H "Content-Type: application/json" \
        --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' \
        "$RPC_URL" > /dev/null; then
        echo ""
        echo "✅ Node is up and responding!"
        exit 0
    fi
    sleep 1
    RETRIES=$((RETRIES-1))
    echo -n "."
done

echo ""
echo "❌ Timeout waiting for node"
exit 1
