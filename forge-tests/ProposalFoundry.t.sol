// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {BaseSetup} from "./BaseSetup.t.sol";
import {Proposal} from "../contracts/Proposal.sol";
import {Validators} from "../contracts/Validators.sol";

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
        try Proposal(PRO).initialize(new address[](0)) { revert("should revert"); } catch (bytes memory e) { err = e; }
        require(err.length > 0, "expected revert");
    }

    function testCreateProposalConstraints() public {
        Proposal p = Proposal(PRO);
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
        Proposal p = Proposal(PRO);
        // freeze timestamp to compute deterministic id
        vm.warp(1_000_000);
        bytes32 id = keccak256(abi.encodePacked(address(this), address(0xBEEF), true, "", block.timestamp));
        // create by anyone (this contract)
        (bool ok, ) = address(p).call(abi.encodeWithSelector(p.createProposal.selector, address(0xBEEF), true, ""));
        require(ok, "create failed");

        // validators vote
        vm.prank(v1); p.voteProposal(id, true);
        vm.prank(v2); p.voteProposal(id, true);
        vm.prank(v3); p.voteProposal(id, true);

        // assert pass recorded
        bool passed = p.pass(address(0xBEEF));
        require(passed, "should pass");
        // also ensure validator became top
        bool isTop = Validators(VAL).isTopValidator(address(0xBEEF));
        require(isTop, "should be top validator");
    }

    function testOnlyValidatorCanVote() public {
    Proposal p = Proposal(PRO);
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
        Proposal p = Proposal(PRO);
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
        Proposal p = Proposal(PRO);
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
        Proposal p = Proposal(PRO);
        uint256[7] memory cids = [uint256(0),1,2,3,4,5,6];
        uint256[6] memory vals = [uint256(100),200,300,400,500,600];
        address recv = makeAddr("recv");
        for (uint i = 0; i < cids.length; i++) {
            vm.warp(4_000_000 + i);
            bytes32 id = keccak256(abi.encodePacked(address(this), cids[i], i==6?uint256(uint160(recv)):vals[i], block.timestamp));
            p.createUpdateConfigProposal(cids[i], i==6?uint256(uint160(recv)):vals[i]);
            vm.prank(v1); p.voteProposal(id, true);
            vm.prank(v2); p.voteProposal(id, true);
            vm.prank(v3); p.voteProposal(id, true);

            if (cids[i]==0) require(p.proposalLastingPeriod()==vals[0], "cid0");
            else if (cids[i]==1) require(p.punishThreshold()==vals[1], "cid1");
            else if (cids[i]==2) require(p.removeThreshold()==vals[2], "cid2");
            else if (cids[i]==3) require(p.decreaseRate()==vals[3], "cid3");
            else if (cids[i]==4) require(p.withdrawProfitPeriod()==vals[4], "cid4");
            else if (cids[i]==5) require(p.increasePeriod()==vals[5], "cid5");
            else if (cids[i]==6) require(p.receiverAddr()==recv, "cid6");
        }
    }
}
