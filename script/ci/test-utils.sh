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