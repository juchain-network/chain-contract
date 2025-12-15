// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {BaseSetup} from "../test/BaseSetup.t.sol";
import {Proposal} from "../contracts/Proposal.sol";

// Minimal script without forge-std dependency; uses BaseSetup helpers for addresses only.
contract AddNewNodeScript is BaseSetup {
    function run() external {
        // Run in test mode, first deploy the system
        address[] memory initialValidators = new address[](3);
        initialValidators[0] = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;
        initialValidators[1] = 0x70997970C51812dc3A010C7d01b50e0d17dc79C8;
        initialValidators[2] = 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC;
        deploySystem(initialValidators);
        
        // Example: Add a new validator node
        address toAdd = 0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc;
        addNewNode(toAdd);
    }
    
    function addNewNode(address toAdd) public {
        // assumes validators already initialized and msg.sender is a validator
        Proposal p = Proposal(PROPOSAL);
        // create proposal
        p.createProposal(toAdd, true, "Adding new validator node");
        // cannot compute id robustly here without freezing timestamp; instead, read pass after external votes
        // This script is a thin wrapper suitable for broadcast as a validator to create the proposal only.
    }
}
