// SPDX-License-Identifier: MIT
pragma solidity 0.8.29;

import {Script, console} from "lib/forge-std/src/Script.sol";
import {Staking} from "../contracts/Staking.sol";
import {Proposal} from "../contracts/Proposal.sol";

/**
 * @title BatchValidatorStakeScript
 * @dev Batch validator staking script - suitable for situations where multiple validator private keys are controlled
 * Warning: Only use when you own all validator private keys
 */
contract BatchValidatorStakeScript is Script {
    function setUp() public {}

    /**
     * @dev Batch validator staking - requires validator private keys to be pre-set
     */
    function run() external {
        address stakingAddress = vm.envOr("STAKING_CONTRACT", address(0));
        require(stakingAddress != address(0), "STAKING_CONTRACT address not set");

        // Validator private key array - obtained from environment variables or secure storage during actual deployment
        uint256[] memory validatorKeys = getValidatorPrivateKeys();

        address proposalAddress = address(Staking(stakingAddress).proposalContract());
        uint256 stakeAmount = Proposal(proposalAddress).minValidatorStake();
        uint256 commissionRate = 500; // 5%

        console.log("=== Batch Validator Staking ===");
        console.log("Staking contract:", stakingAddress);
        console.log("Proposal contract:", proposalAddress);
        console.log("Number of validators:", validatorKeys.length);
        console.log("Stake per validator:", stakeAmount / 1 ether, "JU");

        for (uint256 i = 0; i < validatorKeys.length; i++) {
            registerSingleValidator(stakingAddress, validatorKeys[i], stakeAmount, commissionRate, i + 1);
        }

        // Final status check
        checkFinalStatus(stakingAddress);
    }

    /**
     * @dev Register a single validator
     */
    function registerSingleValidator(
        address stakingAddress,
        uint256 privateKey,
        uint256 stakeAmount,
        uint256 commissionRate,
        uint256 index
    ) internal {
        address validator = vm.addr(privateKey);

        console.log("Registering validator", index, ":", validator);

        // Check balance (if balance is insufficient, funds need to be transferred first
        if (validator.balance < stakeAmount) {
            console.log("Warning: Validator", validator, "has insufficient balance");
            console.log("Required:", stakeAmount / 1 ether, "JU");
            console.log("Current:", validator.balance / 1 ether, "JU");
            return; // Skip this validator
        }

        vm.startBroadcast(privateKey);

        Staking(stakingAddress).registerValidator{value: stakeAmount}(commissionRate);

        vm.stopBroadcast();

        console.log("Validator", index, "registered successfully");
    }

    /**
     * @dev Get validator private keys - secure key management is needed for actual deployment
     */
    function getValidatorPrivateKeys() internal view returns (uint256[] memory) {
        uint256[] memory keys = new uint256[](5);

        // Warning: These are sample private keys, real secure keys must be used for actual deployment
        // Should be obtained from environment variables, key management services, or hardware wallets
        if (block.chainid == 202599) {
            // Testnet sample private keys - for demonstration only
            keys[0] = vm.envOr("VALIDATOR_1_KEY", uint256(0));
            keys[1] = vm.envOr("VALIDATOR_2_KEY", uint256(0));
            keys[2] = vm.envOr("VALIDATOR_3_KEY", uint256(0));
            keys[3] = vm.envOr("VALIDATOR_4_KEY", uint256(0));
            keys[4] = vm.envOr("VALIDATOR_5_KEY", uint256(0));
        }

        // Verify all private keys are set
        for (uint256 i = 0; i < keys.length; i++) {
            require(keys[i] != 0, string(abi.encodePacked("VALIDATOR_", i + 1, "_KEY not set")));
        }

        return keys;
    }

    /**
     * @dev Check final status
     */
    function checkFinalStatus(address stakingAddress) internal view {
        Staking staking = Staking(stakingAddress);

        uint256 totalStaked = staking.totalStaked();
        // Use getValidatorCount() to get the count of all registered validators
        uint256 validatorCount = staking.getValidatorCount();

        console.log("=== Final Status ===");
        console.log("Total staked:", totalStaked / 1 ether, "JU");
        console.log("Total validators:", validatorCount);
        if (validatorCount > 0) {
            console.log("Average stake per validator:", (totalStaked / validatorCount) / 1 ether, "JU");
        } else {
            console.log("No validators registered yet");
        }
    }
}
