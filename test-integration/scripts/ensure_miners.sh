#!/bin/bash
set -euo pipefail

NODES=("$@")
if [ "${#NODES[@]}" -eq 0 ]; then
  NODES=(node0 node1 node2)
fi

start_miner() {
  local node="$1"
  local container="juchain-${node}"
  echo "⛏️  Ensuring miner started on ${container}..."
  for i in $(seq 1 60); do
    if docker exec "${container}" curl -s -X POST -H "Content-Type: application/json" \
      --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' \
      "http://127.0.0.1:8545" >/dev/null 2>&1; then
      # Reset miner to recover from early unlock failures, then verify block progress.
      docker exec "${container}" curl -s -X POST -H "Content-Type: application/json" \
        --data '{"jsonrpc":"2.0","method":"miner_stop","params":[],"id":1}' \
        "http://127.0.0.1:8545" >/dev/null 2>&1 || true
      docker exec "${container}" curl -s -X POST -H "Content-Type: application/json" \
        --data '{"jsonrpc":"2.0","method":"miner_start","params":[],"id":1}' \
        "http://127.0.0.1:8545" >/dev/null 2>&1 || true

      # Wait for block number to advance (node sees chain progress even if not sealing).
      local prev cur
      for j in $(seq 1 30); do
        prev=$(docker exec "${container}" curl -s -X POST -H "Content-Type: application/json" \
          --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' \
          "http://127.0.0.1:8545" | sed -n 's/.*"result":"\\(0x[0-9a-fA-F]*\\)".*/\\1/p')
        sleep 2
        cur=$(docker exec "${container}" curl -s -X POST -H "Content-Type: application/json" \
          --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' \
          "http://127.0.0.1:8545" | sed -n 's/.*"result":"\\(0x[0-9a-fA-F]*\\)".*/\\1/p')
        if [ -n "$prev" ] && [ -n "$cur" ] && [ "$cur" != "$prev" ]; then
          return 0
        fi
        if [ $((j % 5)) -eq 0 ]; then
          docker exec "${container}" curl -s -X POST -H "Content-Type: application/json" \
            --data '{"jsonrpc":"2.0","method":"miner_stop","params":[],"id":1}' \
            "http://127.0.0.1:8545" >/dev/null 2>&1 || true
          docker exec "${container}" curl -s -X POST -H "Content-Type: application/json" \
            --data '{"jsonrpc":"2.0","method":"miner_start","params":[],"id":1}' \
            "http://127.0.0.1:8545" >/dev/null 2>&1 || true
        fi
      done
      echo "⚠️  ${container} did not observe block progress after restart attempts"
      return 1
    fi
    sleep 1
  done
  return 0
}

pids=()
for node in "${NODES[@]}"; do
  start_miner "$node" &
  pids+=($!)
done

for pid in "${pids[@]}"; do
  wait "$pid" || true
done
