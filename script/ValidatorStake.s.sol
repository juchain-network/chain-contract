// SPDX-License-Identifier: MIT
pragma solidity 0.8.29;

import {Script, console} from "lib/forge-std/src/Script.sol";
import {Staking} from "../contracts/Staking.sol";

/**
 * @title ValidatorStakeScript
 * @dev Validator staking script - each validator runs independently
 */
contract ValidatorStakeScript is Script {
    
    function setUp() public {}

    /**
     * @dev Validator staking function
     * Each validator runs this script with their own private key
     */
    function run() external {
        // Get validator private key from environment variables
        uint256 validatorPrivateKey = vm.envOr("VALIDATOR_PRIVATE_KEY", uint256(0));
        require(validatorPrivateKey != 0, "VALIDATOR_PRIVATE_KEY not set");
        
        // Get Staking contract address from environment variables
        address stakingAddress = vm.envOr("STAKING_CONTRACT", address(0));
        require(stakingAddress != address(0), "STAKING_CONTRACT address not set");
        
        address validator = vm.addr(validatorPrivateKey);
        uint256 stakeAmount = 10000 ether; // 10,000 JU
        uint256 commissionRate = vm.envOr("COMMISSION_RATE", uint256(500)); // Default 5%
        
        console.log("=== Validator Staking ===");
        console.log("Validator address:", validator);
        console.log("Staking contract:", stakingAddress);
        console.log("Stake amount:", stakeAmount / 1 ether, "JU");
        console.log("Commission rate:", commissionRate, "/10000");
        
        // Check balance
        require(validator.balance >= stakeAmount, "Insufficient balance for staking");
        
        vm.startBroadcast(validatorPrivateKey);
        
        // Register and stake
        Staking(stakingAddress).registerValidator{value: stakeAmount}(commissionRate);
        
        vm.stopBroadcast();
        
        console.log("Successfully registered as validator!");
        console.log("Staked:", stakeAmount / 1 ether, "JU");
        
        // Verify registration status
        checkRegistrationStatus(stakingAddress, validator);
    }
    
    /**
     * @dev Check validator registration status
     */
    function checkRegistrationStatus(address stakingAddress, address validator) internal view {
        Staking staking = Staking(stakingAddress);
        
        // Check total staked amount
        uint256 totalStaked = staking.totalStaked();
        console.log("Total network stake:", totalStaked / 1 ether, "JU");
        
        // Check validator count
        uint256 validatorCount = staking.getValidatorCount();
        console.log("Total validators:", validatorCount);
        
        console.log("Validator", validator, "is now active in the staking contract");
    }
}
