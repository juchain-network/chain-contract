// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {StakingUtils} from "../utils/StakingUtils.s.sol";
import {DelegatorUtils} from "../utils/DelegatorUtils.s.sol";
import {console} from "forge-std/Test.sol";

contract StakingDelegating is StakingUtils, DelegatorUtils {
    function run() public override {
        console.log("Starting Staking Mechanism Tests...");
        
        // Create test accounts using base class method
        createTestAccounts();
        
        // Deploy and initialize contracts using base class method
        deployAndInitializeContracts();
        
        // Get test accounts
        address validator1 = validatorAccounts[0];
        uint256 validator1Key = validatorKeys[0];
        address validator2 = validatorAccounts[1];
        address delegator1 = delegatorAccounts[0];
        uint256 delegator1Key = delegatorKeys[0];
        
        // Local variable for staking amount
        uint256 stakingAmount = 1000 ether;
        
        console.log("\n=== Test Environment Setup ===");
        console.log("Validator 1:", validator1);
        console.log("Validator 2:", validator2);
        console.log("Delegator 1:", delegator1);
        
        // Test 1: Validator 1 increases stake by 1000
        console.log("\n=== Test 1: Validator 1 increases stake by 1000 ===");
        addStake(validator1Key, validator1, stakingAmount);
        
        // Test 2: Validator 1 decreases stake by 1000
        console.log("\n=== Test 2: Validator 1 decreases stake by 1000 ===");
        decreaseStake(validator1Key, validator1, stakingAmount);
        
        // Test 3: Validator 1 decreases stake by 50000, should fail
        console.log("\n=== Test 3: Validator 1 decreases stake by 50000, should fail ===");
        vm.expectRevert();
        vm.prank(validator1);
        staking.decreaseValidatorStake(50000 ether);
        console.log(unicode"✓ Failed to decrease stake by 50000 as expected");
        
        // Test 4: Delegator 1 delegates 1000 to Validator 1
        console.log("\n=== Test 4: Delegator 1 delegates 1000 to Validator 1 ===");
        delegate(delegator1Key, validator1, stakingAmount);
        
        // Test 5: Delegator 1 delegates 1000 to Validator 2
        console.log("\n=== Test 5: Delegator 1 delegates 1000 to Validator 2 ===");
        delegate(delegator1Key, validator2, stakingAmount);
        
        // Test 6: Delegator 1 undelegates 2000 from Validator 1, should fail
        console.log("\n=== Test 6: Delegator 1 undelegates 2000 from Validator 1, should fail ===");
        vm.expectRevert();
        vm.prank(delegator1);
        staking.undelegate(validator1, 2000 ether);
        console.log(unicode"✓ Failed to undelegate 2000 as expected");
        
        // Test 7: Delegator 1 undelegates 1000 from Validator 1
        console.log("\n=== Test 7: Delegator 1 undelegates 1000 from Validator 1 ===");
        undelegate(delegator1Key, validator1, stakingAmount);
        
        // Test 8: Delegator 1 undelegates 200 from Validator 2
        console.log("\n=== Test 8: Delegator 1 undelegates 200 from Validator 2 ===");
        undelegate(delegator1Key, validator2, 200 ether);
        
        // Results verification
        console.log("\n=== Results Verification ===");
        
        // Verify Validator 1's stake
        console.log("\n1. Validator 1 Stake Verification:");
        (uint256 validator1Stake, , , , , , , , , ) = staking.getValidatorInfo(validator1);
        console.log("   Validator 1 current stake:", validator1Stake / 1 ether, "ETH");
        require(validator1Stake == 100000 ether, "Validator 1 stake should be 100000 ETH");
        
        // Verify Delegator 1's delegation to Validator 1
        console.log("\n2. Delegator 1 Delegation to Validator 1:");
        (uint256 delegator1to1Stake, ) = staking.delegations(delegator1, validator1);
        console.log("   Delegated stake:", delegator1to1Stake / 1 ether, "ETH");
        require(delegator1to1Stake == 0 ether, "Delegator 1 should have no stake left with Validator 1");
        
        // Verify Delegator 1's delegation to Validator 2
        console.log("\n3. Delegator 1 Delegation to Validator 2:");
        (uint256 delegator1to2Stake, ) = staking.delegations(delegator1, validator2);
        console.log("   Delegated stake:", delegator1to2Stake / 1 ether, "ETH");
        require(delegator1to2Stake == 800 ether, "Delegator 1 should have 800 ETH left with Validator 2");
        
        console.log("\nAll Staking Mechanism tests completed successfully!");
    }
}
