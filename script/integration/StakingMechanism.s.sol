// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseTestScript} from "./BaseTestScript.s.sol";
import {console} from "forge-std/Test.sol";

contract StakingMechanismScript is BaseTestScript {
    // Configuration
    uint256 public constant VALIDATOR_STAKE = 100000 ether;
    uint256 public constant DELEGATION_AMOUNT = 200 ether;
    
    // Test accounts
    address public validator;
    address public delegator;
    
    function run() public override {
        console.log("Starting Staking Mechanism Tests...");
        
        // Create test accounts using base class method
        createTestAccounts();
        
        // Deploy and initialize contracts using base class method
        deployAndInitializeContracts();
        
        // Test 1: Validator Self-Staking
        testValidatorSelfStaking();
        
        // Test 2: Delegation to Validator
        testDelegation();
        
        // Test 3: Increase Delegation
        testIncreaseDelegation();
        
        // Test 4: Unstaking Delegation
        testUnstaking();
        
        console.log("\nAll Staking Mechanism tests completed successfully!");
    }
    
    function createTestAccounts() internal override {
        // Call base class method to create initial validators
        super.createTestAccounts();
        
        // Create test accounts for this script
        validator = vm.addr(uint256(keccak256(abi.encodePacked("testStakingValidator"))));
        delegator = vm.addr(uint256(keccak256(abi.encodePacked("testDelegator"))));
        
        // Fund new validator with 110k ETH (100k stake + 10k fees)
        fundNewValidator(uint256(keccak256(abi.encodePacked("testStakingValidator"))));
        vm.deal(delegator, 5000 ether);
        
        console.log("Additional test accounts created:");
        console.log("Validator:", validator);
        console.log("Delegator:", delegator);
    }
    

    
    function testValidatorSelfStaking() internal {
        console.log("\n=== Testing Validator Self-Staking ===");
        
        // Create and pass proposal for new validator
        vm.prank(validatorAccounts[0]);
        
        // Create proposal
        bytes32 proposalId = proposal.createProposal(validator, true, "Add validator for staking test");
        require(proposalId != bytes32(0), "Proposal creation failed");
        
        // Vote for the proposal from all validators
        for (uint256 i = 0; i < validatorAccounts.length; i++) {
            vm.prank(validatorAccounts[i]);
            proposal.voteProposal(proposalId, true);
        }
        
        // Register validator with self-stake and commission rate (10%)
        vm.startBroadcast(validator);
        staking.registerValidator{value: VALIDATOR_STAKE}(1000); // 1000 = 10%
        vm.stopBroadcast();
        
        // Verify self-stake
        (uint256 selfStake, , , , , , , ) = staking.getValidatorInfo(validator);
        require(selfStake == VALIDATOR_STAKE, "Validator should have correct self-stake");
        
        console.log(unicode"✓ Validator self-staked successfully:", selfStake / 1 ether, "ETH");
    }
    
    function testDelegation() internal {
        console.log("\n=== Testing Delegation to Validator ===");
        
        uint256 initialDelegatorBalance = delegator.balance;
        
        // Delegate to validator
        vm.startBroadcast(delegator);
        staking.delegate{value: DELEGATION_AMOUNT}(validator);
        vm.stopBroadcast();
        
        // Verify delegation
        (uint256 delegatedAmount, , , ) = staking.getDelegationInfo(delegator, validator);
        require(delegatedAmount == DELEGATION_AMOUNT, "Delegation amount should be correct");
        
        // Verify balance decrease
        uint256 expectedBalance = initialDelegatorBalance - DELEGATION_AMOUNT;
        require(delegator.balance == expectedBalance, "Delegator balance should decrease by delegation amount");
        
        console.log(unicode"✓ Delegation successful:", delegatedAmount / 1 ether, "ETH");
    }
    
    function testIncreaseDelegation() internal {
        console.log("\n=== Testing Increase Delegation ===");
        
        uint256 additionalDelegation = 100 ether;
        (uint256 initialDelegatedAmount, , , ) = staking.getDelegationInfo(delegator, validator);
        
        // Increase delegation
        vm.startBroadcast(delegator);
        staking.delegate{value: additionalDelegation}(validator);
        vm.stopBroadcast();
        
        // Verify increased delegation
        (uint256 newDelegatedAmount, , , ) = staking.getDelegationInfo(delegator, validator);
        require(newDelegatedAmount == initialDelegatedAmount + additionalDelegation, "Delegation should increase correctly");
        
        console.log(unicode"✓ Delegation increased successfully. New total:", newDelegatedAmount / 1 ether, "ETH");
    }
    
    function testUnstaking() internal {
        console.log("\n=== Testing Unstaking ===");
        
        (uint256 delegatedAmount, , , ) = staking.getDelegationInfo(delegator, validator);
        
        // Request unstaking (undelegate)
        vm.startBroadcast(delegator);
        staking.undelegate(validator, delegatedAmount);
        vm.stopBroadcast();
        
        // Get unbonding period from proposal contract
        uint256 unbondingPeriod = proposal.unbondingPeriod();
        
        // Fast forward to end of unbonding period
        vm.warp(block.timestamp + unbondingPeriod);
        vm.roll(block.number + unbondingPeriod);
        
        // Withdraw unstaked amount
        uint256 initialBalance = delegator.balance;
        vm.startBroadcast(delegator);
        staking.withdrawUnbonded(validator, 1);
        vm.stopBroadcast();
        
        // Verify withdrawal
        uint256 expectedBalance = initialBalance + delegatedAmount;
        require(delegator.balance == expectedBalance, "Delegator should receive unstaked amount");
        
        // Verify delegation is zero
        (uint256 finalDelegatedAmount, , , ) = staking.getDelegationInfo(delegator, validator);
        require(finalDelegatedAmount == 0, "Delegation should be zero after unstaking");
        
        console.log(unicode"✓ Unstaking successful. Delegator received:", delegatedAmount / 1 ether, "ETH");
    }
}
