#!/bin/bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
CONFIG_PATH="${ROOT_DIR}/data/test_config.yaml"

cd "${ROOT_DIR}"

TESTS=$(grep -h "^func Test" "${ROOT_DIR}/tests"/*.go | sed -E 's/^func (Test[^(]+).*/\1/' | grep -v '^TestMain$' | sort -u)

if [ -z "${TESTS}" ]; then
  echo "❌ No tests found in ${ROOT_DIR}/tests"
  exit 1
fi

for T in ${TESTS}; do
  echo "=============================="
  echo "🧪 Running ${T}"
  echo "=============================="
  make clean
  docker system prune -f --volumes
  sleep 5
  make init run ready
  go test ./tests/... -v -run "^${T}$" -count=1 -parallel=1 -timeout 20m -config "${CONFIG_PATH}"
  make stop
  sleep 3
done
