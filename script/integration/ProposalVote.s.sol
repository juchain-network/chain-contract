// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseTestScript} from "./BaseTestScript.s.sol";
import {console} from "forge-std/Test.sol";

// 原子化脚本：仅负责提案投票操作
contract ProposalVoteScript is BaseTestScript {
    function run() public override {
        console.log("Starting Proposal Vote Operation...");
        
        // 获取最新创建的提案ID（假设是最近创建的）
        // 注意：在实际测试中，应该从事件日志中获取提案ID
        // 这里我们创建一个新提案并立即投票
        
        // 创建新提案
        uint256 newValidatorKey = uint256(keccak256(abi.encodePacked("testValidatorForVoteProposal")));
        address newValidator = fundNewValidator(newValidatorKey);
        
        vm.startBroadcast(validatorKeys[0]);
        bytes32 proposalId = proposal.createProposal(newValidator, true, "Add new validator for voting test");
        vm.stopBroadcast();
        
        require(proposalId != bytes32(0), "Proposal creation failed");
        console.log("Proposal created with ID:", toHexString(proposalId));
        
        // 所有验证者投票支持提案
        for (uint256 i = 0; i < validatorAccounts.length; i++) {
            vm.startBroadcast(validatorKeys[i]);
            proposal.voteProposal(proposalId, true);
            vm.stopBroadcast();
            console.log("Validator", i+1, "voted for proposal");
        }
        
        console.log("Proposal Vote Operation completed successfully!");
    }
}