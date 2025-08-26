// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {BaseSetup} from "../test/BaseSetup.t.sol";
import {Proposal} from "../contracts/Proposal.sol";
import {Validators} from "../contracts/Validators.sol";

contract VoteProposalScript is BaseSetup {
    function run() external {
        // 示例：对第一个提案投赞成票
        // 注意：只有活跃验证者才能投票，此脚本仅作演示
        bytes32 sampleProposalId = 0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef;
        
        // 检查当前账户是否为验证者
        bool isValidator = Validators(VALIDATORS).isActiveValidator(msg.sender);
        if (!isValidator) {
            // 仅作演示，实际使用需要验证者账户
            return;
        }
        
        voteOnProposal(sampleProposalId, true);
    }
    
    function voteOnProposal(bytes32 proposalId, bool vote) public {
        // 验证提案 ID 格式 (在 JS 版本中有正则验证)
        require(proposalId != bytes32(0), "Invalid proposal id");
        
        Proposal(PROPOSAL).voteProposal(proposalId, vote);
    }
    
    // 便利函数：投赞成票
    function voteYes(bytes32 proposalId) external {
        Proposal(PROPOSAL).voteProposal(proposalId, true);
    }
    
    // 便利函数：投反对票
    function voteNo(bytes32 proposalId) external {
        Proposal(PROPOSAL).voteProposal(proposalId, false);
    }
}
