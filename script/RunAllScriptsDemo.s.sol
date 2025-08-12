// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {Script, console} from "lib/forge-std/src/Script.sol";
import "../contracts/Validators.sol";
import "../contracts/Proposal.sol";
import "../contracts/Punish.sol";

/**
 * @title RunAllScriptsDemo
 * @dev 运行所有脚本功能的演示脚本
 */
contract RunAllScriptsDemo is Script {
    // 系统合约地址
    address constant VALIDATOR_CONTRACT_ADDR = 0x000000000000000000000000000000000000f000;
    address constant PROPOSAL_CONTRACT_ADDR = 0x000000000000000000000000000000000000F002;
    address constant PUNISH_CONTRACT_ADDR = 0x000000000000000000000000000000000000F001;
    
    // 测试账户
    address constant PROPOSER = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;  // Account 0
    address constant NEW_VALIDATOR = 0x90F79bf6EB2c4f870365E785982E1f101E93b906; // Account 3
    
    function setUp() public {}

    function run() public {
        uint256 deployerPrivateKey = vm.envOr("PRIVATE_KEY", uint256(0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80));
        
        vm.startBroadcast(deployerPrivateKey);

        console.log("=== Running All Script Functions Demo ===");

        // 1. 检查系统状态（DeploySystem功能）
        console.log("\n1. === System Status Check (DeploySystem) ===");
        checkSystemStatus();
        
        // 2. 创建提案（CreateProposal功能）
        console.log("\n2. === Creating Proposal (CreateProposal) ===");
        bytes32 proposalId = createAddValidatorProposal();
        
        // 3. 投票（VoteProposal功能）
        console.log("\n3. === Voting on Proposal (VoteProposal) ===");
        voteOnProposal(proposalId);
        
        // 4. 检查提案结果
        console.log("\n4. === Checking Proposal Results ===");
        checkProposalResults(proposalId);
        
        // 5. 尝试配置更新（UpdateConfig功能）
        console.log("\n5. === Testing Config Update (UpdateConfig) ===");
        testConfigUpdate();

        vm.stopBroadcast();
        
        console.log("\n=== All Script Functions Demo Complete ===");
    }
    
    /**
     * @dev 检查系统状态（DeploySystem脚本功能）
     */
    function checkSystemStatus() internal view {
        // 获取活跃验证器
        address[] memory activeValidators = Validators(VALIDATOR_CONTRACT_ADDR).getActiveValidators();
        console.log("Active validators count:", activeValidators.length);
        
        for (uint i = 0; i < activeValidators.length; i++) {
            console.log("Active validator", i, ":", activeValidators[i]);
        }
        
        // 检查提案配置
        try Proposal(PROPOSAL_CONTRACT_ADDR).proposalLastingPeriod() returns (uint256 period) {
            console.log("Proposal lasting period:", period);
        } catch {
            console.log("Could not get proposal lasting period");
        }
        
        try Proposal(PROPOSAL_CONTRACT_ADDR).receiverAddr() returns (address receiver) {
            console.log("Receiver address:", receiver);
        } catch {
            console.log("Could not get receiver address");
        }
    }
    
    /**
     * @dev 创建添加验证器提案（CreateProposal脚本功能）
     */
    function createAddValidatorProposal() internal returns (bytes32 proposalId) {
        string memory details = "Adding new validator node for network decentralization";
        
        console.log("Creating proposal to add validator:", NEW_VALIDATOR);
        console.log("Details:", details);
        
        try Proposal(PROPOSAL_CONTRACT_ADDR).createProposal(
            NEW_VALIDATOR,
            true,  // flag: true = add, false = remove
            details
        ) returns (bool success) {
            if (success) {
                // 计算提案ID
                proposalId = keccak256(abi.encodePacked(
                    msg.sender,
                    NEW_VALIDATOR,
                    true,
                    details,
                    block.timestamp
                ));
                
                console.log("Proposal created successfully");
                console.log("Proposal ID:");
                console.logBytes32(proposalId);
            } else {
                console.log("Failed to create proposal");
            }
        } catch Error(string memory reason) {
            console.log("Create proposal failed:", reason);
        } catch {
            console.log("Create proposal failed with unknown error");
        }
        
        return proposalId;
    }
    
    /**
     * @dev 对提案投票（VoteProposal脚本功能）
     */
    function voteOnProposal(bytes32 proposalId) internal {
        if (proposalId == bytes32(0)) {
            console.log("No valid proposal ID to vote on");
            return;
        }
        
        console.log("Voting on proposal:");
        console.logBytes32(proposalId);
        
        try Proposal(PROPOSAL_CONTRACT_ADDR).voteProposal(proposalId, true) returns (bool success) {
            if (success) {
                console.log("Vote cast successfully: YES");
            } else {
                console.log("Vote failed");
            }
        } catch Error(string memory reason) {
            console.log("Vote failed:", reason);
        } catch {
            console.log("Vote failed with unknown error");
        }
    }
    
    /**
     * @dev 检查提案结果
     */
    function checkProposalResults(bytes32 proposalId) internal view {
        if (proposalId == bytes32(0)) {
            console.log("No valid proposal ID to check");
            return;
        }
        
        console.log("Checking proposal:");
        console.logBytes32(proposalId);
        
        // 检查提案详情
        try Proposal(PROPOSAL_CONTRACT_ADDR).proposals(proposalId) returns (
            address proposer,
            uint256 createTime,
            uint256 proposalType,
            address dst,
            bool flag,
            string memory /* details */,
            uint256 /* cid */,
            uint256 /* newValue */
        ) {
            console.log("Proposal found:");
            console.log("  Proposer:", proposer);
            console.log("  Target (dst):", dst);
            console.log("  Flag (add/remove):", flag);
            console.log("  Create time:", createTime);
            console.log("  Proposal type:", proposalType);
        } catch {
            console.log("Could not retrieve proposal details");
        }
        
        // 检查是否通过
        try Proposal(PROPOSAL_CONTRACT_ADDR).pass(NEW_VALIDATOR) returns (bool passed) {
            console.log("Target validator passed:", passed);
        } catch {
            console.log("Could not check pass status");
        }
    }
    
    /**
     * @dev 测试配置更新（UpdateConfig脚本功能）
     */
    function testConfigUpdate() internal {
        console.log("Testing configuration update...");
        
        // 尝试创建配置更新提案
        try Proposal(PROPOSAL_CONTRACT_ADDR).createUpdateConfigProposal(1, 100000) returns (bool success) {
            if (success) {
                console.log("Config update proposal created successfully");
                console.log("Config ID: 1, New Value: 100000");
            } else {
                console.log("Config update proposal failed");
            }
        } catch Error(string memory reason) {
            console.log("Config update failed:", reason);
        } catch {
            console.log("Config update failed with unknown error");
        }
    }
    
    /**
     * @dev AddNewNode脚本功能演示
     */
    function demonstrateAddNewNode(address newNode) external {
        console.log("=== Add New Node Demo ===");
        console.log("Adding node:", newNode);
        
        try Proposal(PROPOSAL_CONTRACT_ADDR).createProposal(
            newNode,
            true,
            "Adding new node via AddNewNode script"
        ) returns (bool success) {
            console.log("Add node proposal created:", success);
        } catch Error(string memory reason) {
            console.log("Add node failed:", reason);
        }
    }
    
    /**
     * @dev RemoveNode脚本功能演示
     */
    function demonstrateRemoveNode(address nodeToRemove) external {
        console.log("=== Remove Node Demo ===");
        console.log("Removing node:", nodeToRemove);
        
        try Proposal(PROPOSAL_CONTRACT_ADDR).createProposal(
            nodeToRemove,
            false,
            "Removing node via RemoveNode script"
        ) returns (bool success) {
            console.log("Remove node proposal created:", success);
        } catch Error(string memory reason) {
            console.log("Remove node failed:", reason);
        }
    }
    
    /**
     * @dev 获取系统概览
     */
    function getSystemOverview() external view returns (
        uint256 activeValidatorCount,
        uint256 topValidatorCount,
        address proposalContract,
        address validatorContract,
        address punishContract
    ) {
        address[] memory active = Validators(VALIDATOR_CONTRACT_ADDR).getActiveValidators();
        address[] memory top = Validators(VALIDATOR_CONTRACT_ADDR).getTopValidators();
        
        return (
            active.length,
            top.length,
            PROPOSAL_CONTRACT_ADDR,
            VALIDATOR_CONTRACT_ADDR,
            PUNISH_CONTRACT_ADDR
        );
    }
}
