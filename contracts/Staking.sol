// SPDX-License-Identifier: MIT

pragma solidity ^0.8.20;

import {Params} from './Params.sol';
import {Proposal} from './Proposal.sol';
import {ReentrancyGuard} from './library/ReentrancyGuard.sol';

// Interface for Validators contract to avoid circular dependency
interface IValidators {
    function tryAddValidatorToHighestSet(address validator) external;
    function tryActive(address validator) external returns (bool);
    function isActiveValidator(address who) external view returns (bool);
    function removeFromHighestSet(address validator) external;
}

/**
 * @title Staking Contract for JPoSA Consensus
 * @dev Implements staking mechanism for JuChain Proof of Stake Authorization
 */
contract Staking is Params, ReentrancyGuard {

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
    
    // Maximum number of unbonding entries to process in a single withdrawUnbonded call
    uint256 public constant MAX_UNBONDING_ENTRIES_PER_WITHDRAW = 50;

    struct ValidatorStake {
        uint256 selfStake;          // Validator's own stake
        uint256 totalDelegated;     // Total delegated to this validator
        uint256 commissionRate;     // Commission rate (0-10000, representing 0%-100%)
        uint256 accumulatedRewards; // Accumulated rewards for distribution
        bool isJailed;              // Whether validator is jailed
        uint256 jailUntilBlock;     // Block number until which validator is jailed
        uint256 totalClaimedRewards; // Total claimed rewards (cumulative)
        uint256 lastClaimBlock;      // Block number of last reward claim
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

    // Operations enum for tracking operations done per block
    enum Operations {Distribute}
    // Record the operations is done or not.
    mapping(uint256 => mapping(uint8 => bool)) operationsDone;

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
     * @notice This function automatically performs the same logic as registerValidator for genesis validators
     * @notice Genesis validators are pre-registered with default stake and automatically added to highestValidatorsSet
     * @notice They don't need to pass proposal or wait for 7-day window - they are activated immediately
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
        // This automatically performs the same logic as registerValidator for genesis validators
        for (uint256 i = 0; i < initialValidators.length; i++) {
            address validator = initialValidators[i];
            require(validator != address(0), "Invalid validator address");
            require(validatorStakes[validator].selfStake == 0, "Validator already exists");
            
            // Set up validator stake (same as registerValidator)
            validatorStakes[validator] = ValidatorStake({
                selfStake: MIN_VALIDATOR_STAKE,
                totalDelegated: 0,
                commissionRate: commissionRate,
                accumulatedRewards: 0,
                isJailed: false,
                jailUntilBlock: 0,
                totalClaimedRewards: 0,
                lastClaimBlock: 0
            });
            
            // Add to validators list (same as registerValidator)
            validatorIndex[validator] = allValidators.length;
            allValidators.push(validator);
            totalStaked = totalStaked + MIN_VALIDATOR_STAKE;
            // Emit event for each validator (more accurate than single event)
            emit ValidatorRegistered(validator, MIN_VALIDATOR_STAKE, commissionRate);
        }
        
        initialized = true;
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
            jailUntilBlock: 0,
            totalClaimedRewards: 0,
            lastClaimBlock: 0
        });

        validatorIndex[msg.sender] = allValidators.length;
        allValidators.push(msg.sender);
        totalStaked = totalStaked + msg.value;

        // Activate validator (register = activate)
        // tryActive will add to highestValidatorsSet and emit LogActive event
        validatorsContract.tryActive(msg.sender);

        emit ValidatorRegistered(msg.sender, msg.value, commissionRate);
    }

    /**
     * @dev Add more self-stake to existing validator
     */
    function addValidatorStake() external payable onlyValidValidator(msg.sender) {
        require(msg.value > 0, "Amount must be positive");
        // Check if jailed (jailed validators must unjail first before adding stake)
        require(!validatorStakes[msg.sender].isJailed, "Validator is jailed, must unjail first");
        
        validatorStakes[msg.sender].selfStake = validatorStakes[msg.sender].selfStake + msg.value;
        totalStaked = totalStaked + msg.value;
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
    function withdrawValidatorStake(uint256 amount) external nonReentrant onlyValidValidator(msg.sender) {
        require(amount > 0, "Amount must be positive");
        
        ValidatorStake storage stake = validatorStakes[msg.sender];
        require(stake.selfStake >= amount, "Insufficient self-stake");
        
        // Calculate remaining stake after withdrawal
        uint256 remainingStake = stake.selfStake - amount;
        
        // If partial withdrawal, remaining stake must meet minimum requirement
        // If complete withdrawal (remainingStake == 0), use emergencyExit() instead
        require(remainingStake >= MIN_VALIDATOR_STAKE, "Remaining stake below minimum, use emergencyExit() to withdraw all");
        
        // Effects: update state before external call
        stake.selfStake = remainingStake;
        totalStaked = totalStaked - amount;
        
        emit ValidatorStakeWithdrawn(msg.sender, amount);
        
        // Interactions: external call after state update
        (bool success, ) = payable(msg.sender).call{value: amount}("");
        require(success, "Transfer failed");
    }

    /**
     * @dev Allow validator to resign from validator role
     * @notice This function allows validators to voluntarily resign from their validator role
     * @notice Validator will be removed from highestValidatorsSet and pass will be set to false
     * @notice Validator's transaction fee income (aacIncoming) will NOT be removed
     * @notice After resigning, validator will be excluded from currentValidatorSet at next epoch
     * @notice After being excluded, validator can call emergencyExit() to withdraw all stake
     */
    function resignValidator() external onlyValidValidator(msg.sender) {
        ValidatorStake storage stake = validatorStakes[msg.sender];
        
        // Cannot resign if already jailed/resigned
        require(!stake.isJailed, "Validator already resigned or jailed");
        
        // Mark validator as jailed to ensure exclusion from currentValidatorSet at next epoch
        // This is a technical mechanism to ensure smooth exit, not a punishment
        stake.isJailed = true;
        stake.jailUntilBlock = block.number + 86400;
        emit ValidatorJailed(msg.sender, stake.jailUntilBlock);
        
        // Remove from highestValidatorsSet and set pass = false
        // Note: This does NOT remove transaction fee income (aacIncoming)
        validatorsContract.removeFromHighestSet(msg.sender);
    }

    /**
     * @dev Emergency exit - withdraw all stake and exit validator role
     * @notice Validators in currentValidatorSet (active in current epoch) cannot exit
     * @notice Validators must jail themselves first, then wait until next epoch to exit
     * @notice This ensures smooth exit without disrupting consensus
     */
    function emergencyExit() external nonReentrant onlyValidValidator(msg.sender) {
        ValidatorStake storage stake = validatorStakes[msg.sender];
        uint256 withdrawAmount = stake.selfStake;
        
        // Check if validator is currently participating in consensus (in currentValidatorSet)
        // Note: isActiveValidator() checks currentValidatorSet, not jail status
        bool isInCurrentSet = validatorsContract.isActiveValidator(msg.sender);
        
        // Validators in active set cannot exit - they must resign first
        // After resigning, they will be excluded from currentValidatorSet at next epoch
        require(!isInCurrentSet, "Cannot exit: validator is in active set, resign first and wait until next epoch");
        
        // Note: Since validator is not in currentValidatorSet, their exit won't affect
        // the active validator count, so no need to check MIN_VALIDATORS here
        
        // Effects: update state before external call
        stake.selfStake = 0;
        totalStaked = totalStaked - withdrawAmount;
        
        // Note: If validator called resignValidator() before, they are already removed from
        // highestValidatorsSet and pass is already set to false, so no need to do it again
        // Only remove from highestValidatorsSet and set pass = false if validator didn't call resignValidator
        // Check if validator is still in highestValidatorsSet (by checking if pass is still true)
        if (proposalContract.pass(msg.sender)) {
            // Validator didn't call resignValidator(), so we need to remove them now
            validatorsContract.removeFromHighestSet(msg.sender);
        }
        // If pass is already false, validator already called resignValidator(), no need to do anything
        
        // Remove validator from allValidators array
        _removeFromAllValidators(msg.sender);
        
        emit ValidatorExited(msg.sender, withdrawAmount);
        
        // Interactions: external call after state update
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
        
        _updateRewards(msg.sender, validator);
        
        delegations[msg.sender][validator].amount = delegations[msg.sender][validator].amount + msg.value;
        validatorStakes[validator].totalDelegated = validatorStakes[validator].totalDelegated + msg.value;
        totalStaked = totalStaked + msg.value;
        
        // Update reward debt
        delegations[msg.sender][validator].rewardDebt = delegations[msg.sender][validator].amount
            * rewardPerShare[validator]
            / 1e18;

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
        
        delegations[msg.sender][validator].amount = delegations[msg.sender][validator].amount - amount;
        validatorStakes[validator].totalDelegated = validatorStakes[validator].totalDelegated - amount;
        totalStaked = totalStaked - amount;
        
        // Add to unbonding
        unbondingDelegations[msg.sender][validator].push(UnbondingEntry({
            amount: amount,
            completionBlock: block.number + proposalContract.unbondingPeriod()
        }));
        
        // Update reward debt
        delegations[msg.sender][validator].rewardDebt = delegations[msg.sender][validator].amount
            * rewardPerShare[validator]
            / 1e18;
        
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
                totalWithdraw = totalWithdraw + entries[i].amount;
                
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
        
        // If delegation amount is 0 and all unbonding entries are withdrawn, delete the delegation entry
        // This helps save storage when delegation is completely removed
        if (delegations[msg.sender][validator].amount == 0 && 
            unbondingDelegations[msg.sender][validator].length == 0) {
            delete delegations[msg.sender][validator];
        }
    }

    /**
     * @dev Distribute rewards to the current block miner (validator)
     * @notice Validator address is obtained from block.coinbase
     * @notice Block reward is passed via msg.value (consensus layer reads from Proposal contract)
     * @notice Jailed validators can still produce blocks and receive rewards in the current epoch
     * @notice They will be excluded from the validator set at the next epoch transition
     */
    function distributeRewards() external payable onlyMiner onlyInitialized {
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
        
        // Get block reward from msg.value (consensus layer reads from Proposal contract and passes it here)
        // This avoids duplicate contract calls and saves gas
        uint256 blockReward = msg.value;
        
        // Check if there are rewards to distribute
        if (blockReward == 0) {
            return;
        }
        
        // Get validator address from block.coinbase (similar to Validators.distributeBlockReward())
        address validator = block.coinbase;
        
        // Check if validator exists (has staked)
        ValidatorStake storage stake = validatorStakes[validator];
        if (stake.selfStake == 0) {
            return; // Validator doesn't exist, silently return
        }
        
        // Note: We don't check selfStake >= MIN_VALIDATOR_STAKE here because:
        // 1. Validators in active set cannot exit (emergencyExit() rejects them)
        // 2. Validators cannot reduce stake below MIN_VALIDATOR_STAKE (withdrawValidatorStake() requires remainingStake >= MIN_VALIDATOR_STAKE)
        // 3. If validator can produce a block, they must be in active set and have sufficient stake
        
        uint256 totalStake = stake.selfStake + stake.totalDelegated;
        
        if (totalStake == 0) return;
        
        // 1. Calculate and allocate commission to validator first
        uint256 commission = blockReward * stake.commissionRate / COMMISSION_RATE_BASE;
        stake.accumulatedRewards = stake.accumulatedRewards + commission;
        
        // 2. Calculate remaining rewards after commission
        uint256 remainingRewards = blockReward - commission;
        
        // 3. Calculate validator's share from remaining rewards based on their self-stake proportion
        uint256 validatorShare = remainingRewards * stake.selfStake / totalStake;
        stake.accumulatedRewards = stake.accumulatedRewards + validatorShare;
        
        // 4. Calculate delegator rewards (remaining after validator's share)
        uint256 delegatorRewards = remainingRewards - validatorShare;
        
        // 5. Update reward per share for delegators
        if (stake.totalDelegated > 0) {
            rewardPerShare[validator] = rewardPerShare[validator] + (
                delegatorRewards * 1e18 / stake.totalDelegated
            );
        } else {
            // If no delegators, allocate delegatorRewards to validator
            stake.accumulatedRewards = stake.accumulatedRewards + delegatorRewards;
        }
        
        emit RewardsDistributed(validator, blockReward);
    }

    /**
     * @dev Claim pending rewards
     * @param validator Validator address
     * @notice Validators must wait withdrawProfitPeriod blocks between claims
     * @notice Tracks total claimed rewards and last claim block for statistics
     * @notice First claim is always allowed (when lastClaimBlock == 0)
     */
    function claimRewards(address validator) external nonReentrant {
        _updateRewards(msg.sender, validator);
        
        // For validator claiming their commission
        if (msg.sender == validator) {
            ValidatorStake storage stake = validatorStakes[validator];
            uint256 commission = stake.accumulatedRewards;
            
            if (commission > 0) {
                // Check withdrawal period restriction (allow first claim when lastClaimBlock == 0)
                if (stake.lastClaimBlock > 0) {
                    uint256 withdrawPeriod = proposalContract.withdrawProfitPeriod();
                    require(
                        block.number >= stake.lastClaimBlock + withdrawPeriod,
                        "Must wait withdrawProfitPeriod blocks between claims"
                    );
                }
                
                // Effects: update state before external call
                stake.accumulatedRewards = 0;
                stake.totalClaimedRewards = stake.totalClaimedRewards + commission;
                stake.lastClaimBlock = block.number;
                
                // Interactions: external call after state update
                (bool success, ) = payable(msg.sender).call{value: commission}("");
                require(success, "Transfer failed");
                
                emit RewardsClaimed(msg.sender, validator, commission);
            }
        }
    }

    /**
     * @dev Jail a validator for misbehavior
     * @param validator Validator address to jail
     * @param jailBlocks Number of blocks to jail for
     */
    /**
     * @dev Jail a validator for misbehavior (called by Punish contract)
     * @param validator Validator address to jail
     * @param jailBlocks Number of blocks to jail for
     * @notice Only Punish contract can call this function
     * @notice Validators should use resignValidator() to voluntarily resign from validator role
     */
    function jailValidator(address validator, uint256 jailBlocks) external onlyPunishContract {
        require(validator != address(0), "Invalid validator address");
        require(jailBlocks > 0, "Jail blocks must be positive");
        validatorStakes[validator].isJailed = true;
        validatorStakes[validator].jailUntilBlock = block.number + jailBlocks;
        emit ValidatorJailed(validator, validatorStakes[validator].jailUntilBlock);
    }

    /**
     * @dev Unjail a validator
     * @param validator Validator address to unjail
     * @notice Validator must have passed a proposal (reproposal) before unjailing
     * @notice Once jailed, validator must go through the voting process again to regain validator status
     */
    function unjailValidator(address validator) external {
        require(validator != address(0), "Invalid validator address");
        require(msg.sender == validator, "Only validator can unjail themselves");
        require(validatorStakes[validator].isJailed, "Validator not jailed");
        require(block.number >= validatorStakes[validator].jailUntilBlock, "Jail period not complete");
        
        // Check if validator has sufficient stake to be active
        require(validatorStakes[validator].selfStake >= MIN_VALIDATOR_STAKE, "Insufficient stake, must add stake first");
        
        // Must have passed proposal (reproposal required) - no automatic restoration
        require(proposalContract.pass(validator), "Must pass reproposal first");
        
        // Activate validator first (must succeed before unjailing)
        // This ensures validator is added to highestValidatorsSet before state change
        require(validatorsContract.tryActive(validator), "Failed to activate validator");
        
        validatorStakes[validator].isJailed = false;
        validatorStakes[validator].jailUntilBlock = 0;
        emit ValidatorUnjailed(validator);
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
     * @notice Sorts the provided validators by total stake (selfStake + totalDelegated)
     * @param validators Array of validator addresses to sort
     * @return Top validators list (up to MAX_VALIDATORS), sorted by total stake
     */
    function getTopValidators(address[] memory validators) external view returns (address[] memory) {
        if (validators.length == 0) {
            return new address[](0);
        }
        
        // Create arrays for validators and their total stakes
        address[] memory candidateValidators = new address[](validators.length);
        uint256[] memory totalStakes = new uint256[](validators.length);
        uint256 candidateCount = 0;
        
        // Collect validators and their total stakes for sorting
        for (uint256 i = 0; i < validators.length; i++) {
            address validator = validators[i];
            ValidatorStake storage stake = validatorStakes[validator];
            
            candidateValidators[candidateCount] = validator;
            totalStakes[candidateCount] = stake.selfStake + stake.totalDelegated;
            candidateCount++;
        }
        
        if (candidateCount == 0) {
            return new address[](0);
        }
        
        // Sort by total stake (simple bubble sort for small arrays)
        for (uint256 i = 0; i < candidateCount; i++) {
            for (uint256 j = i + 1; j < candidateCount; j++) {
                if (totalStakes[i] < totalStakes[j]) {
                    // Swap stakes
                    uint256 tempStake = totalStakes[i];
                    totalStakes[i] = totalStakes[j];
                    totalStakes[j] = tempStake;
                    
                    // Swap validators
                    address tempValidator = candidateValidators[i];
                    candidateValidators[i] = candidateValidators[j];
                    candidateValidators[j] = tempValidator;
                }
            }
        }
        
        // Return top validators (up to MAX_VALIDATORS)
        uint256 returnLength = candidateCount < MAX_VALIDATORS ? candidateCount : MAX_VALIDATORS;
        address[] memory topValidators = new address[](returnLength);
        for (uint256 i = 0; i < returnLength; i++) {
            topValidators[i] = candidateValidators[i];
        }
        
        return topValidators;
    }

    /**
     * @dev Update rewards for a delegator
     * @param delegator Delegator address
     * @param validator Validator address
     * @notice This function is internal and called from nonReentrant functions
     * @notice Uses CEI pattern: calculates pending, updates state, then transfers
     */
    function _updateRewards(address delegator, address validator) internal {
        Delegation storage delegation = delegations[delegator][validator];
        if (delegation.amount > 0) {
            uint256 pending = delegation.amount
                * rewardPerShare[validator]
                / 1e18
                - delegation.rewardDebt;
                
            if (pending > 0) {
                // Effects: update state before external call to prevent reentrancy
                delegation.rewardDebt = delegation.amount
                    * rewardPerShare[validator]
                    / 1e18;
                
                // Interactions: external call after state update
                (bool success, ) = payable(delegator).call{value: pending}("");
                require(success, "Transfer failed");
                    
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
     * @notice Jailed validators can still be active in the current epoch until the next epoch transition
     */
    function getValidatorStatus(address validator) external view returns (bool isActive, bool isJailed) {
        ValidatorStake storage stake = validatorStakes[validator];
        isJailed = stake.isJailed;
        // isActive means validator is actually in currentValidatorSet (participating in consensus)
        // Jailed validators remain active in currentValidatorSet until next epoch, so we don't check jailed status
        isActive = validatorsContract.isActiveValidator(validator);
    }

    /**
     * @dev Get validator information
     * @param validator Validator address
     * @return selfStake Validator's self stake
     * @return totalDelegated Total delegated amount
     * @return commissionRate Commission rate
     * @return accumulatedRewards Accumulated rewards available for claim
     * @return isJailed Whether validator is jailed
     * @return jailUntilBlock Block until which validator is jailed
     * @return totalClaimedRewards Total claimed rewards (cumulative)
     * @return lastClaimBlock Block number of last reward claim
     */
    function getValidatorInfo(address validator) external view returns (
        uint256 selfStake,
        uint256 totalDelegated,
        uint256 commissionRate,
        uint256 accumulatedRewards,
        bool isJailed,
        uint256 jailUntilBlock,
        uint256 totalClaimedRewards,
        uint256 lastClaimBlock
    ) {
        ValidatorStake storage stake = validatorStakes[validator];
        return (
            stake.selfStake,
            stake.totalDelegated,
            stake.commissionRate,
            stake.accumulatedRewards,
            stake.isJailed,
            stake.jailUntilBlock,
            stake.totalClaimedRewards,
            stake.lastClaimBlock
        );
    }

    /**
     * @dev Get count of all registered validators
     * @notice Returns the total number of validators that have registered (including inactive ones)
     * @return Count of all registered validators
     */
    function getValidatorCount() external view returns (uint256) {
        return allValidators.length;
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
                * rewardPerShare[validator]
                / 1e18
                - delegation.rewardDebt;
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
}
