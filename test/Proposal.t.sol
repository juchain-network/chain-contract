// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Proposal} from "../contracts/Proposal.sol";

contract ProposalTest {
    Proposal proposal;

    function setUp() public {
        proposal = new Proposal();
        address[] memory vals = new address[](1);
        vals[0] = address(0xBEEF);
        proposal.initialize(vals, address(0xCAFE));
    }
}
