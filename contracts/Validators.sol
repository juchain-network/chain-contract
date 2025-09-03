// SPDX-License-Identifier: MIT

pragma solidity ^0.8.20;

import {Params} from './Params.sol';
import {Proposal} from './Proposal.sol';
import {Punish} from './Punish.sol';
import {Staking} from './Staking.sol';
import {SafeMath} from './library/SafeMath.sol';

contract Validators is Params {
    using SafeMath for uint256;

    enum Status {
        // validator not exist, default status
        NotExist,
        // active
        Active,
        // validator is jailed by system(validator have to repropose)
        Jailed
    }

    struct Description {
        string moniker;
        string identity;
        string website;
        string email;
        string details;
    }

    struct Validator {
        address payable feeAddr;
        Status status;
        Description description;
        uint256 aacIncoming;
        uint256 totalJailedHb;
        uint256 lastWithdrawProfitsBlock;
    }

    mapping(address => Validator) validatorInfo;
    // current validator set used by chain
    // only changed at block epoch
    address[] public currentValidatorSet;
    // highest validator set(dynamic changed)
    address[] public highestValidatorsSet;
    // total jailed hb
    uint256 public totalJailedHb;

    // System contracts
    Proposal proposal;
    Punish punish;
    Staking staking;

    enum Operations {Distribute, UpdateValidators}
    // Record the operations is done or not.
    mapping(uint256 => mapping(uint8 => bool)) operationsDone;

    event LogEditValidator(address indexed val, address indexed fee, uint256 time);
    event LogActive(address indexed val, uint256 time);
    event LogAddToTopValidators(address indexed val, uint256 time);
    event LogRemoveFromTopValidators(address indexed val, uint256 time);
    event LogWithdrawProfits(address indexed val, address indexed fee, uint256 hb, uint256 time);
    event LogRemoveValidator(address indexed val, uint256 hb, uint256 time);
    event LogRemoveValidatorIncoming(address indexed val, uint256 hb, uint256 time);
    event LogDistributeBlockReward(address indexed coinbase, uint256 blockReward, uint256 time);
    event LogUpdateValidator(address[] newSet);

    modifier onlyNotRewarded() {
        require(operationsDone[block.number][uint8(Operations.Distribute)] == false, 'Block is already rewarded');
        _;
    }

    modifier onlyNotUpdated() {
        require(
            operationsDone[block.number][uint8(Operations.UpdateValidators)] == false,
            'Validators already updated'
        );
        _;
    }

    function initialize(
        address[] calldata vals,
        address _proposal,
        address _punish,
        address _staking
    ) external onlyNotInitialized {
        require(_proposal != address(0), "Invalid proposal address");
        require(_punish != address(0), "Invalid punish address");
        require(_staking != address(0), "Invalid staking address");
        
        proposal = Proposal(_proposal);
        punish = Punish(_punish);
        staking = Staking(_staking);

        for (uint256 i = 0; i < vals.length; i++) {
            require(vals[i] != address(0), 'Invalid validator address');

            if (!isActiveValidator(vals[i])) {
                currentValidatorSet.push(vals[i]);
            }
            if (!isTopValidator(vals[i])) {
                highestValidatorsSet.push(vals[i]);
            }
            if (validatorInfo[vals[i]].feeAddr == address(0)) {
                validatorInfo[vals[i]].feeAddr = payable(vals[i]);
            }
            // Important: NotExist validator can't get profits
            if (validatorInfo[vals[i]].status == Status.NotExist) {
                validatorInfo[vals[i]].status = Status.Active;
            }
        }

        initialized = true;
    }

    function createOrEditValidator(
        address payable feeAddr,
        string calldata moniker,
        string calldata identity,
        string calldata website,
        string calldata email,
        string calldata details
    ) external onlyInitialized returns (bool) {
        require(feeAddr != address(0), 'Invalid fee address');
        require(validateDescription(moniker, identity, website, email, details), 'Invalid description');
    address payable validator = payable(msg.sender);
        require(proposal.pass(validator), 'You must be authorized first');

        if (validatorInfo[validator].feeAddr != feeAddr) {
            validatorInfo[validator].feeAddr = feeAddr;
        }

        validatorInfo[validator].description = Description(moniker, identity, website, email, details);

        emit LogEditValidator(validator, feeAddr, block.timestamp);
        return true;
    }

    function tryActive(address validator) external onlyProposalContract onlyInitialized returns (bool) {
        if (validatorInfo[validator].status == Status.Active) {
            return true;
        }

        tryAddValidatorToHighestSet(validator);
        if (validatorInfo[validator].status == Status.Jailed) {
            require(punish.cleanPunishRecord(validator), 'clean failed');
        }
        validatorInfo[validator].status = Status.Active;

        emit LogActive(validator, block.timestamp);

        return true;
    }

    // feeAddr can withdraw profits of it's validator
    function withdrawProfits(address validator) external returns (bool) {
        address payable feeAddr = payable(msg.sender);
        require(validatorInfo[validator].status != Status.NotExist, 'Validator not exist');
        require(validatorInfo[validator].feeAddr == feeAddr, 'You are not the fee receiver of this validator');
        require(
            validatorInfo[validator].lastWithdrawProfitsBlock + proposal.withdrawProfitPeriod() <= block.number,
            'You must wait enough blocks to withdraw your profits after latest withdraw of this validator'
        );
        uint256 aacIncoming = validatorInfo[validator].aacIncoming;
        require(aacIncoming > 0, "You don't have any profits");

        // update info
        validatorInfo[validator].aacIncoming = 0;
        validatorInfo[validator].lastWithdrawProfitsBlock = block.number;

        // send profits to fee address
        if (aacIncoming > 0) {
            feeAddr.transfer(aacIncoming);
        }

        emit LogWithdrawProfits(validator, feeAddr, aacIncoming, block.timestamp);

        return true;
    }

    // distributeBlockReward distributes block reward to all active validators
    function distributeBlockReward() external payable onlyMiner onlyInitialized {
        // Check if block reward has already been distributed for this block
        if (operationsDone[block.number][uint8(Operations.Distribute)] == true) {
            return; // Silently return to avoid consensus issues
        }
        
        // Set distributed flag immediately to prevent reentrancy
        operationsDone[block.number][uint8(Operations.Distribute)] = true;
        
        address val = msg.sender;
        uint256 hb = msg.value;

        // never reach this
        if (validatorInfo[val].status == Status.NotExist) {
            return;
        }

        // Jailed validator can't get profits.
        addProfitsToActiveValidators(hb, address(0));

        emit LogDistributeBlockReward(val, hb, block.timestamp);
    }

    function updateActiveValidatorSet(address[] memory newSet, uint256 epoch)
        public
        onlyMiner
        onlyInitialized
        onlyBlockEpoch(epoch)
    {
        // Check if validators have already been updated for this block
        if (operationsDone[block.number][uint8(Operations.UpdateValidators)] == true) {
            return; // Silently return to avoid consensus issues
        }
        
        // Set updated flag immediately to prevent reentrancy
        operationsDone[block.number][uint8(Operations.UpdateValidators)] = true;
        
        require(newSet.length > 0, 'Validator set empty!');

        currentValidatorSet = newSet;

        emit LogUpdateValidator(newSet);
    }

    /**
     * @dev Update validator set based on staking (JPoSA mechanism)
     * This function is called by the consensus engine to update validators based on stake
     */
    function updateValidatorSetByStake(uint256 epoch)
        public
        onlyMiner
        onlyInitialized
        onlyBlockEpoch(epoch)
    {
        // Check if validators have already been updated for this block
        if (operationsDone[block.number][uint8(Operations.UpdateValidators)] == true) {
            return; // Silently return to avoid consensus issues
        }
        
        // Set updated flag immediately to prevent reentrancy
        operationsDone[block.number][uint8(Operations.UpdateValidators)] = true;
        
        // Get top validators from staking contract
        address[] memory topValidators = staking.getTopValidators(21); // Max 21 validators
        require(topValidators.length > 0, 'No staked validators available');
        
        // Update both current and highest validator sets
        currentValidatorSet = topValidators;
        highestValidatorsSet = topValidators;
        
        emit LogUpdateValidator(topValidators);
    }

    function removeValidator(address val) external onlyPunishContract {
        removeValidatorInternal(val);
    }

    function tryRemoveValidator(address val) external onlyProposalContract {
        removeValidatorInternal(val);
    }

    function removeValidatorInternal(address val) private {
        uint256 hb = validatorInfo[val].aacIncoming;
        validatorInfo[val].status = Status.Jailed;

        tryRemoveValidatorIncoming(val);

        // remove the validator out of active set
        // Note: the jailed validator may in active set if there is only one validator exists
        if (highestValidatorsSet.length > 1) {
            tryRemoveValidatorInHighestSet(val);

            // call proposal contract to set unpass.
            // you have to repropose to be a validator.
            proposal.setUnpassed(val);
            emit LogRemoveValidator(val, hb, block.timestamp);
        }
    }

    function removeValidatorIncoming(address val) external onlyPunishContract {
        tryRemoveValidatorIncoming(val);
    }

    function getValidatorDescription(address val)
        public
        view
        returns (
            string memory,
            string memory,
            string memory,
            string memory,
            string memory
        )
    {
        Validator memory v = validatorInfo[val];

        return (
            v.description.moniker,
            v.description.identity,
            v.description.website,
            v.description.email,
            v.description.details
        );
    }

    function getValidatorInfo(address val)
        public
        view
        returns (
            address payable,
            Status,
            uint256,
            uint256,
            uint256
        )
    {
        Validator memory v = validatorInfo[val];

        return (v.feeAddr, v.status, v.aacIncoming, v.totalJailedHb, v.lastWithdrawProfitsBlock);
    }

    function getActiveValidators() public view returns (address[] memory) {
        return currentValidatorSet;
    }

    function isActiveValidator(address who) public view returns (bool) {
        for (uint256 i = 0; i < currentValidatorSet.length; i++) {
            if (currentValidatorSet[i] == who) {
                return true;
            }
        }

        return false;
    }

    function isTopValidator(address who) public view returns (bool) {
        for (uint256 i = 0; i < highestValidatorsSet.length; i++) {
            if (highestValidatorsSet[i] == who) {
                return true;
            }
        }

        return false;
    }

    function getTopValidators() public view returns (address[] memory) {
        return highestValidatorsSet;
    }

    function validateDescription(
        string memory moniker,
        string memory identity,
        string memory website,
        string memory email,
        string memory details
    ) public pure returns (bool) {
        require(bytes(moniker).length <= 70, 'Invalid moniker length');
        require(bytes(identity).length <= 3000, 'Invalid identity length');
        require(bytes(website).length <= 140, 'Invalid website length');
        require(bytes(email).length <= 140, 'Invalid email length');
        require(bytes(details).length <= 280, 'Invalid details length');

        return true;
    }

    function tryAddValidatorToHighestSet(address val) internal {
        // do nothing if you are already in highestValidatorsSet set
        for (uint256 i = 0; i < highestValidatorsSet.length; i++) {
            if (highestValidatorsSet[i] == val) {
                return;
            }
        }

        highestValidatorsSet.push(val);
        emit LogAddToTopValidators(val, block.timestamp);
    }

    function tryRemoveValidatorIncoming(address val) private {
        // do nothing if validator not exist(impossible)
        if (validatorInfo[val].status == Status.NotExist || currentValidatorSet.length <= 1) {
            return;
        }

        uint256 hb = validatorInfo[val].aacIncoming;
        if (hb > 0) {
            addProfitsToActiveValidators(hb, val);
            // for display purpose
            totalJailedHb = totalJailedHb.add(hb);
            validatorInfo[val].totalJailedHb = validatorInfo[val].totalJailedHb.add(hb);

            validatorInfo[val].aacIncoming = 0;
        }

        emit LogRemoveValidatorIncoming(val, hb, block.timestamp);
    }

    // add profits to all validators by stake percent except the punished validator or jailed validator
    function addProfitsToActiveValidators(uint256 totalReward, address punishedVal) private {
        if (totalReward == 0) {
            return;
        }

        uint256 rewardValsLen = getRewardLen(punishedVal);
        if (rewardValsLen == 0) {
            return;
        }

        uint256 remain;
        address last;
        uint256 per = totalReward.div(rewardValsLen);
        remain = totalReward.sub(per.mul(rewardValsLen));

        for (uint256 i = 0; i < currentValidatorSet.length; i++) {
            address val = currentValidatorSet[i];
            if (validatorInfo[val].status != Status.Jailed && val != punishedVal) {
                validatorInfo[val].aacIncoming = validatorInfo[val].aacIncoming.add(per);

                last = val;
            }
        }

        if (remain > 0 && last != address(0)) {
            validatorInfo[last].aacIncoming = validatorInfo[last].aacIncoming.add(remain);
        }
    }

    function getRewardLen(address punishedVal) private view returns (uint256) {
        uint256 l;
        for (uint256 i = 0; i < currentValidatorSet.length; i++) {
            address val = currentValidatorSet[i];
            if (validatorInfo[val].status != Status.Jailed && val != punishedVal) {
                l++;
            }
        }
        return l;
    }

    function tryRemoveValidatorInHighestSet(address val) private {
        for (
            uint256 i = 0;
            // ensure at least one validator exist
            i < highestValidatorsSet.length && highestValidatorsSet.length > 1;
            i++
        ) {
            if (val == highestValidatorsSet[i]) {
                // remove it
                if (i != highestValidatorsSet.length - 1) {
                    highestValidatorsSet[i] = highestValidatorsSet[highestValidatorsSet.length - 1];
                }

                highestValidatorsSet.pop();
                emit LogRemoveFromTopValidators(val, block.timestamp);

                break;
            }
        }
    }
}
