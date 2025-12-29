// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseTestScript} from "./BaseTestScript.s.sol";
import {console} from "forge-std/Test.sol";

// 原子化脚本：仅负责部署和初始化所有合约
contract ContractDeploymentScript is BaseTestScript {
    function run() public override {
        console.log("Starting Contract Deployment Operation...");
        
        // 调用BaseTestScript中的deployAndInitializeContracts函数部署合约
        deployAndInitializeContracts();
        
        console.log("Contract Deployment Operation completed successfully!");
    }
}