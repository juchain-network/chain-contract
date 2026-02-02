#!/bin/bash
set -euo pipefail

NODES=("$@")
if [ "${#NODES[@]}" -eq 0 ]; then
  NODES=(node0 node1 node2)
fi

check_progress() {
  local node="$1"
  local container="juchain-${node}"
  echo "⛏️  Checking block progress on ${container}..."
  for i in $(seq 1 60); do
    if docker exec "${container}" curl -s -X POST -H "Content-Type: application/json" \
      --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' \
      "http://127.0.0.1:8545" >/dev/null 2>&1; then
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
      done
      echo "⚠️  ${container} did not observe block progress yet"
      return 0
    fi
    sleep 1
  done
  return 0
}

pids=()
for node in "${NODES[@]}"; do
  check_progress "$node" &
  pids+=($!)
done

for pid in "${pids[@]}"; do
  wait "$pid" || true
done
