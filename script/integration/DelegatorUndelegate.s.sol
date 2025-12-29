// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseTestScript} from "./BaseTestScript.s.sol";
import {console} from "forge-std/Test.sol";

// 原子化脚本：仅负责委托者解绑操作
contract DelegatorUndelegateScript is BaseTestScript {
    function run() public override {
        console.log("Starting Delegator Undelegate Operation...");
        
        // 创建委托者账户（与Delegate脚本中使用相同的密钥）
        uint256 delegatorKey = uint256(keccak256(abi.encodePacked("testDelegator1")));
        address delegatorAddr = vm.addr(delegatorKey);
        
        // 获取第一个验证者的地址
        address validatorAddr = validatorAccounts[0];
        
        console.log("Delegator address:", delegatorAddr);
        console.log("Validator address:", validatorAddr);
        
        // 获取当前委托金额
        (uint256 delegatorStake, , , ) = staking.getDelegationInfo(delegatorAddr, validatorAddr);
        require(delegatorStake > 0, "Delegator should have existing stake");
        
        console.log("Current delegation amount:", delegatorStake / 1 ether, "ETH");
        
        // 解绑金额（全部解绑）
        uint256 undelegateAmount = delegatorStake;
        
        // 执行解绑操作
        vm.startBroadcast(delegatorKey);
        staking.undelegate(validatorAddr, undelegateAmount);
        vm.stopBroadcast();
        
        console.log("Delegator undelegate transaction completed");
        
        console.log("Delegator Undelegate Operation completed successfully!");
    }
}