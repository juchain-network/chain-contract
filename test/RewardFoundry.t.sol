// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {BaseSetup} from "./BaseSetup.t.sol";
import {Validators} from "../contracts/Validators.sol";
import {Proposal} from "../contracts/Proposal.sol";
import {Punish} from "../contracts/Punish.sol";

// 完整的奖励分发测试
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
        // Give miner enough ETH for testing
        vm.deal(miner, 100 ether);
    }

    function testRewardEquallyDistributedNoStake() public {
        // 对应 "reward should go directly to block producer if not jailed"
        // New logic: reward goes directly to the block producer (miner = v1)
        uint256 reward = 1 ether;
        vm.prank(miner);
        Validators(VALIDATORS).distributeBlockReward{value: reward}();

        // 检查出块矿工（v1）获得全部奖励
        (, , uint256 v1Incoming,,) = Validators(VALIDATORS).getValidatorInfo(v1);
        require(v1Incoming == reward, "block producer should get full reward");
        
        // 检查其他验证者没有获得奖励
        (, , uint256 v2Incoming,,) = Validators(VALIDATORS).getValidatorInfo(v2);
        (, , uint256 v3Incoming,,) = Validators(VALIDATORS).getValidatorInfo(v3);
        require(v2Incoming == 0, "v2 should get no reward");
        require(v3Incoming == 0, "v3 should get no reward");
    }

    function testRemoveValidatorReward() public {
        // 对应 "remove validator's reward"
        uint256 threshold = Proposal(PROPOSAL).punishThreshold();
        
        // 先给 v1 一些奖励（v1 是 miner，即出块矿工，获得全部奖励）
        vm.prank(miner);
        Validators(VALIDATORS).distributeBlockReward{value: 1 ether}();
        
        (, , uint256 toRemoveBefore,,) = Validators(VALIDATORS).getValidatorInfo(v1);
        (, , uint256 v2Before,,) = Validators(VALIDATORS).getValidatorInfo(v2);
        (, , uint256 v3Before,,) = Validators(VALIDATORS).getValidatorInfo(v3);
        
        require(toRemoveBefore == 1 ether, "v1 should have received the full reward");
        
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
        // 对应 "jailed block producer's reward is distributed to other validators"
        uint256 removeThreshold = Proposal(PROPOSAL).removeThreshold();
        
        // 惩罚 v1 直到被监禁（v1 是 miner，即出块矿工）
        for (uint i = 0; i < removeThreshold; i++) {
            vm.coinbase(miner);
            vm.prank(miner);
            Punish(PUNISH).punish(v1);
            vm.roll(block.number + 1);
        }
        
        // 记录其他验证者奖励前状态
        (, , uint256 v2Before,,) = Validators(VALIDATORS).getValidatorInfo(v2);
        (, , uint256 v3Before,,) = Validators(VALIDATORS).getValidatorInfo(v3);
        
        // 分发新奖励（v1 被监禁，所以奖励分给其他验证者）
        uint256 reward = 1 ether;
        vm.prank(miner); // miner is still v1, but v1 is now jailed
        Validators(VALIDATORS).distributeBlockReward{value: reward}();
        
        // 检查监禁的出块矿工（v1）没有获得奖励
        (, , uint256 v1After,,) = Validators(VALIDATORS).getValidatorInfo(v1);
        require(v1After == 0, "jailed block producer should not get reward");
        
        // 检查奖励被分给其他活跃验证者（v2 和 v3）
        (, , uint256 v2After,,) = Validators(VALIDATORS).getValidatorInfo(v2);
        (, , uint256 v3After,,) = Validators(VALIDATORS).getValidatorInfo(v3);
        
        uint256 inPlan = reward / 2; // 只有 2 个活跃验证者（v2 和 v3）
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
        
        // 给 v2 一些奖励（v1 被监禁，所以奖励分给其他验证者）
        // 注意：miner 仍然是 v1，但 v1 被监禁了，所以奖励会分给 v2 和 v3
        vm.prank(miner);
        Validators(VALIDATORS).distributeBlockReward{value: 1 ether}();
        
        // 由于 v1 被监禁，奖励应该分给 v2 和 v3
        (, , uint256 v1Before,,) = Validators(VALIDATORS).getValidatorInfo(v1);
        (, , uint256 v2Before,,) = Validators(VALIDATORS).getValidatorInfo(v2);
        (, , uint256 v3Before,,) = Validators(VALIDATORS).getValidatorInfo(v3);
        
        require(v1Before == 0, "jailed v1 should not get reward");
        require(v2Before > 0 || v3Before > 0, "v2 or v3 should have received reward");
        
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
