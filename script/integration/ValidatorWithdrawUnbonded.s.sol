// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseTestScript} from "./BaseTestScript.s.sol";
import {console} from "forge-std/Test.sol";

// 原子化脚本：仅负责验证者提取解绑资金操作
contract ValidatorWithdrawUnbondedScript is BaseTestScript {
    function run() public override {
        console.log("Starting Validator Withdraw Unbonded Operation...");
        
        // 获取第一个验证者的密钥和地址
        uint256 validatorKey = validatorKeys[0];
        address validatorAddr = vm.addr(validatorKey);
        
        console.log("Validator address:", validatorAddr);
        
        // 执行提取解绑资金操作
        vm.startBroadcast(validatorKey);
        staking.withdrawUnbonded();
        vm.stopBroadcast();
        
        console.log("Validator withdraw unbonded transaction completed");
        
        // 打印验证者余额
        printBalance(validatorAddr);
        
        console.log("Validator Withdraw Unbonded Operation completed successfully!");
    }
}