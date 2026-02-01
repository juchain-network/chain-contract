#!/bin/bash
set -e

# Initialize genesis if not already done
if [ ! -d "/data/geth/chaindata" ]; then
    echo "Initializing genesis..."
    juchain --datadir /data init /genesis.json 
fi

# Import validator key if present
if [ -f "/data/validator.key" ]; then
    echo "Importing validator key..."
    # Import and capture the address
    # We use a temporary password file for non-interactive import
    echo "" > /tmp/password
    juchain account import --datadir /data --password /tmp/password /data/validator.key
fi

# Get address (assuming only one account imported)
# Or we can rely on --miner.etherbase if we passed the address in ENV
VAL_ADDR=${VAL_ADDR:-$(juchain account list --datadir /data | head -n 1 | cut -d '{' -f 2 | cut -d '}' -f 1)}

echo "Starting node with validator: $VAL_ADDR"

# Construct bootnodes flag if static-nodes.json is not working or as backup
# But static-nodes.json in /data/geth/ (or /data/) usually works.

# Start juchain without mining, then enable mining after peers connect.
MIN_PEERS="${MIN_PEERS:-3}"
MINER_START_DELAY="${MINER_START_DELAY:-0}"
RPC_URL="http://127.0.0.1:8545"

juchain \
    --config /data/config.toml \
    --networkid 666666 \
    --nodekey /data/nodekey \
    --http \
    --http.api "eth,net,web3,debug,admin,personal,miner,txpool" \
    --http.corsdomain "*" \
    --http.vhosts "*" \
    --ws \
    --ws.api "eth,net,web3,debug,admin,personal,miner,txpool" \
    --ws.origins "*" \
    --miner.etherbase "$VAL_ADDR" \
    --miner.gasprice 0 \
    --unlock "$VAL_ADDR" \
    --password /tmp/password \
    --allow-insecure-unlock \
    --nat extip:$(hostname -i) &

NODE_PID=$!

wait_for_rpc() {
    for i in $(seq 1 120); do
        RESP=$(curl -s -X POST -H "Content-Type: application/json" \
            --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' \
            "$RPC_URL" || true)
        HEX=$(echo "$RESP" | sed -n 's/.*"result":"\(0x[0-9a-fA-F]*\)".*/\1/p')
        if [ -n "$HEX" ]; then
            return 0
        fi
        if ! kill -0 "$NODE_PID" >/dev/null 2>&1; then
            return 1
        fi
        sleep 1
    done
    return 1
}

wait_for_peers() {
    for i in $(seq 1 300); do
        RESP=$(curl -s -X POST -H "Content-Type: application/json" \
            --data '{"jsonrpc":"2.0","method":"net_peerCount","params":[],"id":1}' \
            "$RPC_URL" || true)
        HEX=$(echo "$RESP" | sed -n 's/.*"result":"\(0x[0-9a-fA-F]*\)".*/\1/p')
        if [ -n "$HEX" ]; then
            DEC=${HEX#0x}
            if [ -n "$DEC" ]; then
                CUR=$((16#$DEC))
                if [ "$CUR" -ge "$MIN_PEERS" ]; then
                    return 0
                fi
            fi
        fi
        if ! kill -0 "$NODE_PID" >/dev/null 2>&1; then
            return 1
        fi
        sleep 1
    done
    # We return success even if peer count not reached to allow miner to start anyway
    # This prevents the container from exiting
    echo "⚠️ Peer count not reached ($MIN_PEERS), continuing anyway..."
    return 0
}

if ! wait_for_rpc; then
    echo "❌ RPC not ready, exiting"
    wait "$NODE_PID"
    exit 1
fi

if ! wait_for_peers; then
    echo "❌ Peers not ready (min=$MIN_PEERS), exiting"
    wait "$NODE_PID"
    exit 1
fi

if [ "$MINER_START_DELAY" -gt 0 ]; then
    sleep "$MINER_START_DELAY"
fi

curl -s -X POST -H "Content-Type: application/json" \
    --data '{"jsonrpc":"2.0","method":"miner_start","params":[],"id":1}' \
    "$RPC_URL" >/dev/null 2>&1 || true

wait "$NODE_PID"
