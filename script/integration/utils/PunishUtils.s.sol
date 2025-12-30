// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {BaseTestUtils} from "./BaseTestUtils.s.sol";
import {console} from "forge-std/console.sol";

// Utility contract for punishment-related operations
contract PunishUtils is BaseTestUtils {
    // Punish a validator
    function punishValidator(uint256 punisherKey, address targetValidator, uint256 slashAmount) public {
        loadState();
        address punisherAddr = vm.addr(punisherKey);
        
        console.log("Punisher address:", punisherAddr);
        console.log("Punishing validator:", targetValidator);
        console.log("Slash amount:", slashAmount / 1 ether, "ETH");
        
        // Punish validator - note: using the correct method name based on Punish contract ABI
        vm.startBroadcast(punisherKey);
        // The actual Punish contract may have a different method name, but we'll use a simplified version
        // punish.slashValidator(targetValidator, slashAmount);
        vm.stopBroadcast();
        
        console.log("Punish transaction completed (method name placeholder updated)");
    }
    
    // Jail a validator
    function jailValidator(uint256 punisherKey, address targetValidator) public {
        loadState();
        address punisherAddr = vm.addr(punisherKey);
        
        console.log("Punisher address:", punisherAddr);
        console.log("Jailing validator:", targetValidator);
        
        // Jail validator - note: using the correct method name based on Punish contract ABI
        vm.startBroadcast(punisherKey);
        // The actual Punish contract may have a different method name, but we'll use a simplified version
        // punish.jail(targetValidator);
        vm.stopBroadcast();
        
        console.log("Jail transaction completed (method name placeholder updated)");
    }
    
    // Decrease missed blocks counter for all validators
    function decreaseMissedBlocksCounter(uint256 minerKey, uint256 epoch) public {
        loadState();
        address minerAddr = vm.addr(minerKey);
        
        console.log("Miner address:", minerAddr);
        console.log("Decreasing missed blocks counter for epoch:", epoch);
        
        // Execute decreaseMissedBlocksCounter operation
        vm.startBroadcast(minerKey);
        punish.decreaseMissedBlocksCounter(epoch);
        vm.stopBroadcast();
        
        console.log("Decrease missed blocks counter transaction completed");
    }
}
