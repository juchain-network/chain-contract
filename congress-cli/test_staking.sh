#!/bin/bash

# JuChain Congress-CLI Staking Commands Test Script
# This script tests all staking command syntax without requiring a live network

set -e

CLI="./congress-cli"
RPC="http://localhost:8545"
DUMMY_ADDR="0x1234567890123456789012345678901234567890"
DUMMY_VALIDATOR="0x0987654321098765432109876543210987654321"

echo "🚀 Testing JuChain Congress-CLI Staking Commands"
echo "================================================="

# Test main staking help
echo "📋 Testing main staking command..."
$CLI staking --help > /dev/null
echo "✅ Staking command help works"

# Test register-validator help
echo "📋 Testing register-validator command..."
$CLI staking register-validator --help > /dev/null
echo "✅ Register-validator command help works"

# Test delegate help
echo "📋 Testing delegate command..."
$CLI staking delegate --help > /dev/null
echo "✅ Delegate command help works"

# Test undelegate help
echo "📋 Testing undelegate command..."
$CLI staking undelegate --help > /dev/null
echo "✅ Undelegate command help works"

# Test claim-rewards help
echo "📋 Testing claim-rewards command..."
$CLI staking claim-rewards --help > /dev/null
echo "✅ Claim-rewards command help works"

# Test query-validator help
echo "📋 Testing query-validator command..."
$CLI staking query-validator --help > /dev/null
echo "✅ Query-validator command help works"

# Test query-delegation help
echo "📋 Testing query-delegation command..."
$CLI staking query-delegation --help > /dev/null
echo "✅ Query-delegation command help works"

# Test list-top-validators help
echo "📋 Testing list-top-validators command..."
$CLI staking list-top-validators --help > /dev/null
echo "✅ List-top-validators command help works"

echo ""
echo "🎉 All staking commands are properly integrated!"
echo ""
echo "📝 Command Examples:"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "Register Validator:"
echo "  $CLI staking register-validator --rpc_laddr $RPC --proposer $DUMMY_ADDR --stake-amount 10000 --commission-rate 500"
echo ""
echo "Delegate Tokens:"
echo "  $CLI staking delegate --rpc_laddr $RPC --delegator $DUMMY_ADDR --validator $DUMMY_VALIDATOR --amount 1000"
echo ""
echo "Query Validator:"
echo "  $CLI staking query-validator --rpc_laddr $RPC --address $DUMMY_VALIDATOR"
echo ""
echo "List Top Validators:"
echo "  $CLI staking list-top-validators --rpc_laddr $RPC --limit 21"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"

echo ""
echo "🔗 For detailed usage instructions, see: STAKING_USAGE.md"
echo "🏗️  For architecture details, see: contracts/README.md"
