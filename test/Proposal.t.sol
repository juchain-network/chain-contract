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

    function testReceiverAddrInitializedNonZero() public view {
        // 初始化后，receiverAddr 应该为非零地址（见合约 initialize 默认值）
        require(proposal.receiverAddr() != address(0), "receiverAddr should be non-zero");
    }
}
