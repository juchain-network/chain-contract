// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {Script, console} from "lib/forge-std/src/Script.sol";
import {Staking} from "../contracts/Staking.sol";

/**
 * @title ValidatorStakeScript
 * @dev 验证者质押脚本 - 每个验证者独立运行
 */
contract ValidatorStakeScript is Script {
    
    function setUp() public {}

    /**
     * @dev 验证者质押函数
     * 每个验证者使用自己的私钥运行此脚本
     */
    function run() external {
        // 从环境变量获取验证者私钥
        uint256 validatorPrivateKey = vm.envOr("VALIDATOR_PRIVATE_KEY", uint256(0));
        require(validatorPrivateKey != 0, "VALIDATOR_PRIVATE_KEY not set");
        
        // 从环境变量获取Staking合约地址
        address stakingAddress = vm.envOr("STAKING_CONTRACT", address(0));
        require(stakingAddress != address(0), "STAKING_CONTRACT address not set");
        
        address validator = vm.addr(validatorPrivateKey);
        uint256 stakeAmount = 10000 ether; // 10,000 JU
        uint256 commissionRate = vm.envOr("COMMISSION_RATE", uint256(500)); // 默认5%
        
        console.log("=== Validator Staking ===");
        console.log("Validator address:", validator);
        console.log("Staking contract:", stakingAddress);
        console.log("Stake amount:", stakeAmount / 1 ether, "JU");
        console.log("Commission rate:", commissionRate, "/10000");
        
        // 检查余额
        require(validator.balance >= stakeAmount, "Insufficient balance for staking");
        
        vm.startBroadcast(validatorPrivateKey);
        
        // 注册并质押
        Staking(stakingAddress).registerValidator{value: stakeAmount}(commissionRate);
        
        vm.stopBroadcast();
        
        console.log("Successfully registered as validator!");
        console.log("Staked:", stakeAmount / 1 ether, "JU");
        
        // 验证注册状态
        checkRegistrationStatus(stakingAddress, validator);
    }
    
    /**
     * @dev 检查验证者注册状态
     */
    function checkRegistrationStatus(address stakingAddress, address validator) internal view {
        Staking staking = Staking(stakingAddress);
        
        // 检查总质押数量
        uint256 totalStaked = staking.totalStaked();
        console.log("Total network stake:", totalStaked / 1 ether, "JU");
        
        // 检查验证者数量
        uint256 validatorCount = staking.getValidatorCount();
        console.log("Total validators:", validatorCount);
        
        console.log("Validator", validator, "is now active in the staking contract");
    }
}
