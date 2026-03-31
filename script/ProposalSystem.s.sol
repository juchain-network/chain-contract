// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {Script} from "forge-std/Script.sol";
import {Test, console} from "forge-std/Test.sol";
import {Proposal} from "../contracts/Proposal.sol";
import {Punish} from "../contracts/Punish.sol";
import {Staking} from "../contracts/Staking.sol";
import {Validators} from "../contracts/Validators.sol";

contract ProposalSystemScript is Script, Test {
    // Configuration
    uint256 public constant PROPOSAL_LASTING_PERIOD = 7 days;
    uint256 public constant INITIAL_VALIDATORS = 3;

    // Contracts
    Proposal public proposal;
    Punish public punish;
    Staking public staking;
    Validators public validators;

    // Test accounts
    address[] public validatorAccounts;

    // Deployment keys
    uint256 deployerKey = vm.envUint("DEPLOYER_KEY");

    function run() public {
        console.log("Starting Proposal System Tests...");

        // Use default key if not provided
        if (deployerKey == 0) {
            deployerKey = 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80;
        }

        // Create test accounts
        createTestAccounts();

        // Deploy and initialize contracts
        deployAndInitializeContracts();

        // Test 1: Config Proposal Creation
        testConfigProposalCreation();

        // Test 2: Validator Proposal Creation
        testValidatorProposalCreation();

        console.log("\nAll Proposal System tests completed successfully!");
    }

    function createTestAccounts() internal {
        // Create initial validator accounts
        for (uint256 i = 0; i < INITIAL_VALIDATORS; i++) {
            address validator = vm.addr(uint256(keccak256(abi.encodePacked("proposalValidator", i))));
            validatorAccounts.push(validator);
            vm.deal(validator, 1000000 ether);
        }

        console.log(unicode"✓ Test validator accounts created");
        for (uint256 i = 0; i < validatorAccounts.length; i++) {
            console.log("Validator", i + 1, ":", validatorAccounts[i]);
        }
    }

    function deployAndInitializeContracts() internal {
        vm.startBroadcast(deployerKey);

        // Deploy contracts
        proposal = new Proposal();
        punish = new Punish();
        staking = new Staking();
        validators = new Validators();

        // Initialize Proposal contract
        address[] memory initialValidators = new address[](validatorAccounts.length);
        for (uint256 i = 0; i < validatorAccounts.length; i++) {
            initialValidators[i] = validatorAccounts[i];
        }
        uint256 epoch = vm.envOr("EPOCH_DURATION", uint256(600));
        proposal.initialize(initialValidators, address(validators), epoch);

        // Initialize Validators contract
        validators.initialize(
            initialValidators, initialValidators, address(proposal), address(punish), address(staking)
        );

        // Initialize Punish contract
        punish.initialize(address(validators), address(proposal), address(staking));

        // Initialize Staking contract with initial validators
        staking.initializeWithValidators(
            address(validators),
            address(proposal),
            address(punish),
            initialValidators,
            1000 // 10% commission
        );

        // Contract addresses are already set during initialization
        // Miner address is set by the broadcast key, no need to call setMiner on contracts

        vm.stopBroadcast();

        console.log(unicode"✓ Contracts deployed and initialized");
    }

    function testConfigProposalCreation() internal {
        console.log("\n=== Testing Config Proposal Creation ===");

        address proposer = validatorAccounts[0];
        uint256 configId = 5; // blockReward
        uint256 newValue = 0.3 ether;

        // Create config proposal
        vm.startBroadcast(proposer);
        bytes32 proposalId = proposal.createUpdateConfigProposal(configId, newValue);
        vm.stopBroadcast();

        require(proposalId != bytes32(0), "Proposal creation should succeed");

        console.log(unicode"✓ Config proposal created successfully");
    }

    function testValidatorProposalCreation() internal {
        console.log("\n=== Testing Validator Proposal Creation ===");

        address proposer = validatorAccounts[0];
        address newValidator = vm.addr(uint256(keccak256(abi.encodePacked("newValidatorForProposal"))));
        vm.deal(newValidator, 1000000 ether);

        // Create validator proposal
        vm.startBroadcast(proposer);
        bytes32 proposalId = proposal.createProposal(newValidator, true, "Add new validator via proposal");
        vm.stopBroadcast();

        require(proposalId != bytes32(0), "Validator proposal creation should succeed");

        console.log(unicode"✓ Validator proposal created successfully for:", newValidator);
    }
}
