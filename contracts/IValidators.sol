// SPDX-License-Identifier: MIT

pragma solidity ^0.8.29;

interface IValidators {
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
