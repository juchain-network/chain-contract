// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseSetup} from "../test/BaseSetup.t.sol";
import {Proposal} from "../contracts/Proposal.sol";

// Includes address validation and returns proposal ID
contract CreateProposalScript is BaseSetup {
    
    event ProposalCreated(bytes32 indexed id, address indexed proposer, address indexed dst, bool flag);
    
    function run() external {
        // Run in test mode, first deploy the system
        address[] memory initialValidators = new address[](3);
        initialValidators[0] = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;
        initialValidators[1] = 0x70997970C51812dc3A010C7d01b50e0d17dc79C8;
        initialValidators[2] = 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC;
        deploySystem(initialValidators);
        
        // Example: Create a proposal to add a validator
        address target = 0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc;
        bool isAdd = true;
        string memory details = "Add new validator for network expansion";
        
        createProposal(target, isAdd, details);
    }
    
    function createProposal(address target, bool isAdd, string memory details) public returns (bytes32) {
        // Validate address format (corresponds to regex check in JS version)
        require(target != address(0), "Invalid target address");
        require(bytes(details).length <= 3000, "Details too long");
        
        // Calculate proposal ID (needs to be consistent with contract logic)
        bytes32 id = keccak256(abi.encodePacked(msg.sender, target, isAdd, details, block.timestamp));
        
        Proposal(PROPOSAL).createProposal(target, isAdd, details);
        
        emit ProposalCreated(id, msg.sender, target, isAdd);
        return id;
    }
    
    // Convenience function: Add validator
    function addValidator(address validator, string memory details) external returns (bytes32) {
        return createProposal(validator, true, details);
    }
    
    //// Convenience function: Remove validator
    function removeValidator(address validator, string memory details) external returns (bytes32) {
        return createProposal(validator, false, details);
    }
}
