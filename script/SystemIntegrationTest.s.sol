// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {Script, console} from "lib/forge-std/src/Script.sol";
import "../contracts/Validators.sol";
import "../contracts/Proposal.sol";
import "../contracts/Punish.sol";
import "../contracts/Staking.sol";

/**
 * @title SystemIntegrationTest
 * @dev 系统集成测试脚本，验证所有合约功能
 */
contract SystemIntegrationTest is Script {
    // 系统合约地址
    address constant VALIDATOR_CONTRACT_ADDR = 0x000000000000000000000000000000000000f000;
    address constant PROPOSAL_CONTRACT_ADDR = 0x000000000000000000000000000000000000F002;
    address constant PUNISH_CONTRACT_ADDR = 0x000000000000000000000000000000000000F001;
    address constant STAKING_CONTRACT_ADDR = 0x000000000000000000000000000000000000F003;
    
    // 测试账户
    address constant PROPOSER = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;  // Account 0
    address constant NEW_VALIDATOR = 0xa0Ee7A142d267C1f36714E4a8F75612F20a79720; // 使用一个未存在的地址
    
    function setUp() public {}

    function run() public {
        uint256 deployerPrivateKey = vm.envOr("PRIVATE_KEY", uint256(0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80));
        
        vm.startBroadcast(deployerPrivateKey);

        console.log("=== System Integration Test ===");

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
        
        // 5. 测试Staking系统（StakingOperations功能）
        console.log("\n5. === Testing Staking System (StakingOperations) ===");
        testStakingSystem();
        
        // 6. 尝试配置更新（UpdateConfig功能）
        console.log("\n6. === Testing Config Update (UpdateConfig) ===");
        testConfigUpdate();

        vm.stopBroadcast();
        
        console.log("\n=== System Integration Test Complete ===");
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
     * @dev 测试Staking系统（StakingOperations脚本功能）
     */
    function testStakingSystem() internal {
        console.log("Testing Staking system...");
        
        Staking stakingContract = Staking(STAKING_CONTRACT_ADDR);
        Validators validatorsContract = Validators(VALIDATOR_CONTRACT_ADDR);
        
        // 1. 检查当前验证者状态
        console.log("--- Current Validator Status ---");
        address[] memory currentValidators = validatorsContract.getActiveValidators();
        console.log("Current validators count:", currentValidators.length);
        
        // 2. 安全地检查Staking状态
        try stakingContract.MIN_VALIDATORS() returns (uint256 minValidators) {
            console.log("Minimum validators required:", minValidators);
        } catch {
            console.log("Cannot get minimum validators");
        }
        
        try stakingContract.MIN_VALIDATOR_STAKE() returns (uint256 minStake) {
            console.log("Minimum validator stake:", minStake);
        } catch {
            console.log("Cannot get minimum stake");
        }
        
        // 3. 检查顶级验证者
        try stakingContract.getTopValidators(3) returns (address[] memory topValidators) {
            console.log("Top validators count:", topValidators.length);
            
            for (uint i = 0; i < topValidators.length && i < 3; i++) {
                try stakingContract.getValidatorInfo(topValidators[i]) returns (
                    uint256 selfStake,
                    uint256 totalDelegated,
                    uint256, // commissionRate
                    bool, // isJailed
                    uint256 // jailUntilBlock
                ) {
                    console.log("Validator", i, ":", topValidators[i]);
                    console.log("  Self stake:", selfStake);
                    console.log("  Total delegated:", totalDelegated);
                } catch {
                    console.log("Validator", i, ": info unavailable");
                }
            }
        } catch {
            console.log("Cannot get top validators");
        }
        
        // 4. 展示委托信息示例
        if (currentValidators.length > 0) {
            address validator = currentValidators[0];
            try stakingContract.getDelegationInfo(msg.sender, validator) returns (
                uint256 delegatedAmount,
                uint256 rewards,
                uint256 unbondingAmount,
                uint256 unbondingBlock
            ) {
                console.log("Delegation info for", validator);
                console.log("  Delegated amount:", delegatedAmount);
                console.log("  Pending rewards:", rewards);
                console.log("  Unbonding amount:", unbondingAmount);
                console.log("  Unbonding block:", unbondingBlock);
            } catch {
                console.log("Cannot get delegation info for", validator);
            }
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
     * @dev AddNewNode脚本功能测试
     */
    function testAddNewNode(address newNode) external {
        console.log("=== Add New Node Test ===");
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
     * @dev RemoveNode脚本功能测试
     */
    function testRemoveNode(address nodeToRemove) external {
        console.log("=== Remove Node Test ===");
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
        address punishContract,
        address stakingContract
    ) {
        address[] memory active = Validators(VALIDATOR_CONTRACT_ADDR).getActiveValidators();
        address[] memory top = Validators(VALIDATOR_CONTRACT_ADDR).getTopValidators();
        
        return (
            active.length,
            top.length,
            PROPOSAL_CONTRACT_ADDR,
            VALIDATOR_CONTRACT_ADDR,
            PUNISH_CONTRACT_ADDR,
            STAKING_CONTRACT_ADDR
        );
    }
}
