// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {BaseSetup} from "../test/BaseSetup.t.sol";
import {Proposal} from "../contracts/Proposal.sol";

contract UpdateConfigScript is BaseSetup {
    function run() external {
        // 示例：更新配置参数
        uint256 cid = 1; // 配置ID
        uint256 newValue = 7200; // 新值
        updateConfig(cid, newValue);
    }
    
    function updateConfig(uint256 cid, uint256 newValue) public {
        Proposal(PROPOSAL).createUpdateConfigProposal(cid, newValue);
    }
}
