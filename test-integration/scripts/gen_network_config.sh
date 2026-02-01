#!/bin/bash
set -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PROJECT_ROOT="$( cd "$SCRIPT_DIR/../.." && pwd )"
TEST_INT_DIR="$SCRIPT_DIR/.."
DATA_DIR="$TEST_INT_DIR/data"
TEMPLATE_GENESIS="$TEST_INT_DIR/templates/genesis.tpl.json"

# Clean up data directory
rm -rf "$DATA_DIR"
mkdir -p "$DATA_DIR"

echo "=== Generating Network Configuration ==="

# 1. Generate Keys helper (using a temporary Go program)
cat > "$DATA_DIR/genkeys.go" <<EOF
package main

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	key, _ := crypto.GenerateKey()
	addr := crypto.PubkeyToAddress(key.PublicKey)
	priv := hex.EncodeToString(crypto.FromECDSA(key))
	pub := hex.EncodeToString(crypto.FromECDSAPub(&key.PublicKey)[1:]) // Remove 04 prefix
	fmt.Printf("%s,%s,%s\n", addr.Hex(), priv, pub)
}
EOF

# Initialize go module for genkeys if needed
pushd "$TEST_INT_DIR" > /dev/null
go mod tidy > /dev/null 2>&1
popd > /dev/null

generate_key() {
    cd "$TEST_INT_DIR" && go run "$DATA_DIR/genkeys.go" 2>/dev/null
}

# Generate Keys
echo "Generating keys..."
# Funder
IFS=',' read -r FUNDER_ADDR FUNDER_PRIV FUNDER_PUB <<< $(generate_key)
# Trim any potential whitespace/newlines
FUNDER_ADDR=$(echo "$FUNDER_ADDR" | tr -d '[:space:]')
echo "Funder: $FUNDER_ADDR"

# Validators (3 nodes) + Sync Node (1 node) = 4 nodes total
VAL_ADDRS=()
VAL_PRIVS=()
ENODES=()

# We generate for 0..3 (4 nodes). 0-2 are validators, 3 is sync.
for i in {0..3}; do
    # Create node directory
    mkdir -p "$DATA_DIR/node$i/keystore"
    mkdir -p "$DATA_DIR/node$i/geth" 
    
    # Node P2P Key
    IFS=',' read -r NODE_ADDR NODE_PRIV NODE_PUB <<< $(generate_key)
    echo "$NODE_PRIV" > "$DATA_DIR/node$i/nodekey"
    
    # Construct Enode URL
    ENODES+=("enode://$NODE_PUB@node$i:30303")
    
    if [ $i -lt 3 ]; then
        # Validator keys for 0-2
        IFS=',' read -r ADDR PRIV PUB <<< $(generate_key)
        ADDR=$(echo "$ADDR" | tr -d '[:space:]')
        VAL_ADDRS+=($ADDR)
        VAL_PRIVS+=($PRIV)
        echo "Validator $i: $ADDR"
        echo "$PRIV" > "$DATA_DIR/node$i/validator.key"
    else
        echo "Node $i: Sync Node (No validator key)"
    fi
done

# Generate config.toml for each node
echo "Generating config.toml for nodes..."

# Prepare StaticNodes string: ["enode://...", "enode://..."]
STATIC_NODES_TOML="["
for i in {0..3}; do
    if [ $i -ne 0 ]; then STATIC_NODES_TOML+=', ';
    fi
    STATIC_NODES_TOML+="\"${ENODES[$i]}\""
done
STATIC_NODES_TOML+="]"

for i in {0..3}; do
    cat > "$DATA_DIR/node$i/config.toml" <<EOF
[Node]
UserIdent = "juchain-node$i"
DataDir = "/data"
IPCPath = "geth.ipc"
HTTPHost = "0.0.0.0"
HTTPPort = 8545
WSHost = "0.0.0.0"
WSPort = 8546
[Node.P2P]
MaxPeers = 50
ListenAddr = ":30303"
StaticNodes = $STATIC_NODES_TOML
EOF
done

# 2. Build Genesis
echo "Building genesis.json..."

# Compile contracts to ensure we use the latest code
echo "Compiling contracts..."
pushd "$PROJECT_ROOT" > /dev/null
forge build
popd > /dev/null

# Generate System Contracts Alloc using the helper script
echo "Generating system contracts alloc..."
node "$SCRIPT_DIR/build_alloc.js" > "$DATA_DIR/sys_contracts.json"
if [ $? -ne 0 ]; then
    echo "❌ Failed to generate system contracts alloc"
    exit 1
fi

# Add Funder and Validators to Alloc
# Prepare comma-separated validators list
VAL_ADDRS_CSV=$(IFS=,; echo "${VAL_ADDRS[*]}")

echo "Merging alloc with funder and validators..."
ALLOC_JSON=$(node "$SCRIPT_DIR/merge_alloc.js" "$DATA_DIR/sys_contracts.json" "$FUNDER_ADDR" "$VAL_ADDRS_CSV")
if [ $? -ne 0 ]; then
    echo "❌ Failed to merge alloc"
    exit 1
fi

VANITY="0000000000000000000000000000000000000000000000000000000000000000"
VALIDATORS_HEX=""
for addr in "${VAL_ADDRS[@]}"; do
    # Remove 0x prefix
    VALIDATORS_HEX+="${addr:2}"
done
SUFFIX="0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000" # 65 bytes zeros

EXTRA_DATA="0x${VANITY}${VALIDATORS_HEX}${SUFFIX}"

# Replace placeholders
echo "$ALLOC_JSON" > "$DATA_DIR/alloc.json"

jq --slurpfile allocList "$DATA_DIR/alloc.json" --arg extra "$EXTRA_DATA" \
   '.alloc = $allocList[0] | .extraData = $extra' "$TEMPLATE_GENESIS" > "$DATA_DIR/genesis.json"

# 3. Generate test_config.yaml
echo "Generating test_config.yaml..."
cat > "$DATA_DIR/test_config.yaml" <<EOF
rpcs:
  - "http://localhost:18545"

funder:
  address: "$FUNDER_ADDR"
  private_key: "$FUNDER_PRIV"

validators:
EOF

for i in {0..2}; do
cat >> "$DATA_DIR/test_config.yaml" <<EOF
  - address: "${VAL_ADDRS[$i]}"
    private_key: "${VAL_PRIVS[$i]}"
EOF
done

cat >> "$DATA_DIR/test_config.yaml" <<EOF
test:
  funding_amount: "100000000000000000000" # 100 ETH
EOF

echo "✅ Configuration generated at $DATA_DIR"