// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {Script, console} from "forge-std/Script.sol";
import {Proposal} from "../contracts/Proposal.sol";
import {Punish} from "../contracts/Punish.sol";
import {Staking} from "../contracts/Staking.sol";
import {Validators} from "../contracts/Validators.sol";

contract DeploymentScript is Script {
    // Configuration
    uint256 public constant INITIAL_VALIDATORS = 5;
    uint256 public constant INITIAL_STAKE = 100000 ether;
    uint256 public constant BLOCK_REWARD = 0.2 ether;

    // Deployment keys
    uint256 deployerKey = vm.envUint("DEPLOYER_KEY");

    function run() public {
        console.log("Starting PoSA Contract Deployment...");

        // Use default key if not provided
        if (deployerKey == 0) {
            deployerKey = 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80;
        }

        vm.startBroadcast(deployerKey);

        // 1. Deploy contracts (no constructor parameters needed as they inherit from Params)
        Proposal proposal = new Proposal();
        console.log(unicode"✓ Proposal deployed at:", address(proposal));

        Validators validators = new Validators();
        console.log(unicode"✓ Validators deployed at:", address(validators));

        Punish punish = new Punish();
        console.log(unicode"✓ Punish deployed at:", address(punish));

        Staking staking = new Staking();
        console.log(unicode"✓ Staking deployed at:", address(staking));

        uint256 epoch = vm.envOr("EPOCH_DURATION", uint256(600));

        // 2. Initialize Proposal contract first (requires initial validators and validators contract address)
        address[] memory initialValidators = new address[](1);
        initialValidators[0] = vm.addr(deployerKey);
        proposal.initialize(initialValidators, address(validators), epoch);
        console.log(unicode"✓ Proposal contract initialized");

        // 3. Initialize Validators contract (requires initial validators, proposal, punish, and staking addresses)
        validators.initialize(
            initialValidators, initialValidators, address(proposal), address(punish), address(staking)
        );
        console.log(unicode"✓ Validators contract initialized");

        // 4. Initialize Punish contract (requires validators, proposal, and staking addresses)
        punish.initialize(address(validators), address(proposal), address(staking));
        console.log(unicode"✓ Punish contract initialized");

        // 5. Initialize Staking contract with validators
        staking.initializeWithValidators(
            address(validators),
            address(proposal),
            address(punish),
            initialValidators,
            1000 // 10% commission
        );
        console.log(unicode"✓ Staking contract initialized");

        // Initial configuration is set in the initialize functions

        vm.stopBroadcast();

        console.log("\n=== Deployment Summary ===");
        console.log("Proposal: ", address(proposal));
        console.log("Validators: ", address(validators));
        console.log("Punish: ", address(punish));
        console.log("Staking: ", address(staking));
        console.log("\nDeployment completed successfully!");
    }
}
