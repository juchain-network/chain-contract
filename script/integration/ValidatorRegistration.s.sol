// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseTestScript} from "./BaseTestScript.s.sol";
import {console} from "forge-std/Test.sol";

// 原子化脚本：仅负责验证者注册操作
contract ValidatorRegistrationScript is BaseTestScript {
    function run() public override {
        console.log("Starting Validator Registration Operation...");
        
        // 创建新验证者账户和私钥
        uint256 newValidatorKey = uint256(keccak256(abi.encodePacked("testValidator1")));
        address newValidator = fundNewValidator(newValidatorKey); // 使用基类方法为新验证者提供资金
        
        console.log("New validator address:", newValidator);
        
        // 创建并通过添加新验证者的提案
        vm.startBroadcast(validatorKeys[0]);
        bytes32 proposalId = proposal.createProposal(newValidator, true, "Add new validator");
        vm.stopBroadcast();
        require(proposalId != bytes32(0), "Proposal creation failed");
        
        console.log("Proposal created with ID:", toHexString(proposalId));
        
        // 所有验证者投票支持提案
        for (uint256 i = 0; i < validatorAccounts.length; i++) {
            vm.startBroadcast(validatorKeys[i]);
            proposal.voteProposal(proposalId, true);
            vm.stopBroadcast();
        }
        
        // 使用新验证者的私钥注册验证者
        console.log("Registering validator:", newValidator);
        vm.startBroadcast(newValidatorKey);
        staking.registerValidator{value: INITIAL_STAKE}(1000); // 10% 佣金率
        vm.stopBroadcast();
        
        // 验证注册状态
        (uint256 selfStake, , , , , , , ) = staking.getValidatorInfo(newValidator);
        require(selfStake == INITIAL_STAKE, "Validator should have correct self-stake");
        
        console.log("Validator Registration Operation completed successfully!");
    }
}