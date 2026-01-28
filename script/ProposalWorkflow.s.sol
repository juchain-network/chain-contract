// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseSetup} from "../test/BaseSetup.t.sol";
import {Proposal} from "../../contracts/Proposal.sol";
import {Validators} from "../../contracts/Validators.sol";

/**
 * @title ProposalWorkflow
 * @dev Proposal workflow script, demonstrating proposal creation and voting process
 */
contract ProposalWorkflow is BaseSetup {
    
    event WorkflowEvent(string message, address addr, bool result);
    
    function run() external {
        // Demonstrate proposal creation
        address newValidator = 0xa0Ee7A142d267C1f36714E4a8F75612F20a79720;
        
        // 1. Check initial status
        bool isValidatorBefore = Validators(VALIDATORS).isTopValidator(newValidator);
        emit WorkflowEvent("Initial validator status", newValidator, isValidatorBefore);
        
        // 2. Create proposal to add validator
        bytes32 proposalId = Proposal(PROPOSAL).createProposal(newValidator, true, "Workflow: Adding new validator");
        bool createResult = proposalId != bytes32(0);
        emit WorkflowEvent("Proposal creation", newValidator, createResult);
        
        // 3. Check if proposer is a validator
        bool isProposerValidator = Validators(VALIDATORS).isActiveValidator(msg.sender);
        emit WorkflowEvent("Proposer is validator", msg.sender, isProposerValidator);
        
        // 4. Check system status
        address[] memory activeValidators = Validators(VALIDATORS).getActiveValidators();
        emit WorkflowEvent("Active validators count", address(uint160(activeValidators.length)), true);
        
        // 5. Check proposal pass status
        bool passStatus = Proposal(PROPOSAL).pass(newValidator);
        emit WorkflowEvent("Target validator pass status", newValidator, passStatus);
    }
    
    function runConfigWorkflow() external {
        // Demonstrate configuration update proposal
        uint256 configId = 2; // 
        uint256 newValue = 5000; // 
        
        // Create configuration update proposal
        bytes32 proposalId = Proposal(PROPOSAL).createUpdateConfigProposal(configId, newValue);
        bool createResult = proposalId != bytes32(0);
        // casting to 'uint160' is safe because configId is a small configuration ID (typically 1-100)
        // forge-lint: disable-next-line(unsafe-typecast)
        emit WorkflowEvent("Config proposal creation", address(uint160(configId)), createResult);
        
        // Check current configuration
        uint256 currentLastingPeriod = Proposal(PROPOSAL).proposalLastingPeriod();
        // casting to 'uint160' is safe because currentLastingPeriod is a time duration in seconds (well within uint160 limits)
        // forge-lint: disable-next-line(unsafe-typecast)
        emit WorkflowEvent("Current lasting period", address(uint160(currentLastingPeriod)), true);
    }
}
