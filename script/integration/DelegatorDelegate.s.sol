// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseTestScript} from "./BaseTestScript.s.sol";
import {console} from "forge-std/Test.sol";

// 原子化脚本：仅负责委托者委托操作
contract DelegatorDelegateScript is BaseTestScript {
    function run() public override {
        console.log("Starting Delegator Delegate Operation...");
        
        // 直接使用已初始化的委托者账户（索引10-19为委托者账户）
        uint256 delegatorIndex = 10;
        address delegatorAddr = validatorAccounts[delegatorIndex];
        uint256 delegatorKey = validatorKeys[delegatorIndex];
        
        // 获取第一个验证者的地址
        address validatorAddr = validatorAccounts[0];
        
        console.log("Delegator address:", delegatorAddr);
        console.log("Validator address:", validatorAddr);
        
        // 委托金额
        uint256 delegationAmount = 100 ether;
        
        // 执行委托操作
        vm.startBroadcast(delegatorKey);
        staking.delegate{value: delegationAmount}(validatorAddr);
        vm.stopBroadcast();
        
        console.log("Delegator delegate transaction completed");
        
        // 验证委托状态
        (uint256 delegatorStake, , , ) = staking.getDelegationInfo(delegatorAddr, validatorAddr);
        require(delegatorStake == delegationAmount, "Delegator should have correct stake");
        
        console.log("Delegator Delegate Operation completed successfully!");
    }
}