// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseTestScript} from "./BaseTestScript.s.sol";
import {console} from "forge-std/Test.sol";

contract ValidatorManagementScript is BaseTestScript {
    function run() public override {
        console.log("Starting Validator Management Tests...");
        
        // Create test accounts
        createTestAccounts();
        
        // Deploy and initialize contracts
        deployAndInitializeContracts();
        
        // Test 1: Validator Registration
        testValidatorRegistration();
        
        console.log("\nAll Validator Management tests completed successfully!");
    }
    
    function testValidatorRegistration() internal {
        console.log("\n=== Testing Validator Registration ===");
        
        // Create a new validator account and private key
        uint256 newValidatorKey = uint256(keccak256(abi.encodePacked("testValidator1")));
        address newValidator = fundNewValidator(newValidatorKey); // Use base class method to fund new validator
        
        console.log("New validator address:", newValidator);
        console.log("Initial balance:", newValidator.balance / 1 ether, "ETH");
        
        // Create and pass proposal for new validator
        // Record logs to capture the proposal ID from the event
        vm.recordLogs();
        
        // vm.prank(validatorAccounts[0]);
        vm.startBroadcast(validatorKeys[0]);
        bytes32 proposalId = proposal.createProposal(newValidator, true, "Add new validator");
        vm.stopBroadcast();
        require(proposalId != bytes32(0), "Proposal creation failed");
        
        // Proposal ID is already obtained from the createProposal return value
        // No need to extract from logs
        // Convert bytes32 to string for logging (Solidity console doesn't support bytes32 directly)
        console.log("Proposal created with ID:", toHexString(proposalId));
        
        // Vote for the proposal from all validators
        for (uint256 i = 0; i < validatorAccounts.length; i++) {
            // vm.prank(validatorAccounts[i]);
            vm.startBroadcast(validatorKeys[i]);
            proposal.voteProposal(proposalId, true);
            vm.stopBroadcast();
        }
        
        // Simulate miner behavior when needed (setMiner should be called before reward distribution)
        // Example: If we were distributing rewards, we would set the miner temporarily here
        // setMinerTemporarily(newValidator);
        
        // Register validator using the new validator's private key
        console.log("Registering validator:", newValidator);
        vm.startBroadcast(newValidatorKey);
        staking.registerValidator{value: INITIAL_STAKE}(1000); // 1000 = 10%
        vm.stopBroadcast();
        console.log("Validator registration transaction completed");
        
        // Verify registration
        console.log("Verifying validator registration...");
        (uint256 selfStake, , , , , , , ) = staking.getValidatorInfo(newValidator);
        require(selfStake == INITIAL_STAKE, "Validator should have correct self-stake");
        
        console.log(unicode"✓ Validator registered successfully with stake:", selfStake / 1 ether, "ETH");
        console.log("Remaining balance:", newValidator.balance / 1 ether, "ETH");
    }
}
