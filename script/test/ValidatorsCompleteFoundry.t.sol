// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseSetup} from "./BaseSetup.t.sol";
import {Validators} from "../../contracts/Validators.sol";
import {Proposal} from "../../contracts/Proposal.sol";
import {Staking} from "../../contracts/Staking.sol";

// Complete validator tests, corresponding to all functions in test/validators.js
contract ValidatorsCompleteFoundryTest is BaseSetup {

    address v1; address v2; address v3;
    address miner;
    uint256 constant ACTIVE = 1;
    uint256 constant JAILED = 2;

    function setUp() public {
        v1 = makeAddr("v1");
        v2 = makeAddr("v2"); 
        v3 = makeAddr("v3");
        address[] memory initVals = new address[](3);
        initVals[0] = v1; initVals[1] = v2; initVals[2] = v3;
        deploySystem(initVals);
        miner = v1;
        vm.coinbase(miner);
        // Give miner enough ETH for testing
        vm.deal(miner, 100 ether);
    }

    function testCanOnlyInitOnce() public {
        // Corresponds to "can only init once"
        bytes memory err;
        try Validators(VALIDATORS).initialize(new address[](0), address(0xCAFE), address(0xBEEF), address(0xDEAD)) { 
            revert("should revert"); 
        } catch (bytes memory e) { 
            err = e; 
        }
        require(err.length > 0, "expected revert");
    }

    // Create/edited validator related tests
    function testCreateValidatorInvalidFeeAddr() public {
        // Corresponds to "can't create validator if fee addr == address(0)"
        address validator = makeAddr("validator");
        vm.prank(validator);
        (bool ok, ) = address(Validators(VALIDATORS)).call(
            abi.encodeWithSelector(
                Validators.createOrEditValidator.selector,
                address(0), "", "", "", "", ""
            )
        );
        require(!ok, "should fail with zero address");
    }

    function testCreateValidatorInvalidDescription() public {
        // Corresponds to "can't create validator if describe info invalid"
        address validator = makeAddr("validator");
        string memory tooLongMoniker = _generateLongString(71); // > 70 limit
        vm.prank(validator);
        (bool ok, ) = address(Validators(VALIDATORS)).call(
            abi.encodeWithSelector(
                Validators.createOrEditValidator.selector,
                payable(validator), tooLongMoniker, "", "", "", ""
            )
        );
        require(!ok, "should fail with too long moniker");
    }

    function testCreateValidatorNotAuthorized() public {
        // Corresponds to "can't create validator if not pass propose"
        address validator = makeAddr("validator");
        vm.prank(validator);
        (bool ok, ) = address(Validators(VALIDATORS)).call(
            abi.encodeWithSelector(
                Validators.createOrEditValidator.selector,
                payable(validator), "", "", "", "", ""
            )
        );
        require(!ok, "should fail without authorization");
    }

    function testCreateValidatorSuccess() public {
        // Corresponds to "create validator"
        address validator = makeAddr("validator");
        
        // First authorize through proposal
        _passProposal(validator, true);
        
        // Create validator (set validator info)
        vm.prank(validator);
        bool success = Validators(VALIDATORS).createOrEditValidator(payable(validator), "", "", "", "", "");
        require(success, "create validator should succeed");
        
        // In POSA mode, validator must register (stake) after creation to become active
        // Give validator enough ETH and register
        uint256 minStake = Proposal(PROPOSAL).minValidatorStake();
        vm.deal(validator, minStake);
        vm.prank(validator);
        Staking(STAKING).registerValidator{value: minStake}(1000); // 10% commission
        
        // Check if validator is registered (exists)
        require(Validators(VALIDATORS).isValidatorExist(validator), "validator should exist after registration");
        
        // Check if validator is in highestValidatorsSet (automatically added after registration)
        require(Validators(VALIDATORS).isTopValidator(validator), "validator should be in highestValidatorsSet after registration");
        
        // Note: In POSA mode, getValidatorStatus's isActive checks if validator is in currentValidatorSet
        // Newly registered validators won't immediately enter currentValidatorSet, need to wait for next epoch update
        // So here we check if validator exists and is not jailed
        // Validator status should be NotExist (because not yet in currentValidatorSet) or Active (if already in currentValidatorSet)
        // But at least validator should exist and not be jailed
        require(!Validators(VALIDATORS).isValidatorJailed(validator), "validator should not be jailed");
        
        // Validator should be queryable for info (registered)
        (address payable feeAddr, , , ,) = Validators(VALIDATORS).getValidatorInfo(validator);
        require(feeAddr == validator, "validator info should be set");
    }

    function testEditValidatorInfo() public {
        // Corresponds to "edit validator info"
        address validator = makeAddr("validator");
        address feeAddr = makeAddr("feeAddr");
        
        // First authorize and create
        _passProposal(validator, true);
        vm.prank(validator);
        Validators(VALIDATORS).createOrEditValidator(payable(validator), "", "", "", "", "");
        
        // Edit fee address
        vm.prank(validator);
        bool success = Validators(VALIDATORS).createOrEditValidator(payable(feeAddr), "", "", "", "", "");
        require(success, "edit should succeed");
        
        // Check fee address is updated
        (address payable actualFeeAddr,,,,) = Validators(VALIDATORS).getValidatorInfo(validator);
        require(actualFeeAddr == feeAddr, "fee address should be updated");
    }

    // Proposal add/remove validator tests
    function testProposeAddNewValidator() public {
        // Corresponds to "propose add a new val"
        address nval = makeAddr("newval");
        
        // Initially not a validator
        require(!Validators(VALIDATORS).isTopValidator(nval), "should not be validator initially");
        
        // Create and vote to pass proposal
        _passProposal(nval, true);
        
        // In POSA mode, proposal passing only sets pass[nval] = true
        // Validator also needs to register (stake) to become top validator
        require(Proposal(PROPOSAL).pass(nval), "should be marked as passed");
        require(!Validators(VALIDATORS).isTopValidator(nval), "should not be top validator yet (not registered)");
        
        // Give new validator enough ETH and register
        uint256 minStake = Proposal(PROPOSAL).minValidatorStake();
        vm.deal(nval, minStake);
        vm.prank(nval);
        Staking(STAKING).registerValidator{value: minStake}(1000); // 10% commission
        
        // Now should be a validator
        require(Validators(VALIDATORS).isTopValidator(nval), "should be validator after registration");
    }

    function testProposeRemoveValidator() public {
        // Corresponds to "propose remove a val"
        
        // v1 is initially a validator
        require(Validators(VALIDATORS).isTopValidator(v1), "v1 should be validator initially");
        
        // Create and vote to pass removal proposal
        _passProposal(v1, false);
        
        // Should not be a validator now
        require(!Validators(VALIDATORS).isTopValidator(v1), "v1 should not be validator after removal");
        require(!Proposal(PROPOSAL).pass(v1), "v1 should not be marked as passed");
    }

    // Block reward distribution tests
    function testDistributeBlockReward() public {
        // Corresponds to "miner can distribute to validator contract, the profits should be right updated"
        // New logic: reward goes directly to the block producer (miner = v1)
        uint256 fee = 0.3 ether;
        
        vm.prank(miner);
        Validators(VALIDATORS).distributeBlockReward{value: fee}();
        
        // Check block producer (v1) gets full reward
        (, , uint256 v1Incoming,,) = Validators(VALIDATORS).getValidatorInfo(v1);
        require(v1Incoming == fee, "block producer should get full reward");
        
        // Check other validators get no reward
        (, , uint256 v2Incoming,,) = Validators(VALIDATORS).getValidatorInfo(v2);
        (, , uint256 v3Incoming,,) = Validators(VALIDATORS).getValidatorInfo(v3);
        require(v2Incoming == 0, "v2 should get no reward");
        require(v3Incoming == 0, "v3 should get no reward");
    }

    function testUpdateWithdrawProfitPeriod() public {
        // Corresponds to "update withdraw profit wait block"
        vm.warp(5_000_000);
        
        // Create proposal from v1 (active validator) instead of address(this)
        vm.prank(v1);
        bytes32 id = Proposal(PROPOSAL).createUpdateConfigProposal(4, 10);
        
        vm.prank(v1); Proposal(PROPOSAL).voteProposal(id, true);
        vm.prank(v2); Proposal(PROPOSAL).voteProposal(id, true);
        vm.prank(v3); Proposal(PROPOSAL).voteProposal(id, true);
        
        require(Proposal(PROPOSAL).withdrawProfitPeriod() == 10, "withdraw period should be updated");
    }

    function testValidatorWithdrawProfits() public {
        // Corresponds to "validator can withdraw profits"
        uint256 fee = 0.3 ether;
        
        // Distribute rewards (miner = v1 is block producer, gets full reward)
        vm.prank(miner);
        Validators(VALIDATORS).distributeBlockReward{value: fee}();
        
        // Set short withdrawal period
        _updateWithdrawPeriod(1);
        vm.roll(block.number + 2);
        
        uint256 balBefore = miner.balance;
        vm.prank(miner);
        Validators(VALIDATORS).withdrawProfits(miner);
        uint256 balAfter = miner.balance;
        
        require(balAfter > balBefore, "should receive profits");
        
        // Test different fee address
        address feeAddr = makeAddr("feeAddr");
        vm.prank(miner);
        Validators(VALIDATORS).createOrEditValidator(payable(feeAddr), "", "", "", "", "");
        
        // Distribute again (miner = v1 is still block producer)
        vm.prank(miner);
        Validators(VALIDATORS).distributeBlockReward{value: 0.5 ether}();
        
        vm.roll(block.number + 2);
        
        uint256 feeBalBefore = feeAddr.balance;
        vm.prank(feeAddr);
        Validators(VALIDATORS).withdrawProfits(miner);
        uint256 feeBalAfter = feeAddr.balance;
        
        require(feeBalAfter > feeBalBefore, "fee address should receive profits");
    }

    function testCantWithdrawWithoutProfits() public {
        // Corresponds to "Can't call withdrawProfits if you don't have any profits"
        address feeAddr = makeAddr("feeAddr");
        _updateWithdrawPeriod(1);
        vm.roll(block.number + 2);
        
        vm.prank(feeAddr);
        (bool ok, ) = address(Validators(VALIDATORS)).call(
            abi.encodeWithSelector(Validators.withdrawProfits.selector, miner)
        );
        require(!ok, "should fail without profits");
    }

    function testUpdateActiveValidatorSet() public {
        // Corresponds to "update active validator set"
        uint256 epoch = 30;
        address[] memory newSet = new address[](2);
        newSet[0] = v1;
        newSet[1] = v2;
        
        // Simulate reaching epoch boundary (block.number % epoch == 0)
        uint256 targetBlock = ((block.number / epoch) + 1) * epoch;
        vm.roll(targetBlock);
        
        vm.prank(miner);
        Validators(VALIDATORS).updateActiveValidatorSet(newSet, epoch);
        
        // Verify new validator set
        address[] memory activeSet = Validators(VALIDATORS).getActiveValidators();
        require(activeSet.length == 2, "should have 2 active validators");
        require(activeSet[0] == v1, "should contain v1");
        require(activeSet[1] == v2, "should contain v2");
    }

    // Helper functions
    function _passProposal(address target, bool flag) internal {
        vm.warp(block.timestamp + 1000000);
        
        // Create proposal from v1 (active validator) instead of address(this)
        vm.prank(v1);
        bytes32 id = Proposal(PROPOSAL).createProposal(target, flag, "");
        
        vm.prank(v1); Proposal(PROPOSAL).voteProposal(id, true);
        vm.prank(v2); Proposal(PROPOSAL).voteProposal(id, true);
        vm.prank(v3); Proposal(PROPOSAL).voteProposal(id, true);
    }

    function _updateWithdrawPeriod(uint256 period) internal {
        // Create proposal from v1 (active validator) instead of address(this)
        vm.prank(v1);
        bytes32 id = Proposal(PROPOSAL).createUpdateConfigProposal(4, period);
        
        vm.prank(v1); Proposal(PROPOSAL).voteProposal(id, true);
        vm.prank(v2); Proposal(PROPOSAL).voteProposal(id, true);
        vm.prank(v3); Proposal(PROPOSAL).voteProposal(id, true);
    }

    function _generateLongString(uint256 length) internal pure returns (string memory) {
        bytes memory result = new bytes(length);
        for (uint i = 0; i < length; i++) {
            result[i] = "a";
        }
        return string(result);
    }

    function testRemoveValidatorIncoming() public {
        // Test that removeValidatorIncoming correctly distributes incoming rewards
        // when a validator is removed
        
        // Add some incoming rewards to v1
        uint256 reward = 1 ether;
        vm.prank(miner);
        Validators(VALIDATORS).distributeBlockReward{value: reward}();
        
        // Check v1 has incoming rewards
        (, , uint256 v1Incoming,,) = Validators(VALIDATORS).getValidatorInfo(v1);
        require(v1Incoming == reward, "v1 should have incoming rewards");
        
        // Remove v1 from incoming set (must be called from Punish contract)
        vm.prank(PUNISH);
        Validators(VALIDATORS).removeValidatorIncoming(v1);
        
        // Check v1 no longer has incoming rewards
        (, , uint256 v1IncomingAfter,,) = Validators(VALIDATORS).getValidatorInfo(v1);
        require(v1IncomingAfter == 0, "v1 should have no incoming rewards after removal");
        
        // Check that rewards were distributed to other active validators
        (, , uint256 v2Incoming,,) = Validators(VALIDATORS).getValidatorInfo(v2);
        (, , uint256 v3Incoming,,) = Validators(VALIDATORS).getValidatorInfo(v3);
        
        // Calculate expected reward per remaining validator
        uint256 expectedReward = reward / 2; // v2 and v3 should split the reward
        require(v2Incoming >= expectedReward && v3Incoming >= expectedReward, "Rewards should be distributed to remaining validators");
    }

    function testValidatorSetCompetition() public {
        // Test that multiple validators compete for spots in the highestValidatorsSet
        uint256 minStake = Proposal(PROPOSAL).minValidatorStake();
        
        // Create multiple new validators
        address[] memory newValidators = new address[](5);
        for (uint i = 0; i < newValidators.length; i++) {
            newValidators[i] = makeAddr(string(abi.encodePacked("newValidator", i)));
            
            // Authorize through proposal
            _passProposal(newValidators[i], true);
            
            // Create validator
            vm.prank(newValidators[i]);
            Validators(VALIDATORS).createOrEditValidator(payable(newValidators[i]), "", "", "", "", "");
            
            // Register with minimum stake
            vm.deal(newValidators[i], minStake);
            vm.prank(newValidators[i]);
            Staking(STAKING).registerValidator{value: minStake}(1000);
        }
        
        // Add additional stake to some validators to make them competitive
        uint256 additionalStake = minStake * 2;
        vm.deal(newValidators[0], additionalStake);
        vm.prank(newValidators[0]);
        Staking(STAKING).addValidatorStake{value: additionalStake}();
        
        vm.deal(newValidators[1], additionalStake);
        vm.prank(newValidators[1]);
        Staking(STAKING).addValidatorStake{value: additionalStake}();
        
        // Get highest validators set
        address[] memory highestSet = Validators(VALIDATORS).getHighestValidators();
        
        // Check that validators with higher stake are in the highest set
        bool foundValidator0 = false;
        bool foundValidator1 = false;
        for (uint i = 0; i < highestSet.length; i++) {
            if (highestSet[i] == newValidators[0]) foundValidator0 = true;
            if (highestSet[i] == newValidators[1]) foundValidator1 = true;
        }
        
        require(foundValidator0, "Validator with higher stake should be in highest set");
        require(foundValidator1, "Validator with higher stake should be in highest set");
    }
}
