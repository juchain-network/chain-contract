// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseTestScript} from "./BaseTestScript.s.sol";
import {console} from "forge-std/Test.sol";

// 原子化脚本：仅负责创建和初始化测试账户
contract AccountSetupScript is BaseTestScript {
    function run() public override {
        console.log("Starting Account Setup Operation...");
        
        // 调用BaseTestScript中的createTestAccounts函数创建测试账户
        createTestAccounts();
        
        console.log("Account Setup Operation completed successfully!");
    }
}