// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {BaseSetup} from "./BaseSetup.t.sol";
import {Validators} from "../contracts/Validators.sol";
import {Proposal} from "../contracts/Proposal.sol";
import {Punish} from "../contracts/Punish.sol";

// 完整的奖励分发测试，对应 test/reward.js 
contract RewardFoundryTest is BaseSetup {

    address v1; address v2; address v3;
    address miner;

    function setUp() public {
        v1 = makeAddr("v1");
        v2 = makeAddr("v2"); 
        v3 = makeAddr("v3");
        address[] memory initVals = new address[](3);
        initVals[0] = v1; initVals[1] = v2; initVals[2] = v3;
        deploySystem(initVals);
        miner = v1; // 模拟 coinbase
        vm.coinbase(miner);
    }

    function testRewardEquallyDistributedNoStake() public {
        // 对应 "reward should be equally distributed to active validators if no stake"
        uint256 reward = 1 ether;
        vm.prank(miner);
        Validators(VALIDATORS).distributeBlockReward{value: reward}();

        uint256 remain = reward % 3;

        for (uint i = 0; i < 3; i++) {
            address val = i == 0 ? v1 : (i == 1 ? v2 : v3);
            (, , uint256 aacIncoming,,) = Validators(VALIDATORS).getValidatorInfo(val);
            uint256 inPlan = reward / 3;
            
            if (i == 2) { // last validator gets remainder
                require(aacIncoming == inPlan + remain, "last validator should get remainder");
            } else {
                require(aacIncoming == inPlan, "validator should get equal share");
            }
        }
    }

    function testRemoveValidatorReward() public {
        // 对应 "remove validator's reward"
        uint256 threshold = Proposal(PROPOSAL).punishThreshold();
        
        // 先给 v1 一些奖励
        vm.prank(miner);
        Validators(VALIDATORS).distributeBlockReward{value: 1 ether}();
        
        (, , uint256 toRemoveBefore,,) = Validators(VALIDATORS).getValidatorInfo(v1);
        (, , uint256 v2Before,,) = Validators(VALIDATORS).getValidatorInfo(v2);
        (, , uint256 v3Before,,) = Validators(VALIDATORS).getValidatorInfo(v3);
        
        // 惩罚 v1 达到阈值
        for (uint i = 0; i < threshold; i++) {
            vm.coinbase(miner);
            vm.prank(miner);
            Punish(PUNISH).punish(v1);
            vm.roll(block.number + 1);
        }
        
        // v1 的奖励应该被移除并分配给其他验证者
        (, , uint256 v1After,,) = Validators(VALIDATORS).getValidatorInfo(v1);
        (, , uint256 v2After,,) = Validators(VALIDATORS).getValidatorInfo(v2);
        (, , uint256 v3After,,) = Validators(VALIDATORS).getValidatorInfo(v3);
        
        require(v1After == 0, "punished validator should have no reward");
        
        uint256 inPlan = toRemoveBefore / 2;
        uint256 added = inPlan * 2;
        uint256 remain = toRemoveBefore - added;
        
        require(v2After - v2Before == inPlan, "v2 should get equal share");
        require(v3After - v3Before == inPlan + remain, "v3 should get equal share plus remainder");
    }

    function testJailedValidatorCantGetReward() public {
        // 对应 "jailed validator can't get reward"
        uint256 removeThreshold = Proposal(PROPOSAL).removeThreshold();
        
        // 惩罚 v1 直到被监禁
        for (uint i = 0; i < removeThreshold; i++) {
            vm.coinbase(miner);
            vm.prank(miner);
            Punish(PUNISH).punish(v1);
            vm.roll(block.number + 1);
        }
        
        // 记录其他验证者奖励前状态
        (, , uint256 v2Before,,) = Validators(VALIDATORS).getValidatorInfo(v2);
        (, , uint256 v3Before,,) = Validators(VALIDATORS).getValidatorInfo(v3);
        
        // 分发新奖励
        uint256 reward = 1 ether;
        vm.prank(miner);
        Validators(VALIDATORS).distributeBlockReward{value: reward}();
        
        // 检查监禁的验证者没有获得奖励，只有未监禁的获得
        (, , uint256 v1After,,) = Validators(VALIDATORS).getValidatorInfo(v1);
        (, , uint256 v2After,,) = Validators(VALIDATORS).getValidatorInfo(v2);
        (, , uint256 v3After,,) = Validators(VALIDATORS).getValidatorInfo(v3);
        
        require(v1After == 0, "jailed validator should not get reward");
        
        uint256 inPlan = reward / 2; // 只有 2 个活跃验证者
        uint256 remain = reward - inPlan * 2;
        
        require(v2After - v2Before == inPlan, "v2 should get equal share");
        require(v3After - v3Before == inPlan + remain, "v3 should get equal share plus remainder");
    }

    function testJailedValidatorCantGetPunishProfits() public {
        // 对应 "jailed validator can't get profits of punish"
        uint256 threshold = Proposal(PROPOSAL).punishThreshold();
        uint256 removeThreshold = Proposal(PROPOSAL).removeThreshold();
        
        // 先监禁 v1
        for (uint i = 0; i < removeThreshold; i++) {
            vm.coinbase(miner);
            vm.prank(miner);
            Punish(PUNISH).punish(v1);
            vm.roll(block.number + 1);
        }
        
        // 给 v2 一些奖励
        vm.prank(miner);
        Validators(VALIDATORS).distributeBlockReward{value: 1 ether}();
        
        (, , uint256 v1Before,,) = Validators(VALIDATORS).getValidatorInfo(v1);
        (, , uint256 v2Before,,) = Validators(VALIDATORS).getValidatorInfo(v2);
        (, , uint256 v3Before,,) = Validators(VALIDATORS).getValidatorInfo(v3);
        
        // 惩罚 v2 达到阈值，其奖励应该只分给 v3 (v1 被监禁)
        for (uint i = 0; i < threshold; i++) {
            vm.coinbase(miner);
            vm.prank(miner);
            Punish(PUNISH).punish(v2);
            vm.roll(block.number + 1);
        }
        
        (, , uint256 v1After,,) = Validators(VALIDATORS).getValidatorInfo(v1);
        (, , uint256 v2After,,) = Validators(VALIDATORS).getValidatorInfo(v2);
        (, , uint256 v3After,,) = Validators(VALIDATORS).getValidatorInfo(v3);
        
        require(v1After == v1Before, "jailed validator should not benefit from punish");
        require(v2After == 0, "punished validator should lose reward");
        require(v3After - v3Before == v2Before, "v3 should get all v2's reward");
    }
}
