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
        
        // 设置交易参数 - 处理gas问题
        vm.txGasPrice(1000000000); // 1 gwei
        
        vm.startBroadcast(deployerPrivateKey);
        
        console.log("=== Starting Chain Deployment ===");
        console.log("Deployer:", vm.addr(deployerPrivateKey));
        console.log("Chain ID:", block.chainid);
        console.log("Gas Price:", tx.gasprice);
        
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
     * @dev 创建初始验证器数组 - 根据链ID自动选择验证者地址
     * 测试网验证者地址 (Chain ID: 202599):
     * 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b
     * 0x578c39eAf09a4e1aBF428c423970B59BB8baF42E
     * 0xC9eBc132a89AAb349D9232d8Ce70A2c2FEA0A096
     * 0x9e6A23508aa763C709d45F671D7a3A068025ABC0
     * 0x81f7A79A51eDBA249EfA812Eb2D5478F696f7558
     *
     * 主网验证者地址 (Chain ID: 210000):
     * 0x311B37f01c04B84d1f94645BfBd58D82fc03F709
     * 0xDe0e48c5337db3Ca7b3710c27E9728E68Bf220b3
     * 0xccAFA71c31bC11Ba24d526FD27BA57D743152807
     * 0xD5DA2b33C1f620a94bf2039B9Cb540853e7928D7
     * 0x4D432df142823Ca25b21Bc3F9744ED21A275bDEA
     */
    function createInitialValidators() internal view returns (address[] memory) {
        address[] memory initialValidators = new address[](5);

        if (block.chainid == 202599) {
            // 测试网验证者地址
            console.log("Using testnet validators (Chain ID: 202599)");
            initialValidators[0] = 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b;
            initialValidators[1] = 0x578c39eAf09a4e1aBF428c423970B59BB8baF42E;
            initialValidators[2] = 0xC9eBc132a89AAb349D9232d8Ce70A2c2FEA0A096;
            initialValidators[3] = 0x9e6A23508aa763C709d45F671D7a3A068025ABC0;
            initialValidators[4] = 0x81f7A79A51eDBA249EfA812Eb2D5478F696f7558;
        } else if (block.chainid == 210000) {
            // 主网验证者地址
            console.log("Using mainnet validators (Chain ID: 210000)");
            initialValidators[0] = 0x311B37f01c04B84d1f94645BfBd58D82fc03F709;
            initialValidators[1] = 0xDe0e48c5337db3Ca7b3710c27E9728E68Bf220b3;
            initialValidators[2] = 0xccAFA71c31bC11Ba24d526FD27BA57D743152807;
            initialValidators[3] = 0xD5DA2b33C1f620a94bf2039B9Cb540853e7928D7;
            initialValidators[4] = 0x4D432df142823Ca25b21Bc3F9744ED21A275bDEA;
        } else {
            // 默认使用本地开发环境验证者（anvil/hardhat 默认账户）
            console.log("Using default local validators (Chain ID:", block.chainid, ")");
            initialValidators[0] = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266; // VALIDATOR1
            initialValidators[1] = 0x70997970C51812dc3A010C7d01b50e0d17dc79C8; // VALIDATOR2
            initialValidators[2] = 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC; // VALIDATOR3
            initialValidators[3] = 0x90F79bf6EB2c4f870365E785982E1f101E93b906; // VALIDATOR4
            initialValidators[4] = 0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65; // VALIDATOR5
        }

        return initialValidators;
    }

    /**
     * @dev 初始化所有合约 - 使用实际部署的合约地址
     */
    function initializeContracts(
        address validators, 
        address proposal, 
        address punish,
        address staking,
        address[] memory initialValidators
    ) internal {
        console.log("Initializing contracts with actual deployed addresses...");
        
        console.log("Deployed contract addresses:");
        console.log("  Validators:", validators);
        console.log("  Proposal:", proposal);
        console.log("  Punish:", punish);
        console.log("  Staking:", staking);
        
        // 按依赖关系正确初始化合约
        
        // 1. 初始化 Staking (传入 validators 地址)
        console.log("Initializing Staking...");
        Staking(staking).initialize(validators);
        console.log("Staking initialized successfully");

        // 2. 初始化 Proposal (传入 validators 地址)
        console.log("Initializing Proposal...");
        Proposal(proposal).initialize(initialValidators, validators);
        console.log("Proposal initialized successfully");

        // 3. 初始化 Punish (传入 validators 和 proposal 地址)
        console.log("Initializing Punish...");
        Punish(punish).initialize(validators, proposal);
        console.log("Punish initialized successfully");

        // 4. 最后初始化 Validators (传入所有其他合约地址)
        console.log("Initializing Validators...");
        Validators(validators).initialize(initialValidators, proposal, punish, staking);
        console.log("Validators initialized successfully");

        console.log("=== All contracts initialized with correct addresses! ===");
        console.log("Contracts now reference actual deployed addresses instead of hardcoded ones.");

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
