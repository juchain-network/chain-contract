// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseSetup} from "./BaseSetup.t.sol";
import {Punish} from "../contracts/Punish.sol";
import {Proposal} from "../contracts/Proposal.sol";
import {Validators} from "../contracts/Validators.sol";

contract PunishFoundryTest is BaseSetup {

    address miner;
    address v1; address v2; address v3;

    function setUp() public {
        v1 = makeAddr("v1");
        v2 = makeAddr("v2");
        v3 = makeAddr("v3");
        address[] memory initVals = new address[](3);
        initVals[0] = v1; initVals[1] = v2; initVals[2] = v3;
        deploySystem(initVals);
        miner = v1; // simulate coinbase
        vm.coinbase(miner);
    }

    function testPunishThresholdRemovesIncoming() public {
        // read threshold
        uint256 thr = Proposal(PROPOSAL).punishThreshold();
        for (uint256 i = 0; i < thr; i++) {
            vm.coinbase(miner);
            vm.prank(miner);
            Punish(PUNISH).punish(miner);
            vm.roll(block.number + 1);
        }
        (, , uint256 aac,,) = Validators(VALIDATORS).getValidatorInfo(miner);
        require(aac == 0, "incoming removed at punish threshold");
    }

    function testRemoveThresholdJails() public {
        uint256 thr = Proposal(PROPOSAL).removeThreshold();
        for (uint256 i = 0; i < thr; i++) {
            vm.coinbase(miner);
            vm.prank(miner);
            Punish(PUNISH).punish(miner);
            vm.roll(block.number + 1);
        }
        (, Validators.Status status,,,) = Validators(VALIDATORS).getValidatorInfo(miner);
        require(uint256(status) == uint256(Validators.Status.Jailed), "jailed at remove threshold");
    }

    function testDecreaseMissedBlocksCounter() public {
        uint256 thr = Proposal(PROPOSAL).removeThreshold();
        uint256 dec = Proposal(PROPOSAL).decreaseRate();
        // accumulate some missed blocks on v2 and v3
        for (uint256 i = 0; i < thr / dec; i++) {
            vm.coinbase(miner);
            vm.prank(miner);
            Punish(PUNISH).punish(v2);
            vm.roll(block.number + 1);
        }
        vm.coinbase(miner);
        vm.prank(miner);
        Punish(PUNISH).punish(v3);
        vm.roll(block.number + 1);

        uint256 epoch = Punish(PUNISH).epoch();
        uint256 blocksToNext = epoch - (block.number % epoch);
        vm.roll(block.number + blocksToNext);
        vm.prank(miner);
        Punish(PUNISH).decreaseMissedBlocksCounter(epoch);

    uint256 r2 = Punish(PUNISH).getPunishRecord(v2);
    uint256 r3 = Punish(PUNISH).getPunishRecord(v3);
    require(r2 == 0, "v2 decreased to zero or reset");
    require(r3 == 0, "v3 decreased to zero or reset");
    }
}
