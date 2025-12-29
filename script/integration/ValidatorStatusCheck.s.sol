// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseTestScript} from "./BaseTestScript.s.sol";
import {console} from "forge-std/Test.sol";

// 原子化脚本：仅负责检查验证者状态
contract ValidatorStatusCheckScript is BaseTestScript {
    function run() public override {
        console.log("Starting Validator Status Check Operation...");
        
        // 检查所有初始验证者的状态
        for (uint256 i = 0; i < validatorAccounts.length; i++) {
            address validator = validatorAccounts[i];
            console.log("\nChecking validator", i+1, ":", validator);
            
            // 获取验证者信息
            (uint256 selfStake, bool isActive, uint256 commissionRate, 
             uint256 delegatedStake, uint256 totalStake, 
             uint256 rewards, bool isJailed, uint256 jailedUntil) = 
                staking.getValidatorInfo(validator);
            
            console.log("  Self Stake:", selfStake / 1 ether, "ETH");
            console.log("  Is Active:", isActive);
            console.log("  Commission Rate:", commissionRate / 100, "%");
            console.log("  Delegated Stake:", delegatedStake / 1 ether, "ETH");
            console.log("  Total Stake:", totalStake / 1 ether, "ETH");
            console.log("  Rewards:", rewards / 1 ether, "ETH");
            console.log("  Is Jailed:", isJailed);
            console.log("  Jailed Until:", jailedUntil);
            
            // 验证基本状态
            require(selfStake >= MIN_SELF_STAKE, "Validator should have minimum self stake");
        }
        
        console.log("\nValidator Status Check Operation completed successfully!");
    }
}