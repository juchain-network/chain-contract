// SPDX-License-Identifier: MIT
pragma solidity 0.8.29;

import {console} from "lib/forge-std/src/console.sol";
import {BaseSetup} from "../test/BaseSetup.t.sol";
import {Proposal} from "../contracts/Proposal.sol";

/**
 * @title DebugProposalScript
 * @dev Script to debug the Proposal contract
 */
contract DebugProposalScript is BaseSetup {
    
    function run() public view {
        console.log("=== Debugging Proposal Contract ===");
        
        // Check basic status
        console.log("Proposal contract address:", PROPOSAL);
        
        // Check initialization status
        try Proposal(PROPOSAL).initialized() returns (bool init) {
            console.log("Initialized:", init);
        } catch {
            console.log("Cannot check initialized status");
        }
        
        // Check proposalLastingPeriod
        try Proposal(PROPOSAL).proposalLastingPeriod() returns (uint256 period) {
            console.log("Proposal lasting period:", period);
        } catch Error(string memory reason) {
            console.log("proposalLastingPeriod failed:", reason);
        } catch {
            console.log("proposalLastingPeriod failed with unknown error");
        }
        
        // Note: receiverAddr and increasePeriod have been removed, token inflation is no longer supported
        
        // Check punishThreshold
        try Proposal(PROPOSAL).punishThreshold() returns (uint256 threshold) {
            console.log("Punish threshold:", threshold);
        } catch Error(string memory reason) {
            console.log("punishThreshold failed:", reason);
        } catch {
            console.log("punishThreshold failed with unknown error");
        }
    }
    
    function testCreateProposal() external {
        // In test environment, first deploy the system
        address[] memory initVals = new address[](3);
        initVals[0] = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;
        initVals[1] = 0x70997970C51812dc3A010C7d01b50e0d17dc79C8;
        initVals[2] = 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC;
        
        deploySystem(initVals);
        
        console.log("=== Testing Create Proposal ===");
        console.log("Caller:", initVals[0]);
        
        address testTarget = 0x90F79bf6EB2c4f870365E785982E1f101E93b906;
        string memory details = "Test proposal";
        
        vm.prank(initVals[0]);
        try Proposal(PROPOSAL).createProposal(testTarget, true, details) returns (bool success) {
            console.log("Create proposal result:", success);
        } catch Error(string memory reason) {
            console.log("Create proposal failed:", reason);
        } catch {
            console.log("Create proposal failed with unknown error");
        }
    }
}
