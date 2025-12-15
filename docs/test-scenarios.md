# JuChain POSA System Test Scenarios List

This document provides a complete list of test scenarios for manually verifying the correctness of contracts and consensus logic.

## 📋 Test Environment Preparation

### Prerequisites
- ✅ 3 validator nodes running (validator1, validator2, validator3)
- ✅ 1 sync node running
- ✅ `congress-cli` tool compiled
- ✅ All validator accounts have sufficient balance (at least 10,000 JU + Gas)
- ✅ Record current block height: `___________`

### Test Account Preparation
- **Validator1**: `0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266`
- **Validator2**: `0x70997970C51812dc3A010C7d01b50e0d17dc79C8`
- **Validator3**: `0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC`
- **New Validator Candidate**: `0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc` (validator6)
- **Delegator Account**: `0x970e8128ab834e3eac664312d6e30df9e93cb357`

---

## I. Validator Lifecycle Testing

### Test Scenario 1.1: Complete New Validator Registration Process

**Test Objective**: Verify the complete process of a new validator from proposal to activation

**Prerequisites**:
- ✅ 3 validators running normally
- ✅ New validator account has at least 10,000 JU
- ✅ Record current block height: `___________`

**Steps**:

1. **Create Add Validator Proposal**
   ```bash
   ./build/congress-cli create_proposal \
     -c 202599 -l http://localhost:8545 \
     -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
     -t 0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc \
     -o add
   
   ./build/congress-cli sign -f createProposal.json \
     -k /path/to/validator1/keystore/UTC--xxx \
     -p /path/to/validator1/password.txt \
     -c 202599
   
   ./build/congress-cli send -f createProposal_signed.json -c 202599 -l http://localhost:8545
   ```
   - Record proposal ID: `___________`
   - Record proposal creation block: `___________`

2. **Validator Voting (need at least 2/3 agreement)**
   ```bash
   # Validator1 vote (approve)
   ./build/congress-cli vote_proposal \
     -c 202599 -l http://localhost:8545 \
     -s 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
     -i <PROPOSAL_ID> \
     -a
   ./build/congress-cli sign -f voteProposal.json -k validator1_keystore -p password -c 202599
   ./build/congress-cli send -f voteProposal_signed.json -c 202599 -l http://localhost:8545
   
   # Validator2 vote (approve)
   ./build/congress-cli vote_proposal \
     -c 202599 -l http://localhost:8545 \
     -s 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 \
     -i <PROPOSAL_ID> \
     -a
   ./build/congress-cli sign -f voteProposal.json -k validator2_keystore -p password -c 202599
   ./build/congress-cli send -f voteProposal_signed.json -c 202599 -l http://localhost:8545
   
   # Validator3 vote (optional, test if 2 votes are sufficient)
   ```
   - Record vote completion block: `___________`
   - Verify proposal passed: `___________`

3. **Wait 7 Days Registration Period**
   - Record proposal approval timestamp: `___________`
   - Wait 7 days then continue (or modify system time for testing)

4. **Validator Registration and Staking**
   ```bash
   # First transfer 10,000 JU to new validator
   # Then register
   ./build/congress-cli staking register-validator \
     -c 202599 -l http://localhost:8545 \
     --proposer 0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc \
     --stake-amount 10000 \
     --commission-rate 500
   
   ./build/congress-cli sign -f registerValidator.json \
     -k validator6_keystore -p password -c 202599
   
   ./build/congress-cli send -f registerValidator_signed.json -c 202599 -l http://localhost:8545
   ```
   - Record registration block: `___________`

5. **Wait for Next Epoch Update**
   - Current block: `___________`
   - Next Epoch block: `___________` (current block rounded up to multiple of 86400)
   - Wait for Epoch update

6. **Verify Validator Entered Validator Set**
   ```bash
   ./build/congress-cli staking list-top-validators -c 202599 -l http://localhost:8545
   ./build/congress-cli miners -c 202599 -l http://localhost:8545
   ```

**Expected Results**:
- ✅ Proposal created successfully, obtained proposal ID
- ✅ 2 validators voted, proposal passed (`pass[validator] = true`)
- ✅ Registration successful within 7 days (`isProposalValidForStaking()` returns true)
- ✅ Registered validator appears in `allValidators`
- ✅ Next Epoch, validator enters `currentValidatorSet`
- ✅ Validator can start producing blocks

**Actual Results**:
```
[User to fill in]
```

---

### Test Scenario 1.2: Validator Registration Timeout (7-Day Limit)

**Test Objective**: Verify that proposals approved but not registered within 7 days cannot be registered

**Prerequisites**:
- ✅ Have an approved but unregistered proposal
- ✅ Proposal approval time has exceeded 7 days

**Steps**:

1. **Attempt Registration (Over 7 Days)**
   ```bash
   ./build/congress-cli staking register-validator \
     -c 202599 -l http://localhost:8545 \
     --proposer 0xNew validator address \
     --stake-amount 10000 \
     --commission-rate 500
   
   ./build/congress-cli sign -f registerValidator.json -k keystore -p password -c 202599
   ./build/congress-cli send -f registerValidator_signed.json -c 202599 -l http://localhost:8545
   ```

**Expected Results**:
- ❌ Transaction fails, error message contains "Proposal expired, must repropose"
- ❌ `isProposalValidForStaking()` returns false

**Actual Results**:
```
[User to fill in]
```

---

### Test Scenario 1.3: Validator Increase Stake

**Test Objective**: Verify that registered validators can increase their stake

**Prerequisites**:
- ✅ Validator registered and staked at least 10,000 JU
- ✅ Validator account has additional balance

**Steps**:

1. **Query Current Stake**
   ```bash
   ./build/congress-cli staking query-validator \
     -c 202599 -l http://localhost:8545 \
     --address 0xValidator address
   ```
   - Record current self-stake: `___________`

2. **Increase Stake**
   ```bash
   # Note: Need to call contract directly or use other tools
   # Here need to manually construct transaction or use web3
   ```

3. **Verify Stake Increase**
   ```bash
   ./build/congress-cli staking query-validator \
     -c 202599 -l http://localhost:8545 \
     --address 0xValidator address
   ```

**Expected Results**:
- ✅ Stake increased successfully
- ✅ `selfStake` updated to new value
- ✅ Next Epoch, validator ranking may improve

**Actual Results**:
```
[User to fill in]
```

---

### Test Scenario 1.4: Validator Partial Stake Withdrawal

**Test Objective**: Verify that validators can partially withdraw stake (remaining >= 10,000 JU)

**Prerequisites**:
- ✅ Validator staked > 20,000 JU
- ✅ Validator not in `currentValidatorSet` (or can accept temporary exit)

**Steps**:

1. **Partially Withdraw Stake**
   ```bash
   # Need to call Staking.withdrawValidatorStake(amount) directly
   # Ensure remainingStake >= 10,000 JU
   ```

2. **Verify Withdrawal Results**
   ```bash
   ./build/congress-cli staking query-validator \
     -c 202599 -l http://localhost:8545 \
     --address 0xValidator address
   ```

**Expected Results**:
- ✅ Withdrawal successful
- ✅ Remaining stake >= 10,000 JU
- ✅ Validator still valid (if remaining stake >= MIN_VALIDATOR_STAKE)

**Actual Results**:
```
[User to fill in]
```

---

### Test Scenario 1.5: Validator Emergency Exit (emergencyExit)

**Test Objective**: Verify that validators can completely exit, check minimum validator count protection

**Prerequisites**:
- ✅ At least 4 active validators (ensure >= 3 after exit)
- ✅ Target validator registered and staked

**Steps**:

1. **Query Current Active Validator Count**
   ```bash
   ./build/congress-cli miners -c 202599 -l http://localhost:8545
   ```
   - Record active validator count: `___________`

2. **Perform Emergency Exit**
   ```bash
   # Call Staking.emergencyExit()
   # If validator is in currentValidatorSet, will be jailed first
   ```

3. **Verify Exit Results**
   ```bash
   ./build/congress-cli staking query-validator \
     -c 202599 -l http://localhost:8545 \
     --address 0xValidator address
   
   ./build/congress-cli miners -c 202599 -l http://localhost:8545
   ```

**Expected Results**:
- ✅ If validator is in `currentValidatorSet`, first jailed (1 epoch)
- ✅ After exit, remaining validator count >= 3
- ✅ Validator removed from `allValidators`
- ✅ `selfStake` becomes 0
- ✅ Stake amount transferred back to validator account

**Actual Results**:
```
[User to fill in]
```

---

### Test Scenario 1.6: Emergency Exit When Validator Count Insufficient

**Test Objective**: Verify that exit is not allowed when only 3 validators remain

**Prerequisites**:
- ✅ Only 3 active validators
- ✅ Target validator registered and staked

**Steps**:

1. **Attempt Emergency Exit**
   ```bash
   # Call Staking.emergencyExit()
   ```

**Expected Results**:
- ❌ Transaction fails
- ❌ Error message contains "Cannot exit: would leave less than minimum validators"
- ❌ Validator still exists

**Actual Results**:
```
[User to fill in]
```

---

## II. Validator Punishment Testing

### Test Scenario 2.1: Validator Misses Block Production (Minor Punishment)

**Test Objective**: Verify the punishment mechanism when validators miss block production

**Prerequisites**:
- ✅ Validator running normally
- ✅ Record validator's current `missedBlocksCounter`: `___________`

**Steps**:

1. **Stop Validator Node**
   ```bash
   # Stop validator2 node
   pm2 stop ju-chain-validator2
   ```

2. **Wait to Miss Multiple Blocks**
   - Record stop block: `___________`
   - Wait to miss about 10-20 blocks

3. **Check Punishment Records**
   ```bash
   # Query missedBlocksCounter in Punish contract
   # Or check node logs
   ```

4. **Resume Validator Node**
   ```bash
   pm2 start ju-chain-validator2
   ```

**Expected Results**:
- ✅ `missedBlocksCounter` increases
- ✅ If reaches 24 blocks (punishThreshold), validator's income removed
- ✅ Validator can still produce blocks (not reaching removeThreshold)

**Actual Results**:
```
[User to fill in]
```

---

### Test Scenario 2.2: Validator Reaches Punishment Threshold (24 Blocks)

**Test Objective**: Verify that validator's income is removed when reaching punishThreshold

**Prerequisites**:
- ✅ Validator has missed some blocks
- ✅ `missedBlocksCounter` approaching 24

**Steps**:

1. **Stop Validator Node**
   ```bash
   pm2 stop ju-chain-validator2
   ```

2. **Wait to Miss 24 Blocks**
   - Record stop block: `___________`
   - Wait to miss 24 blocks

3. **Check Validator Income**
   ```bash
   ./build/congress-cli miner \
     -c 202599 -l http://localhost:8545 \
     -a 0x70997970C51812dc3A010C7d01b50e0d17dc79C8
   ```

4. **Resume Validator Node**
   ```bash
   pm2 start ju-chain-validator2
   ```

**Expected Results**:
- ✅ Triggered when `missedBlocksCounter % 24 == 0`
- ✅ Validator's income (`aacIncoming`) removed (becomes 0)
- ✅ Validator can still produce blocks (not reaching removeThreshold)

**Actual Results**:
```
[User to fill in]
```

---

### Test Scenario 2.3: Validator Reaches Removal Threshold (48 Blocks) - Jail and Remove

**Test Objective**: Verify that validator is jailed and removed when reaching removeThreshold

**Prerequisites**:
- ✅ Validator running normally
- ✅ At least 4 validators (ensure >= 3 after removal)

**Steps**:

1. **Stop Validator Node**
   ```bash
   pm2 stop ju-chain-validator2
   ```

2. **Wait to Miss 48 Blocks**
   - Record stop block: `___________`
   - Wait to miss 48 blocks (about 48 seconds)

3. **Check Validator Status**
   ```bash
   ./build/congress-cli staking query-validator \
     -c 202599 -l http://localhost:8545 \
     --address 0x70997970C51812dc3A010C7d01b50e0d17dc79C8
   
   ./build/congress-cli miners -c 202599 -l http://localhost:8545
   ```

4. **Verify Validator Cannot Produce Blocks**
   - Check node logs, confirm validator no longer produces blocks

**Expected Results**:
- ✅ Triggered when `missedBlocksCounter % 48 == 0`
- ✅ Validator jailed (`isJailed = true`, `jailUntilBlock = block.number + 86400`)
- ✅ Validator removed from `currentValidatorSet` (next Epoch)
- ✅ Validator removed from `highestValidatorsSet` (if length > 1)
- ✅ `pass[validator] = false` (proposal status cleared)
- ✅ `violationCount[validator]++` (violation count increases)
- ✅ Validator cannot produce blocks (not in `snap.Validators`)

**Actual Results**:
```
[User to fill in]
```

---

### Test Scenario 2.4: Validator Unjail (3 or Fewer Violations)

**Test Objective**: Verify that validators with <= 3 violations can recover automatically

**Prerequisites**:
- ✅ Validator jailed
- ✅ `violationCount <= 3`
- ✅ Imprisonment period passed (`block.number >= jailUntilBlock`)

**Steps**:

1. **Wait for Imprisonment Period to End**
   - Record jail block: `___________`
   - Record jailUntilBlock: `___________`
   - Wait until `jailUntilBlock`

2. **Query Violation Count**
   ```bash
   # Query violationCount in Proposal contract
   ```

3. **Perform Unjail**
   ```bash
   # Call Staking.unjailValidator()
   ```

4. **Verify Recovery Results**
   ```bash
   ./build/congress-cli staking query-validator \
     -c 202599 -l http://localhost:8545 \
     --address 0xValidator address
   
   # Check if pass status automatically recovered
   ```

**Expected Results**:
- ✅ Unjail successful
- ✅ `isJailed = false`
- ✅ `pass[validator] = true` (automatically recovered)
- ✅ `proposalPassedTime[validator] = block.timestamp` (updated time)
- ✅ Validator can re-enter validator set (next Epoch)

**Actual Results**:
```
[User to fill in]
```

---

### Test Scenario 2.5: Validator Unjail Failed (4 or More Violations)

**Test Objective**: Verify that validators with >= 4 violations cannot unjail and need to re-propose

**Prerequisites**:
- ✅ Validator jailed
- ✅ `violationCount >= 4`
- ✅ Imprisonment period passed

**Steps**:

1. **Attempt Unjail**
   ```bash
   # Call Staking.unjailValidator()
   ```

2. **Verify Failure**
   ```bash
   # Check if transaction failed
   ```

3. **Re-propose and Vote**
   ```bash
   # Create proposal
   ./build/congress-cli create_proposal \
     -c 202599 -l http://localhost:8545 \
     -p 0xOther validator address \
     -t 0xJailed validator address \
     -o add
   
   # Vote to pass
   # ...
   ```

4. **Verify Violation Count Reset**
   ```bash
   # Query violationCount, should be 0
   ```

5. **Attempt Unjail Again**
   ```bash
   # Call Staking.unjailValidator()
   ```

**Expected Results**:
- ❌ First unjail fails (require check fails)
- ✅ After re-proposal and voting passes, `violationCount` resets to 0
- ✅ After vote passes, `pass[validator] = true`
- ✅ Second unjail successful (because violationCount reset to 0)

**Actual Results**:
```
[User to fill in]
```

---

### Test Scenario 2.6: Jail-Removed Validators Excluded Immediately During Epoch

**Test Objective**: Verify that validators jailed during Epoch block are immediately excluded from validator set

**Prerequisites**:
- ✅ Validator running normally
- ✅ Approaching Epoch block

**Steps**:

1. **Calculate Next Epoch Block**
   - Current block: `___________`
   - Next Epoch: `___________` (rounded up to multiple of 86400)

2. **Stop Validator Before Epoch Block**
   ```bash
   # Stop 1-2 blocks before Epoch block
   pm2 stop ju-chain-validator2
   ```

3. **Wait for Epoch Block Processing**
   - Observe Epoch block processing

4. **Check Validator Set**
   ```bash
   ./build/congress-cli staking list-top-validators -c 202599 -l http://localhost:8545
   ./build/congress-cli miners -c 202599 -l http://localhost:8545
   ```

**Expected Results**:
- ✅ Validators jailed during Epoch block immediately excluded from `currentValidatorSet`
- ✅ List returned by `getTopValidators()` does not include jailed validators
- ✅ `header.Extra` (based on parent state) may include this validator, but `newValidators` (based on current state) does not
- ✅ Epoch validation allows this inconsistency (POSA mode)

**Actual Results**:
```
[User to fill in]
```

---

## III. Delegation and Reward Testing

### Test Scenario 3.1: Delegate Tokens to Validator

**Test Objective**: Verify that users can delegate tokens to validators

**Prerequisites**:
- ✅ Validator registered and active
- ✅ Delegator account has sufficient balance (at least 1 JU)

**Steps**:

1. **Query Validator Information**
   ```bash
   ./build/congress-cli staking query-validator \
     -c 202599 -l http://localhost:8545 \
     --address 0xValidator address
   ```
   - Record current total delegation: `___________`

2. **Delegate Tokens**
   ```bash
   ./build/congress-cli staking delegate \
     -c 202599 -l http://localhost:8545 \
     --delegator 0x970e8128ab834e3eac664312d6e30df9e93cb357 \
     --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
     --amount 1000
   
   ./build/congress-cli sign -f delegate.json -k delegator_keystore -p password -c 202599
   ./build/congress-cli send -f delegate_signed.json -c 202599 -l http://localhost:8545
   ```

3. **Verify Delegation Results**
   ```bash
   ./build/congress-cli staking query-delegation \
     -c 202599 -l http://localhost:8545 \
     --delegator 0x970e8128ab834e3eac664312d6e30df9e93cb357 \
     --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
   
   ./build/congress-cli staking query-validator \
     -c 202599 -l http://localhost:8545 \
     --address 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
   ```

**Expected Results**:
- ✅ Delegation successful
- ✅ Validator's `totalDelegated` increases
- ✅ Delegator's `delegations[delegator][validator].amount` updated
- ✅ Next Epoch, validator ranking may improve (if total stake increases)

**Actual Results**:
```
[User to fill in]
```

---

### Test Scenario 3.2: Undelegate (Start Unbonding Period)

**Test Objective**: Verify that undelegation enters 7-day unbonding period

**Prerequisites**:
- ✅ Delegator has delegated tokens to validator
- ✅ Record current block: `___________`

**Steps**:

1. **Query Delegation Information**
   ```bash
   ./build/congress-cli staking query-delegation \
     -c 202599 -l http://localhost:8545 \
     --delegator 0x970e8128ab834e3eac664312d6e30df9e93cb357 \
     --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
   ```

2. **Undelegate**
   ```bash
   ./build/congress-cli staking undelegate \
     -c 202599 -l http://localhost:8545 \
     --delegator 0x970e8128ab834e3eac664312d6e30df9e93cb357 \
     --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
     --amount 500
   
   ./build/congress-cli sign -f undelegate.json -k delegator_keystore -p password -c 202599
   ./build/congress-cli send -f undelegate_signed.json -c 202599 -l http://localhost:8545
   ```

3. **Verify Unbonding Status**
   ```bash
   ./build/congress-cli staking query-delegation \
     -c 202599 -l http://localhost:8545 \
     --delegator 0x970e8128ab834e3eac664312d6e30df9e93cb357 \
     --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
   ```
   - Record unbonding completion block: `___________` (current block + unbondingPeriod, default 604800)

**Expected Results**:
- ✅ Undelegation successful
- ✅ Delegation amount decreases
- ✅ Unbonding record created (`unbondingDelegations`)
- ✅ `unbondingAmount` increases
- ✅ `unbondingBlock = block.number + 604800` (after 7 days)
- ✅ During unbonding period, tokens still counted toward validator's total stake

**Actual Results**:
```
[User to fill in]
```

---

### Test Scenario 3.3: Withdraw Unbonded Tokens

**Test Objective**: Verify that tokens can be withdrawn after unbonding period ends

**Prerequisites**:
- ✅ Have unbonding tokens
- ✅ Unbonding period passed (`block.number >= unbondingBlock`)

**Steps**:

1. **Wait for Unbonding Period to End**
   - Current block: `___________`
   - Unbonding completion block: `___________`
   - Wait until unbonding completion block

2. **Withdraw Unbonded Tokens**
   ```bash
   # Call Staking.withdrawUnbonded(validator, maxEntries)
   ```

3. **Verify Withdrawal Results**
   ```bash
   ./build/congress-cli staking query-delegation \
     -c 202599 -l http://localhost:8545 \
     --delegator 0x970e8128ab834e3eac664312d6e30df9e93cb357 \
     --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
   ```

**Expected Results**:
- ✅ Withdrawal successful
- ✅ Tokens transferred to delegator account
- ✅ Unbonding records removed
- ✅ `unbondingAmount` decreases

**Actual Results**:
```
[User to fill in]
```

---

### Test Scenario 3.4: Reward Distribution and Withdrawal

**Test Objective**: Verify that block rewards are correctly distributed to validators and delegators

**Prerequisites**:
- ✅ Validator registered and active
- ✅ Delegators have delegated tokens
- ✅ Validator has produced blocks (received rewards)

**Steps**:

1. **Query Validator Rewards**
   ```bash
   ./build/congress-cli staking query-validator \
     -c 202599 -l http://localhost:8545 \
     --address 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
   ```
   - Record `accumulatedRewards`: `___________`

2. **Query Delegator Rewards**
   ```bash
   ./build/congress-cli staking query-delegation \
     -c 202599 -l http://localhost:8545 \
     --delegator 0x970e8128ab834e3eac664312d6e30df9e93cb357 \
     --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
   ```
   - Record `pendingRewards`: `___________`

3. **Validator Withdraw Rewards**
   ```bash
   ./build/congress-cli staking claim-rewards \
     -c 202599 -l http://localhost:8545 \
     --claimer 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
     --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
   
   ./build/congress-cli sign -f claimRewards.json -k validator_keystore -p password -c 202599
   ./build/congress-cli send -f claimRewards_signed.json -c 202599 -l http://localhost:8545
   ```

4. **Delegator Withdraw Rewards**
   ```bash
   ./build/congress-cli staking claim-rewards \
     -c 202599 -l http://localhost:8545 \
     --claimer 0x970e8128ab834e3eac664312d6e30df9e93cb357 \
     --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
   
   ./build/congress-cli sign -f claimRewards.json -k delegator_keystore -p password -c 202599
   ./build/congress-cli send -f claimRewards_signed.json -c 202599 -l http://localhost:8545
   ```

5. **Verify Withdrawal Results**
   ```bash
   # Query balance changes
   # Verify rewards cleared
   ```

**Expected Results**:
- ✅ Validator receives: commission + validator share
- ✅ Delegator receives: delegation share
- ✅ Reward calculation correct (based on `rewardPerShare` mechanism)
- ✅ After withdrawal `accumulatedRewards` and `pendingRewards` cleared
- ✅ Tokens correctly transferred to accounts

**Actual Results**:
```
[User to fill in]
```

---

### Test Scenario 3.5: Transaction Fee Reward Distribution

**Test Objective**: Verify that transaction fees are correctly distributed to all active validators

**Prerequisites**:
- ✅ Multiple active validators
- ✅ Send some transactions (generate fees)

**Steps**:

1. **Send Transactions to Generate Fees**
   ```bash
   # Send some transactions
   ```

2. **Query Validator Income**
   ```bash
   ./build/congress-cli miner \
     -c 202599 -l http://localhost:8545 \
     -a 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
   ```
   - Record `Accumulated Rewards`: `___________`

3. **Withdraw Transaction Fees**
   ```bash
   ./build/congress-cli withdraw_profits \
     -c 202599 -l http://localhost:8545 \
     -a 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
   
   ./build/congress-cli sign -f withdrawProfits.json -k validator_keystore -p password -c 202599
   ./build/congress-cli send -f withdrawProfits_signed.json -c 202599 -l http://localhost:8545
   ```

**Expected Results**:
- ✅ Transaction fees evenly distributed to all active validators (excluding jailed ones)
- ✅ Each validator receives: `totalReward / activeValidatorCount`
- ✅ Withdrawal successful, tokens transferred to `feeAddr`

**Actual Results**:
```
[User to fill in]
```

---

## IV. Epoch Update Testing

### Test Scenario 4.1: Epoch Update Validator Set

**Test Objective**: Verify that Epoch block correctly updates validator set

**Prerequisites**:
- ✅ Multiple validators (including newly registered)
- ✅ Approaching Epoch block

**Steps**:

1. **Record Current Validator Set**
   ```bash
   ./build/congress-cli staking list-top-validators -c 202599 -l http://localhost:8545
   ./build/congress-cli miners -c 202599 -l http://localhost:8545
   ```
   - Record current set: `___________`

2. **Wait for Epoch Block**
   - Current block: `___________`
   - Next Epoch: `___________`
   - Wait for Epoch block

3. **Verify Epoch Update**
   ```bash
   ./build/congress-cli staking list-top-validators -c 202599 -l http://localhost:8545
   ./build/congress-cli miners -c 202599 -l http://localhost:8545
   ```

**Expected Results**:
- ✅ `currentValidatorSet` updated
- ✅ `currentValidatorSet` updated
- ✅ `highestValidatorsSet` managed by other methods (e.g., `tryAddValidatorToHighestSet`)
- ✅ Newly registered validators (if stake sufficient) enter set
- ✅ Jailed validators excluded
- ✅ Validators sorted by total stake

**Actual Results**:
```
[User to fill in]
```

---

### Test Scenario 4.2: Validator Ranking Changes During Epoch

**Test Objective**: Verify that validator rankings update during Epoch when stake changes

**Prerequisites**:
- ✅ Multiple validators
- ✅ Validators have different stakes

**Steps**:

1. **Record Current Rankings**
   ```bash
   ./build/congress-cli staking list-top-validators -c 202599 -l http://localhost:8545
   ```

2. **Increase Stake or Delegation for a Validator**
   ```bash
   # Increase stake or delegation
   ```

3. **Wait for Next Epoch**
   - Wait for Epoch update

4. **Verify Ranking Changes**
   ```bash
   ./build/congress-cli staking list-top-validators -c 202599 -l http://localhost:8545
   ```

**Expected Results**:
- ✅ Validators re-sorted by total stake (`selfStake + totalDelegated`)
- ✅ Validator with increased stake improves ranking
- ✅ Ranking changes take effect during Epoch update

**Actual Results**:
```
[User to fill in]
```

---

### Test Scenario 4.3: Decrease Punishment Count During Epoch

**Test Objective**: Verify that `missedBlocksCounter` decreases during Epoch

**Prerequisites**:
- ✅ Validators punished (`missedBlocksCounter > 0`)

**Steps**:

1. **Record Punishment Count**
   ```bash
   # Query missedBlocksCounter in Punish contract
   ```
   - Record current count: `___________`

2. **Wait for Epoch Block**
   - Wait for next Epoch

3. **Verify Count Decrease**
   ```bash
   # Query missedBlocksCounter
   ```

**Expected Results**:
- ✅ `missedBlocksCounter` decreases during Epoch (`decreaseMissedBlocksCounter`)
- ✅ Decrease mechanism executes correctly

**Actual Results**:
```
[User to fill in]
```

---

## V. Edge Case Testing

### Test Scenario 5.1: Minimum Validator Count Protection

**Test Objective**: Verify that exit is not allowed when only 3 validators remain

**Prerequisites**:
- ✅ Only 3 active validators

**Steps**:

1. **Query Current Validator Count**
   ```bash
   ./build/congress-cli miners -c 202599 -l http://localhost:8545
   ```

2. **Attempt Emergency Exit**
   ```bash
   # Call Staking.emergencyExit()
   ```

**Expected Results**:
- ❌ Exit fails
- ❌ Error message contains minimum validator count requirement
- ✅ Validator still exists

**Actual Results**:
```
[User to fill in]
```

---

### Test Scenario 5.2: Maximum Validator Count Limit

**Test Objective**: Verify that maximum 21 validators allowed

**Prerequisites**:
- ✅ Nearly 21 validators exist

**Steps**:

1. **Attempt to Register 22nd Validator**
   ```bash
   # Complete proposal, voting, registration process
   ```

2. **Verify Entry into Validator Set**
   ```bash
   ./build/congress-cli staking list-top-validators -c 202599 -l http://localhost:8545
   ```

**Expected Results**:
- ✅ Registration successful (can register)
- ✅ But `getTopValidators()` only returns top 21
- ✅ 22nd validator not in `currentValidatorSet` (if stake insufficient)

**Actual Results**:
```
[User to fill in]
```

---

### Test Scenario 5.3: Delegate to Jailed Validator

**Test Objective**: Verify that delegation to jailed validators is not allowed

**Prerequisites**:
- ✅ Validator jailed

**Steps**:

1. **Attempt Delegation**
   ```bash
   ./build/congress-cli staking delegate \
     -c 202599 -l http://localhost:8545 \
     --delegator 0xDelegator address \
     --validator 0xJailed validator address \
     --amount 1000
   ```

**Expected Results**:
- ❌ Transaction fails
- ❌ Error message contains "Validator is jailed" or "onlyActiveValidator"

**Actual Results**:
```
[User to fill in]
```

---

### Test Scenario 5.4: Validator Removal After Proposal Voting

**Test Objective**: Verify that votes from removed validators are not counted in threshold

**Prerequisites**:
- ✅ Ongoing proposal
- ✅ Validator has voted

**Steps**:

1. **Validator Votes**
   ```bash
   # Validator1 votes
   ```

2. **Remove Validator (Through Punishment)**
   ```bash
   # Have Validator1 jailed and removed
   ```

3. **Check Voting Threshold**
   ```bash
   # Query proposal status
   # Verify vote count is correct
   ```

**Expected Results**:
- ✅ Votes from removed validators not counted in `getActiveVoteCount()`
- ✅ Voting threshold based on current active validator count
- ✅ Proposal pass/reject decision correct

**Actual Results**:
```
[User to fill in]
```

---

### Test Scenario 5.5: Multiple Validators Jailed Simultaneously

**Test Objective**: Verify handling when multiple validators jailed simultaneously

**Prerequisites**:
- ✅ Multiple validators exist

**Steps**:

1. **Stop Multiple Validators Simultaneously**
   ```bash
   pm2 stop ju-chain-validator2
   pm2 stop ju-chain-validator3
   ```

2. **Wait to Reach removeThreshold**
   - Wait 48 blocks

3. **Check Validator Set**
   ```bash
   ./build/congress-cli staking list-top-validators -c 202599 -l http://localhost:8545
   ./build/congress-cli miners -c 202599 -l http://localhost:8545
   ```

**Expected Results**:
- ✅ Multiple validators jailed simultaneously
- ✅ All excluded from validator set
- ✅ At least 1 validator preserved (protection mechanism)
- ✅ Chain continues running

**Actual Results**:
```
[User to fill in]
```

---

## VI. Governance Proposal Testing

### Test Scenario 6.1: Create Configuration Update Proposal

**Test Objective**: Verify that system configuration update proposals can be created

**Prerequisites**:
- ✅ Validators running normally

**Steps**:

1. **Create Configuration Update Proposal**
   ```bash
   ./build/congress-cli create_config_proposal \
     -c 202599 -l http://localhost:8545 \
     -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
     -i 0 \
     -v 86400
   
   # -i 0: proposalLastingPeriod
   # -v 86400: new value (seconds)
   
   ./build/congress-cli sign -f createUpdateConfigProposal.json -k keystore -p password -c 202599
   ./build/congress-cli send -f createUpdateConfigProposal_signed.json -c 202599 -l http://localhost:8545
   ```

2. **Validator Voting**
   ```bash
   # Multiple validators vote
   ```

3. **Verify Configuration Update**
   ```bash
   # Query if configuration updated
   ```

**Expected Results**:
- ✅ Proposal created successfully
- ✅ Configuration updated after voting passes
- ✅ New configuration takes effect

**Actual Results**:
```
[User to fill in]
```

---

### Test Scenario 6.2: Remove Validator Proposal

**Test Objective**: Verify that validators can be removed through proposals

**Prerequisites**:
- ✅ Validator to be removed exists
- ✅ At least 4 validators (ensure >= 3 after removal)

**Steps**:

1. **Create Removal Proposal**
   ```bash
   ./build/congress-cli create_proposal \
     -c 202599 -l http://localhost:8545 \
     -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
     -t 0xValidator address to remove \
     -o remove
   ```

2. **Vote to Pass**
   ```bash
   # Multiple validators vote
   ```

3. **Verify Removal Results**
   ```bash
   ./build/congress-cli miners -c 202599 -l http://localhost:8545
   ```

**Expected Results**:
- ✅ Validator removed after proposal passes
- ✅ `pass[validator] = false`
- ✅ Validator removed from validator set
- ✅ `violationCount[validator]++`

**Actual Results**:
```
[User to fill in]
```

---

## VII. Performance and Security Testing

### Test Scenario 7.1: Large Volume Delegation Operations

**Test Objective**: Verify that system can handle large volume delegation operations

**Prerequisites**:
- ✅ Multiple validators
- ✅ Multiple delegator accounts

**Steps**:

1. **Batch Delegation**
   ```bash
   # Create multiple delegation transactions
   # Sign and send sequentially
   ```

2. **Verify System Performance**
   - Observe transaction confirmation time
   - Check Gas consumption

**Expected Results**:
- ✅ All delegations successful
- ✅ System performance normal
- ✅ State updates correct

**Actual Results**:
```
[User to fill in]
```

---

### Test Scenario 7.2: Reentrancy Attack Protection

**Test Objective**: Verify system protection against reentrancy attacks

**Prerequisites**:
- ✅ Understand reentrancy attack mechanisms
- ✅ Validator account has rewards available for withdrawal

**Steps**:

1. **Check Contract-Level Protection**
   - `Validators` and `Staking` contracts inherit `ReentrancyGuard`
   - Key functions use `nonReentrant` modifier

2. **Test Function-Level Protection**
   - Attempt reentrancy in `withdrawProfits()` (should fail)
   - Attempt reentrancy in `withdrawValidatorStake()` (should fail)
   - Attempt reentrancy in `claimRewards()` (should fail)

3. **Check Block-Level Protection**
   - `distributeBlockReward()` uses `operationsDone` flag
   - `updateActiveValidatorSet()` uses `operationsDone` flag

4. **Verify CEI Pattern**
   - Check if key functions follow Checks-Effects-Interactions pattern
   - State updates before external calls

**Expected Results**:
- ✅ Contract-level protection: `nonReentrant` modifier prevents reentrancy
- ✅ Function-level protection: reentrant calls rejected by `ReentrancyGuard`
- ✅ Block-level protection: same operation can only execute once per block
- ✅ CEI pattern: state updates before external calls, ensuring consistency
- ✅ Reentrancy attacks completely protected

**Actual Results**:
```
[User to fill in]
```

---

## VIII. Comprehensive Scenario Testing

### Test Scenario 8.1: Complete Validator Lifecycle

**Test Objective**: Completely test validator full lifecycle from registration to exit

**Prerequisites**:
- ✅ New validator account

**Steps**:

1. **Proposal and Registration** (refer to scenario 1.1)
2. **Increase Stake** (refer to scenario 1.3)
3. **Receive Delegation** (refer to scenario 3.1)
4. **Withdraw Rewards** (refer to scenario 3.4)
5. **Get Punished and Recover** (refer to scenarios 2.3 and 2.4)
6. **Emergency Exit** (refer to scenario 1.5)

**Expected Results**:
- ✅ All steps executed successfully
- ✅ State transitions correct
- ✅ Data consistency maintained

**Actual Results**:
```
[User to fill in]
```

---

### Test Scenario 8.2: Network Stress Testing

**Test Objective**: Verify system stability under high load

**Prerequisites**:
- ✅ System running normally

**Steps**:

1. **Execute Multiple Operations Simultaneously**
   - Multiple validator registrations
   - Large volume delegation operations
   - Multiple proposals and voting
   - Reward withdrawals

2. **Monitor System Status**
   - Check block production speed
   - Check transaction confirmation time
   - Check Gas consumption

**Expected Results**:
- ✅ System runs stably
- ✅ All operations eventually succeed
- ✅ Performance within acceptable range

**Actual Results**:
```
[User to fill in]
```

---

## Test Checklist

### Functional Checks
- [ ] Complete validator registration process
- [ ] Validator punishment mechanism correct
- [ ] Delegation and unbonding functions normal
- [ ] Reward distribution accurate
- [ ] Epoch updates correct
- [ ] Governance proposals work normally

### Security Checks
- [ ] Minimum validator count protection
- [ ] Reentrancy attack protection (ReentrancyGuard + nonReentrant)
- [ ] Edge case handling
- [ ] State consistency
- [ ] Configuration parameter validation (prevent division by zero, etc.)
- [ ] CEI pattern verification (Checks-Effects-Interactions)

### Performance Checks
- [ ] Transaction confirmation time
- [ ] Gas consumption reasonable
- [ ] System stability

---

## Test Record Template

### Test Environment Information
- **Test Date**: `___________`
- **Tester**: `___________`
- **Test Environment**: `___________` (Mainnet/Testnet/Local)
- **Node Version**: `___________`
- **Contract Version**: `___________`

### Test Results Summary
- **Total Test Scenarios**: `___________`
- **Passed Scenarios**: `___________`
- **Failed Scenarios**: `___________`
- **Skipped Scenarios**: `___________`

### Issues Found
1. `___________`
2. `___________`
3. `___________`

### Improvement Suggestions
1. `___________`
2. `___________`
3. `___________`

---

**Document Version**: v1.1.0  
**Creation Date**: 2025-01-21  
**Last Updated**: 2025-01-21

**Update Content (v1.1.0)**:
- Updated reentrancy attack protection test scenarios: Added ReentrancyGuard and nonReentrant testing
- Updated security mechanism description: Improved CEI pattern verification
- Updated configuration parameter testing: Removed inflation-related tests (cid 5 and 6 removed)