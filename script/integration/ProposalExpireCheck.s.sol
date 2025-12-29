// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseTestScript} from "./BaseTestScript.s.sol";
import {console} from "forge-std/Test.sol";

// 原子化脚本：仅负责检查提案失效状态
contract ProposalExpireCheckScript is BaseTestScript {
    function run() public override {
        console.log("Starting Proposal Expire Check Operation...");
        
        // 创建一个测试提案用于失效检查
        uint256 newValidatorKey = uint256(keccak256(abi.encodePacked("testValidatorForExpireCheck")));
        address newValidator = fundNewValidator(newValidatorKey);
        
        // 创建提案
        vm.startBroadcast(validatorKeys[0]);
        bytes32 proposalId = proposal.createProposal(newValidator, true, "Add new validator for expire check");
        vm.stopBroadcast();
        
        require(proposalId != bytes32(0), "Proposal creation failed");
        console.log("Proposal created with ID:", toHexString(proposalId));
        
        // 检查提案是否存在
        bool exists = proposal.proposalExists(proposalId);
        console.log("Proposal exists:", exists);
        require(exists, "Proposal should exist");
        
        // 获取提案持续时间
        uint256 proposalLastingPeriod = proposal.getProposalLastingPeriod();
        console.log("Proposal lasting period:", proposalLastingPeriod, "seconds");
        
        // 注意：提案失效检查需要时间流逝，这部分将在Shell脚本中完成
        // 这里只创建提案并验证其存在
        
        console.log("\nProposal Expire Check Operation completed successfully!");
        console.log("Note: Proposal expiration time simulation will be handled in shell script using cast rpc");
    }
}