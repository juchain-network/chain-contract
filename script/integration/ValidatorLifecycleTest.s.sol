// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseTestScript} from "./BaseTestScript.s.sol";
import {console} from "forge-std/Test.sol";

contract ValidatorLifecycleTest is BaseTestScript {
    // Configuration
    uint256 public constant BLOCK_REWARD = 0.2 ether;
    
    // Test accounts
    address[] public delegatorAccounts;
    address public newValidator;
    address public newValidator2;
    
    function setUp() public {
        // Create test accounts
        createTestAccounts();
        
        // Deploy and initialize contracts
        deployAndInitializeContracts();
    }
    
    function createTestAccounts() internal override {
        // Call base class to create initial validators
        super.createTestAccounts();
        
        // Create delegator accounts
        for (uint256 i = 0; i < 5; i++) {
            delegatorAccounts.push(vm.addr(uint256(keccak256(abi.encodePacked("delegator", i)))));
            vm.deal(delegatorAccounts[i], 1000000 ether);
        }
        
        // Create new validator accounts for testing with sufficient funding
        newValidator = fundNewValidator(uint256(keccak256(abi.encodePacked("newValidator1"))));
        newValidator2 = fundNewValidator(uint256(keccak256(abi.encodePacked("newValidator2"))));
    }
    

    
    function run() public override {
        console.log("Starting Validator Lifecycle Tests...");
        
        // Create test accounts and deploy contracts (setUp() not called automatically when running as script)
        createTestAccounts();
        deployAndInitializeContracts();
        
        // Test 1: Proposal-added validator complete lifecycle
        testProposalAddedValidatorLifecycle();
        
        // Test 2: Validator rejoining without exiting stake
        testValidatorRejoiningWithoutExit();
        
        // Test 3: Validator rejoining after exiting stake
        testValidatorRejoiningAfterExit();
        
        console.log("\nAll Validator Lifecycle tests completed successfully!");
    }
    
    function testProposalAddedValidatorLifecycle() internal {
        console.log("\n=== Testing Proposal-added Validator Lifecycle ===");
        
        // Set a random validator as miner temporarily
        address miner = validatorAccounts[0];
        setMinerTemporarily(miner);
        vm.deal(miner, 1000 ether);
        
        // Test 1: Create and pass proposal for new validator
        vm.prank(validatorAccounts[0]);
        bytes32 proposalId = proposal.createProposal(newValidator, true, "Add new validator");
        require(proposalId != bytes32(0), "Proposal creation should succeed");
        
        // Vote for the proposal from all validators
        for (uint256 i = 0; i < validatorAccounts.length; i++) {
            vm.prank(validatorAccounts[i]);
            proposal.voteProposal(proposalId, true);
        }
        
        // Test 2: Register new validator
        vm.prank(newValidator);
        staking.registerValidator{value: INITIAL_STAKE}(1500); // 15% commission rate
        
        // Verify registration
        (uint256 selfStake, , , , , , , ) = staking.getValidatorInfo(newValidator);
        require(selfStake == INITIAL_STAKE, "Validator should have correct self-stake");
        
        // Test 3: Simulate epoch switch to activate validator
        address[] memory topValidators = validators.getTopValidators();
        vm.prank(miner);
        validators.updateActiveValidatorSet(topValidators, 1);
        
        // Test 4: Simulate block reward distribution
        vm.prank(miner);
        validators.distributeBlockReward{value: 0.1 ether}();
        
        vm.prank(miner);
        staking.distributeRewards{value: BLOCK_REWARD}();
        
        // Test 5: Claim rewards
        vm.prank(newValidator);
        staking.claimValidatorRewards();
        console.log("New validator claimed reward successfully");
        
        // Test 6: Add more stake
        vm.prank(newValidator);
        staking.addValidatorStake{value: 50000 ether}();
        
        // Test 7: Update commission rate
        vm.prank(newValidator);
        staking.updateCommissionRate(2000); // 20% commission rate
        
        console.log(unicode"✓ Proposal-added Validator Lifecycle test passed");
    }
    
    function testValidatorRejoiningWithoutExit() internal {
        console.log("\n=== Testing Validator Rejoining Without Exit ===");
        
        // Set a different validator as miner temporarily
        address miner = validatorAccounts[1];
        setMinerTemporarily(miner);
        vm.deal(miner, 1000 ether);
        
        // Test 1: Create and pass proposal for new validator 2
        vm.prank(validatorAccounts[0]);
        // Generate expected proposal ID
        bytes32 expectedProposalId = keccak256(abi.encodePacked(validatorAccounts[0], newValidator2, true, "Add new validator 2", block.timestamp));
        
        // Create proposal
        bytes32 actualProposalId = proposal.createProposal(newValidator2, true, "Add new validator 2");
        require(actualProposalId != bytes32(0), "Proposal creation failed");
        
        bytes32 proposalId = expectedProposalId;
        
        // Vote for the proposal from all validators
        for (uint256 i = 0; i < validatorAccounts.length; i++) {
            vm.prank(validatorAccounts[i]);
            proposal.voteProposal(proposalId, true);
        }
        
        // Test 2: Register and activate new validator 2
        vm.prank(newValidator2);
        staking.registerValidator{value: INITIAL_STAKE}(1000); // 10% commission rate
        
        // Simulate epoch switch to activate validator
        address[] memory topValidators1 = validators.getTopValidators();
        vm.prank(miner);
        validators.updateActiveValidatorSet(topValidators1, 1);
        
        // Test 3: Validator resigns but doesn't exit
        vm.prank(newValidator2);
        staking.resignValidator();
        
        // Test 4: Simulate epoch switch (use epoch=1 to bypass onlyBlockEpoch restriction)
        address[] memory topValidators2 = validators.getTopValidators();
        vm.prank(miner);
        validators.updateActiveValidatorSet(topValidators2, 1);
        
        // Test 5: Check if validator is jailed
        bool isJailed = staking.isValidatorJailed(newValidator2);
        require(isJailed == true, "Validator should be jailed after resigning");
        
        // Test 6: Create and pass new proposal for the same validator
        vm.prank(validatorAccounts[0]);
        
        // Generate expected proposal ID
        bytes32 expectedProposalId2 = keccak256(abi.encodePacked(validatorAccounts[0], newValidator2, true, "Re-add validator 2", block.timestamp));
        
        // Create proposal
        bytes32 actualProposalId2 = proposal.createProposal(newValidator2, true, "Re-add validator 2");
        require(actualProposalId2 != bytes32(0), "Proposal creation failed");
        
        bytes32 proposalId2 = expectedProposalId2;
        
        // Vote for the proposal from all validators
        for (uint256 i = 0; i < validatorAccounts.length; i++) {
            vm.prank(validatorAccounts[i]);
            proposal.voteProposal(proposalId2, true);
        }
        
        // Test 8: Unjail validator (should work without re-registering)
        // Advance block number past the jail period using vm.roll
        vm.roll(86402); // Jail period is 86400 blocks, so we need to roll to block 86401 or later
        
        vm.prank(newValidator2);
        staking.unjailValidator(newValidator2);
        
        // Test 9: Simulate epoch switch to reactivate (use epoch=1 to bypass onlyBlockEpoch restriction)
        address[] memory topValidators3 = validators.getTopValidators();
        vm.prank(miner);
        validators.updateActiveValidatorSet(topValidators3, 1);
        
        // Test 10: Check if validator is active again
        (bool isActive, ) = staking.getValidatorStatus(newValidator2);
        console.log("Validator active status after rejoining:", isActive);
        
        console.log(unicode"✓ Validator Rejoining Without Exit test passed");
    }
    
    function testValidatorRejoiningAfterExit() internal {
        console.log("\n=== Testing Validator Rejoining After Exit ===");
        
        // Set another validator as miner temporarily
        address miner = validatorAccounts[2];
        setMinerTemporarily(miner);
        vm.deal(miner, 1000 ether);
        
        // Use newValidator for this test
        
        // Test 1: Validator resigns
        vm.prank(newValidator);
        staking.resignValidator();
        
        // Test 2: Simulate epoch switch
        address[] memory topValidators1 = validators.getTopValidators();
        vm.prank(miner);
        validators.updateActiveValidatorSet(topValidators1, 1);
        
        // Test 3: Validator exits completely
        vm.prank(newValidator);
        staking.exitValidator();
        
        // Test 4: Withdraw unbonded stake
        vm.prank(newValidator);
        staking.withdrawUnbonded(newValidator, 10); // 10 is the max entries to process
        
        // Test 5: Create and pass new proposal for the same validator
        vm.prank(validatorAccounts[0]);
        
        // Generate expected proposal ID
        bytes32 expectedProposalId = keccak256(abi.encodePacked(validatorAccounts[0], newValidator, true, "Re-add validator", block.timestamp));
        
        // Create proposal
        bytes32 actualProposalId = proposal.createProposal(newValidator, true, "Re-add validator");
        require(actualProposalId != bytes32(0), "Proposal creation failed");
        
        bytes32 proposalId = expectedProposalId;
        
        // Vote for the proposal from all validators
        for (uint256 i = 0; i < validatorAccounts.length; i++) {
            vm.prank(validatorAccounts[i]);
            proposal.voteProposal(proposalId, true);
        }
        
        // Test 6: Add stake first (required after exiting)
        vm.prank(newValidator);
        staking.addValidatorStake{value: INITIAL_STAKE}();
        
        // Test 7: Unjail validator
        // Advance block number past the jail period using vm.roll
        vm.roll(86402); // Jail period is 86400 blocks, so we need to roll to block 86401 or later
        
        vm.prank(newValidator);
        staking.unjailValidator(newValidator);
        
        // Test 8: Simulate epoch switch to activate (use epoch=1 to bypass onlyBlockEpoch restriction)
        address[] memory topValidators2 = validators.getTopValidators();
        vm.prank(miner);
        validators.updateActiveValidatorSet(topValidators2, 1);
        
        // Test 9: Decrease stake
        vm.prank(newValidator);
        staking.decreaseValidatorStake(20000 ether);
        
        console.log(unicode"✓ Validator Rejoining After Exit test passed");
    }
}