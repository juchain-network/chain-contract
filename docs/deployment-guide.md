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
      "epoch": 86400,
      "initialValidators": ["0xValidatorCold1", "0xValidatorCold2"],
      "initialSigners": ["0xSignerHot1", "0xSignerHot2"]
    }
  }
}
```

Meaning:

- `period`: target block interval in seconds
- `epoch`: validator-set rotation period in blocks
- `initialValidators`: optional bootstrap validator cold-address list
- `initialSigners`: optional bootstrap signer hot-address list

Validator-set changes only become effective at epoch boundaries.

Fresh PoSA bootstrap mapping rules:

- if both `initialValidators` and `initialSigners` are omitted, fresh PoSA defaults to same-address mode from the
  genesis `extraData` signer list
- if only one side is provided, the missing side defaults to the same addresses
- if both are provided, the two arrays are positional and define the validator-to-signer mapping
- for in-place PoA->PoSA migration, any explicit validator/signer remap must provide both sides together

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

Genesis signer addresses come from the `extraData` list. The field layout is:

```text
extraData = vanity(32 bytes) + signers(20 bytes * N) + signature(65 bytes)
```

Where:

- `vanity` is arbitrary 32-byte padding
- `signers` is the genesis signer hot-address list used by consensus
- `signature` is the genesis block seal placeholder used by the consensus engine

For fresh PoSA:

- `config.congress.initialValidators` / `initialSigners` define the bootstrap cold/hot mapping
- `extraData` must match the effective signer set as a set
- the order of `extraData` does not define cold/hot pairing; the pairing comes from the positional relationship between
  `initialValidators` and `initialSigners`
- if no bootstrap mapping is configured, same-address mode uses the `extraData` signer list as both validators and
  signers

### 3.3 Initialization Order

At block `1`, Congress initializes contracts in this order:

1. `Proposal.initialize(validators, validatorsAddr, epoch)`
2. `Staking.initializeWithValidators(validatorsAddr, proposalAddr, punishAddr, validators, commissionRate)`
3. `Punish.initialize(validatorsAddr, proposalAddr, stakingAddr)`
4. `Validators.initialize(validators, signers, proposalAddr, punishAddr, stakingAddr)`

This ordering matters because:

- `Proposal` must exist before `Staking` can read parameters
- `Staking` must exist before `Punish` and `Validators` can reference real stake and jail state
- `Validators` is initialized last because it references all of the others

Congress always uses the dual-array `Validators.initialize(validators, signers, ...)` path. Same-address mode is
achieved by passing identical validator and signer arrays.

### 3.4 Genesis Validator Bootstrap

Bootstrap specifics:

- Congress initializes `Proposal` first, reads `Proposal.minValidatorStake()`, and moves that amount per bootstrap
  validator from the validator cold-address balances into the `Staking` contract before `initializeWithValidators(...)`
- fresh PoSA genesis alloc must therefore fund every bootstrap validator cold address with at least
  `Proposal.minValidatorStake()`
- genesis validators are registered directly by the bootstrap path
- genesis validators start with `minValidatorStake` bootstrap self-stake and `1000` bps commission
- with current defaults this means `100000 JU` per genesis validator
- validators below the current `minValidatorStake` are excluded from top-validator ranking and from contract-side active / voting / reward eligibility

Bootstrap stake and validator minimum are intentionally aligned at initialization time. The bootstrap path does not mint
synthetic stake into `Staking`; it consumes real cold-address balances.

### 3.5 Bootstrap and Upgrade Validation Checks

Before you initialize nodes from a freshly generated genesis or schedule a PoA->PoSA upgrade, verify:

- chain ID and `config.congress` values are correct
- alloc contains bytecode at `f010` to `f013`
- `extraData` contains the intended signer hot-address set
- the effective bootstrap validator/signer mapping is the one you intend:
  - fresh PoSA and re-genesis upgrades use `config.congress.initialValidators` / `initialSigners`
  - if one side is omitted, the missing side defaults to the other side
  - if both sides are omitted, same-address mode uses the `extraData` signer list for both
- the bootstrap validator/signer count is between `1` and `21`; exceeding `21` will fail during contract
  initialization even if earlier genesis tooling does not reject it
- every bootstrap validator cold address has at least `minValidatorStake` available for bootstrap funding
- no old POA contract addresses are being used as current system-contract allocs

For in-place PoA->PoSA upgrades, Congress resolves bootstrap validator/signer input in this precedence:

1. CLI overrides:
   - `--override.posaValidators`
   - `--override.posaSigners`
2. chain config:
   - `config.congress.initialValidators`
   - `config.congress.initialSigners`
3. default fallback:
   - validator = signer = legacy POA miner address

Additional upgrade rules:

- explicit migration remaps must provide validator and signer arrays together, whether they come from chain config or
  from `--override.posaValidators` / `--override.posaSigners`
- the effective bootstrap signer set must cover the live POA validator/signer set that is being migrated
- PoA->PoSA uses the same `minValidatorStake` bootstrap funding rule from validator cold-address balances
- if a migration bootstrap validator lacks enough balance at the scheduled PoSA time, Congress defers activation by one
  epoch and persists the effective `posaTime` override in the node database so restarts stay consistent

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
- validator nodes can sign blocks with the intended signer hot account
- network peers can discover or statically connect to one another
- RPC is available for contract operations and monitoring

### 4.4 Recommended Validator-Node Settings

Each validator node should have:

- a dedicated data directory
- a dedicated signer hot key or external signer
- RPC enabled on a protected interface
- mining enabled for the signer hot address
- persistent peer connectivity

The exact startup command depends on the chain repository and your runtime environment, but those properties must hold regardless of wrapper scripts.

### 4.5 Account and Signer Management

Recommended approach:

- treat the three roles explicitly:
  - validator cold address: governance, self-stake, signer rotation, unjail, and exit
  - signer hot address: block production only
  - `feeAddr`: transaction-fee withdrawal
- default same-address mode is supported, but production deployments should keep the validator cold key off the
  validator node and only load the signer hot key there
- prefer external signing or hardened key management in production
- keep cold-key storage, signer keystores, and fee-receiver operational controls separate

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
4. the candidate may preconfigure `feeAddr`, metadata, and an optional signer hot address through
   `Validators.createOrEditValidator(...)`
5. the candidate calls `Staking.registerValidator(commissionRate)` with `msg.value >= minValidatorStake`
6. the validator enters `highestValidatorsSet`
7. the validator becomes active in consensus only at the next epoch transition

Important details:

- proposal majority is `getVotingValidatorCount() / 2 + 1`
- proposal approval is immediate on majority
- the candidate must register within `proposalLastingPeriod`
- if the candidate never sets a separate signer, the validator address remains the effective signer
- `registerValidator(...)` is only allowed on non-epoch blocks
- `registerValidator(...)` and `unjailValidator(...)` share the "one validator addition per epoch" limit

### 5.3 Edit Validator Metadata and Signer Binding

Validator operational data is updated through either overload:

- `Validators.createOrEditValidator(feeAddr, moniker, identity, website, email, details)`
- `Validators.createOrEditValidator(feeAddr, signer, moniker, identity, website, email, details)`

Requirements:

- caller must be the validator cold address
- validator must either have governance authorization through `Proposal.pass` or already exist as a registered validator
- `feeAddr` must be non-zero
- metadata field lengths must satisfy the contract validation rules

Behavior:

- the feeAddr-only overload keeps the current signer unchanged
- if a signer is provided for a registered validator, the old signer stays valid through the next epoch checkpoint block
  and the new signer becomes effective from the first block after that checkpoint
- once a rotation has crossed that checkpoint, the new signer node should already be online for block production; if the
  validator is then jailed, removed, or resigns before the next checkpoint, contract cleanup preserves the active
  signer-to-validator mapping long enough for Congress punishment/reward settlement to stay aligned with the current
  epoch snapshot
- if the validator leaves before the pending signer activates, the pending signer reservation is cleared
- existing registered validators may continue to update `feeAddr`, metadata, and signer binding even after `pass` has
  been cleared, as long as the validator still exists in `Staking`
- feeAddr and metadata changes are immediate; signer rotation is epoch-delayed

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

Operational note:

- delegation increases validator ranking weight and reward-sharing capacity
- delegation does not give the delegator direct governance voting rights

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

Already-unbonding principal remains withdrawable after the waiting period and is not directly slashed by the current
validator punishment path.

### 6.5 Reward Withdrawal

There are three separate reward withdrawal surfaces:

1. validator transaction-fee income
   - `Validators.withdrawProfits(validator)`
2. validator coinbase reward share
   - `Staking.claimValidatorRewards()`
3. delegator reward share
   - `Staking.claimRewards(validator)`

Role split:

- transaction-fee income is withdrawn by `feeAddr`
- coinbase reward share is claimed by the validator cold address
- the signer hot address does not own a separate reward bucket

Validator reward and profit withdrawals are each rate-limited by `withdrawProfitPeriod` in their respective paths. If a
validator fee address is misconfigured or cannot receive ETH, the validator can rotate `feeAddr` through either
`createOrEditValidator(...)` overload and then retry `withdrawProfits(...)`, even after `pass` has been cleared, as
long as the validator still exists in `Staking`.

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
- governance updates keep `blockReward` within a safe arithmetic bound
- the producer's stake weight changes the final amount
- jailed validators are excluded from reward eligibility immediately
- if the producer is no longer `isValidatorActive(...)` by block end, the block remains valid but `actualReward = 0`

### 6.8 Current Risk Allocation Model

The current PoSA economics are intentionally asymmetric:

- validators are the high-upside / high-risk role
- validators earn commission, self-stake reward share, and transaction-fee income
- direct slashing applies to validator `selfStake`
- delegated principal and already-unbonding principal are not directly slashed by the current punishment path

This does not make delegators risk-free. Delegators still bear indirect risk through validator underperformance,
jailing or removal, reward interruption, and unbonding time cost.

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

Signer-aware queries that are useful in separated cold/hot deployments:

- `Validators.getValidatorSigner(validator)`
- `Validators.getValidatorBySigner(signer)`
- `Validators.getPendingValidatorSigner(validator)`
- `Validators.getPendingValidatorBySigner(signer)`
- `Validators.getActiveSigners()`
- `Validators.getTopSigners()`
- `Validators.getTopSignersForEpochTransition()`
- `Validators.getRewardEligibleSignersWithStakes()`

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
- pending signer rotations through `getPendingValidatorSigner(...)` / `getPendingValidatorBySigner(...)`
- height-sensitive signer history through `getValidatorBySignerHistoryAt(signer, blockNumber)`

For the pending signer getters:

- `pending = true` means the delayed-rotation record is still stored consistently on-chain
- it does not guarantee the record is still future-dated; a due record may remain observable until a sync or cleanup
  path executes
- use `effectiveBlock` together with the current block height to distinguish “scheduled” from “already due”

### 7.5 Validator-Set Drift Checks

At and around epoch boundaries, compare:

- `highestValidatorsSet`
- `getTopValidators()`
- `currentValidatorSet`
- `getTopSigners()`
- `getTopSignersForEpochTransition()`
- `getActiveSigners()`

Expected behavior:

- `highestValidatorsSet` changes immediately when validators register, resign, or are removed
- `getTopValidators()` reflects current ranking among candidates
- `currentValidatorSet` only changes at epoch rotation
- `getTopSigners()` and `getActiveSigners()` reflect current-block runtime signer semantics
- `getTopSignersForEpochTransition()` is the checkpoint-only query used to build the next epoch header signer set

## 8. Troubleshooting Guide

### 8.1 Validator Not Producing Blocks

Check the following in order:

1. is the validator in `currentValidatorSet`
2. does `Validators.getValidatorSigner(validator)` return the signer hot address you expect
3. if a rotation was recently scheduled, does `Validators.getPendingValidatorSigner(validator)` show the expected
   pending signer and `effectiveBlock`
4. is that signer present in the effective signer set (`getActiveSigners()` / epoch `header.Extra`)
5. is the validator jailed in `Staking`
6. is the signer hot key unlocked or external signer available
7. does the node have stable peer connectivity
8. is the validator being rejected by Congress because the parent state already marks it jailed

Operational note for signer rotation:

- keep the old signer running through the checkpoint block that commits the rotation
- have the new signer node ready before the first post-checkpoint block, otherwise the new signer will start missing
  scheduled turns immediately and can be punished/jail-removed by the normal Congress flow

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
- did the producer fall below `minValidatorStake` or otherwise fail `isValidatorActive(...)` by block end, causing zero
  base reward mint
- did the validator hit `punishThreshold` and lose fee-income eligibility
- is the balance queued in `pendingPayouts`
- is the validator still waiting for `withdrawProfitPeriod`

### 8.7 Epoch Validator Set Mismatch

If the chain rejects an epoch header for validator mismatch:

1. compare `header.Extra` signers with `Validators.getTopSignersForEpochTransition()` at the parent state
2. verify no local node is using a different genesis or chain config
3. verify all nodes agree on the effective bootstrap validator/signer mapping and PoSA upgrade time
4. for in-place upgrades, verify every node is using the same `--override.posaTime`, `--override.posaValidators`, and
   `--override.posaSigners` inputs

### 8.8 Double-Sign Evidence Rejected

Common causes:

- headers are for different heights
- headers are identical
- signer recovery does not match
- signer was not yet effective at the evidence height
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

- `createOrEditValidator(feeAddr, moniker, ...)`
- `createOrEditValidator(feeAddr, signer, moniker, ...)`
- `withdrawProfits(address validator)`
- `getValidatorSigner(address validator)`
- `getValidatorBySigner(address signer)`
- `getPendingValidatorSigner(address validator)`
- `getPendingValidatorBySigner(address signer)`
- `getValidatorBySignerHistory(address signer)`
- `getValidatorBySignerHistoryAt(address signer, uint256 blockNumber)`
- `getActiveSigners()`
- `getActiveValidators()`
- `getActiveValidatorCount()`
- `getVotingValidatorCount()`
- `getRewardEligibleSignersWithStakes()`
- `getRewardEligibleValidatorsWithStakes()`
- `getTopSignersForEpochTransition()`
- `getTopSigners()`
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
- remember that governance cannot raise `minValidatorStake` into a zero-voting-validator state
- treat `currentValidatorSet` as an epoch cache, not the only source of effective validator status
- separate fee-income accounting from coinbase reward accounting in monitoring and operations
- distinguish parent-state block authorization from current-state reward eligibility
- treat validator cold address, signer hot address, and `feeAddr` as separate roles unless you intentionally choose
  same-address mode
- document clearly that delegation boosts validator ranking, while direct slash remains on validator self-stake
- avoid sending validator-management transactions on epoch blocks
- provide `--override.posaValidators` and `--override.posaSigners` together when using in-place PoA->PoSA overrides
- regenerate genesis whenever contract bytecode changes
- never restore old `ju-cli` or legacy POA flows into current operational runbooks without verifying them against the current contracts
