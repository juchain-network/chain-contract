// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {BaseSetup} from "./BaseSetup.t.sol";
import {Proposal} from "../../contracts/Proposal.sol";
import {Validators} from "../../contracts/Validators.sol";
import {Staking} from "../../contracts/Staking.sol";

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

    function testVoteProposalWithResultAlreadyExists() public {
        Proposal p = Proposal(PROPOSAL);
        address newValidator = makeAddr("newValidator");
        
        // Create a proposal
        vm.warp(5_000_000);
        bytes32 id = keccak256(abi.encodePacked(v1, newValidator, true, "", block.timestamp));
        vm.prank(v1); // Use validator v1 as proposer
        p.createProposal(newValidator, true, "");
        
        // Manually set resultExist to true to test the early return path
        // This simulates a situation where the result was already determined by previous votes
        vm.store(
            address(p),
            keccak256(abi.encode(id, uint256(2))), // Slot for results[id]
            bytes32(uint256(1 << 160)) // resultExist = true, agree=0, reject=0
        );
        
        // Try to vote - should return true without processing the vote
        vm.prank(v1);
        bool result = p.voteProposal(id, true);
        require(result, "should return true when result already exists");
    }

    function testVoteProposalRejectPath() public {
        Proposal p = Proposal(PROPOSAL);
        address newValidator = makeAddr("newValidator");
        
        // Create a proposal
        vm.warp(5_000_000);
        bytes32 id = keccak256(abi.encodePacked(v1, newValidator, true, "", block.timestamp));
        vm.prank(v1); // Use validator v1 as proposer
        p.createProposal(newValidator, true, "");
        
        // Vote to make it rejected (majority against)
        vm.prank(v1); p.voteProposal(id, false);
        vm.prank(v2); p.voteProposal(id, false);
        vm.prank(v3); p.voteProposal(id, false);
        
        // Check that reject path was taken
        (uint16 agree, uint16 reject, bool resultExist) = p.results(id);
        require(agree == 0, "should have 0 agree votes");
        require(reject == 3, "should have 3 reject votes");
        require(resultExist, "result should exist");
    }

    function testIsProposalValidForStaking() public {
        Proposal p = Proposal(PROPOSAL);
        
        // Test validator with no proposal (pass is false)
        address newValidator = makeAddr("newValidator");
        bool result = p.isProposalValidForStaking(newValidator);
        require(!result, "should return false for validator with no proposal");
        
        // Test existing validator but no passed time recorded
        // (pass[v1] is true since it was initialized)
        // Let's assume v1 has no passed time (in real code it should, but for testing we'll proceed)
        
        // Create a new validator proposal and pass it
        address testValidator = makeAddr("testValidator");
        vm.warp(5_000_000);
        bytes32 id = keccak256(abi.encodePacked(v1, testValidator, true, "", block.timestamp));
        vm.prank(v1); // Use validator v1 as proposer
        p.createProposal(testValidator, true, "");
        vm.prank(v1); p.voteProposal(id, true);
        vm.prank(v2); p.voteProposal(id, true);
        vm.prank(v3); p.voteProposal(id, true);
        
        // Now test with valid passed time but within deadline
        result = p.isProposalValidForStaking(testValidator);
        require(result, "should return true for valid proposal within deadline");
        
        // Test after deadline has passed
        vm.warp(block.timestamp + p.STAKING_DEADLINE_PERIOD() + 1);
        result = p.isProposalValidForStaking(testValidator);
        require(!result, "should return false for proposal after deadline");
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
        vm.prank(v1);
        vm.expectRevert("Can't add an already exist dst or Can't remove a not passed dst");
        p.createProposal(notPassed, false, "");

        // details too long
        string memory tooLong = new string(3001);
        vm.prank(v1);
        vm.expectRevert("Details too long");
        p.createProposal(address(0xAAA1), true, tooLong);

    // ok to add not passed address; freeze time to compute id
    vm.warp(1_200_000);
    bytes32 id = keccak256(abi.encodePacked(v1, address(0xBBB2), true, "", block.timestamp));
    vm.prank(v1);
    bytes32 proposalId3 = p.createProposal(address(0xBBB2), true, "");
    require(proposalId3 != bytes32(0), "create add should succeed");

    // can't add already exist dst after pass
    vm.prank(v1); p.voteProposal(id, true);
    vm.prank(v2); p.voteProposal(id, true);
    vm.prank(v3); p.voteProposal(id, true);
    vm.prank(v1);
    vm.expectRevert("Can't add an already exist dst or Can't remove a not passed dst");
    p.createProposal(address(0xBBB2), true, "");
    }

    function testCreateAndVoteAddProposalPass() public {
        Proposal p = Proposal(PROPOSAL);
        address newValidator = address(0xBEEF);
        
        // freeze timestamp to compute deterministic id
        vm.warp(1_000_000);
        bytes32 id = keccak256(abi.encodePacked(v1, newValidator, true, "", block.timestamp));
        // create by validator v1
        vm.prank(v1);
        bytes32 proposalId = p.createProposal(newValidator, true, "");
        require(proposalId != bytes32(0), "create failed");

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
        uint256 minStake = Proposal(PROPOSAL).minValidatorStake();
        vm.deal(newValidator, minStake);
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
    bytes32 id = keccak256(abi.encodePacked(v1, address(0xCAFE), true, "", block.timestamp));
    // Create proposal using validator v1
    vm.prank(v1);
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
        bytes32 id = keccak256(abi.encodePacked(v1, address(0xDEAD), true, "", block.timestamp));
        // Create proposal using validator v1
        vm.prank(v1);
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
        bytes32 id = keccak256(abi.encodePacked(v2, v1, false, "", block.timestamp));
        // Create proposal using validator v2
        vm.prank(v2);
        p.createProposal(v1, false, "");
        vm.prank(v1); p.voteProposal(id, true);
        vm.prank(v2); p.voteProposal(id, true);
        vm.prank(v3); p.voteProposal(id, true);
        require(!p.pass(v1), "v1 should be unpassed");
    }

    function testConfigUpdateAll() public {
        Proposal p = Proposal(PROPOSAL);
        uint256[8] memory cids = [uint256(0),1,2,3,4,5,6,7];

        uint256[8] memory vals = [uint256(3600),200,300,400,500,833_000_000_000_000_000,604800,86400];
        for (uint i = 0; i < cids.length; i++) {
            vm.warp(4_000_000 + i);
            bytes32 id = keccak256(abi.encodePacked(v1, cids[i], vals[i], block.timestamp));
            // Create config proposal using validator v1
            vm.prank(v1);
            p.createUpdateConfigProposal(cids[i], vals[i]);
            vm.prank(v1); p.voteProposal(id, true);
            vm.prank(v2); p.voteProposal(id, true);
            vm.prank(v3); p.voteProposal(id, true);

            if (cids[i]==0) require(p.proposalLastingPeriod()==3600, "cid0");
            else if (cids[i]==1) require(p.punishThreshold()==vals[1], "cid1");
            else if (cids[i]==2) require(p.removeThreshold()==vals[2], "cid2");
            else if (cids[i]==3) require(p.decreaseRate()==vals[3], "cid3");
            else if (cids[i]==4) require(p.withdrawProfitPeriod()==vals[4], "cid4");
            else if (cids[i]==5) require(p.blockReward()==vals[5], "cid5");
            else if (cids[i]==6) require(p.unbondingPeriod()==vals[6], "cid6");
            else if (cids[i]==7) require(p.validatorUnjailPeriod()==vals[7], "cid7");
        }
    }

    function testUpdateConfigWithInvalidCID() public {
        // Test updateConfig with an invalid CID (should revert during proposal creation)
        Proposal p = Proposal(PROPOSAL);
        
        // CID 100 is invalid, should revert during proposal creation
        vm.warp(6_000_000);
        vm.prank(v1);
        vm.expectRevert("Invalid config ID");
        p.createUpdateConfigProposal(100, 12345);
    }

    function testUpdateConfigRequireChecks() public {
        Proposal p = Proposal(PROPOSAL);
        
        // Test that the validation checks work during proposal creation
        
        // CID 0 - Invalid proposal period (too small)
        vm.prank(v1);
        vm.expectRevert("Invalid proposal period");
        p.createUpdateConfigProposal(0, 100); // Less than 1 hour
    }
}
