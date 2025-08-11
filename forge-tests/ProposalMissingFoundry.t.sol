// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {BaseSetup} from "./BaseSetup.t.sol";
import {Proposal} from "../contracts/Proposal.sol";
import {Validators} from "../contracts/Validators.sol";

// 补充缺失的 Proposal 测试用例
contract ProposalMissingFoundryTest is BaseSetup {

    address v1; address v2; address v3; address v4; address v5;

    function setUp() public {
        v1 = makeAddr("v1");
        v2 = makeAddr("v2");
        v3 = makeAddr("v3");
        v4 = makeAddr("v4");
        v5 = makeAddr("v5");
        address[] memory initVals = new address[](5); // 使用5个验证者以便测试拒绝场景
        initVals[0] = v1; initVals[1] = v2; initVals[2] = v3; initVals[3] = v4; initVals[4] = v5;
        deploySystem(initVals);
    }

    function testAnyoneCanCreateProposal() public {
        // 对应 "anyone can create proposal"
        Proposal p = Proposal(PRO);
        address candidate = makeAddr("candidate");
        
        // 测试多个不同账户都可以创建提案
        for (uint i = 0; i < 5; i++) {
            address creator = makeAddr(string(abi.encodePacked("creator", i)));
            vm.deal(creator, 1 ether);
            
            vm.warp(1_000_000 + i * 10000); // 不同时间戳避免 ID 冲突
            vm.prank(creator);
            bool success = p.createProposal(candidate, true, "");
            require(success, "should create proposal successfully");
        }
    }

    function testProposalReject() public {
        // 对应 "normal vote(2 agree, 3 reject)"
        Proposal p = Proposal(PRO);
        address candidate = makeAddr("candidate");
        
        vm.warp(2_000_000);
        bytes32 id = keccak256(abi.encodePacked(address(this), candidate, true, "test", block.timestamp));
        p.createProposal(candidate, true, "test");
        
        // 2票同意
        vm.prank(v1); p.voteProposal(id, true);
        vm.prank(v2); p.voteProposal(id, true);
        
        // 3票反对
        vm.prank(v3); p.voteProposal(id, false);
        vm.prank(v4); p.voteProposal(id, false);
        vm.prank(v5); p.voteProposal(id, false);
        
        // 提案应该被拒绝，候选人不应该通过
        require(!p.pass(candidate), "candidate should not pass");
        
        // 检查投票结果
        (uint16 agree, uint16 reject, bool resultExist) = p.results(id);
        require(agree == 2, "should have 2 agree votes");
        require(reject == 3, "should have 3 reject votes");
        require(resultExist, "result should exist");
    }

    function testValidateCandidateInfo() public {
        // 对应 "Validate candidate's info"
        Proposal p = Proposal(PRO);
        address candidate = makeAddr("candidate");
        address proposer = makeAddr("proposer");
        
        vm.warp(3_000_000);
        bytes32 id = keccak256(abi.encodePacked(proposer, candidate, true, "test details", block.timestamp));
        
        vm.prank(proposer);
        p.createProposal(candidate, true, "test details");
        
        // 投票通过
        vm.prank(v1); p.voteProposal(id, true);
        vm.prank(v2); p.voteProposal(id, true);
        vm.prank(v3); p.voteProposal(id, true);
        
        // 验证提案信息
        (address storedProposer, uint256 createTime, uint256 proposalType, address dst, bool flag, string memory details, , ) = p.proposals(id);
        require(storedProposer == proposer, "proposer should match");
        require(dst == candidate, "candidate should match");
        require(flag == true, "flag should be true");
        require(keccak256(bytes(details)) == keccak256(bytes("test details")), "details should match");
        require(proposalType == 1, "should be validator proposal type");
        require(createTime == 3_000_000, "create time should match");
        
        // 验证投票结果
        (uint16 agree, uint16 reject, bool resultExist) = p.results(id);
        require(agree == 3, "should have 3 agree votes");
        require(reject == 0, "should have 0 reject votes");
        require(resultExist, "result should exist");
        
        // 验证候选人状态
        require(p.pass(candidate), "candidate should pass");
    }

    function testSetUnpassedPermission() public {
        // 对应 "only validator can set val unpass"
        Proposal p = Proposal(PRO);
        address candidate = v1; // 使用已存在的验证者
        
        // 非验证者合约调用应该失败
        vm.prank(makeAddr("random"));
        (bool ok, ) = address(p).call(abi.encodeWithSelector(p.setUnpassed.selector, candidate));
        require(!ok, "should fail when called by non-validator contract");
    }

    function testSetUnpassedByValidatorContract() public {
        // 对应 "validator contract can set val unpass"
        // 注意：这个测试需要从 Validators 合约调用，我们在这里模拟这个过程
        
        address candidate = v1;
        
        // 确认候选人初始是通过的
        require(Proposal(PRO).pass(candidate), "candidate should initially pass");
        
        // 模拟 Validators 合约调用 setUnpassed
        // 在实际测试中，这通过惩罚流程或移除提案自动触发
        // 我们通过移除验证者提案来测试这个功能
        vm.warp(4_000_000);
        bytes32 id = keccak256(abi.encodePacked(address(this), candidate, false, "", block.timestamp));
        Proposal(PRO).createProposal(candidate, false, "");
        
        // 投票移除
        vm.prank(v2); Proposal(PRO).voteProposal(id, true);
        vm.prank(v3); Proposal(PRO).voteProposal(id, true);
        vm.prank(v4); Proposal(PRO).voteProposal(id, true);
        
        // 验证候选人现在不通过
        require(!Proposal(PRO).pass(candidate), "candidate should not pass after removal");
    }
}
