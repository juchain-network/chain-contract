// SPDX-License-Identifier: MIT

pragma solidity ^0.8.29;

interface IProposal {
    /**
     * @dev Sets a validator as unpassed (ineligible).
     * @param val The address of the validator to mark as unpassed.
     * @return bool Returns true if the operation was successful.
     */
    function setUnpassed(address val) external returns (bool);

    /**
     * @dev Checks if a validator's proposal is valid for staking.
     * @param validator The address of the validator to check.
     * @return bool Returns true if the validator's proposal is valid for staking.
     */
    function isProposalValidForStaking(address validator) external view returns (bool);

    /**
     * @dev Checks if a validator has passed the proposal process and is eligible to participate.
     * @param validator The address of the validator to check.
     * @return bool Returns true if the validator has passed the proposal process.
     */
    function pass(address validator) external view returns (bool);

    /**
     * @dev Returns the threshold for punishing validators based on missed blocks.
     * @return uint256 The number of missed blocks required to trigger a punishment.
     */
    function punishThreshold() external view returns (uint256);

    /**
     * @dev Returns the threshold for removing validators based on missed blocks.
     * @return uint256 The number of missed blocks required to trigger removal from the validator set.
     */
    function removeThreshold() external view returns (uint256);

    /**
     * @dev Returns the rate at which missed blocks counter decreases.
     * @return uint256 The decrease rate for the missed blocks counter.
     */
    function decreaseRate() external view returns (uint256);

    /**
     * @dev Returns the period (in blocks) for withdrawing validator profits.
     * @return uint256 The profit withdrawal period in blocks.
     */
    function withdrawProfitPeriod() external view returns (uint256);

    /**
     * @dev Returns the reward amount for each block produced.
     * @return uint256 The block reward amount.
     */
    function blockReward() external view returns (uint256);

    /**
     * @dev Returns the unbonding period for delegators (in blocks).
     * @return uint256 The unbonding period in blocks.
     */
    function unbondingPeriod() external view returns (uint256);

    /**
     * @dev Returns the period (in blocks) a validator must wait to unjail after being jailed.
     * @return uint256 The validator unjail period in blocks.
     */
    function validatorUnjailPeriod() external view returns (uint256);

    /**
     * @dev Returns the minimum staking amount required to become a validator.
     * @return uint256 The minimum validator stake amount in wei.
     */
    function minValidatorStake() external view returns (uint256);

    /**
     * @dev Returns the maximum number of validators allowed in the active set.
     * @return uint256 The maximum number of validators.
     */
    function maxValidators() external view returns (uint256);

    /**
     * @dev Returns the epoch duration in blocks.
     * @return uint256 The epoch duration.
     */
    function epoch() external view returns (uint256);

    /**
     * @dev Returns the minimum delegation amount per delegator.
     * @return uint256 The minimum delegation amount in wei.
     */
    function minDelegation() external view returns (uint256);

    /**
     * @dev Returns the minimum undelegation amount per delegator.
     * @return uint256 The minimum undelegation amount in wei.
     */
    function minUndelegation() external view returns (uint256);

    /**
     * @dev Returns the double-sign slash amount (absolute, in wei).
     * @return uint256 The slash amount.
     */
    function doubleSignSlashAmount() external view returns (uint256);

    /**
     * @dev Returns the double-sign reporter reward amount (absolute, in wei).
     * @return uint256 The reward amount.
     */
    function doubleSignRewardAmount() external view returns (uint256);

    /**
     * @dev Returns the double-sign evidence window (in blocks).
     * @return uint256 The evidence window in blocks.
     */
    function doubleSignWindow() external view returns (uint256);

    /**
     * @dev Returns the commission update cooldown (in blocks).
     * @return uint256 The cooldown in blocks.
     */
    function commissionUpdateCooldown() external view returns (uint256);

    /**
     * @dev Returns the burn address for slashed funds after reward.
     * @return address The burn address.
     */
    function burnAddress() external view returns (address);

    /**
     * @dev Returns the maximum commission rate (basis points).
     * @return uint256 The maximum commission rate.
     */
    function maxCommissionRate() external view returns (uint256);
}
