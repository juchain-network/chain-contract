// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {BaseSetup} from "./BaseSetup.t.sol";
import {Proposal} from "../contracts/Proposal.sol";

// Supplement missing Proposal test cases
contract ProposalMissingFoundryTest is BaseSetup {

    address v1; address v2; address v3; address v4; address v5;

    function setUp() public {
        v1 = makeAddr("v1");
        v2 = makeAddr("v2");
        v3 = makeAddr("v3");
        v4 = makeAddr("v4");
        v5 = makeAddr("v5");
        address[] memory initVals = new address[](5); // use 5 validators to test reject scenario
        initVals[0] = v1; initVals[1] = v2; initVals[2] = v3; initVals[3] = v4; initVals[4] = v5;
        deploySystem(initVals);
    }

    function testAnyoneCanCreateProposal() public {
        // Corresponds to "anyone can create proposal"
        Proposal p = Proposal(PROPOSAL);
        address candidate = makeAddr("candidate");
        
        // Test multiple different accounts can create proposals
        for (uint i = 0; i < 5; i++) {
            address creator = makeAddr(string(abi.encodePacked("creator", i)));
            vm.deal(creator, 1 ether);
            
            vm.warp(1_000_000 + i * 10000); // use different timestamps to avoid ID conflicts
            vm.prank(creator);
            bool success = p.createProposal(candidate, true, "");
            require(success, "should create proposal successfully");
        }
    }

    function testProposalReject() public {
        // Test normal vote (2 agree, 3 reject)
        Proposal p = Proposal(PROPOSAL);
        address candidate = makeAddr("candidate");
        
        vm.warp(2_000_000);
        bytes32 id = keccak256(abi.encodePacked(address(this), candidate, true, "test", block.timestamp));
        p.createProposal(candidate, true, "test");
        
        // 2 votes agree
        vm.prank(v1); p.voteProposal(id, true);
        vm.prank(v2); p.voteProposal(id, true);
        
        // 3 votes reject
        vm.prank(v3); p.voteProposal(id, false);
        vm.prank(v4); p.voteProposal(id, false);
        vm.prank(v5); p.voteProposal(id, false);
        
        // The proposal should be rejected, and the candidate should not pass
        require(!p.pass(candidate), "candidate should not pass");
        
        // Check voting results
        (uint16 agree, uint16 reject, bool resultExist) = p.results(id);
        require(agree == 2, "should have 2 agree votes");
        require(reject == 3, "should have 3 reject votes");
        require(resultExist, "result should exist");
    }

    function testValidateCandidateInfo() public {
        // Corresponds to "Validate candidate's info"
        Proposal p = Proposal(PROPOSAL);
        address candidate = makeAddr("candidate");
        address proposer = makeAddr("proposer");
        
        vm.warp(3_000_000);
        bytes32 id = keccak256(abi.encodePacked(proposer, candidate, true, "test details", block.timestamp));
        
        vm.prank(proposer);
        p.createProposal(candidate, true, "test details");
        
        // Vote passes
        vm.prank(v1); p.voteProposal(id, true);
        vm.prank(v2); p.voteProposal(id, true);
        vm.prank(v3); p.voteProposal(id, true);
        
        // Verify proposal information
        (address storedProposer, uint256 createTime, uint256 proposalType, address dst, bool flag, string memory details, , ) = p.proposals(id);
        require(storedProposer == proposer, "proposer should match");
        require(dst == candidate, "candidate should match");
        require(flag == true, "flag should be true");
        require(keccak256(bytes(details)) == keccak256(bytes("test details")), "details should match");
        require(proposalType == 1, "should be validator proposal type");
        require(createTime == 3_000_000, "create time should match");
        
        // Verify voting results
        (uint16 agree, uint16 reject, bool resultExist) = p.results(id);
        require(agree == 3, "should have 3 agree votes");
        require(reject == 0, "should have 0 reject votes");
        require(resultExist, "result should exist");
        
        // Verify candidate status
        require(p.pass(candidate), "candidate should pass");
    }

    function testSetUnpassedPermission() public {
        // Test only validator can set val unpass
        Proposal p = Proposal(PROPOSAL);
        address candidate = v1; // Use existing validator
        
        // Non-validator contract calls should fail
        vm.prank(makeAddr("random"));
        (bool ok, ) = address(p).call(abi.encodeWithSelector(p.setUnpassed.selector, candidate));
        require(!ok, "should fail when called by non-validator contract");
    }

    function testSetUnpassedByValidatorContract() public {
        // Test validator contract can set val unpass
        // Note: This test needs to be called from the Validators contract, we simulate this process here
        
        address candidate = v1;
        
        // Confirm the candidate initially passes
        require(Proposal(PROPOSAL).pass(candidate), "candidate should initially pass");
        
        // Simulate Validators contract calling setUnpassed
        // In actual tests, this is automatically triggered through punishment process or removal proposals
        // We test this functionality through removing validator proposals
        vm.warp(4_000_000);
        bytes32 id = keccak256(abi.encodePacked(address(this), candidate, false, "", block.timestamp));
        Proposal(PROPOSAL).createProposal(candidate, false, "");
        
        // Vote to remove
        vm.prank(v2); Proposal(PROPOSAL).voteProposal(id, true);
        vm.prank(v3); Proposal(PROPOSAL).voteProposal(id, true);
        vm.prank(v4); Proposal(PROPOSAL).voteProposal(id, true);
        
        // Verify the candidate now does not pass
        require(!Proposal(PROPOSAL).pass(candidate), "candidate should not pass after removal");
    }
}
