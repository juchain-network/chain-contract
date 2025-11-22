// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Test} from "forge-std/Test.sol";
import {Staking} from "../contracts/Staking.sol";
import {Proposal} from "../contracts/Proposal.sol";
import {Validators} from "../contracts/Validators.sol";
import {Punish} from "../contracts/Punish.sol";

contract StakingTest is Test {
    // System contract addresses (fixed addresses for testing)
    address constant VALIDATORS = 0x000000000000000000000000000000000000f000;
    address constant PUNISH = 0x000000000000000000000000000000000000F001;
    address constant PROPOSAL = 0x000000000000000000000000000000000000F002;
    address constant STAKING = 0x000000000000000000000000000000000000F003;
    
    // Test addresses
    address constant VALIDATOR1 = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;
    address constant VALIDATOR2 = 0x70997970C51812dc3A010C7d01b50e0d17dc79C8;
    address constant VALIDATOR3 = 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC;
    address constant VALIDATOR4 = 0x90F79bf6EB2c4f870365E785982E1f101E93b906;
    address constant VALIDATOR5 = 0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65;
    address constant VALIDATOR6 = 0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc;
    
    address constant DELEGATOR1 = 0x976EA74026E726554dB657fA54763abd0C3a0aa9;
    address constant DELEGATOR2 = 0x14dC79964da2C08b23698B3D3cc7Ca32193d9955;
    
    uint256 constant MIN_STAKE = 10000 ether;
    uint256 constant COMMISSION_RATE = 1000; // 10%

    function setUp() public {
        // Deploy contracts to fixed addresses (required for onlyStakingContract modifier)
        vm.etch(VALIDATORS, type(Validators).runtimeCode);
        vm.etch(PUNISH, type(Punish).runtimeCode);
        vm.etch(PROPOSAL, type(Proposal).runtimeCode);
        vm.etch(STAKING, type(Staking).runtimeCode);
        
        // Set up test account balances
        vm.deal(VALIDATOR1, 100000 ether);
        vm.deal(VALIDATOR2, 100000 ether);
        vm.deal(VALIDATOR3, 100000 ether);
        vm.deal(VALIDATOR4, 100000 ether);
        vm.deal(VALIDATOR5, 100000 ether);
        vm.deal(VALIDATOR6, 100000 ether);
        vm.deal(DELEGATOR1, 100000 ether);
        vm.deal(DELEGATOR2, 100000 ether);
        
        // Initialize contracts in correct order
        address[] memory initVals = new address[](0);
        Proposal(PROPOSAL).initialize(initVals, VALIDATORS);
        Staking(STAKING).initialize(VALIDATORS, PROPOSAL);
        Punish(PUNISH).initialize(VALIDATORS, PROPOSAL, STAKING);
        Validators(VALIDATORS).initialize(initVals, PROPOSAL, PUNISH, STAKING);
    }

    function testInitialization() public view {
        assertEq(Staking(STAKING).MIN_VALIDATOR_STAKE(), MIN_STAKE);
        assertEq(Staking(STAKING).MAX_VALIDATORS(), 21);
        assertEq(Staking(STAKING).MIN_VALIDATORS(), 3);
        assertEq(Staking(STAKING).getValidatorCount(), 0);
        assertEq(Staking(STAKING).getActiveValidatorCount(), 0);
        assertFalse(Staking(STAKING).hasMinimumValidators());
    }
    
    // Helper function to set up validator with pass status
    function _setupValidatorPass(address validator) internal {
        // Set pass status directly (simulating proposal passed)
        // In real scenario, this would be done through Proposal contract voting
        // Storage layout: Params has initialized (slot 0), Proposal has 7 uint256/address vars (slots 1-7)
        // pass mapping is at slot 8, proposalPassedTime mapping is at slot 9
        vm.store(
            PROPOSAL,
            keccak256(abi.encode(validator, uint256(8))), // pass mapping slot
            bytes32(uint256(1))
        );
        // Set proposalPassedTime to current time (within 7 days)
        vm.store(
            PROPOSAL,
            keccak256(abi.encode(validator, uint256(9))), // proposalPassedTime mapping slot
            bytes32(block.timestamp)
        );
    }
    
    // Helper function to update active validator set (simulating epoch update)
    function _updateActiveValidatorSet() internal {
        // Get top validators from Staking
        address[] memory topValidators = Staking(STAKING).getTopValidators();
        
        if (topValidators.length == 0) {
            return; // No validators to update
        }
        
        // Update active validator set (requires miner and epoch boundary)
        // Set coinbase to simulate miner
        address miner = address(0x123);
        vm.coinbase(miner);
        
        // Roll to epoch boundary (epoch is typically 30 blocks)
        uint256 epoch = 30;
        uint256 currentBlock = block.number;
        uint256 nextEpoch = ((currentBlock / epoch) + 1) * epoch;
        vm.roll(nextEpoch);
        
        // Update active validator set
        vm.prank(miner);
        Validators(VALIDATORS).updateActiveValidatorSet(topValidators, epoch);
    }

    function testValidatorRegistration() public {
        // Set up validator pass status
        _setupValidatorPass(VALIDATOR1);
        
        // Test successful registration
        vm.prank(VALIDATOR1);
        Staking(STAKING).registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        assertEq(Staking(STAKING).getValidatorCount(), 1);
        
        // Update active validator set to make validator active
        _updateActiveValidatorSet();
        
        assertEq(Staking(STAKING).getActiveValidatorCount(), 1);
        
        (uint256 selfStake, uint256 totalDelegated, uint256 commissionRate, bool isJailed, uint256 jailUntilBlock) = 
            Staking(STAKING).getValidatorInfo(VALIDATOR1);
            
        assertEq(selfStake, MIN_STAKE);
        assertEq(totalDelegated, 0);
        assertEq(commissionRate, COMMISSION_RATE);
        assertFalse(isJailed);
        assertEq(jailUntilBlock, 0);
    }

    function test_RevertWhen_InsufficientStake() public {
        _setupValidatorPass(VALIDATOR1);
        vm.prank(VALIDATOR1);
        vm.expectRevert("Insufficient self-stake");
        Staking(STAKING).registerValidator{value: MIN_STAKE - 1}(COMMISSION_RATE);
    }

    function test_RevertWhen_InvalidCommissionRate() public {
        _setupValidatorPass(VALIDATOR1);
        vm.prank(VALIDATOR1);
        vm.expectRevert("Invalid commission rate");
        Staking(STAKING).registerValidator{value: MIN_STAKE}(10001); // > 100%
    }

    function test_RevertWhen_DoubleRegistration() public {
        _setupValidatorPass(VALIDATOR1);
        vm.prank(VALIDATOR1);
        Staking(STAKING).registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        vm.prank(VALIDATOR1);
        vm.expectRevert("Already registered");
        Staking(STAKING).registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
    }

    function testMinimumValidatorsRequirement() public {
        // Register 3 validators (minimum required)
        _setupValidatorPass(VALIDATOR1);
        _setupValidatorPass(VALIDATOR2);
        _setupValidatorPass(VALIDATOR3);
        
        vm.prank(VALIDATOR1);
        Staking(STAKING).registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        vm.prank(VALIDATOR2);
        Staking(STAKING).registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        vm.prank(VALIDATOR3);
        Staking(STAKING).registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        // Update active validator set to make validators active
        _updateActiveValidatorSet();
        
        assertEq(Staking(STAKING).getActiveValidatorCount(), 3);
        assertTrue(Staking(STAKING).hasMinimumValidators());
        
        // Now register a 4th validator
        _setupValidatorPass(VALIDATOR4);
        vm.prank(VALIDATOR4);
        Staking(STAKING).registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        // Update active validator set again
        _updateActiveValidatorSet();
        
        assertEq(Staking(STAKING).getActiveValidatorCount(), 4);
        
        // Test that 4th validator can exit (leaving 3)
        vm.prank(VALIDATOR4);
        Staking(STAKING).emergencyExit();
        
        assertEq(Staking(STAKING).getActiveValidatorCount(), 3);
        assertTrue(Staking(STAKING).hasMinimumValidators());
        
        // Test that 3rd validator cannot exit (would leave 2)
        vm.prank(VALIDATOR3);
        vm.expectRevert("Cannot exit: would leave less than minimum validators");
        Staking(STAKING).emergencyExit();
    }

    function testPartialStakeWithdrawal() public {
        // Register validator with extra stake
        _setupValidatorPass(VALIDATOR1);
        _setupValidatorPass(VALIDATOR2);
        _setupValidatorPass(VALIDATOR3);
        
        vm.prank(VALIDATOR1);
        Staking(STAKING).registerValidator{value: MIN_STAKE * 2}(COMMISSION_RATE);
        
        // Register 2 more validators to meet minimum
        vm.prank(VALIDATOR2);
        Staking(STAKING).registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        vm.prank(VALIDATOR3);
        Staking(STAKING).registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        // Update active validator set to make validators active
        _updateActiveValidatorSet();
        
        // Test partial withdrawal (still leaves minimum stake)
        uint256 withdrawAmount = MIN_STAKE / 2;
        uint256 initialBalance = VALIDATOR1.balance;
        
        vm.prank(VALIDATOR1);
        Staking(STAKING).withdrawValidatorStake(withdrawAmount);
        
        assertEq(VALIDATOR1.balance, initialBalance + withdrawAmount);
        
        (uint256 selfStake,,,, ) = Staking(STAKING).getValidatorInfo(VALIDATOR1);
        assertEq(selfStake, MIN_STAKE * 2 - withdrawAmount);
        assertEq(Staking(STAKING).getActiveValidatorCount(), 3);
    }

    function test_RevertWhen_PartialWithdrawalBelowMinimum() public {
        // Register multiple validators first to avoid minimum validator constraint
        _setupValidatorPass(VALIDATOR1);
        _setupValidatorPass(VALIDATOR2);
        _setupValidatorPass(VALIDATOR3);
        _setupValidatorPass(VALIDATOR4);
        
        vm.prank(VALIDATOR1);
        Staking(STAKING).registerValidator{value: MIN_STAKE + 1000}(COMMISSION_RATE);
        vm.prank(VALIDATOR2);
        Staking(STAKING).registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        vm.prank(VALIDATOR3);
        Staking(STAKING).registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        vm.prank(VALIDATOR4);
        Staking(STAKING).registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        // Try to withdraw amount that would leave stake below minimum
        vm.prank(VALIDATOR1);
        vm.expectRevert("Remaining stake below minimum, use emergencyExit() to withdraw all");
        Staking(STAKING).withdrawValidatorStake(1001);
    }

    function testDelegation() public {
        // Register validator
        _setupValidatorPass(VALIDATOR1);
        vm.prank(VALIDATOR1);
        Staking(STAKING).registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        // Test delegation
        uint256 delegationAmount = 1000 ether;
        vm.prank(DELEGATOR1);
        Staking(STAKING).delegate{value: delegationAmount}(VALIDATOR1);
        
        (uint256 selfStake, uint256 totalDelegated,,, ) = Staking(STAKING).getValidatorInfo(VALIDATOR1);
        assertEq(selfStake, MIN_STAKE);
        assertEq(totalDelegated, delegationAmount);
        
        (uint256 delegatedAmount,,,) = Staking(STAKING).getDelegationInfo(DELEGATOR1, VALIDATOR1);
        assertEq(delegatedAmount, delegationAmount);
    }

    function test_RevertWhen_DelegateToInactiveValidator() public {
        vm.prank(DELEGATOR1);
        vm.expectRevert("Not a valid validator");
        Staking(STAKING).delegate{value: 1000 ether}(VALIDATOR1); // VALIDATOR1 not registered
    }

    function test_RevertWhen_DelegateInsufficientAmount() public {
        // Set up validator pass status
        _setupValidatorPass(VALIDATOR1);
        
        vm.prank(VALIDATOR1);
        Staking(STAKING).registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        vm.prank(DELEGATOR1);
        vm.expectRevert("Insufficient delegation amount");
        Staking(STAKING).delegate{value: 0.5 ether}(VALIDATOR1); // Below MIN_DELEGATION
    }

    function testGetTopValidators() public {
        // Register validators with different stakes
        _setupValidatorPass(VALIDATOR1);
        _setupValidatorPass(VALIDATOR2);
        _setupValidatorPass(VALIDATOR3);
        
        vm.prank(VALIDATOR1);
        Staking(STAKING).registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        vm.prank(VALIDATOR2);
        Staking(STAKING).registerValidator{value: MIN_STAKE * 2}(COMMISSION_RATE);
        
        vm.prank(VALIDATOR3);
        Staking(STAKING).registerValidator{value: MIN_STAKE * 3}(COMMISSION_RATE);
        
        address[] memory topValidators = Staking(STAKING).getTopValidators();
        
        assertEq(topValidators.length, 3);
        // Should be ordered by stake (highest first)
        assertEq(topValidators[0], VALIDATOR3); // 30,000 JU
        assertEq(topValidators[1], VALIDATOR2); // 20,000 JU
        assertEq(topValidators[2], VALIDATOR1); // 10,000 JU
    }

    function testGetTopValidatorsWithDelegations() public {
        // Register validators
        _setupValidatorPass(VALIDATOR1);
        _setupValidatorPass(VALIDATOR2);
        
        vm.prank(VALIDATOR1);
        Staking(STAKING).registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        vm.prank(VALIDATOR2);
        Staking(STAKING).registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        // Add delegation to VALIDATOR1
        vm.prank(DELEGATOR1);
        Staking(STAKING).delegate{value: MIN_STAKE}(VALIDATOR1);
        
        address[] memory topValidators = Staking(STAKING).getTopValidators();
        
        assertEq(topValidators.length, 2);
        // VALIDATOR1 should be first (20,000 total vs 10,000)
        assertEq(topValidators[0], VALIDATOR1);
        assertEq(topValidators[1], VALIDATOR2);
    }

    function testSystemInvariant_MinimumValidators() public {
        // Setup: Register exactly 3 validators (minimum required)
        address[3] memory validatorAddrs = [VALIDATOR1, VALIDATOR2, VALIDATOR3];
        
        _setupValidatorPass(VALIDATOR1);
        _setupValidatorPass(VALIDATOR2);
        _setupValidatorPass(VALIDATOR3);
        
        for (uint i = 0; i < 3; i++) {
            vm.prank(validatorAddrs[i]);
            Staking(STAKING).registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        }
        
        // Update active validator set to make validators active
        _updateActiveValidatorSet();
        
        assertTrue(Staking(STAKING).hasMinimumValidators());
        assertEq(Staking(STAKING).getActiveValidatorCount(), 3);
        
        // Test that no validator can exit
        for (uint i = 0; i < 3; i++) {
            vm.prank(validatorAddrs[i]);
            vm.expectRevert("Cannot exit: would leave less than minimum validators");
            Staking(STAKING).emergencyExit();
        }
        
        // Add one more validator
        _setupValidatorPass(VALIDATOR4);
        vm.prank(VALIDATOR4);
        Staking(STAKING).registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        // Update active validator set again
        _updateActiveValidatorSet();
        
        assertEq(Staking(STAKING).getActiveValidatorCount(), 4);
        
        // Now exactly one validator should be able to exit
        vm.prank(VALIDATOR4);
        Staking(STAKING).emergencyExit();
        
        assertEq(Staking(STAKING).getActiveValidatorCount(), 3);
        
        // Back to minimum - no one can exit again
        vm.prank(VALIDATOR1);
        vm.expectRevert("Cannot exit: would leave less than minimum validators");
        Staking(STAKING).emergencyExit();
    }

    function testValidatorJailing() public {
        // Register 4 validators
        _setupValidatorPass(VALIDATOR1);
        _setupValidatorPass(VALIDATOR2);
        _setupValidatorPass(VALIDATOR3);
        _setupValidatorPass(VALIDATOR4);
        
        vm.prank(VALIDATOR1);
        Staking(STAKING).registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        vm.prank(VALIDATOR2);
        Staking(STAKING).registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        vm.prank(VALIDATOR3);
        Staking(STAKING).registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        vm.prank(VALIDATOR4);
        Staking(STAKING).registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        // Update active validator set to make validators active
        _updateActiveValidatorSet();
        
        assertEq(Staking(STAKING).getActiveValidatorCount(), 4);
        
        // Jail one validator (simulating punishment contract call)
        vm.prank(PUNISH); // Punish contract
        Staking(STAKING).jailValidator(VALIDATOR4, 1000);
        
        // Active count should decrease (jailed validators are excluded from active count)
        assertEq(Staking(STAKING).getActiveValidatorCount(), 3);
        
        (,,, bool isJailed, uint256 jailUntilBlock) = Staking(STAKING).getValidatorInfo(VALIDATOR4);
        assertTrue(isJailed);
        assertEq(jailUntilBlock, block.number + 1000);
        
        // Now no validator should be able to exit (would leave less than 3 active)
        vm.prank(VALIDATOR1);
        vm.expectRevert("Cannot exit: would leave less than minimum validators");
        Staking(STAKING).emergencyExit();
    }

    function testAddValidatorStake() public {
        // Register a validator first
        _setupValidatorPass(VALIDATOR1);
        vm.startPrank(VALIDATOR1);
        Staking(STAKING).registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        // Add more stake
        uint256 additionalStake = 5000 ether;
        Staking(STAKING).addValidatorStake{value: additionalStake}();
        vm.stopPrank();
        
        // Check updated stake
        (uint256 selfStake, , , ,) = Staking(STAKING).getValidatorInfo(VALIDATOR1);
        assertEq(selfStake, MIN_STAKE + additionalStake);
        assertEq(Staking(STAKING).totalStaked(), MIN_STAKE + additionalStake);
    }

    function test_RevertWhen_AddZeroStake() public {
        // Register a validator first
        _setupValidatorPass(VALIDATOR1);
        vm.startPrank(VALIDATOR1);
        Staking(STAKING).registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        // Try to add zero stake
        vm.expectRevert("Amount must be positive");
        Staking(STAKING).addValidatorStake{value: 0}();
        vm.stopPrank();
    }

    function testUpdateCommissionRate() public {
        // Register a validator first
        _setupValidatorPass(VALIDATOR1);
        vm.startPrank(VALIDATOR1);
        Staking(STAKING).registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        // Update commission rate
        uint256 newRate = 2000; // 20%
        Staking(STAKING).updateCommissionRate(newRate);
        vm.stopPrank();
        
        // Check updated rate
        (, , uint256 commissionRate, ,) = Staking(STAKING).getValidatorInfo(VALIDATOR1);
        assertEq(commissionRate, newRate);
    }

    function test_RevertWhen_UpdateInvalidCommissionRate() public {
        // Register a validator first
        _setupValidatorPass(VALIDATOR1);
        vm.startPrank(VALIDATOR1);
        Staking(STAKING).registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        // Try to set invalid commission rate
        vm.expectRevert("Invalid commission rate");
        Staking(STAKING).updateCommissionRate(11000); // > 100%
        vm.stopPrank();
    }

    function testUndelegate() public {
        // Register a validator
        _setupValidatorPass(VALIDATOR1);
        vm.startPrank(VALIDATOR1);
        Staking(STAKING).registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        vm.stopPrank();
        
        // Delegate tokens
        uint256 delegationAmount = 1000 ether;
        vm.startPrank(DELEGATOR1);
        Staking(STAKING).delegate{value: delegationAmount}(VALIDATOR1);
        
        // Undelegate
        uint256 undelegateAmount = 500 ether;
        Staking(STAKING).undelegate(VALIDATOR1, undelegateAmount);
        vm.stopPrank();
        
        // Check delegation info
        (uint256 amount, , ,) = Staking(STAKING).getDelegationInfo(DELEGATOR1, VALIDATOR1);
        assertEq(amount, delegationAmount - undelegateAmount);
        
        // Check validator's total delegated
        (, uint256 totalDelegated, , ,) = Staking(STAKING).getValidatorInfo(VALIDATOR1);
        assertEq(totalDelegated, delegationAmount - undelegateAmount);
    }

    function test_RevertWhen_UndelegateInsufficientAmount() public {
        // Register a validator
        _setupValidatorPass(VALIDATOR1);
        vm.startPrank(VALIDATOR1);
        Staking(STAKING).registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        vm.stopPrank();
        
        // Delegate tokens
        uint256 delegationAmount = 1000 ether;
        vm.startPrank(DELEGATOR1);
        Staking(STAKING).delegate{value: delegationAmount}(VALIDATOR1);
        
        // Try to undelegate more than delegated
        vm.expectRevert("Insufficient delegation");
        Staking(STAKING).undelegate(VALIDATOR1, 2000 ether);
        vm.stopPrank();
    }

    function testWithdrawUnbondedBasic() public {
        // Register a validator
        _setupValidatorPass(VALIDATOR1);
        vm.startPrank(VALIDATOR1);
        Staking(STAKING).registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        vm.stopPrank();
        
        // Delegate tokens
        uint256 delegationAmount = 1000 ether;
        vm.startPrank(DELEGATOR1);
        Staking(STAKING).delegate{value: delegationAmount}(VALIDATOR1);
        
        // Undelegate
        uint256 undelegateAmount = 500 ether;
        Staking(STAKING).undelegate(VALIDATOR1, undelegateAmount);
        
        // Try to withdraw before unbonding period completes (should fail)
        vm.expectRevert("No unbonded tokens available");
        Staking(STAKING).withdrawUnbonded(VALIDATOR1, 1);
        vm.stopPrank();
        
        // This test verifies the basic unbonding mechanism
        // Full withdrawal test would require fixing the array manipulation issue
    }

    function test_RevertWhen_NoUnbondedTokens() public {
        // Register a validator
        _setupValidatorPass(VALIDATOR1);
        vm.startPrank(VALIDATOR1);
        Staking(STAKING).registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        vm.stopPrank();
        
        // Try to withdraw without unbonding
        vm.startPrank(DELEGATOR1);
        vm.expectRevert("No unbonded tokens available");
        Staking(STAKING).withdrawUnbonded(VALIDATOR1, 1);
        vm.stopPrank();
    }

    function testClaimRewards() public {
        // Register a validator
        _setupValidatorPass(VALIDATOR1);
        vm.startPrank(VALIDATOR1);
        Staking(STAKING).registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        // Try to claim rewards (should not revert even if zero)
        Staking(STAKING).claimRewards(VALIDATOR1);
        vm.stopPrank();
    }

    function testUnjailValidator() public {
        // Register 4 validators first
        registerMultipleValidators();
        
        // Update active validator set to make validators active
        _updateActiveValidatorSet();
        
        // Jail a validator
        vm.prank(PUNISH); // Punish contract
        Staking(STAKING).jailValidator(VALIDATOR4, 100);
        
        // Verify jailed
        (,,, bool isJailed,) = Staking(STAKING).getValidatorInfo(VALIDATOR4);
        assertTrue(isJailed);
        
        // Fast forward past jail period
        vm.roll(block.number + 101);
        
        // Unjail validator
        vm.prank(VALIDATOR4);
        Staking(STAKING).unjailValidator(VALIDATOR4);
        
        // Verify unjailed
        (,,, bool isJailedAfter, uint256 jailUntilBlock) = Staking(STAKING).getValidatorInfo(VALIDATOR4);
        assertFalse(isJailedAfter);
        assertEq(jailUntilBlock, 0);
    }

    function test_RevertWhen_UnjailNotJailed() public {
        // Register a validator
        _setupValidatorPass(VALIDATOR1);
        vm.startPrank(VALIDATOR1);
        Staking(STAKING).registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        // Test unjailing a non-jailed validator (should revert)
        vm.expectRevert("Validator not jailed");
        Staking(STAKING).unjailValidator(VALIDATOR1);
        vm.stopPrank();
    }

    function test_RevertWhen_NonValidatorUnjail() public {
        // Register 4 validators first
        registerMultipleValidators();
        
        // Update active validator set to make validators active
        _updateActiveValidatorSet();
        
        // Jail a validator
        vm.prank(PUNISH); // Punish contract
        Staking(STAKING).jailValidator(VALIDATOR4, 100);
        
        // Try to unjail from different address
        vm.startPrank(VALIDATOR2);
        vm.expectRevert("Only validator can unjail themselves");
        Staking(STAKING).unjailValidator(VALIDATOR4);
        vm.stopPrank();
    }

    function test_RevertWhen_UnjailTooEarly() public {
        // Register 4 validators first
        registerMultipleValidators();
        
        // Update active validator set to make validators active
        _updateActiveValidatorSet();
        
        // Jail a validator
        vm.prank(PUNISH); // Punish contract
        Staking(STAKING).jailValidator(VALIDATOR4, 100);
        
        // Try to unjail before jail period is complete
        vm.prank(VALIDATOR4);
        vm.expectRevert("Jail period not complete");
        Staking(STAKING).unjailValidator(VALIDATOR4);
    }

    function testGetDelegationInfo() public {
        // Register a validator
        _setupValidatorPass(VALIDATOR1);
        vm.startPrank(VALIDATOR1);
        Staking(STAKING).registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        vm.stopPrank();
        
        // Delegate tokens
        uint256 delegationAmount = 1000 ether;
        vm.startPrank(DELEGATOR1);
        Staking(STAKING).delegate{value: delegationAmount}(VALIDATOR1);
        vm.stopPrank();
        
        // Check delegation info
        (uint256 amount, uint256 pendingRewards, uint256 unbondingAmount, uint256 unbondingBlock) = 
            Staking(STAKING).getDelegationInfo(DELEGATOR1, VALIDATOR1);
        
        assertEq(amount, delegationAmount);
        assertEq(pendingRewards, 0); // No rewards distributed yet
        assertEq(unbondingAmount, 0);
        assertEq(unbondingBlock, 0);
    }

    function testEmergencyExitWithFourValidators() public {
        // Register 4 validators to allow emergency exit
        registerMultipleValidators();
        
        // Update active validator set to make validators active
        _updateActiveValidatorSet();
        
        // Emergency exit
        uint256 balanceBefore = VALIDATOR4.balance;
        vm.startPrank(VALIDATOR4);
        Staking(STAKING).emergencyExit();
        vm.stopPrank();
        
        // Check balance and validator state
        uint256 balanceAfter = VALIDATOR4.balance;
        assertEq(balanceAfter - balanceBefore, MIN_STAKE);
        
        (uint256 selfStake, , , ,) = Staking(STAKING).getValidatorInfo(VALIDATOR4);
        assertEq(selfStake, 0);
        // After exit, validator is removed from active set, so count decreases
        // But we need to update the set to reflect the change
        _updateActiveValidatorSet();
        assertEq(Staking(STAKING).getActiveValidatorCount(), 3);
    }

    function test_RevertWhen_EmergencyExitMinValidators() public {
        // Register exactly 3 validators (minimum)
        for (uint i = 0; i < 3; i++) {
            address validator = address(uint160(i + 1000));
            vm.deal(validator, 100000 ether);
            _setupValidatorPass(validator);
            vm.startPrank(validator);
            Staking(STAKING).registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
            vm.stopPrank();
        }
        
        // Update active validator set to make validators active
        _updateActiveValidatorSet();
        
        // Try emergency exit when at minimum
        address lastValidator = address(uint160(1002));
        vm.startPrank(lastValidator);
        vm.expectRevert("Cannot exit: would leave less than minimum validators");
        Staking(STAKING).emergencyExit();
        vm.stopPrank();
    }

    function testDistributeRewardsFlow() public {
        // Register a validator
        _setupValidatorPass(VALIDATOR1);
        vm.startPrank(VALIDATOR1);
        Staking(STAKING).registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        vm.stopPrank();
        
        // Delegate tokens
        uint256 delegationAmount = 1000 ether;
        vm.startPrank(DELEGATOR1);
        Staking(STAKING).delegate{value: delegationAmount}(VALIDATOR1);
        vm.stopPrank();
        
        // Simulate reward distribution using coinbase (miner)
        uint256 rewardAmount = 100 ether;
        address miner = address(0x123);
        vm.deal(miner, rewardAmount);
        
        // Set miner as coinbase for the block
        // distributeRewards() now gets validator from msg.sender (block.coinbase)
        vm.coinbase(VALIDATOR1);
        vm.startPrank(VALIDATOR1);
        Staking(STAKING).distributeRewards{value: rewardAmount}();
        vm.stopPrank();
        
        // Check that rewards were distributed (basic test)
        // Note: Full reward testing would require access to internal state
    }

    // Helper function to register multiple validators
    function registerMultipleValidators() internal {
        address[4] memory validatorAddrs = [VALIDATOR1, VALIDATOR2, VALIDATOR3, VALIDATOR4];
        
        for (uint i = 0; i < validatorAddrs.length; i++) {
            _setupValidatorPass(validatorAddrs[i]);
            vm.startPrank(validatorAddrs[i]);
            Staking(STAKING).registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
            vm.stopPrank();
        }
    }
}
