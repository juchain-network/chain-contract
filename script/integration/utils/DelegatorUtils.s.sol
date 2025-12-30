// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseTestUtils} from "./BaseTestUtils.s.sol";
import {console} from "forge-std/console.sol";

// Utility contract for delegator-related operations
contract DelegatorUtils is BaseTestUtils {
    // Delegate tokens to a validator
    function delegate(uint256 delegatorKey, address validatorAddr, uint256 amount) public {
        loadState();
        address delegatorAddr = vm.addr(delegatorKey);
        
        console.log("Delegator address:", delegatorAddr);
        console.log("Delegating to validator:", validatorAddr);
        console.log("Delegation amount:", amount / 1 ether, "ETH");
        
        // Execute delegation
        vm.startBroadcast(delegatorKey);
        staking.delegate{value: amount}(validatorAddr);
        vm.stopBroadcast();
        
        console.log("Delegation transaction completed");
        
        // Print delegator balance
        printBalance(delegatorAddr);
    }
    
    // Undelegate tokens from a validator
    function undelegate(uint256 delegatorKey, address validatorAddr, uint256 amount) public {
        loadState();
        address delegatorAddr = vm.addr(delegatorKey);
        
        console.log("Delegator address:", delegatorAddr);
        console.log("Undelegating from validator:", validatorAddr);
        console.log("Undelegation amount:", amount / 1 ether, "ETH");
        
        // Execute undelegation
        vm.startBroadcast(delegatorKey);
        staking.undelegate(validatorAddr, amount);
        vm.stopBroadcast();
        
        console.log("Undelegation transaction completed");
        
        // Print delegator balance
        printBalance(delegatorAddr);
    }
    
    // Withdraw unbonded funds
    function withdrawUnbonded(uint256 delegatorKey, address validatorAddr) public {
        loadState();
        address delegatorAddr = vm.addr(delegatorKey);
        
        console.log("Delegator address:", delegatorAddr);
        console.log("Withdrawing unbonded from validator:", validatorAddr);
        
        // Execute withdraw unbonded operation
        vm.startBroadcast(delegatorKey);
        staking.withdrawUnbonded(validatorAddr, 100);
        vm.stopBroadcast();
        
        console.log("Delegator withdraw unbonded transaction completed");
        
        // Print delegator balance
        printBalance(delegatorAddr);
    }
    
    // Claim rewards for delegator
    function claimRewards(uint256 delegatorKey, address validatorAddr) public {
        loadState();
        address delegatorAddr = vm.addr(delegatorKey);
        
        console.log("Delegator address:", delegatorAddr);
        console.log("Claiming rewards from validator:", validatorAddr);
        
        // Execute claim rewards operation
        vm.startBroadcast(delegatorKey);
        staking.claimRewards(validatorAddr);
        vm.stopBroadcast();
        
        console.log("Delegator claim rewards transaction completed");
        
        // Print delegator balance
        printBalance(delegatorAddr);
    }
    
    // Check delegator status
    function statusCheck(address delegatorAddr) public {
        loadState();
        console.log("Checking delegator:", delegatorAddr);
        
        // Note: We don't store validator accounts in state anymore
        // This function should be called with specific validator addresses as needed
        // For now, we'll just print a message indicating this change
        console.log("\nDelegator Status Check: Validator accounts are no longer stored in state.");
        console.log("Use getDelegatorStatus(address delegatorAddr, address validatorAddr) for specific validators.");
        
        console.log("\nDelegator Status Check completed successfully!");
    }
    
    // Check delegator status for a specific validator
    function getDelegatorStatus(address delegatorAddr, address validatorAddr) public {
        loadState();
        (uint256 delegatorStake, uint256 rewards, uint256 withdrawableAmount, uint256 unbondingTimestamp) = 
            staking.delegations(delegatorAddr, validatorAddr);
        
        console.log("Validator:", validatorAddr);
        console.log("  Delegated Stake:", delegatorStake / 1 ether, "ETH");
        console.log("  Rewards:", rewards / 1 ether, "ETH");
        console.log("  Withdrawable Amount:", withdrawableAmount / 1 ether, "ETH");
        console.log("  Unbonding Timestamp:", unbondingTimestamp);
    }
}
