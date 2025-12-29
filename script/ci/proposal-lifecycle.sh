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

test_name="Proposal Lifecycle"
test_start "$test_name"

# 1. Environment Setup
test_step "Environment Setup"
run_script "script/integration/AccountSetup.s.sol:AccountSetupScript"
run_script "script/integration/ContractDeployment.s.sol:ContractDeploymentScript"


# 2. Get System Parameters
test_step "Get System Parameters"
# Get parameters from contract or use .env defined values
PROPOSAL_LASTING_PERIOD=${PROPOSAL_LASTING_PERIOD:-$(cast call 0x000000000000000000000000000000000000F012 "getProposalLastingPeriod()(uint256)")}
echo "Proposal Lasting Period: $PROPOSAL_LASTING_PERIOD seconds"

# 3. Create Proposal
test_step "Create Proposal"
run_script "script/integration/ProposalCreate.s.sol:ProposalCreateScript"
verify_result "Proposal created" "script/integration/ProposalStatusCheck.s.sol:ProposalStatusCheckScript"

# 4. Vote on Proposal
test_step "Vote on Proposal"
run_script "script/integration/ProposalVote.s.sol:ProposalVoteScript"
verify_result "Voted on proposal" "script/integration/ProposalStatusCheck.s.sol:ProposalStatusCheckScript"

# 5. Execute Proposal
test_step "Execute Proposal"
run_script "script/integration/ProposalExecute.s.sol:ProposalExecuteScript"
verify_result "Proposal executed" "script/integration/ProposalStatusCheck.s.sol:ProposalStatusCheckScript"

# 6. Wait for Proposal Expiry
test_step "Wait for Proposal Expiry"
run_script "script/integration/ProposalExpireCheck.s.sol:ProposalExpireCheckScript"
advance_time $PROPOSAL_LASTING_PERIOD
advance_blocks 1
verify_result "Proposal expired" "script/integration/ProposalStatusCheck.s.sol:ProposalStatusCheckScript"

test_end