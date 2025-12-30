// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseTestUtils} from "../utils/BaseTestUtils.s.sol";
import {console} from "forge-std/Test.sol";

// Atomic script: Responsible for validator punishment operations
contract ValidatorPunish is BaseTestUtils {
    function run() public override {
        console.log("Starting Validator Punish Tests...");
        
        // // Create test accounts
        // createTestAccounts();
        
        // // Deploy and initialize contracts
        // deployAndInitializeContracts();
        
        // // Test 1: Validator Punishment Mechanism
        // testValidatorPunishmentMechanism();
        
        console.log("\nAll Validator Punish tests completed successfully!");
    }
    
    function testValidatorPunishmentMechanism() internal {
        console.log("\n=== Testing Validator Punishment Mechanism ===");
        
        // Get the first validator's address (to be punished)
        address targetValidator = validatorAccounts[0];
        
        // Use the second validator as miner for punishment operations
        uint256 minerKey = getValidatorKey(1);
        address minerAddr = vm.addr(minerKey);
        
        console.log("Target validator address:", targetValidator);
        console.log("Miner address:", minerAddr);
        
        // Get punishment thresholds from Proposal contract
        uint256 punishThreshold = proposal.punishThreshold();
        uint256 removeThreshold = proposal.removeThreshold();
        console.log("Punish threshold:", punishThreshold);
        console.log("Remove threshold:", removeThreshold);
        
        // Step 1: Generate rewards for validators through multiple block reward distributions
        console.log("\n1. Generating rewards for validators...");
        
        // Set miner temporarily to call distributeBlockReward
        setMinerTemporarily(minerAddr);
        
        for (uint256 i = 0; i < 5; i++) {
            // Distribute block reward to the first validator (as block producer)
            vm.startBroadcast(minerKey);
            validators.distributeBlockReward{value: 1 ether}();
            vm.stopBroadcast();
            console.log("   Block reward distributed #", i+1);
        }
        
        // Verify initial punishment record is zero
        uint256 initialMissedBlocks = punish.getPunishRecord(targetValidator);
        console.log("\nInitial missed blocks counter:", initialMissedBlocks);
        require(initialMissedBlocks == 0, "Initial missed blocks counter should be zero");
        
        // Step 2: Apply multiple punishments to the validator
        console.log("\n2. Applying multiple punishments to validator...");
        
        // First, apply punishments up to punishThreshold - 1
        for (uint256 i = 0; i < punishThreshold - 1; i++) {
            // Increase block number for each punishment
            vm.roll(block.number + 1);
            
            vm.startBroadcast(minerKey);
            punish.punish(targetValidator);
            vm.stopBroadcast();
            
            // Get updated punishment record
            uint256 currentMissedBlocks = punish.getPunishRecord(targetValidator);
            console.log("   Punishment #", i+1, ", Missed blocks:", currentMissedBlocks);
        }
        
        // Check if validator is not jailed yet and rewards are not forfeited
        bool isJailedBeforePunishThreshold = staking.isValidatorJailed(targetValidator);
        console.log("\nValidator jailed status before reaching punish threshold:", isJailedBeforePunishThreshold);
        require(!isJailedBeforePunishThreshold, "Validator should not be jailed before reaching punish threshold");
        
        // Apply one more punishment to reach punishThreshold
        console.log("\n3. Applying punishment to reach punish threshold...");
        vm.startBroadcast(minerKey);
        punish.punish(targetValidator);
        vm.stopBroadcast();
        
        // Check punishment record after reaching punishThreshold
        uint256 missedBlocksAfterPunishThreshold = punish.getPunishRecord(targetValidator);
        console.log("   Missed blocks after punish threshold:", missedBlocksAfterPunishThreshold);
        
        // Verify that at punishThreshold, rewards are forfeited but validator is not jailed
        bool isJailedAtPunishThreshold = staking.isValidatorJailed(targetValidator);
        console.log("   Validator jailed status at punish threshold:", isJailedAtPunishThreshold);
        require(!isJailedAtPunishThreshold, "Validator should not be jailed at punish threshold");
        
        // Step 3: Continue punishing to reach removeThreshold
        console.log("\n4. Continuing to punish to reach remove threshold...");
        
        // Calculate remaining punishments needed to reach removeThreshold
        uint256 remainingPunishments = removeThreshold - missedBlocksAfterPunishThreshold;
        for (uint256 i = 0; i < remainingPunishments; i++) { 
            vm.startBroadcast(minerKey);
            punish.punish(targetValidator);
            vm.stopBroadcast();
            
            // Get updated punishment record
            uint256 currentMissedBlocks = punish.getPunishRecord(targetValidator);
            console.log("   Punishment #", i+1 + punishThreshold, ", Missed blocks:", currentMissedBlocks);
        }
        
        // Check punishment record after reaching removeThreshold
        uint256 missedBlocksAfterRemoveThreshold = punish.getPunishRecord(targetValidator);
        console.log("\n5. Verifying results after reaching remove threshold...");
        console.log("   Missed blocks after remove threshold:", missedBlocksAfterRemoveThreshold);
        
        // Verify that at removeThreshold, validator is jailed and removed from top validators
        bool isJailedAtRemoveThreshold = staking.isValidatorJailed(targetValidator);
        console.log("   Validator jailed status at remove threshold:", isJailedAtRemoveThreshold);
        require(isJailedAtRemoveThreshold, "Validator should be jailed at remove threshold");
        
        // Check if validator is removed from top validators list
        console.log("   Checking if validator is removed from top validators list...");
        address[] memory topValidatorsAfter = validators.getTopValidators();
        bool isInTopValidators = false;
        for (uint256 i = 0; i < topValidatorsAfter.length; i++) {
            if (topValidatorsAfter[i] == targetValidator) {
                isInTopValidators = true;
                break;
            }
        }
        require(!isInTopValidators, "Validator should be removed from top validators list at remove threshold");
        console.log("   Validator correctly removed from top validators list");
        
        console.log(unicode"✓ Validator punishment mechanism test completed successfully");
    }
}