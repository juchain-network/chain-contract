// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {BaseSetup} from "./BaseSetup.t.sol";
import {Punish} from "../contracts/Punish.sol";
import {Validators} from "../contracts/Validators.sol";
import {Proposal} from "../contracts/Proposal.sol";

// 补充缺失的 Punish 合约测试用例
contract PunishMissingFoundryTest is BaseSetup {

    address v1; address v2; address v3;

    function setUp() public {
        v1 = makeAddr("v1");
        v2 = makeAddr("v2");
        v3 = makeAddr("v3");
        address[] memory initVals = new address[](3);
        initVals[0] = v1; initVals[1] = v2; initVals[2] = v3;
        deploySystem(initVals);
    }

    // 辅助函数：检查验证者是否被监禁
    function isJailed(address validator) internal view returns (bool) {
        (, Validators.Status status, , , ) = Validators(VAL).getValidatorInfo(validator);
        return status == Validators.Status.Jailed;
    }

    function testPunishInitialization() public view {
        // 对应 "test punish contract deployment and basic setup"
        Punish punish = Punish(PUN);
        
        // 验证初始状态
        require(punish.initialized(), "contract should be initialized");
        require(punish.getPunishValidatorsLen() == 0, "should have no punish validators initially");
        
        // 验证初始惩罚记录为空
        require(punish.getPunishRecord(v1) == 0, "v1 should have no punish record");
        require(punish.getPunishRecord(v2) == 0, "v2 should have no punish record");
        require(punish.getPunishRecord(v3) == 0, "v3 should have no punish record");
    }

    function testJailedValidatorReactivation() public {
        // 对应 "jailed record will be cleaned if validator repass proposal"
        Punish punish = Punish(PUN);
        
        // 首先惩罚验证者直到被监禁
        vm.coinbase(VAL); // 设置 coinbase 为 VAL 合约地址
        vm.startPrank(VAL);
        for (uint256 i = 0; i < 48; i++) {
            vm.roll(block.number + 1); // 移动到下一个区块
            punish.punish(v1);
        }
        vm.stopPrank();
        
        // 验证者应该被监禁
        require(isJailed(v1), "v1 should be jailed");
        require(punish.getPunishRecord(v1) == 0, "v1 punish record should be reset after removal");
        
        // 创建重新激活验证者的提案
        vm.warp(5_000_000);
        bytes32 id = keccak256(abi.encodePacked(address(this), v1, true, "", block.timestamp));
        vm.prank(address(this));
        require(true == Proposal(PRO).createProposal(v1, true, ""), "should create proposal");
        
        // 投票通过提案
        vm.prank(v2); Proposal(PRO).voteProposal(id, true);
        vm.prank(v3); Proposal(PRO).voteProposal(id, true);
        
        // 验证者应该不再被监禁且惩罚记录被清除
        require(!isJailed(v1), "v1 should not be jailed");
        require(punish.getPunishRecord(v1) == 0, "v1 punish record should remain cleaned");
    }

    function testComplexPunishWorkflow() public {
        // 对应复杂的惩罚工作流测试
        Punish punish = Punish(PUN);
        
        // 测试多个验证者的惩罚流程
        vm.coinbase(VAL);
        vm.startPrank(VAL);
        
        // v1 惩罚到移除收入阈值
        for (uint256 i = 0; i < 24; i++) {
            vm.roll(block.number + 1);
            punish.punish(v1);
        }
        
        // v2 惩罚直到被移除
        for (uint256 i = 0; i < 48; i++) {
            vm.roll(block.number + 1);
            punish.punish(v2);
        }
        
        vm.stopPrank();
        
        // 验证状态
        require(!isJailed(v1), "v1 should not be jailed yet but income removed");
        require(isJailed(v2), "v2 should be jailed after removal");
        require(punish.getPunishRecord(v1) == 24, "v1 should have 24 punish records");
        require(punish.getPunishRecord(v2) == 0, "v2 punish record should be reset after removal");
        
        // 继续惩罚 v1 直到被移除
        vm.coinbase(VAL);
        vm.startPrank(VAL);
        for (uint256 i = 0; i < 24; i++) {
            vm.roll(block.number + 1);
            punish.punish(v1);
        }
        vm.stopPrank();

        require(isJailed(v1), "v1 should now be jailed");
        require(punish.getPunishRecord(v1) == 0, "v1 punish record should be reset after removal");        // 现在两个验证者都被监禁，只有 v3 可以参与投票
        // 需要至少2票才能通过，但只有1个活跃验证者，所以提案无法通过
        vm.warp(6_000_000);
        bytes32 id = keccak256(abi.encodePacked(address(this), v1, true, "", block.timestamp));
        Proposal(PRO).createProposal(v1, true, "");
        
        vm.prank(v3); Proposal(PRO).voteProposal(id, true);
        
        // 提案不应该通过，因为只有1票，需要至少2票
        require(!Proposal(PRO).pass(v1), "proposal should not pass with only 1 vote");
        require(isJailed(v1), "v1 should still be jailed");
    }

    function testPunishPermission() public {
        // 测试只有 Validators 合约可以调用 punish
        Punish punish = Punish(PUN);
        
        // 随机地址调用应该失败
        vm.prank(makeAddr("random"));
        (bool success, ) = address(punish).call(abi.encodeWithSelector(punish.punish.selector, v1));
        require(!success, "should fail when called by non-validator contract");
        
        // 只有 VAL 合约可以调用
        vm.coinbase(VAL);
        vm.prank(VAL);
        punish.punish(v1);
        require(punish.getPunishRecord(v1) == 1, "punish should succeed from VAL contract");
    }

    function testPunishRecordCleaning() public {
        // 测试惩罚记录清理机制
        Punish punish = Punish(PUN);
        
        // 先惩罚但不到移除收入阈值
        vm.coinbase(VAL);
        vm.startPrank(VAL);
        for (uint256 i = 0; i < 10; i++) {
            vm.roll(block.number + 1);
            punish.punish(v1);
        }
        vm.stopPrank();

        require(punish.getPunishRecord(v1) == 10, "should have 10 punish records");
        require(!isJailed(v1), "should not be jailed");
        
        // 直接通过 Validators 合约测试清理记录
        // 这模拟了当验证者重新激活时的场景
        vm.prank(PRO);
        bool success = Validators(VAL).tryActive(v1);
        require(success, "tryActive should succeed");
        
        // 惩罚记录应该保持不变，因为验证者没有被监禁
        require(punish.getPunishRecord(v1) == 10, "punish record should remain unchanged for active validator");
        
        // 现在监禁验证者然后重新激活
        vm.coinbase(VAL);
        vm.startPrank(VAL);
        for (uint256 i = 0; i < 38; i++) { // 总共48次达到移除阈值
            vm.roll(block.number + 1);
            punish.punish(v1);
        }
        vm.stopPrank();
        
        require(isJailed(v1), "v1 should be jailed");
        require(punish.getPunishRecord(v1) == 0, "punish record should be reset after removal");
        
        // 重新激活时记录已经被清理了
        vm.prank(PRO);
        success = Validators(VAL).tryActive(v1);
        require(success, "tryActive should succeed");
        require(punish.getPunishRecord(v1) == 0, "punish record should remain clean");
    }
}
