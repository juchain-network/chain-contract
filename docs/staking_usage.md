# JuChain Congress-CLI Staking Commands Usage Guide

This guide provides detailed instructions for using the `congress-cli` staking commands in JuChain's JPoSA (Juchain Proof of Stake Authority) consensus system.

## Prerequisites

1. **Go Environment**: Ensure Go 1.19+ is installed
2. **Build CLI Tool**:

   ```bash
   cd sys-contract/congress-cli
   go build
   ```

3. **Network Access**: Access to a JuChain RPC endpoint
4. **Keystore**: Ethereum-compatible keystore file and password

## Command Overview

The staking module provides these commands:

- `register-validator`: Register as a validator with self-stake
- `delegate`: Delegate tokens to a validator  
- `undelegate`: Start unbonding delegation
- `claim-rewards`: Claim staking rewards
- `query-validator`: Query validator information
- `query-delegation`: Query delegation details
- `list-top-validators`: List top validators by stake

## Transaction Commands

### 1. Register Validator

Register as a validator by self-staking at least 10,000 JU tokens.

```bash
./congress-cli staking register-validator \
  --rpc_laddr http://localhost:8545 \
  --proposer 0x1234567890123456789012345678901234567890 \
  --stake-amount 10000 \
  --commission-rate 500
```

**Parameters:**

- `--proposer`: Your validator address (required)
- `--stake-amount`: Amount of JU to stake (minimum 10,000)
- `--commission-rate`: Commission rate in basis points (0-10,000, e.g., 500 = 5%)

**Output:** Creates `registerValidator.json` transaction file

### 2. Delegate Tokens

Delegate tokens to an existing validator to earn staking rewards.

```bash
./congress-cli staking delegate \
  --rpc_laddr http://localhost:8545 \
  --delegator 0x1234567890123456789012345678901234567890 \
  --validator 0x0987654321098765432109876543210987654321 \
  --amount 1000
```

**Parameters:**

- `--delegator`: Your account address (required)
- `--validator`: Target validator address (required)
- `--amount`: Amount of JU to delegate (minimum 1)

**Output:** Creates `delegate.json` transaction file

### 3. Undelegate Tokens

Start the 7-day unbonding process to withdraw delegated tokens.

```bash
./congress-cli staking undelegate \
  --rpc_laddr http://localhost:8545 \
  --delegator 0x1234567890123456789012345678901234567890 \
  --validator 0x0987654321098765432109876543210987654321 \
  --amount 500
```

**Parameters:**

- `--delegator`: Your account address (required)
- `--validator`: Target validator address (required)
- `--amount`: Amount of JU to undelegate

**Output:** Creates `undelegate.json` transaction file

### 4. Claim Rewards

Claim accumulated staking rewards from validation or delegation.

```bash
./congress-cli staking claim-rewards \
  --rpc_laddr http://localhost:8545 \
  --claimer 0x1234567890123456789012345678901234567890 \
  --validator 0x0987654321098765432109876543210987654321
```

**Parameters:**

- `--claimer`: Your account address (required)
- `--validator`: Validator address to claim from (required)

**Output:** Creates `claimRewards.json` transaction file

## Query Commands

### 1. Query Validator Information

Get detailed information about a validator.

```bash
./congress-cli staking query-validator \
  --rpc_laddr http://localhost:8545 \
  --address 0x0987654321098765432109876543210987654321
```

**Example Output:**

```
✅ Validator Information
Address: 0x0987654321098765432109876543210987654321
Self Stake: 10000 JU
Total Delegated: 50000 JU
Total Stake: 60000 JU
Commission Rate: 500 basis points
Is Jailed: false
Jail Until Block: 0
```

### 2. Query Delegation Information

Get delegation details between a delegator and validator.

```bash
./congress-cli staking query-delegation \
  --rpc_laddr http://localhost:8545 \
  --delegator 0x1234567890123456789012345678901234567890 \
  --validator 0x0987654321098765432109876543210987654321
```

**Example Output:**

```
✅ Delegation Information
Delegator: 0x1234567890123456789012345678901234567890
Validator: 0x0987654321098765432109876543210987654321
Delegated Amount: 1000 JU
Pending Rewards: 25 JU
Unbonding Amount: 0 JU
Unbonding Block: 0
```

### 3. List Top Validators

Get the list of top validators ranked by total stake.

```bash
./congress-cli staking list-top-validators \
  --rpc_laddr http://localhost:8545 \
  --limit 21
```

**Example Output:**

```
✅ Top Validators
Count: 21
1. 0x0987654321098765432109876543210987654321
2. 0x1234567890123456789012345678901234567890
...
```

## Transaction Workflow

All transaction commands follow a 3-step process:

### Step 1: Create Transaction

Use any staking transaction command to generate an unsigned transaction file:

```bash
./congress-cli staking register-validator \
  --proposer 0x1234567890123456789012345678901234567890 \
  --stake-amount 10000 \
  --commission-rate 500
```

This creates `registerValidator.json` with transaction data.

### Step 2: Sign Transaction

Sign the transaction using your keystore:

```bash
./congress-cli sign \
  --file registerValidator.json \
  --key ./keystore/UTC--2023-... \
  --password ./password.txt \
  --chainId 2025
```

This creates `registerValidator_signed.json` with the signed transaction.

### Step 3: Broadcast Transaction

Send the signed transaction to the network:

```bash
./congress-cli send \
  --file registerValidator_signed.json \
  --rpc_laddr http://localhost:8545
```

The transaction will be broadcast and confirmed on the blockchain.

## Configuration

### RPC Endpoints

Common JuChain RPC endpoints:

- **Mainnet**: `https://rpc.ju.finance`
- **Testnet**: `https://testnet-rpc.ju.finance`
- **Local**: `http://localhost:8545`

### Chain IDs

- **Mainnet**: `2025`
- **Testnet**: `202588`
- **Local**: Your custom chain ID

### Gas Configuration

The CLI automatically estimates gas with a 20% buffer. For custom gas settings, modify the transaction JSON before signing.

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

**RPC Connection Failed**

```bash
❌ Validation Error: invalid RPC URL format: localhost:8545
```

Solution: Use full URL with protocol: `http://localhost:8545`

**Invalid Address Format**

```bash
❌ Validation Error: invalid address format: 0x123
```

Solution: Use complete 40-character hex addresses with 0x prefix

**Insufficient Stake Amount**

```bash
❌ Validation Error: stake amount must be at least 10000 JU
```

Solution: Increase stake amount to meet minimum requirements

**Commission Rate Too High**

```bash
❌ Validation Error: commission rate must be between 0 and 10000 (100%)
```

Solution: Use commission rate between 0-10000 basis points

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
CHAIN_ID="2025"

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
