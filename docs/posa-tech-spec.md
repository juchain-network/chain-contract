# JPoSA Technical Specification

## Scope

This document describes the current JuChain PoSA implementation as defined by:

- `contracts/Proposal.sol`
- `contracts/Staking.sol`
- `contracts/Validators.sol`
- `contracts/Punish.sol`
- `../chain/consensus/congress/congress.go`
- `../chain/consensus/congress/snapshot.go`

It intentionally replaces earlier POA-era and early-PoSA descriptions where behavior has changed.

## Table of Contents

1. [System Overview](#1-system-overview)
2. [Contract Responsibilities](#2-contract-responsibilities)
3. [Contract Collaboration Flows](#3-contract-collaboration-flows)
4. [Consensus Flow and Contract Coordination](#4-consensus-flow-and-contract-coordination)
5. [Major Scenario Flows](#5-major-scenario-flows)
6. [Key Mechanisms](#6-key-mechanisms)
7. [Consensus-to-Contract Sequences](#7-consensus-to-contract-sequences)
8. [Key Parameters and Constants](#8-key-parameters-and-constants)
9. [Safety Mechanisms and Edge Cases](#9-safety-mechanisms-and-edge-cases)
10. [FAQ](#10-faq)

## 1. System Overview

### 1.1 JPoSA Principles

JPoSA combines PoS-style stake competition with PoA-style governance admission.

- stake determines ranking and reward weight
- governance determines who is allowed to become or become-again a validator
- Congress consensus only rotates the active validator set on epoch boundaries

This design splits validator state into three different layers:

- governance authorization
- candidate ranking
- epoch-effective participation

### 1.2 Core Components

| Component | Address | Responsibility | Notes |
| --- | --- | --- | --- |
| `Validators` | `0x000000000000000000000000000000000000F010` | Maintains validator-set caches and fee-income accounting | Owns `currentValidatorSet` and `highestValidatorsSet` |
| `Punish` | `0x000000000000000000000000000000000000F011` | Tracks missed blocks and processes double-sign evidence | Calls into `Staking` and `Validators` |
| `Proposal` | `0x000000000000000000000000000000000000F012` | Governance proposals and governable parameters | Owns `pass` authorization state |
| `Staking` | `0x000000000000000000000000000000000000F013` | Self-stake, delegation, rewards, unbonding, jail state | Source of truth for stake and jail status |
| `Congress` | N/A | Consensus engine in the chain node | Prepares headers, finalizes rewards, updates epoch validator set |

Congress still contains address routing for legacy POA contracts at `0xf000` to `0xf002` for historical compatibility, but current PoSA deployments use `0xf010` to `0xf013`.

### 1.3 Relationship Diagram

```text
┌────────────────────────────────────────────────────────────────┐
│                       Congress Consensus                        │
│  Prepare()   -> set header difficulty, extra-data signers      │
│  Finalize()  -> rewards, punishments, epoch updates            │
│  VerifyHeader() -> validate header.Extra against parent state  │
└────────────────────────────────────────────────────────────────┘
                               │
                               │ system calls
                               ▼
┌────────────────────────────────────────────────────────────────┐
│                        Validators (F010)                        │
│  currentValidatorSet / highestValidatorsSet / fee income       │
└────────────────────────────────────────────────────────────────┘
            │                       │                       │
            │                       │                       │
            ▼                       ▼                       ▼
┌──────────────────┐   ┌──────────────────┐   ┌──────────────────┐
│ Proposal (F012)  │   │ Staking (F013)   │   │ Punish (F011)    │
│ governance       │   │ stake + rewards  │   │ missed blocks    │
│ config params    │   │ jail state       │   │ double-sign       │
└──────────────────┘   └──────────────────┘   └──────────────────┘
```

### 1.4 State Ownership Split

The validator system is intentionally split across four contracts:

- `Proposal.pass[validator]` means governance has authorized the validator
- `Proposal.proposalPassedHeight[validator]` records when that authorization was obtained
- `Validators.highestValidatorsSet` is the validator cold-address candidate cache used for stake-based ranking
- `Validators.currentValidatorSet` is the epoch-effective validator cold-address cache
- `Validators.validatorSigners[validator]` and `Validators.signerValidators[signer]` define the active cold/hot binding
- `Validators.pendingValidatorSigners[validator]` / `pendingSignerEpochs[validator]` hold delayed signer rotations
- `Staking.validatorStakes[validator]` stores self-stake, delegation totals, rewards, and jail state
- `Punish.punishRecords[validator]` stores the missed-block counter and pending punishment state

This split is important: a validator can be present in one cache and absent from another during transitions.

### 1.5 Validator, Signer, and Fee Address Roles

JPoSA supports separated operational identities:

- validator cold address
  - owns self-stake, governance rights, unjail, exit, commission updates, and signer rotation
- signer hot address
  - seals blocks and appears in `header.Coinbase` / epoch `header.Extra`
- fee address
  - withdraws transaction-fee income from `Validators.withdrawProfits(validator)`

Default same-address mode is still supported, but the protocol no longer requires the cold validator key to be the
block-signing key.

### 1.6 Active, Voting, Reward-Eligible, and Top Validators

The code uses several similar-looking concepts that must not be confused.

| Concept | Source | Meaning |
| --- | --- | --- |
| authorized validator | `Proposal.pass` | governance has approved the validator |
| registered validator | `Staking.validatorStakes[validator].isRegistered` | the validator has a stake record |
| top validator | `Validators.getTopValidators()` | candidate selected by stake ranking from `highestValidatorsSet`; requires `isRegistered && selfStake >= minValidatorStake` |
| top signer | `Validators.getTopSigners()` | effective signer set derived from the current top-validator set |
| active validator | `Validators.isValidatorActive()` | in `currentValidatorSet`, registered, not jailed, and `selfStake >= minValidatorStake` |
| active signer | `Validators.getActiveSigners()` | effective signer set derived from `currentValidatorSet` |
| voting validator | `Validators.getVotingValidatorCount()` | governance-counted validator; jailed or below-min-stake validators are excluded immediately |
| reward-eligible validator | `Validators.getRewardEligibleValidatorsWithStakes()` | coinbase-reward-eligible validator; jailed or below-min-stake validators are excluded immediately |
| reward-eligible signer | `Validators.getRewardEligibleSignersWithStakes()` | signer set paired to reward-eligible validators for Congress reward calculation |

## 2. Contract Responsibilities

### 2.1 Proposal

**Role:** governance for validator admission/removal and protocol parameters.

#### Main Features

1. Validator proposals
   - `createProposal(dst, flag, details)`
   - `flag = true` means add validator authorization
   - `flag = false` means remove validator authorization
2. Configuration proposals
   - `createUpdateConfigProposal(cid, newValue)`
   - the proposal validates the new value before creation
3. Voting and finalization
   - `voteProposal(id, auth)`
   - proposal result finalizes immediately once majority is reached
   - voting is only allowed on non-epoch blocks
4. Registration window support
   - `proposalLastingPeriod` is both proposal expiry and registration window
   - `isProposalValidForStaking(validator)` checks whether an approved candidate can still register

#### Main Data

| Field | Meaning |
| --- | --- |
| `proposalLastingPeriod` | proposal expiry and post-pass registration window |
| `pass[address]` | governance authorization bit |
| `proposalPassedHeight[address]` | block height when authorization was granted |
| `lastProposalBlock[address]` | proposer cooldown tracking |
| `proposerNonces[address]` | unique proposal ID support |
| `proposals[id]` | proposal payload |
| `results[id]` | agree/reject counters and finalization bit |
| `votes[voter][id]` | per-voter vote record |

#### Timing and Threshold Rules

- proposal creation is rate-limited by `proposalCooldown`
- voting remains open while `block.number < createBlock + proposalLastingPeriod`
- pass threshold is `agree >= getVotingValidatorCount() / 2 + 1`
- reject threshold is `reject >= getVotingValidatorCount() / 2 + 1`
- results do not wait until expiry; they take effect at the block where threshold is reached

#### Configuration Update Rules

Configuration proposals support `cid = 0` to `19` and become effective immediately once passed. Important constraints include:

- `punishThreshold < removeThreshold`
- `removeThreshold >= decreaseRate`
- `maxValidators <= CONSENSUS_MAX_VALIDATORS`
- `doubleSignSlashAmount >= doubleSignRewardAmount`
- `baseRewardRatio <= 10000`
- `maxCommissionRate <= 10000`
- `burnAddress` must be non-zero

#### Relationships

- depends on `Validators.getVotingValidatorCount()` and `Validators.isValidatorActive()`
- used by `Staking.registerValidator()` and `Staking.unjailValidator()`
- updated by `Validators.removeFromHighestSet()` and `Validators.tryRemoveValidator()`

### 2.2 Staking

**Role:** canonical owner of stake balances, delegation, reward accounting, and jail state.

#### Main Features

1. Genesis bootstrap
   - `initializeWithValidators(...)` pre-registers genesis validators
   - each genesis validator receives `Proposal.minValidatorStake()` bootstrap self-stake
   - Congress satisfies that bootstrap by moving `minValidatorStake` from each bootstrap validator cold address into
     `Staking` before initialization
2. Validator lifecycle
   - `registerValidator(commissionRate)`
   - `addValidatorStake()`
   - `decreaseValidatorStake(amount)`
   - `resignValidator()`
   - `exitValidator()`
   - `unjailValidator(validator)`
3. Delegation lifecycle
   - `delegate(validator)`
   - `undelegate(validator, amount)`
   - `withdrawUnbonded(validator, maxEntries)`
4. Rewards
   - `distributeRewards()` for coinbase rewards
   - `claimRewards(validator)` for delegators
   - `claimValidatorRewards()` for validators
   - `withdrawPendingPayout(recipient)` for delayed transfers
5. Punishment hooks
   - `jailValidator(validator, jailBlocks)`
   - `slashValidator(...)`

#### Main Data

```solidity
struct ValidatorStake {
    uint256 selfStake;
    uint256 totalDelegated;
    uint256 commissionRate;
    uint256 totalRewards;
    uint256 accumulatedRewards;
    bool isJailed;
    uint256 jailUntilBlock;
    uint256 totalClaimedRewards;
    uint256 lastClaimBlock;
    bool isRegistered;
}

struct Delegation {
    uint256 amount;
    uint256 rewardDebt;
}

struct UnbondingEntry {
    uint256 amount;
    uint256 completionBlock;
}
```

Other important storage:

- `lastActiveBlock[validator]`
- `lastCommissionUpdateBlock[validator]`
- `unbondingDelegations[delegator][validator]`
- `rewardPerShare[validator]`
- `validatorsAddedInEpoch`
- `validatorsRemovedInEpoch`
- `pendingPayouts[address]`

#### Key Behavioral Rules

- post-launch `registerValidator(...)` requires:
  - passed proposal
  - valid registration window
  - `msg.value >= minValidatorStake`
  - `commissionRate > 0` and `<= maxCommissionRate`
- `registerValidator(...)` and `unjailValidator(...)` share the "one addition per epoch" guard
- `resignValidator()` consumes the "one removal per epoch" slot
- `decreaseValidatorStake(amount)` is partial-only; if remaining stake would fall below `minValidatorStake`, the call reverts
- `exitValidator()` is only allowed after the validator is no longer in `currentValidatorSet`
- `distributeRewards()` resolves `block.coinbase` signer through `Validators.getValidatorBySigner(...)`
- validator coinbase rewards are claimed by the validator cold address; the signer hot address does not own a separate
  reward bucket

#### Relationships

- depends on `Proposal` for authorization, parameter values, and cooldowns
- depends on `Validators` to add or remove candidates from `highestValidatorsSet`
- called by `Punish` for jailing and slashing
- called by `Congress` for coinbase reward distribution

### 2.3 Validators

**Role:** manages validator-set caches, validator metadata, and transaction-fee income.

#### Main Features

1. Validator metadata and signer binding
   - `createOrEditValidator(feeAddr, moniker, ...)`
   - `createOrEditValidator(feeAddr, signer, moniker, ...)`
   - stores fee address, descriptive metadata, and signer binding
   - callable by a proposal-authorized candidate or any existing registered validator
2. Signer mapping
   - `getValidatorSigner(validator)`
   - `getValidatorBySigner(signer)`
   - `getValidatorBySignerHistory(signer)`
   - pending signer rotation activates from the first block after the next epoch checkpoint
   - historical signer ownership is recorded only after that signer has actually entered the on-chain effective signer set
3. Set management
   - `currentValidatorSet` is the epoch-effective validator cold-address set
   - `highestValidatorsSet` is the validator cold-address candidate cache
   - `getActiveSigners()` / `getTopSigners()` derive the effective signer sets from those validator caches
   - `getTopSignersForEpochTransition()` exposes the signer set that should be committed into the current checkpoint
     header
   - `updateActiveValidatorSet(newSet, epoch)` updates `currentValidatorSet`
   - `tryActive(validator)` inserts a validator into `highestValidatorsSet`
   - `removeFromHighestSet(validator)` removes a validator from `highestValidatorsSet` and clears proposal authorization
4. Transaction-fee income
   - `distributeBlockReward()` accrues transaction-fee reward
   - `withdrawProfits(validator)` lets the validator fee address withdraw `aacIncoming`
5. Query surface
   - `getActiveValidators()`
   - `getActiveSigners()`
   - `getActiveValidatorCount()`
   - `getVotingValidatorCount()`
   - `getRewardEligibleSignersWithStakes()`
   - `getRewardEligibleValidatorsWithStakes()`
   - `getTopSignersForEpochTransition()`
   - `getTopSigners()`
   - `getTopValidators()`

#### Main Data

```solidity
struct Description {
    string moniker;
    string identity;
    string website;
    string email;
    string details;
}

struct Validator {
    address payable feeAddr;
    Description description;
    uint256 aacIncoming;
    uint256 totalJailedHb;
    uint256 lastWithdrawProfitsBlock;
}
```

Other important storage:

- `currentValidatorSet`
- `highestValidatorsSet`
- `validatorSigners[address]`
- `signerValidators[address]`
- `historicalSignerOwners[address]`
- `pendingValidatorSigners[address]`
- `pendingSignerValidators[address]`
- `pendingSignerEpochs[address]`
- `validatorInfo[address]`
- `operationsDone[block][operation]`

#### Key Behavioral Rules

- `currentValidatorSet` and `highestValidatorsSet` store validator cold addresses only
- `getActiveSigners()` and `getTopSigners()` derive signer hot-address sets from the effective cold/hot mapping
- the feeAddr-only `createOrEditValidator(...)` overload preserves the existing signer binding
- if no signer has ever been assigned, the validator address is used as the default signer
- existing registered validators may keep updating `feeAddr`, metadata, and signer binding even after `pass` is cleared
- `getVotingValidatorCount()` excludes jailed validators immediately
- `getRewardEligibleSignersWithStakes()` excludes jailed or below-min-stake validators immediately and returns the
  signer set Congress rewards against
- `getRewardEligibleValidatorsWithStakes()` excludes jailed validators immediately
- `getActiveValidators()` still returns the raw `currentValidatorSet`, which may temporarily include jailed validators until next epoch
- `getTopValidators()` delegates to `Staking.getTopValidators(highestValidatorsSet)`
- when a registered validator rotates signer, the old signer remains valid through the checkpoint block itself and the
  new signer becomes effective from the first block after that checkpoint
- validator removal and voluntary exit clear any pending signer reservation for that validator
- `removeFromHighestSet()` preserves at least one remaining validator in `highestValidatorsSet`

#### Relationships

- depends on `Staking` for real stake and jail status
- depends on `Proposal` to clear `pass` during removal
- called by `Congress` to distribute transaction-fee rewards and rotate epoch validator set
- called by `Punish` and `Proposal` for validator-removal flows

### 2.4 Punish

**Role:** manages missed-block penalties, deferred punishment queues, and double-sign evidence.

#### Main Features

1. Missed-block punishment
   - `punish(val)` increments the counter and triggers threshold logic
   - `val` may be the current validator cold address or the current signer hot address
2. Deferred execution
   - punishment that hits on epoch blocks is deferred into pending queues
   - `executePending(limit)` drains those queues on non-epoch blocks
3. Double-sign evidence
   - `submitDoubleSignEvidence(header1, header2)`
   - resolves the recovered signer through `getValidatorBySignerHistory(...)`
4. Counter decay
   - `decreaseMissedBlocksCounter(epoch)` runs once per epoch

#### Main Data

```solidity
struct PunishRecord {
    uint256 missedBlocksCounter;
    uint256 index;
    bool exist;
}
```

Other important storage:

- `punishValidators`
- `pendingRemove[address]`
- `pendingRemoveIncoming[address]`
- `pendingValidators`
- `doubleSigned[height][validator]`

Upgrade note:

- PoA -> PoSA migration only carries over legacy missed-block state (`punishValidators` / `punishRecords`).
- `pendingRemove`, `pendingRemoveIncoming`, and `pendingValidators` are PoSA runtime queue state and start empty after upgrade.

#### Threshold Semantics

- at `punishThreshold`, `Validators.removeValidatorIncoming(val)` removes fee-income eligibility
- at `removeThreshold`, the validator is jailed and removed from validator-management paths
- if the threshold is hit on an epoch block, the action is queued instead of executed immediately

#### Relationships

- depends on `Proposal` for all threshold and slash parameters
- calls into `Staking` for jail and slash
- calls into `Validators` for fee-income removal and validator removal
- called by `Congress` as a system transaction

## 3. Contract Collaboration Flows

### 3.1 Validator Registration Flow

1. An active validator creates `Proposal.createProposal(candidate, true, details)`.
2. Voting validators call `Proposal.voteProposal(id, true)`.
3. When majority is reached:
   - `pass[candidate] = true`
   - `proposalPassedHeight[candidate] = block.number`
4. Before or after registration, the candidate may call `Validators.createOrEditValidator(...)` to configure `feeAddr`,
   metadata, and an optional signer hot address.
5. The candidate calls `Staking.registerValidator(commissionRate)` with sufficient self-stake.
6. `Staking`:
   - validates proposal state and registration window
   - records `validatorStakes[candidate]`
   - appends the validator to `allValidators`
   - calls `Validators.tryActive(candidate)`
7. `Validators.tryActive(candidate)`:
   - inserts into `highestValidatorsSet`
   - cleans punish record if the validator was previously jailed
8. Congress will include the validator in `currentValidatorSet` only at the next epoch transition.

### 3.2 Governance Removal Flow

1. An active validator creates `Proposal.createProposal(target, false, details)`.
2. Voting validators vote the proposal through `Proposal.voteProposal(...)`.
3. On majority:
   - `pass[target] = false`
   - `proposalPassedHeight[target] = 0`
   - `Validators.tryRemoveValidator(target)` is called
4. `Validators.tryRemoveValidator(target)`:
   - jails the validator first if more than one voting validator exists
   - calls `removeValidatorInternal(target)`
5. `removeValidatorInternal(target)`:
   - removes transaction-fee eligibility if needed
   - removes the validator from `highestValidatorsSet` if at least one candidate remains
   - calls `Proposal.setUnpassed(target)`
6. The validator remains in `currentValidatorSet` until the next epoch update, but loses active privileges immediately once jailed.

### 3.3 Voluntary Exit Flow

1. Validator calls `Staking.resignValidator()`.
2. `Staking`:
   - enforces the one-removal-per-epoch rule
   - checks the `doubleSignWindow` guard based on `lastActiveBlock`
   - marks the validator as jailed for coordinated exit
   - calls `Validators.removeFromHighestSet(msg.sender)`
3. `Validators.removeFromHighestSet(...)`:
   - removes from `highestValidatorsSet`
   - clears `pass[validator]`
4. The validator waits until the next epoch removes it from `currentValidatorSet`.
5. Validator calls `Staking.exitValidator()`.
6. `Staking`:
   - verifies the validator is no longer in `currentValidatorSet`
   - moves all remaining self-stake into the unbonding queue

### 3.4 Transaction-Fee Reward Flow

1. Congress computes block transaction fees and calls `Validators.distributeBlockReward()` with `msg.value`.
2. `Validators.distributeBlockReward()`:
   - resolves `msg.sender` signer to the validator cold address
   - checks whether the validator exists and still satisfies the minimum self-stake floor
   - if the validator is jailed or below `minValidatorStake`, redistributes the fee reward to other active non-jailed
     validators
   - otherwise adds the amount to `validatorInfo[val].aacIncoming`
3. The validator's fee address later withdraws this accumulated income through `Validators.withdrawProfits(validator)`.
   If the current fee address cannot receive ETH, the validator may rotate `feeAddr` via
   `Validators.createOrEditValidator(...)` and retry the withdrawal.

### 3.5 Coinbase Reward Flow

1. Congress reads `Proposal.blockReward()` and `Proposal.baseRewardRatio()`.
2. Congress queries `Validators.getRewardEligibleSignersWithStakes()`.
3. Congress computes `actualReward`.
4. Congress credits the producer and calls `Staking.distributeRewards{value: actualReward}()`.
5. `Staking.distributeRewards()`:
   - resolves `block.coinbase` signer to the validator cold address
   - records validator activity via `lastActiveBlock`
   - takes validator commission
   - allocates validator self-stake share
   - updates `rewardPerShare[validator]` for delegators
6. Later withdrawals happen through:
   - `claimRewards(validator)` for delegators
   - `claimValidatorRewards()` for validators

### 3.6 Epoch Update Flow

1. At epoch block `N`, Congress derives the next validator cold-address set and the corresponding effective signer set
   from the parent block state.
2. `Prepare()` writes the signer set into `header.Extra`.
3. `VerifyHeader()` checks that the epoch header encodes exactly that parent-derived signer set.
4. `Finalize()` calls `handleEpochTransition(...)`.
5. `handleEpochTransition(...)`:
   - calls `updateValidators(newValidatorSet, ...)`
   - calls `decreaseMissedBlocksCounter(...)`
6. `Validators.updateActiveValidatorSet(newValidatorSet, epoch)` replaces `currentValidatorSet`.

### 3.7 Bootstrap and Migration Flow

Fresh PoSA bootstrap:

1. Congress resolves the effective bootstrap mapping from `config.congress.initialValidators` /
   `config.congress.initialSigners`.
2. If one side is omitted, it defaults to the other side. If both are omitted, the genesis signer list from
   `extraData` is used for both validators and signers.
3. `extraData` must match the effective signer set as a set.
4. At initialization, Congress:
   - initializes `Proposal`
   - reads `minValidatorStake`
   - moves one `minValidatorStake` from each bootstrap validator cold address into `Staking`
   - initializes `Staking`, `Punish`, and `Validators.initialize(validators, signers, ...)`

PoA -> PoSA migration:

1. Congress resolves bootstrap validator/signer input in this precedence:
   - CLI overrides `--override.posaValidators` / `--override.posaSigners`
   - chain config `initialValidators` / `initialSigners`
   - default legacy same-address miner mapping
2. Any explicit migration validator/signer remap must supply both arrays together.
3. The effective signer set must cover the live POA validator/signer set being migrated.
4. The same `minValidatorStake` cold-address funding rule applies during migration bootstrap.
5. If any bootstrap validator balance is insufficient at the scheduled upgrade time, Congress defers PoSA activation by
   one epoch and persists the effective `posaTime` override in the node database.
6. Migration rewrites legacy validator and punish state onto validator cold addresses; runtime pending queues start
   empty after upgrade.

## 4. Consensus Flow and Contract Coordination

### 4.1 VerifyHeader

`VerifyHeader()` and `verifyCascadingFields()` enforce the block-level consensus envelope.

Important checks:

- header timestamp is not too far in the future
- `header.Extra` has correct vanity and signature layout
- non-epoch blocks must not contain signer bytes in `header.Extra`
- epoch blocks must contain a non-empty signer list with valid address byte length
- on epoch blocks, `verifyEpochValidators()` compares `header.Extra` with the parent-state contract-derived signer set
- mix digest must be zero
- uncle hash must be empty
- post-fork fields such as `BaseFee`, `WithdrawalsHash`, `ExcessBlobGas`, and `ParentBeaconRoot` are checked when relevant

### 4.2 Prepare

`Prepare()` builds the consensus-specific header fields before transaction execution.

It:

- sets `header.Coinbase` to the local signer hot address
- sets consensus difficulty using `calcDifficulty(...)`
- normalizes vanity bytes in `header.Extra`
- for epoch blocks:
  - resolves the parent-derived validator/signer transition set
  - appends the effective signer set to `header.Extra`
- appends the seal bytes placeholder
- sets header time to at least `parent.Time + period`

Prepare therefore commits the next epoch signer list into `header.Extra` before block execution.

### 4.3 Finalize

`Finalize()` is the consensus-engine execution stage for system behavior.

In current PoSA mode it:

1. rejects blocks from validators that are already jailed in parent state
2. rejects unauthorized external system transactions
3. initializes or migrates system contracts when required
4. punishes out-of-turn validators if needed
5. executes pending punishments
6. distributes transaction-fee rewards
7. distributes coinbase reward
8. handles epoch transition at epoch blocks
9. updates the state root and uncle hash

### 4.4 Snapshot

Congress maintains a consensus snapshot that stores:

- signer hot addresses for signing order and recents-window checks
- recent signers used by spam protection

At epoch blocks, `snapshot.apply(...)`:

- reads signers from `header.Extra`
- rebuilds the signer map from that committed header
- adjusts recents-window state if validator count changed

The snapshot follows the header, not a fresh live contract query.

### 4.5 Validator Selection Source

Validator selection for epoch blocks follows the parent-state rule.

- `resolveEpochTransitionSet(...)` loads the parent header
- `getTopValidatorsAt(chain, parent)` executes `Validators.getTopValidators()` against the parent state root
- `getTopSignersForEpochTransitionFromState(chain, parent, epochHeader)` executes
  `Validators.getTopSignersForEpochTransition()` against the parent state root
- `Validators.getTopValidators()` internally calls `Staking.getTopValidators(highestValidatorsSet)`

This means the epoch block at height `N` commits the signer set derived from state at height `N-1`, while
`Finalize()` still updates `currentValidatorSet` with the corresponding validator cold addresses.

### 4.6 Recent-Signer Window

Congress computes its anti-spam recents window from validator count.

- default rule: `len(validators)/2 + 1`
- when jailed-aware voting count is available in parent state, Congress prefers `getVotingValidatorCount() / 2 + 1`

This avoids liveness loss when many validators are jailed within one epoch but the snapshot still contains the old set.

## 5. Major Scenario Flows

### 5.1 New Validator Onboarding

1. validator community approves a new candidate through governance
2. candidate registers with sufficient self-stake
3. candidate joins `highestValidatorsSet`
4. candidate waits until next epoch to join `currentValidatorSet`
5. after epoch update, candidate becomes:
   - consensus-active
   - governance-eligible
   - reward-eligible unless jailed

### 5.2 Validator Voting

Voting is done through `Proposal.voteProposal(id, auth)`.

Important details:

- only active non-jailed validators count toward the threshold
- a validator may remain in `currentValidatorSet` but be unable to vote if jailed
- once majority is reached, additional votes do not change the result
- voting cannot be performed on epoch blocks

### 5.3 Validator Stake Registration and Increase

Self-stake paths:

- initial registration: `registerValidator(commissionRate)` with `msg.value`
- top-up: `addValidatorStake()` with `msg.value`

Effects:

- registration creates the validator record and candidate-set membership
- top-up changes future ranking immediately because `selfStake + totalDelegated` is the ranking key
- neither operation changes `currentValidatorSet` until epoch rotation

### 5.4 Partial Self-Stake Reduction

`decreaseValidatorStake(amount)`:

- only works for registered valid validators
- cannot reduce remaining self-stake below `minValidatorStake`
- does not transfer funds immediately
- appends an `UnbondingEntry` to the validator's self-owned unbonding queue

This makes validator self-stake reduction use the same delayed principal-withdrawal path as delegator undelegation.

### 5.5 Delegation and Undelegation

Delegation:

- only allowed to active and non-jailed validators
- updates pending reward debt before changing delegation state
- increases validator ranking weight immediately

Undelegation:

- allowed even if the validator has already exited
- reduces delegated amount immediately
- pushes principal into unbonding
- preserves pending reward accounting before the state change

### 5.6 Reward Distribution and Withdrawal

There are three different balances to keep straight:

1. validator fee-income balance in `Validators.aacIncoming`
2. validator coinbase-reward balance in `Staking.accumulatedRewards`
3. delegator reward position tracked by `rewardPerShare`

Withdrawal surface:

- fee-income: `Validators.withdrawProfits(validator)`
- validator coinbase rewards: `Staking.claimValidatorRewards()`
- delegator rewards: `Staking.claimRewards(validator)`
- delayed transfers: `Staking.withdrawPendingPayout(recipient)`

The signer hot address does not own a separate reward ledger. Both fee-income and coinbase reward ultimately resolve
back to the validator cold address and its configured `feeAddr`.

### 5.7 Missed-Block Punishment

When `Punish.punish(val)` runs:

1. missed-block counter increments
2. the input is resolved from current signer hot address to validator cold address when needed
3. if threshold is hit on an epoch block, the action is queued
4. if threshold is hit on a normal block:
   - `punishThreshold` removes fee-income eligibility
   - `removeThreshold` jails and removes the validator

Counter decay happens at epoch blocks through `decreaseMissedBlocksCounter(epoch)`.

### 5.8 Double-Sign Punishment

Double-sign evidence path:

1. reporter submits two conflicting headers
2. `Punish` recovers the signer and checks the evidence window
3. the signer is resolved through `getValidatorBySignerHistory(...)`, so evidence can still target a recently rotated
   signer, but not a signer that was only configured and never became effective on-chain
4. validator is jailed if at least one other voting validator remains
5. `Staking.slashValidator(...)` slashes self-stake
6. reporter reward is paid
7. remaining slash amount is sent to `burnAddress`
8. validator is removed through `Validators.removeValidator(...)`

In the current PoSA design, direct slashing is intentionally limited to validator `selfStake`.
Delegated principal and already-unbonding principal are not slashed by this path.
This means validator economics are deliberately asymmetric:

- validator receives higher upside through commission, self-stake reward share, and fee income
- validator also carries the direct slash risk
- delegators share reward flow and stake-weight support, but do not take direct principal slash

### 5.9 Rejoin After Punishment

To rejoin after jail or governance removal:

1. validator needs fresh governance approval
2. jail period must expire
3. validator must still have at least `minValidatorStake`
4. validator calls `unjailValidator(validator)`
5. validator re-enters `highestValidatorsSet`
6. validator waits until next epoch to re-enter `currentValidatorSet`

### 5.10 Commission and Metadata Updates

Validator operational metadata is updated separately from stake and validator-set membership.

- `Validators.createOrEditValidator(feeAddr, ...)` updates fee address and descriptive fields while preserving signer
- `Validators.createOrEditValidator(feeAddr, signer, ...)` can also bind or rotate signer
- `Staking.updateCommissionRate(newCommissionRate)` updates commission subject to `commissionUpdateCooldown`

`createOrEditValidator(...)` is available both to proposal-authorized pre-registration candidates and to existing
registered validators whose `pass` flag has later been cleared, so fee-income withdrawal and signer recovery can still
be managed while the validator remains registered.

When `createOrEditValidator(...)` is used to rotate a signer, the old signer remains valid through the next epoch
checkpoint block itself, and the new signer becomes the effective consensus signer starting from the first block after
that checkpoint. If the validator leaves before the rotation activates, the pending signer reservation is cleared.

These changes do not require epoch rotation.

## 6. Key Mechanisms

### 6.1 Proposal Mechanics

- proposal IDs use proposer nonces, not timestamps
- add-validator proposals cannot be recreated while a still-valid passed authorization already exists
- remove-validator proposals clear authorization immediately on majority
- configuration proposals validate the new value before proposal creation and again before application

### 6.2 Jail Mechanics

There are two uses of jail state:

1. punitive jail
   - triggered by missed blocks or double-sign evidence
2. technical exit jail
   - triggered by `resignValidator()` to coordinate safe removal at next epoch

In both cases, jail state lives in `Staking`, and other contracts query that state dynamically.

### 6.3 Reward Mechanics

Reward flow is intentionally split:

- `Validators` owns fee-income accounting and signer->validator fee-income resolution
- Congress computes coinbase reward
- `Staking` resolves signer->validator again and splits coinbase reward between validator and delegators

This is why validator rewards are not a single number inside one contract.

### 6.4 Delegation Mechanics

Delegation accounting uses:

- `amount`
- `rewardDebt`
- `rewardPerShare`

This lets the system update delegator reward entitlement without iterating through every delegator on every block reward.
It also reflects the current risk model: delegation increases validator weight and reward-sharing capacity, but direct
slash remains on validator self-stake rather than delegator principal.

### 6.5 Validator Set Updates

Validator-set updates happen in two layers:

- candidate layer: `highestValidatorsSet`
- epoch-effective layer: `currentValidatorSet`

Direct consequences:

- new validators can rank immediately but only become active at next epoch
- jailed validators can lose privileges immediately but still remain visible in `currentValidatorSet` until next epoch

### 6.6 Unified State Ownership

The contracts deliberately do not duplicate full validator state.

- `Proposal` owns authorization and governable parameters
- `Staking` owns stake, delegation, jail, and coinbase reward balances
- `Validators` owns set caches, fee metadata, and transaction-fee income
- `Punish` owns missed-block counters and pending punishment queues

Most "why does this contract not store X?" questions are answered by this split.

### 6.7 Protection Mechanisms

Key protections include:

- `nonReentrant` on the major state-changing and payout functions
- `operationsDone` guards to prevent double reward distribution within one block
- `onlyMiner`, `onlyNotEpoch`, and `onlyBlockEpoch` access gates
- pending payout queue for transfer failures
- last-effective-validator protection during slashing
- at-least-one-validator preservation when removing from `highestValidatorsSet`

## 7. Consensus-to-Contract Sequences

### 7.1 Normal Block Sequence

```text
Prepare
  -> set coinbase / difficulty / time
  -> no signer list in header.Extra

Transactions execute

Finalize
  -> reject jailed producer from parent state
  -> execute pending punishments
  -> distribute fee reward
  -> distribute coinbase reward
  -> no validator-set rotation
```

### 7.2 Epoch Block Sequence

```text
Prepare
  -> query top validators and effective top signers from parent state
  -> write signers into header.Extra

VerifyHeader
  -> compare header.Extra against parent-derived signer set

Finalize
  -> execute rewards as normal
  -> update currentValidatorSet with validator cold addresses
  -> decrease punish counters
  -> validate header extra set again against the derived signer set
```

### 7.3 Validator Turn and Seal Checks

`verifySeal()` uses the consensus snapshot and recent-signer rules.

This means:

- the signer must be in the snapshot signer set
- the signer must not violate the recent-signer window
- jailed-aware recent-limit calculation can shrink the effective recents window when many validators are jailed

### 7.4 Validator Set Update Timeline

```text
Block N-1 state
  -> determine top validators and effective top signers

Epoch Block N Prepare
  -> write parent-derived signers to header.Extra

Epoch Block N execution
  -> still executes the checkpoint block with the previous currentValidatorSet semantics
  -> signer rotations scheduled for block N are not active yet

Epoch Block N Finalize
  -> call Validators.updateActiveValidatorSet(newValidatorSet, epoch)

Block N+1 onward
  -> consensus uses the new signer snapshot from header.Extra
  -> scheduled signer rotations for epoch N become effective
```

## 8. Key Parameters and Constants

### 8.1 Governance and Timing

| Parameter | Default | Meaning |
| --- | --- | --- |
| `proposalLastingPeriod` | `604800` blocks | proposal expiry and registration window |
| `proposalCooldown` | `100` blocks | minimum spacing between proposals from one proposer |
| `commissionUpdateCooldown` | `604800` blocks | minimum spacing between commission updates |
| `withdrawProfitPeriod` | `86400` blocks | interval between validator reward claims |
| `unbondingPeriod` | `604800` blocks | principal unlock delay |
| `validatorUnjailPeriod` | `86400` blocks | jail duration before self-unjail |
| `doubleSignWindow` | `86400` blocks | double-sign evidence freshness window |

### 8.2 Staking and Delegation

| Parameter | Default | Meaning |
| --- | --- | --- |
| `minValidatorStake` | `100000 ether` | minimum validator self-stake; genesis bootstrap uses the same floor at initialization time |
| `maxValidators` | `21` | cap for effective top validators |
| `minDelegation` | `10 ether` | minimum delegation amount |
| `minUndelegation` | `1 ether` | minimum undelegation amount |
| `MAX_UNBONDING_ENTRIES` | `20` | max unbonding entries per delegator-validator pair |

### 8.3 Punishment

| Parameter | Default | Meaning |
| --- | --- | --- |
| `punishThreshold` | `24` | fee-income removal threshold |
| `removeThreshold` | `48` | jail and removal threshold |
| `decreaseRate` | `24` | divisor for epoch decay amount |
| `doubleSignSlashAmount` | `50000 ether` | target slash amount |
| `doubleSignRewardAmount` | `10000 ether` | reporter reward |
| `burnAddress` | `0x000000000000000000000000000000000000dEaD` | remainder recipient |

### 8.4 Rewards

| Parameter | Default | Meaning |
| --- | --- | --- |
| `blockReward` | `0.2 ether` | base reward input for Congress formula |
| `baseRewardRatio` | `3000` | fixed reward share in basis points |
| `maxCommissionRate` | `6000` | validator commission cap in basis points |

### 8.5 Fixed Contract Addresses

| Name | Address |
| --- | --- |
| `VALIDATOR_ADDR` | `0x000000000000000000000000000000000000F010` |
| `PUNISH_ADDR` | `0x000000000000000000000000000000000000F011` |
| `PROPOSAL_ADDR` | `0x000000000000000000000000000000000000F012` |
| `STAKING_ADDR` | `0x000000000000000000000000000000000000F013` |

## 9. Safety Mechanisms and Edge Cases

### 9.1 Reentrancy and Single-Block Guards

- `nonReentrant` protects reward claims, validator lifecycle changes, and payout functions
- `operationsDone[block][operation]` protects against duplicate system actions such as reward distribution and epoch validator updates in one block

### 9.2 Epoch-Sensitive Restrictions

Many operational functions are blocked on epoch blocks through `onlyNotEpoch`, including:

- voting
- validator registration
- resignation
- exit
- unjail
- punishment queue execution

This avoids state races with the epoch transition logic.

### 9.3 Cache Lag vs Immediate Exclusion

`currentValidatorSet` is only updated at epoch boundaries, so it can temporarily lag behind the real effective privileges of a validator.

Immediately after jailing:

- voting eligibility is removed
- coinbase reward eligibility is removed
- block production is rejected by Congress when parent state already shows the jail

Immediately after dropping below `minValidatorStake`:

- top-validator ranking eligibility is removed
- contract-side active / voting / reward-eligible checks return false
- reward sharing skips that validator

What lags is the cache view, not the privilege checks that matter for safety.

### 9.4 Pending Payouts

Reward or principal claims can be queued instead of paid immediately.

This is expected behavior when:

- contract balance is temporarily insufficient
- direct transfer fails

Delayed payout does not mean reward loss; it means the user must later call `withdrawPendingPayout(recipient)`.

### 9.5 Bootstrap Funding and Upgrade Deferral

- fresh PoSA bootstrap requires each bootstrap validator cold address to hold at least `minValidatorStake`
- PoA->PoSA migration applies the same rule to the resolved bootstrap validator cold addresses
- if migration funding is insufficient at the scheduled PoSA time, Congress defers activation by one epoch and persists
  the effective `posaTime` override so restarts keep the same schedule

### 9.6 Last-Validator Protections

Several removal paths preserve liveness by refusing to collapse the effective validator set to zero.

- `Validators.removeFromHighestSet()` requires at least one remaining candidate validator
- `Staking.slashValidator()` preserves `minValidatorStake` for the last effective validator
- punishment removal paths check `getVotingValidatorCount() > 1` before fully jailing and removing

### 9.7 Proposal and Config Safety

- configuration proposals validate before creation and before application
- add-validator proposals cannot be spammed against a still-valid passed authorization
- proposer cooldown avoids rapid proposal churn from one validator

### 9.8 Reward Safety

`blockReward` is a formula input, not a guaranteed exact per-block payout. Documentation, monitoring, and downstream tools must treat:

- transaction-fee income
- actual coinbase reward
- validator commission withdrawal
- delegator reward withdrawal

as separate balances with separate claim paths.

## 10. FAQ

### 10.1 When does a proposal pass?

Immediately when `agree >= getVotingValidatorCount() / 2 + 1`.

### 10.2 When does a newly approved validator start producing blocks?

Only after:

1. the proposal has passed
2. the validator has registered in `Staking`
3. the next epoch has rotated `currentValidatorSet`

### 10.3 Why can a jailed validator still appear in `getActiveValidators()`?

Because `getActiveValidators()` returns the raw epoch cache. Jailed validators are removed from voting, rewards, and block production before that cache is refreshed.

### 10.4 What is the difference between `highestValidatorsSet` and `currentValidatorSet`?

- `highestValidatorsSet` is the candidate cache used for ranking
- `currentValidatorSet` is the epoch-effective validator set used by consensus

### 10.5 How does a validator withdraw rewards?

There are two different reward withdrawal paths:

- transaction-fee income: `Validators.withdrawProfits(validator)`
- coinbase reward share: `Staking.claimValidatorRewards()`

### 10.6 How does a delegator withdraw?

Delegators typically interact with three operations:

- `claimRewards(validator)` for reward claims
- `undelegate(validator, amount)` to start unbonding
- `withdrawUnbonded(validator, maxEntries)` after the unbonding period

### 10.7 Can a validator exit immediately?

No. The validator must:

1. resign
2. wait until no longer present in `currentValidatorSet`
3. call `exitValidator()`
4. wait through unbonding for principal withdrawal

### 10.8 Why does Congress use parent state for epoch validators?

Because the epoch block header must commit to a deterministic validator set before that block's execution mutates current state. Parent-state selection ensures the header and the validator-rotation logic agree.

### 10.9 What is the difference between validator, signer, and feeAddr?

- validator = cold address that owns stake and governance
- signer = hot address that seals blocks for the validator
- feeAddr = address that withdraws transaction-fee income

They may be the same address, but the protocol does not require that.

### 10.10 When does signer rotation take effect?

If a registered validator schedules a new signer for epoch block `N`, the old signer is still valid on block `N`
itself, and the new signer becomes effective starting from block `N+1`.
