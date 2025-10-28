# Clef External Signer Account Management

Clef serves as an external account management and signing tool, and can even run on dedicated secure external hardware USB devices.

## Ethereum Account Encrypted Storage File (Keystore V3 Format)

The user-entered **password** is used to derive a symmetric key through scrypt, then the ciphertext is obtained by encrypting with AES-128-CTR.

```json
{
  "address": "1fe9327d22584e2a8eec4539c541cb0ad897f698",  // Account address (lowercase hex)
  "crypto": {
    "cipher": "aes-128-ctr",                          // Encryption algorithm: AES-128-CTR
    "ciphertext": "811d...",                          // Encrypted private key content
    "cipherparams": {
      "iv": "202cbc5aa0448179538a468b6a26bc55"        // Initialization vector
    },
    "kdf": "scrypt",                                  // Key Derivation Function
    "kdfparams": {
      "dklen": 32,
      "n": 262144,                                    // scrypt parameter, larger is more secure but slower
      "p": 1,
      "r": 8,
      "salt": "280c..."                               // Random salt value
    },
    "mac": "c1d7..."                                  // Hash for verifying decryption correctness
  },
  "id": "1b124e30-73d3-413d-9bd0-4e1a15f6d285",       // UUID
  "version": 3
}

```

Users need to manually review all operations involving sensitive data, and signing is completed locally in Clef.

Create and list accounts, or sign data offline:

```bash
clef init --configdir clefdata

# Create new account
clef newaccount --keystore <path-to-keystore>

# Import raw private key
clef importraw <hexkey>

# List accounts
clef list-accounts --keystore <path-to-keystore>

clef list-wallets --keystore <path-to-keystore>

# Start signer
clef --keystore keys --configdir clefdata --chainid 202599 --http
```

WARNING!

Clef is an account management tool. It may, like any software, contain bugs.

Please take care to

- backup your keystore files,
- verify that the keystore(s) can be opened with your password.

Clef is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY;

without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR

PURPOSE. See the GNU General Public License for more details.

Enter 'ok' to proceed:

> ok

INFO [08-22|14:34:14.621] Using CLI as UI-channel

INFO [08-22|14:34:14.734] Loaded 4byte database                    embeds=268,621 locals=0 local=./4byte-custom.json

WARN [08-22|14:34:14.734] Failed to open master, rules disabled    err="failed stat on clefdata/masterseed.json: stat clefdata/masterseed.json: no such file or directory"

INFO [08-22|14:34:14.734] Starting signer                          chainid=202,599 keystore=keys light-kdf=false advanced=false

INFO [08-22|14:34:14.744] Audit logs configured                    file=audit.log

INFO [08-22|14:34:14.745] HTTP endpoint opened                     url=<http://127.0.0.1:8550/>

INFO [08-22|14:34:14.745] IPC endpoint opened                      url=clefdata/clef.ipc

- ------ Signer info -------
- intapi_version : 7.0.1
- extapi_version : 6.1.0
- extapi_http : <http://127.0.0.1:8550/>
- extapi_ipc : clefdata/clef.ipc
- ------ Available accounts -------

0. 0x3858FfcA201b0A7D75fd23BB302C12332c5e4000 at keystore:///Users/enty/ju-chain-work/chain/build/bin/keys/UTC--2025-08-22T06-33-36.854192000Z--3858ffca201b0a7d75fd23bb302c12332c5e4000

1. 0x3d968443D9B72bCeF4409B3A2D5e31031390FC82 at keystore:///Users/enty/ju-chain-work/chain/build/bin/keys/UTC--2025-08-22T06-33-49.461865000Z--3d968443d9b72bcef4409b3a2d5e31031390fc82

WARN [08-22|14:34:56.398] Served account_signTransaction           conn=127.0.0.1:58911 reqid=1 duration="247.202µs" err="invalid argument 0: json: cannot unmarshal non-string into Go struct field SendTxArgs.chainId of type *hexutil.Big"

ERROR[08-22|14:35:32.655] Signing request with wrong chain id      requested=202,391 configured=202,599

WARN [08-22|14:35:32.656] Served account_signTransaction           conn=127.0.0.1:59774 reqid=1 duration=1.252915ms  err="requested chainid 202391 does not match the configuration of the signer"

ERROR[08-22|14:36:26.262] Signing request with wrong chain id      requested=202,583 configured=202,599

WARN [08-22|14:36:26.262] Served account_signTransaction           conn=127.0.0.1:61145 reqid=1 duration="94.004µs"  err="requested chainid 202583 does not match the configuration of the signer"

ERROR[08-22|14:36:57.379] Signing request with wrong chain id      requested=202,583 configured=202,599

WARN [08-22|14:36:57.379] Served account_signTransaction           conn=127.0.0.1:61854 reqid=1 duration="98.968µs"  err="requested chainid 202583 does not match the configuration of the signer"

ERROR[08-22|14:38:40.566] Signing request with wrong chain id      requested=202,583 configured=202,599

WARN [08-22|14:38:40.566] Served account_signTransaction           conn=127.0.0.1:64264 reqid=1 duration="94.902µs"  err="requested chainid 202583 does not match the configuration of the signer"

- -------- Transaction request-------------

to:    0x1234567890123456789012345678901234567890

from:               0x3858ffca201b0a7d75fd23bb302c12332c5e4000 [chksum INVALID]

value:              1000000000000000000 wei

gas:                0x5208 (21000)

gasprice: 20000000000 wei

nonce:    0x0 (0)

chainid:  0x31767

Request context:

127.0.0.1:49531 -> http -> 127.0.0.1:8550

Additional HTTP header data, provided by the external caller:

User-Agent: "curl/8.1.2"

Origin: ""

- ------------------------------------------

Approve? [y/N]:

> > y

## Account password

Please enter the password for account 0x3858FfcA201b0A7D75fd23BB302C12332c5e4000

>

- ----------------------

Transaction signed:

{

"type": "0x0",

"chainId": "0x31767",

"nonce": "0x0",

"to": "0x1234567890123456789012345678901234567890",

"gas": "0x5208",

"gasPrice": "0x4a817c800",

"maxPriorityFeePerGas": null,

"maxFeePerGas": null,

"value": "0xde0b6b3a7640000",

"input": "0x",

"v": "0x62ef2",

"r": "0xfd80a31a67b54d3d9ffc0c4f27db69f369a22b00792b91102663252d35a89da3",

"s": "0x383db7f2a92dea4322793c2a393b760615ded4e99cfb7489a99bd7667d0bdeff",

"hash": "0xde20680041d236d200aa7dfbe41d2eea5c0d78a7e9fe64dc34ecd2a33c1f766e"

}

## Sign Transaction

```bash
curl -X POST \

-H "Content-Type: application/json" \

--data '{"jsonrpc":"2.0","method":"account_signTransaction","params":[{"from": "0x3858ffca201b0a7d75fd23bb302c12332c5e4000", "to": "0x1234567890123456789012345678901234567890", "value": "0xde0b6b3a7640000", "gas": "0x5208", "gasPrice": "0x4a817c800", "nonce": "0x0", "chainId": "0x31767"}], "id":1}' \

http://127.0.0.1:8550

{"jsonrpc":"2.0","id":1,"result":{"raw":"0xf86f808504a817c800825208941234567890123456789012345678901234567890880de0b6b3a76400008083062ef2a0fd80a31a67b54d3d9ffc0c4f27db69f369a22b00792b91102663252d35a89da3a0383db7f2a92dea4322793c2a393b760615ded4e99cfb7489a99bd7667d0bdeff","tx":{"type":"0x0","chainId":"0x31767","nonce":"0x0","to":"0x1234567890123456789012345678901234567890","gas":"0x5208","gasPrice":"0x4a817c800","maxPriorityFeePerGas":null,"maxFeePerGas":null,"value":"0xde0b6b3a7640000","input":"0x","v":"0x62ef2","r":"0xfd80a31a67b54d3d9ffc0c4f27db69f369a22b00792b91102663252d35a89da3","s":"0x383db7f2a92dea4322793c2a393b760615ded4e99cfb7489a99bd7667d0bdeff","hash":"0xde20680041d236d200aa7dfbe41d2eea5c0d78a7e9fe64dc34ecd2a33c1f766e"}}}

```

## Send Transaction

```bash
curl -X POST \

-H "Content-Type: application/json" \

--data '{"jsonrpc":"2.0","method":"eth_sendRawTransaction","params":["0xf86f808504a817c800825208941234567890123456789012345678901234567890880de0b6b3a76400008083062ef2a0fd80a31a67b54d3d9ffc0c4f27db69f369a22b00792b91102663252d35a89da3a0383db7f2a92dea4322793c2a393b760615ded4e99cfb7489a99bd7667d0bdeff"],"id":1}' \

<http://127.0.0.1:8556>

{"jsonrpc":"2.0","id":1,"result":"0xde20680041d236d200aa7dfbe41d2eea5c0d78a7e9fe64dc34ecd2a33c1f766e"}%

···

## Query Transaction Status Receipt

```bash

curl -X POST \

-H "Content-Type: application/json" \

--data '{"jsonrpc":"2.0","method":"eth_getTransactionReceipt","params":["0xde20680041d236d200aa7dfbe41d2eea5c0d78a7e9fe64dc34ecd2a33c1f766e"],"id":1}' \

http://127.0.0.1:8556

```

## 4. Using Geth Console to Interact with JuChain System Contracts

### 4.1 Preparation: Check Network Status

First, check the current node status in the Geth console:

```jsx
// Check basic network information
eth.blockNumber           // Current block number
eth.mining               // Whether mining
net.version              // Network ID (should be 202599)
admin.peers.length       // Number of connected nodes

// Check preset account balance
var account1 = "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
var account2 = "0x3858ffca201b0a7d75fd23bb302c12332c5e4000"

console.log("Account 1 balance:", web3.fromWei(eth.getBalance(account1), "ether"), "JU")
console.log("Account 2 balance:", web3.fromWei(eth.getBalance(account2), "ether"), "JU")
```

### 4.2 Account Management and Unlocking

```jsx
// View all available accounts
eth.accounts

// Unlock main account for transactions (modify password as needed)
personal.unlockAccount(account1, "password123", 300)  // Unlock for 300 seconds

// Set mining account and start mining (if not already started)
miner.setEtherbase(account1)
if (!eth.mining) {
    miner.start(1)
    console.log("Mining started")
}
```

### 4.3 Send Basic Transaction Test

```jsx
// Send simple transfer transaction
var txHash = eth.sendTransaction({
    from: account1,
    to: account2,
    value: web3.toWei(1, "ether"),
    gas: 21000,
    gasPrice: web3.toWei(20, "gwei")
})

console.log("Transaction hash:", txHash)

// Wait for transaction confirmation and check receipt
setTimeout(function() {
    var receipt = eth.getTransactionReceipt(txHash)
    console.log("Transaction status:", receipt ? "Confirmed" : "Pending")
    if (receipt) {
        console.log("Gas used:", receipt.gasUsed)
        console.log("Block number:", receipt.blockNumber)
    }
}, 3000)
```

## 5. Interact with JuChain System Contracts

### 5.1 Initialize System Contract Connections

JuChain system contracts are pre-deployed at fixed addresses in the genesis block, and we can interact with them directly:

```jsx
// Define system contract addresses
var CONTRACT_ADDRESSES = {
    validators: "0x000000000000000000000000000000000000f000",
    punish: "0x000000000000000000000000000000000000f001", 
    proposal: "0x000000000000000000000000000000000000f002",
    staking: "0x000000000000000000000000000000000000f003"
}

// Validators contract ABI (core validator management)
var validatorsABI = [
    {"inputs":[],"name":"getActiveValidators","outputs":[{"internalType":"address[]","name":"","type":"address[]"}],"stateMutability":"view","type":"function"},
    {"inputs":[],"name":"getTopValidators","outputs":[{"internalType":"address[]","name":"","type":"address[]"}],"stateMutability":"view","type":"function"},
    {"inputs":[{"internalType":"address","name":"val","type":"address"}],"name":"getValidatorInfo","outputs":[{"internalType":"address","name":"feeAddr","type":"address"},{"internalType":"uint256","name":"status","type":"uint256"},{"internalType":"uint256","name":"accumulatedRewards","type":"uint256"},{"internalType":"uint256","name":"totalJailedHB","type":"uint256"},{"internalType":"uint256","name":"lastWithdrawProfitsBlock","type":"uint256"}],"stateMutability":"view","type":"function"}
]

// Staking contract ABI (staking management)
var stakingABI = [
    {"inputs":[{"internalType":"uint256","name":"commissionRate","type":"uint256"}],"name":"register","outputs":[],"stateMutability":"payable","type":"function"},
    {"inputs":[{"internalType":"address","name":"validator","type":"address"}],"name":"getValidatorInfo","outputs":[{"internalType":"uint256","name":"selfStake","type":"uint256"},{"internalType":"uint256","name":"totalDelegated","type":"uint256"},{"internalType":"uint256","name":"totalStake","type":"uint256"},{"internalType":"uint256","name":"commissionRate","type":"uint256"},{"internalType":"bool","name":"isJailed","type":"bool"},{"internalType":"uint256","name":"jailUntilBlock","type":"uint256"}],"stateMutability":"view","type":"function"},
    {"inputs":[{"internalType":"uint256","name":"limit","type":"uint256"}],"name":"getTopValidators","outputs":[{"internalType":"address[]","name":"","type":"address[]"}],"stateMutability":"view","type":"function"}
]

// Create contract instances
var validatorsContract = eth.contract(validatorsABI).at(CONTRACT_ADDRESSES.validators)
var stakingContract = eth.contract(stakingABI).at(CONTRACT_ADDRESSES.staking)

console.log("✅ System contracts connected")
console.log("📊 Validators contract:", CONTRACT_ADDRESSES.validators)
console.log("💰 Staking contract:", CONTRACT_ADDRESSES.staking)
```

### 5.2 Query Validator Information

```jsx
// Query current active validators
console.log("\n=== 📋 Active Validators List ===")
var activeValidators = validatorsContract.getActiveValidators()
console.log("Active validators count:", activeValidators.length)
activeValidators.forEach(function(addr, index) {
    console.log((index + 1) + ".", addr)
})

// Query top validators (from Staking contract)
console.log("\n=== 🏆 Top Validators (by stake) ===")
var topValidators = stakingContract.getTopValidators(21)
console.log("Top validators count:", topValidators.length)

// Query detailed information of specific validator
console.log("\n=== 🔍 Validator Details ===")
var targetValidator = activeValidators[0]  // Use first active validator
console.log("Query validator:", targetValidator)

// Query from Validators contract
var validatorInfo = validatorsContract.getValidatorInfo(targetValidator)
console.log("\n📊 Validators Contract Info:")
console.log("  Fee address:", validatorInfo[0])
console.log("  Status code:", validatorInfo[1].toString())
console.log("  Accumulated rewards:", web3.fromWei(validatorInfo[2], "ether"), "JU")
console.log("  Jail count:", validatorInfo[3].toString())
console.log("  Last withdrawal block:", validatorInfo[4].toString())

// Query from Staking contract
var stakingInfo = stakingContract.getValidatorInfo(targetValidator)
console.log("\n💰 Staking Contract Info:")
console.log("  Self-stake:", web3.fromWei(stakingInfo[0], "ether"), "JU")
console.log("  Total delegated:", web3.fromWei(stakingInfo[1], "ether"), "JU") 
console.log("  Total stake:", web3.fromWei(stakingInfo[2], "ether"), "JU")
console.log("  Commission rate:", (stakingInfo[3].toNumber() / 100).toFixed(2) + "%")
console.log("  Is jailed:", stakingInfo[4])
console.log("  Jailed until block:", stakingInfo[5].toString())
```

### 5.3 Network Status Overview

```jsx
// Comprehensive network status query
console.log("\n=== 🌐 JuChain Network Status ===")
console.log("Current block height:", eth.blockNumber)
console.log("Network ID:", net.version)
console.log("Is mining:", eth.mining)
console.log("Connected nodes:", admin.peers.length)

var latestBlock = eth.getBlock("latest")
console.log("Latest block info:")
console.log("  Block hash:", latestBlock.hash)
console.log("  Miner address:", latestBlock.miner)
console.log("  Transaction count:", latestBlock.transactions.length)
console.log("  Gas used:", latestBlock.gasUsed.toString())
console.log("  Timestamp:", new Date(latestBlock.timestamp * 1000))

// Check if PoSA consensus
var isValidator = activeValidators.indexOf(latestBlock.miner) !== -1
### 5.4 Register New Validator (Advanced)

```jsx
// ⚠️ Warning: This is a real operation requiring significant stake
// Only execute if you actually want to register as validator

console.log("\n=== 💎 Validator Registration Process ===")

// Check account balance (requires at least 10000 JU)
var registrationAccount = account1  // Use previously defined account
var currentBalance = web3.fromWei(eth.getBalance(registrationAccount), "ether")
var requiredStake = 10000  // Minimum stake requirement

console.log("Registration account:", registrationAccount)
console.log("Current balance:", currentBalance, "JU")
console.log("Minimum stake:", requiredStake, "JU")

if (parseFloat(currentBalance) >= requiredStake) {
    console.log("✅ Sufficient balance for registration")
    
    // Set commission rate (in basis points, 500 = 5%)
    var commissionRate = 500  // 5% commission
    
    console.log("\n📝 Prepare registration parameters:")
    console.log("  Stake amount:", requiredStake, "JU")
    console.log("  Commission rate:", (commissionRate/100).toFixed(2) + "%")
    
    // Execute registration (uncomment to actually execute)
    /*
    console.log("\n🚀 Registering validator...")
    var registerTx = stakingContract.register(commissionRate, {
        from: registrationAccount,
        value: web3.toWei(requiredStake, "ether"),
        gas: 500000,
        gasPrice: web3.toWei(20, "gwei")
    })
    
    console.log("Registration transaction hash:", registerTx)
    console.log("Please wait for transaction confirmation...")
    
    // Check transaction status
    setTimeout(function() {
        var receipt = eth.getTransactionReceipt(registerTx)
        if (receipt) {
            console.log("✅ Registration transaction confirmed")
            console.log("Gas used:", receipt.gasUsed)
            console.log("Transaction status:", receipt.status === "0x1" ? "Success" : "Failed")
            
            // Verify registration result
            var newStakingInfo = stakingContract.getValidatorInfo(registrationAccount)
            console.log("\n🎉 Post-registration info:")
            console.log("  Self-stake:", web3.fromWei(newStakingInfo[0], "ether"), "JU")
            console.log("  Commission rate:", (newStakingInfo[3].toNumber() / 100).toFixed(2) + "%")
        } else {
            console.log("⏳ Transaction still confirming...")
        }
    }, 5000)
    */
    
} else {
    console.log("❌ Insufficient balance to register validator")
    console.log("Need additional:", (requiredStake - parseFloat(currentBalance)).toFixed(2), "JU")
}
```

### 5.5 Complete Workflow Summary

```jsx
// 🎯 Complete validator query and management workflow
console.log("\n=== 🎯 JuChain Validator Management Overview ===")

function displayValidatorSummary() {
    var activeVals = validatorsContract.getActiveValidators()
    var topVals = stakingContract.getTopValidators(21)
    
    console.log("📊 Validator Statistics:")
    console.log("  Active validators:", activeVals.length)
    console.log("  Top validators:", topVals.length)
    
    console.log("\n🏆 Top 3 Validators:")
    for (var i = 0; i < Math.min(3, activeVals.length); i++) {
        var addr = activeVals[i]
        var stakingInfo = stakingContract.getValidatorInfo(addr)
        console.log((i+1) + ". " + addr)
        console.log("   Stake:", web3.fromWei(stakingInfo[2], "ether"), "JU")
        console.log("   Commission:", (stakingInfo[3].toNumber() / 100).toFixed(2) + "%")
        console.log("   Status:", stakingInfo[4] ? "Jailed" : "Normal")
    }
    
    console.log("\n📈 Network Health:")
    console.log("  Current block:", eth.blockNumber)
    console.log("  Latest block miner:", eth.getBlock("latest").miner)
    console.log("  PoSA consensus running:", activeVals.length > 0 ? "✅" : "❌")
}

// Execute summary
displayValidatorSummary()
```

**Execution Instructions:**

1. 🚀 **Execute step by step**: Copy the code segment by segment into Geth console and execute, review results after each segment
2. ⏱️ **Wait for confirmation**: Since your node is mining, transactions are usually confirmed within seconds
3. 🔍 **Real-time monitoring**: You can always use `eth.blockNumber` to check current block height
4. 📋 **Status verification**: Detailed results and status information will be displayed after each operation

## 6. Using Congress-CLI Tools

In addition to Geth console, JuChain also provides dedicated command-line tools to manage validators and governance:

### 6.1 Validator Query Commands

```bash
# 🔍 View all validators (Validators contract)
./build/congress-cli miners

# 👤 Query specific validator information
./build/congress-cli miner -a 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266

# 🏆 View top validators in Staking contract
./build/congress-cli staking list-top-validators

# 💰 Query specific validator information in Staking contract
./build/congress-cli staking query-validator --address 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
```

### 6.2 Governance Proposal Management

```bash
# 📝 Create proposal to add validator
./build/congress-cli create_proposal \
    -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
    -t 0xNew_Validator_Address \
    -o add

# 🗳️ Vote to support proposal
./build/congress-cli vote_proposal \
    -s 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
    -i Proposal_ID \
    -a

# 📊 Query proposal status
./build/congress-cli proposal -i 提案ID

# 📋 List all active proposals
./build/congress-cli list_proposals
```

### 6.3 Automation Scripts

```bash
# 🤖 Use automation script to add validator (includes complete process)
./sys-contract/congress-cli/add_validator6.sh

# This script will automatically execute:
# 1. Create proposal to add validator
# 2. Collect sufficient votes
# 3. Execute proposal
# 4. Register validator in Staking contract
# 5. Verify results of all steps
```

## 7. System Contract Address Reference

JuChain system contracts are deployed at the following fixed addresses:

- **Validators Contract**: `0x000000000000000000000000000000000000f000` - Manage validator status and rewards
- **Punish Contract**: `0x000000000000000000000000000000000000f001` - Handle validator punishment
- **Proposal Contract**: `0x000000000000000000000000000000000000f002` - Manage governance proposals
- **Staking Contract**: `0x000000000000000000000000000000000000f003` - Manage staking and delegation

These contracts are automatically initialized at genesis block, no manual deployment required.
