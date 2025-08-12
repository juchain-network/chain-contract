// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {BaseSetup} from "../test/BaseSetup.t.sol";
import {Proposal} from "../contracts/Proposal.sol";

// 对应 scripts/add-node/start_vote.js 的功能
contract VoteProposalScript is BaseSetup {
    function run(bytes32 proposalId, bool vote) external {
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
