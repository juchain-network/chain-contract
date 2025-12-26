// SPDX-License-Identifier: MIT

pragma solidity ^0.8.29;

import {Params} from './Params.sol';
import {IStaking} from './IStaking.sol';
import {IValidators} from './IValidators.sol';
import {IProposal} from './IProposal.sol';
import {ReentrancyGuard} from '@openzeppelin/contracts/utils/ReentrancyGuard.sol';

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
        require(!punished[block.number], 'Already punished');
    }

    function _onlyNotDecreased() internal view {
        require(!decreased[block.number], 'Already decreased');
    }

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

        initialized = true;
    }

    function punish(address val) external onlyMiner onlyInitialized onlyNotPunished nonReentrant {
        punished[block.number] = true;
        if (!punishRecords[val].exist) {
            punishRecords[val].index = punishValidators.length;
            punishValidators.push(val);
            punishRecords[val].exist = true;
        }
        punishRecords[val].missedBlocksCounter++;

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

    // clean validator's punish record if one restake in
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

    function getPunishValidatorsLen() public view returns (uint256) {
        return punishValidators.length;
    }

    function getPunishRecord(address val) public view returns (uint256) {
        return punishRecords[val].missedBlocksCounter;
    }
}
