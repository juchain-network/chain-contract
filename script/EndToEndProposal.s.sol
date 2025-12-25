// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {Script} from "lib/forge-std/src/Script.sol";
import {Test} from "lib/forge-std/src/Test.sol";
import {Proposal} from "../contracts/Proposal.sol";
import {Validators} from "../contracts/Validators.sol";

// End-to-end script: Create proposal + multiple validator voting + check results
contract EndToEndProposalScript is Script, Test {
    
    // Fixed system addresses (consistent with deployment)
    address constant VALIDATORS = 0x000000000000000000000000000000000000f000;
    address constant PUNISH = 0x000000000000000000000000000000000000F001;
    address constant PROPOSAL = 0x000000000000000000000000000000000000F002;
    address constant STAKING = 0x000000000000000000000000000000000000F003;
    
    event ProposalCreated(bytes32 indexed id, address proposer, address target, bool flag);
    event VoteCast(bytes32 indexed id, address voter, bool vote);
    event ProposalResult(bytes32 indexed id, bool passed, address[] topValidators);
    
    function run() external {
        // Example: End-to-end proposal workflow demonstration
        address testTarget = 0xa0Ee7A142d267C1f36714E4a8F75612F20a79720;
        
        // Create a simulated voter array (empty array due to permission restrictions)
        address[] memory voters = new address[](0);
        
        // Directly call internal function to execute validator addition process
        bool success = _runProposalFlow(
            testTarget,
            true, // Add validator
            "End-to-end test: Adding validator",
            voters
        );
        
        if (success) {
            emit ProposalResult(bytes32(uint256(uint160(testTarget))), true, Validators(VALIDATORS).getTopValidators());
        }
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
        bytes32 id = keccak256(abi.encodePacked(msg.sender, configId, newValue, timestamp));
        
        // Create configuration update proposal
        Proposal(PROPOSAL).createUpdateConfigProposal(configId, newValue);
        
        // Validator voting
        uint256 yesVotes = 0;
        for (uint i = 0; i < voters.length; i++) {
            // Check if active validator
            if (Validators(VALIDATORS).isActiveValidator(voters[i])) {
                // Here simplified handling, assuming all votes are yes
                // In actual use, a vote selection array can be passed in
                vm.prank(voters[i]);
                Proposal(PROPOSAL).voteProposal(id, true);
                yesVotes++;
                
                emit VoteCast(id, voters[i], true);
            }
        }
        
        // Check if passed (majority required)
        uint256 requiredVotes = Validators(VALIDATORS).getActiveValidators().length / 2 + 1;
        success = yesVotes >= requiredVotes;
        
        address[] memory topValidators = Validators(VALIDATORS).getTopValidators();
        emit ProposalResult(id, success, topValidators);
        
        return success;
    }
    
    function _runProposalFlow(
        address target,
        bool isAdd,
        string memory details,
        address[] memory voters
    ) internal returns (bool success) {
        // Create proposal and get real ID (via events)
        Proposal(PROPOSAL).createProposal(target, isAdd, details);
        
        // Note: In actual environments, we need to get the real proposal ID from event logs
        // Here for demonstration, we use a simplified ID calculation
        // The real ID should be obtained from LogCreateProposal events
        bytes32 id = keccak256(abi.encodePacked(msg.sender, target, isAdd, details, block.timestamp));
        emit ProposalCreated(id, msg.sender, target, isAdd);
        
        // Validator voting
        uint256 yesVotes = 0;
        for (uint i = 0; i < voters.length; i++) {
            // Check if active validator
            if (Validators(VALIDATORS).isActiveValidator(voters[i])) {
                // Here simplified handling, assuming all votes are yes
                vm.prank(voters[i]);
                Proposal(PROPOSAL).voteProposal(id, true);
                yesVotes++;
                
                emit VoteCast(id, voters[i], true);
            }
        }
        
        // Check if it passes (majority required)
        uint256 requiredVotes = Validators(VALIDATORS).getActiveValidators().length / 2 + 1;
        success = yesVotes >= requiredVotes;
        
        // Verify final status
        if (success) {
            if (isAdd) {
                require(Validators(VALIDATORS).isTopValidator(target), "Validator should be added");
                require(Proposal(PROPOSAL).pass(target), "Target should be marked as passed");
            } else {
                require(!Validators(VALIDATORS).isTopValidator(target), "Validator should be removed");
                require(!Proposal(PROPOSAL).pass(target), "Target should be marked as not passed");
            }
        }
        
        address[] memory topValidators = Validators(VALIDATORS).getTopValidators();
        emit ProposalResult(id, success, topValidators);
        
        return success;
    }
    
    // Convenience function: Get current active validator list for voting
    function getActiveValidators() external view returns (address[] memory) {
        return Validators(VALIDATORS).getActiveValidators();
    }
    
    // Convenience function: Get current top validator list
    function getTopValidators() external view returns (address[] memory) {
        return Validators(VALIDATORS).getTopValidators();
    }
}
