// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseTestScript} from "./BaseTestScript.s.sol";
import {console} from "forge-std/Test.sol";

// 原子化脚本：仅负责执行提案操作
contract ProposalExecuteScript is BaseTestScript {
    function run() public override {
        console.log("Starting Proposal Execute Operation...");
        
        // 创建新提案并执行完整流程（创建+投票+执行）
        // 创建新验证者
        uint256 newValidatorKey = uint256(keccak256(abi.encodePacked("testValidatorForExecuteProposal")));
        address newValidator = fundNewValidator(newValidatorKey);
        
        console.log("New validator address for execute proposal:", newValidator);
        
        // 1. 创建提案
        vm.startBroadcast(validatorKeys[0]);
        bytes32 proposalId = proposal.createProposal(newValidator, true, "Add new validator for execute test");
        vm.stopBroadcast();
        
        require(proposalId != bytes32(0), "Proposal creation failed");
        console.log("Proposal created with ID:", toHexString(proposalId));
        
        // 2. 所有验证者投票支持提案
        for (uint256 i = 0; i < validatorAccounts.length; i++) {
            vm.startBroadcast(validatorKeys[i]);
            proposal.voteProposal(proposalId, true);
            vm.stopBroadcast();
            console.log("Validator", i+1, "voted for proposal");
        }
        
        // 3. 执行提案
        vm.startBroadcast(validatorKeys[0]);
        proposal.executeProposal(proposalId);
        vm.stopBroadcast();
        
        console.log("Proposal Execute Operation completed successfully!");
    }
}