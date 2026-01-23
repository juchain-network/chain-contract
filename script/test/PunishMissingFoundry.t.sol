// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseSetup} from "./BaseSetup.t.sol";
import {Punish} from "../../contracts/Punish.sol";
import {Validators} from "../../contracts/Validators.sol";
import {Proposal} from "../../contracts/Proposal.sol";
import {Staking} from "../../contracts/Staking.sol";

// Supplement missing Punish contract test cases
contract PunishMissingFoundryTest is BaseSetup {

    address v1; address v2; address v3;

    function setUp() public {
        v1 = makeAddr("v1");
        v2 = makeAddr("v2");
        v3 = makeAddr("v3");
        address[] memory initVals = new address[](3);
        initVals[0] = v1; initVals[1] = v2; initVals[2] = v3;
        deploySystem(initVals);
    }

    // Helper function: Check if validator is jailed
    function isJailed(address validator) internal view returns (bool) {
        (, Validators.Status status, , , ) = Validators(VALIDATORS).getValidatorInfo(validator);
        return status == Validators.Status.Jailed;
    }

    function testPunishInitialization() public view {
        // Corresponds to "test punish contract deployment and basic setup"
        Punish punish = Punish(PUNISH);
        
        // Verify initial state
        require(punish.initialized(), "contract should be initialized");
        require(punish.getPunishValidatorsLen() == 0, "should have no punish validators initially");
        
        // Verify initial punish records are empty
        require(punish.getPunishRecord(v1) == 0, "v1 should have no punish record");
        require(punish.getPunishRecord(v2) == 0, "v2 should have no punish record");
        require(punish.getPunishRecord(v3) == 0, "v3 should have no punish record");
    }

    function testJailedValidatorReactivation() public {
        // Corresponds to "jailed record will be cleaned if validator repass proposal"
        Punish punish = Punish(PUNISH);
        
        // First punish validator until jailed
        vm.coinbase(VALIDATORS); // Set coinbase to VALIDATORS contract address
        vm.startPrank(VALIDATORS);
        uint256 currentBlock = block.number;
        for (uint256 i = 0; i < 48; i++) {
            vm.roll(currentBlock + i + 1); // Roll to the next block
            punish.punish(v1);
        }
        vm.stopPrank();
        
        // Validator should be jailed
        require(isJailed(v1), "v1 should be jailed");
        require(punish.getPunishRecord(v1) == 0, "v1 punish record should be reset after removal");
        
        // Get jailUntilBlock, ensure jail period has passed
        (, , , , , uint256 jailUntilBlock, , , , ) = Staking(STAKING).getValidatorInfo(v1);
        require(jailUntilBlock > 0, "v1 should have jailUntilBlock set");
        
        // Create proposal to reactivate validator
        vm.prank(v2);
        bytes32 id = Proposal(PROPOSAL).createProposal(v1, true, "");
        require(id != bytes32(0), "should create proposal");
        
        // Vote to pass proposal
        vm.prank(v2); Proposal(PROPOSAL).voteProposal(id, true);
        vm.prank(v3); Proposal(PROPOSAL).voteProposal(id, true);
        
        // After proposal passes, only sets pass[v1] = true, but won't automatically unjail
        // Validator is still in jailed state, proposal has passed
        require(isJailed(v1), "v1 should still be jailed after proposal passes");
        require(Proposal(PROPOSAL).pass(v1), "v1 should have pass status after proposal passes");
        
        // Wait for jail period to end
        vm.roll(jailUntilBlock + 1);
        
        // In POSA mode, validators need to manually call unjailValidator() to recover
        // unjailValidator() checks pass status (if violations > 3)
        vm.prank(v1);
        Staking(STAKING).unjailValidator(v1);
        
        // Validator should no longer be jailed and punish record cleared
        require(!isJailed(v1), "v1 should not be jailed after unjail");
        require(punish.getPunishRecord(v1) == 0, "v1 punish record should remain cleaned");
    }

    function testOnlyNotPunishedModifier() public {
        // Test that onlyNotPunished modifier prevents multiple punish calls in the same block
        Punish punish = Punish(PUNISH);
        
        // Set coinbase to VALIDATORS contract address
        vm.coinbase(VALIDATORS);
        vm.startPrank(VALIDATORS);
        
        // First punish call should succeed
        punish.punish(v1);
        
        // Second punish call in the same block should revert
        vm.expectRevert("Already punished");
        punish.punish(v1);
        
        // Also test with different validator in same block
        vm.expectRevert("Already punished");
        punish.punish(v2);
        
        vm.stopPrank();
    }
    
    function testComplexPunishWorkflow() public {
        // Corresponds to complex punishment workflow test
        Punish punish = Punish(PUNISH);
        
        // Test punishment workflow for multiple validators
        vm.coinbase(VALIDATORS);
        vm.startPrank(VALIDATORS);
        
        // Punish v1 to income removal threshold
        for (uint256 i = 0; i < 24; i++) {
            vm.roll(block.number + 1);
            punish.punish(v1);
        }
        
        // Punish v2 until removed
        for (uint256 i = 0; i < 48; i++) {
            vm.roll(block.number + 1);
            punish.punish(v2);
        }
        
        vm.stopPrank();
        
        // Verify status
        require(!isJailed(v1), "v1 should not be jailed yet but income removed");
        require(isJailed(v2), "v2 should be jailed after removal");
        require(punish.getPunishRecord(v1) == 24, "v1 should have 24 punish records");
        require(punish.getPunishRecord(v2) == 0, "v2 punish record should be reset after removal");
        
        // Continue punishing v1 until removed
        vm.coinbase(VALIDATORS);
        vm.startPrank(VALIDATORS);
        for (uint256 i = 0; i < 24; i++) {
            vm.roll(block.number + 1);
            punish.punish(v1);
        }
        vm.stopPrank();

        require(isJailed(v1), "v1 should now be jailed");
        require(punish.getPunishRecord(v1) == 0, "v1 punish record should be reset after removal");
        
        // Now both validators are jailed, but they are still in currentValidatorSet (until next epoch)
        // Jailed validators can still vote because they are still in currentValidatorSet
        // According to design logic: threshold = activeValidatorCount / 2 + 1
        // When activeValidatorCount = 3 (v1, v2, v3 all in currentValidatorSet): threshold = 3 / 2 + 1 = 2
        // So at least 2 votes are needed to pass proposal
        vm.warp(6_000_000);
        vm.prank(v3); // v3 is the only unjailed validator
        bytes32 id = Proposal(PROPOSAL).createProposal(v1, true, "");
        
        // v3 votes (1 vote)
        vm.prank(v3); Proposal(PROPOSAL).voteProposal(id, true);
        
        // Proposal should not pass, because threshold = 2, only 1 vote
        // Need one more vote (can be v1 or v2, although they are jailed, they are still in currentValidatorSet)
        vm.prank(v1); Proposal(PROPOSAL).voteProposal(id, true);
        
        // Now with 2 votes, should pass (threshold = 2)
        require(Proposal(PROPOSAL).pass(v1), "proposal should pass with 2 votes when threshold is 2");
        // Note: Proposal passing only sets pass[v1] = true, v1 is still in jailed state
        // v1 needs to unjail first, then register staking to become active validator again
        require(isJailed(v1), "v1 should still be jailed (proposal passing doesn't auto unjail)");
    }

    function testPunishPermission() public {
        // Test only Validators contract can call punish
        Punish punish = Punish(PUNISH);
        
        // Random address calls should fail
        vm.prank(makeAddr("random"));
        (bool success, ) = address(punish).call(abi.encodeWithSelector(punish.punish.selector, v1));
        require(!success, "should fail when called by non-validator contract");
        
        // Only VALIDATORS contract can call
        vm.coinbase(VALIDATORS);
        vm.prank(VALIDATORS);
        punish.punish(v1);
        require(punish.getPunishRecord(v1) == 1, "punish should succeed from VALIDATORS contract");
    }

    function testPunishRecordCleaning() public {
        // Test punish record cleaning mechanism
        Punish punish = Punish(PUNISH);
        
        // First punish but not reaching income removal threshold
        vm.coinbase(VALIDATORS);
        vm.startPrank(VALIDATORS);
        for (uint256 i = 0; i < 10; i++) {
            vm.roll(block.number + 1);
            punish.punish(v1);
        }
        vm.stopPrank();

        require(punish.getPunishRecord(v1) == 10, "should have 10 punish records");
        require(!isJailed(v1), "should not be jailed");
        
        // Directly test record cleaning through Validators contract
        // This simulates the scenario when validators are reactivated
        vm.prank(PROPOSAL);
        bool success = Validators(VALIDATORS).tryActive(v1);
        require(success, "tryActive should succeed");
        
        // Punish records should remain unchanged because validator is not jailed
        require(punish.getPunishRecord(v1) == 10, "punish record should remain unchanged for active validator");
        
        // Now jail validator then reactivate
        vm.coinbase(VALIDATORS);
        vm.startPrank(VALIDATORS);
        for (uint256 i = 0; i < 38; i++) { // Total 48 times to reach removal threshold
            vm.roll(block.number + 1);
            punish.punish(v1);
        }
        vm.stopPrank();
        
        require(isJailed(v1), "v1 should be jailed");
        require(punish.getPunishRecord(v1) == 0, "punish record should be reset after removal");
        
        // Records have been cleaned when reactivating
        vm.prank(PROPOSAL);
        success = Validators(VALIDATORS).tryActive(v1);
        require(success, "tryActive should succeed");
        require(punish.getPunishRecord(v1) == 0, "punish record should remain clean");
    }

    function testCleanPunishRecordWithIndexNotLast() public {
        // Test cleanPunishRecord when validator is not at the end of punishValidators array
        Punish punish = Punish(PUNISH);
        
        // Punish three validators to ensure they're in the array in v1, v2, v3 order
        vm.coinbase(VALIDATORS);
        vm.startPrank(VALIDATORS);
        
        // Punish v1 first
        for (uint256 i = 0; i < 5; i++) {
            vm.roll(block.number + 1);
            punish.punish(v1);
        }
        
        // Punish v2 second
        for (uint256 i = 0; i < 5; i++) {
            vm.roll(block.number + 1);
            punish.punish(v2);
        }
        
        // Punish v3 third
        for (uint256 i = 0; i < 5; i++) {
            vm.roll(block.number + 1);
            punish.punish(v3);
        }
        vm.stopPrank();
        
        // Verify all three are in the array
        require(punish.getPunishValidatorsLen() == 3, "should have 3 validators in punish array");
        require(punish.getPunishRecord(v1) == 5, "v1 should have 5 punish records");
        require(punish.getPunishRecord(v2) == 5, "v2 should have 5 punish records");
        require(punish.getPunishRecord(v3) == 5, "v3 should have 5 punish records");
        
        // Now directly test cleanPunishRecord on v2 (which is not at the end of the array)
        // This should test the branch where index != punishValidators.length - 1
        vm.prank(VALIDATORS);
        bool success = punish.cleanPunishRecord(v2);
        require(success, "cleanPunishRecord should succeed");
        
        // Verify v2's record is cleaned
        require(punish.getPunishRecord(v2) == 0, "v2's punish record should be cleaned");
        
        // Verify the array is correctly maintained
        require(punish.getPunishValidatorsLen() == 2, "should have 2 validators left");
        
        // Verify v1 and v3 are still in the array with correct records
        require(punish.getPunishRecord(v1) == 5, "v1's punish record should remain");
        require(punish.getPunishRecord(v3) == 5, "v3's punish record should remain");
    }

    function testOnlyNotDecreasedModifier() public {
        // Test that onlyNotDecreased modifier prevents multiple decrease calls in the same block
        Punish punish = Punish(PUNISH);
        uint256 epoch = Punish(PUNISH).epoch();
        
        // Roll to a block that is multiple of epoch
        uint256 targetBlock = epoch * 2;
        vm.roll(targetBlock);
        
        // Set coinbase to VALIDATORS contract address (onlyMiner)
        vm.coinbase(VALIDATORS);
        vm.startPrank(VALIDATORS);
        
        // First call should succeed
        punish.decreaseMissedBlocksCounter(epoch);
        
        // Second call in the same block should revert
        vm.expectRevert("Already decreased");
        punish.decreaseMissedBlocksCounter(epoch);
        
        vm.stopPrank();
    }

    function testDecreaseMissedBlocksCounterAtEpochBlock() public {
        // Test decreaseMissedBlocksCounter at epoch block
        Punish punish = Punish(PUNISH);
        uint256 epoch = Punish(PUNISH).epoch();
        
        // Roll to a block that is multiple of epoch
        uint256 targetBlock = epoch * 2;
        vm.roll(targetBlock);
        
        // Call should succeed at epoch block
        vm.coinbase(VALIDATORS);
        vm.prank(VALIDATORS);
        punish.decreaseMissedBlocksCounter(epoch);
    }

    function testInitializeParameterValidation() public {
        // Deploy a new Punish contract to test initialize
        Punish punish = new Punish();
        
        // Test with zero address for validators
        vm.expectRevert("Invalid validators address");
        punish.initialize(address(0), PROPOSAL, STAKING);
        
        // Test with zero address for proposal
        vm.expectRevert("Invalid proposal address");
        punish.initialize(VALIDATORS, address(0), STAKING);
        
        // Test with both zero addresses
        vm.expectRevert("Invalid validators address");
        punish.initialize(address(0), address(0), STAKING);
    }

    function testDecreaseMissedBlocksCounterLogic() public {
        Punish punish = Punish(PUNISH);
        uint256 epoch = Punish(PUNISH).epoch();
        
        // Punish validators to set up different missed block counts
        vm.coinbase(VALIDATORS);
        vm.startPrank(VALIDATORS);
        
        // Case 1: v1 has 3 missed blocks (greater than threshold 2)
        vm.roll(block.number + 1);
        punish.punish(v1);
        vm.roll(block.number + 1);
        punish.punish(v1);
        vm.roll(block.number + 1);
        punish.punish(v1);
        
        // Case 2: v2 has 2 missed blocks (equal to threshold 2)
        vm.roll(block.number + 1);
        punish.punish(v2);
        vm.roll(block.number + 1);
        punish.punish(v2);
        
        // Case 3: v3 has 1 missed block (less than threshold 2)
        vm.roll(block.number + 1);
        punish.punish(v3);
        
        vm.stopPrank();
        
        // Verify initial state
        require(punish.getPunishRecord(v1) == 3, "v1 should have 3 missed blocks");
        require(punish.getPunishRecord(v2) == 2, "v2 should have 2 missed blocks");
        require(punish.getPunishRecord(v3) == 1, "v3 should have 1 missed block");
        
        // Roll to epoch block
        uint256 epochBlock = epoch * 5;
        vm.roll(epochBlock);
        
        // Call decreaseMissedBlocksCounter
        vm.coinbase(VALIDATORS);
        vm.prank(VALIDATORS);
        punish.decreaseMissedBlocksCounter(epoch);
        
        // Verify results:
        // v1: 3 - 2 = 1
        // v2: 2 <= 2, reset to 0
        // v3: 1 <= 2, reset to 0
        require(punish.getPunishRecord(v1) == 1, "v1 should have 1 missed block after decrease");
        require(punish.getPunishRecord(v2) == 0, "v2 should have 0 missed blocks after decrease");
        require(punish.getPunishRecord(v3) == 0, "v3 should have 0 missed blocks after decrease");
        
        // Test with more blocks to cover the other branch
        vm.roll(epochBlock + 1);
        vm.coinbase(VALIDATORS);
        vm.prank(VALIDATORS);
        punish.punish(v1); // v1 now has 2 missed blocks
        
        // Call decrease again at next epoch
        vm.roll(epochBlock + epoch);
        vm.coinbase(VALIDATORS);
        vm.prank(VALIDATORS);
        punish.decreaseMissedBlocksCounter(epoch);
        
        // Now v1 has 2 missed blocks, should be reset to 0
        require(punish.getPunishRecord(v1) == 0, "v1 should have 0 missed blocks after second decrease");
    }

    function testExecutePendingLimitZero() public {
        Punish(PUNISH).executePending(0);
    }

    function testPendingRemoveExecutesOnNextNonEpochPunish() public {
        Punish punish = Punish(PUNISH);
        uint256 removeThreshold = Proposal(PROPOSAL).removeThreshold();
        address miner = VALIDATORS;
        vm.coinbase(miner);

        for (uint256 i = 0; i < removeThreshold - 1; i++) {
            vm.roll(block.number + 1);
            vm.prank(miner);
            punish.punish(v1);
        }

        uint256 epoch = punish.epoch();
        uint256 nextEpoch = ((block.number / epoch) + 1) * epoch;
        vm.roll(nextEpoch);
        vm.prank(miner);
        punish.punish(v1);

        vm.roll(nextEpoch + 1);
        vm.prank(miner);
        punish.punish(v1);

        require(isJailed(v1), "v1 should be jailed after pending remove");
        require(punish.getPunishRecord(v1) == 0, "v1 punish record should be reset");
    }

    function testExecutePendingProcessesIncoming() public {
        Punish punish = Punish(PUNISH);
        uint256 punishThreshold = Proposal(PROPOSAL).punishThreshold();
        address miner = VALIDATORS;
        vm.coinbase(miner);

        for (uint256 i = 0; i < punishThreshold - 1; i++) {
            vm.roll(block.number + 1);
            vm.prank(miner);
            punish.punish(v2);
        }

        uint256 epoch = punish.epoch();
        uint256 nextEpoch = ((block.number / epoch) + 1) * epoch;
        vm.roll(nextEpoch);
        vm.prank(miner);
        punish.punish(v2);

        vm.roll(nextEpoch + 1);
        punish.executePending(1);

        vm.roll(nextEpoch + 2);
        vm.prank(miner);
        punish.punish(v2);
        require(punish.getPunishRecord(v2) == punishThreshold + 1, "v2 should continue to accumulate after pending");
    }

    function testPendingRemoveIncomingExecutesOnNextNonEpochPunish() public {
        Punish punish = Punish(PUNISH);
        uint256 punishThreshold = Proposal(PROPOSAL).punishThreshold();
        address miner = VALIDATORS;
        vm.coinbase(miner);

        for (uint256 i = 0; i < punishThreshold - 1; i++) {
            vm.roll(block.number + 1);
            vm.prank(miner);
            punish.punish(v3);
        }

        uint256 epoch = punish.epoch();
        uint256 nextEpoch = ((block.number / epoch) + 1) * epoch;
        vm.roll(nextEpoch);
        vm.prank(miner);
        punish.punish(v3);

        vm.roll(nextEpoch + 1);
        vm.prank(miner);
        punish.punish(v3);

        require(punish.getPunishRecord(v3) == punishThreshold, "v3 should keep record after pending incoming");
    }

    function testExecutePendingProcessesRemove() public {
        Punish punish = Punish(PUNISH);
        uint256 removeThreshold = Proposal(PROPOSAL).removeThreshold();
        address miner = VALIDATORS;
        vm.coinbase(miner);

        for (uint256 i = 0; i < removeThreshold - 1; i++) {
            vm.roll(block.number + 1);
            vm.prank(miner);
            punish.punish(v2);
        }

        uint256 epoch = punish.epoch();
        uint256 nextEpoch = ((block.number / epoch) + 1) * epoch;
        vm.roll(nextEpoch);
        vm.prank(miner);
        punish.punish(v2);

        vm.roll(nextEpoch + 1);
        punish.executePending(1);

        require(isJailed(v2), "v2 should be jailed after pending remove");
        require(punish.getPunishRecord(v2) == 0, "v2 punish record should be reset after pending remove");
    }

    function testCleanPunishRecordSwap() public {
        Punish punish = Punish(PUNISH);
        address miner = VALIDATORS;
        vm.coinbase(miner);

        vm.roll(block.number + 1);
        vm.prank(miner);
        punish.punish(v1);
        vm.roll(block.number + 1);
        vm.prank(miner);
        punish.punish(v2);

        require(punish.getPunishValidatorsLen() == 2, "should have two punish validators");

        vm.prank(VALIDATORS);
        punish.cleanPunishRecord(v1);

        require(punish.getPunishValidatorsLen() == 1, "should have one punish validator after clean");
        require(punish.getPunishRecord(v1) == 0, "v1 punish record should be cleared");
        require(punish.getPunishRecord(v2) == 1, "v2 punish record should remain");

        vm.prank(VALIDATORS);
        punish.cleanPunishRecord(v3);
    }
}
