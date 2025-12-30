// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {ProposalUtils} from "./ProposalUtils.s.sol";
import {console} from "forge-std/console.sol";

contract InitEnv is ProposalUtils {
    function run() public override {
        console.log("Init environment...");
        
        // Create test accounts
        createTestAccounts();
        
        // Deploy and initialize contracts
        deployAndInitializeContracts();
        
        saveState();

        configProposal(0, 3600);
        configProposal(6, 3600);
        configProposal(4, 3600);
        configProposal(7, 3600);
        
        console.log("\nAll New Validator tests completed successfully!");
    }
}
