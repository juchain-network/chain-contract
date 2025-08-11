// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {BaseSetup} from "../forge-tests/BaseSetup.t.sol";
import {Validators} from "../contracts/Validators.sol";
import {Proposal} from "../contracts/Proposal.sol";
import {Punish} from "../contracts/Punish.sol";
import {Params} from "../contracts/Params.sol";

// 部署脚本：部署或重新初始化系统合约
// 注意：这假设合约已经在固定地址，主要用于初始化
contract DeploySystemScript is BaseSetup {
    
    event SystemDeployed(address validators, address proposal, address punish);
    event SystemInitialized(address[] validators);
    
    // 初始化已部署的系统合约
    function initializeSystem(address[] memory initialValidators) external {
        require(initialValidators.length > 0, "Need at least one validator");
        
        // 初始化各个合约
        Validators(VAL).initialize(initialValidators);
        Proposal(PRO).initialize(initialValidators);
        Punish(PUN).initialize();
        
        emit SystemInitialized(initialValidators);
    }
    
    // 检查系统状态
    function checkSystemStatus() external view returns (
        bool validatorsInit,
        bool proposalInit, 
        bool punishInit,
        address[] memory activeValidators,
        address[] memory topValidators
    ) {
        // 检查各合约初始化状态
        try Validators(VAL).getActiveValidators() returns (address[] memory active) {
            validatorsInit = active.length > 0;
            activeValidators = active;
        } catch {
            validatorsInit = false;
        }
        
        try Validators(VAL).getTopValidators() returns (address[] memory top) {
            topValidators = top;
        } catch {
            // ignore
        }
        
        try Proposal(PRO).proposalLastingPeriod() returns (uint256 period) {
            proposalInit = period > 0;
        } catch {
            proposalInit = false;
        }
        
        try Proposal(PRO).punishThreshold() returns (uint256 threshold) {
            punishInit = threshold > 0;
        } catch {
            punishInit = false;
        }
    }
    
    // 获取系统配置
    function getSystemConfig() external view returns (
        uint256 proposalLastingPeriod,
        uint256 punishThreshold,
        uint256 removeThreshold,
        uint256 decreaseRate,
        uint256 withdrawProfitPeriod,
        uint256 increasePeriod,
        address receiverAddr
    ) {
        Proposal p = Proposal(PRO);
        return (
            p.proposalLastingPeriod(),
            p.punishThreshold(),
            p.removeThreshold(),
            p.decreaseRate(),
            p.withdrawProfitPeriod(),
            p.increasePeriod(),
            p.receiverAddr()
        );
    }
    
    // 便利函数：检查地址是否为验证者
    function isValidator(address validator) external view returns (bool isActive, bool isTop, bool isPassed) {
        isActive = Validators(VAL).isActiveValidator(validator);
        isTop = Validators(VAL).isTopValidator(validator);
        isPassed = Proposal(PRO).pass(validator);
    }
    
    // 便利函数：获取验证者信息
    function getValidatorDetails(address validator) external view returns (
        address payable feeAddr,
        uint256 status,
        uint256 aacIncoming,
        uint256 totalJailedHB,
        uint256 lastWithdrawProfitsBlock
    ) {
        Validators.Status statusEnum;
        (feeAddr, statusEnum, aacIncoming, totalJailedHB, lastWithdrawProfitsBlock) = 
            Validators(VAL).getValidatorInfo(validator);
        status = uint256(statusEnum);
    }
}
