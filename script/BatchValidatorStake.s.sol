// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {Script, console} from "lib/forge-std/src/Script.sol";
import {Staking} from "../contracts/Staking.sol";

/**
 * @title BatchValidatorStakeScript
 * @dev 批量验证者质押脚本 - 适用于控制多个验证者私钥的情况
 * 警告：仅在您拥有所有验证者私钥的情况下使用
 */
contract BatchValidatorStakeScript is Script {
    
    function setUp() public {}

    /**
     * @dev 批量验证者质押 - 需要预先设置验证者私钥
     */
    function run() external {
        address stakingAddress = vm.envOr("STAKING_CONTRACT", address(0));
        require(stakingAddress != address(0), "STAKING_CONTRACT address not set");
        
        // 验证者私钥数组 - 在实际部署时从环境变量或安全存储中获取
        uint256[] memory validatorKeys = getValidatorPrivateKeys();
        
        uint256 stakeAmount = 10000 ether; // 10,000 JU
        uint256 commissionRate = 500; // 5%
        
        console.log("=== Batch Validator Staking ===");
        console.log("Staking contract:", stakingAddress);
        console.log("Number of validators:", validatorKeys.length);
        console.log("Stake per validator:", stakeAmount / 1 ether, "JU");
        
        for (uint256 i = 0; i < validatorKeys.length; i++) {
            registerSingleValidator(
                stakingAddress,
                validatorKeys[i],
                stakeAmount,
                commissionRate,
                i + 1
            );
        }
        
        // 最终状态检查
        checkFinalStatus(stakingAddress);
    }
    
    /**
     * @dev 注册单个验证者
     */
    function registerSingleValidator(
        address stakingAddress,
        uint256 privateKey,
        uint256 stakeAmount,
        uint256 commissionRate,
        uint256 index
    ) internal {
        address validator = vm.addr(privateKey);
        
        console.log("Registering validator", index, ":", validator);
        
        // 检查余额（如果余额不足，需要先转账）
        if (validator.balance < stakeAmount) {
            console.log("Warning: Validator", validator, "has insufficient balance");
            console.log("Required:", stakeAmount / 1 ether, "JU");
            console.log("Current:", validator.balance / 1 ether, "JU");
            return; // 跳过此验证者
        }
        
        vm.startBroadcast(privateKey);
        
        Staking(stakingAddress).registerValidator{value: stakeAmount}(commissionRate);
        
        vm.stopBroadcast();
        
        console.log("Validator", index, "registered successfully");
    }
    
    /**
     * @dev 获取验证者私钥 - 实际部署时需要安全的密钥管理
     */
    function getValidatorPrivateKeys() internal view returns (uint256[] memory) {
        uint256[] memory keys = new uint256[](5);
        
        // 警告：这些是示例私钥，实际部署时必须使用真实的安全密钥
        // 应该从环境变量、密钥管理服务或硬件钱包获取
        if (block.chainid == 202599) {
            // 测试网示例私钥 - 仅用于演示
            keys[0] = vm.envOr("VALIDATOR_1_KEY", uint256(0));
            keys[1] = vm.envOr("VALIDATOR_2_KEY", uint256(0));
            keys[2] = vm.envOr("VALIDATOR_3_KEY", uint256(0));
            keys[3] = vm.envOr("VALIDATOR_4_KEY", uint256(0));
            keys[4] = vm.envOr("VALIDATOR_5_KEY", uint256(0));
        }
        
        // 验证所有私钥都已设置
        for (uint i = 0; i < keys.length; i++) {
            require(keys[i] != 0, string(abi.encodePacked("VALIDATOR_", i + 1, "_KEY not set")));
        }
        
        return keys;
    }
    
    /**
     * @dev 检查最终状态
     */
    function checkFinalStatus(address stakingAddress) internal view {
        Staking staking = Staking(stakingAddress);
        
        uint256 totalStaked = staking.totalStaked();
        uint256 validatorCount = staking.getValidatorCount();
        
        console.log("=== Final Status ===");
        console.log("Total staked:", totalStaked / 1 ether, "JU");
        console.log("Total validators:", validatorCount);
        console.log("Average stake per validator:", (totalStaked / validatorCount) / 1 ether, "JU");
    }
}
