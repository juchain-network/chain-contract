# JPoSA Consensus Whitepaper

## Abstract

JPoSA is JuChain's Proof of Stake Authority design. It combines stake-based validator competition with explicit on-chain governance authorization.

The system is not a plain PoS validator auction and not a plain PoA allowlist:

- stake decides ranking and reward weight
- governance decides who is allowed to enter or re-enter the validator pool
- Congress consensus applies validator-set changes only at epoch boundaries

This document summarizes the current behavior of the PoSA system contracts and the Congress consensus engine.

## Core Highlights

### 1. Hybrid Consensus

JPoSA combines the main benefits of PoS and PoA:

- PoS-style competition through self-stake plus delegation
- PoA-style admission control through governance proposals
- epoch-based validator-set updates for operational stability

### 2. Governed Validator Admission

Becoming a validator after genesis requires two distinct steps:

1. governance approval
2. staking registration

This means stake alone is not enough to join the validator pool.

### 3. Dual Reward Channels

JPoSA separates:

- transaction-fee income
- coinbase reward

The transaction-fee path is handled by `Validators`, while the coinbase reward is computed by Congress and distributed through `Staking`.

### 4. Dynamic but Predictable Validator Set

The validator set is dynamic, but not every state change affects consensus immediately.

- candidate ranking changes when stake changes
- effective block production changes at epoch boundaries
- jailed validators lose key privileges immediately even before the epoch cache rotates

### 5. Governance-Tunable Economics

Core parameters such as block reward input, punishment thresholds, unbonding period, validator caps, commission caps, and slash amounts are governable without redeploying the system contracts.

## Governance Mechanism

JPoSA governance is implemented through four cooperating subsystems.

### Proposal System

The proposal system is the entry point for network policy decisions.

It supports two proposal types:

1. validator proposals
   - add or remove validator authorization
2. configuration proposals
   - update governable protocol parameters

Voting rules:

- only active non-jailed validators vote
- threshold is more than half of the voting validator count
- proposal results finalize immediately once majority is reached
- proposal results do not wait until expiry to take effect

`proposalLastingPeriod` still matters, but as:

- the voting expiry window
- the registration window for an approved new validator

### Validator System

The validator system maintains the validator-set views used by the network.

It tracks:

- the epoch-effective active set
- the candidate cache used for ranking
- validator fee addresses and metadata
- transaction-fee income balances

This makes `Validators` the operational bridge between governance, staking state, and the consensus engine.

### Staking System

The staking system manages all stake-related economic state:

- validator self-stake
- delegation
- undelegation and unbonding
- validator and delegator reward accounting
- jail status

It is also where validators:

- register after governance approval
- resign
- exit
- unjail

### Punish System

The punishment system enforces liveness and safety rules.

It:

- tracks missed-block counters
- removes fee-income eligibility at the lower threshold
- jails and removes validators at the higher threshold
- processes double-sign evidence
- decays missed-block counters at epoch boundaries

Together, these four systems provide governance, economic incentives, and operational enforcement without collapsing all logic into one contract.

## Economic Model

JPoSA uses an incentive model built around stake participation, validator performance, and governance-tunable parameters.

### Reward Sources

There are two distinct reward sources:

1. transaction-fee income
   - routed through `Validators`
   - accrues to the producer unless the producer is jailed
2. coinbase reward
   - computed by Congress
   - distributed through `Staking`

This separation matters operationally because the two balances have different accounting and withdrawal paths.

### Coinbase Reward Formula

The governable parameter `blockReward` is a base input, not a guaranteed fixed payout for every block.

Congress computes the actual producer reward as:

```text
fixedPart    = blockReward * baseRewardRatio / 10000
weightedPool = blockReward * validatorCount * (10000 - baseRewardRatio) / 10000
weightedPart = minerStake * weightedPool / totalStake
actualReward = fixedPart + weightedPart
```

This means:

- each reward-eligible producer gets a fixed component
- the rest depends on stake weight among reward-eligible validators

### Reward Distribution Inside Staking

Once Congress computes `actualReward`, the staking contract splits it into:

- validator commission
- validator self-stake share
- delegator share

If a validator has no delegators, the delegator portion is retained by the validator.

### Reward Eligibility

Reward eligibility excludes jailed validators immediately, even if the epoch-effective set has not yet rotated.

This prevents a jailed validator from continuing to benefit from block rewards solely because it still appears in the current epoch cache.

### Liquidity and Payout Semantics

JPoSA uses delayed principal withdrawal and fault-tolerant payout semantics.

- undelegated principal enters unbonding
- exited validator stake also enters unbonding
- reward claims and principal withdrawals may be queued if immediate transfer is not possible

This means a successful claim may create a pending payout instead of always transferring funds in the same transaction.

## Miner (Validator) Role

In JPoSA, miners are validators. They produce blocks, participate in governance, and maintain validator service quality.

### Process to Become a Validator

After genesis, the validator path is:

1. proposal stage
   - an active validator creates an add-validator proposal
   - active non-jailed validators vote
2. registration stage
   - once majority is reached, the candidate is authorized
   - the candidate calls `registerValidator(...)` within the proposal validity window
   - the candidate provides at least the minimum validator self-stake
3. activation stage
   - the validator enters the candidate set immediately
   - the validator enters the active consensus set at the next epoch

Current default thresholds:

- minimum validator self-stake: `100000 JU`
- maximum active validators: `21`

### Validator Core Responsibilities

Validators are responsible for:

- producing blocks in the Congress turn-taking schedule
- staying online and synchronized
- participating in governance votes
- maintaining competitive stake and delegation
- managing their commission rate and fee address

### Validator Responsibilities and Risks

Validators face several operational and economic risks:

1. missed blocks
   - first threshold removes fee-income eligibility
   - second threshold causes jail and validator removal
2. double signing
   - causes jail and stake slashing
3. delayed exit
   - validators cannot exit directly from the active set
4. governance dependency
   - re-entry after punishment requires a fresh proposal

This is an intentional high-upside / high-risk role in the current PoSA model:

- validators earn commission
- validators keep the validator share of self-stake rewards
- validators receive transaction-fee income through the validator fee address
- direct slash risk falls on validator self-stake

### Validator Exit and Return

Voluntary exit is staged:

1. `resignValidator()`
2. wait until the next epoch removes the validator from the active set
3. `exitValidator()`
4. wait through unbonding to recover principal

Return from jail is also staged:

1. serve the jail period
2. regain governance approval
3. satisfy the minimum self-stake requirement
4. call `unjailValidator(...)`
5. wait for the next epoch to become active again

## User (Delegator) Role

Delegators support validators economically and share in the validator reward flow.
In the current PoSA design, delegators share upside but do not take direct slash on delegated principal or already
unbonding principal.

### Delegation Functions

Delegators can:

- choose active validators to support
- delegate JU to one or more validators
- claim their share of staking rewards
- undelegate and later withdraw principal after unbonding

Current default thresholds:

- minimum delegation: `10 JU`
- minimum undelegation: `1 JU`

### Unbonding and Withdrawal

Delegated principal is not instantly liquid once undelegated.

The path is:

1. `undelegate(...)`
2. principal enters unbonding
3. `withdrawUnbonded(...)` after the unbonding period

Reward withdrawals are separate from principal withdrawals and do not require the same unbonding path, although payout may still be queued if direct transfer fails.

### Delegation Strategy Considerations

Delegators should evaluate validators on:

- uptime and block production quality
- governance participation
- commission rate
- history of punishment
- stake depth and long-term stability

Delegation is not only a yield choice; it also influences validator ranking and the composition of the future active set.
Delegators still carry indirect risk through validator underperformance, jailing, removal, missed reward opportunity,
and the time cost of unbonding, even though direct slashing is borne by validators rather than delegators.

## Governance Committee Role

The governance committee is the set of active non-jailed validators in the current epoch.

### Governance Responsibilities

The committee is responsible for:

- approving or rejecting validator entry and return
- removing validators when necessary through governance
- updating economic and operational parameters
- maintaining validator-set quality and network safety

### Proposal Types

The committee votes on:

1. validator admission and removal proposals
2. parameter update proposals

Parameter updates cover:

- proposal timing
- punishment thresholds
- reward inputs
- unbonding and unjail periods
- validator caps and minimums
- commission caps and cooldowns
- slash and reporter-reward amounts

### Voting Mechanism

The current voting mechanism has four properties:

- only active non-jailed validators vote
- threshold is more than half of the voting validator count
- majority finalizes immediately
- voting is disallowed on epoch blocks

This gives JPoSA a governance process that is decisive but still bounded by the active validator set.

## Governable Parameters

The current governable parameter set can be understood in four groups.

### Time-Related Parameters

- `proposalLastingPeriod = 604800` blocks
- `withdrawProfitPeriod = 86400` blocks
- `unbondingPeriod = 604800` blocks
- `validatorUnjailPeriod = 86400` blocks
- `doubleSignWindow = 86400` blocks
- `commissionUpdateCooldown = 604800` blocks
- `proposalCooldown = 100` blocks

### Punishment-Related Parameters

- `punishThreshold = 24`
- `removeThreshold = 48`
- `decreaseRate = 24`
- `doubleSignSlashAmount = 50000 JU`
- `doubleSignRewardAmount = 10000 JU`
- `burnAddress = 0x000000000000000000000000000000000000dEaD`

### Reward-Related Parameters

- `blockReward = 0.2 JU`
- `baseRewardRatio = 3000` bps
- `maxCommissionRate = 6000` bps

### Validator and Delegation Parameters

- `minValidatorStake = 100000 JU`
- `maxValidators = 21`
- `minDelegation = 10 JU`
- `minUndelegation = 1 JU`

These parameters let the network tune validator economics and operational safety without redeploying the system contracts.

## Summary

### Core Values

JPoSA is built around three core values:

1. controlled decentralization
   - stake matters, but entry is governed
2. deterministic operation
   - consensus uses parent-state validator selection and epoch-based activation
3. operational resilience
   - punishment, reward exclusion, and pending payout mechanics handle adverse conditions explicitly

### Why the Design Matters

The result is a validator model that:

- keeps PoA-style admission control
- adds PoS-style delegation and stake-weighted competition
- exposes major policy levers through governance
- separates candidate ranking from epoch-effective participation

That combination is the defining characteristic of JuChain PoSA.
