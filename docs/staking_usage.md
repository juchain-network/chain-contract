# JuChain Staking System Usage Guide

## 📋 Overview

JuChain adopts an innovative JPoSA (JuChain Proof of Stake Authority) hybrid consensus mechanism, combining the fast confirmation of PoA with the economic incentives of PoS. This guide details how to use the `congress-cli` tool for staking operations.

### 🎯 Staking System Features

- **🏛️ Dual Contract Architecture**: Validators + Staking contracts work in division of labor
- **💰 Economic Incentives**: Validators and delegators share rewards
- **🔒 Security Mechanisms**: 7-day unbonding period and malicious behavior penalties
- **🎖️ Governance Participation**: Stakers participate in network governance decisions
- **📈 Dynamic Adjustment**: Commission rates and stake amounts adjustable in real-time
- **🛡️ Enhanced Security**: ReentrancyGuard protection, configuration parameter validation
- **⚡ Performance Optimization**: Removed SafeMath, using Solidity 0.8+ built-in operators

# Sign Transactions
./build/congress-cli sign
  --file registerValidator.json
  --key ./keystore/UTC--2024-...
  --password ./password.txt
  --chainId 210000

## 🏗️ System Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                  JuChain Staking System                     │
├─────────────────┬─────────────────┬─────────────────────────┤
│   Validators    │   Delegators    │     Staking Pool        │
│                 │                 │                         │
│ • Self-stake    │ • Delegate      │ • Reward Distribution   │
│ • Commission    │ • Undelegate    │ • Slashing Protection   │
│ • Block Rewards │ • Claim Rewards │ • Governance Voting     │
│ • Fee Address   │ • Multiple Val. │ • Unbonding Queue       │
└─────────────────┴─────────────────┴─────────────────────────┘
```

## ⚙️ Environment Preparation

### Software Requirements

| Component | Minimum Version | Recommended Version | Description |
|------|----------|----------|------|
| **Go** | 1.23+ | 1.24+ | Compile Congress CLI |
| **Git** | 2.30+ | Latest | Version control |
| **curl** | 7.0+ | Latest | RPC call testing |

### Build CLI Tool

```bash
# 📁 Enter project directory
cd sys-contract/congress-cli

# 🔧 Compile tool
go build -o build/congress-cli

# ✅ Verify installation
./build/congress-cli version
./build/congress-cli staking --help
```

### Network Configuration

#### RPC Endpoint Configuration

| Network Environment | RPC Address | Chain ID | Description |
|----------|---------|----------|------|
| **Mainnet** | `https://rpc.juchain.org` | 210000 | Production environment |
| **Testnet** | `https://testnet-rpc.juchain.org` | 202599 | Test environment |
| **Local** | `http://localhost:8545` | 202599 | Development environment (default) |

#### Key Management

```bash
# 🔑 Create keystore
mkdir -p keystore
chmod 700 keystore/

# 🔐 Create password file
echo "your secure password" > password.txt
chmod 600 password.txt

# 📄 Prepare key file path
KEYSTORE_FILE="./keystore/UTC--2024-..."
PASSWORD_FILE="./password.txt"
```

## 🎯 Core Command Overview

JuChain Staking system provides the following core commands:

### 📊 Command Categories

#### 🔐 Validator Operations

- `register-validator`: Register as validator and self-stake
- `claim-rewards`: Withdraw validator rewards

#### 💰 Delegation Operations  

- `delegate`: Delegate tokens to validator
- `undelegate`: Start unbonding delegation (7-day cycle)
- `claim-rewards`: Withdraw delegation rewards

#### 🔍 Query Operations

- `query-validator`: Query validator details
- `query-delegation`: Query delegation details  
- `list-top-validators`: Query top validator list

### 📈 Staking Parameters

| Parameter Name | Minimum Value | Maximum Value | Description |
|----------|--------|--------|------|
| **Validator Self-Stake** | 10,000 JU | Unlimited | Minimum requirement to become validator |
| **Delegation Amount** | 1 JU | Unlimited | Minimum amount for single delegation |
| **Commission Rate** | 0% | 100% | Commission rate validators can set |
| **Unbonding Period** | 7 days | 7 days | Waiting time to undelegate |
| **Maximum Validators** | - | 21 | Number of simultaneously active validators in network |

## 🚀 Validator Operations

### Register Validator

Becoming a network validator requires completing the following steps:

#### Step 1: Create Proposal (Initiated by Existing Validator)

Validators must first go through governance proposal to register. Existing validators need to create an add-validator proposal:

```bash
# 📝 Create validator add proposal
./build/congress-cli create_proposal \
  --rpc_laddr http://localhost:8545 \
  --chainId 202599 \
  --proposer 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  --target 0xNew validator address \
  --operation add

# Sign and send proposal
./build/congress-cli sign -f createProposal.json -k keystore -p password --chainId 202599
./build/congress-cli send -f createProposal_signed.json
```

#### Step 2: Validator Voting

Existing validators need to vote on the proposal (requires majority approval):

```bash
# 🗳️ Validator voting (approve)
./build/congress-cli vote_proposal \
  --rpc_laddr http://localhost:8545 \
  --chainId 202599 \
  --signer 0xValidator address \
  --proposalId Proposal ID \
  --approve

# Sign and send vote
./build/congress-cli sign -f voteProposal.json -k keystore -p password --chainId 202599
./build/congress-cli send -f voteProposal_signed.json
```

#### Step 3: Wait 7-Day Registration Period

After proposal passes, validators must complete registration staking within **7 days**, otherwise the qualification expires.

#### Step 4: Register and Stake

```bash
# 📝 Create validator registration transaction (must be within 7 days after proposal passes)
./build/congress-cli staking register-validator \
  --rpc_laddr http://localhost:8545 \
  --chainId 202599 \
  --proposer 0xNew validator address \
  --stake-amount 10000 \
  --commission-rate 500

# Parameter explanation:
# --proposer: Validator account address (required, must be the target address of the passed proposal)
# --stake-amount: Staking amount (minimum 10,000 JU)
# --commission-rate: Commission rate, calculated in basis points (500 = 5%)

# Sign and send
./build/congress-cli sign -f registerValidator.json -k keystore -p password --chainId 202599
./build/congress-cli send -f registerValidator_signed.json
```

**Output File**: `registerValidator.json`

**Important Notes**:
- ⚠️ Must complete registration within **7 days** after proposal passes, otherwise need to re-propose
- ⚠️ Account must have sufficient balance at registration (at least 10,000 JU + Gas fees)
- ⚠️ After registration, need to wait for next Epoch (approximately 24 hours) to start producing blocks

#### Commission Rate Setting Guidelines

| Commission Rate | Basis Points Value | Applicable Scenarios | Competitiveness |
|--------|--------|----------|--------|
| **0-2%** | 0-200 | New validators attracting delegations | ⭐⭐⭐⭐⭐ |
| **3-5%** | 300-500 | Balanced rewards and competition | ⭐⭐⭐⭐ |
| **6-10%** | 600-1000 | Mature validators | ⭐⭐⭐ |
| **10%+** | 1000+ | High-quality services | ⭐⭐ |

### Validator Reward Withdrawal

```bash
# 💰 Withdraw validator rewards
./build/congress-cli staking claim-rewards \
  --rpc_laddr http://localhost:8545 \
  --claimer 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266

# Parameter explanation:
# --claimer: Withdrawal account address (usually same as validator address)
# --validator: Validator address
```

**Output File**: `claimRewards.json`

## 💎 Delegation Operations

### Delegate Tokens

Delegate tokens to trusted validators to earn staking rewards:

```bash
# 🤝 Delegate tokens to validator
./build/congress-cli staking delegate \
  --rpc_laddr http://localhost:8545 \
  --delegator 0x970e8128ab834e3eac664312d6e30df9e93cb357 \
  --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  --amount 1000

# Parameter explanation:
# --delegator: Delegator account address (required)
# --validator: Target validator address (required)
# --amount: Delegation amount (minimum 1 JU)
```

**Output File**: `delegate.json`

#### Delegation Strategy Recommendations

```bash
# 🎯 Diversification delegation strategy
echo "=== Recommended Delegation Strategy ==="
echo "1. Risk diversification: Delegate to 3-5 different validators"
echo "2. Performance priority: Choose validators with high uptime"
echo "3. Commission balance: Consider commission rate and service quality"
echo "4. Governance participation: Support validators actively participating in governance"
```

### Undelegate

Start 7-day unbonding cycle, during which tokens cannot be transferred:

```bash
# 📤 Undelegate (start 7-day unbonding period)
./build/congress-cli staking undelegate \
  --rpc_laddr http://localhost:8545 \
  --delegator 0x970e8128ab834e3eac664312d6e30df9e93cb357 \
  --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  --amount 500

# Parameter explanation:
# --delegator: Delegator account address (required)
# --validator: Target validator address (required)  
# --amount: Unbonding amount
```

**Output File**: `undelegate.json`

#### Unbonding Time Calculation

```bash
# ⏰ Unbonding period explanation
echo "Unbonding cycle: Default 7 days (604,800 blocks), configurable via governance proposal"
echo "Block time: 1 second/block"
echo "Unbonding start: Begins immediately after transaction confirmation"
echo "Fund availability: Can be withdrawn after unbonding period ends (using withdrawUnbonded)"
echo ""
echo "Note: During unbonding period, tokens still count toward validator's total stake but cannot be transferred"
echo "Note: Unbonding period can be modified via proposal (cid = 6)"
```

### Delegator Reward Withdrawal

```bash
# 🎁 Withdraw delegation rewards
./build/congress-cli staking claim-rewards \
  --rpc_laddr http://localhost:8545 \
  --claimer 0x970e8128ab834e3eac664312d6e30df9e93cb357 \
  --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266

# Note: Each delegation relationship needs separate reward withdrawal
```

## 🔍 Query Commands

### Validator Information Query

Get detailed staking and status information for validators:

```bash
# 📊 Query validator details
./build/congress-cli staking query-validator \
  --rpc_laddr http://localhost:8545 \
  --address 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
```

#### Validator Information Output Format

```text
✅ Validator Information
Address: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
Self-Stake: 10,000 JU
Total Delegated: 50,000 JU  
Total Stake: 60,000 JU
Commission Rate: 500 basis points (5%)
Is Jailed: false
Jailed Until Block: 0
Uptime: 99.8%
Last Block: #123456
```

### Delegation Information Query

Query delegation details between specific delegator and validator:

```bash
# 🔍 Query delegation details
./build/congress-cli staking query-delegation \
  --rpc_laddr http://localhost:8545 \
  --delegator 0x970e8128ab834e3eac664312d6e30df9e93cb357 \
  --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
```

#### Delegation Information Output Format

```text
✅ Delegation Information
Delegator: 0x970e8128ab834e3eac664312d6e30df9e93cb357
Validator: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
Delegation Amount: 1,000 JU
Pending Rewards: 25 JU
Unbonding Amount: 0 JU
Unbonding Completion Block: 0
Delegation Time: 2024-08-27 10:30:00
Annual Yield: 8.5%
```

### Top Validator Query

Get validator list sorted by total stake:

```bash
# 🏆 Query top validators
./build/congress-cli staking list-top-validators \
  --rpc_laddr http://localhost:8545 \
  --limit 21
```

#### Top Validator Output Format

```text
✅ Top Validators (sorted by stake)
Total: 21 active validators

Rank | Validator Address                              | Total Stake | Commission | Status
-----|-----------------------------------------------|-------------|------------|-------
1    | 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266     | 60,000 JU   | 5%         | Active
2    | 0x970e8128ab834e3eac664312d6e30df9e93cb357     | 55,000 JU   | 3%         | Active  
3    | 0x6e30df9e93cb3578ec64c67c554dddd8d1da2c25     | 50,000 JU   | 7%         | Active
...
21   | 0x3d968443d9b72bcef4409b3a2d5e31031390fc82     | 15,000 JU   | 10%        | Active
```

## 🔄 Transaction Execution Process

All staking transactions follow the standard three-step process:

### Step 1: Create Transaction

```bash
# 📝 Create unsigned transaction
./build/congress-cli staking register-validator \
  --proposer 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  --stake-amount 10000 \
  --commission-rate 500

echo "✅ Transaction file generated: registerValidator.json"
```

### Step 2: Sign Transaction

```bash
# ✍️ Sign transaction with private key
./build/congress-cli sign \
  --file registerValidator.json \
  --key ./keystore/UTC--2024-... \
  --password ./password.txt \
  --chainId 202599

echo "✅ Signed file generated: registerValidator_signed.json"
```

### Step 3: Broadcast Transaction

```bash
# 📡 Broadcast transaction to network
./build/congress-cli send \
  --file registerValidator_signed.json \
  --rpc_laddr http://localhost:8545

echo "✅ Transaction broadcast, waiting for block confirmation..."
```

### Transaction Status Verification

```bash
# 🔍 Verify transaction results
echo "Check transaction hash: 0x..."
echo "Verify staking status:"
./build/congress-cli staking query-validator \
  --address 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  --rpc_laddr http://localhost:8545
```

## ⚙️ Advanced Configuration

### Gas Fee Optimization

```bash
# 💰 Custom Gas settings (modify JSON file)
echo "Default Gas configuration:"
echo "  gasLimit: Auto-estimated + 20% buffer"
echo "  gasPrice: 20 Gwei"
echo ""
echo "High-priority transactions:"
echo "  gasPrice: 50 Gwei"
echo "  gasLimit: Manually set higher value"
```

### Batch Operations

```bash
# 🔄 Batch delegation example
VALIDATORS=(
  "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
  "0x970e8128ab834e3eac664312d6e30df9e93cb357"
  "0x6e30df9e93cb3578ec64c67c554dddd8d1da2c25"
)

DELEGATOR="0x3858ffca201b0a7d75fd23bb302c12332c5e4000"
AMOUNT=1000

for validator in "${VALIDATORS[@]}"; do
  echo "Delegating $AMOUNT JU to validator: $validator"
  
  # Create delegation transaction
  ./build/congress-cli staking delegate \
    --delegator $DELEGATOR \
    --validator $validator \
    --amount $AMOUNT \
    --rpc_laddr http://localhost:8545
  
  # Rename file to avoid overwrite
  mv delegate.json delegate_${validator:2:8}.json
  
  echo "Transaction file created: delegate_${validator:2:8}.json"
done

echo "✅ All delegation transactions created, please sign and broadcast sequentially"
```

### Automation Scripts

```bash
#!/bin/bash
# auto-stake.sh - Automated staking script

set -e

# Configuration parameters
RPC_URL="http://localhost:8545"
KEYSTORE="./keystore/UTC--2024-..."
PASSWORD="./password.txt"
CHAIN_ID="202599"  # Testnet, use 210000 for mainnet
DELEGATOR="0x3858ffca201b0a7d75fd23bb302c12332c5e4000"

# Function: Execute complete staking process
execute_staking() {
    local operation=$1
    local validator=$2
    local amount=$3
    
    echo "🚀 Executing $operation operation"
    echo "Validator: $validator"
    echo "Amount: $amount JU"
    
    # Create transaction
    ./build/congress-cli staking $operation \
        --delegator $DELEGATOR \
        --validator $validator \
        --amount $amount \
        --rpc_laddr $RPC_URL
    
    # Sign transaction  
    ./build/congress-cli sign \
        --file $operation.json \
        --key $KEYSTORE \
        --password $PASSWORD \
        --chainId $CHAIN_ID
    
    # Broadcast transaction
    ./build/congress-cli send \
        --file ${operation}_signed.json \
        --rpc_laddr $RPC_URL
    
    echo "✅ $operation operation completed"
    echo ""
}

# Example: Delegate to multiple validators
echo "=== Starting batch delegation ==="
execute_staking "delegate" "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266" 1000
execute_staking "delegate" "0x970e8128ab834e3eac664312d6e30df9e93cb357" 1000

echo "=== Batch delegation completed ==="
```

## Configuration

### RPC Endpoints

Common JuChain RPC endpoints:

- **Mainnet**: `https://rpc.juchain.org` (Chain ID: 210000)
- **Testnet**: `https://testnet-rpc.juchain.org` (Chain ID: 202599)
- **Local**: `http://localhost:8545` (Chain ID: 202599)

### Chain IDs

- **Mainnet**: `210000`
- **Testnet**: `202599`
- **Local**: `202599` (default, customizable)

### Gas Configuration

The CLI automatically estimates gas with a 20% buffer. For custom gas settings, modify the transaction JSON before signing.

## Best Practices

### Security

1. **Keystore Safety**: Store keystore files securely and never share passwords
2. **Amount Verification**: Double-check stake amounts before signing
3. **Address Validation**: Verify all addresses are correct before transactions
4. **Reentrancy Protection**: All critical functions are protected by `ReentrancyGuard`
5. **Parameter Validation**: All configuration parameters have range validation to prevent errors

### Staking Strategy

1. **Validator Selection**: Research validators' performance and commission rates
2. **Diversification**: Consider delegating to multiple validators
3. **Reward Timing**: Claim rewards regularly to compound earnings
4. **Unbonding Period**: Plan for the 7-day unbonding period when undelegating

### Monitoring

1. **Regular Queries**: Monitor validator and delegation status regularly
2. **Reward Tracking**: Keep track of accumulated rewards
3. **Network Health**: Monitor network performance and validator uptime

## 🚨 Troubleshooting

### Common Error Handling

Resolve typical issues and error codes in staking operations.

#### RPC Connection Failure

```text
❌ Error: invalid RPC URL format: localhost:8545
Cause: Node not running or incorrect RPC port
```

**Solution:**

```bash
# Check node status
ps aux | grep geth
netstat -tulpn | grep :8545

# Start node
./build/bin/geth --config config-validator1.toml --mine

# Verify connection
curl -X POST -H "Content-Type: application/json" \
  --data '{"jsonrpc":"2.0","method":"net_version","params":[],"id":1}' \
  http://localhost:8545
```

#### Invalid Address Format

```text
❌ Error: invalid address format
Cause: Address does not conform to Ethereum format specification
```

**Solution:**

```bash
# Verify address format (must be 42 characters, starting with 0x)
echo "Correct format: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
echo "Incorrect format: f39Fd6e51aad88F6F4ce6aB8827279cffFb92266"

# Check address checksum
./build/congress-cli utils checksum-address \
  --address 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
```

#### Insufficient Stake Amount

```text
❌ Error: stake amount must be at least 10000 JU
Cause: Staking amount below minimum requirement
```

**Solution:**

```bash
echo "Minimum staking requirements:"
echo "  Validator registration: 10,000 JU"
echo "  Delegation staking: 1 JU"
echo "  Increase staking: 1 JU"

# Check account balance
./build/congress-cli utils get-balance \
  --address 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  --rpc_laddr http://localhost:8545
```

#### Commission Rate Out of Range

```text
❌ Error: commission rate must be between 0 and 10000 (100%)
Cause: Commission rate exceeds 10000 basis points (100%)
```

**Solution:**

```bash
echo "Commission rate setting guide:"
echo "  Minimum: 0 basis points (0%)"
echo "  Maximum: 10000 basis points (100%)"
echo "  Recommended: 100-1000 basis points (1%-10%)"
echo "  Formula: Percentage × 100 = Basis Points"
echo ""
echo "Examples:"
echo "  1% = 100 basis points"
echo "  5% = 500 basis points"
echo "  10% = 1000 basis points"
```

#### Insufficient Transaction Gas Fee

```text
❌ Error: intrinsic gas too low
Cause: Gas fee setting too low
```

**Solution:**

```bash
# View current gas price
echo "Recommended Gas settings:"
echo "  gasPrice: 20 Gwei (normal)"
echo "  gasPrice: 50 Gwei (fast)"
echo "  gasLimit: Automatically estimated by system"
echo ""
echo "Manual Gas setting (modify JSON file):"
echo '{
  "gasPrice": "0x12a05f200",
  "gasLimit": "0x5208"
}'
```

### 🔧 Diagnostic Tools

#### Network Status Check

```bash
#!/bin/bash
# network-health.sh

echo "=== JuChain Network Health Check ==="

# Check RPC connection
check_rpc() {
    echo "🔍 Checking RPC connection..."
    response=$(curl -s -X POST -H "Content-Type: application/json" \
        --data '{"jsonrpc":"2.0","method":"net_version","params":[],"id":1}' \
        http://localhost:8545)
    
    if [[ $response == *"210000"* ]]; then
        echo "✅ RPC connection normal - Mainnet"
    elif [[ $response == *"202599"* ]]; then
        echo "✅ RPC connection normal - Testnet/Localnet"
    else
        echo "❌ RPC connection abnormal"
    fi
}

# Check sync status
check_sync() {
    echo "🔄 Checking sync status..."
    sync_status=$(curl -s -X POST -H "Content-Type: application/json" \
        --data '{"jsonrpc":"2.0","method":"eth_syncing","params":[],"id":1}' \
        http://localhost:8545)
    
    if [[ $sync_status == *"false"* ]]; then
        echo "✅ Node fully synchronized"
    else
        echo "⏳ Node synchronizing..."
    fi
}

# Check validator count
check_validators() {
    echo "👥 Checking active validators..."
    ./build/congress-cli staking list-top-validators \
        --limit 100 \
        --rpc_laddr http://localhost:8545 | \
        grep "Count:" || echo "❌ Unable to get validator information"
}

# Execute checks
check_rpc
check_sync
check_validators
echo "=== Check completed ==="
```

#### Staking Status Monitoring

```bash
#!/bin/bash
# stake-monitor.sh

VALIDATOR="0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
RPC_URL="http://localhost:8545"

echo "=== Staking Status Monitoring ==="
echo "Validator: $VALIDATOR"
echo ""

while true; do
    echo "$(date): Checking staking status..."
    
    # Query validator information
    ./build/congress-cli staking query-validator \
        --address $VALIDATOR \
        --rpc_laddr $RPC_URL | \
        grep -E "(Total Stake|Is Jailed|Commission Rate)"
    
    echo "---"
    sleep 30
done
```

## 📚 Appendix

### System Contract Addresses

| Contract Name | Address | Function Description |
|---------|------|----------|
| Validators | `0x000000000000000000000000000000000000f000` | Validator management and governance |
| Punish | `0x000000000000000000000000000000000000f001` | Punishment mechanism and imprisonment |
| Proposal | `0x000000000000000000000000000000000000f002` | Governance proposal voting |
| Staking | `0x000000000000000000000000000000000000f003` | Staking and delegation management |

### Network Parameters

| Parameter Name | Mainnet Value | Testnet Value | Description |
|---------|--------|----------|------|
| Chain ID | 210000 | 202599 | Network identifier |
| RPC Endpoint | `https://rpc.juchain.org` | `https://testnet-rpc.juchain.org` | Network access point |
| Block Time | 1 second | 1 second | Block generation interval |
| Epoch Cycle | 86400 blocks | 86400 blocks | ~24-hour rotation |
| Minimum Stake | 10,000 JU | 10,000 JU | Validator registration requirement |
| Minimum Delegation | 1 JU | 1 JU | Minimum delegation amount |
| Unbonding Period | 604800 blocks | 604800 blocks | 7-day lock period |
| Registration Period | 7 days | 7 days | Must register within this period after proposal passes |

### Useful Links

- **Official Documentation**: [https://juchain.org/docs](https://juchain.org/docs)
- **Block Explorer**: [https://juchain.org/explorer](https://juchain.org/explorer)
- **GitHub Repository**: [https://github.com/JuChain/go-juchain](https://github.com/JuChain/go-juchain)
- **Community Forum**: [https://forum.juchain.org](https://forum.juchain.org)
- **Technical Support**: [support@juchain.org](mailto:support@juchain.org)

---

**Version**: v1.2.0  
**Update Time**: January 21, 2025  
**Scope**: JuChain mainnet and testnet

**Update Content (v1.2.0):**
- Updated security mechanism description: Added ReentrancyGuard protection description
- Updated configuration parameters: Removed inflation-related configurations (cid 5 and 6)
- Updated technical details: All contracts use Solidity 0.8+ built-in operators (SafeMath removed)

**Update Content (v1.1.0):**
- Corrected contract address format (using correct `0x0000...f000` format)
- Unified Chain ID (mainnet 210000, testnet 202599)
- Updated validator registration process, added proposal prerequisite step description
- Added important note about 7-day registration period
- Corrected unbonding period description (604800 blocks = 7 days)
- Updated RPC endpoint addresses

*This document will be continuously updated, please refer to the latest version for accurate information.*

### File Not Found Errors

If transaction files are missing, ensure:

1. The transaction creation command completed successfully
2. You're in the correct directory
3. The file wasn't moved or deleted

### Network Issues

For network connectivity problems:

1. Verify RPC endpoint is accessible
2. Check firewall settings
3. Ensure correct chain ID
4. Confirm network is operational

## Advanced Usage

### Batch Operations

Create multiple transactions and sign them together:

```bash
# Create multiple delegation transactions
./congress-cli staking delegate --delegator 0x... --validator 0x...1 --amount 1000
./congress-cli staking delegate --delegator 0x... --validator 0x...2 --amount 1000

# Sign all transactions
./congress-cli sign --file delegate.json --key keystore.json --password password.txt
./congress-cli sign --file delegate2.json --key keystore.json --password password.txt

# Broadcast sequentially
./congress-cli send --file delegate_signed.json 
./congress-cli send --file delegate2_signed.json
```

### Scripting

Automate staking operations with shell scripts:

```bash
#!/bin/bash
RPC="http://localhost:8545"
KEYSTORE="./keystore.json"
PASSWORD="./password.txt"
CHAIN_ID="202599"

# Function to stake with a validator
stake_with_validator() {
    local validator=$1
    local amount=$2
    
    echo "Delegating $amount JU to $validator"
    
    # Create transaction
    ./congress-cli staking delegate \
        --rpc_laddr $RPC \
        --delegator 0x1234567890123456789012345678901234567890 \
        --validator $validator \
        --amount $amount
    
    # Sign transaction
    ./congress-cli sign \
        --file delegate.json \
        --key $KEYSTORE \
        --password $PASSWORD \
        --chainId $CHAIN_ID
    
    # Broadcast transaction
    ./congress-cli send \
        --file delegate_signed.json \
        --rpc_laddr $RPC
}

# Delegate to multiple validators
stake_with_validator "0x0987654321098765432109876543210987654321" 1000
stake_with_validator "0x1111111111111111111111111111111111111111" 1000
```

## Support

For additional help:

1. Check the main README in `contracts/README.md`
2. Review the command help: `./congress-cli staking [command] --help`
3. Examine transaction files for debugging
4. Verify network status and RPC connectivity