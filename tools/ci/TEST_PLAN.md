# JuChain 集成测试大纲 (Integration Test Plan)

本文档规划了 JuChain 系统合约的端到端集成测试路径。为确保测试效率，测试执行顺序至关重要。

**执行策略**:
1.  **参数调整 (Phase 0)**: 首先通过治理提案将系统的时间参数（如解绑期、监禁期）缩短，为后续测试创造条件。
2.  **核心功能 (Phase 1)**: 在新参数下执行验证者准入、质押、委托等全流程。
3.  **边界与异常 (Phase 2)**: 穿插进行各种异常场景测试。

---

## 0. 系统参数调整与测试 (System Config & Setup) - **PRIORITY 1**

此阶段既是测试“修改参数”功能的正确性，也是为后续测试做准备。

### 0.1 正常流程 (Setup)
*   **[C-01] 缩短关键周期参数**
    *   **目标**:
        *   `unbondingPeriod`: 设为 100 blocks (原 7天) -> 允许测试取款。
        *   `validatorUnjailPeriod`: 设为 50 blocks (原 1天) -> 允许测试 Unjail。
        *   `proposalCooldown`: 设为 10 blocks (原 100) -> 允许连续提案。
        *   `proposalLastingPeriod`: 设为 200 blocks (原 7天) -> 既保证有时间投票，又方便测试过期。
        *   `withdrawProfitPeriod`: 设为 20 blocks (原 1天) -> 允许快速测试提取收益。
        *   `commissionUpdateCooldown`: 设为 50 blocks (原 7天) -> 允许测试修改佣金。
    *   **步骤**:
        1.  验证者发起 ProposalType=2 的提案 (CID 对应上述参数)。
        2.  其他验证者投票通过。
        3.  检查 `Params` 合约中的变量值是否已更新。
    *   **预期**: 参数更新成功，Event `LogPassProposal` 触发。

---

## 1. 治理与验证者准入 (Governance & Onboarding)

### 1.1 正常流程
*   **[G-01] 新验证者准入流程**: 提案 -> 投票 -> 注册。
*   **[G-02] 移除验证者流程**: 提案 -> 投票 -> 移除。
*   **[G-03] 验证者复活流程 (Re-onboarding)**: 针对已移除地址再次提案准入。
*   **[G-04] 否决提案流程 (Reject Proposal)**: 投票拒绝后的状态验证。

### 1.2 高级与组合
*   **[G-13] 连续添加与移除 (Flip-Flop)**: 验证状态切换的干净程度。
*   **[G-14] 提案与参数修改并行 (Parallel Governance)**: 验证 Nonce 机制确保 ID 唯一。
*   **[G-15] 动态阈值 (Dynamic Threshold)**: 投票期间验证者集合变更，检查通过阈值是否动态调整。
*   **[G-17] 提案 ID 冲突 (Nonce Handling)**: 多个验证者针对同一目标发起相同提案，应生成不同 ID。

---

## 2. 质押与验证者管理 (Staking & Management)

### 2.1 正常流程
*   **[S-01] 增加质押 (Add Stake)**
*   **[S-02] 减少质押 (Decrease Stake)**
*   **[S-03] 修改信息 (Edit Info)**
*   **[S-04] 修改佣金 (Update Commission)**

### 2.2 高级与组合
*   **[S-05] 验证者重生流程 (Reincarnation)**: Resign -> Exit -> Propose -> Register。
*   **[S-17] 频繁质押变更与收益 (Stake Jitter)**: 验证收益在质押变动时的累积。
*   **[S-18] 混合质押操作 (Mixed Stakes)**: 验证者自质押与委托同时变动时的比例分配。
*   **[S-16] 零委托收益归属**: 验证者无委托人时，收益全额归属。

### 2.3 异常与边界
*   **[S-15] 提案有效期限制 (7-Day Rule)**: 提案通过超过 7 天后尝试注册应失败。
*   **[S-19] 质押队列满 (Unbonding Limit)**: 验证者连续减少质押超过 20 次应触发限制。
*   **[S-20] 退出阻塞期 (DoubleSignWindow)**: 出块后立即申请退出应被拦截。
*   **[S-14] Jailed 期间约束**: 监禁状态下禁止修改佣金，但允许追加质押。

---

## 3. 委托与奖励 (Delegation & Rewards)

### 3.1 正常流程
*   **[D-01] 全流程委托与赎回**: Delegate -> Claim -> Undelegate -> Withdraw。
*   **[D-02] 验证者提取佣金**: 周期性提取。

### 3.2 组合场景
*   **[D-04] 多用户委托隔离**: A 和 B 委托同一验证者，收益互不干扰。
*   **[D-15] 身份升级**: 委托人申请并成为验证者。
*   **[D-17] 角色降级**: 验证者退出后继续作为委托人。

---

## 4. 惩罚与退出 (Punishment & Exit)

### 4.1 正常流程
*   **[P-07] 双签证据提交 (Submit Evidence)**: 构造 RLP 证据触发 Slash。
*   **[P-08] 提取手续费收益**: `withdrawProfits` 逻辑。

### 4.2 组合与健壮性
*   **[V-01] 收益分配回退 (Re-distribution)**: 验证者在 Epoch 中被 Jailed，其出块收益应分给其他活跃者。
*   **[P-23] 惩罚计数 Epoch 递减**: 验证计数器随时间自动清理。
*   **[P-24] 待处理队列自动执行**: 验证 `executePending` 在下一区块成功执行。

### 4.3 异常边界
*   **[V-02] 描述信息超长**: 验证 `validateDescription` 的长度校验。
*   **[V-04] 非指定地址提取收益**: 安全拦截。
*   **[P-22] 退出后的双签**: 验证已退出验证者不再接受双签惩罚。

---

**注意**:
*   所有测试执行前，请确保 `tools/ci/config.yaml` 配置正确。
*   关键异常路径必须覆盖 `Revert` 信息断言。
