// SPDX-License-Identifier: MIT

pragma solidity ^0.8.29;

import {Params} from "./Params.sol";
import {IProposal} from "./IProposal.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import {IValidators} from "./IValidators.sol";
import {IStaking} from "./IStaking.sol";

/**
 * @title Staking Contract for JPoSA Consensus
 * @dev Implements staking mechanism for JuChain Proof of Stake Authorization
 */
contract Staking is Params, ReentrancyGuard, IStaking {


    
    // Commission rate precision (10000 = 100%)
    uint256 public constant COMMISSION_RATE_BASE = 10000;    
    
    // Maximum number of unbonding entries per delegator-validator pair
    uint256 public constant MAX_UNBONDING_ENTRIES = 20;

    struct ValidatorStake {
        uint256 selfStake;          // Validator's own stake
        uint256 totalDelegated;     // Total delegated to this validator
        uint256 commissionRate;     // Commission rate (0-10000, representing 0%-100%)
        uint256 accumulatedRewards; // Accumulated rewards for distribution
        bool isJailed;              // Whether validator is jailed
        uint256 jailUntilBlock;     // Block number until which validator is jailed
        uint256 totalClaimedRewards; // Total claimed rewards (cumulative)
        uint256 lastClaimBlock;      // Block number of last reward claim
        bool isRegistered;          // Whether validator is registered
    }

    struct Delegation {
        uint256 amount;             // Delegated amount
        uint256 rewardDebt;         // Reward debt for accurate reward calculation
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
    IProposal public proposalContract;

    // Total number of unique delegators
    uint256 public delegatorCount;
    // Mapping to track if an address is a delegator
    mapping(address => bool) public delegatorExists;
    // Validator address => number of delegators
    mapping(address => uint256) public validatorDelegatorCount;

    event ValidatorRegistered(address indexed validator, uint256 selfStake, uint256 commissionRate);
    event CommissionRateUpdated(address indexed validator, uint256 commissionRate);
    event ValidatorStakeIncreased(address indexed delegator, address indexed validator, uint256 amount);
    event ValidatorStakeDecreased(address indexed delegator, address indexed validator, uint256 amount);
    event ValidatorExited(address indexed validator);
    event Delegated(address indexed delegator, address indexed validator, uint256 amount);
    event Undelegated(address indexed delegator, address indexed validator, uint256 amount);
    event UnbondingCompleted(address indexed delegator, address indexed validator, uint256 amount);
    event RewardsDistributed(address indexed validator, uint256 amount);
    event RewardsClaimed(address indexed delegator, address indexed validator, uint256 amount);
    event ValidatorJailed(address indexed validator, uint256 jailUntilBlock);
    event ValidatorUnjailed(address indexed validator);

    modifier onlyValidValidator(address validator) {
        _onlyValidValidator(validator);
        _;
    }

    function _onlyValidValidator(address validator) internal view {
        require(validatorStakes[validator].isRegistered, "Validator not registered");
        require(validatorStakes[validator].selfStake >= proposalContract.minValidatorStake(), "Not a valid validator");
    }

    modifier onlyActiveValidator(address validator) {
        _onlyActiveValidator(validator);
        _;
    }

    function _onlyActiveValidator(address validator) internal view {
        require(validatorStakes[validator].selfStake >= proposalContract.minValidatorStake(), "Not a valid validator");
        require(!validatorStakes[validator].isJailed, "Validator is jailed");
    }

    /**
     * @dev Initializes the Staking contract with required dependencies.
     * @param validators_ Address of the Validators contract.
     * @param proposal_ Address of the Proposal contract.
     */
    function initialize(address validators_, address proposal_) external onlyNotInitialized {
        require(validators_ != address(0), "Invalid validators address");
        require(proposal_ != address(0), "Invalid proposal address");
        
        validatorsContract = IValidators(validators_);
        proposalContract = IProposal(proposal_);
        initialized = true;
    }

    /**
     * @dev Initialize with pre-registered validators (for genesis deployment)
     * @param validators_ Validators contract address
     * @param proposal_ Proposal contract address
     * @param initialValidators Array of validator addresses to pre-register
     * @param commissionRate Default commission rate for all validators
     * @notice This function automatically performs the same logic as registerValidator for genesis validators
     * @notice Genesis validators are pre-registered with default stake and automatically added to highestValidatorsSet
     * @notice They don't need to pass proposal or wait for 7-day window - they are activated immediately
     */
    function initializeWithValidators(
        address validators_,
        address proposal_,
        address[] calldata initialValidators,
        uint256 commissionRate
    ) external onlyNotInitialized {
        require(validators_ != address(0), "Invalid validators address");
        require(proposal_ != address(0), "Invalid proposal address");
        require(initialValidators.length > 0, "No validators provided");
        require(commissionRate > 0, "Commission rate must be greater than 0");
        require(commissionRate < COMMISSION_RATE_BASE, "Commission rate exceeds maximum allowed");
        
        validatorsContract = IValidators(validators_);
        proposalContract = IProposal(proposal_);
        
        // Cache the minimum validator stake outside the loop
        uint256 minValidatorStake = proposalContract.minValidatorStake();
        
        // Pre-register all initial validators with default stake
        // This automatically performs the same logic as registerValidator for genesis validators
        for (uint256 i = 0; i < initialValidators.length; i++) {
            address validator = initialValidators[i];
            require(validator != address(0), "Invalid validator address");
            require(!validatorStakes[validator].isRegistered, "Validator already exists");
            
            // Set up validator stake (same as registerValidator)
            validatorStakes[validator] = ValidatorStake({
                selfStake: minValidatorStake,
                totalDelegated: 0,
                commissionRate: commissionRate,
                accumulatedRewards: 0,
                isJailed: false,
                jailUntilBlock: 0,
                totalClaimedRewards: 0,
                lastClaimBlock: 0,
                isRegistered: true
            });
            
            // Add to validators list (same as registerValidator)
            allValidators.push(validator);
            totalStaked = totalStaked + minValidatorStake;
            // Emit event for each validator (more accurate than single event)
            emit ValidatorRegistered(validator, minValidatorStake, commissionRate);
        }
        
        initialized = true;
    }

    /**
     * @dev Register as a validator with self-stake
     * @param commissionRate Commission rate (0-10000, representing 0%-100%)
     */
    function registerValidator(uint256 commissionRate) external payable onlyInitialized nonReentrant {
        require(commissionRate > 0, "Commission rate must be greater than 0");
        require(commissionRate <= COMMISSION_RATE_BASE, "Commission rate exceeds maximum allowed");
        require(!validatorStakes[msg.sender].isRegistered, "Already registered");
        require(proposalContract.pass(msg.sender), "Must pass proposal first");
        // Check if proposal is still valid (within 7 days)
        require(proposalContract.isProposalValidForStaking(msg.sender), "Proposal expired, must repropose");
        // Check if staking amount is sufficient
        require(msg.value >= proposalContract.minValidatorStake(), "Insufficient self-stake");

        validatorStakes[msg.sender] = ValidatorStake({
            selfStake: msg.value,
            totalDelegated: 0,
            commissionRate: commissionRate,
            accumulatedRewards: 0,
            isJailed: false,
            jailUntilBlock: 0,
            totalClaimedRewards: 0,
            lastClaimBlock: 0,
            isRegistered: true
        });

        allValidators.push(msg.sender);
        totalStaked = totalStaked + msg.value;

        // Activate validator (register = activate)
        // tryActive will add to highestValidatorsSet and emit LogActive event
        require(validatorsContract.tryActive(msg.sender), "Validator activation failed");

        emit ValidatorRegistered(msg.sender, msg.value, commissionRate);
    }

    /**
     * @dev Add more self-stake to existing validator
     */
    function addValidatorStake() external payable nonReentrant {
        require(msg.value > 0, "Amount must be positive");
        require(validatorStakes[msg.sender].isRegistered, "Validator not registered");
        validatorStakes[msg.sender].selfStake = validatorStakes[msg.sender].selfStake + msg.value;
        totalStaked = totalStaked + msg.value;
        emit ValidatorStakeIncreased(msg.sender, msg.sender, msg.value);
    }

    /**
     * @dev Update validator commission rate
     * @param newCommissionRate New commission rate (0-10000)
     */
    function updateCommissionRate(uint256 newCommissionRate) external onlyValidValidator(msg.sender) nonReentrant {
        require(newCommissionRate > 0, "Commission rate must be greater than 0");
        require(newCommissionRate < COMMISSION_RATE_BASE, "Commission rate exceeds maximum allowed");
        
        validatorStakes[msg.sender].commissionRate = newCommissionRate;
        emit CommissionRateUpdated(msg.sender, newCommissionRate);
    }

    /**
     * @dev Decrease validator's self-stake
     * @param amount Amount to reduce from self-stake
     * @notice If remaining stake would be less than MIN_VALIDATOR_STAKE, the reduction will fail
     * @notice Use withdrawAfterResignation() to withdraw all stake (requires minimum validators requirement)
     */
    function decreaseValidatorStake(uint256 amount) external nonReentrant onlyValidValidator(msg.sender) {
        require(amount > 0, "Amount must be positive");
        
        ValidatorStake storage stake = validatorStakes[msg.sender];
        require(stake.selfStake >= amount, "Insufficient self-stake");
        
        // Calculate remaining stake after reduction
        uint256 remainingStake = stake.selfStake - amount;
        
        // If partial reduction, remaining stake must meet minimum requirement
        // If complete reduction (remainingStake == 0), use exitValidator() instead
        require(remainingStake >= proposalContract.minValidatorStake(), "Remaining stake below minimum, withdraw all stake instead");
        
        // Effects: update state before external call
        stake.selfStake = remainingStake;
        totalStaked = totalStaked - amount;
        
        // Instead of transferring funds directly, add them to unbonding like exitValidator and undelegate
        unbondingDelegations[msg.sender][msg.sender].push(UnbondingEntry({
            amount: amount,
            completionBlock: block.number + proposalContract.unbondingPeriod()
        }));
        
        emit ValidatorStakeDecreased(msg.sender, msg.sender, amount);
    }

    /**
     * @dev Allow validator to resign from validator role
     * @notice This function allows validators to voluntarily resign from their validator role
     * @notice Validator will be removed from highestValidatorsSet and pass will be set to false
     * @notice Validator's transaction fee income (aacIncoming) will NOT be removed
     * @notice After resigning, validator will be excluded from currentValidatorSet at next epoch
     * @notice After being excluded, validator can call emergencyExit() to withdraw all stake
     */
    function resignValidator() external onlyValidValidator(msg.sender) nonReentrant {
        ValidatorStake storage stake = validatorStakes[msg.sender];
        
        // Cannot resign if already jailed/resigned
        require(!stake.isJailed, "Validator already resigned or jailed");
        
        // Mark validator as jailed to ensure exclusion from currentValidatorSet at next epoch
        // This is a technical mechanism to ensure smooth exit, not a punishment
        stake.isJailed = true;
        stake.jailUntilBlock = block.number + proposalContract.validatorUnjailPeriod();
        emit ValidatorJailed(msg.sender, stake.jailUntilBlock);
        
        // Remove from highestValidatorsSet and set pass = false
        // Note: This does NOT remove transaction fee income (aacIncoming)
        validatorsContract.removeFromHighestSet(msg.sender);
    }

    /**
     * @dev Exit validator - withdraw all stake and exit validator role
     * @notice Validators in currentValidatorSet (active in current epoch) cannot exit
     * @notice Validators must jail themselves first, then wait until next epoch to exit
     * @notice This ensures smooth exit without disrupting consensus
     */
    function exitValidator() external nonReentrant onlyValidValidator(msg.sender) {
        ValidatorStake storage stake = validatorStakes[msg.sender];
        uint256 withdrawAmount = stake.selfStake;
        require(withdrawAmount > 0, "Validator has no stake to withdraw");
        
        // Check if validator is currently participating in consensus (in currentValidatorSet)
        // Note: isActiveValidator() checks currentValidatorSet, not jail status
        bool isInCurrentSet = validatorsContract.isActiveValidator(msg.sender);
        
        // Validators in active set cannot exit - they must resign first
        // After resigning, they will be excluded from currentValidatorSet at next epoch
        require(!isInCurrentSet, "Cannot exit: validator is in active set, resign first and wait until next epoch");
        
        // Check if we need to remove from highest set before state changes
        bool needRemoveFromHighestSet = proposalContract.pass(msg.sender);
        
        // Effects: update all state variables first
        stake.selfStake = 0;
        totalStaked = totalStaked - withdrawAmount;
        
        // Instead of transferring funds directly, add them to unbonding like undelegate
        unbondingDelegations[msg.sender][msg.sender].push(UnbondingEntry({
            amount: withdrawAmount,
            completionBlock: block.number + proposalContract.unbondingPeriod()
        }));
        
        // Interactions: external call after state updates
        if (needRemoveFromHighestSet) {
            // Validator didn't call resignValidator(), so we need to remove them now
            validatorsContract.removeFromHighestSet(msg.sender);
        }

        emit ValidatorExited(msg.sender);
    }

    /**
     * @dev Delegate tokens to a validator
     * @param validator Validator address to delegate to
     */
    function delegate(address validator) external payable onlyActiveValidator(validator) nonReentrant {
        require(validator != address(0), "Invalid validator address");
        require(msg.value >= proposalContract.minDelegation(), "Insufficient delegation amount");
        require(validator != msg.sender, "Cannot delegate to yourself");
        
        // Calculate pending rewards
        uint256 pending = _getPendingRewards(msg.sender, validator);
        
        // Check if this is the first delegation to this validator
        bool isFirstForValidator = delegations[msg.sender][validator].amount == 0;
        
        // Update delegation amount
        delegations[msg.sender][validator].amount += msg.value;
        validatorStakes[validator].totalDelegated += msg.value;
        totalStaked += msg.value;
        
        // Update reward debt
        delegations[msg.sender][validator].rewardDebt = delegations[msg.sender][validator].amount * rewardPerShare[validator] / 1e18;

        // Update delegator counts
        if (isFirstForValidator) {
            // Increment validator's delegator count
            validatorDelegatorCount[validator]++;
            
            // Check if this is the first delegation for the delegator
            if (!delegatorExists[msg.sender]) {
                delegatorCount++;
                delegatorExists[msg.sender] = true;
            }
        }

        // Send pending rewards if any
        if (pending > 0) {
            _claimPending(msg.sender, validator, pending);
        }
        
        emit Delegated(msg.sender, validator, msg.value);
    }

    /**
     * @dev Start unbonding delegation from a validator
     * @param validator Validator address to undelegate from
     * @param amount Amount to undelegate
     */
    function undelegate(address validator, uint256 amount) external nonReentrant {
        require(validator != address(0), "Invalid validator address");
        require(amount > 0, "Amount must be positive");
        require(amount >= proposalContract.minUndelegation(), "Insufficient undelegation amount");
        require(unbondingDelegations[msg.sender][validator].length < MAX_UNBONDING_ENTRIES, "Too many unbonding entries");
        require(validator != msg.sender, "Cannot undelegate from yourself");
        require(delegations[msg.sender][validator].amount >= amount, "Insufficient delegation");
        // Allow undelegation even if validator has exited (selfStake == 0)
        // This ensures delegators can always withdraw their funds
        
        // Step 1: Calculate pending rewards (pure view operation, no state changes)
        uint256 pending = _getPendingRewards(msg.sender, validator);
        
        // Check if this will result in no delegation left for this validator
        bool willRemoveFromValidator = delegations[msg.sender][validator].amount == amount;
        
        // Step 2: Effects - Update all state variables before external calls
        delegations[msg.sender][validator].amount -= amount;
        validatorStakes[validator].totalDelegated -= amount;
        totalStaked -= amount;
        
        // Update reward debt based on new delegation amount
        delegations[msg.sender][validator].rewardDebt = delegations[msg.sender][validator].amount * rewardPerShare[validator] / 1e18;
        
        // Update delegator counts if all delegation is removed
        if (willRemoveFromValidator) {
            // Decrement validator's delegator count
            validatorDelegatorCount[validator]--;
            
            // Check if delegator has any remaining delegations
            bool hasRemainingDelegations = false;
            for (uint256 i = 0; i < allValidators.length; i++) {
                address v = allValidators[i];
                if (v != validator && delegations[msg.sender][v].amount > 0) {
                    hasRemainingDelegations = true;
                    break;
                }
            }
            
            // If no remaining delegations, update delegation status
            if (!hasRemainingDelegations) {
                delegatorCount--;
                delete delegatorExists[msg.sender];
            }
        }
        
        // Add to unbonding
        unbondingDelegations[msg.sender][validator].push(UnbondingEntry({
            amount: amount,
            completionBlock: block.number + proposalContract.unbondingPeriod()
        }));
        
        // Step 3: Interactions - Send pending rewards last
        if (pending > 0) {
            _claimPending(msg.sender, validator, pending);
        }
        
        emit Undelegated(msg.sender, validator, amount);
    }

    /**
     * @dev Complete unbonding and withdraw tokens
     * @param validator Validator address
     * @param maxEntries Maximum number of unbonding entries to process
     */
    function withdrawUnbonded(address validator, uint256 maxEntries) external nonReentrant {
        require(maxEntries > 0, "maxEntries must be positive");
        require(maxEntries <= MAX_UNBONDING_ENTRIES, "maxEntries too large");
        
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
        
        // Check if we need to delete the delegation before external call
        bool shouldDeleteDelegation = false;
        if (delegations[msg.sender][validator].amount == 0 && 
            unbondingDelegations[msg.sender][validator].length == 0) {
            shouldDeleteDelegation = true;
        }
        
        // Effects: delete delegation entry if needed before external call
        if (shouldDeleteDelegation) {
            delete delegations[msg.sender][validator];
        }
        
        // Interactions: external call after all state updates
        (bool success, ) = payable(msg.sender).call{value: totalWithdraw}("");
        require(success, "Transfer failed");
        emit UnbondingCompleted(msg.sender, validator, totalWithdraw);
    }

    /**
     * @dev Distribute rewards to the current block miner (validator)
     * @notice Validator address is obtained from block.coinbase
     * @notice Block reward is passed via msg.value (consensus layer reads from Proposal contract)
     * @notice Jailed validators can still produce blocks and receive rewards in the current epoch
     * @notice They will be excluded from the validator set at the next epoch transition
     */
    function distributeRewards() external payable onlyMiner onlyInitialized nonReentrant {
        // Check if block reward has already been distributed for this block
        if (operationsDone[block.number][uint8(Operations.Distribute)]) {
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
     * @dev Claim pending rewards for delegators
     * @param validator Validator address from which to claim rewards
     * @notice Delegators can claim rewards earned from their delegations
     * @notice Uses nonReentrant modifier to prevent reentrancy attacks
     * @notice Follows Checks-Effects-Interactions pattern
     */
    function claimRewards(address validator) external nonReentrant {
        // Checks: Verify validator is registered
        require(validatorStakes[validator].isRegistered, "Validator not registered");
        
        // Checks: Verify caller has delegation to this validator
        require(delegations[msg.sender][validator].amount > 0, "No delegation found");
        
        // Get current pending rewards
        uint256 pending = _getPendingRewards(msg.sender, validator);

        // Effects: Update state before external call to prevent reentrancy
        Delegation storage delegation = delegations[msg.sender][validator];
        if (pending > 0) {
            delegation.rewardDebt = delegation.amount
                * rewardPerShare[validator]
                / 1e18;
            // Interactions: External call after all state updates
            _claimPending(msg.sender, validator, pending);
        }
    }


    /**
     * @dev Claim accumulated rewards for validators
     * @notice Validators can claim their accumulated commission rewards
     * @notice Validators must wait withdrawProfitPeriod blocks between claims
     * @notice First claim is always allowed (when lastClaimBlock == 0)
     * @notice Uses nonReentrant modifier to prevent reentrancy attacks
     * @notice Follows Checks-Effects-Interactions pattern
     * @notice Emits RewardsClaimed event upon successful claim
     */
    function claimValidatorRewards() external nonReentrant {
        address validator = msg.sender;
        
        // Checks: Verify caller is a registered validator
        require(validatorStakes[validator].isRegistered, "Not a registered validator");
        
        // Get current accumulated commission rewards
        ValidatorStake storage stake = validatorStakes[validator];
        uint256 commission = stake.accumulatedRewards;
        
        if (commission > 0) {
            // Checks: Verify withdraw period has passed (if not first claim)
            if (stake.lastClaimBlock > 0) {
                uint256 withdrawPeriod = proposalContract.withdrawProfitPeriod();
                require(
                    block.number >= stake.lastClaimBlock + withdrawPeriod,
                    "Must wait withdrawProfitPeriod blocks between claims"
                );
            }
            
            // Effects: Update all state variables before external calls
            stake.accumulatedRewards = 0;
            stake.totalClaimedRewards = stake.totalClaimedRewards + commission;
            stake.lastClaimBlock = block.number;
            
            // Interactions: External call for commission after all state updates
            _claimPending(msg.sender, validator, commission);
        }
    }

    /**
     * @dev Jail a validator for misbehavior (called by Punish or Validators contract)
     * @param validator Validator address to jail
     * @param jailBlocks Number of blocks to jail for
     * @notice Only Punish or Validators contract can call this function
     * @notice Validators should use resignValidator() to voluntarily resign from validator role
     */
    function jailValidator(address validator, uint256 jailBlocks) external onlyPunishOrValidatorsContract nonReentrant {
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
    function unjailValidator(address validator) external nonReentrant {
        require(validator != address(0), "Invalid validator address");
        require(msg.sender == validator, "Only validator can unjail themselves");
        require(validatorStakes[validator].isJailed, "Validator not jailed");
        require(block.number >= validatorStakes[validator].jailUntilBlock, "Jail period not complete");
        
        // Check if validator has sufficient stake to be active
        require(validatorStakes[validator].selfStake >= proposalContract.minValidatorStake(), "Insufficient stake, must add stake first");
        
        // Must have passed proposal (reproposal required) - no automatic restoration
        require(proposalContract.pass(validator), "Must pass reproposal first");
        
        // Update state first (effects)
        validatorStakes[validator].isJailed = false;
        validatorStakes[validator].jailUntilBlock = 0;
        
        // Interactions: Activate validator after state update
        // This ensures validator is added to highestValidatorsSet after state change
        require(validatorsContract.tryActive(validator), "Failed to activate validator");
        
        emit ValidatorUnjailed(validator);
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
        
        // Cache the minimum validator stake outside the loop
        uint256 minValidatorStake = proposalContract.minValidatorStake();
        
        // Collect validators and their total stakes for sorting
        for (uint256 i = 0; i < validators.length; i++) {
            address validator = validators[i];
            ValidatorStake storage stake = validatorStakes[validator];
            
            // Only include validators with self-stake >= MIN_VALIDATOR_STAKE
            if (stake.selfStake >= minValidatorStake) {
                candidateValidators[candidateCount] = validator;
                totalStakes[candidateCount] = stake.selfStake + stake.totalDelegated;
                candidateCount++;
            }
        }
        
        if (candidateCount == 0) {
            return new address[](0);
        }
        
        // Sort by total stake using heap sort for improved performance
        // Build max heap
        for (uint256 i = candidateCount / 2; i > 0; i--) {
            heapify(candidateValidators, totalStakes, candidateCount, i - 1);
        }
        
        // Extract max and heapify - this puts elements in ascending order (smallest at beginning)
        for (uint256 i = candidateCount - 1; i > 0; i--) {
            // Swap current maximum to end
            (candidateValidators[0], candidateValidators[i]) = (candidateValidators[i], candidateValidators[0]);
            (totalStakes[0], totalStakes[i]) = (totalStakes[i], totalStakes[0]);
            
            // Heapify the reduced heap
            heapify(candidateValidators, totalStakes, i, 0);
        }
        
        // Reverse the array to get descending order (largest at beginning)
        for (uint256 i = 0; i < candidateCount / 2; i++) {
            uint256 j = candidateCount - 1 - i;
            // Swap validator addresses
            (candidateValidators[i], candidateValidators[j]) = (candidateValidators[j], candidateValidators[i]);
            // Swap corresponding stakes
            (totalStakes[i], totalStakes[j]) = (totalStakes[j], totalStakes[i]);
        }
        
        // Return top validators (up to MAX_VALIDATORS)
        uint256 returnLength = candidateCount < proposalContract.maxValidators() ? candidateCount : proposalContract.maxValidators();
        address[] memory topValidators = new address[](returnLength);
        for (uint256 i = 0; i < returnLength; i++) {
            topValidators[i] = candidateValidators[i];
        }
        
        return topValidators;
    }
    
    /**
     * @dev Heapify function to maintain max heap property
     * @param arr Array of validator addresses to heapify
     * @param stakes Array of corresponding total stakes
     * @param n Size of the heap
     * @param i Index of the current root
     */
    function heapify(address[] memory arr, uint256[] memory stakes, uint256 n, uint256 i) internal pure {
        uint256 largest = i;
        uint256 left = 2 * i + 1;
        uint256 right = 2 * i + 2;
        
        // Find largest among root, left child and right child
        if (left < n && stakes[left] > stakes[largest]) {
            largest = left;
        }
        
        if (right < n && stakes[right] > stakes[largest]) {
            largest = right;
        }
        
        // If largest is not root, swap and continue heapifying
        if (largest != i) {
            (arr[i], arr[largest]) = (arr[largest], arr[i]);
            (stakes[i], stakes[largest]) = (stakes[largest], stakes[i]);
            
            heapify(arr, stakes, n, largest);
        }
    }

    function _getPendingRewards(address delegator, address validator) internal view returns (uint256){
        Delegation storage delegation = delegations[delegator][validator];
        if (delegation.amount > 0) {
            return delegation.amount
                * rewardPerShare[validator]
                / 1e18
                - delegation.rewardDebt;
        }
        return 0;
    }
    function _claimPending(address delegator, address validator, uint256 pending) internal {
        if (pending > 0) {
            // Check if contract has enough balance before transfer
            require(address(this).balance >= pending, "Insufficient contract balance");
            // Interactions: external call after state update
            (bool success, ) = payable(delegator).call{value: pending}("");
            require(success, "Transfer failed");
                
            emit RewardsClaimed(delegator, validator, pending);
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
     * @return isRegistered Whether validator is registered
     */
    function getValidatorInfo(address validator) external view returns (
        uint256 selfStake,
        uint256 totalDelegated,
        uint256 commissionRate,
        uint256 accumulatedRewards,
        bool isJailed,
        uint256 jailUntilBlock,
        uint256 totalClaimedRewards,
        uint256 lastClaimBlock,
        bool isRegistered
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
            stake.lastClaimBlock,
            stake.isRegistered
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
     */
    function getDelegationInfo(address delegator, address validator) external view returns (
        uint256 amount,
        uint256 pendingRewards
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
            pending
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
