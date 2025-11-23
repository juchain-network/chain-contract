// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {BaseSetup} from "./BaseSetup.t.sol";
import {Proposal} from "../contracts/Proposal.sol";
import {Validators} from "../contracts/Validators.sol";
import {Staking} from "../contracts/Staking.sol";

contract ProposalFoundryTest is BaseSetup {

    address v1;
    address v2;
    address v3;

    function setUp() public {
        v1 = makeAddr("v1");
        v2 = makeAddr("v2");
        v3 = makeAddr("v3");
        address[] memory initVals = new address[](3);
        initVals[0] = v1; initVals[1] = v2; initVals[2] = v3;
        deploySystem(initVals);
    }

    function testInitOnlyOnce() public {
        bytes memory err;
        try Proposal(PROPOSAL).initialize(new address[](0), address(0xCAFE)) { revert("should revert"); } catch (bytes memory e) { err = e; }
        require(err.length > 0, "expected revert");
    }

    function testCreateProposalConstraints() public {
        Proposal p = Proposal(PROPOSAL);
        // can't remove a not passed dst (choose an address not initialized)
        address notPassed = makeAddr("np");
        (bool ok1, ) = address(p).call(abi.encodeWithSelector(p.createProposal.selector, notPassed, false, ""));
        require(!ok1, "remove not-exist should fail");

        // details too long
        string memory tooLong = new string(3001);
        (bool ok2, ) = address(p).call(abi.encodeWithSelector(p.createProposal.selector, address(0xAAA1), true, tooLong));
        require(!ok2, "details too long should fail");

    // ok to add not passed address; freeze time to compute id
    vm.warp(1_200_000);
    bytes32 id = keccak256(abi.encodePacked(address(this), address(0xBBB2), true, "", block.timestamp));
    (bool ok3, ) = address(p).call(abi.encodeWithSelector(p.createProposal.selector, address(0xBBB2), true, ""));
    require(ok3, "create add should succeed");

    // can't add already exist dst after pass
    vm.prank(v1); p.voteProposal(id, true);
    vm.prank(v2); p.voteProposal(id, true);
    vm.prank(v3); p.voteProposal(id, true);
    (bool ok4, ) = address(p).call(abi.encodeWithSelector(p.createProposal.selector, address(0xBBB2), true, ""));
    require(!ok4, "add already passed should fail");
    }

    function testCreateAndVoteAddProposalPass() public {
        Proposal p = Proposal(PROPOSAL);
        address newValidator = address(0xBEEF);
        
        // freeze timestamp to compute deterministic id
        vm.warp(1_000_000);
        bytes32 id = keccak256(abi.encodePacked(address(this), newValidator, true, "", block.timestamp));
        // create by anyone (this contract)
        (bool ok, ) = address(p).call(abi.encodeWithSelector(p.createProposal.selector, newValidator, true, ""));
        require(ok, "create failed");

        // validators vote
        vm.prank(v1); p.voteProposal(id, true);
        vm.prank(v2); p.voteProposal(id, true);
        vm.prank(v3); p.voteProposal(id, true);

        // assert pass recorded
        bool passed = p.pass(newValidator);
        require(passed, "should pass");
        
        // In POSA mode, validator must register (stake) to become top validator
        // This is the design: proposal passing only grants permission, validator must actively register
        // Give new validator enough ETH and register
        uint256 minStake = Staking(STAKING).MIN_VALIDATOR_STAKE();
        vm.deal(newValidator, 20000 ether);
        vm.prank(newValidator);
        Staking(STAKING).registerValidator{value: minStake}(1000); // 10% commission
        
        // registerValidator() internally calls tryAddValidatorToHighestSet(), so validator should now be top
        // also ensure validator became top
        bool isTop = Validators(VALIDATORS).isTopValidator(newValidator);
        require(isTop, "should be top validator");
    }

    function testOnlyValidatorCanVote() public {
    Proposal p = Proposal(PROPOSAL);
    vm.warp(1_111_111);
    bytes32 id = keccak256(abi.encodePacked(address(this), address(0xCAFE), true, "", block.timestamp));
    p.createProposal(address(0xCAFE), true, "");
        // non-validator
        address nv = makeAddr("nv");
    vm.prank(nv);
        (bool ok, ) = address(p).call(abi.encodeWithSelector(p.voteProposal.selector, id, true));
        require(!ok, "should fail");
    }

    function testValidatorCanOnlyVoteOnceAndExpire() public {
        Proposal p = Proposal(PROPOSAL);
        vm.warp(2_222_222);
        bytes32 id = keccak256(abi.encodePacked(address(this), address(0xDEAD), true, "", block.timestamp));
        p.createProposal(address(0xDEAD), true, "");
        vm.prank(v1); p.voteProposal(id, true);
        // second vote by same validator should fail
        vm.prank(v1);
        (bool ok1, ) = address(p).call(abi.encodeWithSelector(p.voteProposal.selector, id, true));
        require(!ok1, "double vote should fail");

        // expire
        uint256 lasting = p.proposalLastingPeriod();
        vm.warp(block.timestamp + lasting + 1);
        vm.prank(v2);
        (bool ok2, ) = address(p).call(abi.encodeWithSelector(p.voteProposal.selector, id, true));
        require(!ok2, "expired vote should fail");
    }

    function testRemoveProposalPass() public {
        Proposal p = Proposal(PROPOSAL);
        // remove an existing validator (v1)
        vm.warp(3_000_000);
        bytes32 id = keccak256(abi.encodePacked(address(this), v1, false, "", block.timestamp));
        p.createProposal(v1, false, "");
        vm.prank(v1); p.voteProposal(id, true);
        vm.prank(v2); p.voteProposal(id, true);
        vm.prank(v3); p.voteProposal(id, true);
        require(!p.pass(v1), "v1 should be unpassed");
    }

    function testConfigUpdateAll() public {
        Proposal p = Proposal(PROPOSAL);
        // 注意: cid 5 (increasePeriod) 和 cid 6 (receiverAddr) 已移除，系统不再支持代币增发
        uint256[5] memory cids = [uint256(0),1,2,3,4];
        // cid 0 (proposalLastingPeriod) 需要 >= 1 hours && <= 30 days，使用 3600 秒（1小时）
        uint256[5] memory vals = [uint256(3600),200,300,400,500];
        for (uint i = 0; i < cids.length; i++) {
            vm.warp(4_000_000 + i);
            bytes32 id = keccak256(abi.encodePacked(address(this), cids[i], vals[i], block.timestamp));
            p.createUpdateConfigProposal(cids[i], vals[i]);
            vm.prank(v1); p.voteProposal(id, true);
            vm.prank(v2); p.voteProposal(id, true);
            vm.prank(v3); p.voteProposal(id, true);

            if (cids[i]==0) require(p.proposalLastingPeriod()==3600, "cid0");
            else if (cids[i]==1) require(p.punishThreshold()==vals[1], "cid1");
            else if (cids[i]==2) require(p.removeThreshold()==vals[2], "cid2");
            else if (cids[i]==3) require(p.decreaseRate()==vals[3], "cid3");
            else if (cids[i]==4) require(p.withdrawProfitPeriod()==vals[4], "cid4");
        }
    }
}
