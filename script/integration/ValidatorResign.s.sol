// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseTestScript} from "./BaseTestScript.s.sol";
import {console} from "forge-std/Test.sol";

// 原子化脚本：仅负责验证者辞职操作
contract ValidatorResignScript is BaseTestScript {
    function run() public override {
        console.log("Starting Validator Resign Operation...");
        
        // 获取第一个验证者的密钥和地址
        uint256 validatorKey = validatorKeys[0];
        address validatorAddr = vm.addr(validatorKey);
        
        console.log("Validator address:", validatorAddr);
        
        // 执行辞职操作
        vm.startBroadcast(validatorKey);
        staking.resignValidator();
        vm.stopBroadcast();
        
        console.log("Validator resignation transaction completed");
        
        // 验证辞职状态
        (, bool isActive, , , , , , ) = staking.getValidatorInfo(validatorAddr);
        console.log("Validator active status after resignation:", isActive);
        
        console.log("Validator Resign Operation completed successfully!");
    }
}