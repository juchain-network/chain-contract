// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseTestScript} from "./BaseTestScript.s.sol";
import {console} from "forge-std/Test.sol";

// 原子化脚本：仅负责委托者领取奖励操作
contract DelegatorClaimRewardsScript is BaseTestScript {
    function run() public override {
        console.log("Starting Delegator Claim Rewards Operation...");
        
        // 创建委托者账户（与Delegate脚本中使用相同的密钥）
        uint256 delegatorKey = uint256(keccak256(abi.encodePacked("testDelegator1")));
        address delegatorAddr = vm.addr(delegatorKey);
        
        // 获取第一个验证者的地址
        address validatorAddr = validatorAccounts[0];
        
        console.log("Delegator address:", delegatorAddr);
        console.log("Validator address:", validatorAddr);
        
        // 执行领取奖励操作
        vm.startBroadcast(delegatorKey);
        staking.claimRewards(validatorAddr);
        vm.stopBroadcast();
        
        console.log("Delegator claim rewards transaction completed");
        
        // 打印委托者余额
        printBalance(delegatorAddr);
        
        console.log("Delegator Claim Rewards Operation completed successfully!");
    }
}