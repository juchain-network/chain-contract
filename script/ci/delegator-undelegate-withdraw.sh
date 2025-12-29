#!/bin/bash

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

test_name="Delegator Undelegate and Withdraw"
test_start "$test_name"

# 1. Environment Setup
test_step "Environment Setup"
run_script "script/integration/AccountSetup.s.sol:AccountSetupScript"
run_script "script/integration/ContractDeployment.s.sol:ContractDeploymentScript"

# 2. Get System Parameters
test_step "Get System Parameters"
# Get parameters from contract or use .env defined values
UNBONDING_PERIOD=${UNBONDING_PERIOD:-$(cast call 0x000000000000000000000000000000000000F013 "getUnbondingPeriod()(uint256)")}
echo "Unbonding Period: $UNBONDING_PERIOD seconds"

# 3. Delegator Delegate
test_step "Delegator Delegate"
run_script "script/integration/DelegatorDelegate.s.sol:DelegatorDelegateScript"
verify_result "Delegator delegated" "script/integration/DelegatorStatusCheck.s.sol:DelegatorStatusCheckScript"

# 4. Delegator Undelegate
test_step "Delegator Undelegate"
run_script "script/integration/DelegatorUndelegate.s.sol:DelegatorUndelegateScript"
verify_result "Delegator undelegated" "script/integration/DelegatorStatusCheck.s.sol:DelegatorStatusCheckScript"

# 5. Wait for Unbonding Period
test_step "Wait for Unbonding Period"
advance_time $UNBONDING_PERIOD
advance_blocks 1
verify_result "Unbonding period passed" "script/integration/DelegatorStatusCheck.s.sol:DelegatorStatusCheckScript"

# 6. Delegator Withdraw Unbonded
test_step "Delegator Withdraw Unbonded"
run_script "script/integration/DelegatorWithdrawUnbonded.s.sol:DelegatorWithdrawUnbondedScript"
verify_result "Delegator withdrew unbonded funds" "script/integration/DelegatorStatusCheck.s.sol:DelegatorStatusCheckScript"

test_end