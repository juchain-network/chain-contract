// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {Test} from "forge-std/Test.sol";
import {Proposal} from "../contracts/Proposal.sol";

contract MockProposalValidators {
    mapping(address => bool) internal active;

    function setActive(address validator, bool value) external {
        active[validator] = value;
    }

    function isValidatorActive(address validator) external view returns (bool) {
        return active[validator];
    }

    function isTopValidator(address) external pure returns (bool) {
        return false;
    }

    function getVotingValidatorCount() external pure returns (uint256) {
        return 3;
    }
}

contract ProposalHarness is Proposal {
    function exposedOnlyValidator() external view onlyValidator returns (bool) {
        return true;
    }

    function exposedValidateConfig(uint256 cid, uint256 value) external view returns (bool) {
        return validateConfig(cid, value);
    }

    function exposedCheckProposalCooldown() external returns (bool) {
        _checkProposalCooldown();
        return true;
    }

    function setLastProposalBlockForTest(address validator, uint256 blockNumber_) external {
        lastProposalBlock[validator] = blockNumber_;
    }
}

contract ProposalInternalCoverageTest is Test {
    address internal constant VALIDATOR_ADDR = 0x000000000000000000000000000000000000F010;

    ProposalHarness internal proposal;
    MockProposalValidators internal mockValidators;
    address internal v1;

    function setUp() public {
        v1 = makeAddr("proposal-harness-v1");

        mockValidators = new MockProposalValidators();
        vm.etch(VALIDATOR_ADDR, address(mockValidators).code);
        MockProposalValidators(VALIDATOR_ADDR).setActive(v1, true);

        proposal = new ProposalHarness();
        address[] memory vals = new address[](1);
        vals[0] = v1;
        proposal.initialize(vals, VALIDATOR_ADDR, 10);
    }

    function testOnlyValidatorInternalPath() public {
        vm.prank(v1);
        assertTrue(proposal.exposedOnlyValidator());

        address stranger = makeAddr("proposal-harness-stranger");
        vm.prank(stranger);
        vm.expectRevert("Validator only");
        proposal.exposedOnlyValidator();
    }

    function testValidateConfigReturnsTrueForKnownCid() public view {
        assertTrue(proposal.exposedValidateConfig(19, 123));
    }

    function testCheckProposalCooldownInternalPath() public {
        vm.prank(v1);
        assertTrue(proposal.exposedCheckProposalCooldown());

        proposal.setLastProposalBlockForTest(v1, block.number);
        vm.prank(v1);
        vm.expectRevert("Proposal creation too frequent");
        proposal.exposedCheckProposalCooldown();

        vm.roll(block.number + proposal.proposalCooldown() + 1);
        vm.prank(v1);
        assertTrue(proposal.exposedCheckProposalCooldown());
    }
}
