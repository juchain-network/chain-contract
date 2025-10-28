# Congress CLI v1.1.0

Command-line tool for Juchain blockchain governance, used for validator management, proposal voting, and network governance.

## 🚀 What's New (v1.1.0)

- ✅ **Enhanced Input Validation**: Complete parameter validation and error prompts
- ✅ **Better Error Handling**: Structured error messages and detailed error information
- ✅ **Improved User Experience**: Colorized output and clear status indicators
- ✅ **Global Parameter Validation**: Automatic RPC address and chain ID validation
- ✅ **Centralized Configuration Management**: Unified constant and configuration management
- ✅ **Example Commands**: Built-in usage examples and help documentation
- ✅ **Improved Voting System**: Simplified voting syntax
- ✅ **Enhanced Version Information**: Detailed build and version information

## Feature Overview

Congress CLI is a command-line tool for Juchain blockchain governance, providing complete validator management and proposal voting functionality.

### Core Features

- **Proposal Management**: Create validator add/remove proposals and configuration update proposals
- **Voting System**: Vote on proposals (supports simplified voting syntax)
- **Validator Management**: Query validator information and manage rewards
- **Transaction Processing**: Sign and send transactions to the blockchain network
- **Input Validation**: Comprehensive parameter validation and error handling

## Installation and Compilation

### Prerequisites

- Go 1.23.0 or higher
- Solidity compiler (solc 0.8.20)
- abigen tool (for generating Go bindings)

### Compilation Steps

```bash
# Navigate to project directory
cd sys-contract/congress-cli

# Compile contracts and generate Go bindings
make proposal

# Compile executable
make build

# Generated executable located at build/congress-cli
```

### Makefile Targets

- `make build` - Compile complete project
- `make proposal` - Generate Go bindings for Proposal contract
- `make cleanContract` - Clean generated contract files
- `make clean` - Clean build files

## Usage Guide

### Global Parameters

All commands support the following global parameters:

- `-c, --chainId int` - Specify chain ID (testnet: 202599, mainnet: 210000)
- `-l, --rpc_laddr string` - Specify RPC endpoint address
  - Testnet: `https://testnet-rpc.juchain.org`
  - Mainnet: `https://rpc.juchain.org`
  - Local: `http://localhost:8545`

⚠️ **Note**: New version automatically validates these parameters

### Quick Start

1. **View help and examples**:

```bash
./build/congress-cli --help
./build/congress-cli examples
./build/congress-cli [command] --help  # View help for specific command
```

2. **View version information**:

```bash
./build/congress-cli version
```

### Network Configuration

**Test Network**:

```bash
# Global parameter template
./build/congress-cli [command] -c 202599 -l https://testnet-rpc.juchain.org [other parameters]
```

**Main Network**:

```bash
# Global parameter template  
./build/congress-cli [command] -c 210000 -l https://rpc.juchain.org [other parameters]
```

### Command Details

#### 1. Query Validator Information

**Query all validators:**

```bash
# Testnet
./build/congress-cli miners

# Mainnet
./build/congress-cli miners
```

**Query specific validator:**

```bash
# Testnet示例
./build/congress-cli miner -a 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b

# Mainnet示例  
./build/congress-cli miner -a 0x311B37f01c04B84d1f94645BfBd58D82fc03F709
```

Output example:

```text
Address: 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b
Fee Address: 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b
Status: 1
Accumulated Rewards: 5035784561530401884829
Penalized Rewards: 5323260025816819260865
Last Withdraw Block: 1206974
```

**Status explanation:**

- Status 1 = Active
- Status 2 = Inactive

#### 2. Create Proposal

**Create validator add/remove proposal:**

```bash
# Add validator (testnet example)
./build/congress-cli create_proposal \
  -p 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  -t 0x029DAB47e268575D4AC167De64052FB228B5fA41 \
  -o add

# Remove validator (mainnet example)
./build/congress-cli create_proposal \
  -p 0xccafa71c31bc11ba24d526fd27ba57d743152807 \
  -t 0x4d432df142823ca25b21bc3f9744ed21a275bdea \
  -o remove
```

Parameter description:

- `-p, --proposer` - Proposer address (must be a valid validator)
- `-t, --target` - Target address (validator to add or remove)
- `-o, --operation` - Operation type (add or remove)

**Create configuration update proposal:**

```bash
# Testnet示例：修改 proposalLastingPeriod 为 86400
./build/congress-cli create_config_proposal \
  -p 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  -i 0 \
  -v 86400

# Mainnet示例
./build/congress-cli create_config_proposal \
  -p 0xccafa71c31bc11ba24d526fd27ba57d743152807 \
  -i 0 \
  -v 86400
```

Parameter description:

- `-p, --proposer` - Proposer address
- `-i, --cid` - Configuration ID:
  - 0: proposalLastingPeriod (proposal duration)
  - 1: punishThreshold (punishment threshold)
  - 2: removeThreshold (removal threshold)
  - 3: decreaseRate (reduction rate)
- 4: withdrawProfitPeriod (profit withdrawal period)
- `-v, --value` - New configuration value

#### 3. Vote on Proposal

⚠️ **Important**: Voting syntax optimized! Use `-a` flag for approval, omit for rejection.

```bash
# Approval vote (testnet example)
./build/congress-cli vote_proposal \
  -s 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  -i b2be7f3cc702c7a24962df6aed188edbcfdebe20fd55f5670efaedace0e4bcdf \
  -a

# Rejection vote (omit -a parameter)
./build/congress-cli vote_proposal \
  -s 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  -i b2be7f3cc702c7a24962df6aed188edbcfdebe20fd55f5670efaedace0e4bcdf

# Mainnet示例 (赞成票)
./build/congress-cli vote_proposal \
  -s 0xccafa71c31bc11ba24d526fd27ba57d743152807 \
  -i PROPOSAL_ID \
  -a
```

Parameter description:

- `-s, --signer` - Signer address (must be a valid validator)
- `-i, --proposalId` - Proposal ID (obtained from proposal creation output, 64-bit hexadecimal string)
- `-a, --approve` - Approval vote flag (use `-a` for approval, omit for rejection)

#### 4. Sign and Send Transaction

**Sign transaction:**

```bash
# Testnet示例
./build/congress-cli sign \
  -f createProposal.json \
  -k /path/to/keystore/UTC--xxx \
  -p /path/to/password.txt

# Mainnet示例
./build/congress-cli sign \
  -f createProposal.json \
  -k /path/to/keystore/miner1.key \
  -p /path/to/password.file
```

**Send signed transaction:**

```bash
# Testnet示例
./build/congress-cli send \
  -f createProposal_signed.json

# Mainnet示例
./build/congress-cli send \
  -f createProposal_signed.json
```

**Output example after successful send:**

```text
✅ Transaction broadcast successfully!
ℹ️  Transaction hash: 0xb72b3e4f2f4411fd467dcf3a4af16f12e5772a59ec91535ad18283c9a2e32ddf
ℹ️  Waiting for transaction confirmation: 0xb72b3e4f2f4411fd467dcf3a4af16f12e5772a59ec91535ad18283c9a2e32ddf
✅ Transaction confirmed in block 12535222
--------CreateProposal----------
Proposal ID: b2be7f3cc702c7a24962df6aed188edbcfdebe20fd55f5670efaedace0e4bcdf
Proposer: 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b
Destination: 0x029DAB47e268575D4AC167De64052FB228B5fA41
Flag: true
Time: 1754909524
Block: 12535222
-----
```

#### 5. Withdraw Validator Rewards

```bash
# Testnet示例
./build/congress-cli withdraw_profits -a 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b

# Mainnet示例
./build/congress-cli withdraw_profits -a 0xccafa71c31bc11ba24d526fd27ba57d743152807
```

⚠️ **Note**: 

- Profit withdrawal has minimum wait block limit
- Profit withdrawal does not require voting process, validators can directly withdraw their own profits
- Need to wait enough blocks to withdraw (controlled by withdrawProfitPeriod configuration)

## Complete Workflow Example

Below is a complete proposal creation and voting workflow (testnet example):

### 1. Query Current Validator Status

```bash
./build/congress-cli miners
```

### 2. Create Validator Add Proposal

```bash
./build/congress-cli create_proposal \
  -p 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  -t 0x029DAB47e268575D4AC167De64052FB228B5fA41 \
  -o add
```

### 3. Sign Transaction

```bash
./build/congress-cli sign \
  -f createProposal.json \
  -k miner1.key \
  -p password.file \
```

### 4. Send Transaction

```bash
./build/congress-cli send -f createProposal_signed.json
```

**Output example (record proposal ID):**

```text
✅ Transaction broadcast successfully!
ℹ️  Transaction hash: 0x484662b140a0e98ffd629cee763e12c5f79e7dfd312adbe8cd53b49a99e89c06
✅ Transaction confirmed in block 24805
--------CreateProposal----------
Proposal ID: b2be7f3cc702c7a24962df6aed188edbcfdebe20fd55f5670efaedace0e4bcdf
Proposer: 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b
Destination: 0x029DAB47e268575D4AC167De64052FB228B5fA41
Flag: true
Time: 1754905540
Block: 24805
-----
```

### 5. Multiple Validators Vote

Use the proposal ID obtained above to vote:

```bash
# miner1 approval vote
./build/congress-cli vote_proposal \
  -s 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  -i b2be7f3cc702c7a24962df6aed188edbcfdebe20fd55f5670efaedace0e4bcdf \
  -a
./build/congress-cli sign -f voteProposal.json -k miner1.key -p password.file
./build/congress-cli send -f voteProposal_signed.json
# miner2 approval vote
./build/congress-cli vote_proposal \
  -s 0x81f7a79a51edba249efa812eb2d5478f696f7558 \
  -i b2be7f3cc702c7a24962df6aed188edbcfdebe20fd55f5670efaedace0e4bcdf \
  -a
./build/congress-cli sign -f voteProposal.json -k miner2.key -p password.file
./build/congress-cli send -f voteProposal_signed.json

# miner3 approval vote
./build/congress-cli vote_proposal \
  -s 0x578c39eaf09a4e1abf428c423970b59bb8baf42e \
  -i b2be7f3cc702c7a24962df6aed188edbcfdebe20fd55f5670efaedace0e4bcdf \
  -a
./build/congress-cli sign -f voteProposal.json -k miner3.key -p password.file
./build/congress-cli send -f voteProposal_signed.json
```

### 6. Verify Results

```bash
# View new validator information
./build/congress-cli miner -a 0x029DAB47e268575D4AC167De64052FB228B5fA41

# View all validators
./build/congress-cli miners
```

## Configuration Files

### Generated Transaction Files

The tool generates the following JSON files in the current directory:

- `createProposal.json` - Original transaction for creating proposal
- `createProposal_signed.json` - Signed create proposal transaction
- `createUpdateConfigProposal.json` - Original transaction for configuration update proposal
- `createUpdateConfigProposal_signed.json` - Signed configuration update proposal transaction
- `voteProposal.json` - Original voting transaction
- `voteProposal_signed.json` - Signed voting transaction

### Keystore File Format

Uses standard Ethereum keystore format, for example:

```
UTC--202599-08-06T08-30-51.139143000Z--f39fd6e51aad88f6f4ce6ab8827279cffFb92266
```

## Troubleshooting

### Common Errors

**1. EIP-155 错误**

```
send tx error only replay-protected (EIP-155) transactions allowed over RPC
```

Solution: Must specify correct chain ID when signing:

```bash
./build/congress-cli sign -f transaction.json -k keystore -p password
```

**2. Profit Withdrawal Failed**

```
gas estimation failed: execution reverted: You must wait enough blocks to withdraw your profits
```

Solution: Need to wait enough blocks to withdraw profits, this is a normal security mechanism.

**3. RPC Connection Failed**
Ensure:

- RPC endpoint address is correct
- Blockchain node is running
- Network connection is normal

### Debugging Tips

1. Use `--help` parameter to view detailed command usage
2. Check generated JSON file contents
3. Verify keystore file path and password file
4. Confirm chain ID and RPC address are configured correctly

## Technical Architecture

### Project Structure

```
congress-cli/ 
├── cmd/                    # Command implementation
│   ├── proposal.go        # Proposal-related commands
│   ├── tools.go          # Utility functions
│   └── validator.go      # Validator-related commands
├── contracts/            # Contract bindings (symlink)
│   └── generated/       # Auto-generated Go bindings
├── build/               # Build output
│   └── congress-cli    # Executable file
├── Makefile            # Build configuration
├── go.mod              # Go module definition
└── README.md           # This document
```

### Dependencies

- `github.com/ethereum/go-ethereum` - Ethereum client library
- `github.com/spf13/cobra` - CLI framework
- `golang.org/x/crypto` - Cryptographic library

## Contributing

1. Fork the project
2. Create feature branch
3. Commit changes
4. Create Pull Request

## License

This project uses the MIT license.

## Version History

- v1.0.0 - Initial version, supporting basic governance functionality
