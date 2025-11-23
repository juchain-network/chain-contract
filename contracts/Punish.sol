// SPDX-License-Identifier: MIT

pragma solidity ^0.8.20;

import {Params} from './Params.sol';
import {Staking} from './Staking.sol';
import {Validators} from './Validators.sol';
import {Proposal} from './Proposal.sol';

contract Punish is Params {
    struct PunishRecord {
        uint256 missedBlocksCounter;
        uint256 index;
        bool exist;
    }

    Validators validators;
    Proposal proposal;
    Staking staking;

    mapping(address => PunishRecord) punishRecords;
    address[] public punishValidators;

    mapping(uint256 => bool) punished;
    mapping(uint256 => bool) decreased;

    event LogDecreaseMissedBlocksCounter();
    event LogPunishValidator(address indexed val, uint256 time);

    modifier onlyNotPunished() {
        require(!punished[block.number], 'Already punished');
        _;
    }

    modifier onlyNotDecreased() {
        require(!decreased[block.number], 'Already decreased');
        _;
    }

    function initialize(
        address _validators,
        address _proposal,
        address _staking
    ) external onlyNotInitialized {
        require(_validators != address(0), "Invalid validators address");
        require(_proposal != address(0), "Invalid proposal address");
        
        validators = Validators(_validators);
        proposal = Proposal(_proposal);
        staking = Staking(_staking);

        initialized = true;
    }

    function punish(address val) external onlyMiner onlyInitialized onlyNotPunished {
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

        for (uint256 i = 0; i < punishValidators.length; i++) {
            if (
                punishRecords[punishValidators[i]].missedBlocksCounter >
                proposal.removeThreshold() / proposal.decreaseRate()
            ) {
                punishRecords[punishValidators[i]].missedBlocksCounter =
                    punishRecords[punishValidators[i]].missedBlocksCounter -
                    proposal.removeThreshold() /
                    proposal.decreaseRate();
            } else {
                punishRecords[punishValidators[i]].missedBlocksCounter = 0;
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
