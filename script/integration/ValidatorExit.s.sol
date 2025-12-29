// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseTestScript} from "./BaseTestScript.s.sol";
import {console} from "forge-std/Test.sol";

// 原子化脚本：仅负责验证者退出操作
contract ValidatorExitScript is BaseTestScript {
    function run() public override {
        console.log("Starting Validator Exit Operation...");
        
        // 获取第一个验证者的密钥和地址
        uint256 validatorKey = validatorKeys[0];
        address validatorAddr = vm.addr(validatorKey);
        
        console.log("Validator address:", validatorAddr);
        
        // 执行退出操作
        vm.startBroadcast(validatorKey);
        staking.exitValidator();
        vm.stopBroadcast();
        
        console.log("Validator exit transaction completed");
        
        // 验证退出状态
        (, bool isActive, , , , , , ) = staking.getValidatorInfo(validatorAddr);
        console.log("Validator active status after exit:", isActive);
        
        console.log("Validator Exit Operation completed successfully!");
    }
}