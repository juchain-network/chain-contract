#!/usr/bin/env bash

# Load environment variables from .env file
if [ -f "$(dirname "$0")/../../.env" ]; then
    export $(grep -v '^#' "$(dirname "$0")/../../.env" | xargs)
    echo "✅ Environment variables loaded from .env file"
else
    echo "⚠️  .env file not found, using default values"
fi

# Load test framework and utilities
source "$(dirname "$0")/test-framework.sh"
source "$(dirname "$0")/test-utils.sh"

test_name="Validator Punish Flow"
test_start "$test_name"

# 1. Environment Setup
test_step "Environment Setup"
run_script "script/integration/utils/InitEnv.s.sol:InitEnv"

# Get contract addresses from state file
PROPOSAL_ADDRESS=$(read_contract_address "proposal")
PUNISH_ADDRESS=$(read_contract_address "punish")
STAKING_ADDRESS=$(read_contract_address "staking")
VALIDATORS_ADDRESS=$(read_contract_address "validators")

# 2. Get Punishment Thresholds
test_step "Get Punishment Thresholds"

# Get punishThreshold and removeThreshold from proposal contract
PUNISH_THRESHOLD=$(cast call $PROPOSAL_ADDRESS "punishThreshold()(uint256)")
REMOVE_THRESHOLD=$(cast call $PROPOSAL_ADDRESS "removeThreshold()(uint256)")
echo "Punish Threshold: $PUNISH_THRESHOLD"
echo "Remove Threshold: $REMOVE_THRESHOLD"

# 3. Check Initial Validator State
test_step "Check Initial Validator State"

# Get first validator address and key (1-based index)
VALIDATOR=$(get_validator_address 1)
VALIDATOR_KEY=$(get_validator_private_key 1)
echo "Validator to punish: $VALIDATOR"

# Get miner address and key (1-based index)
MINER=$(get_validator_address 2)
MINER_KEY=$(get_validator_private_key 2)
echo "Miner address: $MINER"
echo "Miner key: $MINER_KEY"

# Check initial validator status from staking contract
INITIAL_STAKING_STATUS=$(cast call $STAKING_ADDRESS "getValidatorInfo(address)(uint256,uint256,uint256,uint256,bool,uint256,uint256,uint256,bool,uint256)" $VALIDATOR)
INITIAL_IS_JAILED=$(echo $INITIAL_STAKING_STATUS | cut -d' ' -f5)
INITIAL_JAIL_UNTIL=$(echo $INITIAL_STAKING_STATUS | cut -d' ' -f6)
echo "Initial staking status:"
echo "  Is Jailed: $INITIAL_IS_JAILED"
echo "  Jail Until: $INITIAL_JAIL_UNTIL"

# Check if validator is active in validators contract
INITIAL_IS_ACTIVE=$(cast call $VALIDATORS_ADDRESS "isActiveValidator(address)(bool)" $VALIDATOR)
echo "Initial active status in validators contract: $INITIAL_IS_ACTIVE"

# Verify validator is not jailed initially
if [ "$INITIAL_IS_JAILED" = "0" ]; then
    echo "✅ Validator is not jailed initially"
else
    echo "❌ Validator is already jailed initially"
    exit 1
fi

# 4. Distribute Rewards to Validators
test_step "Distribute Rewards to Validators"

# Set miner temporarily to distribute rewards
echo "Setting miner temporarily to $VALIDATOR"
cast send $PROPOSAL_ADDRESS "setMiner(address)" $VALIDATOR --private-key $VALIDATOR_KEY
cast send $VALIDATORS_ADDRESS "setMiner(address)" $VALIDATOR --private-key $VALIDATOR_KEY
cast send $STAKING_ADDRESS "setMiner(address)" $VALIDATOR --private-key $VALIDATOR_KEY
cast send $PUNISH_ADDRESS "setMiner(address)" $VALIDATOR --private-key $VALIDATOR_KEY

# Distribute block rewards multiple times to generate rewards for validators
echo "Distributing block rewards..."
for i in {1..5}; do
    # Increase block number for each reward distribution
    cast rpc evm_increaseTime 1
    cast rpc evm_mine

    # Distribute block reward
    cast send $VALIDATORS_ADDRESS "distributeBlockReward()()" --private-key $VALIDATOR_KEY --value 1ether
    echo "  Block reward distributed #$i"
done

echo "✅ Rewards distributed successfully"

# Get initial validator rewards and jailed amounts before punishments
echo "Getting initial validator info..."

# Convert scientific notation to plain number
# Usage: scientific_to_plain <scientific_notation_number>
get_validator_info() {
    local validator=$1

    # Capture JSON output using cast call --json
    local result=$(cast call --json "$VALIDATORS_ADDRESS" "getValidatorInfo(address)(address,uint8,uint256,uint256,uint256)" "$validator")

    # Parse JSON using jq
    local fee_addr=$(echo "$result" | jq -r '.[0]')
    local status=$(echo "$result" | jq -r '.[1]')
    local aac_incoming_raw=$(echo "$result" | jq -r '.[2]')
    local total_jailed_hb_raw=$(echo "$result" | jq -r '.[3]')
    local last_withdraw_block=$(echo "$result" | jq -r '.[4]')

    # Return as a string with values separated by spaces
    echo "$fee_addr $status $aac_incoming_raw $total_jailed_hb_raw $last_withdraw_block"
}


# Get initial info for the target validator
INITIAL_INFO=$(get_validator_info $VALIDATOR)
INITIAL_AAC_INCOMING=$(echo $INITIAL_INFO | awk '{print $3}')
INITIAL_TOTAL_JAILED=$(echo $INITIAL_INFO | awk '{print $4}')

echo "Initial validator rewards (aacIncoming): $INITIAL_AAC_INCOMING wei"
echo "Initial jailed amount (totalJailedHb): $INITIAL_TOTAL_JAILED wei"

# Check if rewards are zero using wei_is_zero function
if wei_is_zero "$INITIAL_AAC_INCOMING"; then
    echo "⚠️  WARNING: Validator has zero rewards after distribution"
    echo "   This could be because rewards are distributed to the block producer, not all validators"
fi

# 5. Test punishThreshold Behavior
test_step "Test punishThreshold Behavior"

echo "Applying punishments up to punishThreshold ($PUNISH_THRESHOLD)..."
echo "Setting miner temporarily to $MINER"
cast send $PROPOSAL_ADDRESS "setMiner(address)" $MINER --private-key $MINER_KEY
cast send $VALIDATORS_ADDRESS "setMiner(address)" $MINER --private-key $MINER_KEY
cast send $STAKING_ADDRESS "setMiner(address)" $MINER --private-key $MINER_KEY
cast send $PUNISH_ADDRESS "setMiner(address)" $MINER --private-key $MINER_KEY

# Apply punishments up to punishThreshold
for ((i=1; i<=PUNISH_THRESHOLD; i++)); do
    # Increase block number for each punishment
    cast rpc evm_increaseTime 1
    cast rpc evm_mine

    # Call punish method on punish contract
    cast send $PUNISH_ADDRESS "punish(address)" $VALIDATOR --private-key $MINER_KEY

    # Get updated punishment record
    MISSED_BLOCKS=$(cast call $PUNISH_ADDRESS "getPunishRecord(address)(uint256)" $VALIDATOR)
    echo "  Punishment #$i, Missed blocks: $MISSED_BLOCKS"
done

# Check validator status after reaching punishThreshold
echo "\nChecking validator status after reaching punishThreshold..."

# Check if validator is not jailed yet
IS_JAILED_AFTER_PUNISH_THRESHOLD=$(cast call $STAKING_ADDRESS "isValidatorJailed(address)(bool)" $VALIDATOR)
echo "Validator jailed status after punishThreshold: $IS_JAILED_AFTER_PUNISH_THRESHOLD"

# Validator should not be jailed at punishThreshold, only rewards should be forfeited
if [ "$IS_JAILED_AFTER_PUNISH_THRESHOLD" = "false" ]; then
    echo "✅ Validator is not jailed at punishThreshold (expected)"
else
    echo "❌ Validator is jailed at punishThreshold (unexpected)"
    exit 1
fi

# Get validator info after reaching punishThreshold
AFTER_PUNISH_THRESHOLD_INFO=$(get_validator_info $VALIDATOR)
AFTER_PUNISH_THRESHOLD_AAC=$(echo $AFTER_PUNISH_THRESHOLD_INFO | awk '{print $3}')
AFTER_PUNISH_THRESHOLD_JAILED=$(echo $AFTER_PUNISH_THRESHOLD_INFO | awk '{print $4}')

echo "Validator info after punishThreshold:"
echo "Rewards (aacIncoming): $AFTER_PUNISH_THRESHOLD_AAC wei"
echo "Jailed amount (totalJailedHb): $AFTER_PUNISH_THRESHOLD_JAILED wei"

# Check if validator has been penalized by examining the raw values
# We'll check if the jailed amount has changed from zero to non-zero, or increased
# This avoids issues with scientific notation

# Debug: Print the actual values being compared
 if wei_gt "$AFTER_PUNISH_THRESHOLD_JAILED" "$INITIAL_TOTAL_JAILED"; then
     echo "✅ Validator has been penalized (jailed amount increased)"
     echo "Initial jailed: $INITIAL_TOTAL_JAILED wei"
     echo "After punishThreshold: $AFTER_PUNISH_THRESHOLD_JAILED wei"
     echo "Penalty applied successfully"
 else
     echo "❌ Validator was not penalized (jailed amount did not increase)"
     exit 1
 fi

# 6. Test removeThreshold Behavior
test_step "Test removeThreshold Behavior"

echo "Applying additional punishments to reach removeThreshold ($REMOVE_THRESHOLD)..."

# Calculate remaining punishments needed to reach removeThreshold
CURRENT_MISSED_BLOCKS=$(cast call $PUNISH_ADDRESS "getPunishRecord(address)(uint256)" $VALIDATOR)
REMAINING_PUNISHMENTS=$((REMOVE_THRESHOLD - CURRENT_MISSED_BLOCKS))

# Apply remaining punishments to reach removeThreshold
for ((i=1; i<=REMAINING_PUNISHMENTS; i++)); do
    # Increase block number for each punishment
    cast rpc evm_increaseTime 1
    cast rpc evm_mine

    # Call punish method on punish contract
    cast send $PUNISH_ADDRESS "punish(address)" $VALIDATOR --private-key $MINER_KEY

    # Get updated punishment record
    MISSED_BLOCKS=$(cast call $PUNISH_ADDRESS "getPunishRecord(address)(uint256)" $VALIDATOR)
    echo "  Punishment #$((PUNISH_THRESHOLD + i)), Missed blocks: $MISSED_BLOCKS"
done

# Check validator status after reaching removeThreshold
echo "\nChecking validator status after reaching removeThreshold..."

# Check if validator is now jailed
IS_JAILED_AFTER_REMOVE_THRESHOLD=$(cast call $STAKING_ADDRESS "isValidatorJailed(address)(bool)" $VALIDATOR)
echo "Validator jailed status after removeThreshold: $IS_JAILED_AFTER_REMOVE_THRESHOLD"

# Validator should be jailed at removeThreshold
if [ "$IS_JAILED_AFTER_REMOVE_THRESHOLD" = "true" ]; then
    echo "✅ Validator is jailed at removeThreshold (expected)"
else
    echo "❌ Validator is not jailed at removeThreshold (unexpected)"
    exit 1
fi

# Check missed blocks counter after reaching removeThreshold (should be reset)
MISSED_BLOCKS_AFTER_REMOVE=$(cast call $PUNISH_ADDRESS "getPunishRecord(address)(uint256)" $VALIDATOR)
echo "Missed blocks after removeThreshold: $MISSED_BLOCKS_AFTER_REMOVE"

if [ "$MISSED_BLOCKS_AFTER_REMOVE" -eq 0 ]; then
    echo "✅ Missed blocks counter is reset to 0 after removeThreshold (expected)"
else
    echo "⚠️  Missed blocks counter is not reset to 0 after removeThreshold"
fi

# 7. Verify Validator Removal from Top Validators
test_step "Verify Validator Removal from Top Validators"

# Get top validators list
TOP_VALIDATORS=$(cast call $VALIDATORS_ADDRESS "getTopValidators()(address[])")
echo "Top validators list: $TOP_VALIDATORS"

# Check if validator is removed from top validators list
if ! echo $TOP_VALIDATORS | grep -q $VALIDATOR; then
    echo "✅ Validator is removed from top validators list (expected)"
else
    echo "⚠️  Validator is still in top validators list"
fi

# 8. Verify Validator Status Changes
test_step "Verify Validator Status Changes"

# Get validator status from validators contract
VALIDATOR_INFO=$(cast call $VALIDATORS_ADDRESS "getValidatorInfo(address)(address payable,uint8,uint256,uint256,uint256)" $VALIDATOR)
VALIDATOR_STATUS=$(echo $VALIDATOR_INFO | cut -d' ' -f2)
echo "Validator status from validators contract: $VALIDATOR_STATUS"

# Status enum values: 0=NotExist, 1=Active, 2=Jailed
if [ "$VALIDATOR_STATUS" -eq 2 ]; then
    echo "✅ Validator status from validators contract is Jailed (2)"
elif [ "$VALIDATOR_STATUS" -eq 1 ]; then
    echo "⚠️  Validator status from validators contract is still Active (1), which is expected until next epoch"
elif [ "$VALIDATOR_STATUS" -eq 0 ]; then
    echo "❌ Validator status from validators contract is NotExist (0)"
    exit 1
else
    echo "❌ Invalid validator status from validators contract: $VALIDATOR_STATUS"
    exit 1
fi

test_end
