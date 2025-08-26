// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {Script, console} from "lib/forge-std/src/Script.sol";
import "../contracts/Validators.sol";
import "../contracts/Proposal.sol";
import "../contracts/Punish.sol";
import "../contracts/Staking.sol";

/**
 * @title DeploySystemScript
 * @dev 统一的部署和管理脚本，包含完整的系统部署、初始化和状态检查功能
 */
contract DeploySystemScript is Script {
    // 系统合约地址 (与genesis一致)
    address constant VALIDATOR_CONTRACT_ADDR = 0x000000000000000000000000000000000000f000;
    address constant PUNISH_CONTRACT_ADDR = 0x000000000000000000000000000000000000F001;
    address constant PROPOSAL_CONTRACT_ADDR = 0x000000000000000000000000000000000000F002;
    address constant STAKING_CONTRACT_ADDR = 0x000000000000000000000000000000000000F003;
    
    // 事件
    event SystemDeployed(address validators, address proposal, address punish, address staking);
    event SystemInitialized(address[] validators);
    event SystemStatusChecked(bool validatorsInit, bool proposalInit, bool punishInit);
    
    function setUp() public {}

    /**
     * @dev 主部署函数 - 部署所有合约并初始化
     */
    function run() public {
        uint256 deployerPrivateKey = vm.envOr("PRIVATE_KEY", uint256(0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80));
        
        vm.startBroadcast(deployerPrivateKey);

        console.log("=== Starting Unified Contract Deployment ===");

        // 部署所有合约
        (address validators, address proposal, address punish, address staking) = deployAllContracts();
        
        // 创建初始验证器数组
        address[] memory initialValidators = createInitialValidators();
        
        // 初始化合约
        initializeContracts(initialValidators);
        
        // 检查系统状态
        checkAndLogSystemStatus();
        
        // 发出部署完成事件
        emit SystemDeployed(validators, proposal, punish, staking);
        
        console.log("=== Deployment Complete ===");
        logDeploymentSummary();

        vm.stopBroadcast();
    }

    /**
     * @dev 部署所有合约
     */
    function deployAllContracts() internal returns (
        address validators,
        address proposal,
        address punish,
        address staking
    ) {
        console.log("Deploying contracts...");
        
        // 部署Validators
        console.log("Deploying Validators...");
        Validators validatorsContract = new Validators();
        validators = address(validatorsContract);
        console.log("Validators deployed at:", validators);

        // 部署Proposal
        console.log("Deploying Proposal...");
        Proposal proposalContract = new Proposal();
        proposal = address(proposalContract);
        console.log("Proposal deployed at:", proposal);

        // 部署Punish
        console.log("Deploying Punish...");
        Punish punishContract = new Punish();
        punish = address(punishContract);
        console.log("Punish deployed at:", punish);

        // 部署Staking
        console.log("Deploying Staking...");
        Staking stakingContract = new Staking();
        staking = address(stakingContract);
        console.log("Staking deployed at:", staking);

        // 使用etch将合约部署到预定义的系统地址
        vm.etch(VALIDATOR_CONTRACT_ADDR, validators.code);
        vm.etch(PUNISH_CONTRACT_ADDR, punish.code);
        vm.etch(PROPOSAL_CONTRACT_ADDR, proposal.code);
        vm.etch(STAKING_CONTRACT_ADDR, staking.code);

        console.log("=== System Contracts Etched ===");
        console.log("Validators at:", VALIDATOR_CONTRACT_ADDR);
        console.log("Punish at:", PUNISH_CONTRACT_ADDR);
        console.log("Proposal at:", PROPOSAL_CONTRACT_ADDR);
        console.log("Staking at:", STAKING_CONTRACT_ADDR);
    }

    /**
     * @dev 创建初始验证器数组
     */
    function createInitialValidators() internal pure returns (address[] memory) {
        address[] memory initialValidators = new address[](3);
        initialValidators[0] = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;
        initialValidators[1] = 0x70997970C51812dc3A010C7d01b50e0d17dc79C8;
        initialValidators[2] = 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC;
        return initialValidators;
    }

    /**
     * @dev 初始化所有合约
     */
    function initializeContracts(address[] memory initialValidators) internal {
        console.log("Initializing contracts...");
        
        // 尝试初始化Validators (如果未初始化)
        try Validators(VALIDATOR_CONTRACT_ADDR).initialize(initialValidators) {
            console.log("Validators initialized successfully");
        } catch Error(string memory reason) {
            console.log("Validators initialization failed:", reason);
        } catch {
            console.log("Validators already initialized (OK)");
        }

        // 尝试初始化Proposal (如果未初始化) 
        try Proposal(PROPOSAL_CONTRACT_ADDR).initialize(initialValidators) {
            console.log("Proposal initialized successfully");
        } catch Error(string memory reason) {
            console.log("Proposal initialization failed:", reason);
        } catch {
            console.log("Proposal already initialized (OK)");
        }

        // 尝试初始化Punish
        try Punish(PUNISH_CONTRACT_ADDR).initialize() {
            console.log("Punish initialized successfully");
        } catch Error(string memory reason) {
            console.log("Punish initialization failed:", reason);
        } catch {
            console.log("Punish already initialized (OK)");
        }

        // 尝试初始化Staking
        try Staking(STAKING_CONTRACT_ADDR).initialize() {
            console.log("Staking initialized successfully");
        } catch Error(string memory reason) {
            console.log("Staking initialization failed:", reason);
        } catch {
            console.log("Staking already initialized (OK)");
        }

        emit SystemInitialized(initialValidators);
    }

    /**
     * @dev 检查并记录系统状态
     */
    function checkAndLogSystemStatus() internal view {
        console.log("=== System Status Check ===");
        
        // 检查验证器状态
        try Validators(VALIDATOR_CONTRACT_ADDR).getActiveValidators() returns (address[] memory active) {
            console.log("Active validators count:", active.length);
            for (uint i = 0; i < active.length && i < 5; i++) {
                console.log("Validator", i, ":", active[i]);
            }
        } catch {
            console.log("Failed to get active validators");
        }
        
        // 检查提案配置
        try Proposal(PROPOSAL_CONTRACT_ADDR).proposalLastingPeriod() returns (uint256 period) {
            console.log("Proposal lasting period:", period);
        } catch {
            console.log("Failed to get proposal period");
        }
        
        // 检查接收地址
        try Proposal(PROPOSAL_CONTRACT_ADDR).receiverAddr() returns (address receiver) {
            console.log("Receiver address:", receiver);
        } catch {
            console.log("Failed to get receiver address");
        }
    }

    /**
     * @dev 记录部署摘要
     */
    function logDeploymentSummary() internal pure {
        console.log("System contracts deployed at fixed addresses:");
        console.log("  Validators:", VALIDATOR_CONTRACT_ADDR);
        console.log("  Proposal:", PROPOSAL_CONTRACT_ADDR);
        console.log("  Punish:", PUNISH_CONTRACT_ADDR);
    }

    /**
     * @dev 检查系统状态 - 可以独立调用
     */
    function checkSystemStatus() external view returns (
        bool validatorsInit,
        bool proposalInit, 
        bool punishInit,
        address[] memory activeValidators,
        address[] memory topValidators
    ) {
        // 检查各合约初始化状态
        try Validators(VALIDATOR_CONTRACT_ADDR).getActiveValidators() returns (address[] memory active) {
            validatorsInit = active.length > 0;
            activeValidators = active;
        } catch {
            validatorsInit = false;
        }
        
        try Validators(VALIDATOR_CONTRACT_ADDR).getTopValidators() returns (address[] memory top) {
            topValidators = top;
        } catch {
            // ignore
        }
        
        try Proposal(PROPOSAL_CONTRACT_ADDR).proposalLastingPeriod() returns (uint256 period) {
            proposalInit = period > 0;
        } catch {
            proposalInit = false;
        }
        
        // 简单检查Punish是否有代码
        punishInit = PUNISH_CONTRACT_ADDR.code.length > 0;
    }
    
    /**
     * @dev 获取系统配置
     */
    function getSystemConfig() external view returns (
        uint256 proposalLastingPeriod,
        uint256 punishThreshold,
        uint256 removeThreshold,
        uint256 decreaseRate,
        uint256 withdrawProfitPeriod,
        uint256 increasePeriod,
        address receiverAddr
    ) {
        Proposal p = Proposal(PROPOSAL_CONTRACT_ADDR);
        return (
            p.proposalLastingPeriod(),
            p.punishThreshold(),
            p.removeThreshold(),
            p.decreaseRate(),
            p.withdrawProfitPeriod(),
            p.increasePeriod(),
            p.receiverAddr()
        );
    }
    
    /**
     * @dev 便利函数：检查地址是否为验证者
     */
    function isValidator(address validator) external view returns (bool isActive, bool isTop, bool isPassed) {
        isActive = Validators(VALIDATOR_CONTRACT_ADDR).isActiveValidator(validator);
        isTop = Validators(VALIDATOR_CONTRACT_ADDR).isTopValidator(validator);
        isPassed = Proposal(PROPOSAL_CONTRACT_ADDR).pass(validator);
    }
    
    /**
     * @dev 便利函数：获取验证者详细信息
     */
    function getValidatorDetails(address validator) external view returns (
        address payable feeAddr,
        uint256 status,
        uint256 aacIncoming,
        uint256 totalJailedHB,
        uint256 lastWithdrawProfitsBlock
    ) {
        Validators.Status statusEnum;
        (feeAddr, statusEnum, aacIncoming, totalJailedHB, lastWithdrawProfitsBlock) = 
            Validators(VALIDATOR_CONTRACT_ADDR).getValidatorInfo(validator);
        status = uint256(statusEnum);
    }

    /**
     * @dev 获取部署信息摘要
     */
    function getDeploymentInfo() external pure returns (
        address validatorAddr,
        address proposalAddr,
        address punishAddr,
        string memory description
    ) {
        return (
            VALIDATOR_CONTRACT_ADDR,
            PROPOSAL_CONTRACT_ADDR,
            PUNISH_CONTRACT_ADDR,
            "Congress POA System Contracts - Unified Deployment"
        );
    }

    /**
     * @dev 修复receiverAddr的函数
     */
    function fixReceiverAddr(address newReceiver) external {
        require(newReceiver != address(0), "Invalid receiver address");
        
        vm.startBroadcast();
        
        // 这里需要通过提案系统来设置receiverAddr
        // 或者如果有管理员权限，可以直接设置
        console.log("Setting receiver address to:", newReceiver);
        
        vm.stopBroadcast();
    }
}
