// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseTestScript} from "./BaseTestScript.s.sol";
import {console} from "forge-std/Test.sol";

// 原子化脚本：仅负责检查提案状态
contract ProposalStatusCheckScript is BaseTestScript {
    function run() public override {
        console.log("Starting Proposal Status Check Operation...");
        
        // 创建一个测试提案用于状态检查
        uint256 newValidatorKey = uint256(keccak256(abi.encodePacked("testValidatorForStatusCheck")));
        address newValidator = fundNewValidator(newValidatorKey);
        
        // 创建提案
        vm.startBroadcast(validatorKeys[0]);
        bytes32 proposalId = proposal.createProposal(newValidator, true, "Add new validator for status check");
        vm.stopBroadcast();
        
        require(proposalId != bytes32(0), "Proposal creation failed");
        console.log("Proposal created with ID:", toHexString(proposalId));
        
        // 检查提案状态
        (bool isPassed, bool isExecuted, uint256 yesVotes, uint256 noVotes, uint256 timestamp) = 
            proposal.getProposalStatus(proposalId);
        
        console.log("\nProposal Status:");
        console.log("  Is Passed:", isPassed);
        console.log("  Is Executed:", isExecuted);
        console.log("  Yes Votes:", yesVotes);
        console.log("  No Votes:", noVotes);
        console.log("  Created At:", timestamp);
        
        // 投票支持提案
        vm.startBroadcast(validatorKeys[0]);
        proposal.voteProposal(proposalId, true);
        vm.stopBroadcast();
        
        // 再次检查提案状态
        (isPassed, isExecuted, yesVotes, noVotes, timestamp) = proposal.getProposalStatus(proposalId);
        
        console.log("\nProposal Status After Voting:");
        console.log("  Is Passed:", isPassed);
        console.log("  Is Executed:", isExecuted);
        console.log("  Yes Votes:", yesVotes);
        console.log("  No Votes:", noVotes);
        
        console.log("\nProposal Status Check Operation completed successfully!");
    }
}