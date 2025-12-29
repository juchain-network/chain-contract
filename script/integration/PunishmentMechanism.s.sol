// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseTestScript} from "./BaseTestScript.s.sol";
import {console} from "forge-std/Test.sol";

contract PunishmentMechanismScript is BaseTestScript {
    // Additional configuration specific to this test
    uint256 public constant PUNISH_THRESHOLD = 24;
    
    // Test accounts specific to this test
    address public testValidator;
    
    function run() public override {
        console.log("Starting Punishment Mechanism Tests...");
        
        // Create test accounts using base class method
        createTestAccounts();
        
        // Deploy and initialize contracts using base class method
        deployAndInitializeContracts();
        
        // Create and fund a test validator for punishment tests
        uint256 testValidatorKey = uint256(keccak256(abi.encodePacked("punishTestValidator")));
        testValidator = fundNewValidator(testValidatorKey); // Use base class method to fund new validator
        
        // Register and activate validator
        registerAndActivateValidator();
        
        // Test 1: Missed Blocks Recording
        testMissedBlocksRecording();
        
        console.log("\nAll Punishment Mechanism tests completed successfully!");
    }
    

    
    function registerAndActivateValidator() internal {
        console.log("\nRegistering and activating validator...");
        
        // Create and pass proposal for test validator
        vm.startBroadcast(validatorKeys[0]);
        // Create proposal
        bytes32 proposalId = proposal.createProposal(testValidator, true, "Add validator for punishment test");
        require(proposalId != bytes32(0), "Proposal creation failed");
        vm.stopBroadcast();
        
        // Vote for the proposal from all validators
        for (uint256 i = 0; i < validatorAccounts.length; i++) {
            vm.startBroadcast(validatorKeys[i]);
            proposal.voteProposal(proposalId, true);
            vm.stopBroadcast();
        }
        
        // Get private key for the new validator
        uint256 testValidatorKey = uint256(keccak256(abi.encodePacked("punishTestValidator")));
        
        // Register validator with commission rate (10%)
        vm.startBroadcast(testValidatorKey);
        staking.registerValidator{value: INITIAL_STAKE}(1000); // 1000 = 10%
        vm.stopBroadcast();
        
        console.log(unicode"✓ Validator registered successfully");
    }
    
    function testMissedBlocksRecording() internal view {
        console.log("\n=== Testing Missed Blocks Recording ===");
        
        // Test: Check initial punish record is 0
        uint256 initialMissedBlocks = punish.getPunishRecord(testValidator);
        require(initialMissedBlocks == 0, "Initial missed blocks should be 0");
        
        console.log(unicode"✓ Initial punish record checked successfully");
        
        // Test: Check that validator is not jailed initially
        bool isJailed = staking.isValidatorJailed(testValidator);
        require(isJailed == false, "Validator should not be jailed initially");
        
        console.log(unicode"✓ Initial jail status checked successfully");
    }
}
