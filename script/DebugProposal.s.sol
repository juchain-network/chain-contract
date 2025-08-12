// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {Script, console} from "lib/forge-std/src/Script.sol";
import {BaseSetup} from "../test/BaseSetup.t.sol";
import "../contracts/Proposal.sol";

/**
 * @title DebugProposalScript
 * @dev 调试Proposal合约的脚本
 */
contract DebugProposalScript is BaseSetup {
    
    function run() public view {
        console.log("=== Debugging Proposal Contract ===");
        
        // 检查基本状态
        console.log("Proposal contract address:", PROPOSAL);
        
        // 检查初始化状态
        try Proposal(PROPOSAL).initialized() returns (bool init) {
            console.log("Initialized:", init);
        } catch {
            console.log("Cannot check initialized status");
        }
        
        // 检查proposalLastingPeriod
        try Proposal(PROPOSAL).proposalLastingPeriod() returns (uint256 period) {
            console.log("Proposal lasting period:", period);
        } catch Error(string memory reason) {
            console.log("proposalLastingPeriod failed:", reason);
        } catch {
            console.log("proposalLastingPeriod failed with unknown error");
        }
        
        // 检查receiverAddr
        try Proposal(PROPOSAL).receiverAddr() returns (address receiver) {
            console.log("Receiver address:", receiver);
        } catch Error(string memory reason) {
            console.log("receiverAddr failed:", reason);
        } catch {
            console.log("receiverAddr failed with unknown error");
        }
        
        // 检查punishThreshold
        try Proposal(PROPOSAL).punishThreshold() returns (uint256 threshold) {
            console.log("Punish threshold:", threshold);
        } catch Error(string memory reason) {
            console.log("punishThreshold failed:", reason);
        } catch {
            console.log("punishThreshold failed with unknown error");
        }
    }
    
    function testCreateProposal() external {
        // 在测试环境中，首先部署系统
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
