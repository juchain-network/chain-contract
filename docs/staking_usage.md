# JuChain Staking Usage Guide

## 📋 Overview

JuChain adopts the innovative JPoSA (JuChain Proof of Stake Authority) hybrid consensus mechanism, combining PoA's fast finality with PoS economic incentives. This guide details how to use the `congress-cli` tool for staking operations.

### 🎯 Staking System Features

- **🏛️ Dual Contract Architecture**: Validators + Staking contracts work in coordination
- **💰 Economic Incentives**: Validators and delegators share rewards
- **🔒 Security Mechanisms**: 7-day unbonding period and slashing for misbehavior
- **🎖️ Governance Participation**: Stakers participate in network governance decisions
- **📈 Dynamic Adjustment**: Commission rates and stake amounts adjustable in real-time

## Environment Setup

### Software Requirements

| Component | Minimum Version | Recommended | Description |
|-----------|----------------|-------------|-------------|
| **Go** | 1.23+ | 1.24+ | Build Congress CLI |
| **Git** | 2.30+ | Latest | Version control |
| **curl** | 7.0+ | Latest | RPC call testing |

### Build CLI Tools

```bash
# Navigate to project directory
cd sys-contract/congress-cli

# Compile tool
go build -o build/congress-cli

# Verify installation
./build/congress-cli version
./build/congress-cli staking --help
```

### Network Configuration

#### RPC Endpoint Configuration

| Network | RPC Address | Chain ID | Description |
|---------|-------------|----------|-------------|
| **Mainnet** | `https://rpc.juchain.io` | 202599 | Production environment |
| **Testnet** | `https://testnet-rpc.juchain.io` | 202583 | Test environment |
| **Local** | `http://localhost:8545` | Custom | Development environment |

#### Key Management

```bash
# Create keystore
mkdir -p keystore
chmod 700 keystore/

# Create password file
echo "your-secure-password" > password.txt
chmod 600 password.txt

# Prepare keystore file path
KEYSTORE_FILE="./keystore/UTC--2024-..."
PASSWORD_FILE="./password.txt"
```

## Core Commands Overview

### Command Categories

#### Validator Operations

- `register-validator`: Register as validator with self-stake
- `claim-rewards`: Claim validator rewards

#### Delegation Operations

- `delegate`: Delegate tokens to validators
- `undelegate`: Start unbonding delegation (7-day period)
- `claim-rewards`: Claim delegation rewards

#### Query Operations

- `query-validator`: Query validator details
- `query-delegation`: Query delegation details
- `list-top-validators`: Query top validators list

## Validator Operations

### Register Validator

To become a network validator, stake at least 10,000 JU tokens:

```bash
# Create validator registration transaction
./build/congress-cli staking register-validator \
  --rpc_laddr http://localhost:8545 \
  --proposer 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  --stake-amount 10000 \
  --commission-rate 500

# Parameter description:
# --proposer: Validator account address (required)
# --stake-amount: Stake amount (minimum 10,000 JU)
# --commission-rate: Commission rate in basis points (500 = 5%)
```

**Output file**: `registerValidator.json`

### Claim Validator Rewards

```bash
# Claim validator rewards
./build/congress-cli staking claim-rewards \
  --rpc_laddr http://localhost:8545 \
  --claimer 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266

# Parameter description:
# --claimer: Claim account address (usually same as validator)
# --validator: Validator address
```

## Transaction Flow

All staking transactions follow the standard three-step process:

### Step 1: Create Transaction

```bash
# Create unsigned transaction
./build/congress-cli staking register-validator \
  --proposer 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  --stake-amount 10000 \
  --commission-rate 500

echo "Transaction file generated: registerValidator.json"
```

### Step 2: Sign Transaction

```bash
# Sign transaction with private key
./build/congress-cli sign \
  --file registerValidator.json \
  --key ./keystore/UTC--2024-... \
  --password ./password.txt \
  --chainId 210000

echo "Signed file generated: registerValidator_signed.json"
```

### Step 3: Broadcast Transaction

```bash
# Broadcast transaction to network
./build/congress-cli send \
  --file registerValidator_signed.json \
  --rpc_laddr http://localhost:8545

echo "Transaction broadcasted, waiting for block confirmation..."
```

## System Contract Addresses

| Contract Name | Address | Function Description |
|---------------|---------|----------------------|
| Validators | 0xf000000000000000000000000000000000000000 | Validator management and governance |
| Punish | 0xf000000000000000000000000000000000000001 | Punishment mechanism and jailing |
| Proposal | 0xf000000000000000000000000000000000000002 | Governance proposal voting |
| Staking | 0xf000000000000000000000000000000000000003 | Staking and delegation management |

## Network Parameters

| Parameter | Mainnet | Testnet | Description |
|-----------|---------|---------|-------------|
| Chain ID | 210000 | 202599 | Network identifier |
| RPC Endpoint | `https://rpc.juchain.org` | `https://testnet-rpc.juchain.org` | Network access point |
| Block Time | 1 second | 1 second | Block generation interval |
| Epoch Period | 86400 blocks | 86400 blocks | ~24 hour rotation |
| Minimum Stake | 10,000 JU | 10,000 JU | Validator registration requirement |
| Minimum Delegation | 1 JU | 1 JU | Minimum delegation amount |
| Unbonding Period | 518400 blocks | 518400 blocks | ~6 day lock period |

## Best Practices

### Security

1. **Keystore Safety**: Store keystore files securely and never share passwords
2. **Amount Verification**: Double-check stake amounts before signing
3. **Address Validation**: Verify all addresses are correct before transactions

### Staking Strategy

1. **Validator Selection**: Research validators' performance and commission rates
2. **Diversification**: Consider delegating to multiple validators
3. **Reward Timing**: Claim rewards regularly to compound earnings
4. **Unbonding Period**: Plan for the 7-day unbonding period when undelegating

### Monitoring

1. **Regular Queries**: Monitor validator and delegation status regularly
2. **Reward Tracking**: Keep track of accumulated rewards
3. **Network Health**: Monitor network performance and validator uptime

## Troubleshooting

### Common Errors

#### RPC Connection Failed

**Error**: `invalid RPC URL format: localhost:8545`

**Solution**:
```bash
# Check node status
ps aux | grep geth
netstat -tulpn | grep :8545

# Start node
./build/bin/geth --config config-validator1.toml --mine
```

#### Insufficient Stake Amount

**Error**: `stake amount must be at least 10000 JU`

**Solution**:
```bash
echo "Minimum staking requirements:"
echo "  Validator registration: 10,000 JU"
echo "  Delegation: 1 JU"
echo "  Additional stake: 1 JU"
```

## Support

For additional help:

1. Check the main README in `contracts/README.md`
2. Review command help: `./congress-cli staking [command] --help`
3. Examine transaction files for debugging
4. Verify network status and RPC connectivity

---

**Version**: v1.0.0  
**Last Updated**: August 27, 2024  
**Applicable To**: JuChain Mainnet and Testnet

*This document is continuously updated, please follow the latest version for accurate information.*

