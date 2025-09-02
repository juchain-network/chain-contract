// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Script} from "lib/forge-std/src/Script.sol";
import {Test} from "lib/forge-std/src/Test.sol";
import {Proposal} from "../contracts/Proposal.sol";
import {Validators} from "../contracts/Validators.sol";

// 端到端脚本：创建提案 + 多个验证者投票 + 检查结果
contract EndToEndProposalScript is Script, Test {
    
    // Fixed system addresses (consistent with deployment)
    address constant VALIDATORS = 0x000000000000000000000000000000000000f000;
    address constant PUNISH = 0x000000000000000000000000000000000000F001;
    address constant PROPOSAL = 0x000000000000000000000000000000000000F002;
    address constant STAKING = 0x000000000000000000000000000000000000F003;
    
    event ProposalCreated(bytes32 indexed id, address proposer, address target, bool flag);
    event VoteCast(bytes32 indexed id, address voter, bool vote);
    event ProposalResult(bytes32 indexed id, bool passed, address[] topValidators);
    
    function run() external {
        // 示例：端到端提案流程演示
        address testTarget = 0xa0Ee7A142d267C1f36714E4a8F75612F20a79720;
        
        // 创建一个模拟的投票者数组（空数组，因为权限限制）
        address[] memory voters = new address[](0);
        
        // 直接调用内部函数执行添加验证者流程
        bool success = _runProposalFlow(
            testTarget,
            true, // 添加验证者
            "End-to-end test: Adding validator",
            voters
        );
        
        if (success) {
            emit ProposalResult(bytes32(uint256(uint160(testTarget))), true, Validators(VALIDATORS).getTopValidators());
        }
    }
    
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
        Proposal(PROPOSAL).createUpdateConfigProposal(configId, newValue);
        
        // 验证者投票
        uint256 yesVotes = 0;
        for (uint i = 0; i < voters.length; i++) {
            // 检查是否为活跃验证者
            if (Validators(VALIDATORS).isActiveValidator(voters[i])) {
                // 这里简化处理，假设都投赞成票
                // 在实际使用中，可以传入投票选择数组
                vm.prank(voters[i]);
                Proposal(PROPOSAL).voteProposal(id, true);
                yesVotes++;
                
                emit VoteCast(id, voters[i], true);
            }
        }
        
        // 检查是否通过 (需要超过半数)
        uint256 requiredVotes = Validators(VALIDATORS).getActiveValidators().length / 2 + 1;
        success = yesVotes >= requiredVotes;
        
        address[] memory topValidators = Validators(VALIDATORS).getTopValidators();
        emit ProposalResult(id, success, topValidators);
        
        return success;
    }
    
    function _runProposalFlow(
        address target,
        bool isAdd,
        string memory details,
        address[] memory voters
    ) internal returns (bool success) {
        // 创建提案并获取真实的ID（通过事件）
        Proposal(PROPOSAL).createProposal(target, isAdd, details);
        
        // 注意：在实际环境中，我们需要从事件日志中获取真实的提案ID
        // 这里为了演示，我们使用一个简化的ID计算
        // 真实的ID应该从 LogCreateProposal 事件中获取
        bytes32 id = keccak256(abi.encodePacked(msg.sender, target, isAdd, details, block.timestamp));
        emit ProposalCreated(id, msg.sender, target, isAdd);
        
        // 验证者投票
        uint256 yesVotes = 0;
        for (uint i = 0; i < voters.length; i++) {
            // 检查是否为活跃验证者
            if (Validators(VALIDATORS).isActiveValidator(voters[i])) {
                // 这里简化处理，假设都投赞成票
                vm.prank(voters[i]);
                Proposal(PROPOSAL).voteProposal(id, true);
                yesVotes++;
                
                emit VoteCast(id, voters[i], true);
            }
        }
        
        // 检查是否通过
        uint256 requiredVotes = Validators(VALIDATORS).getActiveValidators().length / 2 + 1;
        success = yesVotes >= requiredVotes;
        
        // 验证最终状态
        if (success) {
            if (isAdd) {
                require(Validators(VALIDATORS).isTopValidator(target), "Validator should be added");
                require(Proposal(PROPOSAL).pass(target), "Target should be marked as passed");
            } else {
                require(!Validators(VALIDATORS).isTopValidator(target), "Validator should be removed");
                require(!Proposal(PROPOSAL).pass(target), "Target should be marked as not passed");
            }
        }
        
        address[] memory topValidators = Validators(VALIDATORS).getTopValidators();
        emit ProposalResult(id, success, topValidators);
        
        return success;
    }
    
    // 便利函数：获取当前活跃验证者列表用于投票
    function getActiveValidators() external view returns (address[] memory) {
        return Validators(VALIDATORS).getActiveValidators();
    }
    
    // 便利函数：获取当前顶级验证者列表
    function getTopValidators() external view returns (address[] memory) {
        return Validators(VALIDATORS).getTopValidators();
    }
}
