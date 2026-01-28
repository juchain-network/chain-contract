#!/bin/bash
set -e

# Initialize genesis if not already done
if [ ! -d "/data/geth" ] && [ ! -d "/data/chaindata" ]; then
    echo "Initializing genesis..."
    juchain init /genesis.json --datadir /data
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

# Start juchain
# Note: Enabling --mine --miner.threads 1 for PoA/PoSA
# Using --allow-insecure-unlock to allow unlocking the account for signing
# Using --config to load StaticNodes and other settings
exec juchain \
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
    --mine \
    --miner.etherbase "$VAL_ADDR" \
    --miner.gasprice 0 \
    --unlock "$VAL_ADDR" \
    --password /tmp/password \
    --allow-insecure-unlock \
    --nat extip:$(hostname -i)
