// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseTestScript} from "./BaseTestScript.s.sol";
import {console} from "forge-std/Test.sol";

// End-to-end script: Create proposal + multiple validator voting + check results
contract EndToEndProposalScript is BaseTestScript {
    

    
    event ProposalCreated(bytes32 indexed id, address proposer, address target, bool flag);
    event VoteCast(bytes32 indexed id, address voter, bool vote);
    event ProposalResult(bytes32 indexed id, bool passed, address[] topValidators);
    
    function run() public override {
        console.log("Starting End-to-End Proposal Tests...");
        
        // Create test accounts
        createTestAccounts();
        
        // Deploy and initialize contracts
        deployAndInitializeContracts();
        
        // Create a new validator for testing
        address newValidator = fundNewValidator(uint256(keccak256(abi.encodePacked("newValidatorEndToEnd"))));
        
        // Use initial validators as voters
        address[] memory voters = validatorAccounts;
        
        // Set a random validator as temporary miner
        address miner = validatorAccounts[0];
        setMinerTemporarily(miner);
        vm.deal(miner, 1000 ether);
        
        // Directly call internal function to execute validator addition process
        bool success = _runProposalFlow(
            newValidator,
            true, // Add validator
            "End-to-end test: Adding validator",
            voters
        );
        
        if (success) {
            emit ProposalResult(bytes32(uint256(uint160(newValidator))), true, validators.getTopValidators());
        }
        
        console.log("End-to-End Proposal Tests completed!");
    }
    
    struct ProposalInfo {
        bytes32 id;
        address proposer;
        address target;
        bool flag;
        string details;
    }
    
    function runAddValidatorFlow(
        address newValidator,
        string memory details,
        address[] memory voters
    ) external returns (bool success) {
        return _runProposalFlow(newValidator, true, details, voters);
    }
    
    function runRemoveValidatorFlow(
        address targetValidator, 
        string memory details,
        address[] memory voters
    ) external returns (bool success) {
        return _runProposalFlow(targetValidator, false, details, voters);
    }
    
    function runConfigUpdateFlow(
        uint256 configId,
        uint256 newValue,
        address[] memory voters
    ) external returns (bool success) {
        // Freeze timestamp to ensure deterministic ID
        uint256 timestamp = block.timestamp;
        bytes32 id = keccak256(abi.encodePacked(validatorAccounts[0], configId, newValue, timestamp));
        
        // Create configuration update proposal
        vm.prank(validatorAccounts[0]);
        proposal.createUpdateConfigProposal(configId, newValue);
        
        // Validator voting
        uint256 yesVotes = 0;
        for (uint i = 0; i < voters.length; i++) {
            // Check if active validator
            if (validators.isActiveValidator(voters[i])) {
                // All votes are yes for demonstration purposes
                vm.prank(voters[i]);
                proposal.voteProposal(id, true);
                yesVotes++;
                
                emit VoteCast(id, voters[i], true);
            }
        }
        
        // Check if passed (majority required)
        uint256 requiredVotes = validators.getActiveValidators().length / 2 + 1;
        success = yesVotes >= requiredVotes;
        
        address[] memory topValidators = validators.getTopValidators();
        emit ProposalResult(id, success, topValidators);
        
        return success;
    }
    
    function _runProposalFlow(
        address target,
        bool isAdd,
        string memory details,
        address[] memory voters
    ) internal returns (bool success) {
        // Create proposal from first validator
        vm.prank(validatorAccounts[0]);
        proposal.createProposal(target, isAdd, details);
        
        // Note: In actual environments, we need to get the real proposal ID from event logs
        // Here for demonstration, we use a simplified ID calculation
        // The real ID should be obtained from LogCreateProposal events
        bytes32 id = keccak256(abi.encodePacked(validatorAccounts[0], target, isAdd, details, block.timestamp));
        emit ProposalCreated(id, validatorAccounts[0], target, isAdd);
        
        // Validator voting
        uint256 yesVotes = 0;
        for (uint i = 0; i < voters.length; i++) {
            // Check if active validator
            if (validators.isActiveValidator(voters[i])) {
                // Here simplified handling, assuming all votes are yes
                vm.prank(voters[i]);
                proposal.voteProposal(id, true);
                yesVotes++;
                
                emit VoteCast(id, voters[i], true);
            }
        }
        
        // Check if it passes (majority required)
        uint256 requiredVotes = validators.getActiveValidators().length / 2 + 1;
        success = yesVotes >= requiredVotes;
        
        // Verify final status
        if (success) {
            if (isAdd) {
                require(proposal.pass(target), "Target should be marked as passed");
                
                // After proposal passes, validator needs to stake to become active
                vm.startBroadcast(target);
                staking.registerValidator{value: proposal.minValidatorStake()}(100); // 1% commission rate
                vm.stopBroadcast();
                
                require(validators.isTopValidator(target), "Validator should be added to top validators");
            } else {
                require(!validators.isTopValidator(target), "Validator should be removed");
                require(!proposal.pass(target), "Target should be marked as not passed");
            }
        }
        
        address[] memory topValidators = validators.getTopValidators();
        emit ProposalResult(id, success, topValidators);
        
        return success;
    }
    
    // Convenience function: Get current active validator list for voting
    function getActiveValidators() external view returns (address[] memory) {
        return validators.getActiveValidators();
    }
    
    // Convenience function: Get current top validator list
    function getTopValidators() external view returns (address[] memory) {
        return validators.getTopValidators();
    }
}
