# AAC System Contracts

## 🏗️ Project Structure

```
sys-contract/
├── contracts/           # Solidity source code
├── forge-scripts/       # Foundry deployment scripts
├── forge-tests/         # Foundry test suites
├── scripts/            # Node.js utility scripts
├── docs/               # Project documentation
├── cmd/                # Go command-line tools
├── foundry.toml        # Foundry configuration
├── package.json        # Node.js dependencies (minimal)
├── generate-contracts.js # Contract template generator
├── go.mod              # Go dependencies
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
# Update genesis.json with system contracts
npm run update-genesis
```

### **Go CLI Tools**

Command-line utilities for blockchain management.

```bash
# Build CLI tools
make build

# View available commands
./bin/contract help
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

## 📋 System Contracts

### **Contract Addresses** (Fixed)

- **Validators**: `0x000000000000000000000000000000000000f000`
- **Punish**: `0x000000000000000000000000000000000000f001`
- **Proposal**: `0x000000000000000000000000000000000000f002`

### **Deployment Flow**

1. Compile contracts: `forge build`
2. Generate contracts: `npm run generate`
3. Update genesis: `npm run update-genesis`
4. Start chain: `cd ../chain && ./start_private_chain.sh`

## 📚 Documentation

- **Deployment Guide**: `docs/deploy.md`
- **Congress Consensus**: `docs/congress.md`
- **Migration Reports**: `docs/FOUNDRY_MIGRATION_COMPLETED.md`
- **Coverage Analysis**: `docs/TEST_COVERAGE_FINAL.md`

## 🔧 Configuration

### **Foundry** (`foundry.toml`)

- Solidity version: 0.8.20
- Source directory: `contracts/`
- Test directory: `forge-tests/`

### **Node.js** (`package.json`)

Minimal dependencies for utility scripts:

- `nunjucks`: Template engine for contract generation

### **Go** (`go.mod`)

CLI tools and blockchain utilities built with Go 1.23+

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

# 5. Update genesis (if needed)
npm run update-genesis
```

## 📖 Further Reading

- [Foundry Book](https://book.getfoundry.sh/)
- [Congress Consensus Documentation](docs/congress.md)
- [Deployment Guide](docs/deploy.md)
