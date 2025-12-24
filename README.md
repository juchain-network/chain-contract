# Ju System Contracts
This repository contains the system contracts for the Ju blockchain's JPoSA (JPoSA) consensus mechanism. It includes validator management, governance proposals, punishment mechanisms, and a comprehensive CLI tool for network administration.

## ✨ Features

- 🏛️ **Hybrid Consensus Mechanism**: Combines the advantages of PoS (Proof of Stake) and PoA (Proof of Authority) to achieve a balance between decentralization and high performance
- 🗳️ **Democratic Governance System**: Proposal governance mechanism supporting protocol upgrades and parameter adjustments with validator participation in network governance decisions
- ⚖️ **Dynamic Validator Set**: Automatically updates the validator list based on staking weight with a limited number of active validators and minimum staking requirement
- 💰 **Economic Incentive Model**: Sustainable economic model with moderate annual issuance rate, fair reward distribution between validators and delegators, and delegation mechanism for ordinary users to participate
- 🔒 **Multi-layer Security Protection**: Comprehensive security mechanisms including unbonding period, automatic validator jailing for misbehavior, and re-entry attack protection

## 🏗️ Project Structure

```
chain-contract/
├── tools/               # Congress JPoSA management CLI tool
├── contracts/           # Solidity source code
├── script/              # Foundry deployment scripts
├── test/                # Foundry test suites
├── docs/               # Project documentation
├── foundry.toml        # Foundry configuration
├── foundry.lock        # Foundry dependency lock file
├── package.json        # Node.js dependencies
├── generate-contracts.js # Contract template generator
├── init_genesis.js     # Genesis file initialization script
├── check_system_status.js # System status checking utility
├── extract-bytecode.js # Bytecode extraction utility
└── README.md           # This file
```

## 📚 Documentation
- [**JPoSA Whitepaper**](docs/posa-whitepaper.md) - Whitepaper of JPoSA consensus mechanism
- [**JPoSA Whitepaper (Chinese)**](docs/posa-whitepaper-zh.md) - Chinese version of JPoSA consensus mechanism whitepaper
- [**JPoSA Technical Specification**](docs/posa-tech-spec.md) - Detailed technical specification of JPoSA consensus mechanism
- [**Deployment Guide**](docs/deployment-guide.md) - Complete deployment and configuration guide
- [**Congress CLI Guide**](docs/ju-cli-guide.md) - Complete guide for Congress JPoSA management



## 📋 System Contracts

### **Deployment Flow**

1. Compile contracts: `forge build`
2. Generate contracts: `npm run generate`
3. Initialize genesis: `npm run init-genesis`
4. Build CLI tools: `cd congress-cli && make build`
5. Start chain: `cd ../chain/local-test && ./pm2-init.sh` or `pm2 start ecosystem.config.js`


## 🌐 Network Information

### RPC Endpoints

- **Testnet**: `https://testnet-rpc.juchain.org` (Chain ID: 202599)
- **Mainnet**: `https://rpc.juchain.org` (Chain ID: 210000)

### Block Explorers

- **Testnet Explorer**: https://testnet.juscan.io
- **Mainnet Explorer**: https://juscan.io


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

## 🏛️ Congress JPoSA Management

### **Validator Operations**

```bash
# Query all active validators
./build/ju-cli validator list
# Query specific validator
./build/ju-cli validator -a VALIDATOR_ADDRESS
```

### **Proposal Management**

```bash
# Create proposal to add new validator
./build/ju-cli proposal create -p PROPOSER_ADDR -t NEW_VALIDATOR_ADDR -o add

# Sign transaction
./build/ju-cli misc sign -f createProposal.json -w keyfile -p passwordfile

# Send transaction
./build/ju-cli misc send -f createProposal_signed.json
```


### **Voting Process**

```bash
# Vote on proposal
./build/ju-cli proposal vote -s VOTER_ADDR -i PROPOSAL_ID -a true

# Sign and send vote
./build/ju-cli misc sign -f voteProposal.json -w keyfile -p passwordfile
./build/ju-cli misc send -f voteProposal_signed.json
```


## 🚀 Quick Start

```bash
# Install nodejs 
# https://nodejs.org/en/download
# Install foundry
# curl -L https://foundry.paradigm.xyz | bash
# foundryup 

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

# 6. Generate contract bytecode (if needed)
npm run build-and-extract

# 7. Build Congress CLI tools
cd tools && make build

# 8. Test CLI tools
./build/ju-cli help
```