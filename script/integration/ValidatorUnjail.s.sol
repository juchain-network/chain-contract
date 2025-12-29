// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseTestScript} from "./BaseTestScript.s.sol";
import {console} from "forge-std/Test.sol";

// 原子化脚本：仅负责验证者解除监禁操作
contract ValidatorUnjailScript is BaseTestScript {
    function run() public override {
        console.log("Starting Validator Unjail Operation...");
        
        // 获取第一个验证者的密钥和地址
        uint256 validatorKey = validatorKeys[0];
        address validatorAddr = vm.addr(validatorKey);
        
        console.log("Validator address:", validatorAddr);
        
        // 执行解除监禁操作
        vm.startBroadcast(validatorKey);
        staking.unjailValidator();
        vm.stopBroadcast();
        
        console.log("Validator unjail transaction completed");
        
        // 检查验证者是否已解除监禁
        bool isJailed = staking.isValidatorJailed(validatorAddr);
        console.log("Validator jailed status after unjail:", isJailed);
        
        console.log("Validator Unjail Operation completed successfully!");
    }
}