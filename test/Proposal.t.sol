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

    // 注意: receiverAddr 和 increasePeriod 已移除，系统不再支持代币增发
    // 以下测试已移除，因为相关功能已删除
    // function testReceiverAddrInitializedNonZero() public view {
    //     require(proposal.receiverAddr() != address(0), "receiverAddr should be non-zero");
    // }
}
