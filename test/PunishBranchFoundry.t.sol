// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseSetup} from "./BaseSetup.t.sol";
import {Punish} from "../contracts/Punish.sol";
import {Proposal} from "../contracts/Proposal.sol";
import {Validators} from "../contracts/Validators.sol";
import {Staking} from "../contracts/Staking.sol";

contract PunishBranchFoundryTest is BaseSetup {
    address internal miner;
    address internal v1;
    address internal v2;
    address internal v3;

    function setUp() public {
        v1 = makeAddr("punish-branch-v1");
        v2 = makeAddr("punish-branch-v2");
        v3 = makeAddr("punish-branch-v3");
        address[] memory initVals = new address[](3);
        initVals[0] = v1;
        initVals[1] = v2;
        initVals[2] = v3;
        deploySystem(initVals, initVals, 10);
        miner = v1;
        vm.coinbase(miner);
    }

    function testPunishAcceptsValidatorAddressDirectly() public {
        vm.prank(miner);
        Punish(PUNISH).punish(v2);

        assertEq(Punish(PUNISH).getPunishRecord(v2), 1);
    }

    function testPunishUnknownValidatorStillResolvesToZeroBeforeRevert() public {
        vm.prank(miner);
        vm.expectRevert("Validator not exist");
        Punish(PUNISH).punish(makeAddr("punish-branch-unknown"));
    }

    function testExecutePendingRejectsZeroLimit() public {
        vm.prank(miner);
        vm.expectRevert("Limit must be positive");
        Punish(PUNISH).executePending(0);
    }

    function testExecutePendingRejectsTooLargeLimit() public {
        vm.prank(miner);
        vm.expectRevert("Limit too large");
        Punish(PUNISH).executePending(6);
    }

    function testPunishRejectsRepeatedCallInSameBlock() public {
        vm.prank(miner);
        Punish(PUNISH).punish(v2);

        vm.prank(miner);
        vm.expectRevert("Already punished");
        Punish(PUNISH).punish(v3);
    }

    function testDecreaseMissedBlocksCounterRejectsRepeatedCallInSameBlock() public {
        uint256 epoch = Proposal(PROPOSAL).epoch();
        vm.roll(epoch);

        vm.prank(miner);
        Punish(PUNISH).decreaseMissedBlocksCounter(epoch);

        vm.prank(miner);
        vm.expectRevert("Already decreased");
        Punish(PUNISH).decreaseMissedBlocksCounter(epoch);
    }

    function testPunishReturnsWhenValidatorNotActive() public {
        address candidate = makeAddr("punish-branch-inactive-candidate");
        uint256 minStake = Proposal(PROPOSAL).minValidatorStake();
        vm.deal(candidate, minStake);

        vm.warp(block.timestamp + 1_000_000);
        vm.roll(block.number + 101);
        vm.prank(v1);
        bytes32 id = Proposal(PROPOSAL).createProposal(candidate, true, "");
        vm.prank(v1);
        Proposal(PROPOSAL).voteProposal(id, true);
        vm.prank(v2);
        Proposal(PROPOSAL).voteProposal(id, true);

        vm.prank(candidate);
        Staking(STAKING).registerValidator{value: minStake}(1000);

        vm.prank(miner);
        Punish(PUNISH).punish(candidate);

        assertEq(Punish(PUNISH).getPunishRecord(candidate), 0);
    }
}
