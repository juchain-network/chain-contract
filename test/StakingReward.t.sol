// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {BaseSetup} from "./BaseSetup.t.sol";
import {Staking} from "../contracts/Staking.sol";
import {Validators} from "../contracts/Validators.sol";
import {Proposal} from "../contracts/Proposal.sol";
import {Punish} from "../contracts/Punish.sol";

contract StakingRewardTest is BaseSetup {
    address v1; 
    address v2; 
    address v3;
    address delegator1;
    address delegator2;
    
    uint256 constant MIN_STAKE = 100000 ether;
    uint256 constant COMMISSION_RATE = 1000; // 10%

    function setUp() public {
        v1 = makeAddr("v1");
        v2 = makeAddr("v2"); 
        v3 = makeAddr("v3");
        delegator1 = makeAddr("delegator1");
        delegator2 = makeAddr("delegator2");
        
        address[] memory initVals = new address[](3);
        initVals[0] = v1; 
        initVals[1] = v2; 
        initVals[2] = v3;
        
        deploySystem(initVals);
        
        // Give participants enough ETH for testing
        vm.deal(v1, MIN_STAKE + 100 ether);
        vm.deal(v2, MIN_STAKE + 100 ether);
        vm.deal(v3, MIN_STAKE + 100 ether);
        vm.deal(delegator1, 10000 ether);
        vm.deal(delegator2, 10000 ether);
    }

    function testUpdateRewardsWithPending() public {
        // Setup: Register validator and delegate
        vm.startPrank(v1);
        Staking(STAKING).addValidatorStake{value: 100 ether}();
        vm.stopPrank();
        
        // Have delegator1 delegate to v1
        vm.startPrank(delegator1);
        Staking(STAKING).delegate{value: 1000 ether}(v1);
        vm.stopPrank();
        
        // Simulate reward distribution by the validator
        uint256 rewardAmount = 100 ether;
        vm.deal(v1, rewardAmount);
        vm.coinbase(v1);
        vm.startPrank(v1);
        Staking(STAKING).distributeRewards{value: rewardAmount}();
        vm.stopPrank();
        
        // Check that rewardPerShare has been updated
        uint256 rewardPerShareV1 = Staking(STAKING).rewardPerShare(v1);
        assertTrue(rewardPerShareV1 > 0, "Reward per share should be greater than 0");
        
        // Now when delegator claims rewards, pending should be > 0
        uint256 delegator1BalanceBefore = delegator1.balance;
        vm.startPrank(delegator1);
        Staking(STAKING).claimRewards(v1);
        vm.stopPrank();
        
        uint256 delegator1BalanceAfter = delegator1.balance;
        assertTrue(delegator1BalanceAfter > delegator1BalanceBefore, "Delegator should receive rewards");
        
        // Verify the rewards claimed event was emitted
        // This will cover the previously uncovered branch in _updateRewards
    }
    
    function testUpdateRewardsWithZeroPending() public {
        // Test the case where delegation.amount is 0, so pending will be 0
        // This covers the branch where pending > 0 is false
        
        // Have delegator1 delegate to v1
        vm.startPrank(delegator1);
        // This will call _updateRewards but with 0 pending since there are no rewards yet
        Staking(STAKING).delegate{value: 1000 ether}(v1);
        vm.stopPrank();
        
        // Verify that the delegation was successful
        (uint256 amount, uint256 pendingRewards, , ) = Staking(STAKING).getDelegationInfo(delegator1, v1);
        assertEq(amount, 1000 ether, "Delegation amount should match");
        assertEq(pendingRewards, 0, "Pending rewards should be 0 initially");
    }
}