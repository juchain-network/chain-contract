// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {ValidatorUtils} from "../utils/ValidatorUtils.s.sol";
import {console} from "forge-std/Test.sol";

contract ValidatorNew is ValidatorUtils {
    function run() public override {
        console.log("Starting New Validator Tests...");
        
        // Create test accounts
        createTestAccounts();
        
        // Deploy and initialize contracts
        deployAndInitializeContracts();

        // Test 1: Validator Registration
        testValidatorRegistration();
        
        console.log("\nAll New Validator tests completed successfully!");
    }
    
    function testValidatorRegistration() internal {
        console.log("\n=== Testing Validator Registration ===");
        
        // Create a new validator account and private key
        uint256 newValidatorKey = getValidatorKey(4);
        address newValidator = getValidatorAddr(4);
        
        console.log("New validator address:", newValidator);
        console.log("Initial balance:", newValidator.balance / 1 ether, "ETH");
        
        // Create and pass proposal for new validator
        vm.startBroadcast(getValidatorKey(0));
        bytes32 proposalId = proposal.createProposal(newValidator, true, "Add new validator");
        vm.stopBroadcast();
        require(proposalId != bytes32(0), "Proposal creation failed");
        
        // No need to extract from logs
        // Convert bytes32 to string for logging (Solidity console doesn't support bytes32 directly)
        console.log("Proposal created with ID:", toHexString(proposalId));
        
        // Vote for the proposal from all validators
        for (uint256 i = 0; i < initialValidators; i++) {
            vm.startBroadcast(getValidatorKey(i));
            proposal.voteProposal(proposalId, true);
            vm.stopBroadcast();
        }
        
        // After vote passes, check top validators list doesn't include new validator
        console.log("Checking top validators list before registration...");
        address[] memory topValidatorsBefore = validators.getTopValidators();
        for (uint256 i = 0; i < topValidatorsBefore.length; i++) {
            require(topValidatorsBefore[i] != newValidator, "New validator should not be in top validators list before registration");
        }
        console.log(unicode"✓ Top validators list correctly doesn't include new validator before registration");
        
        // Register validator using the new validator's private key
        console.log("Registering validator:", newValidator);
        vm.startBroadcast(newValidatorKey);
        staking.registerValidator{value: initialStake}(2000); // 2000 = 20%
        vm.stopBroadcast();
        console.log("Validator registration transaction completed");
        
        // Verify registration
        console.log("Verifying validator registration...");
        (uint256 selfStake, , , , , , , , , ) = staking.getValidatorInfo(newValidator);
        require(selfStake == initialStake, "Validator should have correct self-stake");
        
        // After registration, check top validators list includes new validator
        console.log("Checking top validators list after registration...");
        address[] memory topValidatorsAfter = validators.getTopValidators();
        bool isInTopValidators = false;
        for (uint256 i = 0; i < topValidatorsAfter.length; i++) {
            if (topValidatorsAfter[i] == newValidator) {
                isInTopValidators = true;
                break;
            }
        }
        require(isInTopValidators, "New validator should be in top validators list after registration");
        console.log(unicode"✓ Top validators list correctly includes new validator after registration");
        
        // After registration, check active validators list doesn't include new validator
        console.log("Checking active validators list after registration...");
        address[] memory activeValidators = validators.getActiveValidators();
        for (uint256 i = 0; i < activeValidators.length; i++) {
            require(activeValidators[i] != newValidator, "New validator should not be in active validators list immediately after registration");
        }
        console.log(unicode"✓ Active validators list correctly doesn't include new validator immediately after registration");
        
        // Check new validator's commission rate equals set rate
        console.log("Checking validator commission rate...");
        (, , uint256 commissionRate, , , , , , , ) = staking.getValidatorInfo(newValidator);
        require(commissionRate == 2000, "Validator commission rate should equal set rate (2000 = 20%)");
        console.log(unicode"✓ Validator commission rate correctly set to 20%");
        
        console.log(unicode"✓ Validator registered successfully with stake:", selfStake / 1 ether, "ETH");
        console.log("Remaining balance:", newValidator.balance / 1 ether, "ETH");
    }
}
