// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {BaseSetup} from "../test/BaseSetup.t.sol";
import {Proposal} from "../contracts/Proposal.sol";
import {Validators} from "../contracts/Validators.sol";

/**
 * @title ProposalWorkflow
 * @dev 提案工作流程脚本，演示提案创建和投票流程
 */
contract ProposalWorkflow is BaseSetup {
    
    event WorkflowEvent(string message, address addr, bool result);
    
    function run() external {
        // 演示提案创建
        address newValidator = 0xa0Ee7A142d267C1f36714E4a8F75612F20a79720;
        
        // 1. 检查初始状态
        bool isValidatorBefore = Validators(VALIDATORS).isTopValidator(newValidator);
        emit WorkflowEvent("Initial validator status", newValidator, isValidatorBefore);
        
        // 2. 创建添加验证者的提案
        bool createResult = Proposal(PROPOSAL).createProposal(newValidator, true, "Workflow: Adding new validator");
        emit WorkflowEvent("Proposal creation", newValidator, createResult);
        
        // 3. 检查提案者是否为验证者
        bool isProposerValidator = Validators(VALIDATORS).isActiveValidator(msg.sender);
        emit WorkflowEvent("Proposer is validator", msg.sender, isProposerValidator);
        
        // 4. 检查系统状态
        address[] memory activeValidators = Validators(VALIDATORS).getActiveValidators();
        emit WorkflowEvent("Active validators count", address(uint160(activeValidators.length)), true);
        
        // 5. 检查提案通过状态
        bool passStatus = Proposal(PROPOSAL).pass(newValidator);
        emit WorkflowEvent("Target validator pass status", newValidator, passStatus);
    }
    
    function runConfigWorkflow() external {
        // 演示配置更新提案
        uint256 configId = 2; // 示例配置ID
        uint256 newValue = 5000; // 新的配置值
        
        // 创建配置更新提案
        bool createResult = Proposal(PROPOSAL).createUpdateConfigProposal(configId, newValue);
        emit WorkflowEvent("Config proposal creation", address(uint160(configId)), createResult);
        
        // 检查当前配置
        uint256 currentLastingPeriod = Proposal(PROPOSAL).proposalLastingPeriod();
        emit WorkflowEvent("Current lasting period", address(uint160(currentLastingPeriod)), true);
    }
}
