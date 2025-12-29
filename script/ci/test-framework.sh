#!/bin/bash

# Test result management
test_results=()
test_total=0
test_passed=0
test_failed=0

# Test start
test_start() {
    local test_name=$1
    echo "\n📋 Running Test: $test_name"
    echo "====================================="
    test_total=0
test_passed=0
test_failed=0
    test_results=()
    start_time=$(date +%s)
    
    # Record initial state
    echo "🔍 Initial State:"
    echo "  Block: $(get_current_block)"
    echo "  Timestamp: $(get_current_timestamp) ($(date -r $(get_current_timestamp)))"
}

# Test step
test_step() {
    local step_name=$1
    echo "\n▶ Step: $step_name"
    test_total=$((test_total + 1))
}

# Verify operation result - support custom verification script
verify_with_script() {
    local description=$1
    local script=$2
    local params=${3:-""}
    
    echo "🔍 Verifying: $description"
    if run_script "$script" "$params"; then
        test_passed=$((test_passed + 1))
        test_results+=("✅ $description")
        return 0
    else
        test_failed=$((test_failed + 1))
        test_results+=("❌ $description")
        return 1
    fi
}

# Record state change
record_state_change() {
    local description=$1
    echo "📊 $description:"
    echo "  Block: $(get_current_block)"
    echo "  Timestamp: $(get_current_timestamp) ($(date -r $(get_current_timestamp)))"
}

# Test end
test_end() {
    end_time=$(date +%s)
    duration=$((end_time - start_time))
    
    echo "\n====================================="
    echo "📋 Test Results"
    echo "====================================="
    
    for result in "${test_results[@]}"; do
        echo "$result"
    done
    
    echo "\n📊 Summary:"
    echo "  Total Steps: $test_total"
    echo "  Passed: $test_passed"
    echo "  Failed: $test_failed"
    echo "  Duration: $duration seconds"
    
    if [ $test_failed -eq 0 ]; then
        echo "\n🎉 All tests passed!"
        return 0
    else
        echo "\n❌ Some tests failed!"
        return 1
    fi
}