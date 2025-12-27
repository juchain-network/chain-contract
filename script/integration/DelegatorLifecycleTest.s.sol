// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseTestScript} from "./BaseTestScript.s.sol";
import {console} from "forge-std/Test.sol";

contract DelegatorLifecycleTest is BaseTestScript {
    // Configuration
    uint256 public constant BLOCK_REWARD = 0.2 ether;
    

    

    
    // Test accounts
    address[] public delegatorAccounts;
    address public newValidator;
    
    function setUp() public {
        // Create test accounts using base class method
        createTestAccounts();
        
        // Deploy and initialize contracts using base class method
        deployAndInitializeContracts();
    }
    
    function createTestAccounts() internal override {
        // Call base class method to create initial validators
        super.createTestAccounts();
        
        // Create delegator accounts
        for (uint256 i = 0; i < 5; i++) {
            delegatorAccounts.push(vm.addr(uint256(keccak256(abi.encodePacked("delegator", i)))));
            vm.deal(delegatorAccounts[i], 1000000 ether);
        }
        
        // Create new validator account with 110k ETH (100k stake + 10k fees)
        newValidator = fundNewValidator(uint256(keccak256(abi.encodePacked("newValidator1"))));
    }
    

    
    function run() public override {
        console.log("Starting Delegator Lifecycle Tests...");
        
        // Create test accounts and deploy contracts (setUp() not called automatically when running as script)
        createTestAccounts();
        deployAndInitializeContracts();
        
        // Test 1: Delegate to different status validators
        testDelegateToDifferentStatusValidators();
        
        // Test 2: Complex delegator operations
        testComplexDelegatorOperations();
        
        // Test 3: Delegate reward claiming scenarios
        testDelegateRewardClaiming();
        
        console.log("\nAll Delegator Lifecycle tests completed successfully!");
    }
    
    function testDelegateToDifferentStatusValidators() internal {
        console.log("\n=== Testing Delegate to Different Status Validators ===");
        
        // Set a random validator as miner temporarily
        address miner = validatorAccounts[0];
        setMinerTemporarily(miner);
        vm.deal(miner, 1000 ether);
        address delegator = delegatorAccounts[0];
        uint256 delegateAmount = 100 ether;
        
        // Test 1: Delegate to active validator
        address activeValidator = validatorAccounts[0];
        vm.prank(delegator);
        staking.delegate{value: delegateAmount}(activeValidator);
        
        (uint256 stake, , , ) = staking.getDelegationInfo(delegator, activeValidator);
        require(stake == delegateAmount, "Delegation to active validator should succeed");
        
        // Test 2: Delegate to new proposal-passed validator (not yet registered)
        // Create and pass proposal for new validator
        vm.prank(validatorAccounts[0]);
        bytes32 proposalId = proposal.createProposal(newValidator, true, "Add new validator for delegation test");
        require(proposalId != bytes32(0), "Proposal creation failed");
        require(proposalId != bytes32(0), "Proposal ID should be non-zero");
        
        // Vote for the proposal from all validators
        for (uint256 i = 0; i < validatorAccounts.length; i++) {
            vm.prank(validatorAccounts[i]);
            proposal.voteProposal(proposalId, true);
        }
        
        vm.prank(delegator);
        try staking.delegate{value: delegateAmount}(newValidator) {
            // Should fail because validator is not registered yet
            revert("Delegation to unregistered validator should fail");
        } catch {
            console.log("Expected failure: Delegation to unregistered validator");
        }
        
        // Register the new validator
        vm.prank(newValidator);
        staking.registerValidator{value: INITIAL_STAKE}(1500); // 15% commission rate
        
        // Now delegation should succeed
        vm.prank(delegator);
        staking.delegate{value: delegateAmount}(newValidator);
        
        (stake, , , ) = staking.getDelegationInfo(delegator, newValidator);
        require(stake == delegateAmount, "Delegation to registered validator should succeed");
        
        console.log(unicode"✓ Delegate to Different Status Validators test passed");
    }
    
    function testComplexDelegatorOperations() internal {
        console.log("\n=== Testing Complex Delegator Operations ===");
        
        // Set a different validator as miner temporarily
        address miner = validatorAccounts[1];
        setMinerTemporarily(miner);
        vm.deal(miner, 1000 ether);
        address delegator = delegatorAccounts[1];
        address validator = validatorAccounts[1];
        
        // Test 1: Multiple partial delegations to same validator
        uint256 firstDelegate = 50 ether;
        uint256 secondDelegate = 75 ether;
        uint256 thirdDelegate = 100 ether;
        
        vm.prank(delegator);
        staking.delegate{value: firstDelegate}(validator);
        
        vm.prank(delegator);
        staking.delegate{value: secondDelegate}(validator);
        
        vm.prank(delegator);
        staking.delegate{value: thirdDelegate}(validator);
        
        (uint256 totalStake, , , ) = staking.getDelegationInfo(delegator, validator);
        uint256 expectedStake = firstDelegate + secondDelegate + thirdDelegate;
        require(totalStake == expectedStake, "Multiple delegations should sum correctly");
        
        // Test 2: Delegate to multiple validators
        address validator2 = validatorAccounts[2];
        uint256 delegateToSecond = 60 ether;
        
        vm.prank(delegator);
        staking.delegate{value: delegateToSecond}(validator2);
        
        (uint256 stake2, , , ) = staking.getDelegationInfo(delegator, validator2);
        require(stake2 == delegateToSecond, "Delegation to second validator should succeed");
        
        console.log(unicode"✓ Complex Delegator Operations test passed");
    }
    
    function testDelegateRewardClaiming() internal {
        console.log("\n=== Testing Delegate Reward Claiming ===");
        
        address delegator = delegatorAccounts[2];
        address validator = validatorAccounts[0];
        
        // Test 1: Delegate and claim reward after block production
        uint256 delegateAmount = 200 ether;
        
        vm.prank(delegator);
        staking.delegate{value: delegateAmount}(validator);
        
        // Test 3: Simulate block reward distribution - set miner temporarily
        address miner = validatorAccounts[0];
        setMinerTemporarily(miner);
        vm.deal(miner, 1000 ether);
        
        vm.prank(miner);
        validators.distributeBlockReward{value: 0.1 ether}();
        
        vm.prank(miner);
        staking.distributeRewards{value: BLOCK_REWARD}();
        
        // Claim rewards
        vm.prank(delegator);
        staking.claimRewards(validator);
        console.log("Delegator claimed reward successfully");
        
        // Test 2: Claim rewards after multiple blocks - miner should still be set
        for (uint i = 0; i < 5; i++) {
            vm.prank(miner);
            validators.distributeBlockReward{value: 0.1 ether}();
            vm.prank(miner);
            staking.distributeRewards{value: BLOCK_REWARD}();
        }
        
        vm.prank(delegator);
        staking.claimRewards(validator);
        console.log("Delegator claimed accumulated reward successfully");
        
        console.log(unicode"✓ Delegate Reward Claiming test passed");
    }
}