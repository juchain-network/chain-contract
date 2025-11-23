# POSA 系统合约与共识流程完整文档

## 目录

1. [系统概述](#系统概述)
2. [合约架构](#合约架构)
3. [合约主要功能](#合约主要功能)
4. [合约协作流程](#合约协作流程)
5. [共识流程与合约配合](#共识流程与合约配合)
6. [主要场景处理流程](#主要场景处理流程)
7. [关键机制说明](#关键机制说明)

---

## 一、系统概述

### 1.1 POSA 共识机制

**POSA (Proof of Stake Authorization)** 是结合了 PoS (Proof of Stake) 和 PoA (Proof of Authority) 的混合共识机制：

- **PoS 特性**：基于质押量选择验证者，支持委托和奖励分配
- **PoA 特性**：通过提案机制控制验证者准入，确保网络安全性
- **混合优势**：既保证了去中心化，又维护了网络的稳定性

### 1.2 核心组件

- **4 个核心合约**：Proposal、Staking、Validators、Punish
- **共识引擎**：Congress (Go 实现)
- **Epoch 机制**：每 86400 个区块（约 24 小时）更新一次验证者集合

---

## 二、合约架构

### 2.1 合约关系图

```
┌─────────────────────────────────────────────────────────────┐
│                     共识层 (Congress)                          │
│  - Prepare() → 写入 header.Extra                              │
│  - Finalize() → 调用合约进行奖励分配、验证者更新、惩罚          │
│  - snapshot.apply() → 构建验证者集合（过滤被 jail 的验证者）    │
└─────────────────────────────────────────────────────────────┘
                            │
                            │ 调用
                            ▼
┌─────────────────────────────────────────────────────────────┐
│                    Validators 合约 (0xf000)                  │
│  - 管理验证者集合 (currentValidatorSet, highestValidatorsSet) │
│  - 分配交易手续费奖励                                         │
│  - 提供验证者查询接口                                         │
└─────────────────────────────────────────────────────────────┘
         │                    │                    │
         │                    │                    │
    ┌────▼────┐         ┌────▼────┐         ┌────▼────┐
    │ Proposal│         │ Staking │         │ Punish  │
    │ (0xf002)│         │ (0xf003)│         │ (0xf001)│
    └─────────┘         └─────────┘         └─────────┘
```

### 2.2 合约地址

| 合约名称 | 地址 | 功能描述 |
|---------|------|----------|
| Validators | `0x000000000000000000000000000000000000f000` | 验证者管理、奖励分配 |
| Punish | `0x000000000000000000000000000000000000f001` | 惩罚机制 |
| Proposal | `0x000000000000000000000000000000000000f002` | 提案治理 |
| Staking | `0x000000000000000000000000000000000000f003` | 质押管理 |

---

## 三、合约主要功能

### 3.1 Proposal 合约

**核心职责**：管理验证者提案和系统配置提案

#### 主要功能

1. **验证者提案管理**
   - `createProposal(address dst, bool flag, string details)`: 创建验证者添加/移除提案
   - `voteProposal(bytes32 id, bool auth)`: 验证者投票（仅活动验证者可投票）
   - 提案通过条件：`同意票数 >= 活动验证者数量 / 2 + 1`
   - 提案通过后设置 `pass[address] = true/false`

2. **系统配置提案管理**
   - `createUpdateConfigProposal(uint256 cid, uint256 newValue)`: 创建配置修改提案
   - 可修改的参数：
     - `cid = 0`: 提案有效期（1小时 - 30天）
     - `cid = 1`: 惩罚阈值（必须 > 0）
     - `cid = 2`: 移除阈值（必须 > 0）
     - `cid = 3`: 减少率（必须 > 0，防止除零）
     - `cid = 4`: 提取收益周期（必须 > 0）
   - **注意**：cid 5 和 6（增发相关）已移除，系统不再支持代币增发

3. **提案状态管理**
   - 维护 `pass[address]` 映射，记录地址是否通过提案
   - 维护 `proposalPassedTime[address]`，记录提案通过时间（用于 7 天注册期限）
   - `isProposalValidForStaking(address)`: 检查提案是否在7天有效期内（仅用于新注册验证者）
   - **重要**：7天有效期仅用于新注册，已注册验证者不受此限制

4. **分级惩罚机制（方案C）**
   - `violationCount[address]`: 记录验证者违规次数
   - `setUnpassed(address)`: 验证者被移除时，增加违规计数
   - `autoRestorePass(address)`: 3 次以下违规时自动恢复 pass 状态
   - `getViolationCount(address)`: 查询验证者违规次数
   - 投票通过后自动重置违规计数

#### 关键数据结构

```solidity
mapping(address => bool) public pass;  // 地址是否通过提案
mapping(address => uint256) public proposalPassedTime;  // 提案通过时间
mapping(bytes32 => ProposalInfo) public proposals;  // 提案信息
mapping(bytes32 => ResultInfo) public results;  // 投票结果
mapping(address => uint256) public violationCount;  // 违规次数（方案C: 分级惩罚）
```

#### 与其他合约的关系

- **依赖 Validators**: 获取活动验证者列表用于投票统计
- **被 Staking 依赖**: Staking 检查 `pass[address]` 和 `isProposalValidForStaking()` 判断是否允许新注册质押（仅用于新注册，已注册验证者不受7天有效期限制）
- **被 Validators 调用**: `setUnpassed()` 在验证者被移除时清除提案状态

---

### 3.2 Staking 合约

**核心职责**：管理验证者质押、委托和奖励分配

#### 主要功能

1. **验证者质押管理**
   - `registerValidator(uint256 commissionRate)`: 验证者注册（需先通过提案，等待 7 天）
   - `addValidatorStake()`: 增加质押金额
   - `withdrawValidatorStake(uint256 amount)`: 提取质押（部分提取，剩余质押必须 >= MIN_VALIDATOR_STAKE）
   - `emergencyExit()`: 紧急退出（提取所有质押，检查退出后剩余验证者数量，如果验证者在 currentValidatorSet 中会先 jail）
   - `updateCommissionRate(uint256 newCommissionRate)`: 更新佣金率

2. **委托管理**
   - `delegate(address validator)`: 用户委托代币给验证者
   - `undelegate(address validator, uint256 amount)`: 取消委托（进入解绑期）
   - `withdrawUnbonded(address validator, uint256 maxEntries)`: 提取解绑的代币（7 天解绑期）

3. **奖励分配**
   - `distributeRewards(address validator)`: 出块奖励分配（给验证者和委托者）
   - `claimRewards(address validator)`: 提取奖励（验证者提取佣金，委托者提取委托奖励）

4. **验证者状态管理**
   - `jailValidator(address validator, uint256 jailBlocks)`: 监禁验证者（仅 Punish 合约可调用）
   - `unjailValidator(address validator)`: 解除监禁（验证者自己调用）
   - `getTopValidators()`: 获取顶级验证者列表（用于 epoch 更新，已过滤被 jail 的验证者）
   - `_removeFromAllValidators(address validator)`: 从 allValidators 数组中移除验证者（私有函数，仅在 emergencyExit 时调用）

#### 关键数据结构

```solidity
struct ValidatorStake {
    uint256 selfStake;          // 验证者自身质押
    uint256 totalDelegated;     // 总委托金额
    uint256 commissionRate;     // 佣金率 (0-10000, 代表 0%-100%)
    uint256 accumulatedRewards; // 累积奖励
    bool isJailed;              // 是否被监禁
    uint256 jailUntilBlock;     // 监禁到期区块号
}

struct Delegation {
    uint256 amount;             // 委托金额
    uint256 rewardDebt;          // 奖励债务（用于精确计算奖励）
    uint256 unbondingAmount;     // 解绑中的金额
    uint256 unbondingBlock;      // 解绑完成区块号
}
```

#### 与其他合约的关系

- **依赖 Proposal**: 检查 `pass[address]` 和 `isProposalValidForStaking()` 判断是否允许新注册质押（仅用于新注册，已注册验证者不受7天有效期限制）
- **依赖 Validators**: 调用 `tryAddValidatorToHighestSet()` 添加验证者到最高集合（仅在注册时）
- **被 Punish 调用**: `jailValidator()` 监禁验证者
- **被共识层调用**: `getTopValidators()` 获取验证者列表

---

### 3.3 Validators 合约

**核心职责**：管理验证者集合和交易手续费奖励分配

#### 主要功能

1. **验证者集合管理**
   - `updateActiveValidatorSet(address[] memory newSet, uint256 epoch)`: 更新验证者集合（由共识层调用，传入验证者列表）
   - `getTopValidators()`: 获取顶级验证者列表（返回 `highestValidatorsSet`，用于 POA 模式）
   - **注意**：POSA 模式下，共识层直接调用 `Staking.getTopValidators()` 获取基于质押的验证者列表
   - `getActiveValidators()`: 获取活动验证者列表（基于 `currentValidatorSet`）
   - `getActiveValidatorCount()`: 获取活动验证者数量（基于 `currentValidatorSet`，高效方法）

2. **验证者信息管理**
   - `addValidator(address validator, address feeAddr)`: 添加验证者
   - `removeValidator(address val)`: 移除验证者（仅 Punish 合约可调用）
   - `tryRemoveValidator(address val)`: 尝试移除验证者（仅 Proposal 合约可调用）
   - `editValidator(address feeAddr, ...)`: 编辑验证者信息

3. **奖励分配**
   - `distributeBlockReward()`: 分配交易手续费奖励给所有活动验证者（排除被 jail 的验证者）

4. **验证者状态查询**
   - `isValidatorJailed(address validator)`: 查询验证者是否被监禁（代理到 Staking）
   - `isValidatorActive(address validator)`: 查询验证者是否活跃（代理到 Staking）
   - `isValidatorExist(address validator)`: 查询验证者是否存在（是否有质押）

#### 关键数据结构

```solidity
struct Validator {
    address payable feeAddr;           // 费用接收地址
    Description description;            // 验证者描述信息
    uint256 aacIncoming;               // 累积的交易手续费收入
    uint256 totalJailedHb;            // 总被监禁收入
    uint256 lastWithdrawProfitsBlock;  // 最后提取收益的区块号
    // 注意: status 字段已移除，状态由 Staking 合约管理
    // Status 通过 getValidatorInfo() 动态计算（向后兼容）
}

address[] public currentValidatorSet;  // 当前验证者集合（仅在 epoch 更新）
address[] public highestValidatorsSet; // 最高验证者集合（通过 tryAddValidatorToHighestSet 等方法管理，不在 updateActiveValidatorSet 中同步）
mapping(address => Validator) validatorInfo;  // 验证者信息
```

#### 与其他合约的关系

- **依赖 Staking**: 查询验证者状态（jail、活跃状态）、获取顶级验证者列表
- **依赖 Proposal**: 调用 `setUnpassed()` 清除提案状态
- **被 Punish 调用**: `removeValidator()` 移除验证者
- **被共识层调用**: 更新验证者集合、分配奖励

---

### 3.4 Punish 合约

**核心职责**：惩罚错过出块的验证者

#### 主要功能

1. **惩罚机制**
   - `punish(address val)`: 惩罚验证者（仅矿工可调用）
   - 记录 `missedBlocksCounter`
   - 达到 `punishThreshold` (24块) 时：移除验证者收入
   - 达到 `removeThreshold` (48块) 时：监禁并移除验证者

2. **惩罚记录管理**
   - `decreaseMissedBlocksCounter(uint256 epoch)`: 在 epoch 时减少惩罚记录
   - `cleanPunishRecord(address val)`: 清理验证者的惩罚记录（重新质押时）

#### 关键数据结构

```solidity
struct PunishRecord {
    uint256 missedBlocksCounter;  // 错过块数计数器
    uint256 index;                // 在数组中的索引
    bool exist;                    // 是否存在
}
```

#### 与其他合约的关系

- **依赖 Validators**: 调用 `removeValidator()` 和 `removeValidatorIncoming()` 移除验证者
- **依赖 Staking**: 调用 `jailValidator()` 监禁验证者
- **依赖 Proposal**: 获取 `removeThreshold()` 和 `punishThreshold()`
- **被共识层调用**: `punish()` 惩罚验证者

---

## 四、合约协作流程

### 4.1 验证者注册协作流程

```
1. Proposal.createProposal(dst, true, details)
   └─> 创建提案，设置 proposalType = 1

2. Proposal.voteProposal(id, true) [多个验证者投票]
   └─> 记录投票: results[id].agree++ 或 results[id].reject++
   └─> 计算当前活跃验证者的投票数: getActiveVoteCount()
   └─> 当 当前活跃验证者同意票数 >= getActiveValidatorCount() / 2 + 1 时
       └─> pass[dst] = true
       └─> proposalPassedTime[dst] = block.timestamp
   └─> **注意**：只计算当前仍然活跃的验证者的投票

3. 等待 7 天
   └─> Proposal.isProposalValidForStaking(dst) 检查（仅用于新注册）

4. Staking.registerValidator(commissionRate) {value: >= 10000 ether}
   └─> 检查: proposalContract.pass(msg.sender) == true
   └─> 检查: proposalContract.isProposalValidForStaking(msg.sender) == true（7天内必须注册）
   └─> 创建 ValidatorStake 记录
   └─> validatorsContract.tryAddValidatorToHighestSet(msg.sender)

5. 等待下一个 Epoch
   └─> 共识层调用: staking.getTopValidators() 获取验证者列表
   └─> 共识层调用: Validators.updateActiveValidatorSet(newSet, epoch)
       └─> 更新 currentValidatorSet
       └─> 同步更新 highestValidatorsSet
```

### 4.2 验证者惩罚协作流程

```
1. 共识层检测到验证者错过块
   └─> Congress.punishOutOfTurnValidator()
       └─> Punish.punish(validator)

2. Punish.punish(validator)
   └─> missedBlocksCounter++

3. 达到 punishThreshold (24块)
   └─> missedBlocksCounter % 24 == 0
   └─> Validators.removeValidatorIncoming(validator)
       └─> 移除验证者收入

4. 达到 removeThreshold (48块)
   └─> missedBlocksCounter % 48 == 0
   └─> Staking.jailValidator(validator, 86400)  [先执行]
       └─> isJailed = true
       └─> jailUntilBlock = block.number + 86400
   └─> Validators.removeValidator(validator)  [后执行]
       └─> Proposal.setUnpassed(validator)
           └─> pass[validator] = false
       └─> 从 highestValidatorsSet 移除（如果 length > 1）

5. 下一个 Epoch
   └─> Staking.getTopValidators() 过滤被 jail 的验证者
   └─> 验证者不再出现在验证者集合中
```

### 4.3 奖励分配协作流程

```
1. 共识层 Finalize()
   └─> 如果有交易: distributeFeeReward()
       └─> Validators.distributeBlockReward() {value: 交易手续费}
           └─> 分配给所有活动验证者（排除被 jail 的）
           └─> 每个验证者获得: totalReward / activeValidatorCount

   └─> distributeCoinbaseReward()
       └─> Staking.distributeRewards(validator) {value: 基础出块奖励}
           └─> 验证者获得: commission + (剩余奖励 * selfStake / totalStake)
           └─> 委托者获得: (剩余奖励 * delegatedAmount / totalStake)
           └─> 更新 rewardPerShare
```

### 4.4 Epoch 更新协作流程

```
1. 共识层 Prepare() [Epoch 块]
   └─> getTopValidators() [使用 parent state]
       └─> Staking.getTopValidators()
           └─> 返回已过滤的验证者列表（排除被 jail 的）
   └─> 写入 header.Extra

2. 共识层 Finalize() [Epoch 块]
   └─> handleEpochTransition()
       └─> updateValidators()
           ├─> Staking.getTopValidators() [使用 current state]
           │   └─> 返回已过滤的验证者列表（排除被 jail 的）
           └─> Validators.updateActiveValidatorSet(newSet, epoch)
               └─> 更新 currentValidatorSet 和 highestValidatorsSet

3. 共识层 snapshot.apply() [验证历史区块]
   └─> getTopValidatorsFunc() [使用 parent state]
       └─> 获取已过滤的验证者列表
       └─> 如果失败且是状态不可用错误，fallback 到 header.Extra
       └─> 更新 snap.Validators（不包含被 jail 的验证者）
```

---

## 五、共识流程与合约配合

### 5.1 区块生命周期

#### Prepare 阶段

```go
func (c *Congress) Prepare(chain consensus.ChainHeaderReader, header *types.Header) error {
    // 1. 设置 coinbase 和 nonce
    header.Coinbase = c.validator
    header.Nonce = types.BlockNonce{}
    
    // 2. 获取 snapshot
    snap, err := c.snapshot(chain, number-1, header.ParentHash, nil)
    
    // 3. 如果是 Epoch 块，获取验证者列表并写入 header.Extra
    if number % c.config.Epoch == 0 {
        newSortedValidators, err := c.getTopValidators(chain, header)  // 使用 parent state
        // 写入 header.Extra
        for _, validator := range newSortedValidators {
            header.Extra = append(header.Extra, validator.Bytes()...)
        }
    }
}
```

**关键点**：
- 使用 **parent state** 获取验证者列表
- 写入 `header.Extra`，供后续验证使用

#### Finalize 阶段

```go
func (c *Congress) Finalize(...) error {
    // 1. 初始化系统合约（区块 1）
    if header.Number == 1 {
        initializeSystemContracts(...)
    }
    
    // 2. 惩罚错过块的验证者
    if header.Difficulty != diffInTurn {
        punishOutOfTurnValidator(...)
            └─> Punish.punish(validator)
    }
    
    // 3. 分配交易手续费奖励
    if len(txs) > 0 {
        distributeFeeReward(...)
            └─> Validators.distributeBlockReward() {value: 交易手续费}
    }
    
    // 4. 分配出块基础奖励（POSA 模式）
    distributeCoinbaseReward(...)
        └─> Staking.distributeRewards(validator) {value: 基础出块奖励}
    
    // 5. Epoch 更新（Epoch 块）
    if header.Number % c.config.Epoch == 0 {
        handleEpochTransition(...)
            ├─> Staking.getTopValidators()  // 使用 current state
            └─> Validators.updateActiveValidatorSet(newSet, epoch)
    }
}
```

**关键点**：
- 使用 **current state** 更新验证者集合
- 确保被 jail 的验证者立即被排除

#### VerifyHeader 阶段

```go
func (c *Congress) VerifyHeader(...) error {
    // 1. 验证区块头基本信息
    
    // 2. 如果是 Epoch 块，验证验证者集合
    if header.Number % c.config.Epoch == 0 {
        // POSA 模式：允许 newValidators 是 header.Extra 的子集
        // （因为被 jail 的验证者会被排除）
        // POA 模式：必须完全匹配
    }
}
```

### 5.2 Snapshot 机制

**Snapshot** 存储验证者集合，用于验证区块签名：

```go
type Snapshot struct {
    Validators map[common.Address]struct{}  // 验证者集合
    Recents    map[uint64]common.Address    // 最近签名的验证者
    Number     uint64                       // 区块号
    Hash       common.Hash                  // 区块哈希
}
```

**关键机制**：
- 在 Epoch 块更新 `snap.Validators`
- 使用 `getTopValidatorsFunc()` 获取已过滤的验证者列表（排除被 jail 的）
- 如果状态不可用，fallback 到 `header.Extra`（但会记录警告）

### 5.3 验证者选择机制

**POSA 模式**：
1. `Staking.getTopValidators()` 筛选条件：
   - `selfStake >= MIN_VALIDATOR_STAKE` (10000 ether)
   - `!isJailed` (验证者未被监禁)
   - `proposalContract.pass(validator) == true` (已通过提案)
   - **注意**：不检查 `isProposalValidForStaking()`，7天有效期仅用于新注册验证者（在 `registerValidator()` 中检查）
2. 按总质押排序：`selfStake + totalDelegated`
3. 返回前 `MAX_VALIDATORS` (21) 个验证者

**POA 模式**：
- 使用 `highestValidatorsSet`（缓存）

---

## 六、主要场景处理流程

### 6.1 新验证者加入完整流程

```
┌─────────────────────────────────────────────────────────────┐
│ 阶段 1: 创建提案                                              │
└─────────────────────────────────────────────────────────────┘
用户调用: Proposal.createProposal(dst, true, details)
  ├─> 检查: !pass[dst] (不能重复添加)
  ├─> 生成提案 ID: keccak256(proposer, dst, flag, details, timestamp)
  └─> 创建 ProposalInfo 记录
      └─> proposalType = 1 (验证者提案)
      └─> flag = true (添加验证者)

┌─────────────────────────────────────────────────────────────┐
│ 阶段 2: 验证者投票                                             │
└─────────────────────────────────────────────────────────────┘
多个活动验证者调用: Proposal.voteProposal(id, true)
  ├─> 检查: 调用者是活动验证者 (onlyValidator)
  ├─> 检查: 未投票过
  ├─> 检查: 提案未过期
  ├─> 记录投票: votes[voter][id] = VoteInfo
  ├─> 更新统计: results[id].agree++ (永久记录，用于历史查询)
  └─> 检查是否通过:
      ├─> 计算当前活跃验证者的投票数: getActiveVoteCount(id, true)
      └─> 如果 当前活跃验证者同意票数 >= getActiveValidatorCount() / 2 + 1
          ├─> pass[dst] = true
          ├─> proposalPassedTime[dst] = block.timestamp
          └─> emit LogPassProposal(id)
      └─> **注意**：只计算当前仍然活跃的验证者的投票（被踢出的验证者投票不计入）

┌─────────────────────────────────────────────────────────────┐
│ 阶段 3: 等待 7 天                                              │
└─────────────────────────────────────────────────────────────┘
等待: block.timestamp >= proposalPassedTime[dst] + 7 days
  └─> Proposal.isProposalValidForStaking(dst) 返回 true
  └─> **注意**：7天有效期仅用于新注册，已注册验证者不受此限制

┌─────────────────────────────────────────────────────────────┐
│ 阶段 4: 验证者质押注册                                          │
└─────────────────────────────────────────────────────────────┘
验证者调用: Staking.registerValidator(commissionRate) {value: >= 10000 ether}
  ├─> 检查: msg.value >= MIN_VALIDATOR_STAKE (10000 ether)
  ├─> 检查: commissionRate <= COMMISSION_RATE_BASE (10000)
  ├─> 检查: 未注册过 (selfStake == 0)
  ├─> 检查: proposalContract.pass(msg.sender) == true
  ├─> 检查: proposalContract.isProposalValidForStaking(msg.sender) == true（7天内必须注册）
  ├─> 检查: !isJailed || block.number >= jailUntilBlock
  ├─> 创建 ValidatorStake 记录
  ├─> 添加到 allValidators 数组
  ├─> 更新 totalStaked
  ├─> validatorsContract.tryAddValidatorToHighestSet(msg.sender)
  └─> emit ValidatorRegistered(...)

┌─────────────────────────────────────────────────────────────┐
│ 阶段 5: 等待 Epoch 更新                                         │
└─────────────────────────────────────────────────────────────┘
等待: block.number % 86400 == 0 (下一个 Epoch 块)
  └─> 共识层调用: Staking.getTopValidators() 获取验证者列表
      ├─> 筛选符合条件的验证者
      ├─> 按总质押排序
      └─> 返回前 21 名
   └─> 共识层调用: Validators.updateActiveValidatorSet(newSet, epoch)
      ├─> 更新 currentValidatorSet
      └─> emit LogUpdateValidator(...)
      └─> **注意**：`highestValidatorsSet` 通过其他方法管理（如 `tryAddValidatorToHighestSet`），不会在此处同步更新

┌─────────────────────────────────────────────────────────────┐
│ 阶段 6: 开始出块                                                │
└─────────────────────────────────────────────────────────────┘
验证者出现在验证者集合中
  └─> 可以开始出块
  └─> 可以接收奖励
```

### 6.2 验证者投票流程

```
┌─────────────────────────────────────────────────────────────┐
│ 步骤 1: 验证者投票                                             │
└─────────────────────────────────────────────────────────────┘
活动验证者调用: Proposal.voteProposal(id, auth)
  ├─> 检查: 调用者是活动验证者
  │   └─> Validators.isActiveValidator(msg.sender)
  │       └─> 检查是否在 currentValidatorSet 中
  ├─> 检查: 未投票过
  ├─> 检查: 提案未过期
  ├─> 记录投票
  └─> 更新统计:
      ├─> 如果 auth == true: results[id].agree++
      └─> 如果 auth == false: results[id].reject++

┌─────────────────────────────────────────────────────────────┐
│ 步骤 2: 检查提案是否通过                                         │
└─────────────────────────────────────────────────────────────┘
每次投票后检查:
  ├─> 计算当前活跃验证者的投票数:
  │   ├─> getActiveVoteCount(id, true) - 只统计当前活跃验证者的同意票
  │   └─> getActiveVoteCount(id, false) - 只统计当前活跃验证者的拒绝票
  │
  ├─> 如果 活跃验证者同意票数 >= getActiveValidatorCount() / 2 + 1
  │   ├─> 提案通过
  │   ├─> 如果是验证者提案 (proposalType == 1):
  │   │   ├─> 如果 flag == true: pass[dst] = true
  │   │   └─> 如果 flag == false: pass[dst] = false, removeValidator(dst)
  │   └─> 如果是配置提案 (proposalType == 2): updateConfig(...)
  │
  └─> 如果 活跃验证者拒绝票数 >= getActiveValidatorCount() / 2 + 1
      └─> 提案被拒绝

关键点:
- 投票阈值基于当前活跃验证者数量（已过滤被 jail 的验证者）
- 只计算当前仍然活跃的验证者的投票（被踢出的验证者投票不计入）
- 每个验证者只能投票一次
- 提案有效期: proposalLastingPeriod (默认 7 天)
- **重要**：如果验证者投票后被踢出，其投票不再计入阈值检查
```

### 6.3 验证者质押注册流程

```
┌─────────────────────────────────────────────────────────────┐
│ 前置条件检查                                                   │
└─────────────────────────────────────────────────────────────┘
1. 提案已通过: proposalContract.pass(msg.sender) == true
2. 提案通过后7天内必须注册: proposalContract.isProposalValidForStaking(msg.sender) == true（仅用于新注册）
3. 未注册过: validatorStakes[msg.sender].selfStake == 0
4. 未被监禁或监禁期已过: !isJailed || block.number >= jailUntilBlock

┌─────────────────────────────────────────────────────────────┐
│ 质押注册                                                       │
└─────────────────────────────────────────────────────────────┘
验证者调用: Staking.registerValidator(commissionRate) {value: >= 10000 ether}
  ├─> 检查: msg.value >= MIN_VALIDATOR_STAKE (10000 ether)
  ├─> 检查: commissionRate <= COMMISSION_RATE_BASE (10000)
  ├─> 检查: proposalContract.pass(msg.sender) == true
  ├─> 检查: proposalContract.isProposalValidForStaking(msg.sender) == true（7天内必须注册）
  ├─> 创建 ValidatorStake 记录:
  │   ├─> selfStake = msg.value
  │   ├─> totalDelegated = 0
  │   ├─> commissionRate = commissionRate
  │   ├─> accumulatedRewards = 0
  │   ├─> isJailed = false
  │   └─> jailUntilBlock = 0
  ├─> 添加到 allValidators 数组
  ├─> 更新 totalStaked
  ├─> validatorsContract.tryAddValidatorToHighestSet(msg.sender)
  └─> emit ValidatorRegistered(...)

┌─────────────────────────────────────────────────────────────┐
│ 后续操作                                                       │
└─────────────────────────────────────────────────────────────┘
- 可以增加质押: addValidatorStake()
- 可以更新佣金率: updateCommissionRate()
- 等待下一个 Epoch 进入验证者集合
```

### 6.4 增加质押流程

```
验证者调用: Staking.addValidatorStake() {value: > 0}
  ├─> 检查: 验证者已注册 (selfStake >= MIN_VALIDATOR_STAKE)
  ├─> 检查: 未被监禁或监禁期已过
  ├─> 更新: selfStake += msg.value
  ├─> 更新: totalStaked += msg.value
  └─> emit ValidatorStakeWithdrawn(...)
  └─> **注意**：验证者集合在下一个 Epoch 更新时自动同步
```

### 6.5 提取质押流程

```
┌─────────────────────────────────────────────────────────────┐
│ 部分提取质押                                                   │
└─────────────────────────────────────────────────────────────┘
验证者调用: Staking.withdrawValidatorStake(amount)
  ├─> 检查: amount > 0
  ├─> 检查: selfStake >= amount
  ├─> 计算: remainingStake = selfStake - amount
  ├─> 检查: remainingStake >= MIN_VALIDATOR_STAKE
  │   └─> 如果不满足，交易失败，提示使用 emergencyExit() 完全退出
  ├─> 更新: selfStake = remainingStake
  ├─> 更新: totalStaked -= amount
  ├─> emit ValidatorStakeWithdrawn(...)
  └─> 转账: payable(msg.sender).call{value: amount}("")
      └─> 使用 call() 并检查返回值，确保转账成功

┌─────────────────────────────────────────────────────────────┐
│ 紧急退出（全部提取）                                            │
└─────────────────────────────────────────────────────────────┘
验证者调用: Staking.emergencyExit()
  ├─> 检查: 验证者已注册 (onlyValidValidator)
  ├─> 检查验证者是否在 currentValidatorSet 中
  │   └─> isInCurrentSet = validatorsContract.isActiveValidator(msg.sender)
  ├─> 计算退出后剩余验证者数量
  │   └─> remainingCount = (isInCurrentSet && !isJailed) ? currentActiveCount - 1 : currentActiveCount
  ├─> 检查: remainingCount >= MIN_VALIDATORS (3)
  │   └─> 如果不满足，交易失败
  ├─> 如果验证者在 currentValidatorSet 中且未被 jailed:
  │   ├─> 先 jail 验证者（1 个 epoch，86400 块）
  │   │   ├─> isJailed = true
  │   │   ├─> jailUntilBlock = block.number + 86400
  │   │   └─> emit ValidatorJailed(...)
  │   └─> 确保验证者立即停止出块，平滑退出
  ├─> 更新: selfStake = 0
  ├─> 更新: totalStaked -= withdrawAmount (withdrawAmount 是退出前的 selfStake 值)
  ├─> 从 allValidators 数组中移除验证者
  │   └─> _removeFromAllValidators(msg.sender)
  │       └─> 使用 swap-and-pop 技术安全移除
  ├─> emit ValidatorExited(...)
  └─> 转账: payable(msg.sender).call{value: selfStake}("")
      └─> 使用 call() 并检查返回值，确保转账成功
```

### 6.6 委托流程

```
┌─────────────────────────────────────────────────────────────┐
│ 委托代币给验证者                                                │
└─────────────────────────────────────────────────────────────┘
用户调用: Staking.delegate(validator) {value: >= 1 ether}
  ├─> 检查: validator 是活动验证者 (onlyActiveValidator)
  ├─> 检查: msg.value >= MIN_DELEGATION (1 ether)
  ├─> 检查: validator != msg.sender
  ├─> 更新奖励: _updateRewards(msg.sender, validator)
  ├─> 更新委托记录:
  │   ├─> delegations[delegator][validator].amount += msg.value
  │   └─> delegations[delegator][validator].rewardDebt = amount * rewardPerShare / 1e18
  ├─> 更新验证者总委托: totalDelegated += msg.value
  ├─> 更新总质押: totalStaked += msg.value
  └─> emit Delegated(...)
  └─> **注意**：验证者集合在下一个 Epoch 更新时自动同步

┌─────────────────────────────────────────────────────────────┐
│ 取消委托                                                       │
└─────────────────────────────────────────────────────────────┘
用户调用: Staking.undelegate(validator, amount)
  ├─> 检查: validator 是有效验证者
  ├─> 检查: amount > 0
  ├─> 检查: delegations[msg.sender][validator].amount >= amount
  ├─> 更新奖励: _updateRewards(msg.sender, validator)
  ├─> 更新委托记录:
  │   ├─> delegations[msg.sender][validator].amount -= amount
  │   └─> delegations[msg.sender][validator].rewardDebt = newAmount * rewardPerShare / 1e18
  ├─> 更新验证者总委托: totalDelegated -= amount
  ├─> 更新总质押: totalStaked -= amount
  ├─> 创建解绑记录:
  │   └─> unbondingDelegations[msg.sender][validator].push({
  │       amount: amount,
  │       completionBlock: block.number + proposalContract.unbondingPeriod() (默认7天，可配置)
  │   })
  └─> emit Undelegated(...)
  └─> **注意**：验证者集合在下一个 Epoch 更新时自动同步

┌─────────────────────────────────────────────────────────────┐
│ 提取解绑代币                                                    │
└─────────────────────────────────────────────────────────────┘
用户调用: Staking.withdrawUnbonded(validator, maxEntries)
  ├─> 遍历解绑记录: unbondingDelegations[msg.sender][validator]
  ├─> 找到已到期的记录 (completionBlock <= block.number)
  ├─> 累计可提取金额
  ├─> 移除已完成的记录
  ├─> 检查: totalWithdraw > 0
  └─> 转账: payable(msg.sender).transfer(totalWithdraw)
```

### 6.7 奖励分配和提取流程

```
┌─────────────────────────────────────────────────────────────┐
│ 交易手续费奖励分配（每块）                                       │
└─────────────────────────────────────────────────────────────┘
共识层 Finalize() → distributeFeeReward()
  └─> Validators.distributeBlockReward() {value: 交易手续费}
      ├─> 检查: 调用者是矿工 (onlyMiner)
      ├─> 检查: 验证者已注册 (isValidatorExist)
      ├─> 计算活动验证者数量（排除被 jail 的）: getRewardLen(address(0))
      ├─> 分配: 每个活动验证者获得 totalReward / activeCount
      └─> 记录到 validatorInfo[validator].aacIncoming

┌─────────────────────────────────────────────────────────────┐
│ 出块基础奖励分配（每块，POSA 模式）                              │
└─────────────────────────────────────────────────────────────┘
共识层 Finalize() → distributeCoinbaseReward()
  └─> Staking.distributeRewards(validator) {value: 基础出块奖励}
      ├─> 检查: 调用者是矿工 (onlyMiner)
      ├─> 检查: validator 是活动验证者 (onlyActiveValidator)
      ├─> 计算总质押: totalStake = selfStake + totalDelegated
      ├─> 分配佣金给验证者:
      │   └─> commission = msg.value * commissionRate / 10000
      │   └─> accumulatedRewards += commission
      ├─> 计算剩余奖励: remainingRewards = msg.value - commission
      ├─> 验证者份额: validatorShare = remainingRewards * selfStake / totalStake
      │   └─> accumulatedRewards += validatorShare
      ├─> 委托者份额: delegatorRewards = remainingRewards - validatorShare
      └─> 更新 rewardPerShare: rewardPerShare += delegatorRewards * 1e18 / totalDelegated

┌─────────────────────────────────────────────────────────────┐
│ 提取奖励                                                       │
└─────────────────────────────────────────────────────────────┘
验证者调用: Staking.claimRewards(validator)
  ├─> 更新奖励: _updateRewards(validator, validator)
  ├─> 提取佣金: accumulatedRewards
  └─> 转账: payable(validator).transfer(commission)

委托者调用: Staking.claimRewards(validator)
  ├─> 更新奖励: _updateRewards(delegator, validator)
  ├─> 计算待提取奖励: pending = amount * rewardPerShare / 1e18 - rewardDebt
  └─> 转账: payable(delegator).transfer(pending)

验证者提取交易手续费: Validators.withdrawProfits(validator)
  ├─> 检查: aacIncoming > 0
  ├─> 检查: block.number >= lastWithdrawProfitsBlock + withdrawProfitPeriod
  ├─> 转账到 feeAddr: payable(feeAddr).transfer(aacIncoming)
  └─> 清零: aacIncoming = 0
```

### 6.8 验证者被惩罚流程

```
┌─────────────────────────────────────────────────────────────┐
│ 检测错过出块                                                    │
└─────────────────────────────────────────────────────────────┘
共识层 Finalize() → punishOutOfTurnValidator()
  ├─> 获取当前应该出块的验证者: validators[number % len(validators)]
  ├─> 检查: 该验证者是否最近签过名
  └─> 如果未签名: punishValidator(validator)
      └─> Punish.punish(validator)

┌─────────────────────────────────────────────────────────────┐
│ 惩罚处理                                                       │
└─────────────────────────────────────────────────────────────┘
Punish.punish(validator)
  ├─> 检查: 调用者是矿工 (onlyMiner)
  ├─> 更新: missedBlocksCounter++
  │
  ├─> 如果 missedBlocksCounter % removeThreshold (48) == 0:
  │   ├─> missedBlocksCounter = 0
  │   ├─> Staking.jailValidator(validator, 86400)  [先执行]
  │   │   └─> isJailed = true
  │   │   └─> jailUntilBlock = block.number + 86400
  │   └─> Validators.removeValidator(validator)  [后执行]
  │       ├─> Proposal.setUnpassed(validator)
  │       │   └─> pass[validator] = false
  │       └─> 从 highestValidatorsSet 移除（如果 length > 1）
  │
  └─> 如果 missedBlocksCounter % punishThreshold (24) == 0:
      └─> Validators.removeValidatorIncoming(validator)
          └─> 移除验证者收入 (aacIncoming = 0)

┌─────────────────────────────────────────────────────────────┐
│ 立即生效（Epoch 块）                                            │
└─────────────────────────────────────────────────────────────┘
如果当前块是 Epoch 块:
  └─> 共识层调用: Staking.getTopValidators() 获取验证者列表（已过滤被 jail 的）
  └─> 共识层调用: Validators.updateActiveValidatorSet(newSet, epoch)
      └─> 验证者立即从验证者集合中排除

如果当前块不是 Epoch 块:
  └─> 验证者仍在当前验证者集合中
  └─> 但 snapshot.apply() 会过滤被 jail 的验证者
  └─> 验证者无法出块（不在 snap.Validators 中）
```

### 6.9 验证者重新加入流程（方案C: 分级惩罚机制）

**方案C核心设计**：
- **3 次以下违规**：可以 unjail，自动恢复 `pass` 状态，无需重新投票
- **4 次及以上违规**：不能 unjail，必须先重新提案并通过投票
- **改过机制**：投票通过后重置违规计数

```
┌─────────────────────────────────────────────────────────────┐
│ 等待监禁期结束                                                  │
└─────────────────────────────────────────────────────────────┘
等待: block.number >= jailUntilBlock
  └─> 验证者可以调用 unjailValidator()

┌─────────────────────────────────────────────────────────────┐
│ 解除监禁（方案C: 分级惩罚）                                      │
└─────────────────────────────────────────────────────────────┘
验证者调用: Staking.unjailValidator(validator)
  ├─> 检查: 调用者是验证者自己
  ├─> 检查: 验证者被监禁
  ├─> 检查: 监禁期已过
  ├─> 方案C: 检查违规次数 (violationCount[validator])
  │   └─> 如果 violationCount > 3: 直接 revert，不允许 unjail
  │
  ├─> 更新: isJailed = false
  ├─> 更新: jailUntilBlock = 0
  │
  ├─> 方案C: 自动恢复 pass 状态（3 次以下违规）
  │   └─> Proposal.autoRestorePass(validator)
  │       ├─> 如果 violationCount <= 3:
  │       │   ├─> pass[validator] = true (自动恢复)
  │       │   └─> proposalPassedTime[validator] = block.timestamp
  │       └─> 如果 violationCount > 3: 返回 false（不会执行到这里，因为前面已检查）
  │
  └─> Validators.tryActive(validator)
      └─> 添加到 highestValidatorsSet（如果 pass == true）

┌─────────────────────────────────────────────────────────────┐
│ 重新提案（仅 4 次及以上违规需要）                                  │
└─────────────────────────────────────────────────────────────┘
如果 violationCount >= 4:
  1. 创建提案: Proposal.createProposal(validator, true, details)
  2. 验证者投票: Proposal.voteProposal(id, true)
  3. 提案通过: 
     ├─> pass[validator] = true
     └─> 方案C: 重置违规计数
         └─> violationCount[validator] = 0 (投票通过后重置)

┌─────────────────────────────────────────────────────────────┐
│ 解除监禁（投票通过后，如果验证者已注册质押）                          │
└─────────────────────────────────────────────────────────────┘
**重要场景**：如果验证者已注册质押（selfStake >= MIN_VALIDATOR_STAKE），投票通过后：
  ├─> 状态: pass[validator] = true, violationCount = 0
  ├─> 状态: isJailed = true（仍然被监禁，需要手动 unjail）
  │
  └─> 等待监禁期结束: block.number >= jailUntilBlock
      └─> 调用: Staking.unjailValidator(validator)
          ├─> 检查: violationCount <= 3 ✅（已重置为 0，通过）
          ├─> 更新: isJailed = false
          ├─> 更新: jailUntilBlock = 0
          └─> 调用: Validators.tryActive(validator)
              └─> 添加到 highestValidatorsSet（pass 已经是 true）

**注意**：
- 已注册质押的验证者不需要再次注册（registerValidator 会失败，因为 selfStake > 0）
- 投票通过后，验证者仍然被监禁，必须等待监禁期结束并手动调用 unjailValidator()
- 如果忘记调用 unjailValidator()，即使投票通过也无法恢复

┌─────────────────────────────────────────────────────────────┐
│ 重新质押（仅质押已提取时需要）                                     │
└─────────────────────────────────────────────────────────────┘
如果质押已提取 (selfStake < MIN_VALIDATOR_STAKE):
  └─> Staking.registerValidator(commissionRate) {value: >= 10000 ether}
      ├─> 检查: pass[validator] == true
      ├─> 检查: isProposalValidForStaking(validator) == true（7天内必须注册，仅用于新注册）
      └─> 检查: !isJailed || block.number >= jailUntilBlock
      └─> 注意: 如果验证者仍然被监禁，需要先 unjail 或等待监禁期结束

如果质押未提取 (selfStake >= MIN_VALIDATOR_STAKE):
  └─> 无需重新质押
  └─> 但需要等待监禁期结束并调用 unjailValidator()（如果还未 unjail）

┌─────────────────────────────────────────────────────────────┐
│ 清理惩罚记录                                                    │
└─────────────────────────────────────────────────────────────┘
重新质押时: Validators.addValidator()
  └─> Punish.cleanPunishRecord(validator)
      └─> missedBlocksCounter = 0
      └─> 从 punishValidators 数组移除

或者 unjail 时: Validators.tryActive(validator)
  └─> Punish.cleanPunishRecord(validator)
      └─> missedBlocksCounter = 0
      └─> 从 punishValidators 数组移除

┌─────────────────────────────────────────────────────────────┐
│ 等待 Epoch 更新                                                 │
└─────────────────────────────────────────────────────────────┘
等待下一个 Epoch:
  └─> 共识层调用: Staking.getTopValidators()
      ├─> 检查: selfStake >= MIN_VALIDATOR_STAKE ✅（已注册质押）
      ├─> 检查: !isJailed || block.number >= jailUntilBlock ✅（已 unjail）
      ├─> 检查: pass[validator] == true ✅（投票通过）
      └─> 返回验证者列表
  └─> 共识层调用: Validators.updateActiveValidatorSet(newSet, epoch)
      └─> 验证者重新进入验证者集合
```

### 6.10 Epoch 更新完整流程

```
┌─────────────────────────────────────────────────────────────┐
│ Prepare 阶段（Epoch 块）                                        │
└─────────────────────────────────────────────────────────────┘
共识层 Prepare()
  ├─> 获取验证者列表: getTopValidators() [使用 parent state]
  │   └─> Staking.getTopValidators()
  │       ├─> 筛选条件:
  │       │   - selfStake >= MIN_VALIDATOR_STAKE
  │       │   - !isJailed (验证者未被监禁)
  │       │   - proposalContract.pass(validator) == true
  │       │   - **注意**：不检查 `isProposalValidForStaking()`，7天有效期仅用于新注册（在 `registerValidator()` 中检查）
  │       ├─> 按总质押排序: selfStake + totalDelegated
  │       └─> 返回前 MAX_VALIDATORS (21) 个
  └─> 写入 header.Extra

┌─────────────────────────────────────────────────────────────┐
│ Finalize 阶段（Epoch 块）                                       │
└─────────────────────────────────────────────────────────────┘
共识层 Finalize()
  ├─> 惩罚错过块的验证者（可能 jail）
  ├─> 分配奖励
  └─> handleEpochTransition()
      ├─> updateValidators()
      │   ├─> Staking.getTopValidators() [使用 current state]
      │   │   └─> 返回已过滤的验证者列表（排除被 jail 的）
      │   └─> Validators.updateActiveValidatorSet(newSet, epoch)
      │       ├─> 更新 currentValidatorSet
      │       └─> 同步更新 highestValidatorsSet
      │
      └─> decreaseMissedBlocksCounter()
          └─> Punish.decreaseMissedBlocksCounter(epoch)
              └─> 减少所有验证者的 missedBlocksCounter

┌─────────────────────────────────────────────────────────────┐
│ Snapshot 更新                                                  │
└─────────────────────────────────────────────────────────────┘
共识层 snapshot.apply()
  ├─> 如果是 Epoch 块:
  │   └─> getTopValidatorsFunc() [使用 parent state]
  │       └─> Staking.getTopValidators() [使用 parent state]
  │           └─> 获取已过滤的验证者列表（排除被 jail 的）
  │       └─> 如果失败且是状态不可用错误，fallback 到 header.Extra
  │       └─> 更新 snap.Validators
  └─> 验证者集合更新完成
```

### 6.11 更新佣金率流程

```
验证者调用: Staking.updateCommissionRate(newCommissionRate)
  ├─> 检查: 验证者已注册 (onlyValidValidator)
  ├─> 检查: newCommissionRate <= COMMISSION_RATE_BASE (10000)
  ├─> 检查: 未被监禁或监禁期已过
  ├─> 更新: commissionRate = newCommissionRate
  └─> emit ValidatorUpdated(...)
```

### 6.12 编辑验证者信息流程

```
验证者调用: Validators.editValidator(feeAddr, moniker, identity, website, email, details)
  ├─> 检查: 验证者已注册 (isValidatorExist)
  ├─> 更新: validatorInfo[validator].feeAddr = feeAddr
  ├─> 更新: validatorInfo[validator].description = Description(...)
  └─> emit LogEditValidator(...)
```

---

## 七、关键机制说明

### 7.1 提案机制

**7 天注册期限**：
- 提案通过后，验证者必须在 7 天内完成注册质押，否则资格失效
- 检查: `isProposalValidForStaking(address)` 仅在 `registerValidator()` 中检查
- **重要**：7天有效期仅用于新注册验证者，已注册验证者不受此限制
- 目的：防止恶意提案和快速攻击，确保验证者及时完成注册

**投票阈值**：
- 通过条件：`当前活跃验证者同意票数 >= 活动验证者数量 / 2 + 1`
- 拒绝条件：`当前活跃验证者拒绝票数 >= 活动验证者数量 / 2 + 1`
- 活动验证者数量：`getActiveValidatorCount()`（基于 `currentValidatorSet`，已过滤被 jail 的验证者）
- **投票计数机制**：使用 `getActiveVoteCount()` 只计算当前仍然活跃的验证者的投票
  - 如果验证者投票后被踢出，其投票不再计入阈值检查
  - 这确保了投票阈值始终基于当前活跃验证者状态
  - `results[id].agree` 和 `results[id].reject` 仍然永久记录所有投票（用于历史查询）

### 7.2 Jail 机制

**Jail 触发条件**：
- 连续错过 `removeThreshold` (48) 个块

**Jail 效果**：
- `isJailed = true`
- `jailUntilBlock = block.number + 86400` (1 天)
- 立即从验证者集合中排除（Epoch 块）
- 无法出块（不在 `snap.Validators` 中）
- 无法质押、增加质押、更新佣金率
- **方案C**: `violationCount[validator]++` (增加违规计数)

**Unjail 条件**：
- 等待监禁期结束：`block.number >= jailUntilBlock`
- 验证者自己调用 `unjailValidator()`

**Unjail 后的恢复机制（方案C: 分级惩罚）**：
- **3 次以下违规** (`violationCount <= 3`)：
  - 可以 unjail（检查通过）
  - 自动恢复 `pass[validator] = true`
  - 无需重新提案投票
  - 可直接重新激活（如果质押足够）
- **4 次及以上违规** (`violationCount >= 4`)：
  - 不能 unjail（require 检查失败）
  - 必须先重新提案并通过投票
  - 投票通过后违规计数重置为 0
  - 需要重新质押（如果质押已提取）

**改过机制（方案C）**：
- 投票通过后自动重置违规计数：`violationCount[validator] = 0`
- 给予验证者改过机会，下次违规时重新开始计数

### 7.3 奖励分配机制

**交易手续费奖励**：
- 分配给所有活动验证者（排除被 jail 的）
- 每个验证者获得：`totalReward / activeValidatorCount`
- 记录到 `validatorInfo[validator].aacIncoming`
- 提取：`Validators.withdrawProfits(validator)`

**出块基础奖励（POSA 模式）**：
- 分配给当前出块的验证者和其委托者
- 验证者获得：
  - 佣金：`reward * commissionRate / 10000`
  - 验证者份额：`(reward - commission) * selfStake / totalStake`
- 委托者获得：
  - 委托份额：`(reward - commission) * delegatedAmount / totalStake`
- 提取：`Staking.claimRewards(validator)`

### 7.4 委托机制

**委托要求**：
- 最小委托金额：`MIN_DELEGATION` (1 ether)
- 验证者必须是活动验证者且未被 jail

**解绑机制**：
- 解绑期：`UNBONDING_PERIOD` (604800 块，约 7 天)
- 解绑期间代币仍计入验证者总质押
- 解绑完成后可提取

**奖励计算**：
- 使用 `rewardPerShare` 机制，确保奖励分配精确
- 每次委托/取消委托时更新 `rewardDebt`
- 提取奖励时计算：`pending = amount * rewardPerShare / 1e18 - rewardDebt`

### 7.5 验证者集合更新机制

**更新时机**：
- 仅在 Epoch 块更新（`block.number % 86400 == 0`）
- **重要**：质押、委托等操作不会立即更新验证者集合，需要等待下一个 Epoch 更新

**更新流程**：
1. Prepare 阶段：使用 parent state 获取验证者列表，写入 `header.Extra`
2. Finalize 阶段：使用 current state 更新验证者集合
   - 共识层调用 `Staking.getTopValidators()` 获取验证者列表
   - 共识层调用 `Validators.updateActiveValidatorSet(newSet, epoch)` 更新集合
   - `updateActiveValidatorSet()` 会更新 `currentValidatorSet`
   - **注意**：`highestValidatorsSet` 通过其他方法管理（如 `tryAddValidatorToHighestSet`），不会在 `updateActiveValidatorSet()` 中同步更新
3. Snapshot 更新：使用 parent state 更新 `snap.Validators`

**筛选条件**：
- `selfStake >= MIN_VALIDATOR_STAKE` (10000 ether)
- `!isJailed` (验证者未被监禁，不检查 `jailUntilBlock`，只要 `isJailed == false` 即可)
- `proposalContract.pass(validator) == true` (已通过提案)
- **注意**：不检查 `isProposalValidForStaking()`，7天有效期仅用于新注册验证者（在 `registerValidator()` 中检查）

**排序规则**：
- 按总质押排序：`selfStake + totalDelegated`
- 返回前 `MAX_VALIDATORS` (21) 个验证者

### 7.6 状态统一管理

**Jail 状态管理**：
- `Staking` 合约是 jail 状态的唯一来源
- `Validators` 合约通过代理函数查询：`staking.isValidatorJailed(validator)`
- 确保状态一致性

**验证者状态查询**：
- `Validators.isValidatorJailed()` → `Staking.isValidatorJailed()`
- `Validators.isValidatorActive()` → `Staking.getValidatorStatus()`
- `Validators.isValidatorExist()` → 检查是否有质押

**活动验证者查询**：
- `Validators.getActiveValidators()` → 基于 `currentValidatorSet`，过滤被 jailed 的验证者
- `Validators.getActiveValidatorCount()` → 基于 `currentValidatorSet`，统计未被 jailed 的验证者数量
- **重要**：这两个函数返回的是当前在共识中生效的验证者（`currentValidatorSet`），而不是所有符合条件的验证者
- `getTopValidators()` 可能包含新注册但尚未进入 `currentValidatorSet` 的验证者

### 7.7 保护机制

**最小验证者数量保护**：
- 提取质押时检查：`emergencyExit()` 检查退出后剩余验证者数量 >= `MIN_VALIDATORS` (3)
- 防止所有验证者退出导致链停止

**至少保留 1 个验证者**：
- `removeValidator()` 检查：`highestValidatorsSet.length > 1`
- 如果只有 1 个验证者，不会被移除

**重入攻击防护**：
- `Validators` 和 `Staking` 合约继承 `ReentrancyGuard`
- 关键函数使用 `nonReentrant` 修饰符：
  - `Validators.withdrawProfits()` - 提取收益
  - `Staking.withdrawValidatorStake()` - 提取质押
  - `Staking.emergencyExit()` - 紧急退出
  - `Staking.claimRewards()` - 领取奖励
- 使用 `operationsDone[block.number]` 防止区块级操作重入
- 遵循 CEI 模式（Checks-Effects-Interactions）：先更新状态，后执行外部调用

---

## 八、共识层与合约交互时序

### 8.1 正常出块流程

```
区块 N (非 Epoch 块)
├─> Prepare()
│   └─> 设置 header.Coinbase, header.Nonce
│   └─> 获取 snapshot（验证者集合）
│
├─> 执行交易
│
└─> Finalize()
    ├─> 惩罚错过块的验证者（如果有）
    │   └─> punishOutOfTurnValidator()
    │       └─> Punish.punish(validator)
    ├─> 分配交易手续费奖励（如果有交易）
    │   └─> distributeFeeReward()
    │       └─> Validators.distributeBlockReward() {value: 交易手续费}
    │           └─> 分配给所有活动验证者（排除被 jail 的）
    ├─> 分配出块基础奖励（POSA 模式）
    │   └─> distributeCoinbaseReward()
    │       └─> Staking.distributeRewards(validator) {value: 基础出块奖励}
    │           └─> 分配给验证者和委托者
    └─> 如果是 Epoch 块
    └─> handleEpochTransition()
        ├─> updateValidators()
        │   ├─> Staking.getTopValidators() 获取验证者列表
        │   └─> Validators.updateActiveValidatorSet(newSet, epoch)
        │       └─> 更新验证者集合
        └─> decreaseMissedBlocksCounter()
            └─> Punish.decreaseMissedBlocksCounter(epoch)
```

### 8.2 Epoch 块出块流程

```
区块 N (Epoch 块，N % 86400 == 0)
├─> Prepare()
│   ├─> 设置 header.Coinbase, header.Nonce
│   ├─> 获取 snapshot（验证者集合）
│   └─> 获取验证者列表并写入 header.Extra
│       └─> getTopValidators() [使用 parent state]
│           └─> Staking.getTopValidators()
│               └─> 返回已过滤的验证者列表（排除被 jail 的）
│
├─> 执行交易
│
└─> Finalize()
    ├─> 惩罚错过块的验证者（可能 jail）
    ├─> 分配交易手续费奖励
    ├─> 分配出块基础奖励
      └─> handleEpochTransition()
          ├─> updateValidators()
          │   ├─> getTopValidators() [使用 current state]
          │   │   └─> Staking.getTopValidators() [使用 current state]
          │   │       └─> 返回已过滤的验证者列表（排除被 jail 的）
          │   └─> Validators.updateActiveValidatorSet(newSet, epoch)
          │       └─> 更新 currentValidatorSet
          │       └─> **注意**：`highestValidatorsSet` 通过其他方法管理，不会在此处同步更新
          └─> decreaseMissedBlocksCounter()
              └─> Punish.decreaseMissedBlocksCounter(epoch)
                  └─> 减少所有验证者的 missedBlocksCounter
```

### 8.3 验证者出块时序

```
验证者 A 的轮次（区块 N）
├─> Prepare()
│   ├─> 检查: 是否轮到验证者 A 出块
│   │   └─> number % len(validators) == A 的索引
│   ├─> 检查: 验证者 A 是否在 snap.Validators 中
│   │   └─> 如果被 jail，不在 snap.Validators 中，无法出块
│   └─> 设置 header.Coinbase = A
│
├─> Seal()
│   ├─> 检查: 验证者 A 是否在 snap.Validators 中
│   └─> 签名区块
│
└─> Finalize()
    ├─> 如果验证者 A 成功出块:
    │   ├─> 分配交易手续费奖励（如果有交易）
    │   └─> 分配出块基础奖励给验证者 A
    │
    └─> 如果验证者 A 错过出块:
        └─> punishOutOfTurnValidator()
            └─> Punish.punish(A)
                └─> missedBlocksCounter++
```

### 8.4 验证者集合更新时序

```
Epoch 块 N (N % 86400 == 0)
├─> Prepare() [区块 N-1 的状态]
│   └─> getTopValidators() [使用 parent state]
│       └─> 获取验证者列表（基于区块 N-1 的状态）
│       └─> 写入 header.Extra
│
├─> 执行交易（可能包含质押、委托等操作）
│
└─> Finalize() [区块 N 的状态]
    ├─> 惩罚错过块的验证者（可能 jail）
    │   └─> 如果验证者被 jail，状态立即更新
    │
    └─> handleEpochTransition()
        └─> updateValidators() [使用 current state]
            └─> Staking.getTopValidators() [使用区块 N 的状态]
                └─> 过滤被 jail 的验证者
                └─> 按总质押排序
                └─> 返回前 21 名
            └─> 更新 currentValidatorSet
            └─> 更新 highestValidatorsSet
```

**关键点**：
- Prepare 阶段使用 **parent state**（区块 N-1 的状态）
- Finalize 阶段使用 **current state**（区块 N 的状态）
- 如果验证者在当前块被 jail，会在 Finalize 阶段被排除

---

## 九、关键参数和常量

### 9.1 质押相关参数

| 参数名称 | 值 | 说明 |
|---------|-----|------|
| `MIN_VALIDATOR_STAKE` | 10000 ether | 最小验证者质押金额 |
| `MIN_DELEGATION` | 1 ether | 最小委托金额 |
| `MAX_VALIDATORS` | 21 | 最大验证者数量 |
| `MIN_VALIDATORS` | 3 | 最小验证者数量（保护机制） |
| `COMMISSION_RATE_BASE` | 10000 | 佣金率基数（10000 = 100%） |

### 9.2 时间相关参数

| 参数名称 | 值 | 说明 |
|---------|-----|------|
| `Epoch` | 86400 块 | Epoch 周期（约 24 小时） |
| `VALIDATOR_UNJAIL_PERIOD` | 86400 块 | 验证者监禁期（约 24 小时） |
| `unbondingPeriod` | 604800 块（默认） | 解绑期（约 7 天），可通过提案配置（cid = 6） |
| `PROPOSAL_STAKING_DELAY` | 7 天 | 提案通过后等待期（时间戳） |

### 9.3 惩罚相关参数

| 参数名称 | 值 | 说明 |
|---------|-----|------|
| `punishThreshold` | 24 块 | 惩罚阈值（移除收入） |
| `removeThreshold` | 48 块 | 移除阈值（jail 并移除） |
| 违规自动恢复阈值 | 3 次 | 方案C: 3 次以下违规可自动恢复，4 次及以上需要重新提案 |

### 9.4 合约地址

| 合约名称 | 地址 | 说明 |
|---------|------|------|
| Validators | `0x000000000000000000000000000000000000f000` | 验证者管理合约 |
| Punish | `0x000000000000000000000000000000000000f001` | 惩罚合约 |
| Proposal | `0x000000000000000000000000000000000000f002` | 提案合约 |
| Staking | `0x000000000000000000000000000000000000f003` | 质押合约 |

---

## 十、安全机制和边界情况

### 10.1 重入攻击防护

**机制**：
- **合约级保护**：`Validators` 和 `Staking` 合约继承 `ReentrancyGuard`
- **函数级保护**：关键函数使用 `nonReentrant` 修饰符
  - `Validators.withdrawProfits()` - 提取收益
  - `Staking.withdrawValidatorStake()` - 提取质押
  - `Staking.emergencyExit()` - 紧急退出
  - `Staking.claimRewards()` - 领取奖励
- **区块级保护**：使用 `operationsDone[block.number][operation]` 标志
  - 在同一区块内，同一操作只能执行一次
  - 立即设置标志，防止重入
- **CEI 模式**：所有关键函数遵循 Checks-Effects-Interactions 模式
  - 先执行检查（Checks）
  - 再更新状态（Effects）
  - 最后执行外部调用（Interactions）

**应用场景**：
- `Validators.distributeBlockReward()` - 防止重复分配奖励
- `Validators.updateActiveValidatorSet()` - 防止重复更新验证者集合

### 10.2 状态一致性保证

**Jail 状态统一管理**：
- `Staking` 合约是 jail 状态的唯一来源
- `Validators` 合约通过代理函数查询
- 确保所有查询都返回一致的结果

**验证者集合更新**：
- 仅在 Epoch 块更新
- Prepare 和 Finalize 阶段使用不同的 state（parent vs current）
- Snapshot 机制确保验证者集合的一致性

### 10.3 边界情况处理

**所有验证者被 jail**：
- 保护机制：`removeValidator()` 确保至少保留 1 个验证者
- 如果最后一个验证者被 jail，`getTopValidators()` 可能返回空列表
- 共识层应该检查验证者列表不为空后再调用 `updateActiveValidatorSet()`，但这是可接受的安全机制

**验证者数量低于最小值**：
- `MIN_VALIDATORS` (3) 是业务指导，不是硬性技术限制
- 提取质押时检查：`activeValidatorCount > MIN_VALIDATORS`
- 防止所有验证者退出，但技术上链可以在少于 5 个验证者时运行

**状态数据库不可用**：
- 在轻节点或历史区块验证时，状态数据库可能不可用
- Fallback 机制：使用 `header.Extra` 作为备选
- 区分状态不可用错误和其他错误，只在状态不可用时 fallback

### 10.4 提案机制安全

**7 天注册期限**：
- 提案通过后，验证者必须在 7 天内完成注册质押，否则资格失效
- 防止恶意提案和快速攻击，确保验证者及时完成注册
- 检查：`isProposalValidForStaking(address)` 仅在 `registerValidator()` 中检查
- **重要**：7天有效期仅用于新注册验证者，已注册验证者不受此限制

**投票阈值**：
- 基于当前活动验证者数量（已过滤被 jail 的验证者）
- 通过条件：`当前活跃验证者同意票数 >= 活动验证者数量 / 2 + 1`
- 只计算当前仍然活跃的验证者的投票（使用 `getActiveVoteCount()`）
- **重要机制**：如果验证者投票后被踢出，其投票不再计入阈值检查
- 确保提案需要当前多数活跃验证者同意

### 10.5 奖励分配安全

**精确计算**：
- 使用 `rewardPerShare` 机制确保奖励分配精确
- 每次委托/取消委托时更新 `rewardDebt`
- 防止奖励计算错误
- **注意**：整数除法会产生微小的舍入误差（< 4 wei/次），对 18 位小数代币影响可忽略

**防重复分配**：
- `operationsDone[block.number]` 防止同一块重复分配
- 奖励分配操作在同一块只能执行一次

### 10.6 技术改进和安全增强

**SafeMath 移除**：
- 所有合约已移除 SafeMath 依赖
- 使用 Solidity 0.8+ 内置溢出检查
- 代码更简洁，Gas 成本更低

**重入保护增强**：
- `Validators` 和 `Staking` 合约继承 `ReentrancyGuard`
- 关键函数使用 `nonReentrant` 修饰符
- 遵循 CEI 模式确保状态一致性

**配置参数验证**：
- `Proposal.updateConfig()` 添加了所有参数的范围验证
- 防止配置错误（如 `decreaseRate = 0` 导致除零）
- 提高系统健壮性

**增发功能移除**：
- 移除了 `increasePeriod` 和 `receiverAddr` 配置
- 系统不再支持代币增发/通胀
- 简化了配置管理逻辑

**状态管理优化**：
- `Validator` struct 中移除了冗余的 `status` 字段
- 状态由 `Staking` 合约统一管理（`isJailed`, `jailUntilBlock`）
- `getValidatorInfo()` 动态计算状态（向后兼容）
- 减少存储成本，提高查询效率
- 确保状态一致性（单一数据源）

---

## 十一、常见问题和解答

### 11.1 验证者生命周期

**Q: 验证者从提案到出块需要多长时间？**

A: 
1. 提案创建和投票：取决于验证者投票速度（通常几小时到几天）
2. 7 天注册期限：提案通过后必须在 7 天内完成注册质押（否则资格失效）
3. 质押注册：验证者质押后立即生效
4. Epoch 更新：等待下一个 Epoch 块（最多 24 小时）
5. **总计：至少 7 天 + 投票时间（必须在7天内完成注册）**

**Q: 验证者被 jail 后如何恢复？**

A: **方案C: 分级惩罚机制**
1. 等待监禁期结束（86400 块，约 24 小时）
2. **3 次以下违规** (`violationCount <= 3`)：
   - 可以调用 `unjailValidator()` 解除监禁状态
   - 自动恢复 `pass` 状态，无需重新投票
   - 如果质押足够，可直接等待下一个 Epoch 更新
   - 如果质押不足，需要重新质押
3. **4 次及以上违规** (`violationCount >= 4`)：
   - 不能 unjail（require 检查失败）
   - 必须先重新提案并通过投票
   - 投票通过后违规计数重置为 0
   - **重要**：如果验证者已注册质押（`selfStake >= MIN_VALIDATOR_STAKE`）：
     - 投票通过后，验证者仍然被监禁（`isJailed = true`）
     - 必须等待监禁期结束并手动调用 `unjailValidator()` 解除监禁
     - 不需要再次注册（因为已注册质押）
     - 不需要重新质押（质押仍然存在）
   - 如果质押不足，需要重新质押
4. 等待下一个 Epoch 更新

### 11.2 质押和委托

**Q: 验证者可以提取部分质押吗？**

A: 可以，但需要满足以下条件：
- 提取后剩余质押 >= `MIN_VALIDATOR_STAKE` (10000 ether)
- 如果不满足，交易会失败，提示使用 `emergencyExit()` 完全退出
- `emergencyExit()` 会检查退出后剩余验证者数量 >= `MIN_VALIDATORS` (3)

**Q: 委托者可以随时取消委托吗？**

A: 可以，但代币会进入解绑期（7 天）：
- 取消委托后，代币进入解绑期
- 解绑期间代币仍计入验证者总质押
- 解绑完成后可以提取

### 11.3 奖励分配

**Q: 奖励如何分配？**

A:
- **交易手续费奖励**：分配给所有活动验证者（排除被 jail 的），每个验证者获得 `totalReward / activeCount`
- **出块基础奖励**：分配给当前出块的验证者和其委托者
  - 验证者获得：佣金 + 验证者份额
  - 委托者获得：委托份额

**Q: 如何提取奖励？**

A:
- **验证者提取佣金和验证者份额**：`Staking.claimRewards(validator)`
- **委托者提取委托奖励**：`Staking.claimRewards(validator)`
- **验证者提取交易手续费**：`Validators.withdrawProfits(validator)`

### 11.4 Epoch 更新

**Q: 验证者集合何时更新？**

A: 仅在 Epoch 块更新（`block.number % 86400 == 0`）：
- Prepare 阶段：使用 parent state 获取验证者列表，写入 `header.Extra`
- Finalize 阶段：使用 current state 更新验证者集合
- Snapshot 更新：使用 parent state 更新 `snap.Validators`

**Q: 被 jail 的验证者何时从验证者集合中移除？**

A:
- 如果当前块是 Epoch 块：立即从验证者集合中移除
- 如果当前块不是 Epoch 块：仍在 `currentValidatorSet` 中，但不在 `snap.Validators` 中，无法出块
- 下一个 Epoch 块：从 `currentValidatorSet` 中移除

---

## 十二、总结

### 12.1 核心设计原则

1. **提案机制**：所有验证者必须通过提案才能加入
2. **质押要求**：验证者必须质押至少 10000 ether
3. **7 天注册期限**：提案通过后必须在 7 天内完成注册质押（仅用于新注册，已注册验证者不受此限制）
4. **Epoch 更新**：验证者集合仅在 Epoch 块更新
5. **Jail 机制**：被 jail 的验证者立即无法出块
6. **状态统一管理**：Jail 状态由 Staking 合约统一管理
7. **分级惩罚机制（方案C）**：
   - 3 次以下违规：可以 unjail，自动恢复 pass 状态，无需重新投票
   - 4 次及以上违规：不能 unjail，必须先重新提案并通过投票
   - 改过机制：投票通过后重置违规计数

### 12.2 关键流程

1. **验证者加入**：提案 → 投票 → 等待 7 天 → 质押 → Epoch 更新
2. **验证者惩罚**：错过块 → 惩罚 → Jail → 移除 → `violationCount++` (方案C)
3. **验证者退出**：`emergencyExit()` → 检查剩余验证者数量 >= 3 → 如果在 currentValidatorSet 中先 jail → 从 allValidators 移除
4. **验证者恢复（方案C: 分级惩罚）**：
   - **3 次以下违规**：等待监禁期 → unjail → 自动恢复 pass → 重新激活（如需要重新质押）→ Epoch 更新
   - **4 次及以上违规**：
     - 等待监禁期 → 重新提案 → 投票通过（重置违规计数）
     - **已注册质押的验证者**：投票通过后仍需手动 unjail → Epoch 更新
     - **质押已提取的验证者**：投票通过后需要重新质押 → Epoch 更新
5. **改过机制（方案C）**：投票通过后 → 违规计数重置为 0
6. **奖励分配**：交易手续费（所有活动验证者）+ 出块奖励（当前验证者和委托者）

### 12.3 安全保证

1. **重入攻击防护**：
   - 使用 `ReentrancyGuard` 和 `nonReentrant` 修饰符
   - 使用 `operationsDone` 标志防止区块级重入
   - 遵循 CEI 模式确保状态一致性
2. **状态一致性**：Jail 状态统一管理
3. **边界情况处理**：至少保留 1 个验证者，防止所有验证者退出
4. **精确计算**：使用 `rewardPerShare` 机制确保奖励分配精确

---

**文档版本**：v1.4  
**最后更新**：2025-01-21  
**维护者**：POSA 开发团队

**更新内容（v1.4）：**
- 更新验证者集合更新机制：`updateValidatorSetByStake()` 已删除，改为 `updateActiveValidatorSet()`
- 明确共识层负责调用 `Staking.getTopValidators()` 获取验证者列表，然后调用 `updateActiveValidatorSet()` 更新
- 更新所有相关流程描述和检查清单

**更新内容（v1.3）：**
- 更新重入保护机制：添加 `ReentrancyGuard` 和 `nonReentrant` 修饰符说明
- 更新配置参数验证：添加所有配置参数的范围验证说明
- 移除增发功能：明确 cid 5 和 6 已移除，系统不再支持代币增发
- 更新安全机制：完善 CEI 模式说明

**更新内容（v1.2）：**
- 修正 `getTopValidators()` 筛选条件：移除 `isProposalValidForStaking()` 检查，明确只检查 `pass[validator]` 和 `!isJailed`
- 明确 7 天注册期限仅在 `registerValidator()` 中检查，不在 `getTopValidators()` 中检查
- 更新 `emergencyExit()` 流程描述，修正 `totalStaked` 更新逻辑
- 更新 `getActiveValidators()` 修复方案描述，明确与 `getTopValidators()` 的区别

**更新内容（v1.1）：**
- 更新 `emergencyExit()` 流程：添加退出后剩余验证者数量检查、jail 机制、allValidators 数组清理
- 更新 `withdrawValidatorStake()` 流程：不允许部分退出导致验证者变为非活跃
- 添加 `_removeFromAllValidators()` 函数说明
- 更新验证者退出相关 FAQ