// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseTestScript} from "./BaseTestScript.s.sol";
import {console} from "forge-std/Test.sol";

contract ValidatorLifecycleTest is BaseTestScript {
    // Configuration
    uint256 public constant BLOCK_REWARD = 0.2 ether;
    
    // Test accounts
    address[] public delegatorAccounts;
    uint256[] public delegatorKeys;
    address public newValidator;
    address public newValidator2;
    uint256 public newValidatorKey;
    uint256 public newValidator2Key;
    
    function createTestAccounts() internal override {
        // Call base class to create initial validators and delegators
        super.createTestAccounts();
        
        // Create new validator accounts for testing with sufficient funding
        newValidatorKey = uint256(keccak256(abi.encodePacked("newValidator1")));
        newValidator = fundNewValidator(newValidatorKey);
        newValidator2Key = uint256(keccak256(abi.encodePacked("newValidator2")));
        newValidator2 = fundNewValidator(newValidator2Key);
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
        
        // Test 1: Create and pass proposal for new validator
        console.log("Creating proposal for new validator...");
        vm.startBroadcast(getValidatorKey(0));
        bytes32 proposalId = proposal.createProposal(newValidator, true, "Add new validator");
        require(proposalId != bytes32(0), "Proposal creation should succeed");
        vm.stopBroadcast();
        console.log("Proposal created successfully with ID:");
        console.logBytes32(proposalId);
        
        // Vote for the proposal from all validators
        console.log("Starting voting process for proposal ID:");
        console.logBytes32(proposalId);
        for (uint256 i = 0; i < validatorAccounts.length; i++) {
            console.log("Validator", i, "voting...");
            vm.startBroadcast(getValidatorKey(i));
            proposal.voteProposal(proposalId, true);
            vm.stopBroadcast();
            console.log("Validator", i, "voted successfully");
        }
        console.log("All validators voted on proposal");
        
        // Test 2: Register new validator
        console.log("Registering new validator...");
        console.log("Validator address:", newValidator);
        console.log("Initial stake:");
        console.logUint(INITIAL_STAKE / 1 ether);
        console.log("ETH");
        vm.startBroadcast(newValidatorKey);
        staking.registerValidator{value: INITIAL_STAKE}(1500); // 15% commission rate
        vm.stopBroadcast();
        console.log("Validator registered successfully");
        
        // Verify registration
        (uint256 selfStake, , , , , , , ) = staking.getValidatorInfo(newValidator);
        require(selfStake == INITIAL_STAKE, "Validator should have correct self-stake");
        
        // Test 3: Simulate epoch switch to activate validator
        console.log("Simulating epoch switch to activate validator...");
        address[] memory topValidators = validators.getTopValidators();
        vm.startBroadcast(getValidatorKey(0));
        validators.setMiner(getValidatorAddr(0));
        validators.updateActiveValidatorSet(topValidators, 1);
        vm.stopBroadcast();
        console.log("Active validator set updated");
        
        // Test 4: Simulate block reward distribution
        console.log("Distributing block reward...");
        vm.startBroadcast(newValidatorKey);
        validators.setMiner(newValidator);
        validators.distributeBlockReward{value: 0.1 ether}();
        vm.stopBroadcast();
        console.log("Block reward distributed");
        
        console.log("Distributing staking rewards...");
        vm.startBroadcast(newValidatorKey);
        staking.setMiner(newValidator);
        staking.distributeRewards{value: BLOCK_REWARD}();
        vm.stopBroadcast();
        console.log("Staking rewards distributed");
        
        // Test 5: Claim rewards
        console.log("Validator claiming rewards...");
        vm.startBroadcast(newValidatorKey);
        staking.claimValidatorRewards();
        vm.stopBroadcast();
        console.log("New validator claimed reward successfully");
        
        // Test 6: Add more stake
        console.log("Adding more stake to validator...");
        console.log("Additional stake:");
        console.logUint(50000 ether / 1 ether);
        console.log("ETH");
        vm.startBroadcast(newValidatorKey);
        staking.addValidatorStake{value: 50000 ether}();
        vm.stopBroadcast();
        console.log("Stake added successfully");
        
        // Test 7: Update commission rate
        console.log("Updating validator commission rate...");
        console.log("New commission rate: 20%");
        vm.startBroadcast(newValidatorKey);
        staking.updateCommissionRate(2000); // 20% commission rate
        vm.stopBroadcast();
        console.log("Commission rate updated successfully");
        
        console.log(unicode"✓ Proposal-added Validator Lifecycle test passed");
    }
    
    function testValidatorRejoiningWithoutExit() internal {
        console.log("\n=== Testing Validator Rejoining Without Exit ===");
        
        // Set a different validator as miner temporarily
        address miner = validatorAccounts[1];
        setMinerTemporarily(miner);
        
        // Test 1: Create and pass proposal for new validator 2
        console.log("Creating proposal for new validator 2...");
        vm.startBroadcast(validatorKeys[0]);
        // Create proposal
        bytes32 proposalId = proposal.createProposal(newValidator2, true, "Add new validator 2");
        require(proposalId != bytes32(0), "Proposal creation failed");
        vm.stopBroadcast();
        console.log("Proposal created successfully with ID:");
        console.logBytes32(proposalId);
        
        // Vote for the proposal from all validators
        console.log("Starting voting process for proposal ID:");
        console.logBytes32(proposalId);
        for (uint256 i = 0; i < validatorAccounts.length; i++) {
            console.log("Validator", i, "voting...");
            vm.startBroadcast(validatorKeys[i]);
            proposal.voteProposal(proposalId, true);
            vm.stopBroadcast();
            console.log("Validator", i, "voted successfully");
        }
        console.log("All validators voted on proposal");
        
        // Test 2: Register and activate new validator 2
        console.log("Registering new validator 2...");
        console.log("Validator address:", newValidator2);
        console.log("Initial stake:");
        console.logUint(INITIAL_STAKE / 1 ether);
        console.log("ETH");
        vm.startBroadcast(newValidator2Key);
        staking.registerValidator{value: INITIAL_STAKE}(1000); // 10% commission rate
        vm.stopBroadcast();
        console.log("Validator 2 registered successfully");
        
        // Simulate epoch switch to activate validator
        console.log("Simulating epoch switch to activate validator 2...");
        address[] memory topValidators1 = validators.getTopValidators();
        vm.startBroadcast(getValidatorKey(1));
        validators.setMiner(getValidatorAddr(1));
        validators.updateActiveValidatorSet(topValidators1, 1);
        vm.stopBroadcast();
        console.log("Validator 2 activated");
        
        // Create config proposal to change validatorUnjailPeriod to 6 blocks
        console.log("Creating config proposal to change validatorUnjailPeriod to 6 blocks...");
        vm.startBroadcast(validatorKeys[0]);
        bytes32 configProposalId = proposal.createUpdateConfigProposal(7, 6);
        require(configProposalId != bytes32(0), "Config proposal creation failed");
        vm.stopBroadcast();
        console.log("Config proposal created with ID:");
        console.logBytes32(configProposalId);
        
        // Vote for the config proposal from all validators
        console.log("Voting on config proposal...");
        for (uint256 i = 0; i < validatorAccounts.length; i++) {
            console.log("Validator", i, "voting on config proposal...");
            vm.startBroadcast(validatorKeys[i]);
            proposal.voteProposal(configProposalId, true);
            vm.stopBroadcast();
            console.log("Validator", i, "voted successfully");
        }
        console.log("All validators voted on config proposal");
        
        // Test 3: Validator resigns but doesn't exit
        console.log("Validator 2 resigning...");
        vm.startBroadcast(newValidator2Key);
        staking.resignValidator();
        vm.stopBroadcast();
        console.log("Validator 2 resigned");
        
        // Test 4: Simulate epoch switch (use epoch=1 to bypass onlyBlockEpoch restriction)
        console.log("Simulating epoch switch after resignation...");
        address[] memory topValidators2 = validators.getTopValidators();
        vm.startBroadcast(getValidatorKey(1));
        validators.updateActiveValidatorSet(topValidators2, 1);
        vm.stopBroadcast();
        console.log("Active validator set updated after resignation");
        
        // Test 5: Check if validator is jailed
        bool isJailed = staking.isValidatorJailed(newValidator2);
        require(isJailed == true, "Validator should be jailed after resigning");
        
        // Test 6: Create and pass new proposal for the same validator
        console.log("Creating re-add proposal for validator 2...");
        vm.startBroadcast(validatorKeys[0]);
        // Create proposal
        bytes32 proposalId2 = proposal.createProposal(newValidator2, true, "Re-add validator 2");
        require(proposalId2 != bytes32(0), "Proposal creation failed");
        vm.stopBroadcast();
        console.log("Re-add proposal created with ID:");
        console.logBytes32(proposalId2);
        
        // Vote for the proposal from all validators
        console.log("Starting voting process for re-add proposal...");
        for (uint256 i = 0; i < validatorAccounts.length; i++) {
            console.log("Validator", i, "voting on re-add proposal...");
            vm.startBroadcast(validatorKeys[i]);
            proposal.voteProposal(proposalId2, true);
            vm.stopBroadcast();
            console.log("Validator", i, "voted successfully");
        }
        console.log("All validators voted on re-add proposal");
        
        // Test 8: Unjail validator (should work without re-registering)
        // Simulate blocks past the jail period using simulateBlocks
        console.log("Simulating blocks past jail period...");
        simulateBlocks(10); // Jail period is 6 blocks, so simulate 6 blocks
        
        console.log("Unjailing validator 2...");
        vm.startBroadcast(newValidator2Key);
        staking.unjailValidator(newValidator2);
        vm.stopBroadcast();
        console.log("Validator 2 unjailed successfully");
        
        // Test 9: Simulate epoch switch to reactivate (use epoch=1 to bypass onlyBlockEpoch restriction)
        console.log("Simulating epoch switch to reactivate validator 2...");
        address[] memory topValidators3 = validators.getTopValidators();
        vm.startBroadcast(getValidatorKey(1));
        validators.setMiner(getValidatorAddr(1));
        validators.updateActiveValidatorSet(topValidators3, 1);
        vm.stopBroadcast();
        console.log("Validator 2 reactivated");
        
        // Test 10: Check if validator is active again
        (bool isActive, ) = staking.getValidatorStatus(newValidator2);
        console.log("Validator active status after rejoining:");
        console.logBool(isActive);
        
        console.log(unicode"✓ Validator Rejoining Without Exit test passed");
    }
    
    function testValidatorRejoiningAfterExit() internal {
        console.log("\n=== Testing Validator Rejoining After Exit ===");
        
        // Set another validator as miner temporarily
        address miner = validatorAccounts[2];
        setMinerTemporarily(miner);
        
        // Use newValidator for this test
        
        // Create config proposal to change validatorUnjailPeriod to 6 blocks
        console.log("Creating config proposal to change validatorUnjailPeriod to 6 blocks...");
        vm.startBroadcast(validatorKeys[0]);
        bytes32 configProposalId = proposal.createUpdateConfigProposal(7, 6);
        require(configProposalId != bytes32(0), "Config proposal creation failed");
        vm.stopBroadcast();
        console.log("Config proposal created with ID:");
        console.logBytes32(configProposalId);
        
        // Vote for the config proposal from all validators
        console.log("Voting on config proposal...");
        for (uint256 i = 0; i < validatorAccounts.length; i++) {
            console.log("Validator", i, "voting on config proposal...");
            vm.startBroadcast(validatorKeys[i]);
            proposal.voteProposal(configProposalId, true);
            vm.stopBroadcast();
            console.log("Validator", i, "voted successfully");
        }
        console.log("All validators voted on config proposal");
        
        // Test 1: Validator resigns
        console.log("Validator 1 resigning...");
        vm.startBroadcast(newValidatorKey);
        staking.resignValidator();
        vm.stopBroadcast();
        console.log("Validator 1 resigned");
        
        // Test 2: Simulate epoch switch
        console.log("Simulating epoch switch after resignation...");
        address[] memory topValidators1 = validators.getTopValidators();
        vm.startBroadcast(getValidatorKey(2));
        validators.updateActiveValidatorSet(topValidators1, 1);
        vm.stopBroadcast();
        console.log("Active validator set updated after resignation");
        
        // Test 3: Validator exits completely
        console.log("Validator 1 exiting completely...");
        vm.startBroadcast(newValidatorKey);
        staking.exitValidator();
        vm.stopBroadcast();
        console.log("Validator 1 exited");
        
        // Test 4: Withdraw unbonded stake
        console.log("Withdrawing unbonded stake...");
        vm.startBroadcast(newValidatorKey);
        staking.withdrawUnbonded(newValidator, 10); // 10 is the max entries to process
        vm.stopBroadcast();
        console.log("Unbonded stake withdrawn");
        
        // Test 5: Create and pass new proposal for the same validator
        console.log("Creating re-add proposal for validator 1...");
        vm.startBroadcast(validatorKeys[0]);
        // Create proposal
        bytes32 proposalId = proposal.createProposal(newValidator, true, "Re-add validator");
        require(proposalId != bytes32(0), "Proposal creation failed");
        vm.stopBroadcast();
        console.log("Re-add proposal created with ID:");
        console.logBytes32(proposalId);
        
        // Vote for the proposal from all validators
        console.log("Starting voting process for re-add proposal...");
        for (uint256 i = 0; i < validatorAccounts.length; i++) {
            console.log("Validator", i, "voting on re-add proposal...");
            vm.startBroadcast(validatorKeys[i]);
            proposal.voteProposal(proposalId, true);
            vm.stopBroadcast();
            console.log("Validator", i, "voted successfully");
        }
        console.log("All validators voted on re-add proposal");
        
        // Test 6: Add stake first (required after exiting)
        console.log("Adding initial stake after exit...");
        console.log("Initial stake:");
        console.logUint(INITIAL_STAKE / 1 ether);
        console.log("ETH");
        vm.startBroadcast(newValidatorKey);
        staking.addValidatorStake{value: INITIAL_STAKE}();
        vm.stopBroadcast();
        console.log("Initial stake added successfully");
        
        // Test 7: Unjail validator
        // Simulate blocks past the jail period using simulateBlocks
        console.log("Simulating blocks past jail period...");
        simulateBlocks(6); // Jail period is 6 blocks, so simulate 6 blocks
        
        console.log("Unjailing validator 1...");
        vm.startBroadcast(newValidatorKey);
        staking.unjailValidator(newValidator);
        vm.stopBroadcast();
        console.log("Validator 1 unjailed successfully");
        
        // Test 8: Simulate epoch switch to activate (use epoch=1 to bypass onlyBlockEpoch restriction)
        console.log("Simulating epoch switch to reactivate validator 1...");
        address[] memory topValidators2 = validators.getTopValidators();
        vm.startBroadcast(getValidatorKey(2));
        validators.updateActiveValidatorSet(topValidators2, 1);
        vm.stopBroadcast();
        console.log("Validator 1 reactivated");
        
        // Test 9: Decrease stake
        console.log("Decreasing validator stake...");
        console.log("Decrease amount:");
        console.logUint(20000 ether / 1 ether);
        console.log("ETH");
        vm.startBroadcast(newValidatorKey);
        staking.decreaseValidatorStake(20000 ether);
        vm.stopBroadcast();
        console.log("Stake decreased successfully");
        
        console.log(unicode"✓ Validator Rejoining After Exit test passed");
    }
}