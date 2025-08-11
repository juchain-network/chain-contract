// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {BaseSetup} from "../forge-tests/BaseSetup.t.sol";
import {Proposal} from "../contracts/Proposal.sol";

contract RemoveNodeScript is BaseSetup {
    function run(address toRemove) external {
        Proposal p = Proposal(PRO);
        p.createProposal(toRemove, false, "");
    }
}
