// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseSetup} from "./BaseSetup.t.sol";
import {Validators} from "../contracts/Validators.sol";
import {Proposal} from "../contracts/Proposal.sol";
import {Staking} from "../contracts/Staking.sol";

contract ValidatorSignerFoundryTest is BaseSetup {
    address internal v1;
    address internal v2;
    address internal v3;
    address internal s1;
    address internal s2;
    address internal s3;

    function setUp() public {
        v1 = makeAddr("v1");
        v2 = makeAddr("v2");
        v3 = makeAddr("v3");
        s1 = vm.addr(0x101);
        s2 = vm.addr(0x102);
        s3 = vm.addr(0x103);

        address[] memory initVals = new address[](3);
        initVals[0] = v1;
        initVals[1] = v2;
        initVals[2] = v3;

        address[] memory initSigners = new address[](3);
        initSigners[0] = s1;
        initSigners[1] = s2;
        initSigners[2] = s3;

        deploySystem(initVals, initSigners);
    }

    function testGenesisSeparateSignerBindings() public view {
        assertEq(Validators(VALIDATORS).getValidatorSigner(v1), s1);
        assertEq(Validators(VALIDATORS).getValidatorSigner(v2), s2);
        assertEq(Validators(VALIDATORS).getValidatorSigner(v3), s3);

        assertEq(Validators(VALIDATORS).getValidatorBySigner(s1), v1);
        assertEq(Validators(VALIDATORS).getValidatorBySigner(s2), v2);
        assertEq(Validators(VALIDATORS).getValidatorBySigner(s3), v3);
        assertEq(Validators(VALIDATORS).getValidatorBySignerHistory(s1), v1);
        assertEq(Validators(VALIDATORS).getValidatorBySignerHistory(s2), v2);
        assertEq(Validators(VALIDATORS).getValidatorBySignerHistory(s3), v3);

        address[] memory activeSigners = Validators(VALIDATORS).getActiveSigners();
        assertEq(activeSigners.length, 3);
        assertEq(activeSigners[0], s1);
        assertEq(activeSigners[1], s2);
        assertEq(activeSigners[2], s3);
    }

    function testRewardEligibleSignersWithStakesUsesSignerAddresses() public view {
        (address[] memory signers, uint256[] memory totalStakes) =
            Validators(VALIDATORS).getRewardEligibleSignersWithStakes();

        assertEq(signers.length, 3);
        assertEq(totalStakes.length, 3);
        assertEq(signers[0], s1);
        assertEq(signers[1], s2);
        assertEq(signers[2], s3);

        (uint256 selfStake1, uint256 totalDelegated1,,,,,,,,) = Staking(STAKING).getValidatorInfo(v1);
        (uint256 selfStake2, uint256 totalDelegated2,,,,,,,,) = Staking(STAKING).getValidatorInfo(v2);
        (uint256 selfStake3, uint256 totalDelegated3,,,,,,,,) = Staking(STAKING).getValidatorInfo(v3);

        assertEq(totalStakes[0], selfStake1 + totalDelegated1);
        assertEq(totalStakes[1], selfStake2 + totalDelegated2);
        assertEq(totalStakes[2], selfStake3 + totalDelegated3);
    }

    function testBlockRewardUsesSignerButCreditsValidator() public {
        uint256 fee = 1 ether;
        vm.coinbase(s1);
        vm.deal(s1, fee);

        vm.prank(s1);
        Validators(VALIDATORS).distributeBlockReward{value: fee}();

        (,, uint256 incoming,,) = Validators(VALIDATORS).getValidatorInfo(v1);
        assertEq(incoming, fee);

        (,, uint256 otherIncoming,,) = Validators(VALIDATORS).getValidatorInfo(s1);
        assertEq(otherIncoming, 0);
    }

    function testStakingRewardUsesSignerButCreditsValidator() public {
        uint256 reward = 3 ether;
        vm.coinbase(s1);
        vm.deal(s1, reward);

        vm.prank(s1);
        Staking(STAKING).distributeRewards{value: reward}();

        (,,, uint256 accumulatedRewards,,,,,,) = Staking(STAKING).getValidatorInfo(v1);
        assertGt(accumulatedRewards, 0);

        (,,, uint256 signerRewards,,,,,,) = Staking(STAKING).getValidatorInfo(s1);
        assertEq(signerRewards, 0);
        assertEq(Staking(STAKING).lastActiveBlock(v1), block.number);
    }

    function testRegisterValidatorWithExplicitSigner() public {
        address validator = makeAddr("validator");
        address signer = vm.addr(0x104);
        uint256 minStake = Proposal(PROPOSAL).minValidatorStake();

        _passProposal(validator, true);

        vm.prank(validator);
        Validators(VALIDATORS).createOrEditValidator(payable(validator), signer, "", "", "", "", "");

        vm.deal(validator, minStake);
        vm.prank(validator);
        Staking(STAKING).registerValidator{value: minStake}(1000);

        assertEq(Validators(VALIDATORS).getValidatorSigner(validator), signer);
        assertEq(Validators(VALIDATORS).getValidatorBySigner(signer), validator);
        assertEq(Validators(VALIDATORS).getValidatorBySignerHistory(signer), address(0));
    }

    function _passProposal(address target, bool flag) internal {
        vm.warp(block.timestamp + 1_000_000);
        vm.roll(block.number + 101);

        vm.prank(v1);
        bytes32 id = Proposal(PROPOSAL).createProposal(target, flag, "");

        vm.prank(v1);
        Proposal(PROPOSAL).voteProposal(id, true);
        vm.prank(v2);
        Proposal(PROPOSAL).voteProposal(id, true);
    }
}

contract ValidatorSignerEpochFoundryTest is BaseSetup {
    address internal v1;
    address internal v2;
    address internal v3;
    address internal s1;
    address internal s2;
    address internal s3;

    function setUp() public {
        v1 = makeAddr("epoch-v1");
        v2 = makeAddr("epoch-v2");
        v3 = makeAddr("epoch-v3");
        s1 = vm.addr(0x201);
        s2 = vm.addr(0x202);
        s3 = vm.addr(0x203);

        address[] memory initVals = new address[](3);
        initVals[0] = v1;
        initVals[1] = v2;
        initVals[2] = v3;

        address[] memory initSigners = new address[](3);
        initSigners[0] = s1;
        initSigners[1] = s2;
        initSigners[2] = s3;

        deploySystem(initVals, initSigners, 10);
        vm.roll(5);
    }

    function _runEpochUpdate(address minerSigner) internal {
        vm.roll(10);
        vm.coinbase(minerSigner);
        address[] memory newSet = Validators(VALIDATORS).getTopValidators();
        vm.prank(minerSigner);
        Validators(VALIDATORS).updateActiveValidatorSet(newSet, 10);
    }

    function testSignerRotationTakesEffectNextEpoch() public {
        address newSigner = vm.addr(0x205);

        vm.prank(v1);
        Validators(VALIDATORS).createOrEditValidator(payable(v1), newSigner, "", "", "", "", "");

        assertEq(Validators(VALIDATORS).getValidatorSigner(v1), s1);
        assertEq(Validators(VALIDATORS).getValidatorBySigner(newSigner), address(0));
        assertEq(Validators(VALIDATORS).getValidatorBySignerHistory(newSigner), address(0));

        _runEpochUpdate(s1);
        assertEq(Validators(VALIDATORS).getValidatorSigner(v1), s1);
        assertEq(Validators(VALIDATORS).getValidatorBySigner(s1), v1);
        assertEq(Validators(VALIDATORS).getValidatorBySigner(newSigner), address(0));
        assertEq(Validators(VALIDATORS).getValidatorBySignerHistory(newSigner), v1);

        vm.coinbase(s1);
        vm.deal(s1, 1 ether);
        vm.prank(s1);
        Validators(VALIDATORS).distributeBlockReward{value: 1 ether}();

        (,, uint256 epochIncoming,,) = Validators(VALIDATORS).getValidatorInfo(v1);
        assertEq(epochIncoming, 1 ether);

        vm.roll(11);
        assertEq(Validators(VALIDATORS).getValidatorSigner(v1), newSigner);
        assertEq(Validators(VALIDATORS).getValidatorBySigner(newSigner), v1);
        assertEq(Validators(VALIDATORS).getValidatorBySigner(s1), address(0));

        vm.coinbase(newSigner);
        vm.deal(newSigner, 1 ether);
        vm.prank(newSigner);
        Validators(VALIDATORS).distributeBlockReward{value: 1 ether}();

        (,, uint256 incoming,,) = Validators(VALIDATORS).getValidatorInfo(v1);
        assertEq(incoming, 2 ether);
    }

    function testPendingSignerGettersExposeStoredRotationRecord() public {
        address newSigner = vm.addr(0x20A);

        vm.prank(v1);
        Validators(VALIDATORS).createOrEditValidator(payable(v1), newSigner, "", "", "", "", "");

        (address pendingSigner, uint256 effectiveBlock, bool pending) =
            Validators(VALIDATORS).getPendingValidatorSigner(v1);
        assertEq(pendingSigner, newSigner);
        assertEq(effectiveBlock, 10);
        assertTrue(pending);

        (address pendingValidator, uint256 reverseEffectiveBlock, bool reversePending) =
            Validators(VALIDATORS).getPendingValidatorBySigner(newSigner);
        assertEq(pendingValidator, v1);
        assertEq(reverseEffectiveBlock, 10);
        assertTrue(reversePending);
    }

    function testPendingSignerGettersRemainVisibleAfterDueUntilSyncAndThenClear() public {
        address newSigner = vm.addr(0x20B);

        vm.prank(v1);
        Validators(VALIDATORS).createOrEditValidator(payable(v1), newSigner, "", "", "", "", "");

        vm.roll(11);

        assertEq(Validators(VALIDATORS).getValidatorSigner(v1), newSigner);

        (address pendingSignerBefore, uint256 effectiveBlockBefore, bool pendingBefore) =
            Validators(VALIDATORS).getPendingValidatorSigner(v1);
        assertEq(pendingSignerBefore, newSigner);
        assertEq(effectiveBlockBefore, 10);
        assertTrue(pendingBefore);

        (address pendingValidatorBefore, uint256 reverseEffectiveBlockBefore, bool reversePendingBefore) =
            Validators(VALIDATORS).getPendingValidatorBySigner(newSigner);
        assertEq(pendingValidatorBefore, v1);
        assertEq(reverseEffectiveBlockBefore, 10);
        assertTrue(reversePendingBefore);

        vm.prank(v1);
        Validators(VALIDATORS).createOrEditValidator(payable(v1), "", "", "", "", "");

        (address pendingSignerAfter, uint256 effectiveBlockAfter, bool pendingAfter) =
            Validators(VALIDATORS).getPendingValidatorSigner(v1);
        assertEq(pendingSignerAfter, address(0));
        assertEq(effectiveBlockAfter, 0);
        assertFalse(pendingAfter);

        (address pendingValidatorAfter, uint256 reverseEffectiveBlockAfter, bool reversePendingAfter) =
            Validators(VALIDATORS).getPendingValidatorBySigner(newSigner);
        assertEq(pendingValidatorAfter, address(0));
        assertEq(reverseEffectiveBlockAfter, 0);
        assertFalse(reversePendingAfter);
    }

    function testEpochTransitionSignerQueryExposesNextSignerAtCheckpoint() public {
        address newSigner = vm.addr(0x207);

        vm.prank(v1);
        Validators(VALIDATORS).createOrEditValidator(payable(v1), newSigner, "", "", "", "", "");

        address[] memory runtimeBeforeCheckpoint = Validators(VALIDATORS).getTopSigners();
        address[] memory transitionBeforeCheckpoint = Validators(VALIDATORS).getTopSignersForEpochTransition();
        assertEq(runtimeBeforeCheckpoint[0], s1);
        assertEq(transitionBeforeCheckpoint[0], s1);
        assertEq(Validators(VALIDATORS).getValidatorBySignerHistory(newSigner), address(0));

        _runEpochUpdate(s1);

        address[] memory runtimeAtCheckpoint = Validators(VALIDATORS).getTopSigners();
        address[] memory transitionAtCheckpoint = Validators(VALIDATORS).getTopSignersForEpochTransition();
        assertEq(runtimeAtCheckpoint[0], s1);
        assertEq(transitionAtCheckpoint[0], newSigner);
        assertEq(Validators(VALIDATORS).getValidatorBySignerHistory(newSigner), v1);

        vm.roll(11);

        address[] memory runtimeAfterCheckpoint = Validators(VALIDATORS).getTopSigners();
        address[] memory transitionAfterCheckpoint = Validators(VALIDATORS).getTopSignersForEpochTransition();
        assertEq(runtimeAfterCheckpoint[0], newSigner);
        assertEq(transitionAfterCheckpoint[0], newSigner);
    }

    function testCandidateSignerDoesNotEnterHistoryBeforeActivationAndCanBeReusedAfterEdit() public {
        address validator = makeAddr("candidate-validator");
        address signer = vm.addr(0x208);
        address replacementSigner = vm.addr(0x209);
        address otherValidator = makeAddr("other-candidate-validator");
        uint256 minStake = Proposal(PROPOSAL).minValidatorStake();

        _passProposal(validator, true);
        _passProposal(otherValidator, true);

        vm.prank(validator);
        Validators(VALIDATORS).createOrEditValidator(payable(validator), signer, "", "", "", "", "");

        assertEq(Validators(VALIDATORS).getValidatorBySignerHistory(signer), address(0));

        vm.prank(validator);
        Validators(VALIDATORS).createOrEditValidator(payable(validator), replacementSigner, "", "", "", "", "");

        vm.prank(otherValidator);
        Validators(VALIDATORS).createOrEditValidator(payable(otherValidator), signer, "", "", "", "", "");

        vm.deal(otherValidator, minStake);
        vm.prank(otherValidator);
        Staking(STAKING).registerValidator{value: minStake}(1000);

        assertEq(Validators(VALIDATORS).getValidatorBySignerHistory(signer), address(0));

        _runEpochUpdate(s1);

        assertEq(Validators(VALIDATORS).getValidatorBySignerHistory(signer), otherValidator);
        assertEq(Validators(VALIDATORS).getValidatorBySignerHistory(replacementSigner), address(0));
    }

    function testResignClearsPendingSignerReservation() public {
        address rotatedSigner = vm.addr(0x205);

        vm.prank(v1);
        Validators(VALIDATORS).createOrEditValidator(payable(v1), rotatedSigner, "", "", "", "", "");

        (address pendingSignerBefore, uint256 effectiveBlockBefore, bool pendingBefore) =
            Validators(VALIDATORS).getPendingValidatorSigner(v1);
        assertEq(pendingSignerBefore, rotatedSigner);
        assertEq(effectiveBlockBefore, 10);
        assertTrue(pendingBefore);

        vm.prank(v1);
        Staking(STAKING).resignValidator();

        (address pendingSignerAfter, uint256 effectiveBlockAfter, bool pendingAfter) =
            Validators(VALIDATORS).getPendingValidatorSigner(v1);
        assertEq(pendingSignerAfter, address(0));
        assertEq(effectiveBlockAfter, 0);
        assertFalse(pendingAfter);

        vm.prank(v2);
        Validators(VALIDATORS).createOrEditValidator(payable(v2), rotatedSigner, "", "", "", "", "");

        vm.roll(10);
        assertEq(Validators(VALIDATORS).getValidatorSigner(v2), s2);
        assertEq(Validators(VALIDATORS).getValidatorBySigner(rotatedSigner), address(0));

        vm.roll(11);
        assertEq(Validators(VALIDATORS).getValidatorSigner(v2), rotatedSigner);
        assertEq(Validators(VALIDATORS).getValidatorBySigner(rotatedSigner), v2);
    }

    function testPunishRemovalClearsPendingSignerReservation() public {
        address rotatedSigner = vm.addr(0x206);

        vm.prank(v1);
        Validators(VALIDATORS).createOrEditValidator(payable(v1), rotatedSigner, "", "", "", "", "");

        (address pendingValidatorBefore, uint256 effectiveBlockBefore, bool pendingBefore) =
            Validators(VALIDATORS).getPendingValidatorBySigner(rotatedSigner);
        assertEq(pendingValidatorBefore, v1);
        assertEq(effectiveBlockBefore, 10);
        assertTrue(pendingBefore);

        vm.prank(PUNISH);
        Validators(VALIDATORS).removeValidator(v1);

        (address pendingValidatorAfter, uint256 effectiveBlockAfter, bool pendingAfter) =
            Validators(VALIDATORS).getPendingValidatorBySigner(rotatedSigner);
        assertEq(pendingValidatorAfter, address(0));
        assertEq(effectiveBlockAfter, 0);
        assertFalse(pendingAfter);

        vm.prank(v2);
        Validators(VALIDATORS).createOrEditValidator(payable(v2), rotatedSigner, "", "", "", "", "");

        vm.roll(10);
        assertEq(Validators(VALIDATORS).getValidatorSigner(v2), s2);
        assertEq(Validators(VALIDATORS).getValidatorBySigner(rotatedSigner), address(0));

        vm.roll(11);
        assertEq(Validators(VALIDATORS).getValidatorSigner(v2), rotatedSigner);
        assertEq(Validators(VALIDATORS).getValidatorBySigner(rotatedSigner), v2);
    }

    function _passProposal(address target, bool flag) internal {
        vm.warp(block.timestamp + 1_000_000);
        vm.roll(block.number + 101);

        vm.prank(v1);
        bytes32 id = Proposal(PROPOSAL).createProposal(target, flag, "");

        vm.prank(v1);
        Proposal(PROPOSAL).voteProposal(id, true);
        vm.prank(v2);
        Proposal(PROPOSAL).voteProposal(id, true);
    }
}
