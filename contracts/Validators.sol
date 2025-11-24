// SPDX-License-Identifier: MIT

pragma solidity ^0.8.20;

import {Params} from './Params.sol';
import {Proposal} from './Proposal.sol';
import {Punish} from './Punish.sol';
import {Staking} from './Staking.sol';
import {ReentrancyGuard} from './library/ReentrancyGuard.sol';

contract Validators is Params, ReentrancyGuard {

    /**
     * @dev Validator status enum
     * @notice Status is managed by Staking contract, not stored in Validator struct
     * @notice This enum is used as return type for getValidatorInfo() for backward compatibility
     * @notice Status is calculated dynamically by querying Staking contract
     */
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
        // Note: Status is managed by Staking contract, not stored here
        // Status is calculated dynamically in getValidatorInfo() for backward compatibility
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
            // Important: Initialize validator info for genesis validators
            // Status is now managed by Staking contract, we only set feeAddr here
            // Note: Genesis validators are pre-registered in Staking contract with default stake
            // They are activated by default and don't need separate staking to start producing blocks
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

    /**
     * @dev Activate a validator (called by Proposal or Staking contract)
     * @param validator Validator address to activate
     * @notice This function is called when:
     *   - Validator registers (by Staking contract) - register = activate
     *   - Validator is unjailed (by Staking contract) - called before unjailing state change
     * @notice This function does NOT check jailed status, only checks if validator is in currentValidatorSet
     * @notice Can be called even if validator is still jailed (e.g., in unjailValidator before state change)
     */
    function tryActive(address validator) external onlyInitialized returns (bool) {
        require(
            msg.sender == address(proposal) || msg.sender == address(staking),
            "Only Proposal or Staking contract can call"
        );
        
        // Check if validator is already active (from Staking contract)
        (bool isActive, ) = staking.getValidatorStatus(validator);
        if (isActive) {
            return true;
        }

        // Add validator to highest validators set if not already there
        _tryAddValidatorToHighestSet(validator);
        
        // Clean punish record if validator was previously jailed
        // This clears any missed blocks counter from previous punishments
        punish.cleanPunishRecord(validator);
        
        // Note: Status is now managed by Staking contract, we don't set it here

        emit LogActive(validator, block.timestamp);

        return true;
    }

    // feeAddr can withdraw profits of it's validator
    function withdrawProfits(address validator) external nonReentrant returns (bool) {
        address payable feeAddr = payable(msg.sender);
        // Check if validator exists (has staked) from Staking contract
        require(this.isValidatorExist(validator), 'Validator not exist');
        require(validatorInfo[validator].feeAddr == feeAddr, 'You are not the fee receiver of this validator');
        require(
            validatorInfo[validator].lastWithdrawProfitsBlock + proposal.withdrawProfitPeriod() <= block.number,
            'You must wait enough blocks to withdraw your profits after latest withdraw of this validator'
        );
        uint256 aacIncoming = validatorInfo[validator].aacIncoming;
        require(aacIncoming > 0, "You don't have any profits");

        // update info (Effects: update state before external call)
        validatorInfo[validator].aacIncoming = 0;
        validatorInfo[validator].lastWithdrawProfitsBlock = block.number;

        // send profits to fee address (Interactions: external call after state update)
        (bool success, ) = feeAddr.call{value: aacIncoming}("");
        require(success, "Transfer failed");

        emit LogWithdrawProfits(validator, feeAddr, aacIncoming, block.timestamp);

        return true;
    }

    // distributeBlockReward distributes block reward to all active validators
    function distributeBlockReward() external payable onlyMiner onlyInitialized {
        // Check if block reward has already been distributed for this block
        if (operationsDone[block.number][uint8(Operations.Distribute)] == true) {
            return; // Silently return to avoid consensus issues
        }
        
        // Clean up previous block's data to save storage
        // This prevents storage accumulation while maintaining reentrancy protection
        // Note: We only need to track the current block, historical data is never accessed
        if (block.number > 0) {
            delete operationsDone[block.number - 1][uint8(Operations.Distribute)];
        }
        
        // Set distributed flag immediately to prevent reentrancy
        operationsDone[block.number][uint8(Operations.Distribute)] = true;
        
        address val = msg.sender;
        uint256 hb = msg.value;

        // Check if validator exists (has staked) from Staking contract
        if (!this.isValidatorExist(val)) {
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
        
        // Clean up previous block's data to save storage
        // This prevents storage accumulation while maintaining reentrancy protection
        // Note: We only need to track the current block, historical data is never accessed
        if (block.number > 0) {
            delete operationsDone[block.number - 1][uint8(Operations.UpdateValidators)];
        }
        
        // Set updated flag immediately to prevent reentrancy
        operationsDone[block.number][uint8(Operations.UpdateValidators)] = true;
        
        require(newSet.length > 0, 'Validator set empty!');

        currentValidatorSet = newSet;

        emit LogUpdateValidator(newSet);
    }

    function removeValidator(address val) external onlyPunishContract {
        removeValidatorInternal(val);
    }

    function tryRemoveValidator(address val) external onlyProposalContract {
        removeValidatorInternal(val);
    }

    function removeValidatorInternal(address val) private {
        uint256 hb = validatorInfo[val].aacIncoming;
        // Note: Status is now managed by Staking contract
        // jailValidator() is called before removeValidator(), so isJailed should already be set
        // We don't set status here anymore

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

    /**
     * @dev Get validator information
     * @param val Validator address
     * @return feeAddr Fee address
     * @return status Status (calculated dynamically from Staking contract, not stored)
     * @return aacIncoming Accumulated transaction fee income
     * @return totalJailedHb Total jailed income
     * @return lastWithdrawProfitsBlock Last withdraw profits block
     * @notice Status is calculated dynamically by querying Staking contract.
     *         For better performance, use isValidatorJailed() and isValidatorActive() instead.
     */
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
        
        // Calculate status from Staking contract for backward compatibility
        Status calculatedStatus;
        // Priority: Check jailed status first, even if validator doesn't exist in Staking
        // This ensures that jailed validators (including those without stake) are correctly identified
        if (staking.isValidatorJailed(val)) {
            calculatedStatus = Status.Jailed;
        } else if (!this.isValidatorExist(val)) {
            calculatedStatus = Status.NotExist;
        } else {
            (bool isActive, ) = staking.getValidatorStatus(val);
            calculatedStatus = isActive ? Status.Active : Status.NotExist;
        }

        return (v.feeAddr, calculatedStatus, v.aacIncoming, v.totalJailedHb, v.lastWithdrawProfitsBlock);
    }

    /**
     * @dev Get active validators list
     * @notice Returns validators from currentValidatorSet
     * @notice currentValidatorSet is only updated at epoch blocks, so jailed validators
     *         remain in the set until the next epoch transition
     * @notice This aligns with consensus behavior where jailed validators can still
     *         produce blocks in the current epoch
     * @return Array of validators in currentValidatorSet
     */
    function getActiveValidators() public view returns (address[] memory) {
        // Return currentValidatorSet directly - no need to filter jailed validators
        // Jailed validators will be excluded at the next epoch update
        return currentValidatorSet;
    }

    /**
     * @dev Get count of active validators
     * @notice Returns count of validators in currentValidatorSet
     * @notice More efficient than getActiveValidators().length as it doesn't create an array
     * @notice currentValidatorSet is only updated at epoch blocks, so jailed validators
     *         are counted until the next epoch transition
     * @return Count of validators in currentValidatorSet
     */
    function getActiveValidatorCount() public view returns (uint256) {
        // Return currentValidatorSet length directly - no need to filter jailed validators
        // Jailed validators will be excluded at the next epoch update
        return currentValidatorSet.length;
    }

    function isActiveValidator(address who) public view returns (bool) {
        for (uint256 i = 0; i < currentValidatorSet.length; i++) {
            if (currentValidatorSet[i] == who) {
                return true;
            }
        }

        return false;
    }

    /**
     * @dev Check if validator is jailed (query from Staking contract)
     * @param validator Validator address
     * @return Whether validator is currently jailed
     */
    function isValidatorJailed(address validator) external view returns (bool) {
        return staking.isValidatorJailed(validator);
    }

    /**
     * @dev Check if validator is active (query from Staking contract)
     * @param validator Validator address
     * @return Whether validator is active (can participate in consensus)
     */
    function isValidatorActive(address validator) external view returns (bool) {
        (bool isActive, ) = staking.getValidatorStatus(validator);
        return isActive;
    }

    /**
     * @dev Check if validator exists (has staked)
     * @param validator Validator address
     * @return Whether validator exists (has staked)
     */
    function isValidatorExist(address validator) external view returns (bool) {
        (uint256 selfStake, , , , , , , ) = staking.getValidatorInfo(validator);
        return selfStake > 0;
    }

    function isTopValidator(address who) public view returns (bool) {
        for (uint256 i = 0; i < highestValidatorsSet.length; i++) {
            if (highestValidatorsSet[i] == who) {
                return true;
            }
        }

        return false;
    }

    /**
     * @dev Get highest validators set (returns cached highestValidatorsSet)
     * @return Highest validators list from cached set
     */
    function getHighestValidators() public view returns (address[] memory) {
        return highestValidatorsSet;
    }

    /**
     * @dev Get top validators (unified interface for consensus)
     * @notice Calls Staking.getTopValidators() with highestValidatorsSet for sorting by stake
     * @return Top validators list, sorted by stake in POSA
     */
    function getTopValidators() public view returns (address[] memory) {
        // Get highest validators set
        address[] memory highestValidators = highestValidatorsSet;
        
        // Call Staking contract to sort by stake
        return staking.getTopValidators(highestValidators);
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

    function _tryAddValidatorToHighestSet(address val) internal {
        // do nothing if you are already in highestValidatorsSet set
        for (uint256 i = 0; i < highestValidatorsSet.length; i++) {
            if (highestValidatorsSet[i] == val) {
                return;
            }
        }

        highestValidatorsSet.push(val);
        emit LogAddToTopValidators(val, block.timestamp);
    }

    /**
     * @dev Add validator to highest validators set (called by Staking contract)
     * @param validator Validator address to add
     */
    function tryAddValidatorToHighestSet(address validator) external onlyStakingContract onlyInitialized {
        _tryAddValidatorToHighestSet(validator);
    }

    /**
     * @dev Remove validator from highest validators set (called by Staking contract)
     * @param validator Validator address to remove
     * @notice This function is called when:
     *   - Validator resigns via resignValidator()
     *   - Validator exits via emergencyExit() (if they didn't call resignValidator first)
     * @notice It removes validator from highestValidatorsSet and sets pass[validator] = false
     * @notice It does NOT remove transaction fee income (aacIncoming)
     * @notice This is different from removeValidatorInternal() which calls tryRemoveValidatorIncoming()
     * @notice Ensures at least one validator remains in highestValidatorsSet
     */
    function removeFromHighestSet(address validator) external onlyStakingContract onlyInitialized {
        // Check if validator is in highestValidatorsSet
        bool isInSet = false;
        for (uint256 i = 0; i < highestValidatorsSet.length; i++) {
            if (highestValidatorsSet[i] == validator) {
                isInSet = true;
                break;
            }
        }
        
        // If validator is in set, ensure removing them won't leave less than 1 validator
        if (isInSet) {
            require(highestValidatorsSet.length > 1, "Cannot remove: must keep at least one validator in highestValidatorsSet");
            tryRemoveValidatorInHighestSet(validator);
        }
        
        // Set unpassed so validator must repropose to regain validator status
        proposal.setUnpassed(validator);
        
        // Note: We do NOT remove transaction fee income (aacIncoming) here
        // This is different from removeValidatorInternal() which calls tryRemoveValidatorIncoming()
        emit LogRemoveValidator(validator, validatorInfo[validator].aacIncoming, block.timestamp);
    }


    function tryRemoveValidatorIncoming(address val) private {
        // do nothing if validator not exist or only one validator
        if (!this.isValidatorExist(val) || currentValidatorSet.length <= 1) {
            return;
        }

        uint256 hb = validatorInfo[val].aacIncoming;
        if (hb > 0) {
            addProfitsToActiveValidators(hb, val);
            // for display purpose
            totalJailedHb = totalJailedHb + hb;
            validatorInfo[val].totalJailedHb = validatorInfo[val].totalJailedHb + hb;

            validatorInfo[val].aacIncoming = 0;
        }

        emit LogRemoveValidatorIncoming(val, hb, block.timestamp);
    }

    // add profits to all validators by stake percent except the punished validator and jailed validators
    // Note: Jailed validators should not receive rewards, even though they remain in currentValidatorSet until next epoch
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
        uint256 per = totalReward / rewardValsLen;
        remain = totalReward - (per * rewardValsLen);

        for (uint256 i = 0; i < currentValidatorSet.length; i++) {
            address val = currentValidatorSet[i];
            // Exclude the punished validator and jailed validators
            // Jailed validators remain in currentValidatorSet until next epoch, but should not receive rewards
            if (val != punishedVal && !staking.isValidatorJailed(val)) {
                validatorInfo[val].aacIncoming = validatorInfo[val].aacIncoming + per;

                last = val;
            }
        }

        if (remain > 0 && last != address(0)) {
            validatorInfo[last].aacIncoming = validatorInfo[last].aacIncoming + remain;
        }
    }

    function getRewardLen(address punishedVal) private view returns (uint256) {
        uint256 l;
        for (uint256 i = 0; i < currentValidatorSet.length; i++) {
            address val = currentValidatorSet[i];
            // Exclude the punished validator and jailed validators
            // Jailed validators remain in currentValidatorSet until next epoch, but should not receive rewards
            if (val != punishedVal && !staking.isValidatorJailed(val)) {
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
