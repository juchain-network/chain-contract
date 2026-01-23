// SPDX-License-Identifier: MIT

pragma solidity ^0.8.29;

interface IStaking {
    /**
     * @dev Gets the status of a validator.
     * @param validator Address of the validator to check.
     * @return isActive Whether the validator is active.
     * @return isJailed Whether the validator is jailed.
     */
    function getValidatorStatus(address validator) external view returns (bool isActive, bool isJailed);

    /**
     * @dev Gets the top validators based on stake.
     * @param validators List of validators to evaluate.
     * @return address[] memory List of top validators sorted by stake.
     */
    function getTopValidators(address[] memory validators) external view returns (address[] memory);

    /**
     * @dev Checks if a validator is jailed.
     * @param validator Address of the validator to check.
     * @return bool Returns true if the validator is jailed.
     */
    function isValidatorJailed(address validator) external view returns (bool);

    /**
     * @dev Gets detailed information about a validator.
     * @param validator Address of the validator to check.
     * @return selfStake Amount of tokens staked by the validator.
     * @return totalDelegated Total amount of tokens delegated to the validator.
     * @return commissionRate Commission rate charged by the validator (in precision units).
     * @return accumulatedRewards Total rewards accumulated by the validator.
     * @return isJailed Whether the validator is currently jailed.
     * @return jailUntilBlock Block number until which the validator is jailed.
     * @return totalClaimedRewards Total rewards claimed by the validator.
     * @return lastClaimBlock Block number at which the validator last claimed rewards.
     * @return isRegistered Whether the validator is currently registered.
     * @return totalRewards Total rewards earned by distribution (cumulative).
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
        bool isRegistered,
        uint256 totalRewards
    );

    /**
     * @dev Jails a validator for a specified number of blocks.
     * @param validator Address of the validator to jail.
     * @param jailBlocks Number of blocks the validator will be jailed.
     */
    function jailValidator(address validator, uint256 jailBlocks) external;

    /**
     * @dev Slash a validator's self-stake and distribute reward/burn amounts.
     * @param validator Validator address to slash
     * @param slashAmount Absolute slash amount in wei
     * @param reporter Address to receive the reporter reward
     * @param rewardAmount Reporter reward amount in wei
     * @param burnAddress Address to receive the burn amount
     * @return actualSlash Actual slashed amount
     * @return actualReward Actual reward paid to reporter
     */
    function slashValidator(
        address validator,
        uint256 slashAmount,
        address reporter,
        uint256 rewardAmount,
        address burnAddress
    ) external returns (uint256 actualSlash, uint256 actualReward);


}
