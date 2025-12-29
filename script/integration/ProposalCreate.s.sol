// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseTestScript} from "./BaseTestScript.s.sol";
import {console} from "forge-std/Test.sol";

// 原子化脚本：仅负责创建提案操作
contract ProposalCreateScript is BaseTestScript {
    function run() public override {
        console.log("Starting Proposal Create Operation...");
        
        // 创建新验证者账户（用于提案测试）
        uint256 newValidatorKey = uint256(keccak256(abi.encodePacked("testValidatorForProposal")));
        address newValidator = fundNewValidator(newValidatorKey);
        
        console.log("New validator address for proposal:", newValidator);
        
        // 使用第一个验证者的密钥创建添加新验证者的提案
        vm.startBroadcast(validatorKeys[0]);
        bytes32 proposalId = proposal.createProposal(newValidator, true, "Add new validator");
        vm.stopBroadcast();
        
        require(proposalId != bytes32(0), "Proposal creation failed");
        
        console.log("Proposal created with ID:", toHexString(proposalId));
        
        console.log("Proposal Create Operation completed successfully!");
    }
}