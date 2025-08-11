// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {BaseSetup} from "./BaseSetup.t.sol";
import {Validators} from "../contracts/Validators.sol";
import {Proposal} from "../contracts/Proposal.sol";

// 完整的验证者测试，对应 test/validators.js 的所有功能
contract ValidatorsCompleteFoundryTest is BaseSetup {

    address v1; address v2; address v3;
    address miner;
    uint256 constant Active = 1;
    uint256 constant Jailed = 2;

    function setUp() public {
        v1 = makeAddr("v1");
        v2 = makeAddr("v2"); 
        v3 = makeAddr("v3");
        address[] memory initVals = new address[](3);
        initVals[0] = v1; initVals[1] = v2; initVals[2] = v3;
        deploySystem(initVals);
        miner = v1;
        vm.coinbase(miner);
    }

    function testCanOnlyInitOnce() public {
        // 对应 "can only init once"
        bytes memory err;
        try Validators(VAL).initialize(new address[](0)) { 
            revert("should revert"); 
        } catch (bytes memory e) { 
            err = e; 
        }
        require(err.length > 0, "expected revert");
    }

    // 创建/编辑验证者相关测试
    function testCreateValidatorInvalidFeeAddr() public {
        // 对应 "can't create validator if fee addr == address(0)"
        address validator = makeAddr("validator");
        vm.prank(validator);
        (bool ok, ) = address(Validators(VAL)).call(
            abi.encodeWithSelector(
                Validators.createOrEditValidator.selector,
                address(0), "", "", "", "", ""
            )
        );
        require(!ok, "should fail with zero address");
    }

    function testCreateValidatorInvalidDescription() public {
        // 对应 "can't create validator if describe info invalid"
        address validator = makeAddr("validator");
        string memory tooLongMoniker = _generateLongString(71); // > 70 limit
        vm.prank(validator);
        (bool ok, ) = address(Validators(VAL)).call(
            abi.encodeWithSelector(
                Validators.createOrEditValidator.selector,
                payable(validator), tooLongMoniker, "", "", "", ""
            )
        );
        require(!ok, "should fail with too long moniker");
    }

    function testCreateValidatorNotAuthorized() public {
        // 对应 "can't create validator if not pass propose"
        address validator = makeAddr("validator");
        vm.prank(validator);
        (bool ok, ) = address(Validators(VAL)).call(
            abi.encodeWithSelector(
                Validators.createOrEditValidator.selector,
                payable(validator), "", "", "", "", ""
            )
        );
        require(!ok, "should fail without authorization");
    }

    function testCreateValidatorSuccess() public {
        // 对应 "create validator"
        address validator = makeAddr("validator");
        
        // 先通过提案授权
        _passProposal(validator, true);
        
        // 创建验证者
        vm.prank(validator);
        bool success = Validators(VAL).createOrEditValidator(payable(validator), "", "", "", "", "");
        require(success, "create validator should succeed");
        
        // 检查状态
        (, Validators.Status status,,,) = Validators(VAL).getValidatorInfo(validator);
        require(uint256(status) == Active, "validator should be active");
    }

    function testEditValidatorInfo() public {
        // 对应 "edit validator info"
        address validator = makeAddr("validator");
        address feeAddr = makeAddr("feeAddr");
        
        // 先授权并创建
        _passProposal(validator, true);
        vm.prank(validator);
        Validators(VAL).createOrEditValidator(payable(validator), "", "", "", "", "");
        
        // 编辑 fee 地址
        vm.prank(validator);
        bool success = Validators(VAL).createOrEditValidator(payable(feeAddr), "", "", "", "", "");
        require(success, "edit should succeed");
        
        // 检查 fee 地址已更新
        (address payable actualFeeAddr,,,,) = Validators(VAL).getValidatorInfo(validator);
        require(actualFeeAddr == feeAddr, "fee address should be updated");
    }

    // 提案添加/移除验证者测试
    function testProposeAddNewValidator() public {
        // 对应 "propose add a new val"
        address nval = makeAddr("newval");
        
        // 初始不是验证者
        require(!Validators(VAL).isTopValidator(nval), "should not be validator initially");
        
        // 创建并投票通过提案
        _passProposal(nval, true);
        
        // 现在应该是验证者
        require(Validators(VAL).isTopValidator(nval), "should be validator after proposal");
        require(Proposal(PRO).pass(nval), "should be marked as passed");
    }

    function testProposeRemoveValidator() public {
        // 对应 "propose remove a val"
        
        // v1 初始是验证者
        require(Validators(VAL).isTopValidator(v1), "v1 should be validator initially");
        
        // 创建并投票通过移除提案
        _passProposal(v1, false);
        
        // 现在应该不是验证者
        require(!Validators(VAL).isTopValidator(v1), "v1 should not be validator after removal");
        require(!Proposal(PRO).pass(v1), "v1 should not be marked as passed");
    }

    // 区块奖励分发测试
    function testDistributeBlockReward() public {
        // 对应 "miner can distribute to validator contract, the profits should be right updated"
        uint256 fee = 0.3 ether;
        uint256 expectPerFee = 0.1 ether;
        
        vm.prank(miner);
        Validators(VAL).distributeBlockReward{value: fee}();
        
        // 检查每个验证者获得的奖励
        for (uint i = 0; i < 3; i++) {
            address val = i == 0 ? v1 : (i == 1 ? v2 : v3);
            (, , uint256 aacIncoming,,) = Validators(VAL).getValidatorInfo(val);
            require(aacIncoming == expectPerFee, "should get expected fee");
        }
    }

    function testUpdateWithdrawProfitPeriod() public {
        // 对应 "update withdraw profit wait block"
        vm.warp(5_000_000);
        bytes32 id = keccak256(abi.encodePacked(address(this), uint256(4), uint256(10), block.timestamp));
        Proposal(PRO).createUpdateConfigProposal(4, 10);
        
        vm.prank(v1); Proposal(PRO).voteProposal(id, true);
        vm.prank(v2); Proposal(PRO).voteProposal(id, true);
        vm.prank(v3); Proposal(PRO).voteProposal(id, true);
        
        require(Proposal(PRO).withdrawProfitPeriod() == 10, "withdraw period should be updated");
    }

    function testValidatorWithdrawProfits() public {
        // 对应 "validator can withdraw profits"
        uint256 fee = 0.3 ether;
        
        // 分发奖励
        vm.prank(miner);
        Validators(VAL).distributeBlockReward{value: fee}();
        
        // 设置短的提取周期
        _updateWithdrawPeriod(1);
        vm.roll(block.number + 2);
        
        uint256 balBefore = miner.balance;
        vm.prank(miner);
        Validators(VAL).withdrawProfits(miner);
        uint256 balAfter = miner.balance;
        
        require(balAfter > balBefore, "should receive profits");
        
        // 测试不同的 fee 地址
        address feeAddr = makeAddr("feeAddr");
        vm.prank(miner);
        Validators(VAL).createOrEditValidator(payable(feeAddr), "", "", "", "", "");
        
        // 再次分发
        vm.prank(miner);
        Validators(VAL).distributeBlockReward{value: 0.5 ether}();
        
        vm.roll(block.number + 2);
        
        uint256 feeBalBefore = feeAddr.balance;
        vm.prank(feeAddr);
        Validators(VAL).withdrawProfits(miner);
        uint256 feeBalAfter = feeAddr.balance;
        
        require(feeBalAfter > feeBalBefore, "fee address should receive profits");
    }

    function testCantWithdrawWithoutProfits() public {
        // 对应 "Can't call withdrawProfits if you don't have any profits"
        address feeAddr = makeAddr("feeAddr");
        _updateWithdrawPeriod(1);
        vm.roll(block.number + 2);
        
        vm.prank(feeAddr);
        (bool ok, ) = address(Validators(VAL)).call(
            abi.encodeWithSelector(Validators.withdrawProfits.selector, miner)
        );
        require(!ok, "should fail without profits");
    }

    function testUpdateActiveValidatorSet() public {
        // 对应 "update active validator set"
        uint256 epoch = 30;
        address[] memory newSet = new address[](2);
        newSet[0] = v1;
        newSet[1] = v2;
        
        // 模拟到达 epoch 边界 (block.number % epoch == 0)
        uint256 targetBlock = ((block.number / epoch) + 1) * epoch;
        vm.roll(targetBlock);
        
        vm.prank(miner);
        Validators(VAL).updateActiveValidatorSet(newSet, epoch);
        
        // 验证新的验证者集合
        address[] memory activeSet = Validators(VAL).getActiveValidators();
        require(activeSet.length == 2, "should have 2 active validators");
        require(activeSet[0] == v1, "should contain v1");
        require(activeSet[1] == v2, "should contain v2");
    }

    // 辅助函数
    function _passProposal(address target, bool flag) internal {
        vm.warp(block.timestamp + 1000000);
        bytes32 id = keccak256(abi.encodePacked(address(this), target, flag, "", block.timestamp));
        Proposal(PRO).createProposal(target, flag, "");
        
        vm.prank(v1); Proposal(PRO).voteProposal(id, true);
        vm.prank(v2); Proposal(PRO).voteProposal(id, true);
        vm.prank(v3); Proposal(PRO).voteProposal(id, true);
    }

    function _updateWithdrawPeriod(uint256 period) internal {
        vm.warp(block.timestamp + 1000000);
        bytes32 id = keccak256(abi.encodePacked(address(this), uint256(4), period, block.timestamp));
        Proposal(PRO).createUpdateConfigProposal(4, period);
        
        vm.prank(v1); Proposal(PRO).voteProposal(id, true);
        vm.prank(v2); Proposal(PRO).voteProposal(id, true);
        vm.prank(v3); Proposal(PRO).voteProposal(id, true);
    }

    function _generateLongString(uint256 length) internal pure returns (string memory) {
        bytes memory result = new bytes(length);
        for (uint i = 0; i < length; i++) {
            result[i] = "a";
        }
        return string(result);
    }
}
