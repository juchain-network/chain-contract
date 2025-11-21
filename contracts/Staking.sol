// SPDX-License-Identifier: MIT

pragma solidity ^0.8.20;

import {Params} from './Params.sol';
import {SafeMath} from './library/SafeMath.sol';
import {Proposal} from './Proposal.sol';

// Interface for Validators contract to avoid circular dependency
interface IValidators {
    function tryAddValidatorToHighestSet(address validator) external;
    function tryActive(address validator) external returns (bool);
    function isActiveValidator(address who) external view returns (bool);
    function getActiveValidatorCount() external view returns (uint256);
}

/**
 * @title Staking Contract for JPoSA Consensus
 * @dev Implements staking mechanism for JuChain Proof of Stake Authorization
 */
contract Staking is Params {
    using SafeMath for uint256;

    // Minimum staking amount to become a validator
    uint256 public constant MIN_VALIDATOR_STAKE = 10000 ether; // 10,000 JU
    
    // Minimum delegation amount
    uint256 public constant MIN_DELEGATION = 1 ether; // 1 JU
    
    // Maximum validators in active set
    uint256 public constant MAX_VALIDATORS = 21;
    
    // Minimum validators that must always be active
    // Design Intent: Set to 3 for flexibility, allowing the network to continue operating
    // with fewer validators. This provides more operational flexibility while maintaining
    // basic network functionality. Network security and decentralization may be reduced
    // with fewer validators, but this is an acceptable trade-off for operational flexibility.
    uint256 public constant MIN_VALIDATORS = 3;
    
    // Commission rate precision (10000 = 100%)
    uint256 public constant COMMISSION_RATE_BASE = 10000;
    
    // Unbonding period in blocks (approximately 7 days)
    uint256 public constant UNBONDING_PERIOD = 604800; // 7 days * 24 hours * 3600 seconds / 1 second per block
    
    // Maximum number of unbonding entries to process in a single withdrawUnbonded call
    uint256 public constant MAX_UNBONDING_ENTRIES_PER_WITHDRAW = 50;

    struct ValidatorStake {
        uint256 selfStake;          // Validator's own stake
        uint256 totalDelegated;     // Total delegated to this validator
        uint256 commissionRate;     // Commission rate (0-10000, representing 0%-100%)
        uint256 accumulatedRewards; // Accumulated rewards for distribution
        bool isJailed;              // Whether validator is jailed
        uint256 jailUntilBlock;     // Block number until which validator is jailed
    }

    struct Delegation {
        uint256 amount;             // Delegated amount
        uint256 rewardDebt;         // Reward debt for accurate reward calculation
        uint256 unbondingAmount;    // Amount being unbonded
        uint256 unbondingBlock;     // Block when unbonding completes
    }

    struct UnbondingEntry {
        uint256 amount;
        uint256 completionBlock;
    }

    // Validator address => ValidatorStake
    mapping(address => ValidatorStake) public validatorStakes;
    
    // Delegator => Validator => Delegation
    mapping(address => mapping(address => Delegation)) public delegations;
    
    // Delegator => Validator => UnbondingEntry[]
    mapping(address => mapping(address => UnbondingEntry[])) public unbondingDelegations;
    
    // List of all validators (including inactive ones)
    address[] public allValidators;
    
    // Validator address => index in allValidators array
    mapping(address => uint256) public validatorIndex;
    
    // Total staked amount in the system
    uint256 public totalStaked;
    
    // Rewards per share (scaled by 1e18 for precision)
    mapping(address => uint256) public rewardPerShare;

    // System contracts
    IValidators public validatorsContract;
    Proposal public proposalContract;

    event ValidatorRegistered(address indexed validator, uint256 selfStake, uint256 commissionRate);
    event ValidatorUpdated(address indexed validator, uint256 commissionRate);
    event ValidatorStakeWithdrawn(address indexed validator, uint256 amount);
    event ValidatorExited(address indexed validator, uint256 amount);
    event Delegated(address indexed delegator, address indexed validator, uint256 amount);
    event Undelegated(address indexed delegator, address indexed validator, uint256 amount);
    event UnbondingCompleted(address indexed delegator, address indexed validator, uint256 amount);
    event RewardsDistributed(address indexed validator, uint256 amount);
    event RewardsClaimed(address indexed delegator, address indexed validator, uint256 amount);
    event ValidatorJailed(address indexed validator, uint256 jailUntilBlock);
    event ValidatorUnjailed(address indexed validator);

    modifier onlyValidValidator(address validator) {
        require(validatorStakes[validator].selfStake >= MIN_VALIDATOR_STAKE, "Not a valid validator");
        _;
    }

    modifier onlyActiveValidator(address validator) {
        require(validatorStakes[validator].selfStake >= MIN_VALIDATOR_STAKE, "Not a valid validator");
        require(!validatorStakes[validator].isJailed, "Validator is jailed");
        _;
    }

    function initialize(address _validators, address _proposal) external onlyNotInitialized {
        require(_validators != address(0), "Invalid validators address");
        require(_proposal != address(0), "Invalid proposal address");
        
        validatorsContract = IValidators(_validators);
        proposalContract = Proposal(_proposal);
        initialized = true;
    }

    /**
     * @dev Initialize with pre-registered validators (for genesis deployment)
     * @param _validators Validators contract address
     * @param _proposal Proposal contract address
     * @param initialValidators Array of validator addresses to pre-register
     * @param commissionRate Default commission rate for all validators
     */
    function initializeWithValidators(
        address _validators,
        address _proposal,
        address[] calldata initialValidators,
        uint256 commissionRate
    ) external onlyNotInitialized {
        require(_validators != address(0), "Invalid validators address");
        require(_proposal != address(0), "Invalid proposal address");
        require(initialValidators.length > 0, "No validators provided");
        require(commissionRate <= COMMISSION_RATE_BASE, "Invalid commission rate");
        
        validatorsContract = IValidators(_validators);
        proposalContract = Proposal(_proposal);
        
        // Pre-register all initial validators with default stake
        for (uint256 i = 0; i < initialValidators.length; i++) {
            address validator = initialValidators[i];
            require(validator != address(0), "Invalid validator address");
            require(validatorStakes[validator].selfStake == 0, "Validator already exists");
            
            validatorStakes[validator] = ValidatorStake({
                selfStake: MIN_VALIDATOR_STAKE,
                totalDelegated: 0,
                commissionRate: commissionRate,
                accumulatedRewards: 0,
                isJailed: false,
                jailUntilBlock: 0
            });
            
            // Add to validators list
            validatorIndex[validator] = allValidators.length;
            allValidators.push(validator);
            totalStaked = totalStaked.add(MIN_VALIDATOR_STAKE);
        }
        
        initialized = true;
        
        emit ValidatorRegistered(address(0), MIN_VALIDATOR_STAKE, commissionRate); // Genesis event
    }

    /**
     * @dev Register as a validator with self-stake
     * @param commissionRate Commission rate (0-10000, representing 0%-100%)
     */
    function registerValidator(uint256 commissionRate) external payable onlyInitialized {
        require(msg.value >= MIN_VALIDATOR_STAKE, "Insufficient self-stake");
        require(commissionRate <= COMMISSION_RATE_BASE, "Invalid commission rate");
        require(validatorStakes[msg.sender].selfStake == 0, "Already registered");
        require(proposalContract.pass(msg.sender), "Must pass proposal first");
        // Check if proposal is still valid (within 7 days)
        require(proposalContract.isProposalValidForStaking(msg.sender), "Proposal expired, must repropose");
        // Check if validator is jailed (must unjail first before re-registering)
        require(!validatorStakes[msg.sender].isJailed, "Validator is jailed, must unjail first");

        validatorStakes[msg.sender] = ValidatorStake({
            selfStake: msg.value,
            totalDelegated: 0,
            commissionRate: commissionRate,
            accumulatedRewards: 0,
            isJailed: false,
            jailUntilBlock: 0
        });

        validatorIndex[msg.sender] = allValidators.length;
        allValidators.push(msg.sender);
        totalStaked = totalStaked.add(msg.value);

        // Add to Validators highest validators set
        validatorsContract.tryAddValidatorToHighestSet(msg.sender);

        emit ValidatorRegistered(msg.sender, msg.value, commissionRate);
    }

    /**
     * @dev Add more self-stake to existing validator
     */
    function addValidatorStake() external payable onlyValidValidator(msg.sender) {
        require(msg.value > 0, "Amount must be positive");
        // Check if jailed (jailed validators must unjail first before adding stake)
        require(!validatorStakes[msg.sender].isJailed, "Validator is jailed, must unjail first");
        
        validatorStakes[msg.sender].selfStake = validatorStakes[msg.sender].selfStake.add(msg.value);
        totalStaked = totalStaked.add(msg.value);
    }

    /**
     * @dev Update validator commission rate
     * @param newCommissionRate New commission rate (0-10000)
     */
    function updateCommissionRate(uint256 newCommissionRate) external onlyValidValidator(msg.sender) {
        require(newCommissionRate <= COMMISSION_RATE_BASE, "Invalid commission rate");
        // Check if jailed (jailed validators must unjail first before updating commission)
        require(!validatorStakes[msg.sender].isJailed, "Validator is jailed, must unjail first");
        
        validatorStakes[msg.sender].commissionRate = newCommissionRate;
        emit ValidatorUpdated(msg.sender, newCommissionRate);
    }

    /**
     * @dev Start validator stake withdrawal (validator exit)
     * @param amount Amount to withdraw from self-stake
     * @notice If remaining stake would be less than MIN_VALIDATOR_STAKE, the withdrawal will fail
     * @notice Use emergencyExit() to withdraw all stake (requires minimum validators requirement)
     */
    function withdrawValidatorStake(uint256 amount) external onlyValidValidator(msg.sender) {
        require(amount > 0, "Amount must be positive");
        
        ValidatorStake storage stake = validatorStakes[msg.sender];
        require(stake.selfStake >= amount, "Insufficient self-stake");
        
        // Calculate remaining stake after withdrawal
        uint256 remainingStake = stake.selfStake.sub(amount);
        
        // If partial withdrawal, remaining stake must meet minimum requirement
        // If complete withdrawal (remainingStake == 0), use emergencyExit() instead
        require(remainingStake >= MIN_VALIDATOR_STAKE, "Remaining stake below minimum, use emergencyExit() to withdraw all");
        
        stake.selfStake = remainingStake;
        totalStaked = totalStaked.sub(amount);
        
        emit ValidatorStakeWithdrawn(msg.sender, amount);
        
        // Transfer the withdrawn amount
        (bool success, ) = payable(msg.sender).call{value: amount}("");
        require(success, "Transfer failed");
    }

    /**
     * @dev Emergency exit for validator (withdraw all stake)
     * @notice This function is only callable by validators that meet the minimum validator requirement
     * @notice Removes validator from allValidators array upon exit
     * @notice If validator is currently in currentValidatorSet, they will be jailed first to ensure smooth exit
     */
    function emergencyExit() external onlyValidValidator(msg.sender) {
        ValidatorStake storage stake = validatorStakes[msg.sender];
        uint256 withdrawAmount = stake.selfStake;
        
        // Check if validator is currently participating in consensus (in currentValidatorSet)
        // Note: isActiveValidator() checks currentValidatorSet, not jail status
        bool isInCurrentSet = validatorsContract.isActiveValidator(msg.sender);
        
        // Calculate remaining validators after exit
        uint256 currentActiveCount = getActiveValidatorCount();
        // If validator is in currentValidatorSet and not jailed, they count as active
        // If they exit, remaining count decreases by 1
        uint256 remainingCount = (isInCurrentSet && !stake.isJailed) ? currentActiveCount - 1 : currentActiveCount;
        
        // Ensure remaining validators meet minimum requirement after exit
        require(remainingCount >= MIN_VALIDATORS, "Cannot exit: would leave less than minimum validators");
        
        // If validator is currently active in consensus (in currentValidatorSet and not jailed), jail them first
        // This ensures they stop producing blocks immediately and exit smoothly at next epoch
        if (isInCurrentSet && !stake.isJailed) {
            // Jail validator for 1 epoch (86400 blocks) to ensure smooth exit
            // This prevents immediate disruption to consensus
            stake.isJailed = true;
            stake.jailUntilBlock = block.number + 86400; // 1 epoch
            emit ValidatorJailed(msg.sender, stake.jailUntilBlock);
        }
        
        stake.selfStake = 0;
        totalStaked = totalStaked.sub(withdrawAmount);
        
        // Remove validator from allValidators array
        _removeFromAllValidators(msg.sender);
        
        emit ValidatorExited(msg.sender, withdrawAmount);
        (bool success, ) = payable(msg.sender).call{value: withdrawAmount}("");
        require(success, "Transfer failed");
    }

    /**
     * @dev Delegate tokens to a validator
     * @param validator Validator address to delegate to
     */
    function delegate(address validator) external payable onlyActiveValidator(validator) {
        require(validator != address(0), "Invalid validator address");
        require(msg.value >= MIN_DELEGATION, "Insufficient delegation amount");
        require(validator != msg.sender, "Cannot delegate to yourself");
        require(proposalContract.pass(validator), "Validator must pass proposal");
        
        _updateRewards(msg.sender, validator);
        
        delegations[msg.sender][validator].amount = delegations[msg.sender][validator].amount.add(msg.value);
        validatorStakes[validator].totalDelegated = validatorStakes[validator].totalDelegated.add(msg.value);
        totalStaked = totalStaked.add(msg.value);
        
        // Update reward debt
        delegations[msg.sender][validator].rewardDebt = delegations[msg.sender][validator].amount
            .mul(rewardPerShare[validator])
            .div(1e18);

        emit Delegated(msg.sender, validator, msg.value);
    }

    /**
     * @dev Start unbonding delegation from a validator
     * @param validator Validator address to undelegate from
     * @param amount Amount to undelegate
     */
    function undelegate(address validator, uint256 amount) external {
        require(validator != address(0), "Invalid validator address");
        require(amount > 0, "Amount must be positive");
        require(validator != msg.sender, "Cannot undelegate from yourself");
        require(delegations[msg.sender][validator].amount >= amount, "Insufficient delegation");
        // Allow undelegation even if validator has exited (selfStake == 0)
        // This ensures delegators can always withdraw their funds
        
        _updateRewards(msg.sender, validator);
        
        delegations[msg.sender][validator].amount = delegations[msg.sender][validator].amount.sub(amount);
        validatorStakes[validator].totalDelegated = validatorStakes[validator].totalDelegated.sub(amount);
        totalStaked = totalStaked.sub(amount);
        
        // Add to unbonding
        unbondingDelegations[msg.sender][validator].push(UnbondingEntry({
            amount: amount,
            completionBlock: block.number.add(UNBONDING_PERIOD)
        }));
        
        // Update reward debt
        delegations[msg.sender][validator].rewardDebt = delegations[msg.sender][validator].amount
            .mul(rewardPerShare[validator])
            .div(1e18);

        emit Undelegated(msg.sender, validator, amount);
    }

    /**
     * @dev Complete unbonding and withdraw tokens
     * @param validator Validator address
     * @param maxEntries Maximum number of unbonding entries to process
     */
    function withdrawUnbonded(address validator, uint256 maxEntries) external {
        require(maxEntries > 0, "maxEntries must be positive");
        require(maxEntries <= MAX_UNBONDING_ENTRIES_PER_WITHDRAW, "maxEntries too large");
        
        UnbondingEntry[] storage entries = unbondingDelegations[msg.sender][validator];
        uint256 totalWithdraw = 0;
        uint256 processed = 0;
        
        for (uint256 i = 0; i < entries.length && processed < maxEntries;) {
            if (entries[i].completionBlock <= block.number) {
                totalWithdraw = totalWithdraw.add(entries[i].amount);
                
                // Remove completed entry
                entries[i] = entries[entries.length - 1];
                entries.pop();
                processed++;
                // Don't increment i, continue checking current position (new element moved here)
            } else {
                i++; // Only increment when not deleting
            }
        }
        
        require(totalWithdraw > 0, "No unbonded tokens available");
        
        (bool success, ) = payable(msg.sender).call{value: totalWithdraw}("");
        require(success, "Transfer failed");
        emit UnbondingCompleted(msg.sender, validator, totalWithdraw);
    }

    /**
     * @dev Distribute rewards to a validator
     * @param validator Validator address
     */
    function distributeRewards(address validator) external payable onlyMiner onlyActiveValidator(validator) {
        require(msg.value > 0, "No rewards to distribute");
        
        ValidatorStake storage stake = validatorStakes[validator];
        uint256 totalStake = stake.selfStake.add(stake.totalDelegated);
        
        if (totalStake == 0) return;
        
        // 1. Calculate and allocate commission to validator first
        uint256 commission = msg.value.mul(stake.commissionRate).div(COMMISSION_RATE_BASE);
        stake.accumulatedRewards = stake.accumulatedRewards.add(commission);
        
        // 2. Calculate remaining rewards after commission
        uint256 remainingRewards = msg.value.sub(commission);
        
        // 3. Calculate validator's share from remaining rewards based on their self-stake proportion
        uint256 validatorShare = remainingRewards.mul(stake.selfStake).div(totalStake);
        stake.accumulatedRewards = stake.accumulatedRewards.add(validatorShare);
        
        // 4. Calculate delegator rewards (remaining after validator's share)
        uint256 delegatorRewards = remainingRewards.sub(validatorShare);
        
        // 5. Update reward per share for delegators
        if (stake.totalDelegated > 0) {
            rewardPerShare[validator] = rewardPerShare[validator].add(
                delegatorRewards.mul(1e18).div(stake.totalDelegated)
            );
        } else {
            // If no delegators, allocate delegatorRewards to validator
            stake.accumulatedRewards = stake.accumulatedRewards.add(delegatorRewards);
        }
        
        emit RewardsDistributed(validator, msg.value);
    }

    /**
     * @dev Claim pending rewards
     * @param validator Validator address
     */
    function claimRewards(address validator) external {
        _updateRewards(msg.sender, validator);
        
        // For validator claiming their commission
        if (msg.sender == validator) {
            uint256 commission = validatorStakes[validator].accumulatedRewards;
            if (commission > 0) {
                // Transfer first, then update state to avoid state inconsistency on failure
                (bool success, ) = payable(msg.sender).call{value: commission}("");
                require(success, "Transfer failed");
                validatorStakes[validator].accumulatedRewards = 0;
                emit RewardsClaimed(msg.sender, validator, commission);
            }
        }
    }

    /**
     * @dev Jail a validator for misbehavior
     * @param validator Validator address to jail
     * @param jailBlocks Number of blocks to jail for
     */
    function jailValidator(address validator, uint256 jailBlocks) external onlyPunishContract {
        require(validator != address(0), "Invalid validator address");
        require(jailBlocks > 0, "Jail blocks must be positive");
        validatorStakes[validator].isJailed = true;
        validatorStakes[validator].jailUntilBlock = block.number.add(jailBlocks);
        emit ValidatorJailed(validator, validatorStakes[validator].jailUntilBlock);
    }

    /**
     * @dev Unjail a validator 
     * @param validator Validator address to unjail
     */
    function unjailValidator(address validator) external {
        require(validator != address(0), "Invalid validator address");
        require(msg.sender == validator, "Only validator can unjail themselves");
        require(validatorStakes[validator].isJailed, "Validator not jailed");
        require(block.number >= validatorStakes[validator].jailUntilBlock, "Jail period not complete");
        
        // Check if validator has sufficient stake to be active
        require(validatorStakes[validator].selfStake >= MIN_VALIDATOR_STAKE, "Insufficient stake, must add stake first");
        
        uint256 violations = proposalContract.getViolationCount(validator);
        
        // If violations > 3, must have passed proposal (reproposal required)
        if (violations > proposalContract.MAX_VIOLATIONS_FOR_AUTO_UNJAIL()) {
            require(proposalContract.pass(validator), "Too many violations, must pass reproposal first");
        }
        
        validatorStakes[validator].isJailed = false;
        validatorStakes[validator].jailUntilBlock = 0;
        emit ValidatorUnjailed(validator);
        
        // If violations <= 3, auto restore pass status
        // If violations > 3, pass should already be true from reproposal
        if (violations <= proposalContract.MAX_VIOLATIONS_FOR_AUTO_UNJAIL()) {
            bool restored = false;
            try proposalContract.autoRestorePass(validator) returns (bool success) {
                restored = success;
            } catch {}
            
            if (restored) {
                // Only reactivate if pass was restored
                try validatorsContract.tryActive(validator) {} catch {}
            }
            // Note: If auto restore failed, validator is unjailed but pass=false
            // This is acceptable - validator can create a proposal to regain pass status
            // Creating a proposal doesn't require pass status
        } else {
            // For violations > 3, pass should already be true from reproposal
            // Just try to reactivate
            try validatorsContract.tryActive(validator) {} catch {}
        }
    }

    /**
     * @dev Remove validator from allValidators array
     * @param validator Validator address to remove
     * @notice This function safely removes a validator from the array while preserving array integrity
     * @notice Similar to Punish.sol's cleanPunishRecord implementation
     */
    function _removeFromAllValidators(address validator) private {
        uint256 index = validatorIndex[validator];
        
        // Check if index is valid
        if (index >= allValidators.length) {
            return; // Validator not in array, may have been cleaned already
        }
        
        // If validator is not the last element, move last element to current position
        if (index != allValidators.length - 1) {
            address lastValidator = allValidators[allValidators.length - 1];
            allValidators[index] = lastValidator;
            validatorIndex[lastValidator] = index;
        }
        
        // Remove last element
        allValidators.pop();
        
        // Clear validatorIndex for removed validator
        delete validatorIndex[validator];
    }

    /**
     * @dev Get top validators by total stake
     * @return Top validators list (up to MAX_VALIDATORS)
     */
    function getTopValidators() external view returns (address[] memory) {
        // Create array of all active validators with their total stake
        address[] memory activeValidators = new address[](allValidators.length);
        uint256[] memory totalStakes = new uint256[](allValidators.length);
        uint256 activeCount = 0;
        
        for (uint256 i = 0; i < allValidators.length; i++) {
            address validator = allValidators[i];
            if (validatorStakes[validator].selfStake >= MIN_VALIDATOR_STAKE && 
                !validatorStakes[validator].isJailed &&
                proposalContract.pass(validator)) {
                activeValidators[activeCount] = validator;
                totalStakes[activeCount] = validatorStakes[validator].selfStake.add(validatorStakes[validator].totalDelegated);
                activeCount++;
            }
        }
        
        // Sort by total stake (simple bubble sort for small arrays)
        for (uint256 i = 0; i < activeCount; i++) {
            for (uint256 j = i + 1; j < activeCount; j++) {
                if (totalStakes[i] < totalStakes[j]) {
                    // Swap stakes
                    uint256 tempStake = totalStakes[i];
                    totalStakes[i] = totalStakes[j];
                    totalStakes[j] = tempStake;
                    
                    // Swap validators
                    address tempValidator = activeValidators[i];
                    activeValidators[i] = activeValidators[j];
                    activeValidators[j] = tempValidator;
                }
            }
        }
        
        // Return top validators (up to MAX_VALIDATORS)
        uint256 returnLength = activeCount < MAX_VALIDATORS ? activeCount : MAX_VALIDATORS;
        address[] memory topValidators = new address[](returnLength);
        for (uint256 i = 0; i < returnLength; i++) {
            topValidators[i] = activeValidators[i];
        }
        
        return topValidators;
    }

    /**
     * @dev Update rewards for a delegator
     * @param delegator Delegator address
     * @param validator Validator address
     */
    function _updateRewards(address delegator, address validator) internal {
        Delegation storage delegation = delegations[delegator][validator];
        if (delegation.amount > 0) {
            uint256 pending = delegation.amount
                .mul(rewardPerShare[validator])
                .div(1e18)
                .sub(delegation.rewardDebt);
                
            if (pending > 0) {
                // Transfer first, then update state to avoid state inconsistency on failure
                (bool success, ) = payable(delegator).call{value: pending}("");
                require(success, "Transfer failed");
                
                // Update reward debt after successful transfer to prevent reentrancy
                delegation.rewardDebt = delegation.amount
                    .mul(rewardPerShare[validator])
                    .div(1e18);
                    
                emit RewardsClaimed(delegator, validator, pending);
            }
        }
    }

    /**
     * @dev Check if validator is currently jailed
     * @param validator Validator address
     * @return Whether validator is currently jailed
     */
    function isValidatorJailed(address validator) external view returns (bool) {
        ValidatorStake storage stake = validatorStakes[validator];
        return stake.isJailed;
    }

    /**
     * @dev Get validator status (active and jailed)
     * @param validator Validator address
     * @return isActive Whether validator is active (actually participating in consensus, in currentValidatorSet)
     * @return isJailed Whether validator is currently jailed
     * @notice isActive means validator is in currentValidatorSet and can actually produce blocks
     */
    function getValidatorStatus(address validator) external view returns (bool isActive, bool isJailed) {
        ValidatorStake storage stake = validatorStakes[validator];
        isJailed = stake.isJailed;
        // isActive means validator is actually in currentValidatorSet (participating in consensus)
        // Not just meeting conditions, but actually active in the current epoch
        isActive = validatorsContract.isActiveValidator(validator) && !isJailed;
    }

    /**
     * @dev Get validator information
     * @param validator Validator address
     * @return selfStake Validator's self stake
     * @return totalDelegated Total delegated amount
     * @return commissionRate Commission rate
     * @return isJailed Whether validator is jailed
     * @return jailUntilBlock Block until which validator is jailed
     */
    function getValidatorInfo(address validator) external view returns (
        uint256 selfStake,
        uint256 totalDelegated,
        uint256 commissionRate,
        bool isJailed,
        uint256 jailUntilBlock
    ) {
        ValidatorStake storage stake = validatorStakes[validator];
        return (
            stake.selfStake,
            stake.totalDelegated,
            stake.commissionRate,
            stake.isJailed,
            stake.jailUntilBlock
        );
    }

    /**
     * @dev Get delegation information
     * @param delegator Delegator address
     * @param validator Validator address
     * @return amount Delegated amount
     * @return pendingRewards Pending reward amount
     * @return unbondingAmount Amount being unbonded
     * @return unbondingBlock Block when unbonding completes
     */
    function getDelegationInfo(address delegator, address validator) external view returns (
        uint256 amount,
        uint256 pendingRewards,
        uint256 unbondingAmount,
        uint256 unbondingBlock
    ) {
        Delegation storage delegation = delegations[delegator][validator];
        uint256 pending = 0;
        
        if (delegation.amount > 0) {
            pending = delegation.amount
                .mul(rewardPerShare[validator])
                .div(1e18)
                .sub(delegation.rewardDebt);
        }
        
        return (
            delegation.amount,
            pending,
            delegation.unbondingAmount,
            delegation.unbondingBlock
        );
    }

    /**
     * @dev Get count of unbonding entries for a delegator-validator pair
     * @param delegator Delegator address
     * @param validator Validator address
     * @return Count of unbonding entries
     */
    function getUnbondingEntriesCount(address delegator, address validator) external view returns (uint256) {
        return unbondingDelegations[delegator][validator].length;
    }

    /**
     * @dev Get all unbonding entries for a delegator-validator pair
     * @param delegator Delegator address
     * @param validator Validator address
     * @return Array of unbonding entries
     */
    function getUnbondingEntries(address delegator, address validator) external view returns (UnbondingEntry[] memory) {
        return unbondingDelegations[delegator][validator];
    }

    /**
     * @dev Get total number of validators
     * @return Total validator count
     */
    function getValidatorCount() external view returns (uint256) {
        return allValidators.length;
    }

    /**
     * @dev Get number of active validators (meeting minimum stake and not jailed)
     * @return Active validator count
     */
    /**
     * @dev Get count of active validators (in currentValidatorSet and not jailed)
     * @return Count of validators actually participating in consensus
     * @notice This delegates to Validators contract which counts validators in currentValidatorSet
     */
    function getActiveValidatorCount() public view returns (uint256) {
        return validatorsContract.getActiveValidatorCount();
    }

    /**
     * @dev Check if system has minimum required validators
     * @return Whether system meets minimum validator requirement
     */
    function hasMinimumValidators() external view returns (bool) {
        return getActiveValidatorCount() >= MIN_VALIDATORS;
    }
}
