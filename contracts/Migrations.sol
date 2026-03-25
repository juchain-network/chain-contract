// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

contract Migrations {
    address public owner = msg.sender;
    uint256 public lastCompletedMigration;

    modifier restricted() {
        _restricted();
        _;
    }

    function _restricted() internal view {
        require(msg.sender == owner, "This function is restricted to the contract's owner");
    }

    function setCompleted(uint256 completed) public restricted {
        lastCompletedMigration = completed;
    }
}
