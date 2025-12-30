#!/bin/bash

# Simulate block growth - precise control
advance_blocks() {
    local num_blocks=$1
    echo "🔄 Advancing $num_blocks blocks..."
    cast rpc evm_increaseBlocks $num_blocks
    cast rpc evm_mine
    echo "✅ Advanced $num_blocks blocks to $(get_current_block)"
}

# Simulate time passage - precise control
advance_time() {
    local seconds=$1
    echo "⏰ Advancing time by $seconds seconds..."
    cast rpc evm_increaseTime $seconds
    cast rpc evm_mine
    echo "✅ Advanced time by $seconds seconds to $(get_current_timestamp) ($(date -r $(get_current_timestamp)))"
}

# Simulate to specific block number
advance_to_block() {
    local target_block=$1
    local current_block=$(get_current_block)
    if [ $target_block -le $current_block ]; then
        echo "⚠️  Target block $target_block is less than or equal to current block $current_block"
        return 1
    fi
    local blocks_to_advance=$((target_block - current_block))
    advance_blocks $blocks_to_advance
}

# Simulate to specific timestamp
advance_to_timestamp() {
    local target_timestamp=$1
    local current_timestamp=$(get_current_timestamp)
    if [ $target_timestamp -le $current_timestamp ]; then
        echo "⚠️  Target timestamp $target_timestamp is less than or equal to current timestamp $current_timestamp"
        return 1
    fi
    local seconds_to_advance=$((target_timestamp - current_timestamp))
    advance_time $seconds_to_advance
}

# Get current block number
get_current_block() {
    cast block-number
}

# Get current timestamp
get_current_timestamp() {
    cast block latest --field timestamp
}

# Execute Solidity test script - support parameter passing
run_script() {
    local script=$1
    local params=${2:-""}
    local rpc_url=${3:-"http://localhost:8545"}
    echo "📜 Running script: $script $params"
    forge script $script --rpc-url $rpc_url --broadcast --skip-simulation -vvv $params
    if [ $? -eq 0 ]; then
        echo "✅ Script $script executed successfully"
        return 0
    else
        echo "❌ Script $script failed"
        return 1
    fi
}

# Verify operation result
verify_result() {
    local description=$1
    local script=$2
    echo "🔍 Verifying: $description"
    run_script $script
}

# Read contract address from test state JSON file
read_contract_address() {
    local contract_name=$1
    local state_file=${2:-"$(dirname "$0")/../../state/test_state.json"}
    
    if [ ! -f "$state_file" ]; then
        echo "❌ State file not found: $state_file"
        return 1
    fi
    
    # Use jq to parse JSON and extract contract address
    local address=$(jq -r ".contracts.$contract_name" "$state_file")
    
    if [ -z "$address" ] || [ "$address" = "null" ]; then
        echo "❌ Contract address not found for: $contract_name"
        return 1
    fi
    
    echo "$address"
    return 0
}

# Read deployer address from test state JSON file
read_deployer_address() {
    local state_file=${1:-"$(dirname "$0")/../../state/test_state.json"}
    
    if [ ! -f "$state_file" ]; then
        echo "❌ State file not found: $state_file"
        return 1
    fi
    
    local address=$(jq -r ".deployer.address" "$state_file")
    
    if [ -z "$address" ] || [ "$address" = "null" ]; then
        echo "❌ Deployer address not found"
        return 1
    fi
    
    echo "$address"
    return 0
}

# Read account from test state JSON file
read_account() {
    local account_type=$1  # validatorAccounts or delegatorAccounts
    local index=$2
    local state_file=${3:-"$(dirname "$0")/../../state/test_state.json"}
    
    if [ ! -f "$state_file" ]; then
        echo "❌ State file not found: $state_file"
        return 1
    fi
    
    local address=$(jq -r ".accounts.$account_type[$index]" "$state_file")
    
    if [ -z "$address" ] || [ "$address" = "null" ]; then
        echo "❌ Account not found for: $account_type[$index]"
        return 1
    fi
    
    echo "$address"
    return 0
}

# Get validator private key from environment variables
# Usage: get_validator_private_key <index> (1-based)
get_validator_private_key() {
    local index=$1
    local key_var="VALIDATOR_PRIVATE_KEY_$index"
    
    if [ -z "${!key_var}" ]; then
        echo "❌ Validator private key not found for index $index: $key_var"
        return 1
    fi
    
    echo "${!key_var}"
    return 0
}

# Get validator address from private key
# Usage: get_validator_address <index> (1-based)
get_validator_address() {
    local index=$1
    local private_key=$(get_validator_private_key $index)
    
    if [ $? -ne 0 ]; then
        return 1
    fi
    
    cast wallet address $private_key
    return 0
}

# Get delegator private key from environment variables
# Usage: get_delegator_private_key <index> (1-based)
get_delegator_private_key() {
    local index=$1
    local key_var="DELEGATOR_PRIVATE_KEY_$index"
    
    if [ -z "${!key_var}" ]; then
        echo "❌ Delegator private key not found for index $index: $key_var"
        return 1
    fi
    
    echo "${!key_var}"
    return 0
}

# Get delegator address from private key
# Usage: get_delegator_address <index> (1-based)
get_delegator_address() {
    local index=$1
    local private_key=$(get_delegator_private_key $index)
    
    if [ $? -ne 0 ]; then
        return 1
    fi
    
    cast wallet address $private_key
    return 0
}

# Get all validator private keys as an array
# Usage: get_all_validator_private_keys
# Returns: Array of validator private keys
get_all_validator_private_keys() {
    local keys=()
    local i=1
    
    # Loop through validator private keys up to VALIDATOR_COUNT
    # Using VALIDATOR_COUNT from environment variables for safety
    local max_validators=${VALIDATOR_COUNT:-9}  # Default to 9 if not set
    
    for ((i=1; i<=max_validators; i++)); do
        local key_var="VALIDATOR_PRIVATE_KEY_$i"
        local key="${!key_var}"
        
        if [ -n "$key" ]; then
            keys+=($key)
        else
            # Stop at the first missing key
            break
        fi
    done
    
    echo "${keys[@]}"
    return 0
}

# Get all delegator private keys as an array
# Usage: get_all_delegator_private_keys
# Returns: Array of delegator private keys
get_all_delegator_private_keys() {
    local keys=()
    local i=1
    
    # Loop through delegator private keys up to DELEGATOR_COUNT
    # Using DELEGATOR_COUNT from environment variables for safety
    local max_delegators=${DELEGATOR_COUNT:-9}  # Default to 9 if not set
    
    for ((i=1; i<=max_delegators; i++)); do
        local key_var="DELEGATOR_PRIVATE_KEY_$i"
        local key="${!key_var}"
        
        if [ -n "$key" ]; then
            keys+=($key)
        else
            # Stop at the first missing key
            break
        fi
    done
    
    echo "${keys[@]}"
    return 0
}

# Compare two wei integers using bc
# Usage: wei_gt <a> <b>
# Return: 0 (true) / 1 (false)
wei_gt() {
    local a="$1"
    local b="$2"

    # Ensure both are pure integers
    if ! [[ "$a" =~ ^[0-9]+$ && "$b" =~ ^[0-9]+$ ]]; then
        echo "ERROR: wei_gt expects integer values, got '$a' and '$b'" >&2
        return 2
    fi

    # bc returns 1 if true, 0 if false
    (( $(echo "$a > $b" | bc) ))
}

wei_is_zero() {
    local v="$1"

    [[ "$v" =~ ^[0-9]+$ ]] || return 1
    (( $(echo "$v == 0" | bc) ))
}