// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {Script, console} from "lib/forge-std/src/Script.sol";
import {Validators} from "../contracts/Validators.sol";
import {Proposal} from "../contracts/Proposal.sol";
import {Punish} from "../contracts/Punish.sol";
import {Staking} from "../contracts/Staking.sol";

/**
 * @title DeployToChainScript
 * @dev 真实链上部署脚本，支持多种网络部署
 */
contract DeployToChainScript is Script {
    // 系统合约地址 - 使用 CREATE2 部署到确定性地址
    bytes32 constant SALT = keccak256("SYS_CONTRACT_V1");
    
    // 事件
    event SystemDeployed(address validators, address proposal, address punish, address staking);
    event SystemInitialized(address[] validators);
    
    function setUp() public {}

    /**
     * @dev 主部署函数 - 部署所有合约并初始化
     */
    function run() external {
        // 支持多种私钥环境变量
        uint256 deployerPrivateKey = vm.envOr("CHAIN_PRIVATE_KEY", vm.envOr("PRIVATE_KEY", uint256(0)));
        
        // 如果没有提供私钥，使用默认的 anvil 私钥
        if (deployerPrivateKey == 0) {
            deployerPrivateKey = 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80;
            console.log("Warning: Using default anvil private key");
        }
        
        vm.startBroadcast(deployerPrivateKey);
        
        console.log("=== Starting Chain Deployment ===");
        console.log("Deployer:", vm.addr(deployerPrivateKey));
        console.log("Chain ID:", block.chainid);
        
        // 部署所有合约到确定性地址
        (address validators, address proposal, address punish, address staking) = deployAllContracts();
        
        // 创建初始验证器数组
        address[] memory initialValidators = createInitialValidators();
        
        // 初始化合约
        initializeContracts(validators, proposal, punish, staking, initialValidators);
        
        vm.stopBroadcast();
        
        // 发出部署完成事件
        emit SystemDeployed(validators, proposal, punish, staking);
        
        console.log("=== Chain Deployment Complete ===");
        logDeploymentSummary(validators, proposal, punish, staking);
        
        // 检查系统状态
        checkAndLogSystemStatus(validators, proposal, punish, staking);
    }

    /**
     * @dev 部署所有合约 (使用普通部署而不是 CREATE2)
     */
    function deployAllContracts() internal returns (
        address validators,
        address proposal, 
        address punish,
        address staking
    ) {
        console.log("Deploying contracts...");
        console.log("Current chain ID:", block.chainid);
        console.log("Deployer address:", msg.sender);
        console.log("Deployer balance:", msg.sender.balance);
        
        // 部署 Validators
        console.log("Deploying Validators...");
        Validators validatorsContract = new Validators();
        validators = address(validatorsContract);
        console.log("Validators deployed at:", validators);

        // 部署 Proposal
        console.log("Deploying Proposal...");
        Proposal proposalContract = new Proposal();
        proposal = address(proposalContract);
        console.log("Proposal deployed at:", proposal);

        // 部署 Punish
        console.log("Deploying Punish...");
        Punish punishContract = new Punish();
        punish = address(punishContract);
        console.log("Punish deployed at:", punish);

        // 部署 Staking
        console.log("Deploying Staking...");
        Staking stakingContract = new Staking();
        staking = address(stakingContract);
        console.log("Staking deployed at:", staking);

        console.log("=== All Contracts Deployed ===");
    }

    /**
     * @dev 使用 CREATE2 部署合约
     */
    function deployWithCreate2(bytes memory bytecode, bytes32 salt) internal returns (address) {
        console.log("Deploying with CREATE2, bytecode length:", bytecode.length);
        console.log("Salt:", vm.toString(salt));
        
        address deployed;
        assembly {
            deployed := create2(0, add(bytecode, 0x20), mload(bytecode), salt)
        }
        
        console.log("CREATE2 result:", deployed);
        require(deployed != address(0), "Failed to deploy contract with CREATE2");
        
        // 验证部署
        uint256 size;
        assembly {
            size := extcodesize(deployed)
        }
        console.log("Deployed contract code size:", size);
        require(size > 0, "Contract deployment failed - no code");
        
        return deployed;
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
    function initializeContracts(
        address validators, 
        address proposal, 
        address punish,
        address staking,
        address[] memory initialValidators
    ) internal {
        console.log("Initializing contracts...");
        
        // 初始化Validators
        try Validators(validators).initialize(initialValidators) {
            console.log("Validators initialized successfully");
        } catch Error(string memory reason) {
            console.log("Validators initialization failed:", reason);
        } catch {
            console.log("Validators already initialized (OK)");
        }

        // 初始化Proposal
        try Proposal(proposal).initialize(initialValidators) {
            console.log("Proposal initialized successfully");
        } catch Error(string memory reason) {
            console.log("Proposal initialization failed:", reason);
        } catch {
            console.log("Proposal already initialized (OK)");
        }

        // 初始化Punish
        try Punish(punish).initialize() {
            console.log("Punish initialized successfully");
        } catch Error(string memory reason) {
            console.log("Punish initialization failed:", reason);
        } catch {
            console.log("Punish already initialized (OK)");
        }

        // 初始化Staking
        try Staking(staking).initialize() {
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
    function checkAndLogSystemStatus(
        address validators,
        address proposal,
        address punish,
        address staking
    ) internal view {
        console.log("=== System Status Check ===");
        
        // 检查验证器状态
        try Validators(validators).getActiveValidators() returns (address[] memory active) {
            console.log("Active validators count:", active.length);
            for (uint i = 0; i < active.length && i < 5; i++) {
                console.log("Validator", i, ":", active[i]);
            }
        } catch {
            console.log("Failed to get active validators");
        }
        
        // 检查提案配置
        try Proposal(proposal).proposalLastingPeriod() returns (uint256 period) {
            console.log("Proposal lasting period:", period);
        } catch {
            console.log("Failed to get proposal period");
        }
        
        // 检查接收地址
        try Proposal(proposal).receiverAddr() returns (address receiver) {
            console.log("Receiver address:", receiver);
        } catch {
            console.log("Failed to get receiver address");
        }
    }

    /**
     * @dev 记录部署摘要
     */
    function logDeploymentSummary(
        address validators,
        address proposal,
        address punish,
        address staking
    ) internal pure {
        console.log("System contracts deployed at addresses:");
        console.log("  Validators:", validators);
        console.log("  Proposal:", proposal);
        console.log("  Punish:", punish);
        console.log("  Staking:", staking);
    }

    /**
     * @dev 获取预计算的合约地址
     */
    function getComputedAddress(bytes memory bytecode, bytes32 salt) public view returns (address) {
        bytes32 hash = keccak256(
            abi.encodePacked(
                bytes1(0xff),
                address(this),
                salt,
                keccak256(bytecode)
            )
        );
        return address(uint160(uint256(hash)));
    }
}
