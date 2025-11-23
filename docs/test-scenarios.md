# JuChain POSA 系统测试场景列表

本文档提供了完整的测试场景列表，用于手动验证合约和共识逻辑的正确性。

## 📋 测试环境准备

### 前置要求
- ✅ 3 个验证节点运行中（validator1, validator2, validator3）
- ✅ 1 个同步节点运行中
- ✅ `congress-cli` 工具已编译
- ✅ 所有验证者账户有足够余额（至少 10,000 JU + Gas）
- ✅ 记录当前区块高度：`___________`

### 测试账户准备
- **Validator1**: `0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266`
- **Validator2**: `0x70997970C51812dc3A010C7d01b50e0d17dc79C8`
- **Validator3**: `0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC`
- **新验证者候选**: `0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc` (validator6)
- **委托者账户**: `0x970e8128ab834e3eac664312d6e30df9e93cb357`

---

## 一、验证者生命周期测试

### 测试场景 1.1: 新验证者完整注册流程

**测试目标**: 验证新验证者从提案到激活的完整流程

**前置条件**:
- ✅ 3 个验证者正常运行
- ✅ 新验证者账户有至少 10,000 JU
- ✅ 记录当前区块高度: `___________`

**操作步骤**:

1. **创建添加验证者提案**
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
   - 记录提案ID: `___________`
   - 记录提案创建区块: `___________`

2. **验证者投票（需要至少 2/3 同意）**
   ```bash
   # Validator1 投票（赞成）
   ./build/congress-cli vote_proposal \
     -c 202599 -l http://localhost:8545 \
     -s 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
     -i <PROPOSAL_ID> \
     -a
   ./build/congress-cli sign -f voteProposal.json -k validator1_keystore -p password -c 202599
   ./build/congress-cli send -f voteProposal_signed.json -c 202599 -l http://localhost:8545
   
   # Validator2 投票（赞成）
   ./build/congress-cli vote_proposal \
     -c 202599 -l http://localhost:8545 \
     -s 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 \
     -i <PROPOSAL_ID> \
     -a
   ./build/congress-cli sign -f voteProposal.json -k validator2_keystore -p password -c 202599
   ./build/congress-cli send -f voteProposal_signed.json -c 202599 -l http://localhost:8545
   
   # Validator3 投票（可选，测试是否 2 票即可通过）
   ```
   - 记录投票完成区块: `___________`
   - 验证提案是否通过: `___________`

3. **等待 7 天注册期限**
   - 记录提案通过时间戳: `___________`
   - 等待 7 天后继续（或修改系统时间进行测试）

4. **验证者注册并质押**
   ```bash
   # 先给新验证者转账 10,000 JU
   # 然后注册
   ./build/congress-cli staking register-validator \
     -c 202599 -l http://localhost:8545 \
     --proposer 0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc \
     --stake-amount 10000 \
     --commission-rate 500
   
   ./build/congress-cli sign -f registerValidator.json \
     -k validator6_keystore -p password -c 202599
   
   ./build/congress-cli send -f registerValidator_signed.json -c 202599 -l http://localhost:8545
   ```
   - 记录注册区块: `___________`

5. **等待下一个 Epoch 更新**
   - 当前区块: `___________`
   - 下一个 Epoch 块: `___________` (当前区块向上取整到 86400 的倍数)
   - 等待 Epoch 更新

6. **验证验证者是否进入验证者集合**
   ```bash
   ./build/congress-cli staking list-top-validators -c 202599 -l http://localhost:8545
   ./build/congress-cli miners -c 202599 -l http://localhost:8545
   ```

**期望效果**:
- ✅ 提案创建成功，获得提案ID
- ✅ 2 个验证者投票后提案通过（`pass[validator] = true`）
- ✅ 7 天内注册成功（`isProposalValidForStaking()` 返回 true）
- ✅ 注册后验证者出现在 `allValidators` 中
- ✅ 下一个 Epoch 后验证者进入 `currentValidatorSet`
- ✅ 验证者可以开始出块

**实际结果**:
```
[用户填写]
```

---

### 测试场景 1.2: 验证者注册超时（7 天期限）

**测试目标**: 验证提案通过后超过 7 天无法注册

**前置条件**:
- ✅ 有一个已通过但未注册的提案
- ✅ 提案通过时间已超过 7 天

**操作步骤**:

1. **尝试注册（超过 7 天）**
   ```bash
   ./build/congress-cli staking register-validator \
     -c 202599 -l http://localhost:8545 \
     --proposer 0x新验证者地址 \
     --stake-amount 10000 \
     --commission-rate 500
   
   ./build/congress-cli sign -f registerValidator.json -k keystore -p password -c 202599
   ./build/congress-cli send -f registerValidator_signed.json -c 202599 -l http://localhost:8545
   ```

**期望效果**:
- ❌ 交易失败，错误信息包含 "Proposal expired, must repropose"
- ❌ `isProposalValidForStaking()` 返回 false

**实际结果**:
```
[用户填写]
```

---

### 测试场景 1.3: 验证者增加质押

**测试目标**: 验证已注册验证者可以增加质押

**前置条件**:
- ✅ 验证者已注册并质押至少 10,000 JU
- ✅ 验证者账户有额外余额

**操作步骤**:

1. **查询当前质押**
   ```bash
   ./build/congress-cli staking query-validator \
     -c 202599 -l http://localhost:8545 \
     --address 0x验证者地址
   ```
   - 记录当前自质押: `___________`

2. **增加质押**
   ```bash
   # 注意：需要直接调用合约或使用其他工具
   # 这里需要手动构造交易或使用 web3
   ```

3. **验证质押增加**
   ```bash
   ./build/congress-cli staking query-validator \
     -c 202599 -l http://localhost:8545 \
     --address 0x验证者地址
   ```

**期望效果**:
- ✅ 质押成功增加
- ✅ `selfStake` 更新为新值
- ✅ 下一个 Epoch 后验证者排名可能提升

**实际结果**:
```
[用户填写]
```

---

### 测试场景 1.4: 验证者部分提取质押

**测试目标**: 验证验证者可以部分提取质押（剩余 >= 10,000 JU）

**前置条件**:
- ✅ 验证者质押 > 20,000 JU
- ✅ 验证者不在 `currentValidatorSet` 中（或可以接受暂时退出）

**操作步骤**:

1. **部分提取质押**
   ```bash
   # 需要直接调用 Staking.withdrawValidatorStake(amount)
   # 确保 remainingStake >= 10,000 JU
   ```

2. **验证提取结果**
   ```bash
   ./build/congress-cli staking query-validator \
     -c 202599 -l http://localhost:8545 \
     --address 0x验证者地址
   ```

**期望效果**:
- ✅ 提取成功
- ✅ 剩余质押 >= 10,000 JU
- ✅ 验证者仍然有效（如果剩余质押 >= MIN_VALIDATOR_STAKE）

**实际结果**:
```
[用户填写]
```

---

### 测试场景 1.5: 验证者紧急退出（emergencyExit）

**测试目标**: 验证验证者可以完全退出，检查最小验证者数量保护

**前置条件**:
- ✅ 至少有 4 个活跃验证者（确保退出后 >= 3）
- ✅ 目标验证者已注册质押

**操作步骤**:

1. **查询当前活跃验证者数量**
   ```bash
   ./build/congress-cli miners -c 202599 -l http://localhost:8545
   ```
   - 记录活跃验证者数量: `___________`

2. **执行紧急退出**
   ```bash
   # 调用 Staking.emergencyExit()
   # 如果验证者在 currentValidatorSet 中，会先被 jail
   ```

3. **验证退出结果**
   ```bash
   ./build/congress-cli staking query-validator \
     -c 202599 -l http://localhost:8545 \
     --address 0x验证者地址
   
   ./build/congress-cli miners -c 202599 -l http://localhost:8545
   ```

**期望效果**:
- ✅ 如果验证者在 `currentValidatorSet` 中，先被 jail（1 epoch）
- ✅ 退出后剩余验证者数量 >= 3
- ✅ 验证者从 `allValidators` 中移除
- ✅ `selfStake` 变为 0
- ✅ 质押金额已转账回验证者账户

**实际结果**:
```
[用户填写]
```

---

### 测试场景 1.6: 紧急退出时验证者数量不足

**测试目标**: 验证当只有 3 个验证者时无法退出

**前置条件**:
- ✅ 只有 3 个活跃验证者
- ✅ 目标验证者已注册质押

**操作步骤**:

1. **尝试紧急退出**
   ```bash
   # 调用 Staking.emergencyExit()
   ```

**期望效果**:
- ❌ 交易失败
- ❌ 错误信息包含 "Cannot exit: would leave less than minimum validators"
- ❌ 验证者仍然存在

**实际结果**:
```
[用户填写]
```

---

## 二、验证者惩罚测试

### 测试场景 2.1: 验证者错过出块（轻微惩罚）

**测试目标**: 验证验证者错过出块时的惩罚机制

**前置条件**:
- ✅ 验证者正常运行
- ✅ 记录验证者当前 `missedBlocksCounter`: `___________`

**操作步骤**:

1. **停止验证者节点**
   ```bash
   # 停止 validator2 节点
   pm2 stop ju-chain-validator2
   ```

2. **等待错过多个块**
   - 记录停止区块: `___________`
   - 等待错过约 10-20 个块

3. **检查惩罚记录**
   ```bash
   # 查询 Punish 合约中的 missedBlocksCounter
   # 或查看节点日志
   ```

4. **恢复验证者节点**
   ```bash
   pm2 start ju-chain-validator2
   ```

**期望效果**:
- ✅ `missedBlocksCounter` 增加
- ✅ 如果达到 24 块（punishThreshold），验证者收入被移除
- ✅ 验证者仍然可以出块（未达到 removeThreshold）

**实际结果**:
```
[用户填写]
```

---

### 测试场景 2.2: 验证者达到惩罚阈值（24 块）

**测试目标**: 验证达到 punishThreshold 时移除验证者收入

**前置条件**:
- ✅ 验证者已错过一些块
- ✅ `missedBlocksCounter` 接近 24

**操作步骤**:

1. **停止验证者节点**
   ```bash
   pm2 stop ju-chain-validator2
   ```

2. **等待错过 24 个块**
   - 记录停止区块: `___________`
   - 等待错过 24 个块

3. **检查验证者收入**
   ```bash
   ./build/congress-cli miner \
     -c 202599 -l http://localhost:8545 \
     -a 0x70997970C51812dc3A010C7d01b50e0d17dc79C8
   ```

4. **恢复验证者节点**
   ```bash
   pm2 start ju-chain-validator2
   ```

**期望效果**:
- ✅ `missedBlocksCounter % 24 == 0` 时触发
- ✅ 验证者收入（`aacIncoming`）被移除（变为 0）
- ✅ 验证者仍然可以出块（未达到 removeThreshold）

**实际结果**:
```
[用户填写]
```

---

### 测试场景 2.3: 验证者达到移除阈值（48 块）- Jail 并移除

**测试目标**: 验证达到 removeThreshold 时验证者被 jail 并移除

**前置条件**:
- ✅ 验证者正常运行
- ✅ 至少有 4 个验证者（确保移除后 >= 3）

**操作步骤**:

1. **停止验证者节点**
   ```bash
   pm2 stop ju-chain-validator2
   ```

2. **等待错过 48 个块**
   - 记录停止区块: `___________`
   - 等待错过 48 个块（约 48 秒）

3. **检查验证者状态**
   ```bash
   ./build/congress-cli staking query-validator \
     -c 202599 -l http://localhost:8545 \
     --address 0x70997970C51812dc3A010C7d01b50e0d17dc79C8
   
   ./build/congress-cli miners -c 202599 -l http://localhost:8545
   ```

4. **验证验证者无法出块**
   - 检查节点日志，确认验证者不再出块

**期望效果**:
- ✅ `missedBlocksCounter % 48 == 0` 时触发
- ✅ 验证者被 jail（`isJailed = true`，`jailUntilBlock = block.number + 86400`）
- ✅ 验证者从 `currentValidatorSet` 中移除（下一个 Epoch）
- ✅ 验证者从 `highestValidatorsSet` 中移除（如果 length > 1）
- ✅ `pass[validator] = false`（提案状态被清除）
- ✅ `violationCount[validator]++`（违规计数增加）
- ✅ 验证者无法出块（不在 `snap.Validators` 中）

**实际结果**:
```
[用户填写]
```

---

### 测试场景 2.4: 验证者 Unjail（3 次以下违规）

**测试目标**: 验证违规次数 <= 3 时可以自动恢复

**前置条件**:
- ✅ 验证者被 jail
- ✅ `violationCount <= 3`
- ✅ 监禁期已过（`block.number >= jailUntilBlock`）

**操作步骤**:

1. **等待监禁期结束**
   - 记录 jail 区块: `___________`
   - 记录 jailUntilBlock: `___________`
   - 等待到 `jailUntilBlock`

2. **查询违规次数**
   ```bash
   # 查询 Proposal 合约中的 violationCount
   ```

3. **执行 Unjail**
   ```bash
   # 调用 Staking.unjailValidator()
   ```

4. **验证恢复结果**
   ```bash
   ./build/congress-cli staking query-validator \
     -c 202599 -l http://localhost:8545 \
     --address 0x验证者地址
   
   # 检查 pass 状态是否自动恢复
   ```

**期望效果**:
- ✅ Unjail 成功
- ✅ `isJailed = false`
- ✅ `pass[validator] = true`（自动恢复）
- ✅ `proposalPassedTime[validator] = block.timestamp`（更新时间）
- ✅ 验证者可以重新进入验证者集合（下一个 Epoch）

**实际结果**:
```
[用户填写]
```

---

### 测试场景 2.5: 验证者 Unjail 失败（4 次及以上违规）

**测试目标**: 验证违规次数 >= 4 时无法 unjail，需要重新提案

**前置条件**:
- ✅ 验证者被 jail
- ✅ `violationCount >= 4`
- ✅ 监禁期已过

**操作步骤**:

1. **尝试 Unjail**
   ```bash
   # 调用 Staking.unjailValidator()
   ```

2. **验证失败**
   ```bash
   # 检查交易是否失败
   ```

3. **重新提案并投票**
   ```bash
   # 创建提案
   ./build/congress-cli create_proposal \
     -c 202599 -l http://localhost:8545 \
     -p 0x其他验证者地址 \
     -t 0x被jail的验证者地址 \
     -o add
   
   # 投票通过
   # ...
   ```

4. **验证违规计数重置**
   ```bash
   # 查询 violationCount，应该为 0
   ```

5. **再次尝试 Unjail**
   ```bash
   # 调用 Staking.unjailValidator()
   ```

**期望效果**:
- ❌ 第一次 unjail 失败（require 检查失败）
- ✅ 重新提案并投票通过后，`violationCount` 重置为 0
- ✅ 投票通过后，`pass[validator] = true`
- ✅ 第二次 unjail 成功（因为 violationCount 已重置为 0）

**实际结果**:
```
[用户填写]
```

---

### 测试场景 2.6: Epoch 时被 Jail 的验证者立即排除

**测试目标**: 验证在 Epoch 块被 jail 的验证者立即从验证者集合中排除

**前置条件**:
- ✅ 验证者正常运行
- ✅ 接近 Epoch 块

**操作步骤**:

1. **计算下一个 Epoch 块**
   - 当前区块: `___________`
   - 下一个 Epoch: `___________` (向上取整到 86400 的倍数)

2. **在 Epoch 块前停止验证者**
   ```bash
   # 在 Epoch 块前 1-2 个块停止
   pm2 stop ju-chain-validator2
   ```

3. **等待 Epoch 块处理**
   - 观察 Epoch 块的处理

4. **检查验证者集合**
   ```bash
   ./build/congress-cli staking list-top-validators -c 202599 -l http://localhost:8545
   ./build/congress-cli miners -c 202599 -l http://localhost:8545
   ```

**期望效果**:
- ✅ 在 Epoch 块被 jail 的验证者立即从 `currentValidatorSet` 中排除
- ✅ `getTopValidators()` 返回的列表不包含被 jail 的验证者
- ✅ `header.Extra`（基于 parent state）可能包含该验证者，但 `newValidators`（基于 current state）不包含
- ✅ Epoch 验证允许这种不一致（POSA 模式）

**实际结果**:
```
[用户填写]
```

---

## 三、委托和奖励测试

### 测试场景 3.1: 委托代币给验证者

**测试目标**: 验证用户可以委托代币给验证者

**前置条件**:
- ✅ 验证者已注册并活跃
- ✅ 委托者账户有足够余额（至少 1 JU）

**操作步骤**:

1. **查询验证者信息**
   ```bash
   ./build/congress-cli staking query-validator \
     -c 202599 -l http://localhost:8545 \
     --address 0x验证者地址
   ```
   - 记录当前总委托: `___________`

2. **委托代币**
   ```bash
   ./build/congress-cli staking delegate \
     -c 202599 -l http://localhost:8545 \
     --delegator 0x970e8128ab834e3eac664312d6e30df9e93cb357 \
     --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
     --amount 1000
   
   ./build/congress-cli sign -f delegate.json -k delegator_keystore -p password -c 202599
   ./build/congress-cli send -f delegate_signed.json -c 202599 -l http://localhost:8545
   ```

3. **验证委托结果**
   ```bash
   ./build/congress-cli staking query-delegation \
     -c 202599 -l http://localhost:8545 \
     --delegator 0x970e8128ab834e3eac664312d6e30df9e93cb357 \
     --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
   
   ./build/congress-cli staking query-validator \
     -c 202599 -l http://localhost:8545 \
     --address 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
   ```

**期望效果**:
- ✅ 委托成功
- ✅ 验证者的 `totalDelegated` 增加
- ✅ 委托者的 `delegations[delegator][validator].amount` 更新
- ✅ 下一个 Epoch 后验证者排名可能提升（如果总质押增加）

**实际结果**:
```
[用户填写]
```

---

### 测试场景 3.2: 解除委托（开始解绑期）

**测试目标**: 验证解除委托进入 7 天解绑期

**前置条件**:
- ✅ 委托者已委托代币给验证者
- ✅ 记录当前区块: `___________`

**操作步骤**:

1. **查询委托信息**
   ```bash
   ./build/congress-cli staking query-delegation \
     -c 202599 -l http://localhost:8545 \
     --delegator 0x970e8128ab834e3eac664312d6e30df9e93cb357 \
     --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
   ```

2. **解除委托**
   ```bash
   ./build/congress-cli staking undelegate \
     -c 202599 -l http://localhost:8545 \
     --delegator 0x970e8128ab834e3eac664312d6e30df9e93cb357 \
     --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
     --amount 500
   
   ./build/congress-cli sign -f undelegate.json -k delegator_keystore -p password -c 202599
   ./build/congress-cli send -f undelegate_signed.json -c 202599 -l http://localhost:8545
   ```

3. **验证解绑状态**
   ```bash
   ./build/congress-cli staking query-delegation \
     -c 202599 -l http://localhost:8545 \
     --delegator 0x970e8128ab834e3eac664312d6e30df9e93cb357 \
     --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
   ```
   - 记录解绑完成区块: `___________` (当前区块 + 604800)

**期望效果**:
- ✅ 解除委托成功
- ✅ 委托金额减少
- ✅ 创建解绑记录（`unbondingDelegations`）
- ✅ `unbondingAmount` 增加
- ✅ `unbondingBlock = block.number + 604800` (7 天后)
- ✅ 解绑期间代币仍计入验证者总质押

**实际结果**:
```
[用户填写]
```

---

### 测试场景 3.3: 提取解绑代币

**测试目标**: 验证解绑期结束后可以提取代币

**前置条件**:
- ✅ 有解绑中的代币
- ✅ 解绑期已过（`block.number >= unbondingBlock`）

**操作步骤**:

1. **等待解绑期结束**
   - 当前区块: `___________`
   - 解绑完成区块: `___________`
   - 等待到解绑完成区块

2. **提取解绑代币**
   ```bash
   # 调用 Staking.withdrawUnbonded(validator, maxEntries)
   ```

3. **验证提取结果**
   ```bash
   ./build/congress-cli staking query-delegation \
     -c 202599 -l http://localhost:8545 \
     --delegator 0x970e8128ab834e3eac664312d6e30df9e93cb357 \
     --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
   ```

**期望效果**:
- ✅ 提取成功
- ✅ 代币转账到委托者账户
- ✅ 解绑记录被移除
- ✅ `unbondingAmount` 减少

**实际结果**:
```
[用户填写]
```

---

### 测试场景 3.4: 奖励分配和提取

**测试目标**: 验证出块奖励正确分配给验证者和委托者

**前置条件**:
- ✅ 验证者已注册并活跃
- ✅ 有委托者委托了代币
- ✅ 验证者已出块（获得奖励）

**操作步骤**:

1. **查询验证者奖励**
   ```bash
   ./build/congress-cli staking query-validator \
     -c 202599 -l http://localhost:8545 \
     --address 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
   ```
   - 记录 `accumulatedRewards`: `___________`

2. **查询委托者奖励**
   ```bash
   ./build/congress-cli staking query-delegation \
     -c 202599 -l http://localhost:8545 \
     --delegator 0x970e8128ab834e3eac664312d6e30df9e93cb357 \
     --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
   ```
   - 记录 `pendingRewards`: `___________`

3. **验证者提取奖励**
   ```bash
   ./build/congress-cli staking claim-rewards \
     -c 202599 -l http://localhost:8545 \
     --claimer 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
     --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
   
   ./build/congress-cli sign -f claimRewards.json -k validator_keystore -p password -c 202599
   ./build/congress-cli send -f claimRewards_signed.json -c 202599 -l http://localhost:8545
   ```

4. **委托者提取奖励**
   ```bash
   ./build/congress-cli staking claim-rewards \
     -c 202599 -l http://localhost:8545 \
     --claimer 0x970e8128ab834e3eac664312d6e30df9e93cb357 \
     --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
   
   ./build/congress-cli sign -f claimRewards.json -k delegator_keystore -p password -c 202599
   ./build/congress-cli send -f claimRewards_signed.json -c 202599 -l http://localhost:8545
   ```

5. **验证提取结果**
   ```bash
   # 查询余额变化
   # 查询奖励是否清零
   ```

**期望效果**:
- ✅ 验证者获得：佣金 + 验证者份额
- ✅ 委托者获得：委托份额
- ✅ 奖励计算正确（基于 `rewardPerShare` 机制）
- ✅ 提取后 `accumulatedRewards` 和 `pendingRewards` 清零
- ✅ 代币正确转账到账户

**实际结果**:
```
[用户填写]
```

---

### 测试场景 3.5: 交易手续费奖励分配

**测试目标**: 验证交易手续费正确分配给所有活动验证者

**前置条件**:
- ✅ 有多个活跃验证者
- ✅ 发送一些交易（产生手续费）

**操作步骤**:

1. **发送交易产生手续费**
   ```bash
   # 发送一些交易
   ```

2. **查询验证者收入**
   ```bash
   ./build/congress-cli miner \
     -c 202599 -l http://localhost:8545 \
     -a 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
   ```
   - 记录 `Accumulated Rewards`: `___________`

3. **提取交易手续费**
   ```bash
   ./build/congress-cli withdraw_profits \
     -c 202599 -l http://localhost:8545 \
     -a 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
   
   ./build/congress-cli sign -f withdrawProfits.json -k validator_keystore -p password -c 202599
   ./build/congress-cli send -f withdrawProfits_signed.json -c 202599 -l http://localhost:8545
   ```

**期望效果**:
- ✅ 交易手续费平均分配给所有活动验证者（排除被 jail 的）
- ✅ 每个验证者获得：`totalReward / activeValidatorCount`
- ✅ 提取成功，代币转账到 `feeAddr`

**实际结果**:
```
[用户填写]
```

---

## 四、Epoch 更新测试

### 测试场景 4.1: Epoch 更新验证者集合

**测试目标**: 验证 Epoch 块正确更新验证者集合

**前置条件**:
- ✅ 有多个验证者（包括新注册的）
- ✅ 接近 Epoch 块

**操作步骤**:

1. **记录当前验证者集合**
   ```bash
   ./build/congress-cli staking list-top-validators -c 202599 -l http://localhost:8545
   ./build/congress-cli miners -c 202599 -l http://localhost:8545
   ```
   - 记录当前集合: `___________`

2. **等待 Epoch 块**
   - 当前区块: `___________`
   - 下一个 Epoch: `___________`
   - 等待 Epoch 块

3. **验证 Epoch 更新**
   ```bash
   ./build/congress-cli staking list-top-validators -c 202599 -l http://localhost:8545
   ./build/congress-cli miners -c 202599 -l http://localhost:8545
   ```

**期望效果**:
- ✅ `currentValidatorSet` 更新
- ✅ `currentValidatorSet` 更新
- ✅ `highestValidatorsSet` 通过其他方法管理（如 `tryAddValidatorToHighestSet`）
- ✅ 新注册的验证者（如果质押足够）进入集合
- ✅ 被 jail 的验证者被排除
- ✅ 验证者按总质押排序

**实际结果**:
```
[用户填写]
```

---

### 测试场景 4.2: Epoch 时验证者排名变化

**测试目标**: 验证质押变化后验证者排名在 Epoch 更新

**前置条件**:
- ✅ 有多个验证者
- ✅ 验证者质押不同

**操作步骤**:

1. **记录当前排名**
   ```bash
   ./build/congress-cli staking list-top-validators -c 202599 -l http://localhost:8545
   ```

2. **增加某个验证者的质押或委托**
   ```bash
   # 增加质押或委托
   ```

3. **等待下一个 Epoch**
   - 等待 Epoch 更新

4. **验证排名变化**
   ```bash
   ./build/congress-cli staking list-top-validators -c 202599 -l http://localhost:8545
   ```

**期望效果**:
- ✅ 验证者按总质押（`selfStake + totalDelegated`）重新排序
- ✅ 质押增加的验证者排名提升
- ✅ 排名变化在 Epoch 更新时生效

**实际结果**:
```
[用户填写]
```

---

### 测试场景 4.3: Epoch 时减少惩罚计数

**测试目标**: 验证 Epoch 时 `missedBlocksCounter` 减少

**前置条件**:
- ✅ 有验证者被惩罚（`missedBlocksCounter > 0`）

**操作步骤**:

1. **记录惩罚计数**
   ```bash
   # 查询 Punish 合约中的 missedBlocksCounter
   ```
   - 记录当前计数: `___________`

2. **等待 Epoch 块**
   - 等待下一个 Epoch

3. **验证计数减少**
   ```bash
   # 查询 missedBlocksCounter
   ```

**期望效果**:
- ✅ `missedBlocksCounter` 在 Epoch 时减少（`decreaseMissedBlocksCounter`）
- ✅ 减少机制正确执行

**实际结果**:
```
[用户填写]
```

---

## 五、边界情况测试

### 测试场景 5.1: 最小验证者数量保护

**测试目标**: 验证只有 3 个验证者时无法退出

**前置条件**:
- ✅ 只有 3 个活跃验证者

**操作步骤**:

1. **查询当前验证者数量**
   ```bash
   ./build/congress-cli miners -c 202599 -l http://localhost:8545
   ```

2. **尝试紧急退出**
   ```bash
   # 调用 Staking.emergencyExit()
   ```

**期望效果**:
- ❌ 退出失败
- ❌ 错误信息包含最小验证者数量要求
- ✅ 验证者仍然存在

**实际结果**:
```
[用户填写]
```

---

### 测试场景 5.2: 最大验证者数量限制

**测试目标**: 验证最多 21 个验证者

**前置条件**:
- ✅ 已有接近 21 个验证者

**操作步骤**:

1. **尝试注册第 22 个验证者**
   ```bash
   # 完成提案、投票、注册流程
   ```

2. **验证是否进入验证者集合**
   ```bash
   ./build/congress-cli staking list-top-validators -c 202599 -l http://localhost:8545
   ```

**期望效果**:
- ✅ 注册成功（可以注册）
- ✅ 但 `getTopValidators()` 只返回前 21 个
- ✅ 第 22 个验证者不在 `currentValidatorSet` 中（如果质押不足）

**实际结果**:
```
[用户填写]
```

---

### 测试场景 5.3: 委托给被 Jail 的验证者

**测试目标**: 验证无法委托给被 jail 的验证者

**前置条件**:
- ✅ 验证者被 jail

**操作步骤**:

1. **尝试委托**
   ```bash
   ./build/congress-cli staking delegate \
     -c 202599 -l http://localhost:8545 \
     --delegator 0x委托者地址 \
     --validator 0x被jail的验证者地址 \
     --amount 1000
   ```

**期望效果**:
- ❌ 交易失败
- ❌ 错误信息包含 "Validator is jailed" 或 "onlyActiveValidator"

**实际结果**:
```
[用户填写]
```

---

### 测试场景 5.4: 提案投票时验证者被移除

**测试目标**: 验证投票后被移除的验证者投票不计入阈值

**前置条件**:
- ✅ 有进行中的提案
- ✅ 验证者已投票

**操作步骤**:

1. **验证者投票**
   ```bash
   # Validator1 投票
   ```

2. **移除验证者（通过惩罚）**
   ```bash
   # 让 Validator1 被 jail 并移除
   ```

3. **检查投票阈值**
   ```bash
   # 查询提案状态
   # 验证投票计数是否正确
   ```

**期望效果**:
- ✅ 被移除的验证者投票不计入 `getActiveVoteCount()`
- ✅ 投票阈值基于当前活跃验证者数量
- ✅ 提案通过/拒绝判断正确

**实际结果**:
```
[用户填写]
```

---

### 测试场景 5.5: 同时多个验证者被 Jail

**测试目标**: 验证多个验证者同时被 jail 的处理

**前置条件**:
- ✅ 有多个验证者

**操作步骤**:

1. **同时停止多个验证者**
   ```bash
   pm2 stop ju-chain-validator2
   pm2 stop ju-chain-validator3
   ```

2. **等待达到 removeThreshold**
   - 等待 48 个块

3. **检查验证者集合**
   ```bash
   ./build/congress-cli staking list-top-validators -c 202599 -l http://localhost:8545
   ./build/congress-cli miners -c 202599 -l http://localhost:8545
   ```

**期望效果**:
- ✅ 多个验证者同时被 jail
- ✅ 都被从验证者集合中排除
- ✅ 至少保留 1 个验证者（保护机制）
- ✅ 链继续运行

**实际结果**:
```
[用户填写]
```

---

## 六、治理提案测试

### 测试场景 6.1: 创建配置更新提案

**测试目标**: 验证可以创建系统配置更新提案

**前置条件**:
- ✅ 验证者正常运行

**操作步骤**:

1. **创建配置更新提案**
   ```bash
   ./build/congress-cli create_config_proposal \
     -c 202599 -l http://localhost:8545 \
     -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
     -i 0 \
     -v 86400
   
   # -i 0: proposalLastingPeriod
   # -v 86400: 新值（秒）
   
   ./build/congress-cli sign -f createUpdateConfigProposal.json -k keystore -p password -c 202599
   ./build/congress-cli send -f createUpdateConfigProposal_signed.json -c 202599 -l http://localhost:8545
   ```

2. **验证者投票**
   ```bash
   # 多个验证者投票
   ```

3. **验证配置更新**
   ```bash
   # 查询配置是否更新
   ```

**期望效果**:
- ✅ 提案创建成功
- ✅ 投票通过后配置更新
- ✅ 新配置生效

**实际结果**:
```
[用户填写]
```

---

### 测试场景 6.2: 移除验证者提案

**测试目标**: 验证可以通过提案移除验证者

**前置条件**:
- ✅ 有要移除的验证者
- ✅ 至少有 4 个验证者（确保移除后 >= 3）

**操作步骤**:

1. **创建移除提案**
   ```bash
   ./build/congress-cli create_proposal \
     -c 202599 -l http://localhost:8545 \
     -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
     -t 0x要移除的验证者地址 \
     -o remove
   ```

2. **投票通过**
   ```bash
   # 多个验证者投票
   ```

3. **验证移除结果**
   ```bash
   ./build/congress-cli miners -c 202599 -l http://localhost:8545
   ```

**期望效果**:
- ✅ 提案通过后验证者被移除
- ✅ `pass[validator] = false`
- ✅ 验证者从验证者集合中移除
- ✅ `violationCount[validator]++`

**实际结果**:
```
[用户填写]
```

---

## 七、性能和安全测试

### 测试场景 7.1: 大量委托操作

**测试目标**: 验证系统可以处理大量委托操作

**前置条件**:
- ✅ 有多个验证者
- ✅ 有多个委托者账户

**操作步骤**:

1. **批量委托**
   ```bash
   # 创建多个委托交易
   # 依次签名和发送
   ```

2. **验证系统性能**
   - 观察交易确认时间
   - 检查 Gas 消耗

**期望效果**:
- ✅ 所有委托成功
- ✅ 系统性能正常
- ✅ 状态更新正确

**实际结果**:
```
[用户填写]
```

---

### 测试场景 7.2: 重入攻击防护

**测试目标**: 验证系统防护重入攻击

**前置条件**:
- ✅ 了解重入攻击机制
- ✅ 验证者账户有奖励可提取

**操作步骤**:

1. **检查合约级保护**
   - `Validators` 和 `Staking` 合约继承 `ReentrancyGuard`
   - 关键函数使用 `nonReentrant` 修饰符

2. **测试函数级保护**
   - 尝试在 `withdrawProfits()` 中重入（应该失败）
   - 尝试在 `withdrawValidatorStake()` 中重入（应该失败）
   - 尝试在 `claimRewards()` 中重入（应该失败）

3. **检查区块级保护**
   - `distributeBlockReward()` 使用 `operationsDone` 标志
   - `updateActiveValidatorSet()` 使用 `operationsDone` 标志

4. **验证 CEI 模式**
   - 检查关键函数是否遵循 Checks-Effects-Interactions 模式
   - 状态更新在外部调用之前

**期望效果**:
- ✅ 合约级保护：`nonReentrant` 修饰符防止重入
- ✅ 函数级保护：重入调用会被 `ReentrancyGuard` 拒绝
- ✅ 区块级保护：同一块内同一操作只能执行一次
- ✅ CEI 模式：状态更新在外部调用之前，确保一致性
- ✅ 重入攻击被完全防护

**实际结果**:
```
[用户填写]
```

---

## 八、综合场景测试

### 测试场景 8.1: 完整验证者生命周期

**测试目标**: 完整测试验证者从注册到退出的全生命周期

**前置条件**:
- ✅ 新验证者账户

**操作步骤**:

1. **提案和注册**（参考场景 1.1）
2. **增加质押**（参考场景 1.3）
3. **接收委托**（参考场景 3.1）
4. **提取奖励**（参考场景 3.4）
5. **被惩罚和恢复**（参考场景 2.3 和 2.4）
6. **紧急退出**（参考场景 1.5）

**期望效果**:
- ✅ 所有步骤成功执行
- ✅ 状态转换正确
- ✅ 数据一致性保持

**实际结果**:
```
[用户填写]
```

---

### 测试场景 8.2: 网络压力测试

**测试目标**: 验证系统在高负载下的稳定性

**前置条件**:
- ✅ 系统正常运行

**操作步骤**:

1. **同时执行多个操作**
   - 多个验证者注册
   - 大量委托操作
   - 多个提案和投票
   - 奖励提取

2. **监控系统状态**
   - 检查区块生成速度
   - 检查交易确认时间
   - 检查 Gas 消耗

**期望效果**:
- ✅ 系统稳定运行
- ✅ 所有操作最终成功
- ✅ 性能在可接受范围内

**实际结果**:
```
[用户填写]
```

---

## 测试检查清单

### 功能检查
- [ ] 验证者注册流程完整
- [ ] 验证者惩罚机制正确
- [ ] 委托和解绑功能正常
- [ ] 奖励分配准确
- [ ] Epoch 更新正确
- [ ] 治理提案工作正常

### 安全检查
- [ ] 最小验证者数量保护
- [ ] 重入攻击防护（ReentrancyGuard + nonReentrant）
- [ ] 边界情况处理
- [ ] 状态一致性
- [ ] 配置参数验证（防止除零等错误）
- [ ] CEI 模式验证（Checks-Effects-Interactions）

### 性能检查
- [ ] 交易确认时间
- [ ] Gas 消耗合理
- [ ] 系统稳定性

---

## 测试记录模板

### 测试环境信息
- **测试日期**: `___________`
- **测试人员**: `___________`
- **测试环境**: `___________` (主网/测试网/本地)
- **节点版本**: `___________`
- **合约版本**: `___________`

### 测试结果汇总
- **总测试场景数**: `___________`
- **通过场景数**: `___________`
- **失败场景数**: `___________`
- **跳过场景数**: `___________`

### 发现的问题
1. `___________`
2. `___________`
3. `___________`

### 改进建议
1. `___________`
2. `___________`
3. `___________`

---

**文档版本**: v1.1.0  
**创建日期**: 2025-01-21  
**最后更新**: 2025-01-21

**更新内容（v1.1.0）：**
- 更新重入攻击防护测试场景：添加 ReentrancyGuard 和 nonReentrant 测试
- 更新安全机制说明：完善 CEI 模式验证
- 更新配置参数测试：移除增发相关测试（cid 5 和 6 已移除）

