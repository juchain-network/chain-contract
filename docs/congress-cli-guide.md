# Congress POSA CLI Usage Guide

## Test Accounts (Demo)

For walkthroughs we use these preset test accounts:

```bash
# Validator 1
VALIDATOR1_ADDRESS=0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
VALIDATOR1_PRIVATE_KEY=ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
VALIDATOR1_PASSWORD=123456

# Validator 2
VALIDATOR2_ADDRESS=0x70997970C51812dc3A010C7d01b50e0d17dc79C8
VALIDATOR2_PRIVATE_KEY=59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d
VALIDATOR2_PASSWORD=123456

# Validator 3
VALIDATOR3_ADDRESS=0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC
VALIDATOR3_PRIVATE_KEY=5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a
VALIDATOR3_PASSWORD=123456

# Validator 4
VALIDATOR4_ADDRESS=0x90F79bf6EB2c4f870365E785982E1f101E93b906
VALIDATOR4_PRIVATE_KEY=7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6
VALIDATOR4_PASSWORD=123456

# Validator 5
VALIDATOR5_ADDRESS=0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65
VALIDATOR5_PRIVATE_KEY=47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a
VALIDATOR5_PASSWORD=123456

# Validator 6 (to be added)
VALIDATOR6_ADDRESS=0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc
VALIDATOR6_PRIVATE_KEY=8b3a350cf5c34c9194ca85829a2df0ec3153be0318b5e2d3348e872092edffba
VALIDATOR6_PASSWORD=123456
```

> **⚠️ Important**: These private keys are for testing only—never use them in production.

## Keystore Locations (Private Chain)

Validator keystores are preconfigured here:

```bash
# Private chain root
PRIVATE_CHAIN_PATH=$HOME/ju-chain-work/chain/private-chain

# Validator 1 keystore (wildcard)
VALIDATOR1_KEYSTORE=${PRIVATE_CHAIN_PATH}/data-validator1/keystore/UTC--*--f39fd6e51aad88f6f4ce6ab8827279cfffb92266
VALIDATOR1_PASSWORD=${PRIVATE_CHAIN_PATH}/data-validator1/password.txt

# Validator 2 keystore (wildcard)
VALIDATOR2_KEYSTORE=${PRIVATE_CHAIN_PATH}/data-validator2/keystore/UTC--*--70997970c51812dc3a010c7d01b50e0d17dc79c8
VALIDATOR2_PASSWORD=${PRIVATE_CHAIN_PATH}/data-validator2/password.txt

# Validator 3 keystore (wildcard)
VALIDATOR3_KEYSTORE=${PRIVATE_CHAIN_PATH}/data-validator3/keystore/UTC--*--3c44cdddb6a900fa2b585dd299e03d12fa4293bc
VALIDATOR3_PASSWORD=${PRIVATE_CHAIN_PATH}/data-validator3/password.txt

# Validator 4 keystore (wildcard)
VALIDATOR4_KEYSTORE=${PRIVATE_CHAIN_PATH}/data-validator4/keystore/UTC--*--90f79bf6eb2c4f870365e785982e1f101e93b906
VALIDATOR4_PASSWORD=${PRIVATE_CHAIN_PATH}/data-validator4/password.txt

# Validator 5 keystore (wildcard)
VALIDATOR5_KEYSTORE=${PRIVATE_CHAIN_PATH}/data-validator5/keystore/UTC--*--15d34aaf54267db7d7c367839aaf71a00a2c6a65
VALIDATOR5_PASSWORD=${PRIVATE_CHAIN_PATH}/data-validator5/password.txt

# Validator 6 keystore (wildcard)
VALIDATOR6_KEYSTORE=${PRIVATE_CHAIN_PATH}/data-validator6/keystore/UTC--*--340d92a853ae20a6e7a5b86272fa47aff83a8f7a
VALIDATOR6_PASSWORD=${PRIVATE_CHAIN_PATH}/data-validator6/password.txt
```

All validator passwords are `123456`, stored in each `password.txt`.

## Flexible Keystore Lookup

Avoid hardcoding timestamps by finding files dynamically:

### Method 1: `find`

```bash
# Validator 1
VALIDATOR1_KEY=$(find $HOME/ju-chain-work/chain/private-chain/data-validator1/keystore/ -name "*--f39fd6e51aad88f6f4ce6ab8827279cfffb92266" | head -1)

# Validator 2
VALIDATOR2_KEY=$(find $HOME/ju-chain-work/chain/private-chain/data-validator2/keystore/ -name "*--70997970c51812dc3a010c7d01b50e0d17dc79c8" | head -1)
```

### Method 2: Wildcards

```bash
VALIDATOR1_KEY=$HOME/ju-chain-work/chain/private-chain/data-validator1/keystore/*--f39fd6e51aad88f6f4ce6ab8827279cfffb92266
VALIDATOR2_KEY=$HOME/ju-chain-work/chain/private-chain/data-validator2/keystore/*--70997970c51812dc3a010c7d01b50e0d17dc79c8
```

### Method 3: Helper Function

```bash
get_validator_key() {
    local validator_num=$1
    local address=$(echo $2 | tr '[:upper:]' '[:lower:]')
    find $HOME/ju-chain-work/chain/private-chain/data-validator${validator_num}/keystore/ -name "*--${address}" | head -1
}

VALIDATOR1_KEY=$(get_validator_key 1 f39fd6e51aad88f6f4ce6ab8827279cfffb92266)
VALIDATOR2_KEY=$(get_validator_key 2 70997970c51812dc3a010c7d01b50e0d17dc79c8)
```

> **Why**: No timestamp dependency—matching by address keeps scripts simple and maintainable.

## Build the Tool

Compile congress-cli:

```shell
cd sys-contract/congress-cli
make build
# output: build/congress-cli
```

## Version

This guide targets **Congress CLI v1.2.1**, build date: 2025-08-27.

```shell
./build/congress-cli version
# Congress CLI Version: 1.2.1, Build Date: 2025-08-27
```

> **💡 Flags**  
> - Local test: default `http://127.0.0.1:8545`, chainId `202599`  
> - Testnet: `--chainId 202599 --rpc_laddr https://testnet-rpc.juchain.org`  
> - Mainnet: `--chainId 210000 --rpc_laddr https://rpc.juchain.org`

## Config Commands

### Help

```shell
./build/congress-cli config --help
```

### Get Config

```shell
# all
./build/congress-cli config get
# RPC only
./build/congress-cli config get --rpc
# chain ID only
./build/congress-cli config get --chain-id
```

### Set Config

```shell
# set RPC
./build/congress-cli config set --rpc <RPC_URL>
# set chain ID
./build/congress-cli config set --chain-id <CHAIN_ID>
# set both
./build/congress-cli config set --rpc <RPC_URL> --chain-id <CHAIN_ID>
```

> **Note**: `config` sets CLI defaults (RPC, chainId). On-chain parameter changes are in Section 4.

## 1. Create a Proposal

### 1.1 Build the Unsigned Tx

```shell
# syntax
./build/congress-cli create_proposal -p PROPOSER -t TARGET -o add

# example: validator1 proposes adding validator6
./build/congress-cli create_proposal \
  -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  -t 0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc \
  -o add

# remove example
./build/congress-cli create_proposal \
  -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  -t 0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc \
  -o remove
```

**Params**

- `-p, --proposer`: active validator
- `-t, --target`: address to add/remove
- `-o, --operation`: `add` or `remove`

> Outputs `createProposal.json`

### 1.2 Sign

```shell
./build/congress-cli sign -f createProposal.json -k wallet.key -p password.file
# example
./build/congress-cli sign -f createProposal.json -k validator1.key -p password.file
```

> Outputs `createProposal_signed.json`

### 1.3 Send

```shell
./build/congress-cli send -f createProposal_signed.json
```

> Prints proposal info including the proposal ID.

## 2. Vote on a Proposal

### 2.1 Build Vote Tx

Assume proposal ID `0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c`:

```shell
# approve
./build/congress-cli vote_proposal -s SIGNER -i PROPOSAL_ID -a
# reject (omit -a)
./build/congress-cli vote_proposal -s SIGNER -i PROPOSAL_ID

# examples
./build/congress-cli vote_proposal -s 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 -i 0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c -a
./build/congress-cli vote_proposal -s 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 -i 0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c -a
./build/congress-cli vote_proposal -s 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC -i 0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c -a
./build/congress-cli vote_proposal -s 0x90F79bf6EB2c4f870365E785982E1f101E93b906 -i 0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c -a
./build/congress-cli vote_proposal -s 0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65 -i 0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c -a
```

**Params**

- `-s, --signer`: active validator
- `-i, --proposalId`: 64-hex ID
- `-a, --approve`: approve (omit for reject)

> Outputs `voteProposal.json`

### 2.2 Sign

```shell
./build/congress-cli sign -f voteProposal.json -k validator1.key -p password.file
./build/congress-cli sign -f voteProposal.json -k validator2.key -p password.file
...
```

> Outputs `voteProposal_signed.json`

### 2.3 Send

```shell
./build/congress-cli send -f voteProposal_signed.json
```

### 2.4 End-to-End Voting Example

```shell
# validator1
./build/congress-cli vote_proposal -s 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k validator1.key -p password.file
./build/congress-cli send -f voteProposal_signed.json

# validator2
./build/congress-cli vote_proposal -s 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k validator2.key -p password.file
./build/congress-cli send -f voteProposal_signed.json

# validator3
./build/congress-cli vote_proposal -s 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k validator3.key -p password.file
./build/congress-cli send -f voteProposal_signed.json

# validator4
./build/congress-cli vote_proposal -s 0x90F79bf6EB2c4f870365E785982E1f101E93b906 -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k validator4.key -p password.file
./build/congress-cli send -f voteProposal_signed.json

# validator5
./build/congress-cli vote_proposal -s 0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65 -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k validator5.key -p password.file
./build/congress-cli send -f voteProposal_signed.json
```

> **Notes**  
> - One vote per validator per proposal  
> - Majority of active validators must approve  
> - Replace `PROPOSAL_ID` with the real ID

## 3. Queries

### 3.1 List Active Validators

```shell
./build/congress-cli miners
```

### 3.2 Query One Validator

```shell
./build/congress-cli miner -a <validator_address>
```

Status meanings: Active ✅ / Inactive ❌

### 3.3 List All Proposals

```shell
./build/congress-cli proposals
```

### 3.4 Query One Proposal

```shell
./build/congress-cli proposal -i <proposal_id>
```

Proposal status: Voting / Passed / Failed / Executed

## 4. Change On-Chain Parameters

### 4.1 Create Config Proposal

```shell
./build/congress-cli create_config_proposal -p <proposer> -i <cid> -v <value>

# examples
./build/congress-cli create_config_proposal -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 -i 0 -v 86400   # proposal lasting seconds
./build/congress-cli create_config_proposal -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 -i 4 -v 10     # withdrawProfitPeriod blocks
```

**CID map**

- 0: proposalLastingPeriod (seconds)
- 1: punishThreshold
- 2: removeThreshold
- 3: decreaseRate
- 4: withdrawProfitPeriod (blocks)

**Params**

- `-p, --proposer`: active validator
- `-i, --cid`: 0–4
- `-v, --value`: new value

> Outputs `createUpdateConfigProposal.json`

### 4.2 Sign & Send

```shell
./build/congress-cli sign -f createUpdateConfigProposal.json -k /path/to/validator.key -p password.file
./build/congress-cli send -f createUpdateConfigProposal_signed.json
```

### 4.3 Full Example (shorten withdraw cooldown)

```shell
# create
./build/congress-cli create_config_proposal -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 -i 4 -v 10

# sign
./build/congress-cli sign \
  -f createUpdateConfigProposal.json \
  -k $(find $HOME/ju-chain-work/chain/private-chain/data-validator1/keystore/ -name "*--f39fd6e51aad88f6f4ce6ab8827279cfffb92266" | head -1) \
  -p $HOME/ju-chain-work/chain/private-chain/data-validator1/password.txt

# send
./build/congress-cli send -f createUpdateConfigProposal_signed.json

# record proposal ID from output, then validators vote as usual
```

### 4.4 Common Scenarios

- Shorten withdrawProfitPeriod for testing: `-i 4 -v 10`
- Set proposal duration to 7 days (604800s): `-i 0 -v 604800`

> **Notes**  
> - Config proposals still need majority approval  
> - Changes are immediate after pass  
> - `withdrawProfitPeriod` is in blocks (~1s/block)

## 5. Validator Fee Withdrawal

### 5.1 Build Tx

```shell
./build/congress-cli withdraw_profits -a <validator_address>
```

### 5.2 Sign & Send

```shell
./build/congress-cli sign -f withdrawProfits.json -k miner1.key -p password.file
./build/congress-cli send -f withdrawProfits_signed.json
```

> No voting needed; validators can withdraw directly.

## 6. Staking Commands

### 6.1 Overview

```shell
./build/congress-cli staking --help
```

Subcommands:
1) `register-validator`
2) `delegate`
3) `undelegate`
4) `query-validator`
5) `list-top-validators`
6) `unjail`
7) `withdraw`

### 6.2 Register Validator

```shell
./build/congress-cli staking register-validator \
  --proposer <proposer> \
  --stake-amount <amount> \
  --commission-rate <bps>
```

- `--stake-amount` min 10000 JU  
- `--commission-rate` 0–10000 (bps)

Examples:

```shell
./build/congress-cli staking register-validator \
  --proposer 0x3F9DDeBE20b24B0DEC1d0B5A3c6e8Cb8D3eCF6A5 \
  --stake-amount 10000 \
  --commission-rate 500

./build/congress-cli staking register-validator \
  --proposer 0x3F9DDeBE20b24B0DEC1d0B5A3c6e8Cb8D3eCF6A5 \
  --stake-amount 10000 \
  --commission-rate 500 \
  --rpc https://rpc.juchain.org --chainId 210000
```

### 6.3 Delegate

```shell
./build/congress-cli staking delegate \
  --validator <validator> \
  --amount <amount>
```

### 6.4 Undelegate

```shell
./build/congress-cli staking undelegate \
  --validator <validator> \
  --amount <amount>
```

### 6.5 Query Validator

```shell
./build/congress-cli staking query-validator --address <validator>
```

> Note: Some query commands may hit ABI parsing errors; known issue under fix.

### 6.6 List Top Validators

```shell
./build/congress-cli staking list-top-validators [count]
```

### 6.7 Unjail

```shell
./build/congress-cli staking unjail --validator <validator>
```

### 6.9 Edit Validator Info

```shell
./build/congress-cli staking edit-validator \
  --validator <validator> \
  --fee-addr <fee_address> \
  [--moniker <name>] \
  [--identity <id>] \
  [--website <url>] \
  [--email <email>] \
  [--details <text>]
```

Example:

```shell
./build/congress-cli staking edit-validator \
  --validator 0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc \
  --fee-addr 0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc \
  --moniker "validator6" \
  --details "Validator6 node with fee address configured"

./build/congress-cli sign \
  -f editValidator.json \
  -k $(find $HOME/ju-chain-work/chain/private-chain/data-validator6/keystore/ -name "*--9965507d1a55bcc2695c58ba16fb37d819b0a4dc" | head -1) \
  -p $HOME/ju-chain-work/chain/private-chain/data-validator6/password.txt

./build/congress-cli send -f editValidator_signed.json
```

> **Notes**  
> - Only the validator can edit their info  
> - Fee address receives block rewards  
> - If fee address is `0x0000...`, set a real address to receive rewards

### 6.10 Withdraw Delegation Rewards

```shell
./build/congress-cli staking withdraw --validator <validator>
```

## 7. End-to-End: Add Validator 6

### 7.1 Prereqs

```bash
PRIVATE_CHAIN_PATH=$HOME/ju-chain-work/chain/private-chain

VALIDATOR1_KEY=$(find ${PRIVATE_CHAIN_PATH}/data-validator1/keystore/ -name "*--f39fd6e51aad88f6f4ce6ab8827279cfffb92266" | head -1)
VALIDATOR2_KEY=$(find ${PRIVATE_CHAIN_PATH}/data-validator2/keystore/ -name "*--70997970c51812dc3a010c7d01b50e0d17dc79c8" | head -1)
VALIDATOR3_KEY=$(find ${PRIVATE_CHAIN_PATH}/data-validator3/keystore/ -name "*--3c44cdddb6a900fa2b585dd299e03d12fa4293bc" | head -1)
VALIDATOR4_KEY=$(find ${PRIVATE_CHAIN_PATH}/data-validator4/keystore/ -name "*--90f79bf6eb2c4f870365e785982e1f101e93b906" | head -1)
VALIDATOR5_KEY=$(find ${PRIVATE_CHAIN_PATH}/data-validator5/keystore/ -name "*--15d34aaf54267db7d7c367839aaf71a00a2c6a65" | head -1)

PASSWORD_FILE=${PRIVATE_CHAIN_PATH}/data-validator1/password.txt
```

Passwords are `123456`.

### 7.2 Step 1: Validator1 Creates Proposal

```shell
./build/congress-cli create_proposal \
  -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  -t 0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc \
  -o add

./build/congress-cli sign \
  -f createProposal.json \
  -k $(find $HOME/ju-chain-work/chain/private-chain/data-validator1/keystore/ -name "*--f39fd6e51aad88f6f4ce6ab8827279cfffb92266" | head -1) \
  -p $HOME/ju-chain-work/chain/private-chain/data-validator1/password.txt

./build/congress-cli send -f createProposal_signed.json
```

Record the proposal ID.

### 7.3 Step 2: Five Validators Vote

Replace `PROPOSAL_ID` with the real ID.

```shell
# v1
./build/congress-cli vote_proposal -s 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k $(find $HOME/ju-chain-work/chain/private-chain/data-validator1/keystore/ -name "*--f39fd6e51aad88f6f4ce6ab8827279cfffb92266" | head -1) -p $HOME/ju-chain-work/chain/private-chain/data-validator1/password.txt
./build/congress-cli send -f voteProposal_signed.json

# v2
./build/congress-cli vote_proposal -s 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k $(find $HOME/ju-chain-work/chain/private-chain/data-validator2/keystore/ -name "*--70997970c51812dc3a010c7d01b50e0d17dc79c8" | head -1) -p $HOME/ju-chain-work/chain/private-chain/data-validator2/password.txt
./build/congress-cli send -f voteProposal_signed.json

# v3
./build/congress-cli vote_proposal -s 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k $(find $HOME/ju-chain-work/chain/private-chain/data-validator3/keystore/ -name "*--3c44cdddb6a900fa2b585dd299e03d12fa4293bc" | head -1) -p $HOME/ju-chain-work/chain/private-chain/data-validator3/password.txt
./build/congress-cli send -f voteProposal_signed.json

# v4
./build/congress-cli vote_proposal -s 0x90F79bf6EB2c4f870365E785982E1f101E93b906 -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k $(find $HOME/ju-chain-work/chain/private-chain/data-validator4/keystore/ -name "*--90f79bf6eb2c4f870365e785982e1f101e93b906" | head -1) -p $HOME/ju-chain-work/chain/private-chain/data-validator4/password.txt
./build/congress-cli send -f voteProposal_signed.json

# v5
./build/congress-cli vote_proposal -s 0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65 -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k $(find $HOME/ju-chain-work/chain/private-chain/data-validator5/keystore/ -name "*--15d34aaf54267db7d7c367839aaf71a00a2c6a65" | head -1) -p $HOME/ju-chain-work/chain/private-chain/data-validator5/password.txt
./build/congress-cli send -f voteProposal_signed.json
```

### 7.4 Step 3: Verify

```shell
./build/congress-cli proposal -i PROPOSAL_ID
./build/congress-cli miner -a 0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc
./build/congress-cli miners
```

Expected: proposal Passed/Executed; validator6 listed; total validators 6.

### 7.5 Automation Script Example

```bash
#!/bin/bash
TARGET_ADDRESS="0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc"
PROPOSER_ADDRESS="0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
PRIVATE_CHAIN_PATH="$HOME/ju-chain-work/chain/private-chain"

VALIDATORS=(
  "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
  "0x70997970C51812dc3A010C7d01b50e0d17dc79C8"
  "0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC"
  "0x90F79bf6EB2c4f870365E785982E1f101E93b906"
  "0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65"
)

KEYS=(
  "$(find ${PRIVATE_CHAIN_PATH}/data-validator1/keystore/ -name "*--f39fd6e51aad88f6f4ce6ab8827279cfffb92266" | head -1)"
  "$(find ${PRIVATE_CHAIN_PATH}/data-validator2/keystore/ -name "*--70997970c51812dc3a010c7d01b50e0d17dc79c8" | head -1)"
  "$(find ${PRIVATE_CHAIN_PATH}/data-validator3/keystore/ -name "*--3c44cdddb6a900fa2b585dd299e03d12fa4293bc" | head -1)"
  "$(find ${PRIVATE_CHAIN_PATH}/data-validator4/keystore/ -name "*--90f79bf6eb2c4f870365e785982e1f101e93b906" | head -1)"
  "$(find ${PRIVATE_CHAIN_PATH}/data-validator5/keystore/ -name "*--15d34aaf54267db7d7c367839aaf71a00a2c6a65" | head -1)"
)

PASSWORDS=(
  "${PRIVATE_CHAIN_PATH}/data-validator1/password.txt"
  "${PRIVATE_CHAIN_PATH}/data-validator2/password.txt"
  "${PRIVATE_CHAIN_PATH}/data-validator3/password.txt"
  "${PRIVATE_CHAIN_PATH}/data-validator4/password.txt"
  "${PRIVATE_CHAIN_PATH}/data-validator5/password.txt"
)

echo "=== Step 1: create proposal ==="
./build/congress-cli create_proposal -p $PROPOSER_ADDRESS -t $TARGET_ADDRESS -o add
./build/congress-cli sign -f createProposal.json -k "${KEYS[0]}" -p "${PASSWORDS[0]}"
./build/congress-cli send -f createProposal_signed.json

echo "Enter proposal ID:"
read PROPOSAL_ID

echo "=== Step 2: votes ==="
for i in "${!VALIDATORS[@]}"; do
  echo "Validator $((i+1)) voting..."
  ./build/congress-cli vote_proposal -s ${VALIDATORS[$i]} -i $PROPOSAL_ID -a
  ./build/congress-cli sign -f voteProposal.json -k "${KEYS[$i]}" -p "${PASSWORDS[$i]}"
  ./build/congress-cli send -f voteProposal_signed.json
done

echo "=== Step 3: verify ==="
./build/congress-cli proposal -i $PROPOSAL_ID
./build/congress-cli miner -a $TARGET_ADDRESS
./build/congress-cli miners
```

## 8. Full Test Transaction Example

```shell
# 1) create proposal
./build/congress-cli create_proposal \
  --proposal "test proposal" \
  --action 0 \
  --value 1000 \
  --proposer 0x3F9DDeBE20b24B0DEC1d0B5A3c6e8Cb8D3eCF6A5

# 2) vote (4 validators approve)
./build/congress-cli vote_proposal --proposer 0x3F9DDeBE20b24B0DEC1d0B5A3c6e8Cb8D3eCF6A5 --id 1 --vote 1
./build/congress-cli vote_proposal --proposer 0x4B1E2D4D7C8F5A9B6E3F8A2C7D9E1F4B8C5E6A9 --id 1 --vote 1
./build/congress-cli vote_proposal --proposer 0x5C2F3E5E8D9A6B7C4E5F9B3D8E2A5C9D6F7A8E --id 1 --vote 1
./build/congress-cli vote_proposal --proposer 0x6D3A4F6F9E8B7C5D6A7C8F4E9A3B6D7E8F9C1A --id 1 --vote 1

# 3) query proposal
./build/congress-cli query_proposal --id 1

# 4) delegate
./build/congress-cli staking delegate \
  --validator 0x3F9DDeBE20b24B0DEC1d0B5A3c6e8Cb8D3eCF6A5 \
  --amount 1000
```

## 9. Mainnet: Restore Miner Identity

Miner1 proposes adding `0x029DAB47e268575D4AC167De64052FB228B5fA41`; miner1/2/3 vote yes.

```shell
# step1 create/sign/send
./build/congress-cli create_proposal -p 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b -t 0x029DAB47e268575D4AC167De64052FB228B5fA41 -o add
./build/congress-cli sign -f createProposal.json -k miner1.key -p password.file
./build/congress-cli send -f createProposal_signed.json
# record proposal ID (e.g., b2be7f3c...)

# step2 votes (replace PROPOSAL_ID)
./build/congress-cli vote_proposal -s 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k miner1.key -p password.file
./build/congress-cli send -f voteProposal_signed.json

./build/congress-cli vote_proposal -s 0x81f7a79a51edba249efa812eb2d5478f696f7558 -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k miner2.key -p password.file
./build/congress-cli send -f voteProposal_signed.json

./build/congress-cli vote_proposal -s 0x578c39eaf09a4e1abf428c423970b59bb8baf42e -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k miner3.key -p password.file
./build/congress-cli send -f voteProposal_signed.json

# step3 check
./build/congress-cli miner -a 0x029DAB47e268575D4AC167De64052FB228B5fA41
./build/congress-cli miners
```

## 10. Mainnet: Change Config

### 10.1 Create Config Proposal

```shell
# CIDs: 0 proposalLastingPeriod, 1 punishThreshold, 2 removeThreshold, 3 decreaseRate, 4 withdrawProfitPeriod
./build/congress-cli create_config_proposal -p <proposer> -i <cid> -v <value>

./build/congress-cli create_config_proposal -p 0xccafa71c31bc11ba24d526fd27ba57d743152807 -i 0 -v 86400
./build/congress-cli create_config_proposal -p 0xccafa71c31bc11ba24d526fd27ba57d743152807 -i 4 -v 10

./build/congress-cli sign -f createUpdateConfigProposal.json -k miner1.key -p password.file
./build/congress-cli send -f createUpdateConfigProposal_signed.json
# record proposal ID
```

### 10.2 Validators Vote

```shell
./build/congress-cli vote_proposal -s 0xccafa71c31bc11ba24d526fd27ba57d743152807 -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k miner1.key -p password.file
./build/congress-cli send -f voteProposal_signed.json

./build/congress-cli vote_proposal -s 0x81f7a79a51edba249efa812eb2d5478f696f7558 -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k miner2.key -p password.file
./build/congress-cli send -f voteProposal_signed.json

./build/congress-cli vote_proposal -s 0x578c39eaf09a4e1abf428c423970b59bb8baf42e -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k miner3.key -p password.file
./build/congress-cli send -f voteProposal_signed.json
```

### 10.3 Check Config Proposal Status

```shell
./build/congress-cli proposal -i PROPOSAL_ID
./build/congress-cli proposals
```

> Config changes take effect immediately; `withdrawProfitPeriod` is in blocks. Set carefully.

## 11. Mainnet: Withdraw Validator Fees

```shell
./build/congress-cli withdraw_profits -a <miner_address>
./build/congress-cli sign -f withdrawProfits.json -k miner.key -p password.file
./build/congress-cli send -f withdrawProfits_signed.json
```

## 12. Tool Info

### 12.1 Version

```shell
./build/congress-cli version
```

### 12.2 Help

```shell
./build/congress-cli help
./build/congress-cli [command] --help
```

## 13. Notes

### 13.1 Reminders

- ⚠️ Only active validators can propose/vote  
- ⚠️ Ensure node is fully synced before restoring miner status  
- ⚠️ Each action generates a new proposal ID—use the correct one  
- ⚠️ Protect keystore and password files

### 13.2 Common Errors

1. `"Validator only"`: account is not an active validator  
2. `"You can't vote for a proposal twice"`: duplicate vote  
3. `"gas estimation failed"`: bad params or network issue

### 13.3 System Contract Addresses

- Validators: `0x000000000000000000000000000000000000f000`
- Punish: `0x000000000000000000000000000000000000f001`
- Proposal: `0x000000000000000000000000000000000000f002`

### 13.4 Network Info

- Testnet: `https://testnet-rpc.juchain.org` (Chain ID: 202599)
- Mainnet: `https://rpc.juchain.org` (Chain ID: 210000)
# Congress POSA Consensus Management Tool Usage Guide

## Test Environment Account Configuration

To demonstrate the complete voting process, this document uses the following pre-configured test accounts:

```bash
# Validator 1 Account Configuration
VALIDATOR1_ADDRESS=0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
VALIDATOR1_PRIVATE_KEY=ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
VALIDATOR1_PASSWORD=123456

# Validator 2 Account Configuration
VALIDATOR2_ADDRESS=0x70997970C51812dc3A010C7d01b50e0d17dc79C8
VALIDATOR2_PRIVATE_KEY=59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d
VALIDATOR2_PASSWORD=123456

# Validator 3 Account Configuration
VALIDATOR3_ADDRESS=0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC
VALIDATOR3_PRIVATE_KEY=5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a
VALIDATOR3_PASSWORD=123456

# Validator 4 Account Configuration
VALIDATOR4_ADDRESS=0x90F79bf6EB2c4f870365E785982E1f101E93b906
VALIDATOR4_PRIVATE_KEY=7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6
VALIDATOR4_PASSWORD=123456

# Validator 5 Account Configuration
VALIDATOR5_ADDRESS=0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65
VALIDATOR5_PRIVATE_KEY=47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a
VALIDATOR5_PASSWORD=123456

# Validator 6 Account Configuration (to be added as new validator)
VALIDATOR6_ADDRESS=0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc
VALIDATOR6_PRIVATE_KEY=8b3a350cf5c34c9194ca85829a2df0ec3153be0318b5e2d3348e872092edffba
VALIDATOR6_PASSWORD=123456
```

> **⚠️ Important**: The above private keys are for testing environment only. Do not use them in production environment!

## Private Chain Environment Key File Configuration

In the private chain environment, the validators' key files have been pre-configured in the following locations:

```bash
# Private Chain Data Directory
PRIVATE_CHAIN_PATH=$HOME/ju-chain-work/chain/private-chain

# Validator 1 Key File (using wildcard matching)
VALIDATOR1_KEYSTORE=${PRIVATE_CHAIN_PATH}/data-validator1/keystore/UTC--*--f39fd6e51aad88f6f4ce6ab8827279cfffb92266
VALIDATOR1_PASSWORD=${PRIVATE_CHAIN_PATH}/data-validator1/password.txt

# Validator 2 Key File (using wildcard matching)
VALIDATOR2_KEYSTORE=${PRIVATE_CHAIN_PATH}/data-validator2/keystore/UTC--*--70997970c51812dc3a010c7d01b50e0d17dc79c8
VALIDATOR2_PASSWORD=${PRIVATE_CHAIN_PATH}/data-validator2/password.txt

# Validator 3 Key File (using wildcard matching)
VALIDATOR3_KEYSTORE=${PRIVATE_CHAIN_PATH}/data-validator3/keystore/UTC--*--3c44cdddb6a900fa2b585dd299e03d12fa4293bc
VALIDATOR3_PASSWORD=${PRIVATE_CHAIN_PATH}/data-validator3/password.txt

# Validator 4 Key File (using wildcard matching)
VALIDATOR4_KEYSTORE=${PRIVATE_CHAIN_PATH}/data-validator4/keystore/UTC--*--90f79bf6eb2c4f870365e785982e1f101e93b906
VALIDATOR4_PASSWORD=${PRIVATE_CHAIN_PATH}/data-validator4/password.txt

# Validator 5 Key File (using wildcard matching)
VALIDATOR5_KEYSTORE=${PRIVATE_CHAIN_PATH}/data-validator5/keystore/UTC--*--15d34aaf54267db7d7c367839aaf71a00a2c6a65
VALIDATOR5_PASSWORD=${PRIVATE_CHAIN_PATH}/data-validator5/password.txt

# Validator 6 Key File (using wildcard matching)
VALIDATOR6_KEYSTORE=${PRIVATE_CHAIN_PATH}/data-validator6/keystore/UTC--*--340d92a853ae20a6e7a5b86272fa47aff83a8f7a
VALIDATOR6_PASSWORD=${PRIVATE_CHAIN_PATH}/data-validator6/password.txt
```

All validators' passwords are `123456`, stored in their respective `password.txt` files.

## Flexible Key File Lookup

To avoid hardcoding specific timestamps, we provide flexible key file lookup methods:

### Method 1: Using find command

```bash
# Get validator 1 key file
VALIDATOR1_KEY=$(find $HOME/ju-chain-work/chain/private-chain/data-validator1/keystore/ -name "*--f39fd6e51aad88f6f4ce6ab8827279cfffb92266" | head -1)

# Get validator 2 key file
VALIDATOR2_KEY=$(find $HOME/ju-chain-work/chain/private-chain/data-validator2/keystore/ -name "*--70997970c51812dc3a010c7d01b50e0d17dc79c8" | head -1)
```

### Method 2: Using wildcards

```bash
# Using wildcards in shell scripts
VALIDATOR1_KEY=$HOME/ju-chain-work/chain/private-chain/data-validator1/keystore/*--f39fd6e51aad88f6f4ce6ab8827279cfffb92266
VALIDATOR2_KEY=$HOME/ju-chain-work/chain/private-chain/data-validator2/keystore/*--70997970c51812dc3a010c7d01b50e0d17dc79c8
```

### Method 3: Utility function

```bash
# Define function to get key file
get_validator_key() {
    local validator_num=$1
    local address=$(echo $2 | tr '[:upper:]' '[:lower:]')
    find $HOME/ju-chain-work/chain/private-chain/data-validator${validator_num}/keystore/ -name "*--${address}" | head -1
}

# Usage example
VALIDATOR1_KEY=$(get_validator_key 1 f39fd6e51aad88f6f4ce6ab8827279cfffb92266)
VALIDATOR2_KEY=$(get_validator_key 2 70997970c51812dc3a010c7d01b50e0d17dc79c8)
```

> **Advantages**: This approach does not depend on specific timestamps. As long as the address matches, the corresponding key file can be found, which is more flexible and maintainable.

## Tool Compilation

First, you need to compile the congress-cli tool:

```shell
cd sys-contract/congress-cli
make build
# The generated executable is located at build/congress-cli
```

## Version Information

This document has been updated to **Congress CLI v1.2.1**, build date: 2025-08-27.

```shell
./build/congress-cli version
# Output: Congress CLI Version: 1.2.1, Build Date: 2025-08-27
```

> **💡 Syntax Description**:
>
> - Local testing: Default connection `http://127.0.0.1:8545`, chain ID `202599`
> - Testnet: `--chainId 202599 --rpc_laddr https://testnet-rpc.juchain.org`
> - Mainnet: `--chainId 210000 --rpc_laddr https://rpc.juchain.org`

## Configuration Management Commands

### View Configuration Help

Congress CLI v1.2.0 provides convenient configuration query and setting functions.

```shell
# View config subcommand help
./build/congress-cli config --help
```

### Query System Configuration

Query current system configuration parameters:

```shell
# Query all configurations
./build/congress-cli config get

# Query RPC endpoint
./build/congress-cli config get --rpc

# Query chain ID
./build/congress-cli config get --chain-id
```

**Example**:

```shell
# Query all configurations
./build/congress-cli config get
# Output:
# RPC endpoint: https://rpc.juchain.org
# Chain ID: 210000

# Query RPC endpoint only
./build/congress-cli config get --rpc
# Output: RPC endpoint: https://rpc.juchain.org

# Query chain ID only
./build/congress-cli config get --chain-id
# Output: Chain ID: 210000
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

> **Note**: The config command here is used to set the configuration of the congress-cli tool itself (such as RPC endpoint, chain ID), not to modify blockchain system parameters. For blockchain system parameter modification, please refer to Chapter 4.

## 1. Create Proposal

### 1.1. Create Original Transaction

```shell
# Basic syntax
./build/congress-cli create_proposal -p proposer address -t new miner address -o add

# Complete example: Validator 1 creates an add proposal for validator 6
./build/congress-cli create_proposal -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 -t 0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc -o add

# Other example: Remove validator
./build/congress-cli create_proposal -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 -t 0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc -o remove
```

**Parameter Description**:

- `-p, --proposer`: Proposer address (must be a valid validator)
- `-t, --target`: Target address (validator to add or remove)
- `-o, --operation`: Operation type (add or remove)

> A `createProposal.json` file will be generated after successful execution

### 1.2. Sign Transaction

```shell
./build/congress-cli sign -f createProposal.json -k wallet file -p wallet password file

# Validator 1 signature example
./build/congress-cli sign -f createProposal.json -k validator1.key -p password.file
```

> A `createProposal_signed.json` file will be generated after successful execution

### 1.3. Send Transaction

```shell
./build/congress-cli send -f createProposal_signed.json
```

> After successful execution, proposal information will be output, including the important proposal ID:

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

> **⚠️ Important**: Record the proposal ID, it will be needed for voting!

## 2. Proposal Voting

### 2.1. Create Voting Transaction

Now multiple validators need to vote on the proposal. Assume the proposal ID is `0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c`

```shell
# Basic syntax
./build/congress-cli vote_proposal -s signer address -i proposal ID -a  # Approval vote
./build/congress-cli vote_proposal -s signer address -i proposal ID     # Rejection vote

# Validator 1 vote (proposer must also vote)
./build/congress-cli vote_proposal -s 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 -i 0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c -a

# Validator 2 vote
./build/congress-cli vote_proposal -s 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 -i 0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c -a

# Validator 3 vote
./build/congress-cli vote_proposal -s 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC -i 0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c -a

# Validator 4 vote
./build/congress-cli vote_proposal -s 0x90F79bf6EB2c4f870365E785982E1f101E93b906 -i 0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c -a

# Validator 5 vote
./build/congress-cli vote_proposal -s 0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65 -i 0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c -a
```

**Parameter Description**:

- `-s, --signer`: Signer address (must be a valid validator)
- `-i, --proposalId`: Proposal ID (64-character hexadecimal string)
- `-a, --approve`: Approval flag (use -a for YES, omit for NO)

> A `voteProposal.json` file will be generated after successful execution

### 2.2. Sign Transaction

Each vote needs to be signed with the corresponding validator's private key:

```shell
# Validator 1 signature
./build/congress-cli sign -f voteProposal.json -k validator1.key -p password.file

# Validator 2 signature  
./build/congress-cli sign -f voteProposal.json -k validator2.key -p password.file

# Validator 3 signature
./build/congress-cli sign -f voteProposal.json -k validator3.key -p password.file

# Validator 4 signature
./build/congress-cli sign -f voteProposal.json -k validator4.key -p password.file

# Validator 5 signature
./build/congress-cli sign -f voteProposal.json -k validator5.key -p password.file
```

> A `voteProposal_signed.json` file will be generated after successful execution

### 2.3. Send Voting Transaction

Each validator needs to send their own voting transaction:

```shell
# Send each validator's voting transaction in turn
./build/congress-cli send -f voteProposal_signed.json
```

> Confirmation information will be output after successful execution:

```text
✅ Transaction confirmed in block 8830
read sender from signed tx is 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
✅ Transaction broadcast successfully!
```

### 2.4. Complete Voting Process Example

The following is the complete voting process to add validator 6 as a new validator:

```shell
# Step 1: Validator 1 vote
./build/congress-cli vote_proposal -s 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k validator1.key -p password.file
./build/congress-cli send -f voteProposal_signed.json

# Step 2: Validator 2 vote
./build/congress-cli vote_proposal -s 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k validator2.key -p password.file
./build/congress-cli send -f voteProposal_signed.json

# Step 3: Validator 3 vote
./build/congress-cli vote_proposal -s 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k validator3.key -p password.file
./build/congress-cli send -f voteProposal_signed.json

# Step 4: Validator 4 vote
./build/congress-cli vote_proposal -s 0x90F79bf6EB2c4f870365E785982E1f101E93b906 -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k validator4.key -p password.file
./build/congress-cli send -f voteProposal_signed.json

# Step 5: Validator 5 vote
./build/congress-cli vote_proposal -s 0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65 -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k validator5.key -p password.file
./build/congress-cli send -f voteProposal_signed.json
```

> **Note**:
>
> - Each validator can only vote once on the same proposal
> - Sufficient validators need to vote in favor for the proposal to pass
> - Please replace "PROPOSAL_ID" with the actual proposal ID in the above commands

## 3. Query Operations

### 3.1 Query All Active Miners

```shell
./build/congress-cli miners
```

> Output example:

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
./build/congress-cli miner -a <validator address>

# Example
./build/congress-cli miner -a 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
```

> Output example:

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
- Status: Inactive ❌ = Abnormal status

### 3.3 Query All Proposals

```shell
./build/congress-cli proposals
```

> Output example:

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
./build/congress-cli proposal -i <proposal ID>

# Example
./build/congress-cli proposal -i 0x0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c
```

> Output example:

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
./build/congress-cli create_config_proposal -p <proposer address> -i <configuration item ID> -v <configuration value>

# Example: Modify proposal duration to 86400 seconds
./build/congress-cli create_config_proposal -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 -i 0 -v 86400

# Example: Modify withdrawal cooldown period to 10 blocks (about 10 seconds)
./build/congress-cli create_config_proposal -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 -i 4 -v 10
```

**Configuration Item ID Mapping Table:**

- 0: proposalLastingPeriod (proposal duration, unit: seconds)
- 1: punishThreshold (punishment threshold)
- 2: removeThreshold (removal threshold)
- 3: decreaseRate (decrease rate)
- 4: withdrawProfitPeriod (profit withdrawal period, unit: blocks)

**Parameter Description:**

- `-p, --proposer`: Proposer address (must be a valid validator)
- `-i, --cid`: Configuration item ID (0-4)
- `-v, --value`: New value of configuration item

> A `createUpdateConfigProposal.json` file will be generated after successful execution

### 4.2 Sign and Send Transaction

The signing and sending process for configuration modification proposals is the same as for regular proposals:

```shell
# Sign transaction (note the filename is createUpdateConfigProposal.json)
./build/congress-cli sign -f createUpdateConfigProposal.json -k /path/to/validator.key -p password.file

# Send transaction
./build/congress-cli send -f createUpdateConfigProposal_signed.json
```

### 4.3 Complete Configuration Modification Process Example

The following is the complete process for modifying the withdrawal cooldown period:

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

# Step 4: Record proposal ID (obtained from output)
# Example proposal ID from output: 0xd87a55165c909c9b4ef949a3d697e3b26a6a66eee38b2ed519f52f8acd342539

# Step 5: Validator voting (requires sufficient validator votes)
./build/congress-cli vote_proposal -s 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 -i 0xd87a55165c909c9b4ef949a3d697e3b26a6a66eee38b2ed519f52f8acd342539 -a
./build/congress-cli sign -f voteProposal.json -k $(find $HOME/ju-chain-work/chain/private-chain/data-validator1/keystore/ -name "*--f39fd6e51aad88f6f4ce6ab8827279cfffb92266" | head -1) -p password.file
./build/congress-cli send -f voteProposal_signed.json

# Repeat voting process for other validators...

# Step 6: Check proposal status
./build/congress-cli proposal -i 0xd87a55165c909c9b4ef949a3d697e3b26a6a66eee38b2ed519f52f8acd342539
```

### 4.4 Common Configuration Modification Scenarios

#### Scenario 1: Shorten Withdrawal Cooldown Period (for testing)

```shell
# Change withdrawal cooldown period from default 86400 blocks (24 hours) to 10 blocks (about 10 seconds)
./build/congress-cli create_config_proposal -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 -i 4 -v 10
```

#### Scenario 2: Adjust Proposal Duration

```shell
# Change proposal duration to 7 days (604800 seconds)
./build/congress-cli create_config_proposal -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 -i 0 -v 604800
```

> **Note:**
>
> - Configuration modification proposals also require sufficient validator votes to pass and execute
> - Configuration changes take effect immediately, please set parameter values carefully
> - The unit of withdrawProfitPeriod is blocks, assuming 1 block per second for time calculation

## 5. Miner Profit Withdrawal

### 5.1 Create Withdrawal Transaction

```shell
./build/congress-cli withdraw_profits -a <miner address>

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

> **Note:** Profit withdrawal does not require a voting process, miners can directly withdraw their own profits

## 6. Staking Operations

### 6.1 Staking Command Overview

Congress CLI v1.2.0 has added a complete staking function module that supports validator registration, delegation, querying and other operations.

```shell
# View staking subcommand help
./build/congress-cli staking --help
```

**Available Staking Subcommands:**

1. `register-validator` - Register validator
2. `delegate` - Delegate stake to validator
3. `undelegate` - Cancel delegation
4. `query-validator` - Query specified validator information
5. `list-top-validators` - List top validators
6. `unjail` - Unjail validator
7. `withdraw` - Withdraw delegation rewards

### 6.2 Register Validator

Registering as a new validator requires meeting the minimum staking requirements.

```shell
./build/congress-cli staking register-validator \
  --proposer <proposer address> \
  --stake-amount <stake amount> \
  --commission-rate <commission rate>
```

**Parameter Description**:

- `--proposer`: Proposer address (required)
- `--stake-amount`: JU stake amount (required, minimum 10000 JU)
- `--commission-rate`: Commission rate in basis points (0-10000, e.g., 500 represents 5%)

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
  --validator <validator address> \
  --amount <delegate amount>
```

**Example**:

```shell
# Delegate 1000 JU to validator
./build/congress-cli staking delegate \
  --validator 0x3F9DDeBE20b24B0DEC1d0B5A3c6e8Cb8D3eCF6A5 \
  --amount 1000
```

### 6.4 Undelegate

Cancel previous delegation, need to wait for unbonding period.

```shell
./build/congress-cli staking undelegate \
  --validator <validator address> \
  --amount <undelegate amount>
```

**Example**:

```shell
# Undelegate 500 JU
./build/congress-cli staking undelegate \
  --validator 0x3F9DDeBE20b24B0DEC1d0B5A3c6e8Cb8D3eCF6A5 \
  --amount 500
```

### 6.5 Query Validator Information

Query detailed information of a specified validator, including stake amount, commission rate, status, etc.

```shell
./build/congress-cli staking query-validator --address <validator address>
```

**Example**:

```shell
# Query validator information
./build/congress-cli staking query-validator \
  --address 0x3F9DDeBE20b24B0DEC1d0B5A3c6e8Cb8D3eCF6A5
```

> **Note**: Some query commands in the current version may encounter ABI parsing errors, which is a known issue being fixed.

### 6.6 List Top Validators

View the list of currently active top validators.

```shell
./build/congress-cli staking list-top-validators [count]
```

**Example**:

```shell
# View top 15 validators (default)
./build/congress-cli staking list-top-validators

# View top 10 validators
./build/congress-cli staking list-top-validators 10
```

### 6.7 Unjail

After being jailed for violations, validators can apply to unjail themselves.

```shell
./build/congress-cli staking unjail --validator <validator address>
```

### 6.9 Edit Validator Information

Edit information of existing validators, including fee address and description information.

```shell
./build/congress-cli staking edit-validator \
  --validator <validator address> \
  --fee-addr <fee address> \
  [--moniker <display name>] \
  [--identity <identity identifier>] \
  [--website <website URL>] \
  [--email <contact email>] \
  [--details <detailed description>]
```

**Parameter Description**:

- `--validator`: Validator address to edit (required)
- `--fee-addr`: Fee address for receiving mining rewards (required)
- `--moniker`: Validator display name (optional)
- `--identity`: Validator identity identifier, such as Keybase signature (optional)
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

# Sign transaction (using validator6's own key)
./build/congress-cli sign \
  -f editValidator.json \
  -k $(find $HOME/ju-chain-work/chain/private-chain/data-validator6/keystore/ -name "*--9965507d1a55bcc2695c58ba16fb37d819b0a4dc" | head -1) \
  -p $HOME/ju-chain-work/chain/private-chain/data-validator6/password.txt

# Send transaction
./build/congress-cli send -f editValidator_signed.json
```

> **Note**:
>
> - Editing validator information can only be performed by the validator itself
> - Fee address is used to receive block mining rewards
> - If the validator's fee address is `0x0000...`, the correct fee address should be set in time to ensure rewards can be received

### 6.10 Withdraw Rewards

Withdraw rewards generated by delegation.

```shell
./build/congress-cli staking withdraw --validator <validator address>
```

## 7. End-to-End Process: Adding Validator 6

This chapter demonstrates the complete process from scratch, letting validators 1-5 vote for validator 6 to make it a new validator.

### 7.1 Prerequisites

Key files in the private chain environment have been pre-configured in the following locations:

```bash
# Set private chain path variable
PRIVATE_CHAIN_PATH=$HOME/ju-chain-work/chain/private-chain

# Validator key file paths (using wildcard pattern)
VALIDATOR1_KEY=$(find ${PRIVATE_CHAIN_PATH}/data-validator1/keystore/ -name "*--f39fd6e51aad88f6f4ce6ab8827279cfffb92266" | head -1)
VALIDATOR2_KEY=$(find ${PRIVATE_CHAIN_PATH}/data-validator2/keystore/ -name "*--70997970c51812dc3a010c7d01b50e0d17dc79c8" | head -1)
VALIDATOR3_KEY=$(find ${PRIVATE_CHAIN_PATH}/data-validator3/keystore/ -name "*--3c44cdddb6a900fa2b585dd299e03d12fa4293bc" | head -1)
VALIDATOR4_KEY=$(find ${PRIVATE_CHAIN_PATH}/data-validator4/keystore/ -name "*--90f79bf6eb2c4f870365e785982e1f101e93b906" | head -1)
VALIDATOR5_KEY=$(find ${PRIVATE_CHAIN_PATH}/data-validator5/keystore/ -name "*--15d34aaf54267db7d7c367839aaf71a00a2c6a65" | head -1)

# Password file path (all validators use the same password file)
PASSWORD_FILE=${PRIVATE_CHAIN_PATH}/data-validator1/password.txt
```

> **Note**: All validators' passwords are `123456`, and you can use the `password.txt` file from any validator directory.

### 7.2 Step 1: Validator 1 Creates Proposal

```shell
# Create proposal to add validator 6
./build/congress-cli create_proposal \
  -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  -t 0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc \
  -o add

# Validator 1 signs (using dynamically found key file path)
./build/congress-cli sign \
  -f createProposal.json \
  -k $(find $HOME/ju-chain-work/chain/private-chain/data-validator1/keystore/ -name "*--f39fd6e51aad88f6f4ce6ab8827279cfffb92266" | head -1) \
  -p $HOME/ju-chain-work/chain/private-chain/data-validator1/password.txt

# Send transaction
./build/congress-cli send -f createProposal_signed.json
```

> **Important**: Record the proposal ID from the output, for example: `0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c`

### 7.3 Step 2: 5 Validators Vote

Replace `PROPOSAL_ID` in the following commands with the actual proposal ID obtained in the first step.

```shell
# Validator 1 vote
./build/congress-cli vote_proposal -s 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 -i PROPOSAL_ID -a
./build/congress-cli sign \
  -f voteProposal.json \
  -k $(find $HOME/ju-chain-work/chain/private-chain/data-validator1/keystore/ -name "*--f39fd6e51aad88f6f4ce6ab8827279cfffb92266" | head -1) \
  -p $HOME/ju-chain-work/chain/private-chain/data-validator1/password.txt
./build/congress-cli send -f voteProposal_signed.json

# Validator 2 vote
./build/congress-cli vote_proposal -s 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 -i PROPOSAL_ID -a
./build/congress-cli sign \
  -f voteProposal.json \
  -k $(find $HOME/ju-chain-work/chain/private-chain/data-validator2/keystore/ -name "*--70997970c51812dc3a010c7d01b50e0d17dc79c8" | head -1) \
  -p $HOME/ju-chain-work/chain/private-chain/data-validator2/password.txt
./build/congress-cli send -f voteProposal_signed.json

# Validator 3 vote
./build/congress-cli vote_proposal -s 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC -i PROPOSAL_ID -a
./build/congress-cli sign \
  -f voteProposal.json \
  -k $(find $HOME/ju-chain-work/chain/private-chain/data-validator3/keystore/ -name "*--3c44cdddb6a900fa2b585dd299e03d12fa4293bc" | head -1) \
  -p $HOME/ju-chain-work/chain/private-chain/data-validator3/password.txt
./build/congress-cli send -f voteProposal_signed.json

# Validator 4 vote
./build/congress-cli vote_proposal -s 0x90F79bf6EB2c4f870365E785982E1f101E93b906 -i PROPOSAL_ID -a
./build/congress-cli sign \
  -f voteProposal.json \
  -k $(find $HOME/ju-chain-work/chain/private-chain/data-validator4/keystore/ -name "*--90f79bf6eb2c4f870365e785982e1f101e93b906" | head -1) \
  -p $HOME/ju-chain-work/chain/private-chain/data-validator4/password.txt
./build/congress-cli send -f voteProposal_signed.json

# Validator 5 vote
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

# Query validator 6 status
./build/congress-cli miner -a 0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc

# Query all validators list
./build/congress-cli miners
```

> **Expected Results**:
>
> - Proposal status should show as "Passed" or "Executed"
> - Validator 6 should appear in the active validator list
> - Total number of validators should change from 5 to 6

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

# Key file array (using dynamic lookup)
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

echo "=== Step 2: Validators Voting ==="
for i in "${!VALIDATORS[@]}"; do
    echo "Validator $((i+1)) voting..."
    ./build/congress-cli vote_proposal -s ${VALIDATORS[$i]} -i $PROPOSAL_ID -a
    ./build/congress-cli sign -f voteProposal.json -k "${KEYS[$i]}" -p "${PASSWORDS[$i]}"
    ./build/congress-cli send -f voteProposal_signed.json
    echo "Validator $((i+1)) vote completed"
done

echo "=== Step 3: Verify Results ==="
./build/congress-cli proposal -i $PROPOSAL_ID
./build/congress-cli miner -a $TARGET_ADDRESS
./build/congress-cli miners
```

## 8. Complete Test Transaction Example

> **The following is a complete test process example:**

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

> miner1 creates a proposal to add 0x029DAB47e268575D4AC167De64052FB228B5fA41 as a new miner. After creating the proposal, miner1, miner2, and miner3 vote to approve it.

```shell
# step1 Create proposal transaction, sign and send
./build/congress-cli create_proposal -p 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b -t 0x029DAB47e268575D4AC167De64052FB228B5fA41 -o add
./build/congress-cli sign -f createProposal.json -k miner1.key -p password.file
./build/congress-cli send -f createProposal_signed.json
# After executing this command, you can obtain the proposal ID, for example: b2be7f3cc702c7a24962df6aed188edbcfdebe20fd55f5670efaedace0e4bcdf

# step2 3 miners vote on the proposal (please replace PROPOSAL_ID with the actual proposal ID obtained in the previous step)
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

# step3 View information of the newly added miner
./build/congress-cli miner -a 0x029DAB47e268575D4AC167De64052FB228B5fA41
./build/congress-cli miners
```

## 10. Mainnet Configuration Modification

### 10.1 Create Configuration Proposal

```shell
# Configuration item information corresponding to configuration item IDs
# 0 proposalLastingPeriod, 1 punishThreshold, 2 removeThreshold, 3 decreaseRate, 4 withdrawProfitPeriod
./build/congress-cli create_config_proposal -p proposer address -i configuration item ID -v configuration item value

# Example: Modify proposalLastingPeriod to 86400 seconds
./build/congress-cli create_config_proposal -p 0xccafa71c31bc11ba24d526fd27ba57d743152807 -i 0 -v 86400

# Example: Modify withdrawProfitPeriod to 10 blocks
./build/congress-cli create_config_proposal -p 0xccafa71c31bc11ba24d526fd27ba57d743152807 -i 4 -v 10

# Sign transaction (note the filename is createUpdateConfigProposal.json)
./build/congress-cli sign -f createUpdateConfigProposal.json -k miner1.key -p password.file

# Send transaction
./build/congress-cli send -f createUpdateConfigProposal_signed.json
# After executing this command, you can obtain the proposal ID. Record the proposal ID for subsequent voting
```

### 10.2 Validator Voting

The voting process for configuration proposals is the same as for adding validator proposals:

```shell
# Example: Vote on configuration proposal (replace PROPOSAL_ID with the actual proposal ID)
# miner1 vote
./build/congress-cli vote_proposal -s 0xccafa71c31bc11ba24d526fd27ba57d743152807 -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k miner1.key -p password.file
./build/congress-cli send -f voteProposal_signed.json

# miner2 vote
./build/congress-cli vote_proposal -s 0x81f7a79a51edba249efa812eb2d5478f696f7558 -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k miner2.key -p password.file
./build/congress-cli send -f voteProposal_signed.json

# miner3 vote
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

> **Important Reminder:**
>
> - Once a configuration modification proposal passes, it takes effect immediately
> - The unit of withdrawProfitPeriod is blocks, not seconds
> - Please set configuration parameters carefully to avoid affecting normal network operation

## 11. Mainnet Miner Profit Withdrawal

```shell
# step1 Create original transaction
./build/congress-cli withdraw_profits -a miner address

# step2 Transaction signature
./build/congress-cli sign -f withdrawProfits.json -k miner.key -p password.file

# step3 Send transaction
./build/congress-cli send -f withdrawProfits_signed.json
```

## 12. Tool Information

### 12.1 Version Check

```shell
./build/congress-cli version
```

### 12.2 Help Information

```shell
./build/congress-cli help
./build/congress-cli [command] --help  # View help for specific command
```

## 13. Notes

### 13.1 Important Reminders

- ⚠️ **Validator Requirements**: Only currently active validators can create proposals and vote
- ⚠️ **Network Sync**: Ensure the node is fully synchronized to the latest state before recovering miner identity
- ⚠️ **Proposal ID**: A new proposal ID is generated for each operation, be sure to use the correct ID
- ⚠️ **Key Security**: Properly secure key files and password files

### 13.2 Common Errors

1. **"Validator only"**: Current account is not a valid validator
2. **"You can't vote for a proposal twice"**: This validator has already voted on this proposal
3. **"gas estimation failed"**: Transaction parameters error or network issue

### 13.3 System Contract Addresses

- **Validators**: `0x000000000000000000000000000000000000f000`
- **Punish**: `0x000000000000000000000000000000000000f001`
- **Proposal**: `0x000000000000000000000000000000000000f002`

### 13.4 Network Information

- **Testnet**: `https://testnet-rpc.juchain.org` (Chain ID: 202599)
- **Mainnet**: `https://rpc.juchain.org` (Chain ID: 210000)
