// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseTestScript} from "./BaseTestScript.s.sol";
import {console} from "forge-std/Test.sol";

// 原子化脚本：仅负责惩罚验证者操作
contract ValidatorPunishScript is BaseTestScript {
    function run() public override {
        console.log("Starting Validator Punish Operation...");
        
        // 获取第一个验证者的地址
        address validator = validatorAccounts[0];
        
        // 使用第二个验证者作为矿工执行惩罚操作
        uint256 minerKey = validatorKeys[1];
        
        console.log("Punishing validator:", validator);
        console.log("Miner address:", vm.addr(minerKey));
        
        // 执行惩罚操作
        vm.startBroadcast(minerKey);
        punish.punish(validator); // 惩罚验证者（错过区块）
        vm.stopBroadcast();
        
        console.log("Validator punishment transaction completed");
        
        // 检查验证者是否被监禁
        bool isJailed = staking.isValidatorJailed(validator);
        console.log("Validator jailed status after punishment:", isJailed);
        
        console.log("Validator Punish Operation completed successfully!");
    }
}