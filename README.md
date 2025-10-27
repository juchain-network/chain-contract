# Ju System Contracts

Congress Proof-of-Authority (PoA) consensus system contracts for JuChain.

## Overview

This project contains:
- **System Contracts**: Validators, Proposal, and Punish contracts for network governance
- **CLI Tool**: Go-based command-line interface for managing validators and proposals

## System Contracts

### Contract Addresses
- Validators: `0x000000000000000000000000000000000000f000`
- Punish: `0x000000000000000000000000000000F001`
- Proposal: `0x000000000000000000000000000000000000f002`

### Key Parameters
- `punishThreshold = 24`: Threshold for confiscating validator rewards
- `removeThreshold = 48`: Threshold for removing validator from active set
- `decreaseRate = 24`: Error count reduction rate per epoch
- `withdrawProfitPeriod = 28800`: Block interval between profit withdrawals
- `proposalLastingPeriod = 7 days`: Validity period for proposals

### Consensus Mechanism

Congress PoA combines elements of `clique` and `DPoS` algorithms with system contracts for validator management.

**Key Features:**
1. Automatic validator set updates based on stake ranking
2. Reward distribution among validators
3. Punishment mechanism for missed blocks
4. Governance through proposals and voting

### Becoming a Validator

1. Create a proposal via `Proposal.createProposal(dst, flag, details)`
2. Current validators vote via `Proposal.voteProposal(id, auth)`
3. When majority approve, call `Validators.createOrEditValidator(...)` to register information
4. New validator becomes active in the next epoch cycle

## CLI Tool

Build the CLI:
```bash
make build
```

### Network Configuration
```bash
# Testnet
./congress-cli --chainId 202599 --rpc_laddr https://testnet-rpc.juchain.org

# Mainnet
./congress-cli --chainId 210000 --rpc_laddr https://rpc.juchain.org
```

### Common Operations

#### 1. Create Proposal
```bash
./congress-cli create_proposal -p <proposer> -t <target> -o add --rpc_laddr <rpc>
./congress-cli sign -f createProposal.json -k <key> -p <password> --chainId <id>
./congress-cli send -f createProposal_signed.json --rpc_laddr <rpc>
```

#### 2. Vote on Proposal
```bash
./congress-cli vote_proposal -s <signer> -i <proposalId> -a <true/false> --rpc_laddr <rpc>
./congress-cli sign -f voteProposal.json -k <key> -p <password> --chainId <id>
./congress-cli send -f voteProposal_signed.json --rpc_laddr <rpc>
```

#### 3. Query Validators
```bash
# List all active validators
./congress-cli miners --rpc_laddr <rpc>

# Get validator details
./congress-cli miner --rpc_laddr <rpc> -a <validator_address>
```

#### 4. Withdraw Profits
```bash
./congress-cli withdraw_profits -a <validator_address> --rpc_laddr <rpc>
./congress-cli sign -f withdrawProfits.json -k <key> -p <password> --chainId <id>
./congress-cli send -f withdrawProfits_signed.json --rpc_laddr <rpc>
```

#### 5. Update Configuration
```bash
# Configuration IDs:
# 0: proposalLastingPeriod, 1: punishThreshold, 2: removeThreshold
# 3: decreaseRate, 4: withdrawProfitPeriod
./congress-cli create_config_proposal -p <proposer> -c <configId> -v <value> --rpc_laddr <rpc>
# Then vote using the same process as validator proposal
```


## Genesis Configuration

Configure consensus algorithm as `congress`:
```json
"congress": {
    "period": 3,
    "epoch": 200
}
```

Set contract bytecode in genesis for system contract addresses.