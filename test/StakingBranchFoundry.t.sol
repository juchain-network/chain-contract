// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseSetup} from "./BaseSetup.t.sol";
import {Proposal} from "../contracts/Proposal.sol";
import {Staking} from "../contracts/Staking.sol";
import {Validators} from "../contracts/Validators.sol";
import {StdStorage, stdStorage} from "forge-std/StdStorage.sol";

contract StakingBranchFoundryTest is BaseSetup {
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

    function testInitializeRejectsWrongValidatorsContractAddress() public {
        Staking testStaking = new Staking();
        vm.expectRevert("Invalid validators contract address");
        testStaking.initialize(address(0x1234), PROPOSAL, PUNISH);
    }

    function testInitializeRejectsWrongProposalContractAddress() public {
        Staking testStaking = new Staking();
        vm.expectRevert("Invalid proposal contract address");
        testStaking.initialize(VALIDATORS, address(0x1234), PUNISH);
    }

    function testInitializeRejectsWrongPunishContractAddress() public {
        Staking testStaking = new Staking();
        vm.expectRevert("Invalid punish contract address");
        testStaking.initialize(VALIDATORS, PROPOSAL, address(0x1234));
    }

    function testInitializeWithValidatorsRejectsWrongValidatorsContractAddress() public {
        Staking testStaking = new Staking();
        address[] memory vals = new address[](1);
        vals[0] = v1;
        vm.expectRevert("Invalid validators contract address");
        testStaking.initializeWithValidators(address(0x1234), PROPOSAL, PUNISH, vals, 1000);
    }

    function testInitializeWithValidatorsRejectsWrongProposalContractAddress() public {
        Staking testStaking = new Staking();
        address[] memory vals = new address[](1);
        vals[0] = v1;
        vm.expectRevert("Invalid proposal contract address");
        testStaking.initializeWithValidators(VALIDATORS, address(0x1234), PUNISH, vals, 1000);
    }

    function testInitializeWithValidatorsRejectsWrongPunishContractAddress() public {
        Staking testStaking = new Staking();
        address[] memory vals = new address[](1);
        vals[0] = v1;
        vm.expectRevert("Invalid punish contract address");
        testStaking.initializeWithValidators(VALIDATORS, PROPOSAL, address(0x1234), vals, 1000);
    }

    function testInitializeWithValidatorsRejectsZeroCommissionRate() public {
        Staking testStaking = new Staking();
        address[] memory vals = new address[](1);
        vals[0] = v1;
        vm.expectRevert("Commission rate must be greater than 0");
        testStaking.initializeWithValidators(VALIDATORS, PROPOSAL, PUNISH, vals, 0);
    }

    function testInitializeWithValidatorsRejectsDuplicateValidator() public {
        Staking testStaking = new Staking();
        address[] memory vals = new address[](2);
        vals[0] = v1;
        vals[1] = v1;
        vm.deal(address(testStaking), Proposal(PROPOSAL).minValidatorStake() * 2);

        vm.expectRevert("Validator already exists");
        testStaking.initializeWithValidators(VALIDATORS, PROPOSAL, PUNISH, vals, 1000);
    }

    function testReinitializeV2StakingRejectsSecondCall() public {
        address miner = makeAddr("miner");
        vm.coinbase(miner);

        vm.prank(miner);
        Staking(STAKING).reinitializeV2();

        vm.prank(miner);
        vm.expectRevert("Already reinitialized");
        Staking(STAKING).reinitializeV2();
    }

    function testRegisterValidatorRejectsZeroCommissionRate() public {
        address candidate = makeAddr("candidate");
        uint256 minStake = Proposal(PROPOSAL).minValidatorStake();
        vm.deal(candidate, minStake);
        _passProposal(candidate);

        vm.prank(candidate);
        vm.expectRevert("Commission rate must be greater than 0");
        Staking(STAKING).registerValidator{value: minStake}(0);
    }

    function testRegisterValidatorRejectsAlreadyRegistered() public {
        uint256 minStake = Proposal(PROPOSAL).minValidatorStake();
        vm.deal(v1, minStake);

        vm.prank(v1);
        vm.expectRevert("Already registered");
        Staking(STAKING).registerValidator{value: minStake}(1000);
    }

    function testWithdrawUnbondedRejectsZeroMaxEntries() public {
        vm.expectRevert("maxEntries must be positive");
        Staking(STAKING).withdrawUnbonded(v1, 0);
    }

    function testWithdrawUnbondedRejectsTooLargeMaxEntries() public {
        vm.expectRevert("maxEntries too large");
        Staking(STAKING).withdrawUnbonded(v1, 65);
    }

    function testWithdrawUnbondedRejectsWhenNothingAvailable() public {
        vm.expectRevert("No unbonded tokens available");
        Staking(STAKING).withdrawUnbonded(v1, 1);
    }

    function testClaimRewardsRejectsWhenNoDelegationExists() public {
        vm.prank(v1);
        vm.expectRevert("No delegation found");
        Staking(STAKING).claimRewards(v2);
    }

    function testDistributeRewardsRejectsAlreadyDistributedInSameBlock() public {
        vm.coinbase(v1);
        vm.deal(v1, 1 ether);

        vm.prank(v1);
        Staking(STAKING).distributeRewards{value: 1 ether}();

        vm.deal(v1, 2 ether);
        vm.prank(v1);
        vm.expectRevert("Already distributed");
        Staking(STAKING).distributeRewards{value: 2 ether}();
    }

    function testJailValidatorRejectsInvalidArguments() public {
        vm.prank(PUNISH);
        vm.expectRevert("Invalid validator address");
        Staking(STAKING).jailValidator(address(0), 1);

        vm.prank(PUNISH);
        vm.expectRevert("Jail blocks must be positive");
        Staking(STAKING).jailValidator(v1, 0);
    }

    function testSlashValidatorRejectsInvalidArguments() public {
        vm.prank(PUNISH);
        vm.expectRevert("Invalid validator address");
        Staking(STAKING).slashValidator(address(0), 1, v1, 0, address(0xdead));

        vm.prank(PUNISH);
        vm.expectRevert("Slash amount must be positive");
        Staking(STAKING).slashValidator(v1, 0, v1, 0, address(0xdead));

        vm.prank(PUNISH);
        vm.expectRevert("Invalid burn address");
        Staking(STAKING).slashValidator(v1, 1, v1, 0, address(0));
    }

    function testUpdateLastActiveBlockRejectsZeroValidator() public {
        vm.prank(VALIDATORS);
        vm.expectRevert("Invalid validator address");
        Staking(STAKING).updateLastActiveBlock(address(0));
    }

    function testWithdrawPendingPayoutRejectsInvalidRecipient() public {
        stdstore.target(STAKING).sig("pendingPayouts(address)").with_key(v1).checked_write(1 ether);

        vm.prank(v1);
        vm.expectRevert("Invalid recipient");
        Staking(STAKING).withdrawPendingPayout(payable(address(0)));
    }

    function testWithdrawPendingPayoutRejectsWhenEmpty() public {
        vm.prank(v1);
        vm.expectRevert("No pending payout");
        Staking(STAKING).withdrawPendingPayout(payable(v2));
    }

    function testWithdrawPendingPayoutRejectsFailedTransfer() public {
        RejectingRecipient rejecting = new RejectingRecipient();
        stdstore.target(STAKING).sig("pendingPayouts(address)").with_key(v1).checked_write(1 ether);

        vm.prank(v1);
        vm.expectRevert("Transfer failed");
        Staking(STAKING).withdrawPendingPayout(payable(address(rejecting)));
    }

    function testUnjailValidatorRejectsInsufficientStake() public {
        address candidate = makeAddr("candidate_insufficient");
        uint256 minStake = Proposal(PROPOSAL).minValidatorStake();
        vm.deal(candidate, minStake);
        _passProposal(candidate);

        vm.prank(candidate);
        Staking(STAKING).registerValidator{value: minStake}(1000);

        vm.prank(PUNISH);
        Staking(STAKING).jailValidator(candidate, 1);

        vm.prank(PUNISH);
        Staking(STAKING).slashValidator(candidate, minStake, v1, 0, address(0xdead));

        vm.roll(Proposal(PROPOSAL).epoch() + 1);
        vm.prank(candidate);
        vm.expectRevert("Insufficient stake, must add stake first");
        Staking(STAKING).unjailValidator(candidate);
    }

    function testUnjailValidatorRejectsMissingReproposal() public {
        address candidate = makeAddr("candidate_reproposal");
        uint256 minStake = Proposal(PROPOSAL).minValidatorStake();
        vm.deal(candidate, minStake);
        _passProposal(candidate);

        vm.prank(candidate);
        Staking(STAKING).registerValidator{value: minStake}(1000);

        vm.prank(PUNISH);
        Staking(STAKING).jailValidator(candidate, 1);

        vm.prank(VALIDATORS);
        Proposal(PROPOSAL).setUnpassed(candidate);

        vm.roll(Proposal(PROPOSAL).epoch() + 1);
        vm.prank(candidate);
        vm.expectRevert("Must pass reproposal first");
        Staking(STAKING).unjailValidator(candidate);
    }

    function testUnjailValidatorRejectsActivationFailure() public {
        address candidate = makeAddr("candidate_activation_fail");
        uint256 minStake = Proposal(PROPOSAL).minValidatorStake();
        vm.deal(candidate, minStake);
        _passProposal(candidate);

        vm.prank(candidate);
        Staking(STAKING).registerValidator{value: minStake}(1000);

        vm.prank(PUNISH);
        Staking(STAKING).jailValidator(candidate, 1);

        vm.roll(Proposal(PROPOSAL).epoch() + 1);
        vm.mockCall(VALIDATORS, abi.encodeWithSelector(Validators.tryActive.selector, candidate), abi.encode(false));

        vm.prank(candidate);
        vm.expectRevert("Failed to activate validator");
        Staking(STAKING).unjailValidator(candidate);
    }

    function testSlashValidatorReturnsZeroForEmptyStake() public {
        address candidate = makeAddr("candidate_empty");

        vm.prank(PUNISH);
        (uint256 actualSlash, uint256 actualReward) =
            Staking(STAKING).slashValidator(candidate, 1 ether, v1, 0, address(0xdead));

        assertEq(actualSlash, 0);
        assertEq(actualReward, 0);
    }

    function testDistributeRewardsReturnsWhenSignerUnmapped() public {
        address stranger = makeAddr("stranger_signer");
        vm.coinbase(stranger);
        vm.deal(stranger, 1 ether);

        vm.prank(stranger);
        Staking(STAKING).distributeRewards{value: 1 ether}();
    }

    function testDistributeRewardsReturnsWhenValidatorJailed() public {
        vm.prank(PUNISH);
        Staking(STAKING).jailValidator(v1, 10);

        vm.coinbase(v1);
        vm.deal(v1, 1 ether);
        vm.prank(v1);
        Staking(STAKING).distributeRewards{value: 1 ether}();
    }

    function testDistributeRewardsReturnsWhenValidatorStakeBelowMinimum() public {
        uint256 minStake = Proposal(PROPOSAL).minValidatorStake();
        address candidate = makeAddr("candidate_below_min");
        vm.deal(candidate, minStake);
        _passProposal(candidate);

        vm.prank(candidate);
        Staking(STAKING).registerValidator{value: minStake}(1000);

        vm.prank(PUNISH);
        Staking(STAKING).slashValidator(candidate, 1, v1, 0, address(0xdead));

        vm.coinbase(candidate);
        vm.deal(candidate, 1 ether);
        vm.prank(candidate);
        Staking(STAKING).distributeRewards{value: 1 ether}();
    }

    function testClaimValidatorRewardsRejectsUnregisteredValidator() public {
        address stranger = makeAddr("stranger_validator");
        vm.prank(stranger);
        vm.expectRevert("Not a registered validator");
        Staking(STAKING).claimValidatorRewards();
    }

    function testClaimValidatorRewardsNoOpWhenCommissionIsZero() public {
        vm.prank(v1);
        Staking(STAKING).claimValidatorRewards();

        (,,,,,, uint256 totalClaimedRewards, uint256 lastClaimBlock,,) = Staking(STAKING).getValidatorInfo(v1);
        assertEq(totalClaimedRewards, 0);
        assertEq(lastClaimBlock, 0);
    }

    function testDistributeRewardsAllocatesAllToValidatorWhenNoDelegators() public {
        vm.coinbase(v1);
        vm.deal(v1, 10 ether);

        vm.prank(v1);
        Staking(STAKING).distributeRewards{value: 10 ether}();

        (,, uint256 commissionRate, uint256 accumulatedRewards,,,,,, uint256 totalRewards) =
            Staking(STAKING).getValidatorInfo(v1);
        uint256 expectedCommission = 10 ether * commissionRate / 10000;

        assertEq(totalRewards, 10 ether);
        assertEq(accumulatedRewards, 10 ether);
        assertGt(expectedCommission, 0);
    }

    function _passProposal(address candidate) internal {
        Proposal p = Proposal(PROPOSAL);
        vm.prank(v1);
        bytes32 id = p.createProposal(candidate, true, "");
        vm.prank(v1);
        p.voteProposal(id, true);
        vm.prank(v2);
        p.voteProposal(id, true);
        vm.prank(v3);
        p.voteProposal(id, true);
    }
}

contract RejectingRecipient {
    receive() external payable {
        revert("reject");
    }
}
