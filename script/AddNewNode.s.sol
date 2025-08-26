// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {BaseSetup} from "../test/BaseSetup.t.sol";
import {Proposal} from "../contracts/Proposal.sol";
import {Validators} from "../contracts/Validators.sol";

// Minimal script without forge-std dependency; uses BaseSetup helpers for addresses only.
contract AddNewNodeScript is BaseSetup {
    function run() external {
        // 示例：添加一个新的验证者节点
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
