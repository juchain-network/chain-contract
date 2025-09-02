
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Test} from "lib/forge-std/src/Test.sol";
import {Validators} from "../contracts/Validators.sol";
import {Punish} from "../contracts/Punish.sol";
import {Proposal} from "../contracts/Proposal.sol";
import {Staking} from "../contracts/Staking.sol";

abstract contract BaseSetup is Test {
    address constant VALIDATORS = 0x000000000000000000000000000000000000f000;
    address constant PUNISH = 0x000000000000000000000000000000000000F001;
    address constant PROPOSAL = 0x000000000000000000000000000000000000F002;
    address constant STAKING = 0x000000000000000000000000000000000000F003;

    // Deploy runtime code of contracts to fixed addresses and initialize them
    function deploySystem(address[] memory initVals) internal {
        // place runtime code at fixed addresses (consistent with deployment)
        vm.etch(VALIDATORS, type(Validators).runtimeCode);
        vm.etch(PUNISH, type(Punish).runtimeCode);
        vm.etch(PROPOSAL, type(Proposal).runtimeCode);
        vm.etch(STAKING, type(Staking).runtimeCode);

        // initialize with provided validators
        Proposal(PROPOSAL).initialize(initVals);
        Validators(VALIDATORS).initialize(initVals);
        Punish(PUNISH).initialize();
        Staking(STAKING).initialize();
    }
}
