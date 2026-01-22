// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseTestUtils} from "./BaseTestUtils.s.sol";
import {console} from "forge-std/console.sol";

// Utility contract for staking-related operations
contract StakingUtils is BaseTestUtils {
    // Withdraw unbonded funds
    function withdrawUnbonded(uint256 delegatorKey, address validatorAddr, uint256 maxEntries) public {
        loadState();
        address delegatorAddr = vm.addr(delegatorKey);
        
        vm.startBroadcast(delegatorKey);
        staking.withdrawUnbonded(validatorAddr, maxEntries);
        vm.stopBroadcast();
        
        console.log("Withdraw unbonded completed for delegator", delegatorAddr, "from validator", validatorAddr);
    }
    
    // Claim rewards
    // function claimRewards(uint256 delegatorKey, address validatorAddr) public {
    //     loadState();
    //     address delegatorAddr = vm.addr(delegatorKey);
        
    //     vm.startBroadcast(delegatorKey);
    //     staking.claimRewards(validatorAddr);
    //     vm.stopBroadcast();
        
    //     console.log("Rewards claimed for delegator", delegatorAddr, "from validator", validatorAddr);
    // }


    // Add stake to validator
    function addStake(uint256 validatorKey, address validatorAddr, uint256 amount) public {
        loadState();
        console.log("Adding stake to validator:", validatorAddr);
        console.log("Amount to add:", amount / 1 ether, "ETH");
        
        // Execute add stake operation
        vm.startBroadcast(validatorKey);
        staking.addValidatorStake{value: amount}();
        vm.stopBroadcast();
        
        console.log("Add stake transaction completed");
        
        // Verify stake amount
        (uint256 selfStake, , , , , , , , , ) = staking.getValidatorInfo(validatorAddr);
        console.log("Updated self stake:", selfStake / 1 ether, "ETH");
    }
    
    // Decrease stake from validator
    function decreaseStake(uint256 validatorKey, address validatorAddr, uint256 amount) public {
        loadState();
        console.log("Decreasing stake from validator:", validatorAddr);
        console.log("Amount to decrease:", amount / 1 ether, "ETH");
        
        // Execute decrease stake operation
        vm.startBroadcast(validatorKey);
        staking.decreaseValidatorStake(amount);
        vm.stopBroadcast();
        
        console.log("Decrease stake transaction completed");
        
        // Verify stake amount
        (uint256 selfStake, , , , , , , , , ) = staking.getValidatorInfo(validatorAddr);
        console.log("Updated self stake:", selfStake / 1 ether, "ETH");
    }
    
    // Set validator commission rate
    function setCommissionRate(uint256 validatorKey, address validatorAddr, uint256 rate) public {
        loadState();
        console.log("Setting commission rate for validator:", validatorAddr);
        console.log("New commission rate:", rate / 100, "%");
        
        // Execute set commission rate operation
        vm.startBroadcast(validatorKey);
        staking.updateCommissionRate(rate);
        vm.stopBroadcast();
        
        console.log("Set commission rate transaction completed");
        
        // Verify commission rate
        (, , , , , uint256 commissionRate, , , , ) = staking.getValidatorInfo(validatorAddr);
        console.log("Updated commission rate:", commissionRate / 100, "%");
    }
    
    // Claim validator rewards
    function claimValidatorRewards(uint256 validatorKey, address validatorAddr) public {
        loadState();
        console.log("Claiming rewards for validator:", validatorAddr);
        
        // Execute claim rewards operation for validator
        vm.startBroadcast(validatorKey);
        staking.claimRewards(validatorAddr);
        vm.stopBroadcast();
        
        console.log("Validator rewards claimed successfully");
    }
    
    // Distribute rewards to validators
    function distributeRewards(uint256 minerKey, uint256 blockReward) public {
        loadState();
        address minerAddr = vm.addr(minerKey);
        
        console.log("Miner address:", minerAddr);
        console.log("Distributing rewards:", blockReward / 1 ether, "ETH");
        
        // Execute distributeRewards operation
        vm.startBroadcast(minerKey);
        staking.distributeRewards{value: blockReward}();
        vm.stopBroadcast();
        
        console.log("Rewards distributed successfully");
    }
}
