
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Validators} from "../contracts/Validators.sol";
import {Punish} from "../contracts/Punish.sol";
import {Proposal} from "../contracts/Proposal.sol";
import {Staking} from "../contracts/Staking.sol";

/// Minimal VM cheatcode interface (subset) to avoid external deps
interface Vm {
    function prank(address) external;
    function startPrank(address) external;
    function stopPrank() external;
    function warp(uint256) external;
    function roll(uint256) external;
    function deal(address, uint256) external;
    function coinbase(address) external;
    function etch(address, bytes calldata) external;
}

abstract contract BaseSetup {
    // hevm cheatcode address
    Vm constant vm = Vm(address(uint160(uint256(keccak256("hevm cheat code")))));

    // Fixed system addresses (consistent with deployment)
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

    // helper to make deterministic pseudo addresses and fund them
    function makeAddr(string memory salt) internal returns (address a) {
        a = address(uint160(uint256(keccak256(abi.encodePacked(salt)))));
        vm.deal(a, 100 ether);
    }
}
