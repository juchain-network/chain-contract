// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseTestScript} from "./BaseTestScript.s.sol";
import {console} from "forge-std/Test.sol";

// 原子化脚本：仅负责委托者提取解绑资金操作
contract DelegatorWithdrawUnbondedScript is BaseTestScript {
    function run() public override {
        console.log("Starting Delegator Withdraw Unbonded Operation...");
        
        // 创建委托者账户（与Delegate脚本中使用相同的密钥）
        uint256 delegatorKey = uint256(keccak256(abi.encodePacked("testDelegator1")));
        address delegatorAddr = vm.addr(delegatorKey);
        
        console.log("Delegator address:", delegatorAddr);
        
        // 执行提取解绑资金操作
        vm.startBroadcast(delegatorKey);
        staking.withdrawUnbonded();
        vm.stopBroadcast();
        
        console.log("Delegator withdraw unbonded transaction completed");
        
        // 打印委托者余额
        printBalance(delegatorAddr);
        
        console.log("Delegator Withdraw Unbonded Operation completed successfully!");
    }
}