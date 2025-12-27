// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {BaseSetup} from "./BaseSetup.t.sol";
import {Proposal} from "../../contracts/Proposal.sol";

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

    function testOnlyValidatorCanCreateProposal() public {
        // Test that only validators can create proposals
        Proposal p = Proposal(PROPOSAL);
        address candidate = makeAddr("candidate");
        
        // Test non-validator cannot create proposal
        address nonValidator = makeAddr("nonValidator");
        vm.warp(1_000_000);
        vm.prank(nonValidator);
        vm.expectRevert("Validator only");
        p.createProposal(candidate, true, "");
        
        // Test validators can create proposals
        for (uint i = 0; i < 3; i++) {
            address validator = [v1, v2, v3][i];
            vm.warp(1_000_000 + i * 10000); // use different timestamps to avoid ID conflicts
            vm.prank(validator);
            bytes32 proposalId = p.createProposal(candidate, true, "");
            require(proposalId != bytes32(0), "validator should create proposal successfully");
        }
    }

    function testProposalReject() public {
        // Test normal vote (2 agree, 3 reject)
        Proposal p = Proposal(PROPOSAL);
        address candidate = makeAddr("candidate");
        
        vm.warp(2_000_000);
        bytes32 id = keccak256(abi.encodePacked(v1, candidate, true, "test", block.timestamp));
        vm.prank(v1);
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
        address proposer = v1; // Use validator v1 as proposer
        
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
        bytes32 id = keccak256(abi.encodePacked(v1, candidate, false, "", block.timestamp));
        vm.prank(v1);
        Proposal(PROPOSAL).createProposal(candidate, false, "");
        
        // Vote to remove
        vm.prank(v2); Proposal(PROPOSAL).voteProposal(id, true);
        vm.prank(v3); Proposal(PROPOSAL).voteProposal(id, true);
        vm.prank(v4); Proposal(PROPOSAL).voteProposal(id, true);
        
        // Verify the candidate now does not pass
        require(!Proposal(PROPOSAL).pass(candidate), "candidate should not pass after removal");
    }
    
    function testUpdateConfigCID0() public {
        // Test updating proposalLastingPeriod (cid=0) with valid value
        Proposal p = Proposal(PROPOSAL);
        
        // Test with valid value (1 day = 86400 seconds)
        uint256 timestamp = 5_000_000;
        vm.warp(timestamp);
        uint256 validValue = 86400; // 1 day in seconds
        
        vm.prank(v1);
        bytes32 proposalId = p.createUpdateConfigProposal(0, validValue); // cid=0, value=86400 seconds (valid)
        require(proposalId != bytes32(0), "should create valid proposal");
        
        bytes32 id = keccak256(abi.encodePacked(v1, uint256(0), validValue, uint256(timestamp)));
        
        vm.prank(v1); p.voteProposal(id, true);
        vm.prank(v2); p.voteProposal(id, true);
        vm.prank(v3); p.voteProposal(id, true); // Majority vote
        
        require(p.proposalLastingPeriod() == validValue, "should update to valid value");
    }

    function testIsProposalValidForStakingBranches() public {
        Proposal p = Proposal(PROPOSAL);
        address validator = makeAddr("testValidator");
        
        // Case 1: pass[validator] is false -> returns false
        bool result1 = p.isProposalValidForStaking(validator);
        require(!result1, "should return false when pass is false");
        
        // Case 2: Test with an existing validator
        // Existing validators have pass=true and proposalPassedTime set
        bool result2 = p.isProposalValidForStaking(v1);
        require(result2, "should return true for existing validator");
        
        // Case 3: Test with existing validator after time warp
        vm.warp(block.timestamp + 8 days);
        bool result3 = p.isProposalValidForStaking(v1);
        require(!result3, "should return false for existing validator after time warp");
        
        // Reset time
        vm.warp(block.timestamp - 8 days);
    }
    
    // We can't test updateConfig directly since it's private
    // Instead, we'll test configuration proposals which indirectly call updateConfig
    function testCreateAndVoteConfigProposal() public {
        Proposal p = Proposal(PROPOSAL);
        
        // Test creating a configuration proposal
        vm.prank(v1); // Use validator v1 as proposer
        bytes32 id = keccak256(abi.encodePacked(v1, uint256(0), uint256(24 hours), block.timestamp));
        bytes32 proposalId = p.createUpdateConfigProposal(0, 24 hours);
        assert(proposalId != bytes32(0));
        
        // Vote on the proposal
        vm.prank(v1);
        p.voteProposal(id, true);
        
        vm.prank(v2);
        p.voteProposal(id, true);
        
        vm.prank(v3);
        p.voteProposal(id, true);
        
        // Since we have 3 validators and all voted, the proposal should pass
        // The actual updateConfig call happens internally when the proposal is processed
    }
    
    function testCreateProposalWithInvalidConditions() public {
        Proposal p = Proposal(PROPOSAL);
        address existingValidator = v1;
        
        // Test trying to add an already existing validator
        vm.prank(v1); // Use validator v1 as proposer
        vm.expectRevert("Can't add an already exist dst or Can't remove a not passed dst");
        p.createProposal(existingValidator, true, "Try to add existing validator");
        
        // Test details too long
        string memory longDetails = new string(3001);
        vm.prank(v1); // Use validator v1 as proposer
        vm.expectRevert("Details too long");
        p.createProposal(makeAddr("newValidator"), true, longDetails);
    }
    
    function testVoteProposalWithExpiredProposal() public {
        Proposal p = Proposal(PROPOSAL);
        address newValidator = makeAddr("newValidator");
        
        // Create a proposal
        vm.prank(v1); // Use validator v1 as proposer
        bytes32 id = keccak256(abi.encodePacked(v1, newValidator, true, "Add new validator", block.timestamp));
        p.createProposal(newValidator, true, "Add new validator");
        
        // Warp time to after the proposal period
        vm.warp(block.timestamp + p.proposalLastingPeriod() + 1);
        
        // Try to vote - should revert
        vm.prank(v1);
        vm.expectRevert("Proposal expired");
        p.voteProposal(id, true);
    }
    
    function testVoteProposalTwice() public {
        Proposal p = Proposal(PROPOSAL);
        address newValidator = makeAddr("newValidator");
        
        // Create a proposal
        vm.prank(v1); // Use validator v1 as proposer
        bytes32 id = keccak256(abi.encodePacked(v1, newValidator, true, "Add new validator", block.timestamp));
        p.createProposal(newValidator, true, "Add new validator");
        
        // First vote
        vm.prank(v1);
        p.voteProposal(id, true);
        
        // Try to vote again - should revert
        vm.prank(v1);
        vm.expectRevert("You can't vote for a proposal twice");
        p.voteProposal(id, true);
    }
    
    function testVoteProposalForNonExistent() public {
        Proposal p = Proposal(PROPOSAL);
        bytes32 nonExistentId = keccak256(abi.encodePacked("non-existent", block.timestamp));
        
        // Try to vote for non-existent proposal
        vm.prank(v1);
        vm.expectRevert("Proposal not exist");
        p.voteProposal(nonExistentId, true);
    }
    
    function testVoteProposalWithExistingResult() public {
        Proposal p = Proposal(PROPOSAL);
        address newValidator = makeAddr("newValidator");
        
        // Create a proposal
        vm.prank(v1); // Use validator v1 as proposer
        bytes32 id = keccak256(abi.encodePacked(v1, newValidator, true, "Add new validator", block.timestamp));
        p.createProposal(newValidator, true, "Add new validator");
        
        // Vote on the proposal
        vm.prank(v1);
        p.voteProposal(id, true);
        vm.prank(v2);
        p.voteProposal(id, true);
        vm.prank(v3);
        p.voteProposal(id, true);
        
        // Now try to vote again - should NOT revert, but just return true
        // because the function handles resultExist by returning early
        vm.prank(v4);
        bool result = p.voteProposal(id, true);
        require(result, "should return true even if proposal already has result");
    }
    
    function testCreateProposalToRemoveNonExistentValidator() public {
        Proposal p = Proposal(PROPOSAL);
        address nonExistentValidator = makeAddr("nonExistentValidator");
        
        // Try to create a proposal to remove a validator that hasn't passed
        vm.prank(v1); // Use validator v1 as proposer
        vm.expectRevert("Can't add an already exist dst or Can't remove a not passed dst");
        p.createProposal(nonExistentValidator, false, "Remove non-existent validator");
    }
    
    function testUpdateConfigCID1() public {
        Proposal p = Proposal(PROPOSAL);
        
        // Test updating punishThreshold (cid=1) with valid value
        uint256 validValue = 48; // New punish threshold
        
        // Create config proposal
        bytes32 id = keccak256(abi.encodePacked(v1, uint256(1), validValue, block.timestamp));
        vm.prank(v1);
        p.createUpdateConfigProposal(1, validValue);
        
        // Vote to pass
        vm.prank(v1); p.voteProposal(id, true);
        vm.prank(v2); p.voteProposal(id, true);
        vm.prank(v3); p.voteProposal(id, true);
        
        // Verify the config was updated
        require(p.punishThreshold() == validValue, "punishThreshold should be updated");
    }
    
    function testIsProposalValidForStakingWithZeroPassedTime() public {
        Proposal p = Proposal(PROPOSAL);
        address newValidator = makeAddr("newValidator");
        
        // First, test with a validator that has pass=false (and thus proposalPassedTime=0)
        bool result1 = p.isProposalValidForStaking(newValidator);
        require(!result1, "should return false when pass is false");
        
        // Check that proposalPassedTime is 0 initially
        require(p.proposalPassedTime(newValidator) == 0, "proposalPassedTime should be 0 initially");
        
        // Create a proposal and vote for it
        vm.prank(v1); // Use validator v1 as proposer
        p.createProposal(newValidator, true, "Add validator");
        
        // Get the proposal ID correctly by recreating it
        bytes32 id;
        uint256 createTime = block.timestamp;
        string memory details = "Add validator";
        id = keccak256(abi.encodePacked(v1, newValidator, true, details, createTime));
        
        // Vote to pass
        vm.prank(v1); p.voteProposal(id, true);
        vm.prank(v2); p.voteProposal(id, true);
        vm.prank(v3); p.voteProposal(id, true);
        
        // Now the validator should have pass=true and proposalPassedTime > 0
        require(p.pass(newValidator), "validator should pass after proposal");
        require(p.proposalPassedTime(newValidator) > 0, "proposalPassedTime should be set after passing");
        
        // Test that it returns true within the deadline
        bool result2 = p.isProposalValidForStaking(newValidator);
        require(result2, "should return true when within deadline");
        
        // Test that it returns false after the deadline
        vm.warp(block.timestamp + 8 days); // More than 7 days
        bool result3 = p.isProposalValidForStaking(newValidator);
        require(!result3, "should return false when after deadline");
        
        // Clean up using setUnpassed
        vm.prank(VALIDATORS);
        p.setUnpassed(newValidator);
    }
    
    function testUpdateConfigInvalidCID() public {
        Proposal p = Proposal(PROPOSAL);
        
        // Test updating with an invalid config ID (e.g., 100)
        uint256 invalidCid = 100;
        uint256 value = 86400;
        
        // Create config proposal - should revert immediately with invalid CID
        vm.expectRevert("Invalid config ID");
        vm.prank(v1); // Use validator v1 as proposer
        p.createUpdateConfigProposal(invalidCid, value);
    }

    // Additional tests to cover missing branches
    function testInitializeWithZeroValidatorContract() public {
        // Test initializing with _validators = address(0)
        // Deploy a new Proposal contract instance instead of using the already initialized one
        Proposal p = new Proposal();
        address[] memory vals = new address[](1);
        vals[0] = makeAddr("val1");
        vm.expectRevert("Invalid validators address");
        p.initialize(vals, address(0));
    }

    function testInitializeWithZeroAddressInVals() public {
        // Test initializing with zero address in vals array
        // Deploy a new Proposal contract instance instead of using the already initialized one
        Proposal p = new Proposal();
        address[] memory vals = new address[](2);
        vals[0] = makeAddr("val1");
        vals[1] = address(0); // Invalid address
        vm.expectRevert("Invalid validator address");
        p.initialize(vals, makeAddr("validators"));
    }

    function testCreateDuplicateProposal() public {
        // Test creating duplicate proposal (same parameters, different timestamps)
        Proposal p = Proposal(PROPOSAL);
        address candidate = makeAddr("candidate");
        
        // Create first proposal
        vm.warp(1_000_000);
        bytes32 id1 = keccak256(abi.encodePacked(v1, candidate, true, "test details", block.timestamp));
        vm.prank(v1); // Use validator v1 as proposer
        bytes32 proposalId1 = p.createProposal(candidate, true, "test details");
        require(proposalId1 != bytes32(0), "should create first proposal");
        
        // Create second proposal with different timestamp (should succeed)
        vm.warp(1_000_001);
        bytes32 id2 = keccak256(abi.encodePacked(v1, candidate, true, "test details", block.timestamp));
        vm.prank(v1); // Use validator v1 as proposer
        bytes32 proposalId2 = p.createProposal(candidate, true, "test details");
        require(proposalId2 != bytes32(0), "should create second proposal with different timestamp");
        
        // Verify both proposals exist
        (, uint256 createTime1, , , , , , ) = p.proposals(id1);
        (, uint256 createTime2, , , , , , ) = p.proposals(id2);
        require(createTime1 == 1_000_000, "first proposal should exist");
        require(createTime2 == 1_000_001, "second proposal should exist");

    }

    function testUpdateConfigCID2() public {
        // Test updating removeThreshold (cid=2)
        Proposal p = Proposal(PROPOSAL);
        uint256 validValue = 96; // New remove threshold
        
        bytes32 id = keccak256(abi.encodePacked(v1, uint256(2), validValue, block.timestamp));
        vm.prank(v1);
        p.createUpdateConfigProposal(2, validValue);
        
        // Vote to pass
        vm.prank(v1); p.voteProposal(id, true);
        vm.prank(v2); p.voteProposal(id, true);
        vm.prank(v3); p.voteProposal(id, true);
        
        // Verify the config was updated
        require(p.removeThreshold() == validValue, "removeThreshold should be updated");
    }

    function testUpdateConfigCID3() public {
        // Test updating decreaseRate (cid=3)
        Proposal p = Proposal(PROPOSAL);
        uint256 validValue = 48; // New decrease rate
        
        bytes32 id = keccak256(abi.encodePacked(v1, uint256(3), validValue, block.timestamp));
        vm.prank(v1);
        p.createUpdateConfigProposal(3, validValue);
        
        // Vote to pass
        vm.prank(v1); p.voteProposal(id, true);
        vm.prank(v2); p.voteProposal(id, true);
        vm.prank(v3); p.voteProposal(id, true);
        
        // Verify the config was updated
        require(p.decreaseRate() == validValue, "decreaseRate should be updated");
    }

    function testUpdateConfigCID4() public {
        // Test updating withdrawProfitPeriod (cid=4)
        Proposal p = Proposal(PROPOSAL);
        uint256 validValue = 172800; // New withdraw profit period (2 days)
        
        bytes32 id = keccak256(abi.encodePacked(v1, uint256(4), validValue, block.timestamp));
        vm.prank(v1);
        p.createUpdateConfigProposal(4, validValue);
        
        // Vote to pass
        vm.prank(v1); p.voteProposal(id, true);
        vm.prank(v2); p.voteProposal(id, true);
        vm.prank(v3); p.voteProposal(id, true);
        
        // Verify the config was updated
        require(p.withdrawProfitPeriod() == validValue, "withdrawProfitPeriod should be updated");
    }

    function testUpdateConfigCID5() public {
        // Test updating blockReward (cid=5)
        Proposal p = Proposal(PROPOSAL);
        uint256 validValue = 400_000_000_000_000_000; // New block reward (0.4 ether)
        
        bytes32 id = keccak256(abi.encodePacked(v1, uint256(5), validValue, block.timestamp));
        vm.prank(v1);
        p.createUpdateConfigProposal(5, validValue);
        
        // Vote to pass
        vm.prank(v1); p.voteProposal(id, true);
        vm.prank(v2); p.voteProposal(id, true);
        vm.prank(v3); p.voteProposal(id, true);
        
        // Verify the config was updated
        require(p.blockReward() == validValue, "blockReward should be updated");
    }

    function testUpdateConfigCID6() public {
        // Test updating unbondingPeriod (cid=6)
        Proposal p = Proposal(PROPOSAL);
        uint256 validValue = 1209600; // New unbonding period (14 days)
        
        bytes32 id = keccak256(abi.encodePacked(v1, uint256(6), validValue, block.timestamp));
        vm.prank(v1);
        p.createUpdateConfigProposal(6, validValue);
        
        // Vote to pass
        vm.prank(v1); p.voteProposal(id, true);
        vm.prank(v2); p.voteProposal(id, true);
        vm.prank(v3); p.voteProposal(id, true);
        
        // Verify the config was updated
        require(p.unbondingPeriod() == validValue, "unbondingPeriod should be updated");
    }

    function testUpdateConfigCID7() public {
        // Test updating validatorUnjailPeriod (cid=7)
        Proposal p = Proposal(PROPOSAL);
        uint256 validValue = 172800; // New validator unjail period (2 days)
        
        bytes32 id = keccak256(abi.encodePacked(v1, uint256(7), validValue, block.timestamp));
        vm.prank(v1);
        p.createUpdateConfigProposal(7, validValue);
        
        // Vote to pass
        vm.prank(v1); p.voteProposal(id, true);
        vm.prank(v2); p.voteProposal(id, true);
        vm.prank(v3); p.voteProposal(id, true);
        
        // Verify the config was updated
        require(p.validatorUnjailPeriod() == validValue, "validatorUnjailPeriod should be updated");
    }

    function testCreateProposalWithSameID() public {
        // Test creating proposal with same ID (same parameters and timestamp)
        Proposal p = Proposal(PROPOSAL);
        address candidate = makeAddr("candidate");
        string memory details = "test proposal";
        
        vm.warp(5_000_000);
        
        // Create first proposal
        vm.prank(v1);
        bytes32 proposalId = p.createProposal(candidate, true, details);
        require(proposalId != bytes32(0), "should create proposal successfully");
        
        // Try to create identical proposal with same timestamp
        vm.expectRevert("Proposal already exists");
        vm.prank(v1);
        p.createProposal(candidate, true, details);
    }

    function testUpdateConfigCID0Invalid() public {
        // Test updating proposalLastingPeriod with invalid values (cid=0)
        Proposal p = Proposal(PROPOSAL);
        
        // Test value too small (less than 1 hour)
        uint256 invalidValue1 = 3599; // 59 minutes 59 seconds
        vm.expectRevert("Invalid proposal period");
        vm.prank(v1); // Use validator v1 as proposer
        p.createUpdateConfigProposal(0, invalidValue1);
        
        // Test value too large (more than 30 days)
        uint256 invalidValue2 = 31 days + 1;
        vm.expectRevert("Invalid proposal period");
        vm.prank(v1); // Use validator v1 as proposer
        p.createUpdateConfigProposal(0, invalidValue2);
    }

    function testUpdateConfigCID1Invalid() public {
        // Test updating punishThreshold with invalid value (cid=1)
        Proposal p = Proposal(PROPOSAL);
        
        // Test zero value
        uint256 invalidValue = 0;
        vm.expectRevert("Config value must be positive");
        vm.prank(v1); // Use validator v1 as proposer
        p.createUpdateConfigProposal(1, invalidValue);
    }

    function testUpdateConfigCID2Invalid() public {
        // Test updating removeThreshold with invalid value (cid=2)
        Proposal p = Proposal(PROPOSAL);
        
        // Test zero value
        uint256 invalidValue = 0;
        vm.expectRevert("Config value must be positive");
        vm.prank(v1); // Use validator v1 as proposer
        p.createUpdateConfigProposal(2, invalidValue);
    }

    function testUpdateConfigCID3Invalid() public {
        // Test updating decreaseRate with invalid value (cid=3)
        Proposal p = Proposal(PROPOSAL);
        
        // Test zero value
        uint256 invalidValue = 0;
        vm.expectRevert("Config value must be positive");
        vm.prank(v1); // Use validator v1 as proposer
        p.createUpdateConfigProposal(3, invalidValue);
    }

    function testUpdateConfigCID4Invalid() public {
        // Test updating withdrawProfitPeriod with invalid value (cid=4)
        Proposal p = Proposal(PROPOSAL);
        
        // Test zero value
        uint256 invalidValue = 0;
        vm.expectRevert("Config value must be positive");
        vm.prank(v1); // Use validator v1 as proposer
        p.createUpdateConfigProposal(4, invalidValue);
    }

    function testUpdateConfigCID5Invalid() public {
        // Test updating blockReward with invalid value (cid=5)
        Proposal p = Proposal(PROPOSAL);
        
        // Test zero value
        uint256 invalidValue = 0;
        vm.expectRevert("Config value must be positive");
        vm.prank(v1); // Use validator v1 as proposer
        p.createUpdateConfigProposal(5, invalidValue);
    }

    function testUpdateConfigCID6Invalid() public {
        // Test updating unbondingPeriod with invalid value (cid=6)
        Proposal p = Proposal(PROPOSAL);
        
        // Test zero value
        uint256 invalidValue = 0;
        vm.expectRevert("Config value must be positive");
        vm.prank(v1); // Use validator v1 as proposer
        p.createUpdateConfigProposal(6, invalidValue);
    }

    function testUpdateConfigCID7Invalid() public {
        // Test updating validatorUnjailPeriod with invalid value (cid=7)
        Proposal p = Proposal(PROPOSAL);
        
        // Test zero value
        uint256 invalidValue = 0;
        vm.expectRevert("Config value must be positive");
        vm.prank(v1); // Use validator v1 as proposer
        p.createUpdateConfigProposal(7, invalidValue);
    }

    function testIsProposalValidForStakingExpired() public {
        // Test proposal expiration detection
        Proposal p = Proposal(PROPOSAL);
        
        address candidate = makeAddr("candidate");
        
        // Create and pass a proposal to add the validator
        vm.warp(6_000_000);
        vm.prank(v1);
        p.createProposal(candidate, true, "add validator");
        bytes32 id = keccak256(abi.encodePacked(v1, candidate, true, "add validator", block.timestamp));
        
        // Vote to pass
        vm.prank(v1); p.voteProposal(id, true);
        vm.prank(v2); p.voteProposal(id, true);
        vm.prank(v3); p.voteProposal(id, true);
        
        // Verify the proposal is valid initially
        require(p.isProposalValidForStaking(candidate), "proposal should be valid initially");
        
        // Fast forward time by 7 days + 1 second (exceeding STAKING_DEADLINE_PERIOD)
        vm.warp(block.timestamp + 7 days + 1);
        
        // Verify the proposal is now invalid due to expiration
        require(!p.isProposalValidForStaking(candidate), "proposal should be invalid after expiration");
    }
}