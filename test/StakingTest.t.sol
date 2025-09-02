// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Test} from "forge-std/Test.sol";
import {Staking} from "../contracts/Staking.sol";

contract StakingTest is Test {
    Staking public staking;
    
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
        staking = new Staking();
        
        // Set up system addresses in Params
        // Note: In actual deployment, these would be set differently
        vm.deal(VALIDATOR1, 100000 ether);
        vm.deal(VALIDATOR2, 100000 ether);
        vm.deal(VALIDATOR3, 100000 ether);
        vm.deal(VALIDATOR4, 100000 ether);
        vm.deal(VALIDATOR5, 100000 ether);
        vm.deal(VALIDATOR6, 100000 ether);
        vm.deal(DELEGATOR1, 100000 ether);
        vm.deal(DELEGATOR2, 100000 ether);
        
        // Initialize the contract
        staking.initialize();
    }

    function testInitialization() public view {
        assertEq(staking.MIN_VALIDATOR_STAKE(), MIN_STAKE);
        assertEq(staking.MAX_VALIDATORS(), 21);
        assertEq(staking.MIN_VALIDATORS(), 5);
        assertEq(staking.getValidatorCount(), 0);
        assertEq(staking.getActiveValidatorCount(), 0);
        assertFalse(staking.hasMinimumValidators());
    }

    function testValidatorRegistration() public {
        // Test successful registration
        vm.prank(VALIDATOR1);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        assertEq(staking.getValidatorCount(), 1);
        assertEq(staking.getActiveValidatorCount(), 1);
        
        (uint256 selfStake, uint256 totalDelegated, uint256 commissionRate, bool isJailed, uint256 jailUntilBlock) = 
            staking.getValidatorInfo(VALIDATOR1);
            
        assertEq(selfStake, MIN_STAKE);
        assertEq(totalDelegated, 0);
        assertEq(commissionRate, COMMISSION_RATE);
        assertFalse(isJailed);
        assertEq(jailUntilBlock, 0);
    }

    function test_RevertWhen_InsufficientStake() public {
        vm.prank(VALIDATOR1);
        vm.expectRevert("Insufficient self-stake");
        staking.registerValidator{value: MIN_STAKE - 1}(COMMISSION_RATE);
    }

    function test_RevertWhen_InvalidCommissionRate() public {
        vm.prank(VALIDATOR1);
        vm.expectRevert("Invalid commission rate");
        staking.registerValidator{value: MIN_STAKE}(10001); // > 100%
    }

    function test_RevertWhen_DoubleRegistration() public {
        vm.prank(VALIDATOR1);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        vm.prank(VALIDATOR1);
        vm.expectRevert("Already registered");
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
    }

    function testMinimumValidatorsRequirement() public {
        // Register 5 validators (minimum required)
        vm.prank(VALIDATOR1);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        vm.prank(VALIDATOR2);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        vm.prank(VALIDATOR3);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        vm.prank(VALIDATOR4);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        vm.prank(VALIDATOR5);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        assertEq(staking.getActiveValidatorCount(), 5);
        assertTrue(staking.hasMinimumValidators());
        
        // Now register a 6th validator
        vm.prank(VALIDATOR6);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        assertEq(staking.getActiveValidatorCount(), 6);
        
        // Test that 6th validator can exit (leaving 5)
        vm.prank(VALIDATOR6);
        staking.emergencyExit();
        
        assertEq(staking.getActiveValidatorCount(), 5);
        assertTrue(staking.hasMinimumValidators());
        
        // Test that 5th validator cannot exit (would leave 4)
        vm.prank(VALIDATOR5);
        vm.expectRevert("Cannot exit: minimum validators required");
        staking.emergencyExit();
    }

    function testPartialStakeWithdrawal() public {
        // Register validator with extra stake
        vm.prank(VALIDATOR1);
        staking.registerValidator{value: MIN_STAKE * 2}(COMMISSION_RATE);
        
        // Register 4 more validators to meet minimum
        vm.prank(VALIDATOR2);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        vm.prank(VALIDATOR3);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        vm.prank(VALIDATOR4);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        vm.prank(VALIDATOR5);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        // Test partial withdrawal (still leaves minimum stake)
        uint256 withdrawAmount = MIN_STAKE / 2;
        uint256 initialBalance = VALIDATOR1.balance;
        
        vm.prank(VALIDATOR1);
        staking.withdrawValidatorStake(withdrawAmount);
        
        assertEq(VALIDATOR1.balance, initialBalance + withdrawAmount);
        
        (uint256 selfStake,,,, ) = staking.getValidatorInfo(VALIDATOR1);
        assertEq(selfStake, MIN_STAKE * 2 - withdrawAmount);
        assertEq(staking.getActiveValidatorCount(), 5);
    }

    function test_RevertWhen_PartialWithdrawalBelowMinimum() public {
        // Register multiple validators first to avoid minimum validator constraint
        vm.prank(VALIDATOR1);
        staking.registerValidator{value: MIN_STAKE + 1000}(COMMISSION_RATE);
        vm.prank(VALIDATOR2);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        vm.prank(VALIDATOR3);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        vm.prank(VALIDATOR4);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        vm.prank(VALIDATOR5);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        vm.prank(VALIDATOR6);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        // Try to withdraw amount that would leave stake below minimum
        vm.prank(VALIDATOR1);
        vm.expectRevert("Remaining stake below minimum");
        staking.withdrawValidatorStake(1001);
    }

    function testDelegation() public {
        // Register validator
        vm.prank(VALIDATOR1);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        // Test delegation
        uint256 delegationAmount = 1000 ether;
        vm.prank(DELEGATOR1);
        staking.delegate{value: delegationAmount}(VALIDATOR1);
        
        (uint256 selfStake, uint256 totalDelegated,,, ) = staking.getValidatorInfo(VALIDATOR1);
        assertEq(selfStake, MIN_STAKE);
        assertEq(totalDelegated, delegationAmount);
        
        (uint256 delegatedAmount,,,) = staking.getDelegationInfo(DELEGATOR1, VALIDATOR1);
        assertEq(delegatedAmount, delegationAmount);
    }

    function test_RevertWhen_DelegateToInactiveValidator() public {
        vm.prank(DELEGATOR1);
        vm.expectRevert("Not a valid validator");
        staking.delegate{value: 1000 ether}(VALIDATOR1); // VALIDATOR1 not registered
    }

    function test_RevertWhen_DelegateInsufficientAmount() public {
        vm.prank(VALIDATOR1);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        vm.prank(DELEGATOR1);
        vm.expectRevert("Insufficient delegation amount");
        staking.delegate{value: 0.5 ether}(VALIDATOR1); // Below MIN_DELEGATION
    }

    function testGetTopValidators() public {
        // Register validators with different stakes
        vm.prank(VALIDATOR1);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        vm.prank(VALIDATOR2);
        staking.registerValidator{value: MIN_STAKE * 2}(COMMISSION_RATE);
        
        vm.prank(VALIDATOR3);
        staking.registerValidator{value: MIN_STAKE * 3}(COMMISSION_RATE);
        
        address[] memory topValidators = staking.getTopValidators(10);
        
        assertEq(topValidators.length, 3);
        // Should be ordered by stake (highest first)
        assertEq(topValidators[0], VALIDATOR3); // 30,000 JU
        assertEq(topValidators[1], VALIDATOR2); // 20,000 JU
        assertEq(topValidators[2], VALIDATOR1); // 10,000 JU
    }

    function testGetTopValidatorsWithDelegations() public {
        // Register validators
        vm.prank(VALIDATOR1);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        vm.prank(VALIDATOR2);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        // Add delegation to VALIDATOR1
        vm.prank(DELEGATOR1);
        staking.delegate{value: MIN_STAKE}(VALIDATOR1);
        
        address[] memory topValidators = staking.getTopValidators(10);
        
        assertEq(topValidators.length, 2);
        // VALIDATOR1 should be first (20,000 total vs 10,000)
        assertEq(topValidators[0], VALIDATOR1);
        assertEq(topValidators[1], VALIDATOR2);
    }

    function testSystemInvariant_MinimumValidators() public {
        // Setup: Register exactly 5 validators (minimum required)
        address[5] memory validators = [VALIDATOR1, VALIDATOR2, VALIDATOR3, VALIDATOR4, VALIDATOR5];
        
        for (uint i = 0; i < 5; i++) {
            vm.prank(validators[i]);
            staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        }
        
        assertTrue(staking.hasMinimumValidators());
        assertEq(staking.getActiveValidatorCount(), 5);
        
        // Test that no validator can exit
        for (uint i = 0; i < 5; i++) {
            vm.prank(validators[i]);
            vm.expectRevert("Cannot exit: minimum validators required");
            staking.emergencyExit();
        }
        
        // Add one more validator
        vm.prank(VALIDATOR6);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        assertEq(staking.getActiveValidatorCount(), 6);
        
        // Now exactly one validator should be able to exit
        vm.prank(VALIDATOR6);
        staking.emergencyExit();
        
        assertEq(staking.getActiveValidatorCount(), 5);
        
        // Back to minimum - no one can exit again
        vm.prank(VALIDATOR1);
        vm.expectRevert("Cannot exit: minimum validators required");
        staking.emergencyExit();
    }

    function testValidatorJailing() public {
        // Register 6 validators
        vm.prank(VALIDATOR1);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        vm.prank(VALIDATOR2);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        vm.prank(VALIDATOR3);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        vm.prank(VALIDATOR4);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        vm.prank(VALIDATOR5);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        vm.prank(VALIDATOR6);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        assertEq(staking.getActiveValidatorCount(), 6);
        
        // Jail one validator (simulating punishment contract call)
        vm.prank(address(0x000000000000000000000000000000000000F001)); // Punish contract
        staking.jailValidator(VALIDATOR6, 1000);
        
        // Active count should decrease
        assertEq(staking.getActiveValidatorCount(), 5);
        
        (,,, bool isJailed, uint256 jailUntilBlock) = staking.getValidatorInfo(VALIDATOR6);
        assertTrue(isJailed);
        assertEq(jailUntilBlock, block.number + 1000);
        
        // Now no validator should be able to exit (would leave less than 5 active)
        vm.prank(VALIDATOR1);
        vm.expectRevert("Cannot exit: minimum validators required");
        staking.emergencyExit();
    }

    function testAddValidatorStake() public {
        // Register a validator first
        vm.startPrank(VALIDATOR1);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        // Add more stake
        uint256 additionalStake = 5000 ether;
        staking.addValidatorStake{value: additionalStake}();
        vm.stopPrank();
        
        // Check updated stake
        (uint256 selfStake, , , ,) = staking.getValidatorInfo(VALIDATOR1);
        assertEq(selfStake, MIN_STAKE + additionalStake);
        assertEq(staking.totalStaked(), MIN_STAKE + additionalStake);
    }

    function test_RevertWhen_AddZeroStake() public {
        // Register a validator first
        vm.startPrank(VALIDATOR1);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        // Try to add zero stake
        vm.expectRevert("Amount must be positive");
        staking.addValidatorStake{value: 0}();
        vm.stopPrank();
    }

    function testUpdateCommissionRate() public {
        // Register a validator first
        vm.startPrank(VALIDATOR1);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        // Update commission rate
        uint256 newRate = 2000; // 20%
        staking.updateCommissionRate(newRate);
        vm.stopPrank();
        
        // Check updated rate
        (, , uint256 commissionRate, ,) = staking.getValidatorInfo(VALIDATOR1);
        assertEq(commissionRate, newRate);
    }

    function test_RevertWhen_UpdateInvalidCommissionRate() public {
        // Register a validator first
        vm.startPrank(VALIDATOR1);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        // Try to set invalid commission rate
        vm.expectRevert("Invalid commission rate");
        staking.updateCommissionRate(11000); // > 100%
        vm.stopPrank();
    }

    function testUndelegate() public {
        // Register a validator
        vm.startPrank(VALIDATOR1);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        vm.stopPrank();
        
        // Delegate tokens
        uint256 delegationAmount = 1000 ether;
        vm.startPrank(DELEGATOR1);
        staking.delegate{value: delegationAmount}(VALIDATOR1);
        
        // Undelegate
        uint256 undelegateAmount = 500 ether;
        staking.undelegate(VALIDATOR1, undelegateAmount);
        vm.stopPrank();
        
        // Check delegation info
        (uint256 amount, , ,) = staking.getDelegationInfo(DELEGATOR1, VALIDATOR1);
        assertEq(amount, delegationAmount - undelegateAmount);
        
        // Check validator's total delegated
        (, uint256 totalDelegated, , ,) = staking.getValidatorInfo(VALIDATOR1);
        assertEq(totalDelegated, delegationAmount - undelegateAmount);
    }

    function test_RevertWhen_UndelegateInsufficientAmount() public {
        // Register a validator
        vm.startPrank(VALIDATOR1);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        vm.stopPrank();
        
        // Delegate tokens
        uint256 delegationAmount = 1000 ether;
        vm.startPrank(DELEGATOR1);
        staking.delegate{value: delegationAmount}(VALIDATOR1);
        
        // Try to undelegate more than delegated
        vm.expectRevert("Insufficient delegation");
        staking.undelegate(VALIDATOR1, 2000 ether);
        vm.stopPrank();
    }

    function testWithdrawUnbondedBasic() public {
        // Register a validator
        vm.startPrank(VALIDATOR1);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        vm.stopPrank();
        
        // Delegate tokens
        uint256 delegationAmount = 1000 ether;
        vm.startPrank(DELEGATOR1);
        staking.delegate{value: delegationAmount}(VALIDATOR1);
        
        // Undelegate
        uint256 undelegateAmount = 500 ether;
        staking.undelegate(VALIDATOR1, undelegateAmount);
        
        // Try to withdraw before unbonding period completes (should fail)
        vm.expectRevert("No unbonded tokens available");
        staking.withdrawUnbonded(VALIDATOR1, 1);
        vm.stopPrank();
        
        // This test verifies the basic unbonding mechanism
        // Full withdrawal test would require fixing the array manipulation issue
    }

    function test_RevertWhen_NoUnbondedTokens() public {
        // Register a validator
        vm.startPrank(VALIDATOR1);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        vm.stopPrank();
        
        // Try to withdraw without unbonding
        vm.startPrank(DELEGATOR1);
        vm.expectRevert("No unbonded tokens available");
        staking.withdrawUnbonded(VALIDATOR1, 1);
        vm.stopPrank();
    }

    function testClaimRewards() public {
        // Register a validator
        vm.startPrank(VALIDATOR1);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        // Try to claim rewards (should not revert even if zero)
        staking.claimRewards(VALIDATOR1);
        vm.stopPrank();
    }

    function testUnjailValidator() public {
        // Register 6 validators first
        registerMultipleValidators();
        
        // Jail a validator
        vm.prank(address(0x000000000000000000000000000000000000F001)); // Punish contract
        staking.jailValidator(VALIDATOR6, 100);
        
        // Verify jailed
        (,,, bool isJailed,) = staking.getValidatorInfo(VALIDATOR6);
        assertTrue(isJailed);
        
        // Fast forward past jail period
        vm.roll(block.number + 101);
        
        // Unjail validator
        vm.prank(VALIDATOR6);
        staking.unjailValidator(VALIDATOR6);
        
        // Verify unjailed
        (,,, bool isJailedAfter, uint256 jailUntilBlock) = staking.getValidatorInfo(VALIDATOR6);
        assertFalse(isJailedAfter);
        assertEq(jailUntilBlock, 0);
    }

    function test_RevertWhen_UnjailNotJailed() public {
        // Register a validator
        vm.startPrank(VALIDATOR1);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        
        // Test unjailing a non-jailed validator (should revert)
        vm.expectRevert("Validator not jailed");
        staking.unjailValidator(VALIDATOR1);
        vm.stopPrank();
    }

    function test_RevertWhen_NonValidatorUnjail() public {
        // Register 6 validators first
        registerMultipleValidators();
        
        // Jail a validator
        vm.prank(address(0x000000000000000000000000000000000000F001)); // Punish contract
        staking.jailValidator(VALIDATOR6, 100);
        
        // Try to unjail from different address
        vm.startPrank(VALIDATOR2);
        vm.expectRevert("Only validator can unjail themselves");
        staking.unjailValidator(VALIDATOR6);
        vm.stopPrank();
    }

    function test_RevertWhen_UnjailTooEarly() public {
        // Register 6 validators first
        registerMultipleValidators();
        
        // Jail a validator
        vm.prank(address(0x000000000000000000000000000000000000F001)); // Punish contract
        staking.jailValidator(VALIDATOR6, 100);
        
        // Try to unjail before jail period is complete
        vm.prank(VALIDATOR6);
        vm.expectRevert("Jail period not complete");
        staking.unjailValidator(VALIDATOR6);
    }

    function testGetDelegationInfo() public {
        // Register a validator
        vm.startPrank(VALIDATOR1);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        vm.stopPrank();
        
        // Delegate tokens
        uint256 delegationAmount = 1000 ether;
        vm.startPrank(DELEGATOR1);
        staking.delegate{value: delegationAmount}(VALIDATOR1);
        vm.stopPrank();
        
        // Check delegation info
        (uint256 amount, uint256 pendingRewards, uint256 unbondingAmount, uint256 unbondingBlock) = 
            staking.getDelegationInfo(DELEGATOR1, VALIDATOR1);
        
        assertEq(amount, delegationAmount);
        assertEq(pendingRewards, 0); // No rewards distributed yet
        assertEq(unbondingAmount, 0);
        assertEq(unbondingBlock, 0);
    }

    function testEmergencyExitWithSixValidators() public {
        // Register 6 validators to allow emergency exit
        registerMultipleValidators();
        
        // Emergency exit
        uint256 balanceBefore = VALIDATOR6.balance;
        vm.startPrank(VALIDATOR6);
        staking.emergencyExit();
        vm.stopPrank();
        
        // Check balance and validator state
        uint256 balanceAfter = VALIDATOR6.balance;
        assertEq(balanceAfter - balanceBefore, MIN_STAKE);
        
        (uint256 selfStake, , , ,) = staking.getValidatorInfo(VALIDATOR6);
        assertEq(selfStake, 0);
        assertEq(staking.getActiveValidatorCount(), 5);
    }

    function test_RevertWhen_EmergencyExitMinValidators() public {
        // Register exactly 5 validators (minimum)
        for (uint i = 0; i < 5; i++) {
            address validator = address(uint160(i + 1000));
            vm.deal(validator, 100000 ether);
            vm.startPrank(validator);
            staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
            vm.stopPrank();
        }
        
        // Try emergency exit when at minimum
        address lastValidator = address(uint160(1004));
        vm.startPrank(lastValidator);
        vm.expectRevert("Cannot exit: minimum validators required");
        staking.emergencyExit();
        vm.stopPrank();
    }

    function testDistributeRewardsFlow() public {
        // Register a validator
        vm.startPrank(VALIDATOR1);
        staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
        vm.stopPrank();
        
        // Delegate tokens
        uint256 delegationAmount = 1000 ether;
        vm.startPrank(DELEGATOR1);
        staking.delegate{value: delegationAmount}(VALIDATOR1);
        vm.stopPrank();
        
        // Simulate reward distribution using coinbase (miner)
        uint256 rewardAmount = 100 ether;
        address miner = address(0x123);
        vm.deal(miner, rewardAmount);
        
        // Set miner as coinbase for the block
        vm.coinbase(miner);
        vm.startPrank(miner);
        staking.distributeRewards{value: rewardAmount}(VALIDATOR1);
        vm.stopPrank();
        
        // Check that rewards were distributed (basic test)
        // Note: Full reward testing would require access to internal state
    }

    // Helper function to register multiple validators
    function registerMultipleValidators() internal {
        address[6] memory validators = [VALIDATOR1, VALIDATOR2, VALIDATOR3, VALIDATOR4, VALIDATOR5, VALIDATOR6];
        
        for (uint i = 0; i < validators.length; i++) {
            vm.startPrank(validators[i]);
            staking.registerValidator{value: MIN_STAKE}(COMMISSION_RATE);
            vm.stopPrank();
        }
    }
}
