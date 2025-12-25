// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {BaseSetup} from "../test/BaseSetup.t.sol";
import {Proposal} from "../contracts/Proposal.sol";

contract UpdateConfigScript is BaseSetup {
    function run() external {
        // Example: Update configuration parameters
        uint256 cid = 1; // Configuration ID
        uint256 newValue = 7200; // New value
        updateConfig(cid, newValue);
    }
    
    function updateConfig(uint256 cid, uint256 newValue) public {
        Proposal(PROPOSAL).createUpdateConfigProposal(cid, newValue);
    }
}
