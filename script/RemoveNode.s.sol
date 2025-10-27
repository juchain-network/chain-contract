// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {BaseSetup} from "../test/BaseSetup.t.sol";
import {Proposal} from "../contracts/Proposal.sol";

contract RemoveNodeScript is BaseSetup {
    function run() external {
        // 示例：移除一个验证者节点
        address toRemove = 0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65;
        removeNode(toRemove);
    }
    
    function removeNode(address toRemove) public {
        Proposal p = Proposal(PROPOSAL);
        p.createProposal(toRemove, false, "Removing validator node");
    }
}
