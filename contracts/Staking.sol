// SPDX-License-Identifier: MIT

pragma solidity ^0.8.20;

import './Params.sol';
import './library/SafeMath.sol';

// Interface for Validators contract to avoid circular dependency
interface IValidators {
    function updateValidatorSetByStake(uint256 epoch) external;
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
    uint256 public constant MIN_VALIDATORS = 5;
    
    // Commission rate precision (10000 = 100%)
    uint256 public constant COMMISSION_RATE_BASE = 10000;
    
    // Unbonding period in blocks (approximately 7 days)
    uint256 public constant UNBONDING_PERIOD = 604800; // 7 days * 24 hours * 3600 seconds / 1 second per block

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
        require(!validatorStakes[validator].isJailed || block.number >= validatorStakes[validator].jailUntilBlock, "Validator is jailed");
        _;
    }

    function initialize() external onlyNotInitialized {
        validatorsContract = IValidators(ValidatorContractAddr);
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

        emit ValidatorRegistered(msg.sender, msg.value, commissionRate);
    }

    /**
     * @dev Add more self-stake to existing validator
     */
    function addValidatorStake() external payable onlyValidValidator(msg.sender) {
        require(msg.value > 0, "Amount must be positive");
        
        validatorStakes[msg.sender].selfStake = validatorStakes[msg.sender].selfStake.add(msg.value);
        totalStaked = totalStaked.add(msg.value);
    }

    /**
     * @dev Update validator commission rate
     * @param newCommissionRate New commission rate (0-10000)
     */
    function updateCommissionRate(uint256 newCommissionRate) external onlyValidValidator(msg.sender) {
        require(newCommissionRate <= COMMISSION_RATE_BASE, "Invalid commission rate");
        
        validatorStakes[msg.sender].commissionRate = newCommissionRate;
        emit ValidatorUpdated(msg.sender, newCommissionRate);
    }

    /**
     * @dev Start validator stake withdrawal (validator exit)
     * @param amount Amount to withdraw from self-stake
     */
    function withdrawValidatorStake(uint256 amount) external onlyValidValidator(msg.sender) {
        require(amount > 0, "Amount must be positive");
        
        ValidatorStake storage stake = validatorStakes[msg.sender];
        require(stake.selfStake >= amount, "Insufficient self-stake");
        
        // Calculate remaining stake after withdrawal
        uint256 remainingStake = stake.selfStake.sub(amount);
        
        // If this would make validator inactive, check minimum validator requirement
        if (remainingStake < MIN_VALIDATOR_STAKE) {
            uint256 activeValidatorCount = getActiveValidatorCount();
            require(activeValidatorCount > MIN_VALIDATORS, "Cannot exit: minimum validators required");
        }
        
        // If partial withdrawal, ensure remaining stake meets minimum
        if (remainingStake > 0) {
            require(remainingStake >= MIN_VALIDATOR_STAKE, "Remaining stake below minimum");
        }
        
        stake.selfStake = remainingStake;
        totalStaked = totalStaked.sub(amount);
        
        // If validator becomes inactive, handle cleanup
        if (remainingStake < MIN_VALIDATOR_STAKE) {
            // Note: We don't remove from allValidators array to preserve indices
            // Validator will be filtered out in getTopValidators and other functions
            emit ValidatorExited(msg.sender, amount);
        } else {
            emit ValidatorStakeWithdrawn(msg.sender, amount);
        }
        
        // Transfer the withdrawn amount
        payable(msg.sender).transfer(amount);
    }

    /**
     * @dev Emergency exit for validator (withdraw all stake)
     */
    function emergencyExit() external onlyValidValidator(msg.sender) {
        uint256 activeValidatorCount = getActiveValidatorCount();
        require(activeValidatorCount > MIN_VALIDATORS, "Cannot exit: minimum validators required");
        
        ValidatorStake storage stake = validatorStakes[msg.sender];
        uint256 withdrawAmount = stake.selfStake;
        
        stake.selfStake = 0;
        totalStaked = totalStaked.sub(withdrawAmount);
        
        emit ValidatorExited(msg.sender, withdrawAmount);
        payable(msg.sender).transfer(withdrawAmount);
    }

    /**
     * @dev Delegate tokens to a validator
     * @param validator Validator address to delegate to
     */
    function delegate(address validator) external payable onlyActiveValidator(validator) {
        require(validator != address(0), "Invalid validator address");
        require(msg.value >= MIN_DELEGATION, "Insufficient delegation amount");
        
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
    function undelegate(address validator, uint256 amount) external onlyValidValidator(validator) {
        require(validator != address(0), "Invalid validator address");
        require(amount > 0, "Amount must be positive");
        require(delegations[msg.sender][validator].amount >= amount, "Insufficient delegation");
        
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
        UnbondingEntry[] storage entries = unbondingDelegations[msg.sender][validator];
        uint256 totalWithdraw = 0;
        uint256 processed = 0;
        
        for (uint256 i = 0; i < entries.length && processed < maxEntries; i++) {
            if (entries[i].completionBlock <= block.number) {
                totalWithdraw = totalWithdraw.add(entries[i].amount);
                
                // Remove completed entry
                entries[i] = entries[entries.length - 1];
                entries.pop();
                i--; // Adjust index since we removed an element
                processed++;
            }
        }
        
        require(totalWithdraw > 0, "No unbonded tokens available");
        
        payable(msg.sender).transfer(totalWithdraw);
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
        
        // Calculate commission
        uint256 commission = msg.value.mul(stake.commissionRate).div(COMMISSION_RATE_BASE);
        uint256 delegatorRewards = msg.value.sub(commission);
        
        // Add commission to validator's accumulated rewards
        stake.accumulatedRewards = stake.accumulatedRewards.add(commission);
        
        // Update reward per share for delegators
        if (stake.totalDelegated > 0) {
            rewardPerShare[validator] = rewardPerShare[validator].add(
                delegatorRewards.mul(1e18).div(stake.totalDelegated)
            );
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
                validatorStakes[validator].accumulatedRewards = 0; // Set to 0 before transfer
                payable(msg.sender).transfer(commission);
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
        
        validatorStakes[validator].isJailed = false;
        validatorStakes[validator].jailUntilBlock = 0;
        emit ValidatorUnjailed(validator);
    }

    /**
     * @dev Get top validators by total stake
     * @param limit Maximum number of validators to return
     * @return Top validators list
     */
    function getTopValidators(uint256 limit) external view returns (address[] memory) {
        if (limit > MAX_VALIDATORS) {
            limit = MAX_VALIDATORS;
        }
        
        // Create array of all active validators with their total stake
        address[] memory activeValidators = new address[](allValidators.length);
        uint256[] memory totalStakes = new uint256[](allValidators.length);
        uint256 activeCount = 0;
        
        for (uint256 i = 0; i < allValidators.length; i++) {
            address validator = allValidators[i];
            if (validatorStakes[validator].selfStake >= MIN_VALIDATOR_STAKE && 
                (!validatorStakes[validator].isJailed || block.number >= validatorStakes[validator].jailUntilBlock)) {
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
        
        // Return top validators
        uint256 returnLength = activeCount < limit ? activeCount : limit;
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
                // Update reward debt first to prevent reentrancy
                delegation.rewardDebt = delegation.amount
                    .mul(rewardPerShare[validator])
                    .div(1e18);
                    
                payable(delegator).transfer(pending);
                emit RewardsClaimed(delegator, validator, pending);
            }
        }
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
    function getActiveValidatorCount() public view returns (uint256) {
        uint256 activeCount = 0;
        
        for (uint256 i = 0; i < allValidators.length; i++) {
            address validator = allValidators[i];
            if (validatorStakes[validator].selfStake >= MIN_VALIDATOR_STAKE && 
                (!validatorStakes[validator].isJailed || block.number >= validatorStakes[validator].jailUntilBlock)) {
                activeCount++;
            }
        }
        
        return activeCount;
    }

    /**
     * @dev Check if system has minimum required validators
     * @return Whether system meets minimum validator requirement
     */
    function hasMinimumValidators() external view returns (bool) {
        return getActiveValidatorCount() >= MIN_VALIDATORS;
    }
}
