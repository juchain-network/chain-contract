// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseSetup} from "../test/BaseSetup.t.sol";
import {Staking} from "../contracts/Staking.sol";
import {Validators} from "../contracts/Validators.sol";
import {Proposal} from "../contracts/Proposal.sol";

// Staking system operation script
contract StakingOperationsScript is BaseSetup {
    event StakingInfo(string info, uint256 value);
    event ValidatorInfo(string info, address validator, uint256 stake);

    function run() external {
        // Run in test mode, deploy the system first
        address[] memory initialValidators = new address[](3);
        initialValidators[0] = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;
        initialValidators[1] = 0x70997970C51812dc3A010C7d01b50e0d17dc79C8;
        initialValidators[2] = 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC;
        deploySystem(initialValidators);

        emit StakingInfo("=== Staking System Operations ===", 0);

        // Display current validator status
        showValidatorStatus();

        // Demonstrate validator registration staking
        demonstrateValidatorRegistration();

        // Demonstrate delegation function
        demonstrateDelegation();

        // Demonstrate getting top validators
        demonstrateTopValidators();

        emit StakingInfo("=== Operations Complete ===", 0);
    }

    function showValidatorStatus() internal {
        emit StakingInfo("--- Current Validator Status ---", 0);

        Staking stakingContract = Staking(STAKING);
        Validators validatorsContract = Validators(VALIDATORS);

        // Safely check current validators
        address[] memory currentValidators = validatorsContract.getActiveValidators();
        emit StakingInfo("Current validators count", currentValidators.length);

        // Only check Staking status when there are validators
        if (currentValidators.length > 0) {
            for (uint256 i = 0; i < currentValidators.length && i < 5; i++) {
                address validator = currentValidators[i];

                // Safely get validator information
                try stakingContract.getValidatorInfo(validator) returns (
                    uint256 selfStake,
                    uint256 totalDelegated,
                    uint256, /* commissionRate */
                    uint256, /* accumulatedRewards */
                    bool, /* isJailed */
                    uint256, /* jailUntilBlock */
                    uint256, /* totalClaimedRewards */
                    uint256, /* lastClaimBlock */
                    bool, /* isRegistered */
                    uint256 /* totalRewards */
                ) {
                    emit ValidatorInfo("Validator", validator, selfStake + totalDelegated);
                } catch {
                    emit ValidatorInfo("Validator (no stake info)", validator, 0);
                }
            }
        } else {
            emit StakingInfo("No validators registered yet", 0);
        }
    }

    function demonstrateValidatorRegistration() internal {
        emit StakingInfo("--- Validator Registration Demo ---", 0);

        Proposal proposalContract = Proposal(PROPOSAL);

        // Safely get minimum staking requirement
        try proposalContract.minValidatorStake() returns (uint256 minStake) {
            emit StakingInfo("Minimum stake required", minStake);
        } catch {
            emit StakingInfo("Cannot get minimum stake", 0);
        }

        // This is just to demonstrate the interface, actual registration requires sending ETH
        emit StakingInfo("To register call: staking.registerValidator{value: minStake}(commissionRate)", 0);
    }

    function demonstrateDelegation() internal {
        emit StakingInfo("--- Delegation Demo ---", 0);

        Staking stakingContract = Staking(STAKING);
        Validators validatorsContract = Validators(VALIDATORS);

        // Safely get top validators (through Validators contract unified interface)
        try validatorsContract.getTopValidators() returns (address[] memory topValidators) {
            if (topValidators.length > 0) {
                address validator = topValidators[0];
                emit ValidatorInfo("Example delegation to validator", validator, 0);

                // Safely display delegation information
                try stakingContract.getDelegationInfo(msg.sender, validator) returns (
                    uint256 delegatedAmount, uint256 rewards
                ) {
                    emit StakingInfo("Delegated amount", delegatedAmount);
                    emit StakingInfo("Pending rewards", rewards);
                } catch {
                    emit StakingInfo("Cannot get delegation info", 0);
                }
            } else {
                emit StakingInfo("No validators available for delegation", 0);
            }
        } catch {
            emit StakingInfo("Cannot get top validators", 0);
        }
    }

    function demonstrateTopValidators() internal {
        emit StakingInfo("--- Top Validators ---", 0);

        Staking stakingContract = Staking(STAKING);
        Validators validatorsContract = Validators(VALIDATORS);

        // Safely get top validators (through Validators contract unified interface)
        try validatorsContract.getTopValidators() returns (address[] memory topValidators) {
            emit StakingInfo("Total top validators", topValidators.length);

            for (uint256 i = 0; i < topValidators.length && i < 3; i++) {
                try stakingContract.getValidatorInfo(topValidators[i]) returns (
                    uint256 selfStake,
                    uint256 totalDelegated,
                    uint256 commissionRate,
                    uint256, /* accumulatedRewards */
                    bool isJailed,
                    uint256 jailUntilBlock,
                    uint256, /* totalClaimedRewards */
                    uint256, /* lastClaimBlock */
                    bool, /* isRegistered */
                    uint256 /* totalRewards */
                ) {
                    emit ValidatorInfo("Validator info", topValidators[i], selfStake);
                    emit StakingInfo("Total delegated", totalDelegated);
                    emit StakingInfo("Commission rate %", commissionRate);
                    emit StakingInfo("Jailed", isJailed ? 1 : 0);
                    emit StakingInfo("Jail until block", jailUntilBlock);
                } catch {
                    emit ValidatorInfo("Validator (info unavailable)", topValidators[i], 0);
                }
            }
        } catch {
            emit StakingInfo("Cannot get top validators list", 0);
        }
    }
}
