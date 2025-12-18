// SPDX-License-Identifier: MIT

pragma solidity ^0.8.20;

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
     * @dev Checks if an address is an active validator.
     * @param who Address to check.
     * @return bool Returns true if the address is an active validator.
     */
    function isActiveValidator(address who) external view returns (bool);

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
}
