// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {ProposalUtils} from "../utils/ProposalUtils.s.sol";
import {console} from "forge-std/Test.sol";

// End-to-end script: Create proposal + multiple validator voting + check results
contract ProposalFlow is ProposalUtils {
    

    
    event ProposalCreated(bytes32 indexed id, address proposer, address target, bool flag);
    event VoteCast(bytes32 indexed id, address voter, bool vote);
    event ProposalResult(bytes32 indexed id, bool passed, address[] topValidators);
    
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
    ) public returns (bool success) {
        // Create configuration update proposal
        vm.startBroadcast(validatorKeys[0]);
        bytes32 id = proposal.createUpdateConfigProposal(configId, newValue);
        vm.stopBroadcast();
        
        // Validator voting
        uint256 yesVotes = 0;
        for (uint i = 0; i < voters.length; i++) {
            // Check if active validator
            if (validators.isActiveValidator(voters[i])) {
                // All votes are yes for demonstration purposes
                vm.startBroadcast(getValidatorKey(i));
                proposal.voteProposal(id, true);
                vm.stopBroadcast();
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
    
    function run() public override {
        console.log("Starting Governance Parameter Update Tests...");
        
        // Create test accounts
        createTestAccounts();
        
        // Deploy and initialize contracts
        deployAndInitializeContracts();
        
        // Use initial validators as voters
        address[] memory voters = validatorAccounts;
        
        // Set a random validator as temporary miner
        address miner = validatorAccounts[0];
        setMinerTemporarily(miner);
        
        // Define new values for each governance parameter
        // Each config ID corresponds to a parameter in Proposal.sol
        // Config IDs: 0-9 (10 parameters total)
        uint256[] memory newValues = new uint256[](10);
        newValues[0] = 3600;          // 0: proposalLastingPeriod (blocks)
        newValues[1] = 30;             // 1: punishThreshold (blocks)
        newValues[2] = 60;             // 2: removeThreshold (blocks)
        newValues[3] = 50;             // 3: decreaseRate (percentage * 100)
        newValues[4] = 86400;          // 4: withdrawProfitPeriod (blocks)
        newValues[5] = 0.1 ether;      // 5: blockReward (wei)
        newValues[6] = 7200;           // 6: unbondingPeriod (blocks)
        newValues[7] = 1000;           // 7: validatorUnjailPeriod (blocks)
        newValues[8] = 200000 ether;   // 8: minValidatorStake (wei)
        newValues[9] = 25;             // 9: maxValidators (number)
        
        // Update each governance parameter through proposal
        for (uint256 i = 0; i < 10; i++) {
            console.log("\n=== Updating Governance Parameter %d ===", i);
            bool success = runConfigUpdateFlow(i, newValues[i], voters);
            require(success, string(abi.encodePacked("Failed to update parameter ", vm.toString(i))));
            console.log("Parameter %d updated successfully!", i);
        }
        
        console.log("\nAll Governance Parameter Update Tests completed successfully!");
    }
    
    function _runProposalFlow(
        address target,
        bool isAdd,
        string memory details,
        address[] memory voters
    ) internal returns (bool success) {
        // Create proposal from first validator
        vm.startBroadcast(validatorKeys[0]);
        proposal.createProposal(target, isAdd, details);
        vm.stopBroadcast();
        
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
                vm.startBroadcast(getValidatorKey(i));
                proposal.voteProposal(id, true);
                vm.stopBroadcast();
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
                vm.startBroadcast(uint256(keccak256(abi.encodePacked("newValidatorEndToEnd"))));
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
