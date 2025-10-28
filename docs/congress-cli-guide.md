# Congress JPoSA Consensus Management Tool User Guide

## Test Environment Account Configuration

To demonstrate the complete voting process, this document uses the following pre-configured test accounts:

```bash
# Validator 1 account configuration
VALIDATOR1_ADDRESS=0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
VALIDATOR1_PRIVATE_KEY=ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
VALIDATOR1_PASSWORD=123456

# Validator 2 account configuration
VALIDATOR2_ADDRESS=0x70997970C51812dc3A010C7d01b50e0d17dc79C8
VALIDATOR2_PRIVATE_KEY=59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d
VALIDATOR2_PASSWORD=123456

# Validator 3 account configuration
VALIDATOR3_ADDRESS=0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC
VALIDATOR3_PRIVATE_KEY=5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a
VALIDATOR3_PASSWORD=123456

# Validator 4 account configuration
VALIDATOR4_ADDRESS=0x90F79bf6EB2c4f870365E785982E1f101E93b906
VALIDATOR4_PRIVATE_KEY=7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6
VALIDATOR4_PASSWORD=123456

# Validator 5 account configuration
VALIDATOR5_ADDRESS=0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65
VALIDATOR5_PRIVATE_KEY=47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a
VALIDATOR5_PASSWORD=123456

# Validator 6 account configuration (to be added as new validator)
VALIDATOR6_ADDRESS=0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc
VALIDATOR6_PRIVATE_KEY=8b3a350cf5c34c9194ca85829a2df0ec3153be0318b5e2d3348e872092edffba
VALIDATOR6_PASSWORD=123456
```

> **⚠️ Important**: The above private keys are for testing purposes only, never use in production!

## Private Chain Environment Keystore Configuration

In private chain environment, validator keystore files are pre-configured at the following locations:

```bash
# Private chain data directory
PRIVATE_CHAIN_PATH=$HOME/ju-chain-work/chain/private-chain

# Validator 1 keystore file (using wildcard)
VALIDATOR1_KEYSTORE=${PRIVATE_CHAIN_PATH}/data-validator1/keystore/UTC--*--f39fd6e51aad88f6f4ce6ab8827279cfffb92266
VALIDATOR1_PASSWORD=${PRIVATE_CHAIN_PATH}/data-validator1/password.txt

# Validator 2 keystore file (using wildcard)
VALIDATOR2_KEYSTORE=${PRIVATE_CHAIN_PATH}/data-validator2/keystore/UTC--*--70997970c51812dc3a010c7d01b50e0d17dc79c8
VALIDATOR2_PASSWORD=${PRIVATE_CHAIN_PATH}/data-validator2/password.txt

# Validator 3 keystore file (using wildcard)
VALIDATOR3_KEYSTORE=${PRIVATE_CHAIN_PATH}/data-validator3/keystore/UTC--*--3c44cdddb6a900fa2b585dd299e03d12fa4293bc
VALIDATOR3_PASSWORD=${PRIVATE_CHAIN_PATH}/data-validator3/password.txt

# Validator 4 keystore file (using wildcard)
VALIDATOR4_KEYSTORE=${PRIVATE_CHAIN_PATH}/data-validator4/keystore/UTC--*--90f79bf6eb2c4f870365e785982e1f101e93b906
VALIDATOR4_PASSWORD=${PRIVATE_CHAIN_PATH}/data-validator4/password.txt

# Validator 5 keystore file (using wildcard)
VALIDATOR5_KEYSTORE=${PRIVATE_CHAIN_PATH}/data-validator5/keystore/UTC--*--15d34aaf54267db7d7c367839aaf71a00a2c6a65
VALIDATOR5_PASSWORD=${PRIVATE_CHAIN_PATH}/data-validator5/password.txt

# Validator 6 keystore file (using wildcard)
VALIDATOR6_KEYSTORE=${PRIVATE_CHAIN_PATH}/data-validator6/keystore/UTC--*--340d92a853ae20a6e7a5b86272fa47aff83a8f7a
VALIDATOR6_PASSWORD=${PRIVATE_CHAIN_PATH}/data-validator6/password.txt
```

All validators use password `123456`, stored in their respective `password.txt` files.

## Flexible Keystore File Lookup

To avoid hardcoding specific timestamps, we provide flexible keystore file lookup methods:

### Method 1: Using find command

```bash
# Get Validator 1 keystore file
VALIDATOR1_KEY=$(find $HOME/ju-chain-work/chain/private-chain/data-validator1/keystore/ -name "*--f39fd6e51aad88f6f4ce6ab8827279cfffb92266" | head -1)

# Get Validator 2 keystore file
VALIDATOR2_KEY=$(find $HOME/ju-chain-work/chain/private-chain/data-validator2/keystore/ -name "*--70997970c51812dc3a010c7d01b50e0d17dc79c8" | head -1)
```

### Method 2: Using wildcards

```bash
# Use wildcards in shell script
VALIDATOR1_KEY=$HOME/ju-chain-work/chain/private-chain/data-validator1/keystore/*--f39fd6e51aad88f6f4ce6ab8827279cfffb92266
VALIDATOR2_KEY=$HOME/ju-chain-work/chain/private-chain/data-validator2/keystore/*--70997970c51812dc3a010c7d01b50e0d17dc79c8
```

### Method 3: Utility function

```bash
# Define function to get keystore file
get_validator_key() {
    local validator_num=$1
    local address=$(echo $2 | tr '[:upper:]' '[:lower:]')
    find $HOME/ju-chain-work/chain/private-chain/data-validator${validator_num}/keystore/ -name "*--${address}" | head -1
}

# Usage example
VALIDATOR1_KEY=$(get_validator_key 1 f39fd6e51aad88f6f4ce6ab8827279cfffb92266)
VALIDATOR2_KEY=$(get_validator_key 2 70997970c51812dc3a010c7d01b50e0d17dc79c8)
```

> **Advantage**: This method does not depend on specific timestamps, as long as the address matches, the corresponding keystore file can be found, making it more flexible and maintainable.

## Tool Compilation

First, compile the congress-cli tool:

```shell
cd sys-contract/congress-cli
make build
# The generated executable is located at build/congress-cli
```

## Version Information

This document has been updated to **Congress CLI v1.2.1**, Build Date: 2025-08-27.

```shell
./build/congress-cli version
# Output: Congress CLI Version: 1.2.1, Build Date: 2025-08-27
```

> **💡 Syntax Instructions**:
>
> - Local test: defaults to `http://127.0.0.1:8545`，Chain ID `202599`
> - Testnet:`--chainId 202599 --rpc_laddr https://testnet-rpc.juchain.org`
> - Mainnet:`--chainId 210000 --rpc_laddr https://rpc.juchain.org`

## Configuration Management Commands

### View Configuration Help

Congress CLI v1.2.0 provides convenient configuration query and setup functions.

```shell
# View config subcommand help
./build/congress-cli config --help
```

### Query System Configuration

Query current system configuration parameters:

```shell
# Query all configuration
./build/congress-cli config get

# Query RPC endpoint
./build/congress-cli config get --rpc

# Query chain ID
./build/congress-cli config get --chain-id
```

**Example**:

```shell
# Query all configuration
./build/congress-cli config get
# Output:
# RPC endpoint: https://rpc.juchain.org
# Chain ID: 210000

# Query RPC endpoint only
./build/congress-cli config get --rpc
# Output:RPC endpoint: https://rpc.juchain.org

# Query chain ID only
./build/congress-cli config get --chain-id
# Output:Chain ID: 210000
```

### Modify System Configuration

Set RPC endpoint and chain ID configuration:

```shell
# Set RPC endpoint
./build/congress-cli config set --rpc <RPC_URL>

# Set chain ID
./build/congress-cli config set --chain-id <CHAIN_ID>

# Set both RPC and chain ID
./build/congress-cli config set --rpc <RPC_URL> --chain-id <CHAIN_ID>
```

**Example**:

```shell
# Set local test environment
./build/congress-cli config set --rpc http://127.0.0.1:8545 --chain-id 202599

# Set testnet environment
./build/congress-cli config set --rpc https://testnet-rpc.juchain.org --chain-id 202599

# Set mainnet environment
./build/congress-cli config set --rpc https://rpc.juchain.org --chain-id 210000

# Set RPC endpoint only
./build/congress-cli config set --rpc https://rpc.juchain.org

# Set chain ID only
./build/congress-cli config set --chain-id 210000
```

> **Note**: The config command here is for setting the congress-cli tool configuration (such as RPC endpoint, chain ID), not for modifying blockchain system parameters. For blockchain system parameter modification, refer to Chapter 4.

## 1. Create Proposal

### 1.1. Create Original Transaction

```shell
# Basic syntax
./build/congress-cli create_proposal -p proposer_address -t new_miner_address -o add

# Complete example: Validator 1 creates add proposal for Validator 6
./build/congress-cli create_proposal -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 -t 0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc -o add

# Another example: Remove validator
./build/congress-cli create_proposal -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 -t 0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc -o remove
```

**Parameter Description**:

- `-p, --proposer`: Proposer address (must be a valid validator)
- `-t, --target`: Target address (validator to add or remove)
- `-o, --operation`: Operation type (add or remove)

> Successfully generates `createProposal.json` file

### 1.2. Sign Transaction

```shell
./build/congress-cli sign -f createProposal.json -k keystore_file -p password_file

# Validator 1 signing example
./build/congress-cli sign -f createProposal.json -k validator1.key -p password.file
```

> Successfully generates `createProposal_signed.json` file

### 1.3. Send Transaction

```shell
./build/congress-cli send -f createProposal_signed.json
```

> Successfully outputs proposal information, including important proposal ID:

```text
✅ Transaction confirmed in block 8758
read sender from signed tx is 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
--------CreateProposal----------
Proposal ID: 0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c
Proposer: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
Destination: 0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc
Flag: true
Time: 1756110615
Block: 8758
-----
✅ Transaction broadcast successfully!
```

> **⚠️ Important**: Record the proposal ID, it is needed for voting!

## 2. Proposal Voting

### 2.1. Create Voting Transaction

Now multiple validators need to vote on the proposal. Assuming the proposal ID is `0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c`

```shell
# Basic syntax
./build/congress-cli vote_proposal -s signer_address -i PROPOSAL_ID -a  # Vote in favor
./build/congress-cli vote_proposal -s signer_address -i PROPOSAL_ID     # Vote against

# Validator 1 votes (the proposer must also vote)
./build/congress-cli vote_proposal -s 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 -i 0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c -a

# Validator 2 votes
./build/congress-cli vote_proposal -s 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 -i 0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c -a

# Validator 3 votes
./build/congress-cli vote_proposal -s 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC -i 0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c -a

# Validator 4 votes
./build/congress-cli vote_proposal -s 0x90F79bf6EB2c4f870365E785982E1f101E93b906 -i 0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c -a

# Validator 5 votes
./build/congress-cli vote_proposal -s 0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65 -i 0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c -a
```

**Parameter Description**:

- `-s, --signer`: Signer address (must be a valid validator)
- `-i, --proposalId`: Proposal ID (64-character hexadecimal string)
- `-a, --approve`: Approval flag (use -a for YES, omit for NO)

> Successfully generates `voteProposal.json` file

### 2.2. Sign Transaction

Each vote requires the corresponding validator private key to sign:

```shell
# Validator 1 signs
./build/congress-cli sign -f voteProposal.json -k validator1.key -p password.file

# Validator 2 signs  
./build/congress-cli sign -f voteProposal.json -k validator2.key -p password.file

# Validator 3 signs
./build/congress-cli sign -f voteProposal.json -k validator3.key -p password.file

# Validator 4 signs
./build/congress-cli sign -f voteProposal.json -k validator4.key -p password.file

# Validator 5 signs
./build/congress-cli sign -f voteProposal.json -k validator5.key -p password.file
```

> Successfully generates `voteProposal_signed.json` file

### 2.3. Send Voting Transaction

Each validator needs to send their own voting transaction:

```shell
# Send voting transaction for each validator in sequence
./build/congress-cli send -f voteProposal_signed.json
```

> Successfully outputs confirmation information:

```text
✅ Transaction confirmed in block 8830
read sender from signed tx is 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
✅ Transaction broadcast successfully!
```

### 2.4. Complete Voting Process Example

Below is the complete voting process for adding Validator 6 as a new validator:

```shell
# Step 1: Validator 1 votes
./build/congress-cli vote_proposal -s 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k validator1.key -p password.file
./build/congress-cli send -f voteProposal_signed.json

# Step 2: Validator 2 votes
./build/congress-cli vote_proposal -s 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k validator2.key -p password.file
./build/congress-cli send -f voteProposal_signed.json

# Step 3: Validator 3 votes
./build/congress-cli vote_proposal -s 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k validator3.key -p password.file
./build/congress-cli send -f voteProposal_signed.json

# Step 4: Validator 4 votes
./build/congress-cli vote_proposal -s 0x90F79bf6EB2c4f870365E785982E1f101E93b906 -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k validator4.key -p password.file
./build/congress-cli send -f voteProposal_signed.json

# Step 5: Validator 5 votes
./build/congress-cli vote_proposal -s 0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65 -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k validator5.key -p password.file
./build/congress-cli send -f voteProposal_signed.json
```

> **Note**:
>
> - Each validator can only vote once on the same proposal
> - Enough validators need to vote in favor for the proposal to pass
> - Replace "ProposalID" in the above commands with the actual proposal ID

## 3. Query Operations

### 3.1 Query All Active Miners

```shell
./build/congress-cli miners
```

> Example output:

```text
ℹ️  Fetching validator information...
ℹ️  Found 5 validators:

--- Validator 1 ---
Address: 0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65
Fee Address: 0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65
Status: Active ✅
Accumulated Rewards: 54174000000000
Penalized Rewards: 0
Last Withdraw Block: 0

--- Validator 2 ---
Address: 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC
Fee Address: 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC
Status: Active ✅
Accumulated Rewards: 54174000000000
Penalized Rewards: 0
Last Withdraw Block: 0
...
```

### 3.2 Query Single Miner

```shell
./build/congress-cli miner -a <validator_address>

# Example
./build/congress-cli miner -a 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
```

> Example output:

```text
ℹ️  Querying validator information for: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
Address: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
Fee Address: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
Status: Active ✅
Accumulated Rewards: 54174000000000
Penalized Rewards: 0
Last Withdraw Block: 0
```

**Status Description:**

- Status: Active ✅ = Active validator
- Status: Inactive ❌ = Abnormal state

### 3.3 Query All Proposals

```shell
./build/congress-cli proposals
```

> Example output:

```text
ℹ️  Fetching all proposals...
ℹ️  Found 1 proposal(s):

--- Proposal 1 ---
ID: 0x0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c
Subject: Test Proposal Subject
Content: Test Proposal Content
Type: 4
Status: Voting
Block Number: 8829
Content Hash: 0xed2e9ba8a0b3ca2b9b7a2c4b8f9a7b5c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0
Contract Address: 0x0000000000000000000000000000000000001000
Current: 1
Start Time: 202599-01-22 20:20:21 +0000 UTC
End Time: 202599-01-23 20:20:21 +0000 UTC
```

### 3.4 Query Single Proposal

```shell
./build/congress-cli proposal -i <proposal_id>

# Example
./build/congress-cli proposal -i 0x0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c
```

> Example output:

```text
ℹ️  Fetching proposal details...
ID: 0x0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c
Subject: Test Proposal Subject
Content: Test Proposal Content
Type: 4
Status: Voting
Block Number: 8829
Content Hash: 0xed2e9ba8a0b3ca2b9b7a2c4b8f9a7b5c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0
Contract Address: 0x0000000000000000000000000000000000001000
Current: 1
Start Time: 202599-01-22 20:20:21 +0000 UTC
End Time: 202599-01-23 20:20:21 +0000 UTC
```

**Proposal Status Description:**

- Status: Voting = Voting in progress
- Status: Passed = Passed
- Status: Failed = Rejected
- Status: Executed = Executed

## 4. Modify Parameter Configuration

### 4.1 Create Configuration Modification Proposal

```shell
./build/congress-cli create_config_proposal -p <proposer_address> -i <config_item_id> -v <config_value>

# Example: Modify proposal duration to 86400 seconds
./build/congress-cli create_config_proposal -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 -i 0 -v 86400

# Example: Modify withdrawal cooldown period to 10 blocks (approx. 10 seconds)
./build/congress-cli create_config_proposal -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 -i 4 -v 10
```

**Configuration Item ID Reference:**

- 0: proposalLastingPeriod (proposal duration, in seconds)
- 1: punishThreshold (punishment threshold)
- 2: removeThreshold (removal threshold)
- 3: decreaseRate (decrease rate)
- 4: withdrawProfitPeriod (withdrawal profit period, in blocks)

**Parameter Description:**

- `-p, --proposer`: Proposer address (must be a valid validator)
- `-i, --cid`: Configuration item ID (0-4)
- `-v, --value`: New value for configuration item

> Successfully generates `createUpdateConfigProposal.json` file

### 4.2 Sign and Send Transaction

Configuration modification proposal signing and sending process is the same as regular proposals:

```shell
# Sign transaction (note filename is createUpdateConfigProposal.json)
./build/congress-cli sign -f createUpdateConfigProposal.json -k /path/to/validator.key -p password.file

# Send transaction
./build/congress-cli send -f createUpdateConfigProposal_signed.json
```

### 4.3 Complete Configuration Modification Process Example

Below is the complete process for modifying withdrawal cooldown period:

```shell
# Step 1: Create configuration proposal
./build/congress-cli create_config_proposal -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 -i 4 -v 10

# Step 2: Sign proposal
./build/congress-cli sign \
  -f createUpdateConfigProposal.json \
  -k $(find $HOME/ju-chain-work/chain/private-chain/data-validator1/keystore/ -name "*--f39fd6e51aad88f6f4ce6ab8827279cfffb92266" | head -1) \
  -p $HOME/ju-chain-work/chain/private-chain/data-validator1/password.txt

# Step 3: Send proposal
./build/congress-cli send -f createUpdateConfigProposal_signed.json

# Step 4: Record proposal ID (from output)
# Example proposal ID from output: 0xd87a55165c909c9b4ef949a3d697e3b26a6a66eee38b2ed519f52f8acd342539

# Step 5: Validator voting (requires sufficient validators to vote)
./build/congress-cli vote_proposal -s 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 -i 0xd87a55165c909c9b4ef949a3d697e3b26a6a66eee38b2ed519f52f8acd342539 -a
./build/congress-cli sign -f voteProposal.json -k $(find $HOME/ju-chain-work/chain/private-chain/data-validator1/keystore/ -name "*--f39fd6e51aad88f6f4ce6ab8827279cfffb92266" | head -1) -p password.file
./build/congress-cli send -f voteProposal_signed.json

# Repeat voting process for other validators...

# Step 6: View proposal status
./build/congress-cli proposal -i 0xd87a55165c909c9b4ef949a3d697e3b26a6a66eee38b2ed519f52f8acd342539
```

### 4.4 Common Configuration Modification Scenarios

#### Scenario 1: Shorten Withdrawal Cooldown Period (for testing)

```shell
# Change withdrawal cooldown period from default 86400 blocks (24 hours) to 10 blocks (approx. 10 seconds)
./build/congress-cli create_config_proposal -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 -i 4 -v 10
```

#### Scenario 2: Adjust Proposal Duration

```shell
# Change proposal duration to 7 days (604800 seconds)
./build/congress-cli create_config_proposal -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 -i 0 -v 604800
```

> **Note:**
>
> - Configuration modification proposals also require sufficient validators to vote before they can pass and execute
> - Configuration modifications take effect immediately, please set parameter values carefully
> - withdrawProfitPeriod is in blocks, assuming 1 block per second calculation

## 5. Miner Profit Withdrawal

### 5.1 Create Withdrawal Transaction

```shell
./build/congress-cli withdraw_profits -a <miner_address>

# Example
./build/congress-cli withdraw_profits -a 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
```

**Parameter Description:**

- `-a, --address`: Miner address to withdraw profits from

### 5.2 Sign and Send Transaction

```shell
# Sign transaction
./build/congress-cli sign -f withdrawProfits.json -k miner1.key -p password.file

# Send transaction
./build/congress-cli send -f withdrawProfits_signed.json
```

> **Note:** Profit withdrawal does not require voting process, miners can directly withdraw their own profits

## 6. Staking Operations

### 6.1 Staking Command Overview

Congress CLI v1.2.0 has added a complete staking module, supporting validator registration, delegation, query and other operations.

```shell
# View staking subcommand help
./build/congress-cli staking --help
```

**Available Staking Subcommands:**

1. `register-validator` - Register validator
2. `delegate` - Delegate staking to validator
3. `undelegate` - Undelegate
4. `query-validator` - Query specific validator information
5. `list-top-validators` - List top validators
6. `unjail` - Unjail validator
7. `withdraw` - Withdraw delegation rewards

### 6.2 Register Validator

To become a new validator, you must meet the minimum stake requirement.

```shell
./build/congress-cli staking register-validator \
  --proposer <proposer_address> \
  --stake-amount <stake_amount> \
  --commission-rate <commission_rate>
```

**Parameter Description**:

- `--proposer`: Proposer address (required)
- `--stake-amount`: Amount of JU to stake (required, minimum 10000 JU)
- `--commission-rate`: Commission rate in basis points (0-10000, e.g., 500 means 5%)

**Example**:

```shell
# Local test environment
./build/congress-cli staking register-validator \
  --proposer 0x3F9DDeBE20b24B0DEC1d0B5A3c6e8Cb8D3eCF6A5 \
  --stake-amount 10000 \
  --commission-rate 500

# Mainnet environment
./build/congress-cli staking register-validator \
  --proposer 0x3F9DDeBE20b24B0DEC1d0B5A3c6e8Cb8D3eCF6A5 \
  --stake-amount 10000 \
  --commission-rate 500 \
  --rpc https://rpc.juchain.org --chainId 210000
```

### 6.3 Delegate Stake

Delegate JU tokens to validators to earn rewards.

```shell
./build/congress-cli staking delegate \
  --validator <validator_address> \
  --amount <delegation_amount>
```

**Example**:

```shell
# Delegate 1000 JU to validator
./build/congress-cli staking delegate \
  --validator 0x3F9DDeBE20b24B0DEC1d0B5A3c6e8Cb8D3eCF6A5 \
  --amount 1000
```

### 6.4 Undelegate

Undelegate previous delegation, requires waiting for unbonding period.

```shell
./build/congress-cli staking undelegate \
  --validator <validator_address> \
  --amount <undelegation_amount>
```

**Example**:

```shell
# Undelegate 500 JU
./build/congress-cli staking undelegate \
  --validator 0x3F9DDeBE20b24B0DEC1d0B5A3c6e8Cb8D3eCF6A5 \
  --amount 500
```

### 6.5 Query Validator Information

Query detailed information of specific validator, including stake, commission rate, status, etc.

```shell
./build/congress-cli staking query-validator --address <validator_address>
```

**Example**:

```shell
# Query validator information
./build/congress-cli staking query-validator \
  --address 0x3F9DDeBE20b24B0DEC1d0B5A3c6e8Cb8D3eCF6A5
```

> **Note**: Current version may encounter ABI parsing errors with some query commands, this is a known issue being fixed.

### 6.6 List Top Validators

View current active top validators list.

```shell
./build/congress-cli staking list-top-validators [limit]
```

**Example**:

```shell
# View top 15 validators (default)
./build/congress-cli staking list-top-validators

# View top 10 validators
./build/congress-cli staking list-top-validators 10
```

### 6.7 Unjail

Validators jailed for violations can apply to remove jail status.

```shell
./build/congress-cli staking unjail --validator <validator_address>
```

### 6.9 Edit Validator Information

Edit existing validator information, including fee address and description.

```shell
./build/congress-cli staking edit-validator \
  --validator <validator_address> \
  --fee-addr <fee_address> \
  [--moniker <display_name>] \
  [--identity <identity>] \
  [--website <website_url>] \
  [--email <contact_email>] \
  [--details <detailed_description>]
```

**Parameter Description**:

- `--validator`: Validator address to edit (required)
- `--fee-addr`: Fee address for receiving mining rewards (required)
- `--moniker`: Validator display name (optional)
- `--identity`: Validator identity, such as Keybase signature (optional)
- `--website`: Validator website URL (optional)
- `--email`: Validator contact email (optional)
- `--details`: Validator detailed description (optional)

**Example**:

```shell
# Set fee address for validator6
./build/congress-cli staking edit-validator \
  --validator 0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc \
  --fee-addr 0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc \
  --moniker "validator6" \
  --details "Validator6 node with fee address configured"

# Sign transaction (using validator6 own key)
./build/congress-cli sign \
  -f editValidator.json \
  -k $(find $HOME/ju-chain-work/chain/private-chain/data-validator6/keystore/ -name "*--9965507d1a55bcc2695c58ba16fb37d819b0a4dc" | head -1) \
  -p $HOME/ju-chain-work/chain/private-chain/data-validator6/password.txt

# Send transaction
./build/congress-cli send -f editValidator_signed.json
```

> **Note**:
>
> - Only validators themselves can edit their information
> - Fee address is used to receive block mining rewards
> - If validator fee address is `0x0000...`, should set correct fee address to ensure rewards are received

### 6.10 Withdraw Rewards

Withdraw delegation rewards.

```shell
./build/congress-cli staking withdraw --validator <validator_address>
```

## 7. Complete End-to-End Process: Adding Validator 6

This chapter demonstrates the complete process from start to finish, having validators 1-5 vote for validator 6 to make it a new validator.

### 7.1 Prerequisites

Keystore files in private chain environment are pre-configured at the following locations:

```bash
# Set private chain path variable
PRIVATE_CHAIN_PATH=$HOME/ju-chain-work/chain/private-chain

# Validator keystore file paths (using wildcard pattern)
VALIDATOR1_KEY=$(find ${PRIVATE_CHAIN_PATH}/data-validator1/keystore/ -name "*--f39fd6e51aad88f6f4ce6ab8827279cfffb92266" | head -1)
VALIDATOR2_KEY=$(find ${PRIVATE_CHAIN_PATH}/data-validator2/keystore/ -name "*--70997970c51812dc3a010c7d01b50e0d17dc79c8" | head -1)
VALIDATOR3_KEY=$(find ${PRIVATE_CHAIN_PATH}/data-validator3/keystore/ -name "*--3c44cdddb6a900fa2b585dd299e03d12fa4293bc" | head -1)
VALIDATOR4_KEY=$(find ${PRIVATE_CHAIN_PATH}/data-validator4/keystore/ -name "*--90f79bf6eb2c4f870365e785982e1f101e93b906" | head -1)
VALIDATOR5_KEY=$(find ${PRIVATE_CHAIN_PATH}/data-validator5/keystore/ -name "*--15d34aaf54267db7d7c367839aaf71a00a2c6a65" | head -1)

# Password file path (all validators use same password file)
PASSWORD_FILE=${PRIVATE_CHAIN_PATH}/data-validator1/password.txt
```

> **Note**: All validators use password `123456`, can use any validator directory `password.txt` file.

### 7.2 Step 1: Validator 1 Creates Proposal

```shell
# Create proposal to add Validator 6
./build/congress-cli create_proposal \
  -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  -t 0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc \
  -o add

# Validator 1 signs (using dynamically found keystore file path)
./build/congress-cli sign \
  -f createProposal.json \
  -k $(find $HOME/ju-chain-work/chain/private-chain/data-validator1/keystore/ -name "*--f39fd6e51aad88f6f4ce6ab8827279cfffb92266" | head -1) \
  -p $HOME/ju-chain-work/chain/private-chain/data-validator1/password.txt

# Send transaction
./build/congress-cli send -f createProposal_signed.json
```

> **Important**: Record the proposal ID from output, e.g.: `0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c`

### 7.3 Step 2: 5 Validators Vote

Replace `PROPOSAL_ID` in the following commands with the actual proposal ID obtained in step 1.

```shell
# Validator 1 votes
./build/congress-cli vote_proposal -s 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 -i PROPOSAL_ID -a
./build/congress-cli sign \
  -f voteProposal.json \
  -k $(find $HOME/ju-chain-work/chain/private-chain/data-validator1/keystore/ -name "*--f39fd6e51aad88f6f4ce6ab8827279cfffb92266" | head -1) \
  -p $HOME/ju-chain-work/chain/private-chain/data-validator1/password.txt
./build/congress-cli send -f voteProposal_signed.json

# Validator 2 votes
./build/congress-cli vote_proposal -s 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 -i PROPOSAL_ID -a
./build/congress-cli sign \
  -f voteProposal.json \
  -k $(find $HOME/ju-chain-work/chain/private-chain/data-validator2/keystore/ -name "*--70997970c51812dc3a010c7d01b50e0d17dc79c8" | head -1) \
  -p $HOME/ju-chain-work/chain/private-chain/data-validator2/password.txt
./build/congress-cli send -f voteProposal_signed.json

# Validator 3 votes
./build/congress-cli vote_proposal -s 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC -i PROPOSAL_ID -a
./build/congress-cli sign \
  -f voteProposal.json \
  -k $(find $HOME/ju-chain-work/chain/private-chain/data-validator3/keystore/ -name "*--3c44cdddb6a900fa2b585dd299e03d12fa4293bc" | head -1) \
  -p $HOME/ju-chain-work/chain/private-chain/data-validator3/password.txt
./build/congress-cli send -f voteProposal_signed.json

# Validator 4 votes
./build/congress-cli vote_proposal -s 0x90F79bf6EB2c4f870365E785982E1f101E93b906 -i PROPOSAL_ID -a
./build/congress-cli sign \
  -f voteProposal.json \
  -k $(find $HOME/ju-chain-work/chain/private-chain/data-validator4/keystore/ -name "*--90f79bf6eb2c4f870365e785982e1f101e93b906" | head -1) \
  -p $HOME/ju-chain-work/chain/private-chain/data-validator4/password.txt
./build/congress-cli send -f voteProposal_signed.json

# Validator 5 votes
./build/congress-cli vote_proposal -s 0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65 -i PROPOSAL_ID -a
./build/congress-cli sign \
  -f voteProposal.json \
  -k $(find $HOME/ju-chain-work/chain/private-chain/data-validator5/keystore/ -name "*--15d34aaf54267db7d7c367839aaf71a00a2c6a65" | head -1) \
  -p $HOME/ju-chain-work/chain/private-chain/data-validator5/password.txt
./build/congress-cli send -f voteProposal_signed.json
```

### 7.4 Step 3: Verify Results

```shell
# Query proposal status
./build/congress-cli proposal -i PROPOSAL_ID

# Query Validator 6 status
./build/congress-cli miner -a 0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc

# Query all validators list
./build/congress-cli miners
```

> **Expected Results**:
>
> - Proposal status should show "Passed" or "Executed"
> - Validator 6 should appear in active validator list
> - Total validators should increase from 5 to 6

### 7.5 Automation Script Example

You can also create a script to automate the entire process:

```bash
#!/bin/bash

# Set proposal target address
TARGET_ADDRESS="0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc"
PROPOSER_ADDRESS="0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"

# Private chain path
PRIVATE_CHAIN_PATH="$HOME/ju-chain-work/chain/private-chain"

# Validator address array
VALIDATORS=(
    "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
    "0x70997970C51812dc3A010C7d01b50e0d17dc79C8"
    "0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC"
    "0x90F79bf6EB2c4f870365E785982E1f101E93b906"
    "0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65"
)

# Keystore file array (using dynamic lookup)
KEYS=(
    "$(find ${PRIVATE_CHAIN_PATH}/data-validator1/keystore/ -name "*--f39fd6e51aad88f6f4ce6ab8827279cfffb92266" | head -1)"
    "$(find ${PRIVATE_CHAIN_PATH}/data-validator2/keystore/ -name "*--70997970c51812dc3a010c7d01b50e0d17dc79c8" | head -1)"
    "$(find ${PRIVATE_CHAIN_PATH}/data-validator3/keystore/ -name "*--3c44cdddb6a900fa2b585dd299e03d12fa4293bc" | head -1)"
    "$(find ${PRIVATE_CHAIN_PATH}/data-validator4/keystore/ -name "*--90f79bf6eb2c4f870365e785982e1f101e93b906" | head -1)"
    "$(find ${PRIVATE_CHAIN_PATH}/data-validator5/keystore/ -name "*--15d34aaf54267db7d7c367839aaf71a00a2c6a65" | head -1)"
)

# Password file array
PASSWORDS=(
    "${PRIVATE_CHAIN_PATH}/data-validator1/password.txt"
    "${PRIVATE_CHAIN_PATH}/data-validator2/password.txt"
    "${PRIVATE_CHAIN_PATH}/data-validator3/password.txt"
    "${PRIVATE_CHAIN_PATH}/data-validator4/password.txt"
    "${PRIVATE_CHAIN_PATH}/data-validator5/password.txt"
)

echo "=== Step 1: Create Proposal ==="
./build/congress-cli create_proposal -p $PROPOSER_ADDRESS -t $TARGET_ADDRESS -o add
./build/congress-cli sign -f createProposal.json -k "${KEYS[0]}" -p "${PASSWORDS[0]}"
./build/congress-cli send -f createProposal_signed.json

echo "Please enter proposal ID:"
read PROPOSAL_ID

echo "=== Step 2: Validators Vote ==="
for i in "${!VALIDATORS[@]}"; do
    echo "Validator $((i+1)) voting..."
    ./build/congress-cli vote_proposal -s ${VALIDATORS[$i]} -i $PROPOSAL_ID -a
    ./build/congress-cli sign -f voteProposal.json -k "${KEYS[$i]}" -p "${PASSWORDS[$i]}"
    ./build/congress-cli send -f voteProposal_signed.json
    echo "Validator $((i+1)) voting completed"
done

echo "=== Step 3: Verify Results ==="
./build/congress-cli proposal -i $PROPOSAL_ID
./build/congress-cli miner -a $TARGET_ADDRESS
./build/congress-cli miners
```

## 8. Complete Test Transaction Example

> **Below is a complete test flow example:**

```shell
# 1. Create proposal
./build/congress-cli create_proposal 
  --proposal "test proposal" 
  --action 0 
  --value 1000 
  --proposer 0x3F9DDeBE20b24B0DEC1d0B5A3c6e8Cb8D3eCF6A5

# 2. Vote (4 validators vote in favor)
./build/congress-cli vote_proposal --proposer 0x3F9DDeBE20b24B0DEC1d0B5A3c6e8Cb8D3eCF6A5 --id 1 --vote 1
./build/congress-cli vote_proposal --proposer 0x4B1E2D4D7C8F5A9B6E3F8A2C7D9E1F4B8C5E6A9 --id 1 --vote 1  
./build/congress-cli vote_proposal --proposer 0x5C2F3E5E8D9A6B7C4E5F9B3D8E2A5C9D6F7A8E --id 1 --vote 1
./build/congress-cli vote_proposal --proposer 0x6D3A4F6F9E8B7C5D6A7C8F4E9A3B6D7E8F9C1A --id 1 --vote 1

# 3. Query proposal status
./build/congress-cli query_proposal --id 1

# 4. Delegate stake
./build/congress-cli staking delegate 
  --validator 0x3F9DDeBE20b24B0DEC1d0B5A3c6e8Cb8D3eCF6A5 
  --amount 1000
```

## 9. Mainnet Miner Identity Recovery Operation

> miner1 creates proposal to add 0x029DAB47e268575D4AC167De64052FB228B5fA41 as new miner, after creating proposal, miner1,miner2,miner3 vote to approve

```shell
# step1 Create proposal transaction, sign and send
./build/congress-cli create_proposal -p 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b -t 0x029DAB47e268575D4AC167De64052FB228B5fA41 -o add
./build/congress-cli sign -f createProposal.json -k miner1.key -p password.file
./build/congress-cli send -f createProposal_signed.json
# This command can obtain proposal ID after execution, e.g.: b2be7f3cc702c7a24962df6aed188edbcfdebe20fd55f5670efaedace0e4bcdf

# step2 3 miners vote on proposal (replace PROPOSAL_ID with actual proposal ID from previous step)
# miner1
./build/congress-cli vote_proposal -s 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k miner1.key -p password.file
./build/congress-cli send -f voteProposal_signed.json

# miner2
./build/congress-cli vote_proposal -s 0x81f7a79a51edba249efa812eb2d5478f696f7558 -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k miner2.key -p password.file
./build/congress-cli send -f voteProposal_signed.json

# miner3
./build/congress-cli vote_proposal -s 0x578c39eaf09a4e1abf428c423970b59bb8baf42e -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k miner3.key -p password.file
./build/congress-cli send -f voteProposal_signed.json

# step3 View new miner information
./build/congress-cli miner -a 0x029DAB47e268575D4AC167De64052FB228B5fA41
./build/congress-cli miners
```

## 10. Mainnet Configuration Modification

### 10.1 Create Configuration Proposal

```shell
# Configuration item ID and corresponding information
# 0 proposalLastingPeriod, 1 punishThreshold, 2 removeThreshold, 3 decreaseRate, 4 withdrawProfitPeriod
./build/congress-cli create_config_proposal -p Proposer_Miner_Address -i Config_Item_ID -v Config_Item_Value

# Example: Modify proposalLastingPeriod to 86400 seconds
./build/congress-cli create_config_proposal -p 0xccafa71c31bc11ba24d526fd27ba57d743152807 -i 0 -v 86400

# Example: Modify withdrawProfitPeriod to 10 blocks
./build/congress-cli create_config_proposal -p 0xccafa71c31bc11ba24d526fd27ba57d743152807 -i 4 -v 10

# Sign transaction (note filename is createUpdateConfigProposal.json)
./build/congress-cli sign -f createUpdateConfigProposal.json -k miner1.key -p password.file

# Send transaction
./build/congress-cli send -f createUpdateConfigProposal_signed.json
# This command can obtain proposal ID after execution, record proposal ID for subsequent voting
```

### 10.2 Validators Vote

Configuration proposal voting process is the same as adding validator proposal:

```shell
# Example: Vote on configuration proposal (replace PROPOSAL_ID with actual proposal ID)
# miner1
./build/congress-cli vote_proposal -s 0xccafa71c31bc11ba24d526fd27ba57d743152807 -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k miner1.key -p password.file
./build/congress-cli send -f voteProposal_signed.json

# miner2
./build/congress-cli vote_proposal -s 0x81f7a79a51edba249efa812eb2d5478f696f7558 -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k miner2.key -p password.file
./build/congress-cli send -f voteProposal_signed.json

# miner3
./build/congress-cli vote_proposal -s 0x578c39eaf09a4e1abf428c423970b59bb8baf42e -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k miner3.key -p password.file
./build/congress-cli send -f voteProposal_signed.json
```

### 10.3 View Configuration Proposal Status

```shell
# View proposal details
./build/congress-cli proposal -i PROPOSAL_ID

# View all proposals
./build/congress-cli proposals
```

> **Important Reminders**:
>
> - Configuration modification proposals take effect immediately upon passing
> - withdrawProfitPeriod is in blocks, not seconds
> - Please carefully set configuration parameters to avoid affecting normal network operation

## 11. Mainnet Miner Profit Withdrawal

```shell
# step1 Create original transaction
./build/congress-cli withdraw_profits -a Miner_Address

# step2 Sign transaction
./build/congress-cli sign -f withdrawProfits.json -k miner.key -p password.file

# step3 Send transaction
./build/congress-cli send -f withdrawProfits_signed.json
```

## 12. Tool Information

### 12.1 Version Information

```shell
./build/congress-cli version
```

### 12.2 Help Information

```shell
./build/congress-cli help
./build/congress-cli [command] --help  # View help for specific command
```

## 13. Important Notes

### 13.1 Important Reminders

- ⚠️ **Validator Requirement**: Only current active validators can create proposals and vote
- ⚠️ **Network Sync**: Before recovering miner identity, ensure node is fully synced to latest state
- ⚠️ **Proposal ID**: Each operation generates new proposal ID, must use correct ID
- ⚠️ **Key Security**: Properly secure keystore and password files

### 13.2 Common Errors

1. **"Validator only"**: Current account is not a valid validator
2. **"You cannot vote for a proposal twice"**: This validator has already voted on this proposal
3. **"gas estimation failed"**: Transaction parameter error or network issue

### 13.3 System Contract Addresses

- **Validators**: `0x000000000000000000000000000000000000f000`
- **Punish**: `0x000000000000000000000000000000000000f001`
- **Proposal**: `0x000000000000000000000000000000000000f002`

### 13.4 Network Information

- **Testnet**: `https://testnet-rpc.juchain.org` (Chain ID: 202599)
- **Mainnet**: `https://rpc.juchain.org` (Chain ID: 210000)
