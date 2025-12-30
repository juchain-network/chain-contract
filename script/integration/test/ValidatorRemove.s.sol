// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {ValidatorUtils} from "../utils/ValidatorUtils.s.sol";
import {console} from "forge-std/Test.sol";

contract ValidatorRemove is ValidatorUtils {
    function run() public override {
        console.log("Starting Remove Validator Tests...");
        
        // Create test accounts
        createTestAccounts();
        
        // Deploy and initialize contracts
        deployAndInitializeContracts();
        
        // Test 1: Validator Removal
        testValidatorRemoval();
        
        console.log("\nAll Remove Validator tests completed successfully!");
    }
    
    function testValidatorRemoval() internal {
        console.log("\n=== Testing Validator Removal ===");
        
        // Get the 4th validator's address
        address validatorAddr = getValidatorAddr(3);
        
        console.log("Validator to remove address:", validatorAddr);
        console.log("Validator initial balance:", validatorAddr.balance / 1 ether, "ETH");
        
        // Verify the validator is in the initial set
        console.log("Verifying validator is in initial set...");
        (uint256 selfStakeBefore, , , , , , , , ) = staking.getValidatorInfo(validatorAddr);
        require(selfStakeBefore > 0, "Validator should be in initial set with stake");
        console.log(unicode"✓ Validator is in initial set with stake:", selfStakeBefore / 1 ether, "ETH");
        
        // Create and pass proposal to remove the validator
        vm.startBroadcast(getValidatorKey(0));
        bytes32 proposalId = proposal.createProposal(validatorAddr, false, "Remove validator");
        vm.stopBroadcast();
        require(proposalId != bytes32(0), "Proposal creation failed");
        
        // Convert bytes32 to string for logging
        console.log("Proposal created with ID:", toHexString(proposalId));
        
        // Vote for the proposal from all validators except the one being removed
        for (uint256 i = 0; i < initialValidators; i++) {
            if (i != 3) { // Skip the validator being removed
                vm.startBroadcast(getValidatorKey(i));
                proposal.voteProposal(proposalId, true);
                vm.stopBroadcast();
            }
        }
        
        // Verify the proposal passed
        console.log("Verifying proposal passed...");
        // Note: The proposal.pass() function checks if the proposal passed
        // For removal proposals, it should return false after passing
        require(!proposal.pass(validatorAddr), "Proposal should have passed for removal");
        console.log(unicode"✓ Proposal passed successfully");
        
        // Remove the validator
        console.log("Removing validator:", validatorAddr);
        // The removal should be handled by the proposal passing, no need for explicit call
        // But we can verify the validator is no longer in the active set
        
        // Verify the validator is no longer in top validators list
        console.log("Checking top validators list after removal...");
        address[] memory topValidatorsAfter = validators.getTopValidators();
        bool isInTopValidators = false;
        for (uint256 i = 0; i < topValidatorsAfter.length; i++) {
            if (topValidatorsAfter[i] == validatorAddr) {
                isInTopValidators = true;
                break;
            }
        }
        require(!isInTopValidators, "Validator should not be in top validators list after removal");
        console.log(unicode"✓ Validator correctly removed from top validators list");
        
        // Verify the validator is jailed
        console.log("Checking validator jail status after removal...");
        (, bool isJailed) = staking.getValidatorStatus(validatorAddr);
        require(isJailed, "Validator should be jailed after removal");
        console.log(unicode"✓ Validator correctly marked as jailed after removal");
        
        console.log(unicode"✓ Validator removal test completed successfully");
    }
}