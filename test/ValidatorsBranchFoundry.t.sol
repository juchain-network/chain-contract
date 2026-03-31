// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseSetup} from "./BaseSetup.t.sol";
import {Proposal} from "../contracts/Proposal.sol";
import {Validators} from "../contracts/Validators.sol";
import {Staking} from "../contracts/Staking.sol";

contract ValidatorsBranchFoundryTest is BaseSetup {
    address v1;
    address v2;
    address v3;
    address s1;
    address s2;
    address s3;

    function setUp() public {
        v1 = makeAddr("validators-branch-v1");
        v2 = makeAddr("validators-branch-v2");
        v3 = makeAddr("validators-branch-v3");
        s1 = vm.addr(0x301);
        s2 = vm.addr(0x302);
        s3 = vm.addr(0x303);

        address[] memory initVals = new address[](3);
        initVals[0] = v1;
        initVals[1] = v2;
        initVals[2] = v3;

        address[] memory initSigners = new address[](3);
        initSigners[0] = s1;
        initSigners[1] = s2;
        initSigners[2] = s3;

        deploySystem(initVals, initSigners, 10);
    }

    function testInitializeRejectsLengthMismatch() public {
        Validators testValidators = new Validators();
        address[] memory vals = new address[](1);
        vals[0] = v1;
        address[] memory signers = new address[](0);

        vm.expectRevert("Length mismatch");
        testValidators.initialize(vals, signers, PROPOSAL, PUNISH, STAKING);
    }

    function testInitializeRejectsWrongProposalContractAddress() public {
        Validators testValidators = new Validators();
        address[] memory vals = new address[](1);
        vals[0] = v1;
        address[] memory signers = new address[](1);
        signers[0] = s1;

        vm.expectRevert("Invalid proposal contract address");
        testValidators.initialize(vals, signers, address(0x1234), PUNISH, STAKING);
    }

    function testInitializeRejectsWrongPunishContractAddress() public {
        Validators testValidators = new Validators();
        address[] memory vals = new address[](1);
        vals[0] = v1;
        address[] memory signers = new address[](1);
        signers[0] = s1;

        vm.expectRevert("Invalid punish contract address");
        testValidators.initialize(vals, signers, PROPOSAL, address(0x1234), STAKING);
    }

    function testInitializeRejectsWrongStakingContractAddress() public {
        Validators testValidators = new Validators();
        address[] memory vals = new address[](1);
        vals[0] = v1;
        address[] memory signers = new address[](1);
        signers[0] = s1;

        vm.expectRevert("Invalid staking contract address");
        testValidators.initialize(vals, signers, PROPOSAL, PUNISH, address(0x1234));
    }

    function testInitializeRejectsZeroValidatorAddress() public {
        Validators testValidators = new Validators();
        address[] memory vals = new address[](1);
        vals[0] = address(0);
        address[] memory signers = new address[](1);
        signers[0] = s1;

        vm.expectRevert("Invalid validator address");
        testValidators.initialize(vals, signers, PROPOSAL, PUNISH, STAKING);
    }

    function testReinitializeV2ValidatorsRejectsSecondCall() public {
        address miner = makeAddr("validators-branch-miner");
        vm.coinbase(miner);

        vm.prank(miner);
        Validators(VALIDATORS).reinitializeV2();

        vm.prank(miner);
        vm.expectRevert("Already reinitialized");
        Validators(VALIDATORS).reinitializeV2();
    }

    function testDistributeBlockRewardRejectsSameBlockRepeat() public {
        vm.coinbase(s1);
        vm.deal(s1, 1 ether);

        vm.prank(s1);
        Validators(VALIDATORS).distributeBlockReward{value: 1 ether}();

        vm.prank(s1);
        (bool ok,) = address(Validators(VALIDATORS)).call{value: 1 ether}(
            abi.encodeWithSelector(Validators.distributeBlockReward.selector)
        );
        require(!ok, "second reward in same block should fail");
    }

    function testDistributeBlockRewardReturnsWhenSignerUnmapped() public {
        address stranger = makeAddr("validators-branch-stranger");
        vm.coinbase(stranger);
        vm.deal(stranger, 1 ether);

        vm.prank(stranger);
        Validators(VALIDATORS).distributeBlockReward{value: 1 ether}();
    }

    function testDistributeBlockRewardReturnsWhenMappedValidatorNotRegistered() public {
        address candidate = makeAddr("validators-branch-unregistered");
        address candidateSigner = vm.addr(0x398);
        _passProposal(candidate);

        vm.prank(candidate);
        Validators(VALIDATORS).createOrEditValidator(payable(candidate), candidateSigner, "", "", "", "", "");

        vm.coinbase(candidateSigner);
        vm.deal(candidateSigner, 1 ether);
        vm.prank(candidateSigner);
        Validators(VALIDATORS).distributeBlockReward{value: 1 ether}();
    }

    function testUpdateActiveValidatorSetReturnsWhenRepeatedInSameBlock() public {
        uint256 epoch = Validators(VALIDATORS).epoch();
        address[] memory newSet = Validators(VALIDATORS).getTopValidators();
        uint256 targetBlock = ((block.number / epoch) + 1) * epoch;
        vm.roll(targetBlock);

        vm.coinbase(s1);
        vm.prank(s1);
        Validators(VALIDATORS).updateActiveValidatorSet(newSet, epoch);

        vm.coinbase(s1);
        vm.prank(s1);
        Validators(VALIDATORS).updateActiveValidatorSet(newSet, epoch);
    }

    function testCreateOrEditValidatorRejectsSignerAlreadyAssigned() public {
        address candidate = makeAddr("validators-branch-candidate");
        _passProposal(candidate);

        vm.prank(candidate);
        vm.expectRevert("Signer already used");
        Validators(VALIDATORS).createOrEditValidator(payable(candidate), s1, "", "", "", "", "");
    }

    function testCreateOrEditValidatorKeepsPendingSignerWhenZeroSignerProvided() public {
        address candidate = makeAddr("validators-branch-zero-signer");
        _passProposal(candidate);
        uint256 minStake = Proposal(PROPOSAL).minValidatorStake();
        vm.deal(candidate, minStake);

        vm.prank(candidate);
        Staking(STAKING).registerValidator{value: minStake}(1000);

        address nextSigner = vm.addr(0x399);
        vm.prank(candidate);
        Validators(VALIDATORS).createOrEditValidator(payable(candidate), nextSigner, "", "", "", "", "");

        (address pendingSigner,, bool exists) = Validators(VALIDATORS).getPendingValidatorSigner(candidate);
        assertTrue(exists);
        assertEq(pendingSigner, nextSigner);

        vm.prank(candidate);
        Validators(VALIDATORS).createOrEditValidator(payable(candidate), address(0), "", "", "", "", "");

        (pendingSigner,, exists) = Validators(VALIDATORS).getPendingValidatorSigner(candidate);
        assertTrue(exists);
        assertEq(pendingSigner, nextSigner);
    }

    function testCreateOrEditValidatorClearsPendingSignerWhenResetToCurrentSigner() public {
        address candidate = makeAddr("validators-branch-clear-pending");
        _passProposal(candidate);
        uint256 minStake = Proposal(PROPOSAL).minValidatorStake();
        vm.deal(candidate, minStake);

        vm.prank(candidate);
        Staking(STAKING).registerValidator{value: minStake}(1000);

        address originalSigner = Validators(VALIDATORS).getValidatorSigner(candidate);
        address nextSigner = vm.addr(0x39A);

        vm.prank(candidate);
        Validators(VALIDATORS).createOrEditValidator(payable(candidate), nextSigner, "", "", "", "", "");

        (address pendingSigner,, bool pending) = Validators(VALIDATORS).getPendingValidatorSigner(candidate);
        assertTrue(pending);
        assertEq(pendingSigner, nextSigner);

        vm.prank(candidate);
        Validators(VALIDATORS).createOrEditValidator(payable(candidate), originalSigner, "", "", "", "", "");

        (pendingSigner,, pending) = Validators(VALIDATORS).getPendingValidatorSigner(candidate);
        assertFalse(pending);
        assertEq(pendingSigner, address(0));
    }

    function testCreateOrEditValidatorReplacesOldPendingSignerReservation() public {
        address candidate = makeAddr("validators-branch-replace-pending");
        _passProposal(candidate);
        uint256 minStake = Proposal(PROPOSAL).minValidatorStake();
        vm.deal(candidate, minStake);

        vm.prank(candidate);
        Staking(STAKING).registerValidator{value: minStake}(1000);

        address signerA = vm.addr(0x39B);
        address signerB = vm.addr(0x39C);

        vm.prank(candidate);
        Validators(VALIDATORS).createOrEditValidator(payable(candidate), signerA, "", "", "", "", "");

        vm.prank(candidate);
        Validators(VALIDATORS).createOrEditValidator(payable(candidate), signerB, "", "", "", "", "");

        (address validatorA,, bool pendingA) = Validators(VALIDATORS).getPendingValidatorBySigner(signerA);
        assertEq(validatorA, address(0));
        assertFalse(pendingA);

        (address validatorB,, bool pendingB) = Validators(VALIDATORS).getPendingValidatorBySigner(signerB);
        assertEq(validatorB, candidate);
        assertTrue(pendingB);
    }

    function testValidateDescriptionRejectsIdentityTooLong() public {
        string memory tooLongIdentity = _generateString(3001);
        vm.expectRevert("Invalid identity length");
        Validators(VALIDATORS).validateDescription("", tooLongIdentity, "", "", "");
    }

    function testValidateDescriptionRejectsWebsiteTooLong() public {
        string memory tooLongWebsite = _generateString(141);
        vm.expectRevert("Invalid website length");
        Validators(VALIDATORS).validateDescription("", "", tooLongWebsite, "", "");
    }

    function testValidateDescriptionRejectsEmailTooLong() public {
        string memory tooLongEmail = _generateString(141);
        vm.expectRevert("Invalid email length");
        Validators(VALIDATORS).validateDescription("", "", "", tooLongEmail, "");
    }

    function testValidateDescriptionRejectsDetailsTooLong() public {
        string memory tooLongDetails = _generateString(281);
        vm.expectRevert("Invalid details length");
        Validators(VALIDATORS).validateDescription("", "", "", "", tooLongDetails);
    }

    function testIsActiveValidatorReturnsFalseForUnknownAddress() public {
        assertFalse(Validators(VALIDATORS).isActiveValidator(makeAddr("validators-branch-not-active")));
    }

    function testIsTopValidatorReturnsFalseForUnknownAddress() public {
        assertFalse(Validators(VALIDATORS).isTopValidator(makeAddr("validators-branch-not-top")));
    }

    function testDistributeBlockRewardZeroValueWithAllValidatorsIneligible() public {
        vm.prank(PUNISH);
        Staking(STAKING).jailValidator(v1, 10);
        vm.prank(PUNISH);
        Staking(STAKING).jailValidator(v2, 10);
        vm.prank(PUNISH);
        Staking(STAKING).jailValidator(v3, 10);

        vm.coinbase(s1);
        vm.prank(s1);
        Validators(VALIDATORS).distributeBlockReward();
    }

    function testRemoveFromHighestSetBreaksWhenValidatorFound() public {
        vm.prank(STAKING);
        Validators(VALIDATORS).removeFromHighestSet(v1);

        assertFalse(Validators(VALIDATORS).isTopValidator(v1));
    }

    function testStandaloneInitializeRecordsHistoricalSignerFromCurrentBlock() public {
        Validators fresh = new Validators();
        address[] memory vals = new address[](1);
        vals[0] = makeAddr("validators-branch-fresh");
        address[] memory signers = new address[](1);
        signers[0] = vm.addr(0x39D);

        vm.roll(5);
        fresh.initialize(vals, signers, PROPOSAL, PUNISH, STAKING);

        assertEq(fresh.getValidatorBySignerHistoryAt(signers[0], 4), address(0));
        assertEq(fresh.getValidatorBySignerHistoryAt(signers[0], 5), vals[0]);
    }

    function _passProposal(address candidate) internal {
        vm.warp(block.timestamp + 1_000_000);
        vm.roll(block.number + 101);

        vm.prank(v1);
        bytes32 id = Proposal(PROPOSAL).createProposal(candidate, true, "");

        vm.prank(v1);
        Proposal(PROPOSAL).voteProposal(id, true);
        vm.prank(v2);
        Proposal(PROPOSAL).voteProposal(id, true);
    }

    function _generateString(uint256 length) internal pure returns (string memory) {
        bytes memory result = new bytes(length);
        for (uint256 i = 0; i < length; i++) {
            result[i] = "a";
        }
        return string(result);
    }
}
