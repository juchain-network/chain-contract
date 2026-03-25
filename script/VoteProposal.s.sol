// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseSetup} from "../test/BaseSetup.t.sol";
import {Proposal} from "../contracts/Proposal.sol";
import {Validators} from "../contracts/Validators.sol";

contract VoteProposalScript is BaseSetup {
    function run() external {
        // Run in test mode, first deploy the system
        address[] memory initialValidators = new address[](3);
        initialValidators[0] = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;
        initialValidators[1] = 0x70997970C51812dc3A010C7d01b50e0d17dc79C8;
        initialValidators[2] = 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC;
        deploySystem(initialValidators);

        // Example: Vote yes on the first proposal
        // Note: Only active validators can vote, this script is for demonstration only
        bytes32 sampleProposalId = 0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef;

        // Check if current account is a validator
        bool isValidator = Validators(VALIDATORS).isActiveValidator(msg.sender);
        if (!isValidator) {
            // For demonstration only, actual use requires validator account
            return;
        }

        voteOnProposal(sampleProposalId, true);
    }

    function voteOnProposal(bytes32 proposalId, bool vote) public {
        // Validate proposal ID format (there's regex validation in the JS version)
        require(proposalId != bytes32(0), "Invalid proposal id");

        Proposal(PROPOSAL).voteProposal(proposalId, vote);
    }

    // Convenience function: Vote yes
    function voteYes(bytes32 proposalId) external {
        Proposal(PROPOSAL).voteProposal(proposalId, true);
    }

    // Convenience function: Vote no
    function voteNo(bytes32 proposalId) external {
        Proposal(PROPOSAL).voteProposal(proposalId, false);
    }
}
