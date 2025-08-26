// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {BaseSetup} from "../test/BaseSetup.t.sol";
import {Staking} from "../contracts/Staking.sol";
import {Validators} from "../contracts/Validators.sol";

// Staking 系统操作脚本
contract StakingOperationsScript is BaseSetup {
    
    event StakingInfo(string info, uint256 value);
    event ValidatorInfo(string info, address validator, uint256 stake);
    
    function run() external {
        emit StakingInfo("=== Staking System Operations ===", 0);
        
        // 展示当前验证者状态
        showValidatorStatus();
        
        // 演示验证者注册质押
        demonstrateValidatorRegistration();
        
        // 演示委托功能
        demonstrateDelegation();
        
        // 演示获取顶级验证者
        demonstrateTopValidators();
        
        emit StakingInfo("=== Operations Complete ===", 0);
    }
    
    function showValidatorStatus() internal {
        emit StakingInfo("--- Current Validator Status ---", 0);
        
        Staking stakingContract = Staking(STAKING);
        Validators validatorsContract = Validators(VALIDATORS);
        
        // 安全地检查当前验证者
        address[] memory currentValidators = validatorsContract.getActiveValidators();
        emit StakingInfo("Current validators count", currentValidators.length);
        
        // 仅在有验证者时才检查Staking状态
        if (currentValidators.length > 0) {
            for (uint i = 0; i < currentValidators.length && i < 5; i++) {
                address validator = currentValidators[i];
                
                // 安全地获取验证者信息
                try stakingContract.getValidatorInfo(validator) returns (
                    uint256 selfStake,
                    uint256 totalDelegated,
                    uint256, /* commissionRate */
                    bool, /* isJailed */
                    uint256 /* jailUntilBlock */
                ) {
                    emit ValidatorInfo("Validator", validator, selfStake + totalDelegated);
                } catch {
                    emit ValidatorInfo("Validator (no stake info)", validator, 0);
                }
            }
        } else {
            emit StakingInfo("No validators registered yet", 0);
        }
        
        // 安全地检查Staking系统状态
        try stakingContract.MIN_VALIDATORS() returns (uint256 minValidators) {
            emit StakingInfo("Minimum validators required", minValidators);
        } catch {
            emit StakingInfo("Cannot get minimum validators", 0);
        }
    }
    
    function demonstrateValidatorRegistration() internal {
        emit StakingInfo("--- Validator Registration Demo ---", 0);
        
        Staking stakingContract = Staking(STAKING);
        
        // 安全地获取最小质押要求
        try stakingContract.MIN_VALIDATOR_STAKE() returns (uint256 minStake) {
            emit StakingInfo("Minimum stake required", minStake);
        } catch {
            emit StakingInfo("Cannot get minimum stake", 0);
        }
        
        // 这里只是展示接口，实际注册需要发送ETH
        emit StakingInfo("To register call: staking.registerValidator{value: minStake}(commissionRate)", 0);
    }
    
    function demonstrateDelegation() internal {
        emit StakingInfo("--- Delegation Demo ---", 0);
        
        Staking stakingContract = Staking(STAKING);
        
        // 安全地获取顶级验证者
        try stakingContract.getTopValidators(5) returns (address[] memory topValidators) {
            if (topValidators.length > 0) {
                address validator = topValidators[0];
                emit ValidatorInfo("Example delegation to validator", validator, 0);
                
                // 安全地显示委托信息
                try stakingContract.getDelegationInfo(msg.sender, validator) returns (
                    uint256 delegatedAmount,
                    uint256 rewards,
                    uint256 unbondingAmount,
                    uint256 unbondingBlock
                ) {
                    emit StakingInfo("Delegated amount", delegatedAmount);
                    emit StakingInfo("Pending rewards", rewards);
                    emit StakingInfo("Unbonding amount", unbondingAmount);
                    emit StakingInfo("Unbonding block", unbondingBlock);
                } catch {
                    emit StakingInfo("Cannot get delegation info", 0);
                }
            } else {
                emit StakingInfo("No validators available for delegation", 0);
            }
        } catch {
            emit StakingInfo("Cannot get top validators", 0);
        }
    }
    
    function demonstrateTopValidators() internal {
        emit StakingInfo("--- Top Validators ---", 0);
        
        Staking stakingContract = Staking(STAKING);
        
        // 安全地获取顶级验证者
        try stakingContract.getTopValidators(5) returns (address[] memory topValidators) {
            emit StakingInfo("Total top validators", topValidators.length);
            
            for (uint i = 0; i < topValidators.length && i < 3; i++) {
                try stakingContract.getValidatorInfo(topValidators[i]) returns (
                    uint256 selfStake,
                    uint256 totalDelegated,
                    uint256 commissionRate,
                    bool isJailed,
                    uint256 jailUntilBlock
                ) {
                    emit ValidatorInfo("Validator info", topValidators[i], selfStake);
                    emit StakingInfo("Total delegated", totalDelegated);
                    emit StakingInfo("Commission rate %", commissionRate);
                    emit StakingInfo("Jailed", isJailed ? 1 : 0);
                    emit StakingInfo("Jail until block", jailUntilBlock);
                } catch {
                    emit ValidatorInfo("Validator (info unavailable)", topValidators[i], 0);
                }
            }
        } catch {
            emit StakingInfo("Cannot get top validators list", 0);
        }
    }
}
