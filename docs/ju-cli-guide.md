# ju-cli Usage Guide

## Overview

ju-cli is a command-line tool for Juchain blockchain governance that provides comprehensive functionality for validator management, proposal voting, staking management, and network governance.

## 🚀 Core Features

- **Proposal Management**: Create, vote on, and query validator and configuration proposals
- **Validator Management**: Query validator information and manage earnings
- **Staking Management**: Validator registration, delegation, undelegation, and reward withdrawal
- **Transaction Processing**: Secure three-stage transaction workflow (generate → sign → send)


## 🔒 Secure Three-Stage Operation

For enhanced security, ju-cli implements a three-stage transaction workflow:

1. **Generate (Online)**: Create unsigned transaction data while connected to the network
2. **Sign (Offline)**: Sign the transaction using a private key or keystore file offline
3. **Send (Online)**: Broadcast the signed transaction to the network

This approach minimizes exposure of sensitive private keys while ensuring transaction integrity.

## Installation and Compilation

### Prerequisites

- Go 1.23.0 or higher
- Solidity compiler (solc 0.8.20)
- abigen tool (for generating Go bindings)

### Compilation Steps

```bash
# Compile contracts and generate Go bindings
make generate-go-client

# Enter command line directory
cd tools

# Compile executable
make build

# The generated executable is located at build/ju-cli
```

## Usage

### Global Parameters

All commands support the following global parameter:

- `-r, --rpc` - RPC endpoint URL
  - Testnet: `https://testnet-rpc.juchain.org`
  - Mainnet: `https://rpc.juchain.org`
  - Local: `http://localhost:8545`

### Command Structure

```
ju-cli
├── proposal      # Proposal management commands
├── validator     # Validator management commands
├── staking       # Staking and delegation commands
└── misc          # Miscellaneous commands (sign, send)
```

## 1. Proposal Management

### 1.1 Create Proposal

Create a proposal to add or remove a validator:

```bash
# Create add validator proposal
./build/ju-cli proposal create \
  -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  -t 0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc \
  -o add \
  -r https://testnet-rpc.juchain.org

# Create remove validator proposal
./build/ju-cli proposal create \
  -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  -t 0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc \
  -o remove \
  -r https://testnet-rpc.juchain.org
```

**Parameters:**
- `-p, --proposer` - Proposer address (must be a valid validator)
- `-t, --target` - Target address (validator to add or remove)
- `-o, --operation` - Operation type (`add` or `remove`)
- `-r, --rpc` - RPC endpoint URL

**Output:** `createProposal.json`

### 1.2 Create Configuration Proposal

Create a proposal to update system parameters:

```bash
./build/ju-cli proposal create-config \
  -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  -i 0 \
  -v 86400 \
  -r https://testnet-rpc.juchain.org
```

**Parameters:**
- `-p, --proposer` - Proposer address (must be a valid validator)
- `-i, --cid` - Configuration ID (0-9)
- `-v, --value` - New configuration value
- `-r, --rpc` - RPC endpoint URL

**Output:** `createUpdateConfigProposal.json`

### 1.3 Vote on Proposal

Vote on an existing proposal:

```bash
# Approve vote
./build/ju-cli proposal vote \
  -s 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  -i 0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c \
  -a \
  -r https://testnet-rpc.juchain.org

# Reject vote (omit -a)
./build/ju-cli proposal vote \
  -s 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  -i 0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c \
  -r https://testnet-rpc.juchain.org
```

**Parameters:**
- `-s, --signer` - Signer address (must be a valid validator)
- `-i, --proposalId` - Proposal ID (64-character hex string)
- `-a, --approve` - Approval flag (use -a for approve, omit for reject)
- `-r, --rpc` - RPC endpoint URL

**Output:** `voteProposal.json`

### 1.4 Query Proposals

```bash
# Query all proposals
./build/ju-cli proposal list -r https://testnet-rpc.juchain.org

# Query specific proposal
./build/ju-cli proposal query \
  -i 0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c \
  -r https://testnet-rpc.juchain.org
```

## 2. Validator Management

### 2.1 Query Validators

```bash
# Query all validators
./build/ju-cli validator list -r https://testnet-rpc.juchain.org

# Query specific validator
./build/ju-cli validator query \
  -a 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  -r https://testnet-rpc.juchain.org
```

### 2.2 Withdraw Validator Profits

```bash
./build/ju-cli validator withdraw-profits \
  -a 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  -r https://testnet-rpc.juchain.org
```

**Output:** `withdrawProfits.json`

## 3. Staking Management

### 3.1 Register Validator

```bash
./build/ju-cli staking validator-register \
  -p 0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc \
  -s 100000 \
  -c 500 \
  -r https://testnet-rpc.juchain.org
```

**Parameters:**
- `-p, --proposer` - Validator account address
- `-s, --stake-amount` - Staking amount (minimum 100,000 JU)
- `-c, --commission-rate` - Commission rate in basis points (0-10000)
- `-r, --rpc` - RPC endpoint URL

**Output:** `registerValidator.json`

### 3.2 Delegate

```bash
./build/ju-cli staking delegate \
  -d 0x970e8128ab834e3eac664312d6e30df9e93cb357 \
  -v 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  -s 1000 \
  -r https://testnet-rpc.juchain.org
```

**Parameters:**
- `-d, --delegator` - Delegator address
- `-v, --validator` - Validator address
- `-s, --amount` - Delegation amount (minimum 1 JU)
- `-r, --rpc` - RPC endpoint URL

**Output:** `delegate.json`

### 3.3 Undelegate

```bash
./build/ju-cli staking undelegate \
  -d 0x970e8128ab834e3eac664312d6e30df9e93cb357 \
  -v 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  -s 500 \
  -r https://testnet-rpc.juchain.org
```

**Parameters:**
- `-d, --delegator` - Delegator address
- `-v, --validator` - Validator address
- `-s, --amount` - Unbonding amount
- `-r, --rpc` - RPC endpoint URL

**Output:** `undelegate.json`

### 3.4 Claim Rewards

```bash
./build/ju-cli staking claim-rewards \
  -c 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  -v 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  -r https://testnet-rpc.juchain.org
```

**Parameters:**
- `-c, --claimer` - Claimer address
- `-v, --validator` - Validator address
- `-r, --rpc` - RPC endpoint URL

**Output:** `claimRewards.json`

### 3.5 Withdraw Unbonded Stakes

```bash
./build/ju-cli staking withdraw-unbonded \
  -c 0x970e8128ab834e3eac664312d6e30df9e93cb357 \
  -v 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  -r https://testnet-rpc.juchain.org
```

**Parameters:**
- `-c, --claimer` - Claimer address
- `-v, --validator` - Validator address
- `-r, --rpc` - RPC endpoint URL

**Output:** `withdrawUnbonded.json`

## 4. Transaction Signing and Sending

### 4.1 Sign Transaction

```bash
# Using keystore file
./build/ju-cli misc sign \
  -f createProposal.json \
  -w /path/to/keystore/UTC--xxx \
  -p /path/to/password.txt

# Using private key directly
./build/ju-cli misc sign \
  -f createProposal.json \
  -k 0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d
```

**Parameters:**
- `-f, --file` - Transaction file path
- `-w, --wallet` - Keystore file path (required if not using private key)
- `-k, --private-key` - Private key in hex format (required if not using wallet)
- `-p, --password` - Password file path (required when using wallet file)

**Output:** `[original filename]_signed.json`

### 4.2 Send Transaction

```bash
./build/ju-cli misc send \
  -f createProposal_signed.json \
  -r https://testnet-rpc.juchain.org
```

**Parameters:**
- `-f, --file` - Signed transaction file path
- `-r, --rpc` - RPC endpoint URL

## 5. Three-Stage Operation Example

```bash
# Stage 1: Generate transaction (online)
./build/ju-cli proposal create \
  -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  -t 0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc \
  -o add \
  -r https://testnet-rpc.juchain.org

# Stage 2: Sign transaction (offline recommended)
./build/ju-cli misc sign \
  -f createProposal.json \
  -w /path/to/keystore \
  -p /path/to/password.txt

# Stage 3: Send transaction (online)
./build/ju-cli misc send \
  -f createProposal_signed.json \
  -r https://testnet-rpc.juchain.org
```

## 6. Network Endpoints

- **Testnet**: `https://testnet-rpc.juchain.org` (Chain ID: 202599)
- **Mainnet**: `https://rpc.juchain.org` (Chain ID: 210000)
- **Local**: `http://localhost:8545` (Chain ID: 202599)

## 7. Troubleshooting

### Common Errors

1. **EIP-155 Error**: Ensure you're using the correct chain ID when signing transactions
2. **Gas Estimation Failed**: Check parameters and network connectivity
3. **Validator Only**: Only active validators can propose or vote
4. **Proposal Expired**: Validators must register within 7 days after proposal approval

### Debugging Tips

- Use `--help` with any command for detailed usage information
- Check generated JSON files for transaction details
- Verify keystore and password file paths
- Ensure RPC endpoint is correct and node is synced

## 8. Tool Information

### Version

```bash
./build/ju-cli version
```

### Help

```bash
./build/ju-cli --help
./build/ju-cli [command] --help
```
