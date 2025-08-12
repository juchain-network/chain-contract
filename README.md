# AAC System Contracts

> Congress POA 共识系统合约

This repository contains the system contracts for the AAC blockchain's Congress Proof-of-Authority (POA) consensus mechanism. It includes validator management, governance proposals, punishment mechanisms, and a comprehensive CLI tool for network administration.

## ✨ Features

- 🏛️ **Congress POA Consensus**: Validator-based consensus with democratic governance
- 🗳️ **Proposal System**: Create and vote on network changes
- ⚖️ **Punishment Mechanism**: Automatic validator jailing for misbehavior  
- 💰 **Reward Distribution**: Fair fee sharing among active validators
- 🛠️ **CLI Management**: Complete command-line toolset for network operations
- 🧪 **Comprehensive Testing**: 40+ test cases with full coverage

## 🏗️ Project Structure

```
sys-contract/
├── contracts/           # Solidity source code
├── congress-cli/        # Congress POA management CLI tool
├── forge-scripts/       # Foundry deployment scripts
├── test/         # Foundry test suites
├── legacy-scripts/      # Legacy Node.js scripts
├── docs/               # Project documentation
├── out/                # Compiled contract artifacts (Foundry)
├── cache/              # Build cache
├── foundry.toml        # Foundry configuration
├── package.json        # Node.js dependencies (minimal)
├── generate-contracts.js # Contract template generator
├── init_genesis.js     # Genesis file initialization script
├── genesis.json        # Blockchain genesis configuration
└── README.md           # This file
```

## 🛠️ Development Tools

### **Primary: Foundry**

The main development and testing framework.

```bash
# Build contracts
forge build

# Run tests
forge test

# Run specific test
forge test --match-test testProposalCreation

# Run tests with verbosity
forge test -vvv
```

### **Contract Generation**

Generate production and mock contracts using templates.

```bash
# Generate production contracts
npm run generate

# Generate mock contracts for testing
npm run generate:mock
```

### **Genesis Configuration**

Update blockchain genesis file with compiled contract bytecode.

```bash
# Compile contracts first
forge build

# Update genesis.json with system contracts
npm run init-genesis
```

### **Congress CLI Tools**

Command-line utilities for Congress POA consensus management.

```bash
# Build CLI tools
cd congress-cli
make build

# View available commands
./build/congress-cli help

# Query miners
./build/congress-cli miners --rpc_laddr https://testnet-rpc.juchain.org

# Create proposal
./build/congress-cli create_proposal -p PROPOSER_ADDR -t TARGET_ADDR -o add --rpc_laddr RPC_ENDPOINT
```

## 🧪 Testing

### **Foundry Tests** (Primary)

Complete test coverage with 40+ test cases:

```bash
# Run all tests
forge test

# Test specific contracts
forge test --match-contract Proposal
forge test --match-contract Validators
forge test --match-contract Punish
```

### **Test Coverage**

- **Proposal Management**: Creation, voting, config updates
- **Validator Lifecycle**: Registration, rewards, punishment
- **Punishment System**: Thresholds, jailing, missed blocks
- **Reward Distribution**: Fee sharing, profit withdrawal

## 🏛️ Congress POA Management

### **Validator Operations**

```bash
# Query all active validators
./build/congress-cli miners --rpc_laddr https://testnet-rpc.juchain.org

# Query specific validator
./build/congress-cli miner --rpc_laddr https://testnet-rpc.juchain.org -a VALIDATOR_ADDRESS
```

### **Proposal Management**

```bash
# Create proposal to add new validator
./build/congress-cli create_proposal -p PROPOSER_ADDR -t NEW_VALIDATOR_ADDR -o add --rpc_laddr RPC_ENDPOINT

# Sign transaction
./build/congress-cli sign -f createProposal.json -k keyfile -p passwordfile --chainId CHAIN_ID

# Send transaction
./build/congress-cli send -f createProposal_signed.json --rpc_laddr RPC_ENDPOINT
```

### **Voting Process**

```bash
# Vote on proposal
./build/congress-cli vote_proposal -s VOTER_ADDR -i PROPOSAL_ID -a true --rpc_laddr RPC_ENDPOINT

# Sign and send vote
./build/congress-cli sign -f voteProposal.json -k keyfile -p passwordfile --chainId CHAIN_ID
./build/congress-cli send -f voteProposal_signed.json --rpc_laddr RPC_ENDPOINT
```

## 📋 System Contracts

### **Contract Addresses** (Fixed)

- **Validators**: `0x000000000000000000000000000000000000f000`
- **Punish**: `0x000000000000000000000000000000000000f001`
- **Proposal**: `0x000000000000000000000000000000000000f002`

### **Deployment Flow**

1. Compile contracts: `forge build`
2. Generate contracts: `npm run generate`
3. Initialize genesis: `npm run init-genesis`
4. Build CLI tools: `cd congress-cli && make build`
5. Start chain: `cd ../chain && ./start_private_chain.sh`

## 📚 Documentation

- **Congress CLI Guide**: `docs/congress.md` - Complete guide for Congress POA management
- **Deployment Guide**: `docs/deploy.md`
- **Migration Reports**: `docs/FOUNDRY_MIGRATION_COMPLETED.md`
- **Coverage Analysis**: `docs/TEST_COVERAGE_FINAL.md`

## 🌐 Network Information

- **Testnet**: `https://testnet-rpc.juchain.org` (Chain ID: 202599)
- **Mainnet**: `https://rpc.juchain.org` (Chain ID: 210000)

## 🔧 Configuration

### **Foundry** (`foundry.toml`)

- Solidity version: 0.8.20
- Source directory: `contracts/`
- Test directory: `test/`
- Output directory: `out/`

### **Node.js** (`package.json`)

Minimal dependencies for utility scripts:

- `nunjucks`: Template engine for contract generation

### **Congress CLI** (`congress-cli/`)

Go-based command-line tools for Congress POA consensus management.

## 🚀 Quick Start

```bash
# 1. Install dependencies
npm install

# 2. Build contracts
forge build

# 3. Run tests
forge test

# 4. Generate contracts
npm run generate

# 5. Initialize genesis (if needed)
npm run init-genesis

# 6. Build Congress CLI tools
cd congress-cli && make build

# 7. Test CLI tools
./build/congress-cli help
```

## 📖 Further Reading

- [Foundry Book](https://book.getfoundry.sh/)
- [Congress Consensus Documentation](docs/congress.md)
- [Deployment Guide](docs/deploy.md)
