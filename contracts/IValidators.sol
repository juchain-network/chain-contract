// SPDX-License-Identifier: MIT

pragma solidity ^0.8.29;

interface IValidators {
    /**
     * @dev Gets the effective signer for a validator.
     * @param validator Validator cold address.
     * @return signer Effective hot signer address.
     */
    function getValidatorSigner(address validator) external view returns (address signer);

    /**
     * @dev Resolves a signer address to its validator.
     * @param signer Signer hot address.
     * @return validator Validator cold address, or zero if not found.
     */
    function getValidatorBySigner(address signer) external view returns (address validator);

    /**
     * @dev Gets the pending signer rotation record for a validator.
     * @param validator Validator cold address.
     * @return signer Pending signer hot address.
     * @return effectiveBlock Checkpoint block after which the pending signer becomes runtime-effective.
     * @return pending True when a consistent pending rotation record is still stored on-chain.
     */
    function getPendingValidatorSigner(address validator)
        external
        view
        returns (address signer, uint256 effectiveBlock, bool pending);

    /**
     * @dev Gets the pending signer rotation record by signer address.
     * @param signer Pending signer hot address.
     * @return validator Validator cold address that reserved the signer.
     * @return effectiveBlock Checkpoint block after which the pending signer becomes runtime-effective.
     * @return pending True when a consistent pending rotation record is still stored on-chain.
     */
    function getPendingValidatorBySigner(address signer)
        external
        view
        returns (address validator, uint256 effectiveBlock, bool pending);

    /**
     * @dev Resolves a signer address to its validator using historical effective bindings.
     * @param signer Signer hot address.
     * @return validator Validator cold address, or zero if signer has never been effective.
     */
    function getValidatorBySignerHistory(address signer) external view returns (address validator);

    /**
     * @dev Gets effective active signer set derived from current validator set.
     * @return signers Active signer addresses.
     */
    function getActiveSigners() external view returns (address[] memory signers);

    /**
     * @dev Gets effective top signer set derived from effective top validators.
     * @return signers Top signer addresses.
     */
    function getTopSigners() external view returns (address[] memory signers);

    /**
     * @dev Gets the signer set that should be committed into the current epoch checkpoint header.
     * @notice Uses checkpoint-transition semantics: a signer scheduled for the current epoch block
     *         is exposed here even though it does not become the runtime-effective signer until the next block.
     * @return signers Top signer addresses for epoch transition/header extra construction.
     */
    function getTopSignersForEpochTransition() external view returns (address[] memory signers);

    /**
     * @dev Tries to add a validator to the highest set based on stake.
     * @param validator Address of the validator to add.
     */
    function tryAddValidatorToHighestSet(address validator) external;

    /**
     * @dev Tries to activate a validator.
     * @param validator Address of the validator to activate.
     * @return bool Returns true if the validator was successfully activated.
     */
    function tryActive(address validator) external returns (bool);

    /**
     * @dev Get count of active validators.
     * @return Count of validators in currentValidatorSet.
     */
    function getActiveValidatorCount() external view returns (uint256);

    /**
     * @dev Get count of voting validators (active and not jailed).
     * @return Count of validators eligible to vote.
     */
    function getVotingValidatorCount() external view returns (uint256);

    /**
     * @dev Get the effective top validators after applying stake-based filtering.
     * @return Validators currently eligible to enter the next epoch set.
     */
    function getEffectiveTopValidators() external view returns (address[] memory);

    /**
     * @dev Get count of effective top validators.
     * @return Count of validators returned by getEffectiveTopValidators().
     */
    function getEffectiveTopValidatorCount() external view returns (uint256);

    /**
     * @dev Check whether a validator is the only effective top validator left.
     * @param validator Address to check.
     * @return bool Returns true when validator is the sole effective top validator.
     */
    function isLastEffectiveValidator(address validator) external view returns (bool);

    /**
     * @dev Get reward-eligible validators and their total stakes.
     * @return validators Active validators that are not jailed.
     * @return totalStakes Total stake amounts for each validator.
     */
    function getRewardEligibleValidatorsWithStakes()
        external
        view
        returns (address[] memory validators, uint256[] memory totalStakes);

    /**
     * @dev Get reward-eligible signers and their corresponding validator total stakes.
     * @return signers Active signer addresses that can receive block rewards.
     * @return totalStakes Total stake amounts for each signer's validator.
     */
    function getRewardEligibleSignersWithStakes()
        external
        view
        returns (address[] memory signers, uint256[] memory totalStakes);

    /**
     * @dev Checks if an address is an active validator.
     * @param who Address to check.
     * @return bool Returns true if the address is an active validator.
     */
    function isActiveValidator(address who) external view returns (bool);

    /**
     * @dev Checks if a validator is active (in current set and not jailed).
     * @param validator Address to check.
     * @return bool Returns true if the validator is active.
     */
    function isValidatorActive(address validator) external view returns (bool);

    /**
     * @dev Checks if an address is a top validator.
     * @param who Address to check.
     * @return bool Returns true if the address is a top validator.
     */
    function isTopValidator(address who) external view returns (bool);

    /**
     * @dev Removes a validator from the highest set.
     * @param validator Address of the validator to remove.
     */
    function removeFromHighestSet(address validator) external;

    /**
     * @dev Removes a validator from the active set.
     * @param val Address of the validator to remove.
     */
    function removeValidator(address val) external;

    /**
     * @dev Removes a validator from the incoming validator set.
     * @param val Address of the validator to remove from the incoming set.
     */
    function removeValidatorIncoming(address val) external;

    /**
     * @dev Tries to remove a validator.
     * @param val Address of the validator to remove.
     */
    function tryRemoveValidator(address val) external;

    /**
     * @dev Checks if validator exists (has staked).
     * @param validator Address to check.
     * @return bool Returns true if validator exists.
     */
    function isValidatorExist(address validator) external view returns (bool);
}
