// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;

import {Script} from "forge-std/Script.sol";
import {stdJson} from "forge-std/StdJson.sol";
import {Test, console} from "forge-std/Test.sol";
import {Proposal} from "../generated/Proposal.sol";
import {Punish} from "../generated/Punish.sol";
import {Staking} from "../generated/Staking.sol";
import {Validators} from "../generated/Validators.sol";

// Base contract for all test utility contracts, providing common functionality
contract BaseTestUtils is Script, Test {
    using stdJson for string;
    
    // Test state management
    string public constant TEST_STATE_FILE = "./state/test_state.json"; // State file path
    
    // Contracts (only these are saved to state file)
    Proposal public proposal;
    Punish public punish;
    Staking public staking;
    Validators public validators;
    
    // Deployment keys
    uint256 public deployerKey;
    address public deployer;
    
    // Configuration from environment variables
    uint256 public constant COMMISSION_RATE_BASE = 10000;
    uint256 public initialStake;
    uint256 public initialValidators;
    uint256 public validatorMinBalance;
    uint256 public newValidatorFunding;
    uint256 public commissionRate;
    uint256 public minSelfStake;
    uint256 public minDelegation;
    uint256 public epochDuration;
    
    // Test accounts - managed by test scripts
    address[] public validatorAccounts;
    uint256[] public validatorKeys;
    address[] public delegatorAccounts;
    uint256[] public delegatorKeys;

    // Load test state from JSON file
    function loadState() public {
        // 1. Read the entire JSON file
        string memory json = vm.readFile(TEST_STATE_FILE); // forge-ignore unsafe-cheatcode
        
        // 2. Restore deployed contracts
        proposal = Proposal(json.readAddress(".contracts.proposal"));
        punish = Punish(json.readAddress(".contracts.punish"));
        staking = Staking(json.readAddress(".contracts.staking"));
        validators = Validators(json.readAddress(".contracts.validators"));

        // 3. Initialize test state from environment variables
        initializeState();
    }

    // Save current state to JSON file
    function saveState() public {
        string memory json = buildStateJson();
        vm.writeFile(TEST_STATE_FILE, json); // forge-ignore unsafe-cheatcode
        console.log("\nTest state saved to", TEST_STATE_FILE);
    }

    // Helper function to convert address to hex string
    function toHexString(address addr) internal pure returns (string memory) {
        return vm.toString(addr);
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
    
    // Create test accounts with sufficient funding
    function createTestAccounts() internal virtual {
        loadTestAccounts();
        
        // Fund deployer with sufficient ETH
        vm.deal(deployer, 100000000 ether);
        
        vm.startBroadcast(deployerKey);
        // Create initial validator accounts
        for (uint256 i = 0; i < 9; i++) {
            address validatorAddr = getValidatorAddr(i);
            address delegatorAddr = getDelegatorAddr(i);
            
            // Give validator enough ETH for transactions
            fundInitAccounts(validatorAddr, 110000);
            
            // Give delegator enough ETH for transactions
            fundInitAccounts(delegatorAddr, 10000);
        }
        vm.stopBroadcast();
        
        console.log("Test accounts created");
        console.log("Deployer address:", deployer);
        console.log("Deployer balance:", deployer.balance / 1 ether, "ETH");
    }

    // Create test accounts with sufficient funding
    function loadTestAccounts() internal virtual {
        // Initialize state from environment variables
        initializeState();
        // Create initial validator accounts
        for (uint256 i = 1; i < 10; i++) {
            uint256 validatorKey;
            uint256 delegatorKey;
            // Use vm.toString() to convert uint to string and then concatenate
            string memory minerVarName = string.concat("VALIDATOR_PRIVATE_KEY_", vm.toString(i));
            string memory delegatorVarName = string.concat("DELEGATOR_PRIVATE_KEY_", vm.toString(i));
            
            try vm.envUint(minerVarName) returns (uint256 key) {
                validatorKey = key;
            } catch {
            }
            try vm.envUint(delegatorVarName) returns (uint256 key) {
                delegatorKey = key;
            } catch {
            }
            
            address validatorAddr = vm.addr(validatorKey);
            address delegatorAddr = vm.addr(delegatorKey);
            
            // Add validator account to state
            validatorAccounts.push(validatorAddr);
            validatorKeys.push(validatorKey);
            
            // Add delegator account to state
            delegatorAccounts.push(delegatorAddr);
            delegatorKeys.push(delegatorKey);
        }
        console.log("Test accounts loaded");
    }

    
    function fundInitAccounts(address to, uint256 amountEth) internal {
        (bool success, ) = payable(to).call{value: amountEth * 1 ether}("");
        require(success, "fundAddress: transfer failed");
    }
    
    // Deploy and initialize all contracts
    function deployAndInitializeContracts() internal virtual {
        vm.startBroadcast(deployerKey);
        
        // Deploy contracts
        proposal = new Proposal();
        punish = new Punish();
        staking = new Staking();
        validators = new Validators();
        
        // Initialize Proposal contract
        address[] memory initialValidatorsArray = new address[](initialValidators);
        // Fill initialValidatorsArray with addresses from validatorAccounts
        for (uint256 i = 0; i < initialValidators; i++) {
            initialValidatorsArray[i] = validatorAccounts[i];
        }
        
        proposal.initialize(initialValidatorsArray, address(validators));
        
        // Initialize Validators contract
        validators.initialize(initialValidatorsArray, address(proposal), address(punish), address(staking));
        
        // Initialize Punish contract
        punish.initialize(address(validators), address(proposal), address(staking));
        
        // Initialize Staking contract with initial validators
        staking.initializeWithValidators(
            address(validators),
            address(proposal),
            initialValidatorsArray,
            commissionRate // Commission rate from environment variable
        );

        // Set contract addresses for all contracts
        validators.setContracts(
            address(validators),
            address(punish),
            address(proposal),
            address(staking)
        );
        
        staking.setContracts(
            address(validators),
            address(punish),
            address(proposal),
            address(staking)
        );
        
        punish.setContracts(
            address(validators),
            address(punish),
            address(proposal),
            address(staking)
        );
        
        proposal.setContracts(
            address(validators),
            address(punish),
            address(proposal),
            address(staking)
        );

        // Initial validators are automatically staked with minValidatorStake() during initializeWithValidators
        // No need to transfer ETH directly to Staking contract
        
        vm.stopBroadcast();

        saveState();
        
        console.log("Contracts deployed and initialized");
        console.log("Initial validators registered with minValidatorStake:", initialValidators, "validators");
        console.log("Staking contract balance:", address(staking).balance / 1 ether, "ETH");
    }
    
    // Fund a new validator with sufficient ETH for stake and fees
    function fundNewValidator(uint256 newValidatorKey) internal returns (address) {
        address newValidatorAddr = vm.addr(newValidatorKey);
        fundAddress(newValidatorAddr, newValidatorFunding);
        return newValidatorAddr;
    }
    
    // Fund an address with specified amount of wei
    function fundAddress(address to, uint256 amountWei) public {
        require(to != address(0), "fundAddress: zero address");
        require(amountWei > 0, "fundAddress: zero amount");

        vm.startBroadcast(deployerKey);

        (bool success, ) = payable(to).call{value: amountWei}("");
        require(success, "fundAddress: transfer failed");

        vm.stopBroadcast();

        console.log("Funded address:", to);
        console.log("Amount:", amountWei / 1 ether, "ETH");
    }
    
    // Print account balance
    function printBalance(address account) public view {
        uint256 balanceEth = account.balance / 1 ether;
        console.log("Account", account, "balance:", balanceEth);
    }

    // Build JSON string from current state (only contracts are saved)
    function buildStateJson() internal view returns (string memory) {
        // Start JSON object
        string memory json = "{\n";
        
        // Add contracts section (only these are saved)
        json = string.concat(json, "  \"contracts\": {\n");
        json = string.concat(json, "    \"proposal\": \"", toHexString(address(proposal)), "\",\n");
        json = string.concat(json, "    \"punish\": \"", toHexString(address(punish)), "\",\n");
        json = string.concat(json, "    \"staking\": \"", toHexString(address(staking)), "\",\n");
        json = string.concat(json, "    \"validators\": \"", toHexString(address(validators)), "\"\n");
        json = string.concat(json, "  }\n");
        
        // Close JSON object
        json = string.concat(json, "}");
        
        return json;
    }

    // Initialize state from environment variables
    function initializeState() internal {
        // Try to get VALIDATOR_COUNT from environment variables, default to 5 if not set
        try vm.envUint("VALIDATOR_COUNT") returns (uint256 count) {
            initialValidators = count;
        } catch {
            initialValidators = 5;
        }
        
        // Initialize other configuration from environment variables
        try vm.envUint("MIN_SELF_STAKE") returns (uint256 stake) {
            minSelfStake = stake * 1 ether; // Convert from ETH to wei
        } catch {
            minSelfStake = 100000 ether;
        }
        
        // Load commission rate from environment variables with proper unit conversion
        try vm.envUint("COMMISSION_RATE") returns (uint256 rate) {
            // If rate is greater than COMMISSION_RATE_BASE, it's likely in wei format (e.g., 50000000000000000 for 5%)
            if (rate > COMMISSION_RATE_BASE) {
                // Convert from wei format (e.g., 50000000000000000) to base format (e.g., 500)
                commissionRate = (rate * COMMISSION_RATE_BASE) / 1e18;
            } else {
                // Already in base format (e.g., 5000 for 50%)
                commissionRate = rate;
            }
        } catch {
            // Default to 5% in base format if environment variable not found
            commissionRate = 500;
        }
        
        // Load minimum delegation from environment variables
        try vm.envUint("MIN_DELEGATION") returns (uint256 delegation) {
            minDelegation = delegation * 1 ether; // Convert from ETH to wei
        } catch {
            // Default to 1 ether if environment variable not found
            minDelegation = 1 ether;
        }
        
        try vm.envUint("EPOCH_DURATION") returns (uint256 duration) {
            epochDuration = duration;
        } catch {
            epochDuration = 10;
        }
        
        try vm.envUint("INITIAL_STAKE") returns (uint256 stake) {
            initialStake = stake * 1 ether; // Convert from ETH to wei
        } catch {
            initialStake = minSelfStake; // Default to same as minSelfStake
        }
        
        try vm.envUint("VALIDATOR_MIN_BALANCE") returns (uint256 balance) {
            validatorMinBalance = balance;
        } catch {
            validatorMinBalance = 100 ether; // Default for transaction fees
        }
        
        try vm.envUint("NEW_VALIDATOR_FUNDING") returns (uint256 funding) {
            newValidatorFunding = funding;
        } catch {
            newValidatorFunding = initialStake + 10000 ether; // Default: Stake + fees
        }
        
        // Load deployer key from environment variables
        try vm.envUint("DEPLOYER_KEY") returns (uint256 key) {
            deployerKey = key;
        } catch {
            // Default deployer key if not found in environment
            deployerKey = 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80;
        }
        
        deployer = vm.addr(deployerKey);
    }

    // Get validator key by index
    function getValidatorKey(uint256 index) internal view returns (uint256) {
        require(index < validatorKeys.length, "Validator index out of bounds");
        return validatorKeys[index];
    }
    
    // Get validator address by index
    function getValidatorAddr(uint256 index) internal view returns (address) {
        require(index < validatorKeys.length, "Validator index out of bounds");
        return vm.addr(validatorKeys[index]);
    }

    // Get delegator key by index
    function getDelegatorKey(uint256 index) internal view returns (uint256) {
        require(index < delegatorKeys.length, "Delegator index out of bounds");
        return delegatorKeys[index];
    }
    
    // Get delegator address by index
    function getDelegatorAddr(uint256 index) internal view returns (address) {
        require(index < delegatorKeys.length, "Delegator index out of bounds");
        return vm.addr(delegatorKeys[index]);
    }
    
    // Set miner address temporarily for specific operations
    function setMinerTemporarily(address miner) internal {
        proposal.setMiner(miner);
        validators.setMiner(miner);
        staking.setMiner(miner);
        punish.setMiner(miner);
    }
    
    // Main run function to be implemented by derived classes
    function run() public virtual {
        // To be implemented by derived classes
    }
}
