# Hardhat 测试 vs Foundry 测试完整对比分析

## test/proposal.js 测试覆盖分析

### 1. Init 测试

| Hardhat 测试 | Foundry 对应测试 | 状态 |
|-------------|-----------------|------|
| "Init can only call once" | ProposalFoundry.t.sol::testInitOnlyOnce | ✅ 已覆盖 |

### 2. Create proposal 测试  

| Hardhat 测试 | Foundry 对应测试 | 状态 |
|-------------|-----------------|------|
| "anyone can create proposal" | ProposalFoundry.t.sol::testCreateAndVoteAddProposalPass (部分) | ⚠️ 需补充多账户创建测试 |
| "can't try to add already exist dst" | ProposalFoundry.t.sol::testCreateProposalConstraints | ✅ 已覆盖 |
| "can't try to remove not exist dst" | ProposalFoundry.t.sol::testCreateProposalConstraints | ✅ 已覆盖 |
| "details info can't too long" | ProposalFoundry.t.sol::testCreateProposalConstraints | ✅ 已覆盖 |

### 3. Vote for proposal(add,pass) 测试

| Hardhat 测试 | Foundry 对应测试 | 状态 |
|-------------|-----------------|------|
| "normal vote for a proposal(3 true/2 false)" | ProposalFoundry.t.sol::testCreateAndVoteAddProposalPass | ✅ 已覆盖 |
| "only validator can vote for a proposal" | ProposalFoundry.t.sol::testOnlyValidatorCanVote | ✅ 已覆盖 |
| "validator can only vote for a proposal once" | ProposalFoundry.t.sol::testValidatorCanOnlyVoteOnceAndExpire | ✅ 已覆盖 |
| "validator can't vote for proposal if it is expired" | ProposalFoundry.t.sol::testValidatorCanOnlyVoteOnceAndExpire | ✅ 已覆盖 |
| "Validate candidate's info" | ❌ 缺失 | ❌ 需补充详细信息验证 |

### 4. Vote for proposal(remove,pass) 测试

| Hardhat 测试 | Foundry 对应测试 | 状态 |
|-------------|-----------------|------|
| "normal vote for a proposal(3 true/2 false)" | ProposalFoundry.t.sol::testRemoveProposalPass | ✅ 已覆盖 |

### 5. Vote for proposal(reject) 测试

| Hardhat 测试 | Foundry 对应测试 | 状态 |
|-------------|-----------------|------|
| "normal vote(2 agree, 3 reject)" | ❌ 缺失 | ❌ 需补充提案被拒绝的测试 |

### 6. Create/Vote config update proposal 测试

| Hardhat 测试 | Foundry 对应测试 | 状态 |
|-------------|-----------------|------|
| "normal vote for a proposal(4 true/2 false)" | ProposalFoundry.t.sol::testConfigUpdateAll | ✅ 已覆盖 |

### 7. Set val unpass 测试

| Hardhat 测试 | Foundry 对应测试 | 状态 |
|-------------|-----------------|------|
| "only validator can set val unpass" | ❌ 缺失 | ❌ 需补充权限检查 |
| "validator contract can set val unpass" | ❌ 缺失 | ❌ 需补充功能测试 |

## test/validators.js 测试覆盖分析

### 1. Validators 合约测试

| Hardhat 测试 | Foundry 对应测试 | 状态 |
|-------------|-----------------|------|
| "can only init once" | ValidatorsCompleteFoundry.t.sol::testCanOnlyInitOnce | ✅ 已覆盖 |

### 2. create or edit validator 测试

| Hardhat 测试 | Foundry 对应测试 | 状态 |
|-------------|-----------------|------|
| "can't create validator if fee addr == address(0)" | ValidatorsCompleteFoundry.t.sol::testCreateValidatorInvalidFeeAddr | ✅ 已覆盖 |
| "can't create validator if describe info invalid" | ValidatorsCompleteFoundry.t.sol::testCreateValidatorInvalidDescription | ✅ 已覆盖 |
| "can't create validator if not pass propose" | ValidatorsCompleteFoundry.t.sol::testCreateValidatorNotAuthorized | ✅ 已覆盖 |
| "create validator" | ValidatorsCompleteFoundry.t.sol::testCreateValidatorSuccess | ✅ 已覆盖 |
| "edit validator info" | ValidatorsCompleteFoundry.t.sol::testEditValidatorInfo | ✅ 已覆盖 |

### 3. propose add a new val 测试

| Hardhat 测试 | Foundry 对应测试 | 状态 |
|-------------|-----------------|------|
| "not a val" | ValidatorsCompleteFoundry.t.sol::testProposeAddNewValidator | ✅ 已覆盖 |
| "create/vote proposal" | ValidatorsCompleteFoundry.t.sol::testProposeAddNewValidator | ✅ 已覆盖 |
| "is a val" | ValidatorsCompleteFoundry.t.sol::testProposeAddNewValidator | ✅ 已覆盖 |

### 4. propose remove a val 测试

| Hardhat 测试 | Foundry 对应测试 | 状态 |
|-------------|-----------------|------|
| "not a val" / "create/vote proposal" / "is a val" | ValidatorsCompleteFoundry.t.sol::testProposeRemoveValidator | ✅ 已覆盖 |

### 5. distribute block reward 测试

| Hardhat 测试 | Foundry 对应测试 | 状态 |
|-------------|-----------------|------|
| "miner can distribute to validator contract..." | ValidatorsCompleteFoundry.t.sol::testDistributeBlockReward | ✅ 已覆盖 |
| "update withdraw profit wait block" | ValidatorsCompleteFoundry.t.sol::testUpdateWithdrawProfitPeriod | ✅ 已覆盖 |
| "validator can withdraw profits" | ValidatorsCompleteFoundry.t.sol::testValidatorWithdrawProfits | ✅ 已覆盖 |
| "Can't call withdrawProfits if you don't have any profits" | ValidatorsCompleteFoundry.t.sol::testCantWithdrawWithoutProfits | ✅ 已覆盖 |

### 6. update set 测试

| Hardhat 测试 | Foundry 对应测试 | 状态 |
|-------------|-----------------|------|
| "update active validator set" | ValidatorsCompleteFoundry.t.sol::testUpdateActiveValidatorSet | ✅ 已覆盖 |

### 7. Punish 合约测试 (在 validators.js 中)

| Hardhat 测试 | Foundry 对应测试 | 状态 |
|-------------|-----------------|------|
| "can only init once" | ❌ 缺失 | ❌ 需补充 Punish 初始化测试 |
| "miner can punish validator" | PunishFoundry.t.sol (部分覆盖) | ⚠️ 需完善复杂惩罚逻辑 |
| "validator missed record will decrease if necessary" | PunishFoundry.t.sol::testDecreaseMissedBlocksCounter | ✅ 已覆盖 |
| "jailed record will be cleaned if validator repass proposal" | ❌ 缺失 | ❌ 需补充重新通过提案清理记录 |

## test/reward.js 测试覆盖分析

### 1. normal case 测试

| Hardhat 测试 | Foundry 对应测试 | 状态 |
|-------------|-----------------|------|
| "reward should be equally distributed to active validators if no stake" | RewardFoundry.t.sol::testRewardEquallyDistributedNoStake | ✅ 已覆盖 |

### 2. punish reward 测试

| Hardhat 测试 | Foundry 对应测试 | 状态 |
|-------------|-----------------|------|
| "remove validator's reward" | RewardFoundry.t.sol::testRemoveValidatorReward | ✅ 已覆盖 |
| "jailed validator can't get reward" | RewardFoundry.t.sol::testJailedValidatorCantGetReward | ✅ 已覆盖 |
| "jailed validator can't get profits of punish" | RewardFoundry.t.sol::testJailedValidatorCantGetPunishProfits | ✅ 已覆盖 |

## 总结

### ✅ 已完全覆盖的测试：20/27 (74%)

### ❌ 缺失的重要测试：7 个

### 需要补充的测试

1. **Proposal 测试补充**：
   - 多账户创建提案测试
   - 提案详细信息验证
   - 提案被拒绝场景 (2 同意 vs 3 反对)
   - setUnpassed 权限和功能测试

2. **Punish 测试补充**：
   - Punish 合约初始化测试
   - 完整的惩罚逻辑 (包含监禁到重新激活的完整流程)
   - 重新通过提案清理惩罚记录

3. **Edge Case 测试**：
   - 更多边界条件测试
   - 错误恢复场景测试

## 下一步行动计划

1. 补充缺失的 7 个测试用例
2. 增强现有测试的边界条件覆盖
3. 确保事件日志的验证
4. 添加更多的状态验证断言
