// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseSetup} from "./BaseSetup.t.sol";
import {Proposal} from "../contracts/Proposal.sol";
import {Staking} from "../contracts/Staking.sol";
import {Validators} from "../contracts/Validators.sol";
import {Punish} from "../contracts/Punish.sol";

contract ReinitializeV2Test is BaseSetup {
    address miner;

    function setUp() public {
        address[] memory initVals = new address[](3);
        initVals[0] = makeAddr("v1");
        initVals[1] = makeAddr("v2");
        initVals[2] = makeAddr("v3");
        deploySystem(initVals);
        miner = makeAddr("miner");
        vm.coinbase(miner);
    }

    function testReinitializeV2Proposal() public {
        Proposal p = Proposal(PROPOSAL);
        vm.prank(miner);
        p.reinitializeV2();
        assertEq(p.revision(), 2);
    }

    function testReinitializeV2Staking() public {
        Staking s = Staking(STAKING);
        vm.prank(miner);
        s.reinitializeV2();
        assertEq(s.revision(), 2);
    }

    function testReinitializeV2Validators() public {
        Validators v = Validators(VALIDATORS);
        vm.prank(miner);
        v.reinitializeV2();
        assertEq(v.revision(), 2);
    }

    function testReinitializeV2Punish() public {
        Punish p = Punish(PUNISH);
        vm.prank(miner);
        p.reinitializeV2();
        assertEq(p.revision(), 2);
    }
}
