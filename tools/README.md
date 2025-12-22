# Congress CLI v1.2.0

Juchain blockchain governance command-line tool for validator management, proposal voting, staking management, and network governance.

## 🚀 New Features (v1.2.0)

- ✅ **Complete Staking Management**: Supports validator registration, delegation, reward withdrawal, and other complete functions
- ✅ **Configuration Management**: Supports configuration query and setting
- ✅ **Enhanced Proposal Query**: Supports querying individual proposals and all proposals
- ✅ **Improved Input Validation**: Complete parameter validation and error prompts
- ✅ **Better Error Handling**: Structured error messages and detailed error information
- ✅ **Enhanced User Experience**: Colorful output and clear status indicators
- ✅ **Global Parameter Validation**: Automatic validation of RPC addresses and chain IDs
- ✅ **Centralized Configuration Management**: Unified constants and configuration management

## Feature Overview

Congress CLI is a command-line tool for Juchain blockchain governance that provides complete validator management, proposal voting, and staking management functions.

### Core Features

- **Proposal Management**: Create validator add/remove proposals and configuration update proposals
- **Voting System**: Vote on proposals (supports simplified voting syntax)
- **Validator Management**: Query validator information, manage earnings, edit validator information
- **Staking Management**: Validator registration, delegation, undelegation, reward withdrawal
- **Transaction Processing**: Sign and send transactions to the blockchain network
- **Configuration Management**: Set and query RPC endpoints and chain IDs
- **Input Validation**: Comprehensive parameter validation and error handling

## Installation and Compilation

### Prerequisites

- Go 1.23.0 or higher
- Solidity compiler (solc 0.8.20)
- abigen tool (for generating Go bindings)

### Compilation Steps

```bash
# Enter project directory
cd sys-contract/congress-cli

# Compile contracts and generate Go bindings
make proposal

# Compile executable
make build

# The generated executable is located at build/congress-cli
```

### Makefile Targets

- `make build` - Compile the entire project
- `make proposal` - Generate Go bindings for the Proposal contract
- `make cleanContract` - Clean up generated contract files
- `make clean` - Clean up build files

## Usage Guide

### Global Parameters

All commands support the following global parameters:

- `-c, --chainId int` - Specify chain ID (Testnet: 202599, Mainnet: 210000)
- `-l, --rpc_laddr string` - Specify RPC endpoint address
  - Testnet: `https://testnet-rpc.juchain.org`
  - Mainnet: `https://rpc.juchain.org`
  - Local: `http://localhost:8545`

⚠️ **Note**: The new version automatically validates these parameters

### Configuration Management

You can use a configuration file to manage default RPC endpoints and chain IDs:

```bash
# Set default RPC endpoint
./build/congress-cli config set --rpc https://testnet-rpc.juchain.org

# Set default chain ID
./build/congress-cli config set --chain-id 202599

# View current configuration
./build/congress-cli config list

# Get specific configuration item
./build/congress-cli config get --rpc
./build/congress-cli config get --chain-id
```

### Quick Start

1. **View help and examples**:

```bash
./build/congress-cli --help
./build/congress-cli [command] --help  # View specific command help
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

---

## I. Proposal Management

### 1.1 Create Validator Add/Remove Proposal

**Create add validator proposal:**

```bash
# Testnet example
./build/congress-cli create_proposal \
  -p 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  -t 0x029DAB47e268575D4AC167De64052FB228B5fA41 \
  -o add \
  -c 202599 \
  -l https://testnet-rpc.juchain.org

# Mainnet example
./build/congress-cli create_proposal \
  -p 0xccafa71c31bc11ba24d526fd27ba57d743152807 \
  -t 0x4d432df142823ca25b21bc3f9744ed21a275bdea \
  -o add \
  -c 210000 \
  -l https://rpc.juchain.org
```

**Create remove validator proposal:**

```bash
./build/congress-cli create_proposal \
  -p 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  -t 0x029DAB47e268575D4AC167De64052FB228B5fA41 \
  -o remove \
  -c 202599 \
  -l https://testnet-rpc.juchain.org
```

**Parameter Description:**
- `-p, --proposer` - Proposer address (must be a valid validator)
- `-t, --target` - Target address (validator to add or remove)
- `-o, --operation` - Operation type (`add` or `remove`)

**Output File**: `createProposal.json`

### 1.2 Create Configuration Update Proposal

**Supported Configuration Items:**

| CID | Configuration Item | Description | Value Range |
|-----|-------------------|-------------|-------------|
| 0 | proposalLastingPeriod | Proposal duration (seconds) | 3600 - 2592000 (1 hour - 30 days) |
| 1 | punishThreshold | Punishment threshold (blocks) | > 0 |
| 2 | removeThreshold | Removal threshold (blocks) | > 0 |
| 3 | decreaseRate | Decrease rate | > 0 |
| 4 | withdrawProfitPeriod | Profit withdrawal period (blocks) | > 0 |
| 5 | blockReward | Block reward (wei) | > 0 |

**Create configuration update proposal:**

```bash
# Modify proposal duration to 86400 seconds (1 day)
./build/congress-cli create_config_proposal \
  -p 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  -i 0 \
  -v 86400 \
  -c 202599 \
  -l https://testnet-rpc.juchain.org

# Modify block reward to 0.833 ether
./build/congress-cli create_config_proposal \
  -p 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  -i 5 \
  -v 833000000000000000 \
  -c 202599 \
  -l https://testnet-rpc.juchain.org
```

**Parameter Description:**
- `-p, --proposer` - Proposer address
- `-i, --cid` - Configuration ID (0-5)
- `-v, --value` - New configuration value

**Output File**: `createUpdateConfigProposal.json`

### 1.3 Vote on Proposal

⚠️ **Important**: Voting syntax has been optimized! Use the `-a` flag for approval, omit for rejection.

```bash
# Approval vote (Testnet example)
./build/congress-cli vote_proposal \
  -s 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  -i b2be7f3cc702c7a24962df6aed188edbcfdebe20fd55f5670efaedace0e4bcdf \
  -a \
  -c 202599 \
  -l https://testnet-rpc.juchain.org

# Rejection vote (omit -a parameter)
./build/congress-cli vote_proposal \
  -s 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  -i b2be7f3cc702c7a24962df6aed188edbcfdebe20fd55f5670efaedace0e4bcdf \
  -c 202599 \
  -l https://testnet-rpc.juchain.org
```

**Parameter Description:**
- `-s, --signer` - Signer address (must be a valid validator)
- `-i, --proposalId` - Proposal ID (64-character hexadecimal string)
- `-a, --approve` - Approval flag (use `-a` for approval, omit for rejection)

**Output File**: `voteProposal.json`

### 1.4 Query Proposals

**Query a single proposal:**

```bash
./build/congress-cli proposal \
  -i b2be7f3cc702c7a24962df6aed188edbcfdebe20fd55f5670efaedace0e4bcdf \
  -l https://testnet-rpc.juchain.org
```

**Query all proposals:**

```bash
./build/congress-cli proposals \
  -l https://testnet-rpc.juchain.org
```

**Output Example:**

```text
📋 Proposal Details:
Proposal ID: b2be7f3cc702c7a24962df6aed188edbcfdebe20fd55f5670efaedace0e4bcdf
Proposer: 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b (Validator Address)
Target Address: 0x029DAB47e268575D4AC167De64052FB228B5fA41 (Validator to Add)
Action: Add New Validator (Flag: true)
Proposal Type: 1 (Validator Management)
Create Time: 2025-01-21 10:30:00 UTC
Status: ✅ Passed
Votes: 👍 3 agree, 👎 0 reject
```

---

## II. Validator Management

### 2.1 Query Validator Information

**Query all validators:**

```bash
./build/congress-cli miners \
  -l https://testnet-rpc.juchain.org
```

**Query specific validator:**

```bash
./build/congress-cli miner \
  -a 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  -l https://testnet-rpc.juchain.org
```

**Output Example:**

```text
Address: 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b
Fee Address: 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b
Status: Active ✅
Accumulated Rewards: 5035784561530401884829
Penalized Rewards: 5323260025816819260865
Last Withdraw Block: 1206974
```

**Status Explanation:**
- Status 1 = Active
- Status 2 = Inactive

### 2.2 Withdraw Validator Profits

```bash
# Testnet example
./build/congress-cli withdraw_profits \
  -a 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  -c 202599 \
  -l https://testnet-rpc.juchain.org

# Mainnet example
./build/congress-cli withdraw_profits \
  -a 0xccafa71c31bc11ba24d526fd27ba57d743152807 \
  -c 210000 \
  -l https://rpc.juchain.org
```

⚠️ **Note**:
- Profit withdrawal has a minimum waiting block limit (controlled by the `withdrawProfitPeriod` configuration)
- Profit withdrawal does not require a voting process, validators can directly withdraw their own profits
- Need to wait for enough blocks before withdrawing

**Output File**: `withdrawProfits.json`

### 2.3 Edit Validator Information

```bash
./build/congress-cli staking edit-validator \
  --validator 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  --fee-addr 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  --moniker "My Validator" \
  --identity "keybase_identity" \
  --website "https://validator.example.com" \
  --email "validator@example.com" \
  --details "Professional validator node" \
  -l https://testnet-rpc.juchain.org
```

**Parameter Description:**
- `--validator` - Validator address (required)
- `--fee-addr` - Fee receiving address (required)
- `--moniker` - Validator display name (optional)
- `--identity` - Validator identity identifier (optional, Keybase signature)
- `--website` - Validator website (optional)
- `--email` - Validator contact email (optional)
- `--details` - Validator description (optional)

**Output File**: `editValidator.json`

---

## III. Staking Management

### 3.1 Validator Registration

**Prerequisites:**
1. Validator must pass proposal (`pass[validator] = true`)
2. Registration must be completed within 7 days after proposal passes
3. Account must have sufficient balance (at least 10,000 JU + gas fees)

**Register Validator:

```bash
./build/congress-cli staking register-validator \
  --proposer 0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc \
  --stake-amount 10000 \
  --commission-rate 500 \
  -c 202599 \
  -l https://testnet-rpc.juchain.org
```

**Parameter Description:**
- `--proposer` - Validator account address (required, must be the target address that passed the proposal)
- `--stake-amount` - Staking amount (required, minimum 10,000 JU)
- `--commission-rate` - Commission rate in basis points (required, 500 = 5%, range 0-10000)

**Important Notes**:
- ⚠️ Registration must be completed within **7 days** after proposal passes, otherwise a new proposal is required
- ⚠️ Account must have sufficient balance (at least 10,000 JU + gas fees) at registration
- ⚠️ Need to wait for the next Epoch (approximately 24 hours) after registration to start producing blocks

**Output File**: `registerValidator.json`

### 3.2 Delegate Tokens

Delegate tokens to trusted validators to earn staking rewards:

```bash
./build/congress-cli staking delegate \
  --delegator 0x970e8128ab834e3eac664312d6e30df9e93cb357 \
  --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  --amount 1000 \
  -l https://testnet-rpc.juchain.org
```

**Parameter Description:**
- `--delegator` - Delegator account address (required)
- `--validator` - Target validator address (required)
- `--amount` - Delegation amount (required, minimum 1 JU)

**Output File**: `delegate.json`

### 3.3 Undelegate Tokens

Start a 7-day unbonding period during which tokens cannot be transferred:

```bash
./build/congress-cli staking undelegate \
  --delegator 0x970e8128ab834e3eac664312d6e30df9e93cb357 \
  --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  --amount 500 \
  -l https://testnet-rpc.juchain.org
```

**Parameter Description:**
- `--delegator` - Delegator account address (required)
- `--validator` - Target validator address (required)
- `--amount` - Unbonding amount (required)

**Note**:
- Unbonding period: 7 days (604,800 blocks)
- During unbonding, tokens still count toward validator's total stake but cannot be transferred
- After unbonding is complete, tokens can be withdrawn (using `withdrawUnbonded`, requires direct contract call)

**Output File**: `undelegate.json`

### 3.4 Claim Rewards

**Validator claims commission and validator share:

```bash
./build/congress-cli staking claim-rewards \
  --claimer 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  -c 202599 \
  -l https://testnet-rpc.juchain.org
```

**Delegator claims delegation rewards:

```bash
./build/congress-cli staking claim-rewards \
  --claimer 0x970e8128ab834e3eac664312d6e30df9e93cb357 \
  --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  -c 202599 \
  -l https://testnet-rpc.juchain.org
```

**Parameter Description:**
- `--claimer` - Claiming account address (required)
- `--validator` - Validator address (required)

**Note**: Rewards must be claimed separately for each delegation relationship

**Output File**: `claimRewards.json`

### 3.5 Query Validator Information

Get detailed staking and status information for validators:

```bash
./build/congress-cli staking query-validator \
  --address 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  -l https://testnet-rpc.juchain.org
```

**Output Example:**

```text
✅ Validator Information
Address: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
Self Stake: 10000 JU
Total Delegated: 50000 JU
Total Stake: 60000 JU
Commission Rate: 500 basis points
Is Jailed: false
Jail Until Block: 0
```

### 3.6 Query Delegation Information

Query delegation details between a specific delegator and validator:

```bash
./build/congress-cli staking query-delegation \
  --delegator 0x970e8128ab834e3eac664312d6e30df9e93cb357 \
  --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  -l https://testnet-rpc.juchain.org
```

**Output Example:**

```text
✅ Delegation Information
Delegator: 0x970e8128ab834e3eac664312d6e30df9e93cb357
Validator: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
Delegated Amount: 1000 JU
Pending Rewards: 25 JU
Unbonding Amount: 0 JU
Unbonding Block: 0
```

### 3.7 Query Top Validators

Get a list of validators sorted by total staked amount:

```bash
./build/congress-cli staking list-top-validators \
  --limit 21 \
  -l https://testnet-rpc.juchain.org
```

**Parameter Description:**
- `--limit` - Maximum number of validators to display (optional, default 21, for display purposes only)

**Output Example:**

```text
✅ Top Validators
Total Count: 21
1. 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
2. 0x970e8128ab834e3eac664312d6e30df9e93cb357
3. 0x6e30df9e93cb3578ec64c67c554dddd8d1da2c25
...
```

---

## IV. Transaction Processing

### 4.1 Sign Transactions

**Sign transactions:**

```bash
# Testnet example
./build/congress-cli sign \
  -f createProposal.json \
  -k /path/to/keystore/UTC--xxx \
  -p /path/to/password.txt \
  -c 202599

# Mainnet example
./build/congress-cli sign \
  -f createProposal.json \
  -k /path/to/keystore/miner1.key \
  -p /path/to/password.file \
  -c 210000
```

**Parameter Description:**
- `-f, --file` - Transaction file path (required)
- `-k, --key` - Keystore file path (required)
- `-p, --password` - Password file path (required)
- `-c, --chainId` - Chain ID (required)

**Output File**: `[original filename]_signed.json`

### 4.2 Send Signed Transactions

**Send signed transactions:

```bash
# Testnet example
./build/congress-cli send \
  -f createProposal_signed.json \
  -l https://testnet-rpc.juchain.org

# Mainnet example
./build/congress-cli send \
  -f createProposal_signed.json \
  -l https://rpc.juchain.org
```

**Example output after successful sending:

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

---

## V. Complete Workflow Examples

### 5.1 Adding a New Validator Complete Process

#### Step 1: Query Current Validator Status

```bash
./build/congress-cli miners -l https://testnet-rpc.juchain.org
```

#### Step 2: Create Validator Addition Proposal

```bash
./build/congress-cli create_proposal \
  -p 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  -t 0x029DAB47e268575D4AC167De64052FB228B5fA41 \
  -o add \
  -c 202599 \
  -l https://testnet-rpc.juchain.org
```

#### Step 3: Sign and Send Proposal

```bash
./build/congress-cli sign \
  -f createProposal.json \
  -k miner1.key \
  -p password.file \
  -c 202599

./build/congress-cli send \
  -f createProposal_signed.json \
  -l https://testnet-rpc.juchain.org
```

**Record Proposal ID** (obtained from output):
```
Proposal ID: b2be7f3cc702c7a24962df6aed188edbcfdebe20fd55f5670efaedace0e4bcdf
```

#### Step 4: Multiple Validator Voting

```bash
# miner1 approval vote
./build/congress-cli vote_proposal \
  -s 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  -i b2be7f3cc702c7a24962df6aed188edbcfdebe20fd55f5670efaedace0e4bcdf \
  -a \
  -c 202599 \
  -l https://testnet-rpc.juchain.org
./build/congress-cli sign -f voteProposal.json -k miner1.key -p password.file -c 202599
./build/congress-cli send -f voteProposal_signed.json -l https://testnet-rpc.juchain.org

# miner2 approval vote
./build/congress-cli vote_proposal \
  -s 0x81f7a79a51edba249efa812eb2d5478f696f7558 \
  -i b2be7f3cc702c7a24962df6aed188edbcfdebe20fd55f5670efaedace0e4bcdf \
  -a \
  -c 202599 \
  -l https://testnet-rpc.juchain.org
./build/congress-cli sign -f voteProposal.json -k miner2.key -p password.file -c 202599
./build/congress-cli send -f voteProposal_signed.json -l https://testnet-rpc.juchain.org

# miner3 approval vote (if needed)
./build/congress-cli vote_proposal \
  -s 0x578c39eaf09a4e1abf428c423970b59bb8baf42e \
  -i b2be7f3cc702c7a24962df6aed188edbcfdebe20fd55f5670efaedace0e4bcdf \
  -a \
  -c 202599 \
  -l https://testnet-rpc.juchain.org
./build/congress-cli sign -f voteProposal.json -k miner3.key -p password.file -c 202599
./build/congress-cli send -f voteProposal_signed.json -l https://testnet-rpc.juchain.org
```

#### Step 5: Wait for 7-Day Registration Period

After the proposal passes, validators must complete registration and staking within **7 days**, otherwise the qualification expires.

#### Step 6: Validator Registration and Staking

```bash
./build/congress-cli staking register-validator \
  --proposer 0x029DAB47e268575D4AC167De64052FB228B5fA41 \
  --stake-amount 10000 \
  --commission-rate 500 \
  -c 202599 \
  -l https://testnet-rpc.juchain.org

./build/congress-cli sign \
  -f registerValidator.json \
  -k new_validator.key \
  -p password.file \
  -c 202599

./build/congress-cli send \
  -f registerValidator_signed.json \
  -l https://testnet-rpc.juchain.org
```

#### Step 7: Wait for Next Epoch Update

After registration, need to wait for the next Epoch (approximately 24 hours) before starting to produce blocks.

#### Step 8: Verify Validator Entered Validator Set

```bash
./build/congress-cli staking list-top-validators -l https://testnet-rpc.juchain.org
./build/congress-cli miners -l https://testnet-rpc.juchain.org
```

### 5.2 Updating System Configuration Complete Process

#### Step 1: Create Configuration Update Proposal

```bash
./build/congress-cli create_config_proposal \
  -p 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  -i 5 \
  -v 833000000000000000 \
  -c 202599 \
  -l https://testnet-rpc.juchain.org
```

#### Step 2: Sign and Send Proposal

```bash
./build/congress-cli sign \
  -f createUpdateConfigProposal.json \
  -k miner1.key \
  -p password.file \
  -c 202599

./build/congress-cli send \
  -f createUpdateConfigProposal_signed.json \
  -l https://testnet-rpc.juchain.org
```

#### Step 3: Validator Voting (Same as Adding Validator Process)

#### Step 4: Verify Configuration Update

After the proposal passes, the configuration will be automatically updated.

---

## VI. Configuration Files

### Generated Transaction Files

The tool will generate the following JSON files in the current directory:

**Proposal Related:**
- `createProposal.json` - Original transaction for creating proposal
- `createProposal_signed.json` - Signed transaction for creating proposal
- `createUpdateConfigProposal.json` - Original transaction for creating configuration update proposal
- `createUpdateConfigProposal_signed.json` - Signed transaction for configuration update proposal
- `voteProposal.json` - Original transaction for voting
- `voteProposal_signed.json` - Signed transaction for voting

**Validator Related:**
- `withdrawProfits.json` - Original transaction for profit withdrawal
- `withdrawProfits_signed.json` - Signed transaction for profit withdrawal
- `editValidator.json` - Original transaction for editing validator information
- `editValidator_signed.json` - Signed transaction for editing validator information

**Staking Related:**
- `registerValidator.json` - Original transaction for validator registration
- `registerValidator_signed.json` - Signed transaction for validator registration
- `delegate.json` - Original transaction for delegation
- `delegate_signed.json` - Signed transaction for delegation
- `undelegate.json` - Original transaction for undelegation
- `undelegate_signed.json` - Signed transaction for undelegation
- `claimRewards.json` - Original transaction for claiming rewards
- `claimRewards_signed.json` - Signed transaction for claiming rewards

### Keystore File Format

Uses standard Ethereum keystore format, for example:

```
UTC--202599-08-06T08-30-51.139143000Z--f39fd6e51aad88f6f4ce6ab8827279cffFb92266
```

---

## VII. Troubleshooting

### Common Errors

**1. EIP-155 Error

```
send tx error only replay-protected (EIP-155) transactions allowed over RPC
```

Solution: The correct chain ID must be specified when signing:

```bash
./build/congress-cli sign -f transaction.json -k keystore -p password -c 202599
```

**2. Profit Withdrawal Failure

```
gas estimation failed: execution reverted: You must wait enough blocks to withdraw your profits
```

Solution: Need to wait for enough blocks before withdrawing profits, this is a normal security mechanism.

**3. Validator Registration Failure

```
execution reverted: Proposal expired, must repropose
```

Solution: Registration must be completed within 7 days after the proposal passes, otherwise a new proposal is required.

**4. RPC Connection Failure

Ensure:
- RPC endpoint address is correct
- Blockchain node is running
- Network connection is normal

### Debugging Tips

1. Use the `--help` parameter to view detailed command usage
2. Check the content of generated JSON files
3. Verify keystore file path and password file
4. Confirm chain ID and RPC address configuration is correct
5. Use `config list` to view current configuration

---

## VIII. Technical Architecture

### Project Structure

```
congress-cli/
├── cmd/                    # Command implementations
│   ├── proposal.go        # Proposal-related commands
│   ├── validator.go       # Validator-related commands
│   ├── staking.go         # Staking-related commands
│   ├── config_cmd.go      # Configuration management commands
│   ├── tools.go           # Utility functions
│   └── utils.go           # Utility functions
├── contracts/             # Contract bindings (symbolic links)
│   └── generated/         # Auto-generated Go bindings
├── build/                 # Build output
│   └── congress-cli      # Executable file
├── Makefile              # Build configuration
├── go.mod                # Go module definition
└── README.md             # This document
```

### Dependencies

- `github.com/ethereum/go-ethereum` - Ethereum client library
- `github.com/spf13/cobra` - CLI framework
- `golang.org/x/crypto` - Cryptography library

---

## IX. Contribution Guidelines

1. Fork the project
2. Create a feature branch
3. Commit changes
4. Create a Pull Request

---

## X. License

This project is licensed under the MIT License.

---

## XI. Version History

- **v1.2.0** - Added complete staking management functions, configuration management, enhanced proposal query
- **v1.1.0** - Improved input validation, better error handling, enhanced user experience
- **v1.0.0** - Initial version, supporting basic governance functions

---

**Document Version**: v1.2.0  
**Last Updated**: 2025-01-21  
**Maintainer**: POSA Development Team
