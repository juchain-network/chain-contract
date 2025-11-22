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
    
    // Track consecutive normal epochs for validators to reset violation count
    // consecutiveNormalEpochs[validator] = N: validator has run normally for N consecutive epochs
    // When N >= RESET_VIOLATION_THRESHOLD (10), reset violation count
    mapping(address => uint256) public consecutiveNormalEpochs;
    // Threshold for resetting violation count
    // Design Intent: Validators must run normally for 10 consecutive epochs to reset their violation count.
    // This ensures that validators demonstrate consistent good behavior before getting a fresh start.
    // If a validator is removed or jailed, consecutiveNormalEpochs is reset to 0, requiring them to
    // start over. This is intentional - it encourages stable, long-term participation.
    uint256 public constant RESET_VIOLATION_THRESHOLD = 10; // Run normally for 10 epochs to reset violation count

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
     *   - Proposal passes (by Proposal contract)
     *   - Validator is unjailed (by Staking contract)
     * @notice Validator should already be unjailed before calling this function
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
    function withdrawProfits(address validator) external returns (bool) {
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

        // update info
        validatorInfo[validator].aacIncoming = 0;
        validatorInfo[validator].lastWithdrawProfitsBlock = block.number;

        // send profits to fee address
        if (aacIncoming > 0) {
            (bool success, ) = feeAddr.call{value: aacIncoming}("");
            require(success, "Transfer failed");
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

    /**
     * @dev Update validator set based on staking (JPoSA mechanism)
     * This function is called by the consensus engine to update validators based on stake
     * @param epoch Current epoch number
     * @return Updated validator list (based on current state, excludes jailed validators)
     */
    function updateValidatorSetByStake(uint256 epoch)
        public
        onlyMiner
        onlyInitialized
        onlyBlockEpoch(epoch)
        returns (address[] memory)
    {
        // Check if validators have already been updated for this block
        if (operationsDone[block.number][uint8(Operations.UpdateValidators)] == true) {
            // Return current validator set if already updated
            return currentValidatorSet;
        }
        
        // Clean up previous block's data to save storage
        // This prevents storage accumulation while maintaining reentrancy protection
        // Note: We only need to track the current block, historical data is never accessed
        if (block.number > 0) {
            delete operationsDone[block.number - 1][uint8(Operations.UpdateValidators)];
        }
        
        // Set updated flag immediately to prevent reentrancy
        operationsDone[block.number][uint8(Operations.UpdateValidators)] = true;
        
        // Get top validators from staking contract (based on current state)
        // This will automatically exclude jailed validators since getTopValidators checks jail status
        address[] memory topValidators = staking.getTopValidators(); // Returns up to MAX_VALIDATORS
        require(topValidators.length > 0, 'No staked validators available');
        
        // Update both current and highest validator sets
        currentValidatorSet = topValidators;
        highestValidatorsSet = topValidators;
        
        // Track consecutive normal epochs for validators to reset violation count
        _updateConsecutiveNormalEpochs(topValidators);
        
        emit LogUpdateValidator(topValidators);
        
        return topValidators;
    }
    
    /**
     * @dev Update consecutive normal epochs for validators
     * @param activeValidators Current active validators list
     * @notice This function tracks validators that run normally for consecutive epochs
     * @notice When a validator runs normally for RESET_VIOLATION_THRESHOLD epochs, reset violation count
     * @notice Design Intent: Only validators that run continuously for 10 epochs can reset their violation count.
     *         If a validator is removed or jailed, their consecutive count is reset to 0, requiring them to
     *         start over. This encourages stable, long-term participation and discourages frequent violations.
     */
    function _updateConsecutiveNormalEpochs(address[] memory activeValidators) private {
        // Track which validators are in the new active set
        for (uint256 i = 0; i < activeValidators.length; i++) {
            address validator = activeValidators[i];
            // Increment consecutive normal epochs for active validators
            consecutiveNormalEpochs[validator] = consecutiveNormalEpochs[validator] + 1;
            
            // If validator has run normally for enough epochs, reset violation count
            if (consecutiveNormalEpochs[validator] >= RESET_VIOLATION_THRESHOLD) {
                try proposal.resetViolationCount(validator) {} catch {}
                // Reset counter after resetting violation count
                consecutiveNormalEpochs[validator] = 0;
            }
        }
        
        // Reset consecutive normal epochs for validators not in active set
        // (they were removed or jailed)
        // Design Intent: This ensures that validators must run continuously without interruption
        // to reset their violation count. If they are removed or jailed, they must start over.
        for (uint256 i = 0; i < currentValidatorSet.length; i++) {
            address validator = currentValidatorSet[i];
            bool isActive = false;
            for (uint256 j = 0; j < activeValidators.length; j++) {
                if (activeValidators[j] == validator) {
                    isActive = true;
                    break;
                }
            }
            if (!isActive) {
                // Validator was removed or jailed, reset consecutive normal epochs
                // This is intentional - validators must demonstrate continuous good behavior
                consecutiveNormalEpochs[validator] = 0;
            }
        }
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
     * @return status Status (deprecated, use isValidatorJailed/isValidatorActive instead)
     * @return aacIncoming Accumulated transaction fee income
     * @return totalJailedHb Total jailed income
     * @return lastWithdrawProfitsBlock Last withdraw profits block
     * @notice Status field is deprecated. Use isValidatorJailed() and isValidatorActive() instead.
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
     * @dev Get active validators list (excluding jailed validators)
     * @notice Returns validators from currentValidatorSet that are not jailed
     * @notice currentValidatorSet represents validators that are actually active in consensus
     * @notice This ensures that jailed validators are excluded immediately, even before next epoch
     * @return Array of active validators from currentValidatorSet (excluding jailed validators)
     */
    function getActiveValidators() public view returns (address[] memory) {
        // Always use currentValidatorSet as base, which represents validators actually active in consensus
        // Filter out jailed validators (if Staking contract is available)
        if (address(staking) != address(0) && initialized) {
            // POSA mode: Filter out jailed validators from currentValidatorSet
            // Count non-jailed validators
            uint256 activeCount = 0;
            for (uint256 i = 0; i < currentValidatorSet.length; i++) {
                if (!staking.isValidatorJailed(currentValidatorSet[i])) {
                    activeCount++;
                }
            }
            
            // Create filtered array
            address[] memory filteredValidators = new address[](activeCount);
            uint256 index = 0;
            for (uint256 i = 0; i < currentValidatorSet.length; i++) {
                if (!staking.isValidatorJailed(currentValidatorSet[i])) {
                    filteredValidators[index] = currentValidatorSet[i];
                    index++;
                }
            }
            return filteredValidators;
        }
        
        // POA mode without Staking: return currentValidatorSet as-is
        return currentValidatorSet;
    }

    /**
     * @dev Get count of active validators (excluding jailed validators)
     * @notice Returns count of validators in currentValidatorSet that are not jailed
     * @notice More efficient than getActiveValidators().length as it doesn't create an array
     * @return Count of active validators in currentValidatorSet (excluding jailed validators)
     */
    function getActiveValidatorCount() public view returns (uint256) {
        // Always count non-jailed validators in currentValidatorSet
        // currentValidatorSet represents validators that are actually active in consensus
        // getTopValidators() may include newly registered validators that haven't entered currentValidatorSet yet
        if (address(staking) != address(0) && initialized) {
            // POSA mode: Filter out jailed validators from currentValidatorSet
            uint256 count = 0;
            for (uint256 i = 0; i < currentValidatorSet.length; i++) {
                if (!staking.isValidatorJailed(currentValidatorSet[i])) {
                    count++;
                }
            }
            return count;
        }
        
        // POA mode without Staking: return currentValidatorSet length
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
        (uint256 selfStake, , , , ) = staking.getValidatorInfo(validator);
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
     * @dev Get top validators (POA mode - returns cached highestValidatorsSet)
     * @return Top validators list from cached set
     */
    function getTopValidators() public view returns (address[] memory) {
        return highestValidatorsSet;
    }

    /**
     * @dev Get top validators by stake (POSA mode - calls Staking contract)
     * @return Top validators list based on stake (up to MAX_VALIDATORS)
     * @notice This function provides a unified interface for POSA mode
     *         It calls Staking.getTopValidators() to get real-time validator list
     */
    function getTopValidatorsByStake() external view returns (address[] memory) {
        return staking.getTopValidators();
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


    function tryRemoveValidatorIncoming(address val) private {
        // do nothing if validator not exist or only one validator
        if (!this.isValidatorExist(val) || currentValidatorSet.length <= 1) {
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
            // Check if validator is not jailed (from Staking contract) and not punished
            if (!staking.isValidatorJailed(val) && val != punishedVal) {
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
            // Check if validator is not jailed (from Staking contract) and not punished
            if (!staking.isValidatorJailed(val) && val != punishedVal) {
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
