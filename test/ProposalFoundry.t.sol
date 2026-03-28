// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseSetup} from "./BaseSetup.t.sol";
import {Proposal} from "../contracts/Proposal.sol";
import {Validators} from "../contracts/Validators.sol";
import {Staking} from "../contracts/Staking.sol";
import {StdStorage, stdStorage} from "forge-std/StdStorage.sol";

contract ProposalFoundryTest is BaseSetup {
    using stdStorage for StdStorage;

    address v1;
    address v2;
    address v3;

    function setUp() public {
        v1 = makeAddr("v1");
        v2 = makeAddr("v2");
        v3 = makeAddr("v3");
        address[] memory initVals = new address[](3);
        initVals[0] = v1;
        initVals[1] = v2;
        initVals[2] = v3;
        deploySystem(initVals);
    }

    function testVoteProposalWithResultAlreadyExists() public {
        Proposal p = Proposal(PROPOSAL);
        address newValidator = makeAddr("newValidator");

        // Create a proposal
        vm.warp(5_000_000);
        vm.prank(v1); // Use validator v1 as proposer
        bytes32 id = p.createProposal(newValidator, true, "");

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
        vm.prank(v1); // Use validator v1 as proposer
        bytes32 id = p.createProposal(newValidator, true, "");

        // Vote to make it rejected (majority against)
        vm.prank(v1);
        p.voteProposal(id, false);
        vm.prank(v2);
        p.voteProposal(id, false);
        vm.prank(v3);
        p.voteProposal(id, false);

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
        vm.prank(v1); // Use validator v1 as proposer
        bytes32 id = p.createProposal(testValidator, true, "");
        vm.prank(v1);
        p.voteProposal(id, true);
        vm.prank(v2);
        p.voteProposal(id, true);
        vm.prank(v3);
        p.voteProposal(id, true);

        // Now test with valid passed time but within deadline
        result = p.isProposalValidForStaking(testValidator);
        require(result, "should return true for valid proposal within deadline");

        // Test after deadline has passed
        uint256 proposalLastingPeriod = p.proposalLastingPeriod();
        vm.roll(block.number + proposalLastingPeriod + 1);
        result = p.isProposalValidForStaking(testValidator);
        require(!result, "should return false for proposal after deadline");
    }

    function testInitOnlyOnce() public {
        bytes memory err;
        try Proposal(PROPOSAL).initialize(new address[](0), address(0xCAFE), TEST_EPOCH) {
            revert("should revert");
        } catch (bytes memory e) {
            err = e;
        }
        require(err.length > 0, "expected revert");
    }

    function testCreateProposalConstraints() public {
        Proposal p = Proposal(PROPOSAL);
        uint256 cooldown = p.proposalCooldown();
        // can remove a not passed dst (choose an address not initialized) - now allowed
        address notPassed = makeAddr("np");
        vm.prank(v1);
        bytes32 removeId = p.createProposal(notPassed, false, "");
        require(removeId != bytes32(0), "remove should succeed even for not passed dst");

        // details too long
        string memory tooLong = new string(3001);
        vm.roll(block.number + cooldown);
        vm.prank(v1);
        vm.expectRevert("Details too long");
        p.createProposal(address(0xAAA1), true, tooLong);

        // ok to add not passed address
        vm.prank(v1);
        bytes32 id = p.createProposal(address(0xBBB2), true, "");
        require(id != bytes32(0), "create add should succeed");

        // can't add already exist dst after pass
        vm.prank(v1);
        p.voteProposal(id, true);
        vm.prank(v2);
        p.voteProposal(id, true);
        vm.prank(v3);
        p.voteProposal(id, true);
        vm.roll(block.number + cooldown);
        vm.prank(v1);
        vm.expectRevert("Can't add an already passed dst");
        p.createProposal(address(0xBBB2), true, "");
    }

    function testCreateAndVoteAddProposalPass() public {
        Proposal p = Proposal(PROPOSAL);
        address newValidator = address(0xBEEF);

        // create by validator v1
        vm.prank(v1);
        bytes32 id = p.createProposal(newValidator, true, "");
        require(id != bytes32(0), "create failed");

        // validators vote
        vm.prank(v1);
        p.voteProposal(id, true);
        vm.prank(v2);
        p.voteProposal(id, true);
        vm.prank(v3);
        p.voteProposal(id, true);

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
        // Create proposal using validator v1
        vm.prank(v1);
        bytes32 id = p.createProposal(address(0xCAFE), true, "");
        // non-validator
        address nv = makeAddr("nv");
        vm.prank(nv);
        (bool ok,) = address(p).call(abi.encodeWithSelector(p.voteProposal.selector, id, true));
        require(!ok, "should fail");
    }

    function testValidatorCanOnlyVoteOnceAndExpire() public {
        Proposal p = Proposal(PROPOSAL);
        // Create proposal using validator v1
        vm.prank(v1);
        bytes32 id = p.createProposal(address(0xDEAD), true, "");
        vm.prank(v1);
        p.voteProposal(id, true);
        // second vote by same validator should fail
        vm.prank(v1);
        (bool ok1,) = address(p).call(abi.encodeWithSelector(p.voteProposal.selector, id, true));
        require(!ok1, "double vote should fail");

        // expire by increasing block number
        uint256 lasting = p.proposalLastingPeriod();
        vm.roll(block.number + lasting + 1);
        vm.prank(v2);
        (bool ok2,) = address(p).call(abi.encodeWithSelector(p.voteProposal.selector, id, true));
        require(!ok2, "expired vote should fail");
    }

    function testRemoveProposalPass() public {
        Proposal p = Proposal(PROPOSAL);
        // remove an existing validator (v1)
        // Create proposal using validator v2
        vm.prank(v2);
        bytes32 id = p.createProposal(v1, false, "");
        vm.prank(v1);
        p.voteProposal(id, true);
        vm.prank(v2);
        p.voteProposal(id, true);
        vm.prank(v3);
        p.voteProposal(id, true);
        require(!p.pass(v1), "v1 should be unpassed");
    }

    function testConfigUpdateProposalLastingPeriod() public {
        _updateConfigAndAssert(0, 3600);
        require(Proposal(PROPOSAL).proposalLastingPeriod() == 3600, "cid0");
    }

    function testConfigUpdatePunishThreshold() public {
        _updateConfigAndAssert(1, 20);
        require(Proposal(PROPOSAL).punishThreshold() == 20, "cid1");
    }

    function testConfigUpdateRemoveThreshold() public {
        _updateConfigAndAssert(2, 60);
        require(Proposal(PROPOSAL).removeThreshold() == 60, "cid2");
    }

    function testConfigUpdateDecreaseRate() public {
        _updateConfigAndAssert(3, 30);
        require(Proposal(PROPOSAL).decreaseRate() == 30, "cid3");
    }

    function testConfigUpdateWithdrawProfitPeriod() public {
        _updateConfigAndAssert(4, 500);
        require(Proposal(PROPOSAL).withdrawProfitPeriod() == 500, "cid4");
    }

    function testConfigUpdateBlockReward() public {
        uint256 blockReward = 833_000_000_000_000_000;
        _updateConfigAndAssert(5, blockReward);
        require(Proposal(PROPOSAL).blockReward() == blockReward, "cid5");
    }

    function testConfigUpdateUnbondingPeriod() public {
        _updateConfigAndAssert(6, 604800);
        require(Proposal(PROPOSAL).unbondingPeriod() == 604800, "cid6");
    }

    function testConfigUpdateValidatorUnjailPeriod() public {
        _updateConfigAndAssert(7, 86400);
        require(Proposal(PROPOSAL).validatorUnjailPeriod() == 86400, "cid7");
    }

    function testConfigUpdateProposalCooldown() public {
        _updateConfigAndAssert(19, 123);
        require(Proposal(PROPOSAL).proposalCooldown() == 123, "cid19");
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

        // CID 0 - Invalid proposal period (zero value)
        vm.prank(v1);
        vm.expectRevert("Config value must be positive");
        p.createUpdateConfigProposal(0, 0); // Zero value is invalid
    }

    function testSetUnpassed() public {
        Proposal p = Proposal(PROPOSAL);
        Validators v = Validators(VALIDATORS);

        // Test that only Validators contract can call setUnpassed
        vm.prank(v1);
        vm.expectRevert();
        p.setUnpassed(v1);

        // Now test with Validators contract as caller
        vm.startPrank(address(v));
        bool result = p.setUnpassed(v1);
        vm.stopPrank();

        require(result, "setUnpassed should return true");
        require(!p.pass(v1), "validator should be unpassed");
    }

    function testIsProposalValidForStakingWithInvalidValidator() public {
        Proposal p = Proposal(PROPOSAL);

        // Test with invalid validator (pass is false)
        address invalidValidator = makeAddr("invalidValidator");
        bool result = p.isProposalValidForStaking(invalidValidator);
        require(!result, "should return false for invalid validator");

        // Test with valid validator that passed proposal and is within period
        address validValidator = makeAddr("validValidator");
        vm.warp(7_000_000);
        vm.prank(v1);
        bytes32 id = p.createProposal(validValidator, true, "");
        vm.prank(v1);
        p.voteProposal(id, true);
        vm.prank(v2);
        p.voteProposal(id, true);
        vm.prank(v3);
        p.voteProposal(id, true);

        result = p.isProposalValidForStaking(validValidator);
        require(result, "should return true for valid validator within period");

        // Test with valid validator that passed proposal and is outside period
        uint256 proposalLastingPeriod = p.proposalLastingPeriod();
        vm.roll(block.number + proposalLastingPeriod + 1);
        result = p.isProposalValidForStaking(validValidator);
        require(!result, "should return false for valid validator outside period");
    }

    function testCreateProposalExpiredPass() public {
        Proposal p = Proposal(PROPOSAL);
        address newValidator = makeAddr("newValidator");

        // 1. 创建一个添加验证者的提案，并使其通过
        vm.warp(8_000_000);
        vm.prank(v1);
        bytes32 id = p.createProposal(newValidator, true, "");
        vm.prank(v1);
        p.voteProposal(id, true);
        vm.prank(v2);
        p.voteProposal(id, true);
        vm.prank(v3);
        p.voteProposal(id, true);

        // 验证提案初始状态
        require(p.pass(newValidator), "validator should be passed initially");

        // 2. 等待提案过期
        uint256 proposalLastingPeriod = p.proposalLastingPeriod();
        vm.roll(block.number + proposalLastingPeriod + 1);

        // 3. 再次尝试创建相同的添加验证者提案
        // 这应该成功，因为提案已经过期，第182-185行的分支会被执行
        vm.prank(v1);
        bytes32 newId = p.createProposal(newValidator, true, "");

        // 4. 验证新提案成功创建
        require(newId != bytes32(0), "should be able to create new proposal after expiration");
        require(newId != id, "new proposal should have different ID");

        // 验证计数器状态被重置
        require(!p.pass(newValidator), "pass should be reset to false");
        require(p.proposalPassedHeight(newValidator) == 0, "proposalPassedHeight should be reset to 0");
    }

    function testVoteProposalInvalidTypeReverts() public {
        Proposal p = Proposal(PROPOSAL);
        address candidate = makeAddr("invalidTypeCandidate");

        vm.prank(v1);
        bytes32 id = p.createProposal(candidate, true, "");

        stdstore.target(address(p)).sig("proposals(bytes32)").with_key(id).depth(3).checked_write(uint256(99));

        vm.prank(v1);
        p.voteProposal(id, true);

        vm.prank(v2);
        vm.expectRevert("Invalid proposal type");
        p.voteProposal(id, true);
    }

    function _updateConfigAndAssert(uint256 cid, uint256 value) internal {
        Proposal p = Proposal(PROPOSAL);
        vm.prank(v1);
        bytes32 id = p.createUpdateConfigProposal(cid, value);
        vm.prank(v1);
        p.voteProposal(id, true);
        vm.prank(v2);
        p.voteProposal(id, true);
        vm.prank(v3);
        p.voteProposal(id, true);
    }
}
