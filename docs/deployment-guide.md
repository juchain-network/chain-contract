# JuChain PoSA Deployment and Operations Guide

## Scope

This guide reflects the current repository layout and the current PoSA implementation.

- this repository contains system contracts, Foundry tests, and genesis-generation scripts
- `ju-cli` and generated Go bindings have been moved to a standalone project
- the behavior described here is sourced from the Solidity contracts in `contracts/` and the Congress consensus implementation in `/home/litian/juchain/github/chain/consensus/congress`

If your external tooling exposes a different command surface, the contract function names and semantics in this document remain the source of truth.

## Table of Contents

1. [What Gets Deployed](#1-what-gets-deployed)
2. [Environment Setup and Build](#2-environment-setup-and-build)
3. [Genesis Generation and Validation](#3-genesis-generation-and-validation)
4. [Node Deployment and Network Layout](#4-node-deployment-and-network-layout)
5. [Validator and Governance Operations](#5-validator-and-governance-operations)
6. [Rewards, Staking, and Withdrawal Operations](#6-rewards-staking-and-withdrawal-operations)
7. [Monitoring and Day-2 Operations](#7-monitoring-and-day-2-operations)
8. [Troubleshooting Guide](#8-troubleshooting-guide)
9. [Contract Interface Summary](#9-contract-interface-summary)
10. [Best Practices Summary](#10-best-practices-summary)

## 1. What Gets Deployed

### 1.1 System Contracts

| Contract | Address | Responsibility |
| --- | --- | --- |
| `Validators` | `0x000000000000000000000000000000000000F010` | validator-set caches, validator metadata, fee-income accounting |
| `Punish` | `0x000000000000000000000000000000000000F011` | missed-block punishment and double-sign evidence |
| `Proposal` | `0x000000000000000000000000000000000000F012` | governance proposals and governable parameters |
| `Staking` | `0x000000000000000000000000000000000000F013` | validator self-stake, delegation, rewards, unbonding, jail state |

Historical POA addresses at `0xf000` to `0xf002` still exist in Congress for upgrade compatibility, but new PoSA deployments should use `0xf010` to `0xf013`.

### 1.2 Consensus Parameters

At genesis, Congress is configured under `config.congress`:

```json
{
  "config": {
    "congress": {
      "period": 1,
      "epoch": 86400
    }
  }
}
```

Meaning:

- `period`: target block interval in seconds
- `epoch`: validator-set rotation period in blocks

Validator-set changes only become effective at epoch boundaries.

### 1.3 Repository Boundaries

This repository owns:

- Solidity contracts
- Foundry tests
- contract bytecode generation
- genesis generation

This repository does not own:

- validator-node binaries
- long-running node orchestration
- external CLI wrappers around contract calls

Those parts live in the chain repository and any external tooling repository used by your deployment.

## 2. Environment Setup and Build

### 2.1 Prerequisites

Required tools:

- Node.js and npm
- Foundry
- access to the JuChain node software from the chain repository

Recommended checks:

```bash
node --version
npm --version
forge --version
```

### 2.2 Install Dependencies

```bash
npm install
```

This installs the JavaScript dependencies used by the local generation scripts.

### 2.3 Build and Test Contracts

```bash
forge build
forge test
```

Expected outputs:

- compiled artifacts under `out/`
- green unit and contract-level tests

### 2.4 Refresh Generated Artifacts and Genesis

If you modify Solidity source, the expected workflow is:

```bash
forge build
npm run generate
npm run init-genesis
```

What each step does:

- `forge build` compiles contracts into `out/`
- `npm run generate` refreshes generated contract artifacts used by the repo scripts
- `npm run init-genesis` injects compiled bytecode into `genesis.json`

If a local network is already running, wipe node data and re-initialize all nodes after regenerating genesis.

### 2.5 Build Verification Checklist

After building, confirm:

- `out/Proposal.sol/Proposal.json` exists
- `out/Staking.sol/Staking.json` exists
- `out/Validators.sol/Validators.json` exists
- `out/Punish.sol/Punish.json` exists
- `genesis.json` contains alloc entries for `f010` to `f013`

## 3. Genesis Generation and Validation

### 3.1 Genesis Contract Allocations

The generated genesis must predeploy the four PoSA contracts at:

```text
0x000000000000000000000000000000000000f010
0x000000000000000000000000000000000000f011
0x000000000000000000000000000000000000f012
0x000000000000000000000000000000000000f013
```

Do not hand-copy old `0xf000` to `0xf003` examples from legacy POA documents.

### 3.2 `extraData` Format

Genesis validators come from the `extraData` validator list. The field layout is:

```text
extraData = vanity(32 bytes) + validators(20 bytes * N) + signature(65 bytes)
```

Where:

- `vanity` is arbitrary 32-byte padding
- `validators` is the genesis validator address list
- `signature` is the genesis block seal placeholder used by the consensus engine

Congress reads those validators as the initial validator set and uses them to bootstrap the PoSA contracts.

### 3.3 Initialization Order

At block `1`, Congress initializes contracts in this order:

1. `Proposal.initialize(validators, validatorsAddr, epoch)`
2. `Staking.initializeWithValidators(validatorsAddr, proposalAddr, punishAddr, validators, commissionRate)`
3. `Punish.initialize(validatorsAddr, proposalAddr, stakingAddr)`
4. `Validators.initialize(validators, proposalAddr, punishAddr, stakingAddr)`

This ordering matters because:

- `Proposal` must exist before `Staking` can read parameters
- `Staking` must exist before `Punish` and `Validators` can reference real stake and jail state
- `Validators` is initialized last because it references all of the others

### 3.4 Genesis Validator Bootstrap

Bootstrap specifics:

- Congress pre-funds the staking contract with `1 ether` per genesis validator for `initializeWithValidators(...)`
- genesis validators are registered directly by the bootstrap path
- genesis validators start with `1 ether` bootstrap self-stake and `1000` bps commission
- post-launch validator registration still requires `minValidatorStake = 100000 JU` by default

Do not misread the bootstrap amount as the post-launch validator minimum.

### 3.5 Recommended Genesis Validation Checks

Before you initialize nodes from a newly generated genesis, verify:

- chain ID and `config.congress` values are correct
- alloc contains bytecode at `f010` to `f013`
- `extraData` contains the intended validator list
- the initial validator count is non-zero and does not exceed the consensus maximum
- no old POA contract addresses are being used as current system-contract allocs

## 4. Node Deployment and Network Layout

### 4.1 Deployment Shapes

Common layouts:

1. Single-node development environment
   - good for local contract validation and simple end-to-end checks
2. Multi-node validator network
   - required for realistic epoch rotation, governance, missed-block punishment, and failover testing

### 4.2 Minimum Practical Validator Layout

For multi-validator testing, use enough physical nodes to keep consensus live when one validator is jailed or removed.

Operational rule:

- do not configure more active validators than the physical network can realistically support
- remember that Congress uses `N/2 + 1` style recent-signer and quorum calculations

### 4.3 Node Configuration Requirements

Your node deployment must ensure:

- all validators start from the same `genesis.json`
- chain config in the node matches the generated genesis
- validators can sign blocks with the intended mining account
- network peers can discover or statically connect to one another
- RPC is available for contract operations and monitoring

### 4.4 Recommended Validator-Node Settings

Each validator node should have:

- a dedicated data directory
- a dedicated validator key or external signer
- RPC enabled on a protected interface
- mining enabled for the validator address
- persistent peer connectivity

The exact startup command depends on the chain repository and your runtime environment, but those properties must hold regardless of wrapper scripts.

### 4.5 Account and Signer Management

Recommended approach:

- use one validator account per node
- keep fee-receiver and validator key responsibilities explicit
- prefer external signing or hardened key management in production
- keep validator keystores and passwords separate per node

### 4.6 Network Connectivity

For multi-node deployments, decide on one of:

- static peer lists
- trusted peer sets
- controlled peer discovery with firewall rules

Whatever method is used, the validator nodes must have durable connectivity across epoch boundaries. Unstable peer connectivity is one of the fastest ways to create false missed-block punishment.

## 5. Validator and Governance Operations

### 5.1 Validator Lifecycle Overview

The current validator lifecycle is:

```text
governance approval
  -> registerValidator
  -> enters highestValidatorsSet
  -> next epoch enters currentValidatorSet
  -> active operation
  -> resignValidator or punishment
  -> next epoch leaves currentValidatorSet
  -> exitValidator
  -> unbonding
```

### 5.2 Add a Validator

The current validator-admission flow is:

1. an active validator creates an add-validator proposal with `Proposal.createProposal(candidate, true, details)`
2. active non-jailed validators vote with `Proposal.voteProposal(id, true)`
3. as soon as majority is reached, `Proposal.pass[candidate]` becomes `true`
4. the candidate calls `Staking.registerValidator(commissionRate)` with `msg.value >= minValidatorStake`
5. the validator enters `highestValidatorsSet`
6. the validator becomes active in consensus only at the next epoch transition

Important details:

- proposal majority is `getVotingValidatorCount() / 2 + 1`
- proposal approval is immediate on majority
- the candidate must register within `proposalLastingPeriod`
- `registerValidator(...)` is only allowed on non-epoch blocks
- `registerValidator(...)` and `unjailValidator(...)` share the "one validator addition per epoch" limit

### 5.3 Edit Validator Metadata

Validator descriptive data is updated through:

- `Validators.createOrEditValidator(feeAddr, moniker, identity, website, email, details)`

Requirements:

- caller must be the validator address
- validator must have governance authorization through `Proposal.pass`
- metadata field lengths must satisfy the contract validation rules

This operation does not change stake or active-set membership.

### 5.4 Update Commission Rate

Validators change commission through:

- `Staking.updateCommissionRate(newCommissionRate)`

Requirements:

- validator must be registered and valid
- `newCommissionRate > 0`
- `newCommissionRate <= maxCommissionRate`
- cooldown since the previous update must have passed

### 5.5 Remove or Resign a Validator

There are two main removal paths.

Governance removal:

1. create `Proposal.createProposal(candidate, false, details)`
2. validators vote
3. majority approval clears `pass[candidate]`
4. validator is jailed and removed from candidate management paths

Voluntary exit:

1. validator calls `Staking.resignValidator()`
2. validator is marked jailed for technical exit coordination
3. validator is removed from `highestValidatorsSet`
4. validator waits until the next epoch removes it from `currentValidatorSet`
5. validator calls `Staking.exitValidator()`
6. principal enters unbonding

Current valid exit names are:

- `decreaseValidatorStake`
- `resignValidator`
- `exitValidator`

Do not automate against old names such as `withdrawValidatorStake` or `emergencyExit`.

### 5.6 Re-entry After Jail

A jailed validator does not automatically regain status.

To come back, the validator must:

1. wait until `validatorUnjailPeriod` ends
2. pass a fresh governance proposal
3. keep at least `minValidatorStake`
4. call `Staking.unjailValidator(validator)` on a non-epoch block

The validator re-enters the candidate set immediately but only re-enters the active consensus set after the next epoch rotation.

### 5.7 System Parameter Updates

Governable parameters are changed through:

- `Proposal.createUpdateConfigProposal(cid, newValue)`
- `Proposal.voteProposal(id, true)`

Operationally important facts:

- validation happens before proposal creation and again on execution
- the update becomes effective immediately once majority is reached
- it does not wait until `proposalLastingPeriod` expires

## 6. Rewards, Staking, and Withdrawal Operations

### 6.1 Delegation

Users delegate through:

- `Staking.delegate(validator)` with `msg.value`

Requirements:

- validator must be active and not jailed
- minimum delegation is `10 JU` by default
- self-delegation through the public delegation path is rejected

### 6.2 Undelegation

Users undelegate through:

- `Staking.undelegate(validator, amount)`

Requirements:

- amount must be at least `minUndelegation`
- delegation amount must exist
- undelegated principal enters an unbonding queue

### 6.3 Validator Self-Stake Changes

Validators can:

- add self-stake with `addValidatorStake()`
- partially reduce self-stake with `decreaseValidatorStake(amount)`

Important behavior:

- the remaining self-stake after reduction must still satisfy `minValidatorStake`
- reduced principal is not paid immediately; it moves into the validator's unbonding queue

### 6.4 Principal Withdrawal

After the unbonding period:

- call `Staking.withdrawUnbonded(validator, maxEntries)`

This applies to:

- undelegated principal
- validator self-stake reduced via `decreaseValidatorStake(...)`
- validator self-stake exited via `exitValidator()`

### 6.5 Reward Withdrawal

There are three separate reward withdrawal surfaces:

1. validator transaction-fee income
   - `Validators.withdrawProfits(validator)`
2. validator coinbase reward share
   - `Staking.claimValidatorRewards()`
3. delegator reward share
   - `Staking.claimRewards(validator)`

Validator reward and profit withdrawals are each rate-limited by `withdrawProfitPeriod` in their respective paths.

### 6.6 Pending Payouts

The staking contract may queue a payment instead of transferring immediately. When that happens, the recipient must later call:

- `Staking.withdrawPendingPayout(recipient)`

Do not assume reward claims or unbond withdrawals are always paid in the same transaction.

### 6.7 Coinbase Reward Formula

Congress computes the actual producer payout as:

```text
fixedPart    = blockReward * baseRewardRatio / 10000
weightedPool = blockReward * validatorCount * (10000 - baseRewardRatio) / 10000
weightedPart = minerStake * weightedPool / totalStake
actualReward = fixedPart + weightedPart
```

Implications:

- `blockReward` is a base input, not the guaranteed final payout for every block
- the producer's stake weight changes the final amount
- jailed validators are excluded from reward eligibility immediately

## 7. Monitoring and Day-2 Operations

### 7.1 Basic Network Health

At minimum, monitor:

- block height growth
- peer count and peer stability
- validator nodes' signing and mining status
- block production cadence relative to `period`
- epoch-boundary transitions

### 7.2 Contract-Level Health Signals

Useful contract reads include:

- `Validators.getActiveValidators()`
- `Validators.getVotingValidatorCount()`
- `Validators.getRewardEligibleValidatorsWithStakes()`
- `Validators.getTopValidators()`
- `Staking.getValidatorInfo(validator)`
- `Staking.getValidatorStatus(validator)`
- `Proposal.pass(validator)`
- `Proposal.proposals(id)` and `Proposal.results(id)`

These queries are enough to reconstruct most validator-lifecycle and governance state without external CLI wrappers.

### 7.3 Epoch-Sensitive Operating Rules

Treat epoch blocks specially.

Operations blocked on epoch blocks include:

- validator registration
- governance voting
- unjail
- resign
- exit
- execution of pending punishments

If an automation flow unexpectedly reverts with an epoch-related message, first check whether the current block number is divisible by `epoch`.

### 7.4 Reward and Punishment Monitoring

Track these separately:

- transaction-fee income in `Validators`
- validator reward accumulation in `Staking`
- delegator pending reward state via delegation and `rewardPerShare`
- missed-block counters and pending punishment queues in `Punish`
- pending payout balances in `Staking.pendingPayouts`

### 7.5 Validator-Set Drift Checks

At and around epoch boundaries, compare:

- `highestValidatorsSet`
- `getTopValidators()`
- `currentValidatorSet`

Expected behavior:

- `highestValidatorsSet` changes immediately when validators register, resign, or are removed
- `getTopValidators()` reflects current ranking among candidates
- `currentValidatorSet` only changes at epoch rotation

## 8. Troubleshooting Guide

### 8.1 Validator Not Producing Blocks

Check the following in order:

1. is the validator in `currentValidatorSet`
2. is the validator jailed in `Staking`
3. is the validator key unlocked or external signer available
4. does the node have stable peer connectivity
5. is the validator being rejected by Congress because the parent state already marks it jailed

### 8.2 Proposal Creation or Voting Fails

Common causes:

- caller is not an active validator
- proposal cooldown not satisfied
- proposal already expired
- duplicate vote attempt
- transaction sent on an epoch block

### 8.3 `registerValidator` Fails

Common causes:

- candidate has not passed governance
- proposal registration window expired
- insufficient self-stake
- commission rate exceeds `maxCommissionRate`
- another validator was already added in the same epoch
- call was made on an epoch block

### 8.4 `resignValidator` or `exitValidator` Fails

Common causes:

- validator is still within the `doubleSignWindow` since its last active block
- another validator already resigned in the same epoch
- validator is still present in `currentValidatorSet`
- call was made on an epoch block

### 8.5 `unjailValidator` Fails

Common causes:

- jail period not complete
- fresh governance proposal not passed
- self-stake below `minValidatorStake`
- another validator was already added in the current epoch
- call was made on an epoch block

### 8.6 Reward Looks Wrong

Check the reward surface first:

- is the difference in transaction-fee income or coinbase reward
- was the validator jailed and therefore excluded from reward eligibility
- did the validator hit `punishThreshold` and lose fee-income eligibility
- is the balance queued in `pendingPayouts`
- is the validator still waiting for `withdrawProfitPeriod`

### 8.7 Epoch Validator Set Mismatch

If the chain rejects an epoch header for validator mismatch:

1. compare `header.Extra` validators with `Validators.getTopValidators()` at the parent state
2. verify no local node is using a different genesis or chain config
3. verify all nodes agree on the PoSA upgrade time and system contract addresses

### 8.8 Double-Sign Evidence Rejected

Common causes:

- headers are for different heights
- headers are identical
- signer recovery does not match
- evidence is outside `doubleSignWindow`
- target validator no longer exists
- call was made on an epoch block

## 9. Contract Interface Summary

### 9.1 Proposal

High-value operational functions:

- `createProposal(address dst, bool flag, string details)`
- `createUpdateConfigProposal(uint256 cid, uint256 newValue)`
- `voteProposal(bytes32 id, bool auth)`
- `isProposalValidForStaking(address validator)`
- `pass(address validator)`

### 9.2 Staking

High-value operational functions:

- `registerValidator(uint256 commissionRate)`
- `addValidatorStake()`
- `decreaseValidatorStake(uint256 amount)`
- `resignValidator()`
- `exitValidator()`
- `delegate(address validator)`
- `undelegate(address validator, uint256 amount)`
- `withdrawUnbonded(address validator, uint256 maxEntries)`
- `claimRewards(address validator)`
- `claimValidatorRewards()`
- `withdrawPendingPayout(address payable recipient)`
- `unjailValidator(address validator)`
- `getValidatorInfo(address validator)`

### 9.3 Validators

High-value operational functions:

- `createOrEditValidator(...)`
- `withdrawProfits(address validator)`
- `getActiveValidators()`
- `getActiveValidatorCount()`
- `getVotingValidatorCount()`
- `getRewardEligibleValidatorsWithStakes()`
- `getTopValidators()`
- `isValidatorActive(address validator)`
- `isValidatorJailed(address validator)`

### 9.4 Punish

System and evidence functions:

- `punish(address val)`
- `executePending(uint256 limit)`
- `submitDoubleSignEvidence(bytes header1, bytes header2)`
- `decreaseMissedBlocksCounter(uint256 epoch)`

## 10. Best Practices Summary

- treat `blockReward` as a formula input, not a fixed payout promise
- treat `currentValidatorSet` as an epoch cache, not the only source of effective validator status
- separate fee-income accounting from coinbase reward accounting in monitoring and operations
- avoid sending validator-management transactions on epoch blocks
- regenerate genesis whenever contract bytecode changes
- never restore old `ju-cli` or legacy POA flows into current operational runbooks without verifying them against the current contracts
