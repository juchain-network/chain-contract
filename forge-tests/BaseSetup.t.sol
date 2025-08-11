// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Validators} from "../contracts/Validators.sol";
import {Punish} from "../contracts/Punish.sol";
import {Proposal} from "../contracts/Proposal.sol";

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

    // Fixed system addresses
    address constant VAL = 0x000000000000000000000000000000000000f000;
    address constant PUN = 0x000000000000000000000000000000000000F001;
    address constant PRO = 0x000000000000000000000000000000000000F002;

    // Deploy runtime code of contracts to fixed addresses and initialize them
    function deploySystem(address[] memory initVals) internal {
        // place runtime code at fixed addresses
        vm.etch(VAL, type(Validators).runtimeCode);
        vm.etch(PUN, type(Punish).runtimeCode);
        vm.etch(PRO, type(Proposal).runtimeCode);

        // initialize with provided validators
        Proposal(PRO).initialize(initVals);
        Validators(VAL).initialize(initVals);
        Punish(PUN).initialize();
    }

    // helper to make deterministic pseudo addresses and fund them
    function makeAddr(string memory salt) internal returns (address a) {
        a = address(uint160(uint256(keccak256(abi.encodePacked(salt)))));
        vm.deal(a, 100 ether);
    }
}
