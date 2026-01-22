// SPDX-License-Identifier: MIT

pragma solidity ^0.8.29;

import {Params} from "./Params.sol";
import {IStaking} from "./IStaking.sol";
import {IValidators} from "./IValidators.sol";
import {IProposal} from "./IProposal.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

contract Punish is Params, ReentrancyGuard {
    struct PunishRecord {
        uint256 missedBlocksCounter;
        uint256 index;
        bool exist;
    }

    IValidators validators;
    IProposal proposal;
    IStaking staking;

    mapping(address => PunishRecord) punishRecords;
    address[] public punishValidators;

    mapping(uint256 => bool) punished;
    mapping(uint256 => bool) decreased;
    mapping(address => bool) pendingRemove;
    mapping(address => bool) pendingRemoveIncoming;

    event LogDecreaseMissedBlocksCounter();
    event LogPunishValidator(address indexed val, uint256 time);

    modifier onlyNotPunished() {
        _onlyNotPunished();
        _;
    }

    modifier onlyNotDecreased() {
        _onlyNotDecreased();
        _;
    }

    function _onlyNotPunished() internal view {
        require(!punished[block.number], "Already punished");
    }

    function _onlyNotDecreased() internal view {
        require(!decreased[block.number], "Already decreased");
    }

    /**
     * @dev Initializes the Punish contract with required dependencies.
     * @param validators_ Address of the Validators contract.
     * @param proposal_ Address of the Proposal contract.
     * @param staking_ Address of the Staking contract.
     */
    function initialize(
        address validators_,
        address proposal_,
        address staking_
    ) external onlyNotInitialized {
        require(validators_ != address(0), "Invalid validators address");
        require(proposal_ != address(0), "Invalid proposal address");
        
        validators = IValidators(validators_);
        proposal = IProposal(proposal_);
        staking = IStaking(staking_);
        _initializeEpoch(proposal.epoch());

        initialized = true;
    }

    /**
     * @dev Punishes a validator for missing blocks.
     * @param val Address of the validator to punish.
     * @notice Only the miner can call this function.
     * @notice Punishment is applied based on missed blocks threshold.
     */
    function punish(address val) external onlyMiner onlyInitialized onlyNotPunished nonReentrant {
        punished[block.number] = true;
        require(epoch > 0, "Epoch not set");
        bool isEpochBlock = block.number % epoch == 0;

        if (!isEpochBlock) {
            if (pendingRemove[val]) {
                pendingRemove[val] = false;
                pendingRemoveIncoming[val] = false;
                punishRecords[val].missedBlocksCounter = 0;
                staking.jailValidator(val, proposal.validatorUnjailPeriod());
                validators.removeValidator(val);
                emit LogPunishValidator(val, block.timestamp);
                return;
            }
            if (pendingRemoveIncoming[val]) {
                pendingRemoveIncoming[val] = false;
                validators.removeValidatorIncoming(val);
                emit LogPunishValidator(val, block.timestamp);
                return;
            }
        }

        if (!punishRecords[val].exist) {
            punishRecords[val].index = punishValidators.length;
            punishValidators.push(val);
            punishRecords[val].exist = true;
        }
        punishRecords[val].missedBlocksCounter++;

        if (isEpochBlock) {
            if (punishRecords[val].missedBlocksCounter % proposal.removeThreshold() == 0) {
                pendingRemove[val] = true;
            } else if (punishRecords[val].missedBlocksCounter % proposal.punishThreshold() == 0) {
                pendingRemoveIncoming[val] = true;
            }
            emit LogPunishValidator(val, block.timestamp);
            return;
        }

        if (punishRecords[val].missedBlocksCounter % proposal.removeThreshold() == 0) {
            // reset validator's missed blocks counter
            punishRecords[val].missedBlocksCounter = 0;
            // jail validator first (sets isJailed in Staking contract)
            staking.jailValidator(val, proposal.validatorUnjailPeriod());
            // then remove validator (which will check isJailed status)
            validators.removeValidator(val);
        } else if (punishRecords[val].missedBlocksCounter % proposal.punishThreshold() == 0) {
            validators.removeValidatorIncoming(val);
        }

        emit LogPunishValidator(val, block.timestamp);
    }

    /**
     * @dev Decreases the missed blocks counter for all validators at the end of an epoch.
     * @param epoch The epoch number for which to decrease the counter.
     * @notice Only the miner can call this function.
     * @notice This function is called once per epoch to reduce punishment counters.
     */
    function decreaseMissedBlocksCounter(uint256 epoch)
        external
        onlyMiner
        onlyNotDecreased
        onlyInitialized
        onlyBlockEpoch(epoch)
    {
        decreased[block.number] = true;
        if (punishValidators.length == 0) {
            return;
        }

        // Cache external results outside the loop
        uint256 removeThreshold = proposal.removeThreshold();
        uint256 decreaseRate = proposal.decreaseRate();
        uint256 decreaseAmount = removeThreshold / decreaseRate;

        uint256 punishValidatorsLength = punishValidators.length;
        for (uint256 i = 0; i < punishValidatorsLength; i++) {
            address validator = punishValidators[i];
            if (punishRecords[validator].missedBlocksCounter > decreaseAmount) {
                punishRecords[validator].missedBlocksCounter -= decreaseAmount;
            } else {
                punishRecords[validator].missedBlocksCounter = 0;
            }
        }

        emit LogDecreaseMissedBlocksCounter();
    }

    /**
     * @dev Cleans the punishment record for a validator.
     * @param val Address of the validator whose record to clean.
     * @return bool Returns true if the operation was successful.
     * @notice This function is called when a validator restakes.
     */
    function cleanPunishRecord(address val) public onlyInitialized onlyValidatorsContract returns (bool) {
        if (punishRecords[val].missedBlocksCounter != 0) {
            punishRecords[val].missedBlocksCounter = 0;
        }

        // remove it out of array if exist
        if (punishRecords[val].exist && punishValidators.length > 0) {
            if (punishRecords[val].index != punishValidators.length - 1) {
                address uval = punishValidators[punishValidators.length - 1];
                punishValidators[punishRecords[val].index] = uval;

                punishRecords[uval].index = punishRecords[val].index;
            }
            punishValidators.pop();
            punishRecords[val].index = 0;
            punishRecords[val].exist = false;
        }

        return true;
    }

    /**
     * @dev Gets the number of validators currently being punished.
     * @return uint256 The number of validators in the punishment list.
     */
    function getPunishValidatorsLen() public view returns (uint256) {
        return punishValidators.length;
    }

    /**
     * @dev Gets the punishment record for a specific validator.
     * @param val Address of the validator to check.
     * @return uint256 The number of missed blocks for the validator.
     */
    function getPunishRecord(address val) public view returns (uint256) {
        return punishRecords[val].missedBlocksCounter;
    }
}
