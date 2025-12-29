// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseTestScript} from "./BaseTestScript.s.sol";
import {console} from "forge-std/Test.sol";

// 原子化脚本：仅负责检查委托者状态
contract DelegatorStatusCheckScript is BaseTestScript {
    function run() public override {
        console.log("Starting Delegator Status Check Operation...");
        
        // 创建委托者账户（与Delegate脚本中使用相同的密钥）
        uint256 delegatorKey = uint256(keccak256(abi.encodePacked("testDelegator1")));
        address delegatorAddr = vm.addr(delegatorKey);
        
        console.log("Checking delegator:", delegatorAddr);
        
        // 检查委托者对所有验证者的委托状态
        for (uint256 i = 0; i < validatorAccounts.length; i++) {
            address validatorAddr = validatorAccounts[i];
            
            (uint256 delegatorStake, uint256 rewards, uint256 withdrawableAmount, uint256 unbondingTimestamp) = 
                staking.getDelegationInfo(delegatorAddr, validatorAddr);
            
            if (delegatorStake > 0 || rewards > 0 || withdrawableAmount > 0) {
                console.log("\n  Validator:", validatorAddr);
                console.log("    Delegated Stake:", delegatorStake / 1 ether, "ETH");
                console.log("    Rewards:", rewards / 1 ether, "ETH");
                console.log("    Withdrawable Amount:", withdrawableAmount / 1 ether, "ETH");
                console.log("    Unbonding Timestamp:", unbondingTimestamp);
            }
        }
        
        console.log("\nDelegator Status Check Operation completed successfully!");
    }
}