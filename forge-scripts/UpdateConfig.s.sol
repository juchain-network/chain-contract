// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {BaseSetup} from "../forge-tests/BaseSetup.t.sol";
import {Proposal} from "../contracts/Proposal.sol";

contract UpdateConfigScript is BaseSetup {
    function run(uint256 cid, uint256 newValue) external {
        Proposal(PRO).createUpdateConfigProposal(cid, newValue);
    }
}
