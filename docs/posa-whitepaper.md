# JPoSA Consensus Mechanism Whitepaper

## Abstract

JPoSA (Juchain Proof of Stake Authority) is an innovative hybrid consensus mechanism that combines the advantages of PoS (Proof of Stake) and PoA (Proof of Authority). This mechanism aims to provide security and fairness for decentralized networks while ensuring high performance and fast transaction confirmation.

The core of the JPoSA consensus mechanism lies in selecting block producers through economic incentives and governance mechanisms, while utilizing an established authority verification system to ensure network security. This two-layer mechanism preserves the efficiency of traditional PoA while introducing the decentralization characteristics and economic incentives of PoS.

## Core Highlights

### 1. Hybrid Consensus Mechanism
- Combines the advantages of PoS proof-of-stake and PoA proof-of-authority
- Selects validators based on staking weight while maintaining the stability of authority validators
- Achieves a balance between decentralization and high performance

### 2. Economic Incentive Model
- Validators and delegators share block rewards
- Validators can set commission rates, with a default of 10%, extracting a certain percentage from total rewards as miner commissions
- Remaining rewards are distributed to validators and their delegators based on staking weight

### 3. Dynamic Validator Set
- Automatically updates the validator list based on staking weight
- Supports up to MAX_VALIDATORS active validators (governable parameter, default: 21)
- Minimum validator staking requirement: MIN_VALIDATOR_STAKE (governable parameter, default: 100,000 JU)

### 4. Governance System
- Proposal governance mechanism supports protocol upgrades and parameter adjustments
- Validators participate in network governance decisions- Punishment mechanism prevents malicious behavior

### 5. Security Mechanisms
- Sets unbonding period to prevent sudden exits
- Punishment mechanism targets bad behavior
- Double-sign evidence slashing with reporter reward and burn
- Re-entry attack protection and parameter validation

## Governance Mechanism

The JPoSA consensus mechanism adopts an advanced governance model to ensure the democracy and transparency of network decisions. This governance system consists of four core components:

### Proposal System
The proposal system is the core of the governance mechanism, allowing validators to propose and vote on important network changes:

1. **Validator Proposals**
   - Adding or removing validator qualifications requires governance proposals
   - New validators must receive majority approval from existing validators
   - Proposal passing threshold: more than half of active validators agree

2. **System Configuration Proposals**
   - Parameter adjustments, such as punishment thresholds, unbonding periods, etc.
   - Block reward amount adjustments
   - Other protocol-level changes

### Validator System
The validator management system is responsible for maintaining the status of the current validator set:

- Managing the active validator list
- Distributing transaction fee income
- Maintaining basic validator information

### Staking System
The staking system manages all equity-related functions:

- Validator self-staking management
- User delegation staking processing
- Reward distribution mechanism
- Punishment execution

### Punish System
The punishment system monitors validator behavior and executes corresponding penalties:

- Tracking validators' missed block situations
- Implementing punishment measures when thresholds are reached
- Automatically jailing jailed validators
- Double-sign evidence slashing (jail + slash), reporter rewarded, remainder burned

The governance mechanism achieves decentralized network management through the collaborative work of these four systems, ensuring that all major decisions are fully discussed and voted on by the validator community.

## Economic Model

The JPoSA consensus mechanism adopts a sustainable economic model that incentivizes validators and delegators to actively participate in network maintenance through reasonable reward distribution.

### Block Rewards
The network incentivizes validators to maintain network security and process transactions through block rewards. The quantity of block rewards can be adjusted through governance proposals:

- **Basic Block Reward**: Each block generates 0.2 JU reward (sample value, adjustable through governance)
- **Daily Total Issuance**: 17,280 JU (0.2 JU/block × 86,400 blocks/day)
- **Annual Total Issuance**: Approximately 6,307,200 JU (17,280 JU/day × 365 days)

### Reward Distribution Mechanism
Block rewards are distributed in two parts:

1. **Transaction Fee Distribution** (100%)
   - All transaction fees are distributed to the block-producing validator
   - Jailed validators cannot receive transaction fees

2. **Basic Reward Distribution** (Proportional Distribution)
   - Block-producing validators can set commission rates, extracting a certain percentage from basic rewards as commissions (default 10%)
   - The remaining portion is distributed to validators and their delegators based on staking weight

### Staking Yield
Staking yield depends on the network's total staking volume and transaction activity levels:

- **Validator Self-Staking Income** = Personal share of basic rewards + Transaction fee share + Delegator commissions
- **Delegator Income** = Reward share obtained based on delegated amount, validator's total staking, and validator's commission rate

Assuming the network has 100,000,000 JU total staking and an annual inflation rate of 3%, the annualized yield is approximately 5-6%, depending specifically on the validator's commission rate and network usage.

### Inflation Control
The JPoSA mechanism controls inflation through a variable annual issuance rate:

- Annual issuance rate: Approximately 3%
- Inflation rate dynamically adjusts with network total token supply growth
- Partial inflation impact offset through transaction fee burning

This inflation model ensures the network has sufficient incentives to maintain secure operation while avoiding excessive dilution of token holders' equity.

## Miner (Validator) Role

In the JPoSA network, miners are called "validators" and are responsible for producing new blocks and maintaining network security. Becoming a validator requires meeting strict conditions and going through a governance process.

### Process to Become a Validator

1. **Proposal Stage**
   - Addition of new validators must be proposed by existing validators (or applicants)
   - Proposals need to receive voting support from more than half of active validators

2. **Staking Stage**
   - After proposal passes, applicants must complete staking registration within 7 days
   - Minimum self-staking requirement: MIN_VALIDATOR_STAKE (governable parameter, default: 100,000 JU)
   - Set commission rate (0-100%, default 10%, adjustable according to personal strategy)

3. **Activation Stage**
   - After staking is completed, wait for the next cycle (approximately 24 hours)
   - The system selects the top MAX_VALIDATORS validators based on staking weight
   - Officially begin block production and receive rewards### Validator Core Functions

1. **Block Production**
   - Produce new blocks in rotation order
   - Verify and package transactions
   - Maintain network security and stability

2. **Staking Management**
   - Increase self-staking to improve ranking
   - Adjust commission rates to attract delegators
   - Reduce staking when necessary (must meet minimum requirements)

3. **Reward Withdrawal**
   - Regularly withdraw block rewards and commissions
   - Rewards include basic rewards and transaction fee shares

4. **Delegation Management**
   - View delegator information and support
   - Maintain good service to attract more delegations

### Validator Responsibilities and Risks

1. **Online Time**
   - Must keep nodes continuously online and synchronized
   - Missing block production will be punished

2. **Punishment Mechanism**
   - Over 24 blocks, suspend transaction fee income and forfeit transaction fees, distributed to other validators, but basic staking income remains
   - Over 48 blocks, jailed and removed from validator set, but in the current epoch can still continue block production and receive basic staking income, but in the next epoch will be permanently kicked out, requiring re-proposal application to become a validator again
   - Epoch cycle is 24 hours

3. **Exit Mechanism**
   - Reduce staking: As long as minimum staking quantity is met, can reduce anytime
   - Exit mechanism: Active validators in current epoch are not allowed to exit, must first request exit application, wait until next epoch, will be removed from active validator list, then can exit
   - Staking refund: After validator exits, staking needs to wait 7 days before unlocking and returning

Validators play a crucial role in the network, serving as both infrastructure maintainers and important participants in the governance system.

## User (Delegator) Role

Ordinary users can participate in network maintenance and earn rewards by delegating JU tokens to validators. These users are called "delegators".

### Delegation Functions

1. **Selecting Validators**
   - Choose suitable validators based on their historical performance, commission rates, and reputation
   - Can delegate tokens to multiple validators to diversify risk
   - Minimum delegation amount: 1 JU

2. **Delegation Staking**
   - Send JU tokens to selected validator addresses for delegation
   - Delegated amounts are immediately counted toward the validator's total staking weight
   - Delegators share the validator's rewards based on staking proportion

3. **Reward Withdrawal**
   - Regularly withdraw delegation earnings
   - Rewards include the portion belonging to oneself from the 90% of basic rewards
   - Reward withdrawal requires no unbonding period and can be withdrawn anytime

### Unbonding and Withdrawal

1. **Delegation Unbonding**
   - After initiating an unbonding request, delegated amounts enter a 7-day unbonding period
   - During the unbonding period (unbonded portion) corresponding rewards will not be received
   - Principal can be withdrawn after unbonding period ends

2. **Fund Withdrawal**
   - Principal can be withdrawn after unbonding period expires
   - Withdrawal operations are completed instantly with no additional delay

### Delegation Strategy Recommendations

1. **Diversified Delegation**
   - Don't delegate all tokens to a single validator
   - Diversify delegation among 3-5 well-performing validators

2. **Monitor Validator Performance**
   - Regularly check validators' online time and reward distribution situations
   - Timely replace poorly performing validators

3. **Commission Rate Considerations**
   - Low commission rate validators can bring higher returns
   - But also need to consider service quality, not just commission rates

The delegator mechanism enables users who don't have the conditions to run validator nodes to participate in network maintenance and earn rewards, greatly enhancing the network's degree of decentralization and user participation.

## Governance Committee Role

The governance committee of the JPoSA network consists of all active validators and they are the core participants in network governance decisions.

### Governance Responsibilities

1. **Proposal Review**
   - Review and evaluate submitted various proposals
   - Analyze the impact of proposals on network security and performance
   - Vote on whether proposals should pass

2. **Parameter Adjustment**
   - Adjust key parameters according to network development needs
   - Balance degree of decentralization and network performance
   - Control inflation rate and reward distribution mechanism

3. **Validator Management**
   - Review qualifications of new validators
   - Make penalty decisions for jailed validators
   - Maintain quality of validator set

### Proposal Types

1. **Validator-Related Proposals**
   - Adding new validators
   - Removing unqualified validators

2. **System Parameter Proposals**
   - Adjusting punishment thresholds
   - Modifying unbonding period duration
   - Updating block reward amounts
   - Adjusting commission rate caps

### Voting Mechanism

1. **Voting Threshold**
   - All proposals (including ordinary proposals and major changes): Require agreement from more than half of active validators

2. **Voting Period**
   - Default proposal validity period is 7 days
   - Can vote anytime during validity period
   - Vote results are tallied after expiration

3. **Voting Weight**
   - Each active validator has equal one vote
   - Voting weights are not allocated based on staking volume
   - Ensures equality of governance power

The governance committee mechanism ensures that important network decisions are made jointly by validators who actually maintain network security, embodying the concept of true decentralized governance.

## Governable Parameters

The JPoSA network provides rich governable parameters, allowing validators to adjust according to network development needs:

### Time-Related Parameters

1. **Proposal Validity Period** (proposalLastingPeriod)
   - Range: 3,600 blocks (approximately 1 hour) to 2,592,000 blocks (approximately 30 days)
   - Default value: 604,800 blocks (approximately 7 days)
   - Used to control proposal voting cycles

2. **Unbonding Period** (unbondingPeriod)
   - Default value: 604,800 blocks (approximately 7 days)
   - Controls locking time for delegated funds and staked funds
   - Affects fund liquidity

3. **Validator Unjail Period** (validatorUnjailPeriod)
   - Default value: 86,400 blocks (approximately 24 hours)
   - Time required for jailed validators to return to normal state

4. **Commission Update Cooldown** (commissionUpdateCooldown)
   - Default value: 604,800 blocks (approximately 7 days)
   - Minimum interval between validator commission rate updates

5. **Proposal Cooldown** (proposalCooldown)
   - Default value: 100 blocks
   - Minimum interval between proposals created by the same validator

### Punishment-Related Parameters

1. **Punishment Threshold** (punishThreshold)
   - Default value: 24 blocks
   - Income suspension for validators when this threshold is reached

2. **Removal Threshold** (removeThreshold)
   - Default value: 48 blocks
   - Jailing and removal of validators when this threshold is reached

3. **Decrease Rate** (decreaseRate)
   - Default value: 24
   - Controls punishment counter reduction ratio, used to mitigate punishments for jailed validators, reducing punished blocks by removeThreshold/decreaseRate per epoch

4. **Validator Unjail Period** (validatorUnjailPeriod)
   - Default value: 86,400 blocks (approximately 24 hours)
   - Time required for jailed validators to return to normal state after removal, only after this time can validators reapply to join the validator set

5. **Double-Sign Slash Amount** (doubleSignSlashAmount)
   - Default value: 50,000 JU
   - Absolute slash applied to validator self-stake upon double-sign evidence

6. **Double-Sign Reporter Reward** (doubleSignRewardAmount)
   - Default value: 10,000 JU
   - Reward paid to the evidence reporter (<= slash amount)

7. **Double-Sign Evidence Window** (doubleSignWindow)
   - Default value: 86,400 blocks (approximately 24 hours)
   - Evidence must be submitted within this window

8. **Burn Address** (burnAddress)
   - Default value: `0x000000000000000000000000000000000000dEaD`
   - Receives the slashed remainder after reporter reward

### Reward-Related Parameters

1. **Profit Withdrawal Cycle** (withdrawProfitPeriod)
   - Controls minimum interval for validators to withdraw transaction fees
   - Default value: 86,400 blocks (approximately 24 hours)

2. **Block Reward** (blockReward)
   - Basic reward amount produced per block
   - Default value: 0.2 JU

3. **Base Reward Ratio** (baseRewardRatio)
   - Default value: 3000 (30.00%)
   - Base reward ratio used for reward distribution (0-10000)

4. **Max Commission Rate** (maxCommissionRate)
   - Default value: 6000 (60.00%)
   - Upper bound for validator commission rates (0-10000)

5. **Commission Rate Base** (COMMISSION_RATE_BASE)
   - Used to calculate validator commission rates
   - 10000 represents 100%

### Technical Parameters

1. **Maximum Validator Count** (maxValidators)
   - Upper limit of simultaneously active validators in the network
   - Default value: 21

2. **Minimum Validator Stake** (minValidatorStake)
   - Minimum staking amount required to become a validator
   - Default value: 100,000 JU
   - Ensures validators have sufficient economic incentives to maintain network security

3. **Minimum Delegation** (minDelegation)
   - Minimum delegation amount per delegator
   - Default value: 10 JU

4. **Minimum Undelegation** (minUndelegation)
   - Minimum undelegation amount per delegator
   - Default value: 1 JU

All parameter adjustments require governance proposals and sufficient votes to take effect, ensuring transparency and security of network parameter changes.

## Summary

The JPoSA consensus mechanism successfully combines the advantages of PoS and PoA through innovative hybrid design, providing blockchain networks with high performance, high security, and good decentralization characteristics.

### Core Values

1. **Balance Between Decentralization and Performance**
   - Select validators based on staking weight to enhance degree of decentralization
   - Maintain PoA efficiency to ensure fast transaction confirmation

2. **Sustainable Economic Model**
   - Reasonable reward distribution mechanism incentivizes all parties to actively participate
   - 3% annual issuance rate balances incentives and inflation control
   - Delegation mechanism allows ordinary users to participate in network maintenance and benefit

3. **Sound Governance System**
   - Governance committee composed of validators who actually maintain the network
   - Transparent proposal and voting mechanisms
   - Rich adjustable parameters adapt to network development needs

4. **Strong Security Assurance**
   - Multi-layer security protection mechanisms
   - Sound punishment system constrains validator behavior
   - 7-day unbonding period prevents malicious exits
