// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {BaseSetup} from "../forge-tests/BaseSetup.t.sol";
import {Proposal} from "../contracts/Proposal.sol";
import {Validators} from "../contracts/Validators.sol";

// 端到端脚本：创建提案 + 多个验证者投票 + 检查结果
// 对应 Hardhat 脚本中的完整流程
contract EndToEndProposalScript is BaseSetup {
    
    event ProposalCreated(bytes32 indexed id, address proposer, address target, bool flag);
    event VoteCast(bytes32 indexed id, address voter, bool vote);
    event ProposalResult(bytes32 indexed id, bool passed, address[] topValidators);
    
    struct ProposalInfo {
        bytes32 id;
        address proposer;
        address target;
        bool flag;
        string details;
    }
    
    function runAddValidatorFlow(
        address newValidator,
        string memory details,
        address[] memory voters
    ) external returns (bool success) {
        return _runProposalFlow(newValidator, true, details, voters);
    }
    
    function runRemoveValidatorFlow(
        address targetValidator, 
        string memory details,
        address[] memory voters
    ) external returns (bool success) {
        return _runProposalFlow(targetValidator, false, details, voters);
    }
    
    function runConfigUpdateFlow(
        uint256 configId,
        uint256 newValue,
        address[] memory voters
    ) external returns (bool success) {
        // 冻结时间戳以确保确定性 ID
        uint256 timestamp = block.timestamp;
        bytes32 id = keccak256(abi.encodePacked(msg.sender, configId, newValue, timestamp));
        
        // 创建配置更新提案
        Proposal(PRO).createUpdateConfigProposal(configId, newValue);
        
        // 验证者投票
        uint256 yesVotes = 0;
        for (uint i = 0; i < voters.length; i++) {
            // 检查是否为活跃验证者
            if (Validators(VAL).isActiveValidator(voters[i])) {
                // 这里简化处理，假设都投赞成票
                // 在实际使用中，可以传入投票选择数组
                vm.prank(voters[i]);
                Proposal(PRO).voteProposal(id, true);
                yesVotes++;
                
                emit VoteCast(id, voters[i], true);
            }
        }
        
        // 检查是否通过 (需要超过半数)
        uint256 requiredVotes = Validators(VAL).getActiveValidators().length / 2 + 1;
        success = yesVotes >= requiredVotes;
        
        address[] memory topValidators = Validators(VAL).getTopValidators();
        emit ProposalResult(id, success, topValidators);
        
        return success;
    }
    
    function _runProposalFlow(
        address target,
        bool isAdd,
        string memory details,
        address[] memory voters
    ) internal returns (bool success) {
        // 冻结时间戳以确保确定性 ID  
        uint256 timestamp = block.timestamp;
        bytes32 id = keccak256(abi.encodePacked(msg.sender, target, isAdd, details, timestamp));
        
        // 创建提案
        Proposal(PRO).createProposal(target, isAdd, details);
        emit ProposalCreated(id, msg.sender, target, isAdd);
        
        // 验证者投票
        uint256 yesVotes = 0;
        for (uint i = 0; i < voters.length; i++) {
            // 检查是否为活跃验证者
            if (Validators(VAL).isActiveValidator(voters[i])) {
                // 这里简化处理，假设都投赞成票
                vm.prank(voters[i]);
                Proposal(PRO).voteProposal(id, true);
                yesVotes++;
                
                emit VoteCast(id, voters[i], true);
            }
        }
        
        // 检查是否通过
        uint256 requiredVotes = Validators(VAL).getActiveValidators().length / 2 + 1;
        success = yesVotes >= requiredVotes;
        
        // 验证最终状态
        if (success) {
            if (isAdd) {
                require(Validators(VAL).isTopValidator(target), "Validator should be added");
                require(Proposal(PRO).pass(target), "Target should be marked as passed");
            } else {
                require(!Validators(VAL).isTopValidator(target), "Validator should be removed");
                require(!Proposal(PRO).pass(target), "Target should be marked as not passed");
            }
        }
        
        address[] memory topValidators = Validators(VAL).getTopValidators();
        emit ProposalResult(id, success, topValidators);
        
        return success;
    }
    
    // 便利函数：获取当前活跃验证者列表用于投票
    function getActiveValidators() external view returns (address[] memory) {
        return Validators(VAL).getActiveValidators();
    }
    
    // 便利函数：获取当前顶级验证者列表
    function getTopValidators() external view returns (address[] memory) {
        return Validators(VAL).getTopValidators();
    }
}
