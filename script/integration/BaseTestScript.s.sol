// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {Script} from "forge-std/Script.sol";
import {Test, console} from "forge-std/Test.sol";
import {Proposal} from "../../contracts/Proposal.sol";
import {Punish} from "../../contracts/Punish.sol";
import {Staking} from "../../contracts/Staking.sol";
import {Validators} from "../../contracts/Validators.sol";

// Base class for all test scripts, providing common functionality
contract BaseTestScript is Script, Test {
    // Commission rate base value (matches Staking.sol)
    uint256 public constant COMMISSION_RATE_BASE = 10000;
    
    // Configuration
    uint256 public immutable INITIAL_STAKE;
    uint256 public immutable INITIAL_VALIDATORS;
    uint256 public immutable VALIDATOR_MIN_BALANCE;
    uint256 public immutable NEW_VALIDATOR_FUNDING;
    uint256 public immutable COMMISSION_RATE;
    uint256 public immutable MIN_SELF_STAKE;
    uint256 public immutable MIN_DELEGATION;
    uint256 public immutable EPOCH_DURATION;
    
    // Constructor to initialize from environment variables
    constructor() {
        // Try to get VALIDATOR_COUNT from environment variables, default to 5 if not set
        try vm.envUint("VALIDATOR_COUNT") returns (uint256 count) {
            INITIAL_VALIDATORS = count;
        } catch {
            INITIAL_VALIDATORS = 5;
        }
        
        // Initialize other configuration from environment variables
        try vm.envUint("MIN_SELF_STAKE") returns (uint256 stake) {
            MIN_SELF_STAKE = stake * 1 ether; // Convert from ETH to wei
        } catch {
            MIN_SELF_STAKE = 100000 ether;
        }
        
        // Load commission rate from environment variables with proper unit conversion
        try vm.envUint("COMMISSION_RATE") returns (uint256 rate) {
            // If rate is greater than COMMISSION_RATE_BASE, it's likely in wei format (e.g., 50000000000000000 for 5%)
            if (rate > COMMISSION_RATE_BASE) {
                // Convert from wei format (e.g., 50000000000000000) to base format (e.g., 500)
                COMMISSION_RATE = (rate * COMMISSION_RATE_BASE) / 1e18;
            } else {
                // Already in base format (e.g., 5000 for 50%)
                COMMISSION_RATE = rate;
            }
        } catch {
            // Default to 5% in base format if environment variable not found
            COMMISSION_RATE = 500;
        }
        
        // Load minimum delegation from environment variables
        try vm.envUint("MIN_DELEGATION") returns (uint256 delegation) {
            MIN_DELEGATION = delegation * 1 ether; // Convert from ETH to wei
        } catch {
            // Default to 1 ether if environment variable not found
            MIN_DELEGATION = 1 ether;
        }
        
        try vm.envUint("EPOCH_DURATION") returns (uint256 duration) {
            EPOCH_DURATION = duration;
        } catch {
            EPOCH_DURATION = 10;
        }
        
        try vm.envUint("INITIAL_STAKE") returns (uint256 stake) {
            INITIAL_STAKE = stake * 1 ether; // Convert from ETH to wei
        } catch {
            INITIAL_STAKE = MIN_SELF_STAKE; // Default to same as MIN_SELF_STAKE
        }
        
        try vm.envUint("VALIDATOR_MIN_BALANCE") returns (uint256 balance) {
            VALIDATOR_MIN_BALANCE = balance;
        } catch {
            VALIDATOR_MIN_BALANCE = 100 ether; // Default for transaction fees
        }
        
        try vm.envUint("NEW_VALIDATOR_FUNDING") returns (uint256 funding) {
            NEW_VALIDATOR_FUNDING = funding;
        } catch {
            NEW_VALIDATOR_FUNDING = INITIAL_STAKE + 10000 ether; // Default: Stake + fees
        }
        
        // Load deployer key from environment variables
        try vm.envUint("DEPLOYER_KEY") returns (uint256 key) {
            deployerKey = key;
        } catch {
            // Default deployer key if not found in environment
            deployerKey = 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80;
        }
    }
    
    // Contracts
    Proposal public proposal;
    Punish public punish;
    Staking public staking;
    Validators public validators;
    
    // Test accounts
    address[] public validatorAccounts;
    uint256[] public validatorKeys;
    
    // Deployment keys
    uint256 deployerKey;
    address deployer;
    
    // Load deployer key from environment variables in constructor
    // Already handled in constructor via try/catch for vm.envUint("DEPLOYER_KEY")
    
    function run() public virtual {
        // To be implemented by derived classes
    }
    
    // Create test accounts with sufficient funding
    function createTestAccounts() internal virtual {
        // Use default key if not provided
        if (deployerKey == 0) {
            deployerKey = 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80;
        }
        
        deployer = vm.addr(deployerKey);
        
        // Create initial validator accounts
        for (uint256 i = 0; i < INITIAL_VALIDATORS; i++) {
            uint256 validatorKey;
            
            // Try to get validator private key from environment variables
            // Construct environment variable name using switch case for numbers 1-9
            string memory envVarName;
            if (i + 1 == 1) envVarName = "VALIDATOR_PRIVATE_KEY_1";
            else if (i + 1 == 2) envVarName = "VALIDATOR_PRIVATE_KEY_2";
            else if (i + 1 == 3) envVarName = "VALIDATOR_PRIVATE_KEY_3";
            else if (i + 1 == 4) envVarName = "VALIDATOR_PRIVATE_KEY_4";
            else if (i + 1 == 5) envVarName = "VALIDATOR_PRIVATE_KEY_5";
            
            try vm.envUint(envVarName) returns (uint256 key) {
                validatorKey = key;
            } catch {
                // Fall back to generated key if environment variable not found
                validatorKey = uint256(keccak256(abi.encodePacked("validator", i)));
            }
            
            address validatorAddr = vm.addr(validatorKey);
            validatorAccounts.push(validatorAddr);
            validatorKeys.push(validatorKey);
            
            // Give validator enough ETH for transactions
            fundAddress(validatorAddr, VALIDATOR_MIN_BALANCE);
            
        }
        
        // Fund deployer with sufficient ETH
        fundAddress(deployer, 100000000 ether);
        
        console.log(unicode"✓ Test accounts created");
        console.log("Deployer address:", deployer);
        console.log("Deployer balance:", deployer.balance / 1 ether, "ETH");
        console.log("Initial validators count:", validatorAccounts.length);
        
        // Print validator accounts and their balances
        for (uint256 i = 0; i < validatorAccounts.length; i++) {
            console.log("Validator", i+1, "address:", validatorAccounts[i]);
            uint256 balanceEth = validatorAccounts[i].balance / 1 ether;
            console.log("Validator", i+1, "balance:", balanceEth);
        }
    }
    
    // Deploy and initialize all contracts
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
        proposal.initialize(initialValidators, address(validators));
        
        // Initialize Validators contract
        validators.initialize(initialValidators, address(proposal), address(punish), address(staking));
        
        // Initialize Punish contract
        punish.initialize(address(validators), address(proposal), address(staking));
        
        // Initialize Staking contract with initial validators
        staking.initializeWithValidators(
            address(validators),
            address(proposal),
            initialValidators,
            COMMISSION_RATE // Commission rate from environment variable
        );
        

        
        // Initial validators are automatically staked with minValidatorStake() during initializeWithValidators
        // No need to transfer ETH directly to Staking contract
        
        vm.stopBroadcast();
        
        console.log(unicode"✓ Contracts deployed and initialized");
        console.log("Initial validators registered with minValidatorStake:", INITIAL_VALIDATORS, "validators");
        console.log("Staking contract balance:", address(staking).balance / 1 ether, "ETH");
    }
    

    
    // Fund a new validator with sufficient ETH for stake and fees
    function fundNewValidator(uint256 newValidatorKey) internal returns (address) {
        address newValidatorAddr = vm.addr(newValidatorKey);
        fundAddress(newValidatorAddr, NEW_VALIDATOR_FUNDING);
        return newValidatorAddr;
    }
    
    // Set miner address temporarily for specific operations
    function setMinerTemporarily(address miner) internal {
        // Miner address is set by the broadcast key, no need to call setMiner on contracts
        // console.log removed to avoid state mutability warning
    }
    
    // Get validator key by index
    function getValidatorKey(uint256 index) internal view returns (uint256) {
        require(index < validatorKeys.length, "Validator index out of bounds");
        return validatorKeys[index];
    }
    
    // Helper function to convert bytes32 to hex string
    function toHexString(bytes32 data) internal pure returns (string memory) {
        bytes memory alphabet = "0123456789abcdef";
        bytes memory str = new bytes(66);
        str[0] = '0';
        str[1] = 'x';
        for (uint256 i = 0; i < 32; i++) {
            uint8 b = uint8(uint256(data) / (2**(8*(31 - i))));
            str[2 + 2*i] = alphabet[b / 16];
            str[3 + 2*i] = alphabet[b % 16];
        }
        return string(str);
    }

    function printBalance(address account) public view {
        uint256 balanceEth = account.balance / 1 ether;
        console.log("Account", account, "balance:", balanceEth);
    }

    function fundAddress(address to, uint256 amountWei) public {
        require(to != address(0), "fundAddress: zero address");
        require(amountWei > 0, "fundAddress: zero amount");

        vm.startBroadcast(deployerKey);

        (bool success, ) = payable(to).call{value: amountWei}("");
        require(success, "fundAddress: transfer failed");

        vm.stopBroadcast();

        console.log("Funded address:", to);
        console.log("Amount:", amountWei, "wei");
    }

}