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

test_name="Validator Full Lifecycle"
test_start "$test_name"

# 1. Environment Setup
test_step "Environment Setup"
run_script "script/integration/AccountSetup.s.sol:AccountSetupScript"
run_script "script/integration/ContractDeployment.s.sol:ContractDeploymentScript"

# 2. Get System Parameters
test_step "Get System Parameters"
# Get parameters from contract or use .env defined values
EPOCH_DURATION=${EPOCH_DURATION:-$(cast call 0x000000000000000000000000000000000000f010 "getEpochDuration()(uint256)")}
UNBONDING_PERIOD=${UNBONDING_PERIOD:-$(cast call 0x000000000000000000000000000000000000F013 "getUnbondingPeriod()(uint256)")}

echo "System Parameters:"
echo "  Epoch Duration: $EPOCH_DURATION seconds"
echo "  Unbonding Period: $UNBONDING_PERIOD seconds"

# 3. Validator Registration
test_step "Validator Registration"
run_script "script/integration/ValidatorRegistration.s.sol:ValidatorRegistrationScript"
verify_result "Validator registered" "script/integration/ValidatorStatusCheck.s.sol:ValidatorStatusCheckScript"

# 4. Validator Resign
test_step "Validator Resign"
run_script "script/integration/ValidatorResign.s.sol:ValidatorResignScript"
verify_result "Validator resigned" "script/integration/ValidatorStatusCheck.s.sol:ValidatorStatusCheckScript"

# 5. Wait for Next Epoch
test_step "Wait for Next Epoch"
advance_time $EPOCH_DURATION
advance_blocks 1
verify_result "Epoch advanced" "script/integration/ValidatorStatusCheck.s.sol:ValidatorStatusCheckScript"

# 6. Validator Exit
test_step "Validator Exit"
run_script "script/integration/ValidatorExit.s.sol:ValidatorExitScript"
verify_result "Validator exited" "script/integration/ValidatorStatusCheck.s.sol:ValidatorStatusCheckScript"

# 7. Wait for Unbonding Period
test_step "Wait for Unbonding Period"
advance_time $UNBONDING_PERIOD
advance_blocks 1
verify_result "Unbonding period passed" "script/integration/ValidatorStatusCheck.s.sol:ValidatorStatusCheckScript"

# 8. Validator Withdraw Unbonded
test_step "Validator Withdraw Unbonded"
run_script "script/integration/ValidatorWithdrawUnbonded.s.sol:ValidatorWithdrawUnbondedScript"
verify_result "Validator withdrew unbonded funds" "script/integration/ValidatorStatusCheck.s.sol:ValidatorStatusCheckScript"

test_end