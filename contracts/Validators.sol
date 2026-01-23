// SPDX-License-Identifier: MIT

pragma solidity ^0.8.29;

import {Params} from "./Params.sol";
import {IProposal} from "./IProposal.sol";
import {IPunish} from "./IPunish.sol";
import {IStaking} from "./IStaking.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import {IValidators} from "./IValidators.sol";

contract Validators is Params, ReentrancyGuard, IValidators {

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
    IProposal proposal;
    IPunish punish;
    IStaking staking;

    uint256 private constant PENDING_EXECUTION_LIMIT = 5;

    enum Operations {Distribute, UpdateValidators}
    // Record the operations is done or not.
    mapping(uint256 => mapping(uint8 => bool)) operationsDone;

    uint256 public revision;
    uint256[50] private __gap;

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
        _onlyNotRewarded();
        _;
    }

    function _onlyNotRewarded() internal view {
        require(!operationsDone[block.number][uint8(Operations.Distribute)], "Block is already rewarded");
    }

    /**
     * @dev Initializes the Validators contract with validators and dependencies.
     * @param vals Array of initial validator addresses.
     * @param proposal_ Address of the Proposal contract.
     * @param punish_ Address of the Punish contract.
     * @param staking_ Address of the Staking contract.
     */
    function initialize(
        address[] calldata vals,
        address proposal_,
        address punish_,
        address staking_
    ) external onlyNotInitialized {
        require(proposal_ != address(0), "Invalid proposal address");
        require(punish_ != address(0), "Invalid punish address");
        require(staking_ != address(0), "Invalid staking address");
        
        proposal = IProposal(proposal_);
        punish = IPunish(punish_);
        staking = IStaking(staking_);
        _initializeEpoch(proposal.epoch());

        for (uint256 i = 0; i < vals.length; i++) {
            require(vals[i] != address(0), "Invalid validator address");

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

        revision = 1;
        initialized = true;
    }

    /**
     * @dev Reinitialize for upgrades (v2).
     * @notice Only miner can call. Can be called once.
     */
    function reinitializeV2() external onlyInitialized onlyMiner {
        require(revision < 2, "Already reinitialized");
        revision = 2;
    }

    /**
     * @dev Creates or edits a validator's information.
     * @param feeAddr Address where validator fees will be sent.
     * @param moniker Validator's display name.
     * @param identity Validator's identity (e.g., Keybase ID).
     * @param website Validator's website URL.
     * @param email Validator's email address.
     * @param details Additional details about the validator.
     * @return bool Returns true if the operation was successful.
     */
    function createOrEditValidator(
        address payable feeAddr,
        string calldata moniker,
        string calldata identity,
        string calldata website,
        string calldata email,
        string calldata details
    ) external onlyInitialized returns (bool) {
        require(feeAddr != address(0), "Invalid fee address");
        require(validateDescription(moniker, identity, website, email, details), "Invalid description");
        address payable validator = payable(msg.sender);
        require(proposal.pass(validator), "You must be authorized first");

        if (validatorInfo[validator].feeAddr != feeAddr) {
            validatorInfo[validator].feeAddr = feeAddr;
        }

        validatorInfo[validator].description = Description({moniker: moniker, identity: identity, website: website, email: email, details: details});

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
    function tryActive(address validator) external onlyInitialized onlyNotEpoch nonReentrant returns (bool) {
        require(
            msg.sender == address(proposal) || msg.sender == address(staking),
            "Only Proposal or Staking contract can call"
        );
        
        // Add validator to highest validators set if not already there
        _tryAddValidatorToHighestSet(validator);
        
        // Only clean punish record if validator was previously jailed
        (, , , , bool isJailed, , , , , ) = staking.getValidatorInfo(validator);
        if (isJailed) {
            require(punish.cleanPunishRecord(validator), "Punish record clean failed");
        }
        
        emit LogActive(validator, block.timestamp);

        return true;
    }

    /**
     * @dev Allows the fee address to withdraw profits for a validator.
     * @param validator Address of the validator whose profits are being withdrawn.
     * @return bool Returns true if the operation was successful.
     * @notice Only the validator's designated fee address can call this function.
     * @notice There's a minimum waiting period between withdrawals.
     */
    function withdrawProfits(address validator) external nonReentrant returns (bool) {
        address payable feeAddr = payable(msg.sender);
        // Check if validator exists (has staked) from Staking contract
        require(this.isValidatorExist(validator), "Validator does not exist");
        require(validatorInfo[validator].feeAddr == feeAddr, "You are not the fee receiver of this validator");
        require(
            validatorInfo[validator].lastWithdrawProfitsBlock + proposal.withdrawProfitPeriod() <= block.number,
            "You must wait enough blocks to withdraw your profits after latest withdraw of this validator"
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

    /**
     * @dev Distributes block reward to the block producer (validator).
     * @notice If the block producer is jailed, the reward is distributed to other active validators.
     * @notice Only the miner can call this function.
     * @notice Block reward is passed via msg.value.
     */
    function distributeBlockReward() external payable onlyMiner onlyInitialized onlyNotRewarded {
        // Check is now handled by onlyNotRewarded modifier
        
        // Clean up previous block's data to save storage
        // This prevents storage accumulation while maintaining reentrancy protection
        // Note: We only need to track the current block, historical data is never accessed
        if (block.number > 0) {
            delete operationsDone[block.number - 1][uint8(Operations.Distribute)];
        }
        
        // Set distributed flag immediately to prevent reentrancy
        operationsDone[block.number][uint8(Operations.Distribute)] = true;

        if (block.number % epoch != 0) {
            try punish.executePending(PENDING_EXECUTION_LIMIT) {} catch {}
        }
        
        address val = msg.sender;
        uint256 hb = msg.value;

        // Check if validator exists (has staked) from Staking contract
        // Note: This branch should be unreachable in normal operation
        // as block producers must be validators according to consensus rules
        // This check is added for code robustness
        if (!this.isValidatorExist(val)) {
            return;
        }
        staking.updateLastActiveBlock(val);

        // Check if the block producer is jailed
        if (staking.isValidatorJailed(val)) {
            // If jailed, distribute reward to other active validators (excluding the jailed producer)
            addProfitsToActiveValidators(hb, val);
        } else {
            // If not jailed, reward goes directly to the block producer
            validatorInfo[val].aacIncoming = validatorInfo[val].aacIncoming + hb;
        }

        emit LogDistributeBlockReward(val, hb, block.timestamp);
    }

    /**
     * @dev Updates the active validator set.
     * @param newSet Array of addresses for the new validator set.
     * @param epoch Epoch number for which the update is happening.
     * @notice Only the miner can call this function.
     * @notice Validators are updated at specific epoch boundaries.
     */
    function updateActiveValidatorSet(address[] memory newSet, uint256 epoch)
        public
        onlyMiner
        onlyInitialized
        onlyBlockEpoch(epoch)
    {
        // Check if validators have already been updated for this block
        if (operationsDone[block.number][uint8(Operations.UpdateValidators)]) {
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
        
        address[] memory expected = staking.getTopValidators(highestValidatorsSet);
        require(expected.length > 0, "Validator set empty!");
        uint256 maxValidators = proposal.maxValidators();
        if (maxValidators > CONSENSUS_MAX_VALIDATORS) {
            maxValidators = CONSENSUS_MAX_VALIDATORS;
        }
        require(expected.length <= maxValidators, "Validator set too large");
        require(newSet.length == expected.length, "Validator set mismatch");
        _validateValidatorSet(newSet);
        _validateValidatorSet(expected);
        _requireSameSet(newSet, expected);

        currentValidatorSet = newSet;

        emit LogUpdateValidator(newSet);
    }

    /**
     * @dev Removes a validator from the active set.
     * @param val Address of the validator to remove.
     * @notice Only the Punish contract can call this function.
     */
    function removeValidator(address val) external onlyPunishContract onlyNotEpoch nonReentrant {
        removeValidatorInternal(val);
    }

    /**
     * @dev Tries to remove a validator from the active set.
     * @param val Address of the validator to remove.
     * @notice Only the Proposal contract can call this function.
     */
    function tryRemoveValidator(address val) external onlyProposalContract onlyNotEpoch nonReentrant {
        // Jail validator first to ensure no more rewards are distributed
        if (getVotingValidatorCount() > 1) {
            staking.jailValidator(val, proposal.validatorUnjailPeriod());
        }
        // Remove validator from active set
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
            require(proposal.setUnpassed(val), "Validator unpass set failed");
            emit LogRemoveValidator(val, hb, block.timestamp);
        }
    }

    /**
     * @dev Removes a validator from the incoming validator set.
     * @param val Address of the validator to remove.
     * @notice Only the Punish contract can call this function.
     */
    function removeValidatorIncoming(address val) external onlyPunishContract onlyNotEpoch nonReentrant {
        tryRemoveValidatorIncoming(val);
    }

    /**
     * @dev Gets the description of a validator.
     * @param val Address of the validator.
     * @return moniker Validator's display name.
     * @return identity Validator's identity (e.g., Keybase ID).
     * @return website Validator's website URL.
     * @return email Validator's email address.
     * @return details Additional details about the validator.
     */
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
        // Check if validator has staked
        (uint256 selfStake, , , , , , , , bool isRegistered, ) = staking.getValidatorInfo(val);
        
        // Check if validator is not registered or has no stake
        if (!isRegistered || selfStake == 0) {
            calculatedStatus = Status.NotExist;
        } else if (staking.isValidatorJailed(val)) {
            calculatedStatus = Status.Jailed;
        } else {
            calculatedStatus = Status.Active;
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
     * @dev Get active validators list with their total stake amounts
     * @notice Returns validators from currentValidatorSet along with their total stake (selfStake + totalDelegated)
     * @notice currentValidatorSet is only updated at epoch blocks, so jailed validators
     *         remain in the set until the next epoch transition
     * @notice This aligns with consensus behavior where jailed validators can still
     *         produce blocks in the current epoch
     * @return validators Array of validators in currentValidatorSet
     * @return totalStakes Array of total stake amounts for each validator (selfStake + totalDelegated)
     */
    function getActiveValidatorsWithStakes() public view returns (address[] memory validators, uint256[] memory totalStakes) {
        uint256 length = currentValidatorSet.length;
        validators = new address[](length);
        totalStakes = new uint256[](length);
        
        for (uint256 i = 0; i < length; i++) {
            address validator = currentValidatorSet[i];
            validators[i] = validator;
            
            // Get validator info from Staking contract
            (uint256 selfStake, uint256 totalDelegated, , , , , , , , ) = 
                staking.getValidatorInfo(validator);
            
            // Calculate total stake (selfStake + totalDelegated)
            totalStakes[i] = selfStake + totalDelegated;
        }
        
        return (validators, totalStakes);
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

    /**
     * @dev Get count of voting validators (active and not jailed)
     * @return Count of validators eligible to vote
     */
    function getVotingValidatorCount() public view returns (uint256) {
        uint256 currentSetLength = currentValidatorSet.length;
        uint256 count = 0;
        for (uint256 i = 0; i < currentSetLength; i++) {
            address validator = currentValidatorSet[i];
            if (!staking.isValidatorJailed(validator)) {
                count++;
            }
        }
        return count;
    }

    /**
     * @dev Checks if an address is an active validator.
     * @param who Address to check.
     * @return bool Returns true if the address is in the current validator set.
     */
    function isActiveValidator(address who) public view returns (bool) {
        uint256 currentSetLength = currentValidatorSet.length;
        for (uint256 i = 0; i < currentSetLength; i++) {
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
     * @notice Active means validator is in currentValidatorSet and not jailed
     * @notice Validators in currentValidatorSet remain there until next epoch even if jailed
     * @notice This function checks both current participation and jail status
     */
    function isValidatorActive(address validator) external view returns (bool) {
        // Check if validator is in currentValidatorSet (actively participating in consensus)
        if (!isActiveValidator(validator)) {
            return false;
        }
        
        // Check if validator is jailed
        // Jailed validators can still be in currentValidatorSet until next epoch
        // but should not be considered active for most purposes
        return !staking.isValidatorJailed(validator);
    }

    /**
     * @dev Check if validator exists (has staked)
     * @param validator Validator address
     * @return Whether validator exists (has staked)
     */
    function isValidatorExist(address validator) external view returns (bool) {
        (uint256 selfStake, , , , , , , , , ) = 
            staking.getValidatorInfo(validator);
        return selfStake > 0;
    }

    /**
     * @dev Checks if an address is a top validator.
     * @param who Address to check.
     * @return bool Returns true if the address is in the highest validators set.
     */
    function isTopValidator(address who) public view returns (bool) {
        uint256 highestSetLength = highestValidatorsSet.length;
        for (uint256 i = 0; i < highestSetLength; i++) {
            if (highestValidatorsSet[i] == who) {
                return true;
            }
        }

        return false;
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

    /**
     * @dev Get highest validators set (returns cached highestValidatorsSet)
     * @return Highest validators list from cached set
     */
    function getHighestValidators() public view returns (address[] memory) {
        return highestValidatorsSet;
    }

    /**
     * @dev Validates a validator's description fields.
     * @param moniker Validator's display name.
     * @param identity Validator's identity (e.g., Keybase ID).
     * @param website Validator's website URL.
     * @param email Validator's email address.
     * @param details Additional details about the validator.
     * @return bool Returns true if all fields are valid.
     */
    function validateDescription(
        string memory moniker,
        string memory identity,
        string memory website,
        string memory email,
        string memory details
    ) public pure returns (bool) {
        require(bytes(moniker).length <= 70, "Invalid moniker length");
        require(bytes(identity).length <= 3000, "Invalid identity length");
        require(bytes(website).length <= 140, "Invalid website length");
        require(bytes(email).length <= 140, "Invalid email length");
        require(bytes(details).length <= 280, "Invalid details length");

        return true;
    }

    function _tryAddValidatorToHighestSet(address val) internal {
        // do nothing if you are already in highestValidatorsSet set
        uint256 highestSetLength = highestValidatorsSet.length;
        for (uint256 i = 0; i < highestSetLength; i++) {
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
    function tryAddValidatorToHighestSet(address validator) external onlyStakingContract onlyInitialized onlyNotEpoch {
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
    function removeFromHighestSet(address validator) external onlyStakingContract onlyInitialized onlyNotEpoch nonReentrant {
        // Check if validator is in highestValidatorsSet
        bool isInSet = false;
        uint256 highestSetLength = highestValidatorsSet.length;
        for (uint256 i = 0; i < highestSetLength; i++) {
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
        bool success = proposal.setUnpassed(validator);
        require(success, "Failed to update validator status");
        
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

        uint256 currentSetLength = currentValidatorSet.length;
        
        // Cache jailed status for all validators in a single pass
        bool[] memory isJailed = new bool[](currentSetLength);
        uint256 validValidatorCount = 0;
        
        for (uint256 i = 0; i < currentSetLength; i++) {
            address val = currentValidatorSet[i];
            bool jailed = staking.isValidatorJailed(val);
            isJailed[i] = jailed;
            if (val != punishedVal && !jailed) {
                validValidatorCount++;
            }
        }

        // Check if there are any valid validators to distribute rewards to
        // Note: This branch should be unreachable in normal operation
        // as there should always be at least one active validator in the network
        // This check is added for code robustness to prevent division by zero
        if (validValidatorCount == 0) {
            return;
        }

        // Calculate per-validator reward without divide-before-multiply
        uint256 per = totalReward / validValidatorCount;
        uint256 remainder = totalReward % validValidatorCount;

        // Distribute rewards using cached jailed status
        for (uint256 i = 0; i < currentSetLength; i++) {
            address val = currentValidatorSet[i];
            if (val != punishedVal && !isJailed[i]) {
                uint256 reward = per;
                if (remainder > 0) {
                    reward += 1;
                    remainder -= 1;
                }
                validatorInfo[val].aacIncoming += reward;
            }
        }
    }

    function _validateValidatorSet(address[] memory set) private pure {
        uint256 length = set.length;
        for (uint256 i = 0; i < length; i++) {
            address validator = set[i];
            require(validator != address(0), "Invalid validator address");
            for (uint256 j = i + 1; j < length; j++) {
                require(set[j] != validator, "Duplicate validator");
            }
        }
    }

    function _requireSameSet(address[] memory left, address[] memory right) private pure {
        uint256 length = left.length;
        for (uint256 i = 0; i < length; i++) {
            bool found = false;
            for (uint256 j = 0; j < length; j++) {
                if (left[i] == right[j]) {
                    found = true;
                    break;
                }
            }
            require(found, "Validator set mismatch");
        }
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
