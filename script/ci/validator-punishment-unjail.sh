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

test_name="Validator Punishment and Unjail"
test_start "$test_name"

# 1. Environment Setup
test_step "Environment Setup"
run_script "script/integration/AccountSetup.s.sol:AccountSetupScript"
run_script "script/integration/ContractDeployment.s.sol:ContractDeploymentScript"

# 2. Get System Parameters
test_step "Get System Parameters"
# Get parameters from contract or use .env defined values
VALIDATOR_UNJAIL_PERIOD=${VALIDATOR_UNJAIL_PERIOD:-$(cast call 0x000000000000000000000000000000000000F013 "getValidatorUnjailPeriod()(uint256)")}
echo "Validator Unjail Period: $VALIDATOR_UNJAIL_PERIOD seconds"

# 3. Punish Validator
test_step "Punish Validator"
run_script "script/integration/ValidatorPunish.s.sol:ValidatorPunishScript"
verify_result "Validator punished" "script/integration/ValidatorStatusCheck.s.sol:ValidatorStatusCheckScript"

# 4. Wait for Unjail Period
test_step "Wait for Unjail Period"
advance_time $VALIDATOR_UNJAIL_PERIOD
advance_blocks 1
verify_result "Unjail period passed" "script/integration/ValidatorStatusCheck.s.sol:ValidatorStatusCheckScript"

# 5. Validator Unjail
test_step "Validator Unjail"
run_script "script/integration/ValidatorUnjail.s.sol:ValidatorUnjailScript"
verify_result "Validator unjailed" "script/integration/ValidatorStatusCheck.s.sol:ValidatorStatusCheckScript"

test_end