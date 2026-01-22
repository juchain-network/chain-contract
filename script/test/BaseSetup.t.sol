
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {Test} from "lib/forge-std/src/Test.sol";
import {Validators} from "../../contracts/Validators.sol";
import {Punish} from "../../contracts/Punish.sol";
import {Proposal} from "../../contracts/Proposal.sol";
import {Staking} from "../../contracts/Staking.sol";

abstract contract BaseSetup is Test {
    address constant VALIDATORS = 0x000000000000000000000000000000000000F010;
    address constant PUNISH = 0x000000000000000000000000000000000000F011;
    address constant PROPOSAL = 0x000000000000000000000000000000000000F012;
    address constant STAKING = 0x000000000000000000000000000000000000F013;
    uint256 constant TEST_EPOCH = 1_000_000;

    // Deploy runtime code of contracts to fixed addresses and initialize them
    function deploySystem(address[] memory initVals) internal {
        // place runtime code at fixed addresses (consistent with deployment)
        vm.etch(VALIDATORS, type(Validators).runtimeCode);
        vm.etch(PUNISH, type(Punish).runtimeCode);
        vm.etch(PROPOSAL, type(Proposal).runtimeCode);
        vm.etch(STAKING, type(Staking).runtimeCode);

        // Initialize contracts in correct order
        // 1. Proposal first (needed by others)
        Proposal(PROPOSAL).initialize(initVals, VALIDATORS, TEST_EPOCH);
        
        // 2. Staking with genesis validators (but don't call tryAddValidatorToHighestSet yet)
        // Use initializeWithValidators to automatically register genesis validators in Staking
        // This ensures genesis validators are immediately available without needing to register separately
        Staking(STAKING).initializeWithValidators(VALIDATORS, PROPOSAL, initVals, 1000); // 10% commission
        
        // 3. Punish (needs Staking)
        Punish(PUNISH).initialize(VALIDATORS, PROPOSAL, STAKING);
        
        // 4. Validators last (needs all others)
        Validators(VALIDATORS).initialize(initVals, PROPOSAL, PUNISH, STAKING);
        
        // 5. Now that Validators is initialized, add genesis validators to highestValidatorsSet
        // This completes the registration process that was started in initializeWithValidators
        // Must call from STAKING address to satisfy onlyStakingContract modifier
        for (uint256 i = 0; i < initVals.length; i++) {
            vm.prank(STAKING);
            Validators(VALIDATORS).tryAddValidatorToHighestSet(initVals[i]);
        }
    }
}
