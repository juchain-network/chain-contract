// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {BaseSetup} from "../test/BaseSetup.t.sol";
import {Proposal} from "../contracts/Proposal.sol";

// 对应 scripts/add-node/create_proposal.js 的增强版本
// 包含地址验证和返回提案 ID
contract CreateProposalScript is BaseSetup {
    
    event ProposalCreated(bytes32 indexed id, address indexed proposer, address indexed dst, bool flag);
    
    function run(address target, bool isAdd, string memory details) external returns (bytes32) {
        // 验证地址格式 (对应 JS 版本的正则检查)
        require(target != address(0), "Invalid target address");
        require(bytes(details).length <= 3000, "Details too long");
        
        // 计算提案 ID (需要与合约逻辑一致)
        bytes32 id = keccak256(abi.encodePacked(msg.sender, target, isAdd, details, block.timestamp));
        
        Proposal(PROPOSAL).createProposal(target, isAdd, details);
        
        emit ProposalCreated(id, msg.sender, target, isAdd);
        return id;
    }
    
    // 便利函数：添加验证者
    function addValidator(address validator, string memory details) external returns (bytes32) {
        return this.run(validator, true, details);
    }
    
    // 便利函数：移除验证者  
    function removeValidator(address validator, string memory details) external returns (bytes32) {
        return this.run(validator, false, details);
    }
}
