// SPDX-License-Identifier: MIT

pragma solidity ^0.8.20;

interface IPunish {
    /**
     * @dev Initializes the Punish contract with required dependencies.
     * @param _validators Address of the Validators contract.
     * @param _proposal Address of the Proposal contract.
     * @param _staking Address of the Staking contract.
     */
    function initialize(
        address _validators,
        address _proposal,
        address _staking
    ) external;

    /**
     * @dev Punishes a validator for missing blocks.
     * @param val Address of the validator to punish.
     */
    function punish(address val) external;

    /**
     * @dev Decreases the missed blocks counter for validators at the end of an epoch.
     * @param epoch The epoch number for which to decrease the counter.
     */
    function decreaseMissedBlocksCounter(uint256 epoch) external;

    /**
     * @dev Cleans the punishment record for a validator.
     * @param val Address of the validator whose record to clean.
     * @return bool Returns true if the operation was successful.
     */
    function cleanPunishRecord(address val) external returns (bool);

    /**
     * @dev Gets the number of validators currently being punished.
     * @return uint256 The number of validators in the punishment list.
     */
    function getPunishValidatorsLen() external view returns (uint256);

    /**
     * @dev Gets the punishment record for a specific validator.
     * @param val Address of the validator to check.
     * @return uint256 The number of missed blocks for the validator.
     */
    function getPunishRecord(address val) external view returns (uint256);

    /**
     * @dev Gets a validator from the punishment list by index.
     * @param index The index in the punishment list.
     * @return address The address of the validator at the specified index.
     */
    function punishValidators(uint256 index) external view returns (address);
}
