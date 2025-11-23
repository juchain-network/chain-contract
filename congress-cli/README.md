# Congress CLI v1.2.0

Juchain 区块链治理命令行工具，用于验证者管理、提案投票、质押管理和网络治理。

## 🚀 新功能 (v1.2.0)

- ✅ **完整的质押管理**：支持验证者注册、委托、奖励提取等完整功能
- ✅ **配置管理**：支持配置查询和设置
- ✅ **提案查询增强**：支持查询单个提案和所有提案
- ✅ **改进的输入验证**：完整的参数验证和错误提示
- ✅ **更好的错误处理**：结构化错误消息和详细错误信息
- ✅ **增强的用户体验**：彩色输出和清晰的状态指示
- ✅ **全局参数验证**：自动验证 RPC 地址和链 ID
- ✅ **配置集中管理**：统一的常量和配置管理

## 功能概述

Congress CLI 是一个用于 Juchain 区块链治理的命令行工具，提供了完整的验证者管理、提案投票和质押管理功能。

### 核心功能

- **提案管理**：创建验证者添加/移除提案和配置更新提案
- **投票系统**：对提案进行投票表决（支持简化的投票语法）
- **验证者管理**：查询验证者信息、管理收益、编辑验证者信息
- **质押管理**：验证者注册、委托、解除委托、奖励提取
- **交易处理**：签名和发送交易到区块链网络
- **配置管理**：设置和查询 RPC 端点和链 ID
- **输入验证**：全面的参数验证和错误处理

## 安装和编译

### 前置要求

- Go 1.23.0 或更高版本
- Solidity 编译器 (solc 0.8.20)
- abigen 工具（用于生成 Go 绑定）

### 编译步骤

```bash
# 进入项目目录
cd sys-contract/congress-cli

# 编译合约并生成 Go 绑定
make proposal

# 编译可执行文件
make build

# 生成的可执行文件位于 build/congress-cli
```

### Makefile 目标

- `make build` - 编译完整项目
- `make proposal` - 生成 Proposal 合约的 Go 绑定
- `make cleanContract` - 清理生成的合约文件
- `make clean` - 清理构建文件

## 使用指南

### 全局参数

所有命令都支持以下全局参数：

- `-c, --chainId int` - 指定链ID（测试网：202599，主网：210000）
- `-l, --rpc_laddr string` - 指定RPC端点地址
  - 测试网：`https://testnet-rpc.juchain.org`
  - 主网：`https://rpc.juchain.org`
  - 本地：`http://localhost:8545`

⚠️ **注意**：新版本会自动验证这些参数的有效性

### 配置管理

可以使用配置文件来管理默认的 RPC 端点和链 ID：

```bash
# 设置默认 RPC 端点
./build/congress-cli config set --rpc https://testnet-rpc.juchain.org

# 设置默认链 ID
./build/congress-cli config set --chain-id 202599

# 查看当前配置
./build/congress-cli config list

# 获取特定配置项
./build/congress-cli config get --rpc
./build/congress-cli config get --chain-id
```

### 快速开始

1. **查看帮助和示例**：

```bash
./build/congress-cli --help
./build/congress-cli [command] --help  # 查看特定命令帮助
```

2. **查看版本信息**：

```bash
./build/congress-cli version
```

### 网络配置

**测试网络**：

```bash
# 全局参数模板
./build/congress-cli [command] -c 202599 -l https://testnet-rpc.juchain.org [其他参数]
```

**主网络**：

```bash
# 全局参数模板  
./build/congress-cli [command] -c 210000 -l https://rpc.juchain.org [其他参数]
```

---

## 一、提案管理

### 1.1 创建验证者添加/移除提案

**创建添加验证者提案：**

```bash
# 测试网示例
./build/congress-cli create_proposal \
  -p 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  -t 0x029DAB47e268575D4AC167De64052FB228B5fA41 \
  -o add \
  -c 202599 \
  -l https://testnet-rpc.juchain.org

# 主网示例
./build/congress-cli create_proposal \
  -p 0xccafa71c31bc11ba24d526fd27ba57d743152807 \
  -t 0x4d432df142823ca25b21bc3f9744ed21a275bdea \
  -o add \
  -c 210000 \
  -l https://rpc.juchain.org
```

**创建移除验证者提案：**

```bash
./build/congress-cli create_proposal \
  -p 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  -t 0x029DAB47e268575D4AC167De64052FB228B5fA41 \
  -o remove \
  -c 202599 \
  -l https://testnet-rpc.juchain.org
```

**参数说明：**
- `-p, --proposer` - 提案者地址（必须是有效验证者）
- `-t, --target` - 目标地址（要添加或移除的验证者）
- `-o, --operation` - 操作类型（`add` 或 `remove`）

**输出文件**：`createProposal.json`

### 1.2 创建配置更新提案

**支持的配置项：**

| CID | 配置项 | 说明 | 取值范围 |
|-----|--------|------|----------|
| 0 | proposalLastingPeriod | 提案持续期（秒） | 3600 - 2592000 (1小时 - 30天) |
| 1 | punishThreshold | 惩罚阈值（块数） | > 0 |
| 2 | removeThreshold | 移除阈值（块数） | > 0 |
| 3 | decreaseRate | 减少率 | > 0 |
| 4 | withdrawProfitPeriod | 提取收益周期（块数） | > 0 |
| 5 | blockReward | 区块奖励（wei） | > 0 |

**创建配置更新提案：**

```bash
# 修改提案持续期为 86400 秒（1天）
./build/congress-cli create_config_proposal \
  -p 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  -i 0 \
  -v 86400 \
  -c 202599 \
  -l https://testnet-rpc.juchain.org

# 修改区块奖励为 0.833 ether
./build/congress-cli create_config_proposal \
  -p 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  -i 5 \
  -v 833000000000000000 \
  -c 202599 \
  -l https://testnet-rpc.juchain.org
```

**参数说明：**
- `-p, --proposer` - 提案者地址
- `-i, --cid` - 配置ID（0-5）
- `-v, --value` - 新的配置值

**输出文件**：`createUpdateConfigProposal.json`

### 1.3 投票提案

⚠️ **重要**：投票语法已优化！使用 `-a` 标志表示赞成，省略表示反对。

```bash
# 赞成票（测试网示例）
./build/congress-cli vote_proposal \
  -s 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  -i b2be7f3cc702c7a24962df6aed188edbcfdebe20fd55f5670efaedace0e4bcdf \
  -a \
  -c 202599 \
  -l https://testnet-rpc.juchain.org

# 反对票（省略 -a 参数）
./build/congress-cli vote_proposal \
  -s 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  -i b2be7f3cc702c7a24962df6aed188edbcfdebe20fd55f5670efaedace0e4bcdf \
  -c 202599 \
  -l https://testnet-rpc.juchain.org
```

**参数说明：**
- `-s, --signer` - 签名者地址（必须是有效验证者）
- `-i, --proposalId` - 提案ID（64位十六进制字符串）
- `-a, --approve` - 赞成票标志（使用 `-a` 表示赞成，省略表示反对）

**输出文件**：`voteProposal.json`

### 1.4 查询提案

**查询单个提案：**

```bash
./build/congress-cli proposal \
  -i b2be7f3cc702c7a24962df6aed188edbcfdebe20fd55f5670efaedace0e4bcdf \
  -l https://testnet-rpc.juchain.org
```

**查询所有提案：**

```bash
./build/congress-cli proposals \
  -l https://testnet-rpc.juchain.org
```

**输出示例：**

```text
📋 Proposal Details:
Proposal ID: b2be7f3cc702c7a24962df6aed188edbcfdebe20fd55f5670efaedace0e4bcdf
Proposer: 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b (验证者地址)
Target Address: 0x029DAB47e268575D4AC167De64052FB228B5fA41 (待添加验证者)
Action: Add New Validator (Flag: true)
Proposal Type: 1 (Validator Management 验证者管理)
Create Time: 2025-01-21 10:30:00 UTC
Status: ✅ Passed (提案通过)
Votes: 👍 3 agree, 👎 0 reject
```

---

## 二、验证者管理

### 2.1 查询验证者信息

**查询所有验证者：**

```bash
./build/congress-cli miners \
  -l https://testnet-rpc.juchain.org
```

**查询特定验证者：**

```bash
./build/congress-cli miner \
  -a 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  -l https://testnet-rpc.juchain.org
```

**输出示例：**

```text
Address: 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b
Fee Address: 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b
Status: Active ✅
Accumulated Rewards: 5035784561530401884829
Penalized Rewards: 5323260025816819260865
Last Withdraw Block: 1206974
```

**状态说明：**
- Status 1 = Active (活跃)
- Status 2 = Inactive (异常)

### 2.2 提取验证者收益

```bash
# 测试网示例
./build/congress-cli withdraw_profits \
  -a 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  -c 202599 \
  -l https://testnet-rpc.juchain.org

# 主网示例
./build/congress-cli withdraw_profits \
  -a 0xccafa71c31bc11ba24d526fd27ba57d743152807 \
  -c 210000 \
  -l https://rpc.juchain.org
```

⚠️ **注意**：
- 提取收益有最小等待块数限制（由 `withdrawProfitPeriod` 配置项控制）
- 收益提取不需要投票流程，验证者可以直接提取自己的收益
- 需要等待足够的区块数才能提取

**输出文件**：`withdrawProfits.json`

### 2.3 编辑验证者信息

```bash
./build/congress-cli staking edit-validator \
  --validator 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  --fee-addr 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  --moniker "My Validator" \
  --identity "keybase_identity" \
  --website "https://validator.example.com" \
  --email "validator@example.com" \
  --details "Professional validator node" \
  -l https://testnet-rpc.juchain.org
```

**参数说明：**
- `--validator` - 验证者地址（必需）
- `--fee-addr` - 费用接收地址（必需）
- `--moniker` - 验证者显示名称（可选）
- `--identity` - 验证者身份标识（可选，Keybase 签名）
- `--website` - 验证者网站（可选）
- `--email` - 验证者联系邮箱（可选）
- `--details` - 验证者描述（可选）

**输出文件**：`editValidator.json`

---

## 三、质押管理

### 3.1 验证者注册

**前置条件：**
1. 验证者必须通过提案（`pass[validator] = true`）
2. 提案通过后必须在 7 天内完成注册
3. 账户必须有足够的余额（至少 10,000 JU + Gas 费用）

**注册验证者：**

```bash
./build/congress-cli staking register-validator \
  --proposer 0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc \
  --stake-amount 10000 \
  --commission-rate 500 \
  -c 202599 \
  -l https://testnet-rpc.juchain.org
```

**参数说明：**
- `--proposer` - 验证者账户地址（必需，必须是提案通过的目标地址）
- `--stake-amount` - 质押金额（必需，最低 10,000 JU）
- `--commission-rate` - 佣金率，以基点计算（必需，500 = 5%，范围 0-10000）

**重要提示**：
- ⚠️ 必须在提案通过后 **7 天内**完成注册，否则需要重新提案
- ⚠️ 注册时账户必须有足够的余额（至少 10,000 JU + Gas 费用）
- ⚠️ 注册后需要等待下一个 Epoch（约 24 小时）才能开始出块

**输出文件**：`registerValidator.json`

### 3.2 委托代币

将代币委托给信任的验证者以获得质押奖励：

```bash
./build/congress-cli staking delegate \
  --delegator 0x970e8128ab834e3eac664312d6e30df9e93cb357 \
  --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  --amount 1000 \
  -l https://testnet-rpc.juchain.org
```

**参数说明：**
- `--delegator` - 委托者账户地址（必需）
- `--validator` - 目标验证者地址（必需）
- `--amount` - 委托金额（必需，最低 1 JU）

**输出文件**：`delegate.json`

### 3.3 解除委托

开始 7 天解绑周期，期间代币无法转移：

```bash
./build/congress-cli staking undelegate \
  --delegator 0x970e8128ab834e3eac664312d6e30df9e93cb357 \
  --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  --amount 500 \
  -l https://testnet-rpc.juchain.org
```

**参数说明：**
- `--delegator` - 委托者账户地址（必需）
- `--validator` - 目标验证者地址（必需）
- `--amount` - 解绑金额（必需）

**注意**：
- 解绑周期：7 天（604,800 个区块）
- 解绑期间代币仍计入验证者总质押，但无法转移
- 解绑完成后可提取（使用 `withdrawUnbonded`，需要直接调用合约）

**输出文件**：`undelegate.json`

### 3.4 提取奖励

**验证者提取佣金和验证者份额：**

```bash
./build/congress-cli staking claim-rewards \
  --claimer 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  -c 202599 \
  -l https://testnet-rpc.juchain.org
```

**委托者提取委托奖励：**

```bash
./build/congress-cli staking claim-rewards \
  --claimer 0x970e8128ab834e3eac664312d6e30df9e93cb357 \
  --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  -c 202599 \
  -l https://testnet-rpc.juchain.org
```

**参数说明：**
- `--claimer` - 提取账户地址（必需）
- `--validator` - 验证者地址（必需）

**注意**：每个委托关系需要单独提取奖励

**输出文件**：`claimRewards.json`

### 3.5 查询验证者信息

获取验证者的详细质押和状态信息：

```bash
./build/congress-cli staking query-validator \
  --address 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  -l https://testnet-rpc.juchain.org
```

**输出示例：**

```text
✅ Validator Information
Address: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
Self Stake: 10000 JU
Total Delegated: 50000 JU
Total Stake: 60000 JU
Commission Rate: 500 basis points
Is Jailed: false
Jail Until Block: 0
```

### 3.6 查询委托信息

查询特定委托者与验证者之间的委托详情：

```bash
./build/congress-cli staking query-delegation \
  --delegator 0x970e8128ab834e3eac664312d6e30df9e93cb357 \
  --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  -l https://testnet-rpc.juchain.org
```

**输出示例：**

```text
✅ Delegation Information
Delegator: 0x970e8128ab834e3eac664312d6e30df9e93cb357
Validator: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
Delegated Amount: 1000 JU
Pending Rewards: 25 JU
Unbonding Amount: 0 JU
Unbonding Block: 0
```

### 3.7 查询顶级验证者

获取按总质押量排序的验证者列表：

```bash
./build/congress-cli staking list-top-validators \
  --limit 21 \
  -l https://testnet-rpc.juchain.org
```

**参数说明：**
- `--limit` - 显示的最大验证者数量（可选，默认 21，仅用于显示）

**输出示例：**

```text
✅ Top Validators
Total Count: 21
1. 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
2. 0x970e8128ab834e3eac664312d6e30df9e93cb357
3. 0x6e30df9e93cb3578ec64c67c554dddd8d1da2c25
...
```

---

## 四、交易处理

### 4.1 签名交易

**签名交易：**

```bash
# 测试网示例
./build/congress-cli sign \
  -f createProposal.json \
  -k /path/to/keystore/UTC--xxx \
  -p /path/to/password.txt \
  -c 202599

# 主网示例
./build/congress-cli sign \
  -f createProposal.json \
  -k /path/to/keystore/miner1.key \
  -p /path/to/password.file \
  -c 210000
```

**参数说明：**
- `-f, --file` - 交易文件路径（必需）
- `-k, --key` - Keystore 文件路径（必需）
- `-p, --password` - 密码文件路径（必需）
- `-c, --chainId` - 链ID（必需）

**输出文件**：`[原文件名]_signed.json`

### 4.2 发送已签名的交易

**发送已签名的交易：**

```bash
# 测试网示例
./build/congress-cli send \
  -f createProposal_signed.json \
  -l https://testnet-rpc.juchain.org

# 主网示例
./build/congress-cli send \
  -f createProposal_signed.json \
  -l https://rpc.juchain.org
```

**发送成功后的输出示例：**

```text
✅ Transaction broadcast successfully!
ℹ️  Transaction hash: 0xb72b3e4f2f4411fd467dcf3a4af16f12e5772a59ec91535ad18283c9a2e32ddf
ℹ️  Waiting for transaction confirmation: 0xb72b3e4f2f4411fd467dcf3a4af16f12e5772a59ec91535ad18283c9a2e32ddf
✅ Transaction confirmed in block 12535222
--------CreateProposal----------
Proposal ID: b2be7f3cc702c7a24962df6aed188edbcfdebe20fd55f5670efaedace0e4bcdf
Proposer: 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b
Destination: 0x029DAB47e268575D4AC167De64052FB228B5fA41
Flag: true
Time: 1754909524
Block: 12535222
-----
```

---

## 五、完整工作流程示例

### 5.1 添加新验证者完整流程

#### 步骤 1: 查询当前验证者状态

```bash
./build/congress-cli miners -l https://testnet-rpc.juchain.org
```

#### 步骤 2: 创建验证者添加提案

```bash
./build/congress-cli create_proposal \
  -p 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  -t 0x029DAB47e268575D4AC167De64052FB228B5fA41 \
  -o add \
  -c 202599 \
  -l https://testnet-rpc.juchain.org
```

#### 步骤 3: 签名并发送提案

```bash
./build/congress-cli sign \
  -f createProposal.json \
  -k miner1.key \
  -p password.file \
  -c 202599

./build/congress-cli send \
  -f createProposal_signed.json \
  -l https://testnet-rpc.juchain.org
```

**记录提案ID**（从输出中获取）：
```
Proposal ID: b2be7f3cc702c7a24962df6aed188edbcfdebe20fd55f5670efaedace0e4bcdf
```

#### 步骤 4: 多个验证者投票

```bash
# miner1 赞成票
./build/congress-cli vote_proposal \
  -s 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  -i b2be7f3cc702c7a24962df6aed188edbcfdebe20fd55f5670efaedace0e4bcdf \
  -a \
  -c 202599 \
  -l https://testnet-rpc.juchain.org
./build/congress-cli sign -f voteProposal.json -k miner1.key -p password.file -c 202599
./build/congress-cli send -f voteProposal_signed.json -l https://testnet-rpc.juchain.org

# miner2 赞成票
./build/congress-cli vote_proposal \
  -s 0x81f7a79a51edba249efa812eb2d5478f696f7558 \
  -i b2be7f3cc702c7a24962df6aed188edbcfdebe20fd55f5670efaedace0e4bcdf \
  -a \
  -c 202599 \
  -l https://testnet-rpc.juchain.org
./build/congress-cli sign -f voteProposal.json -k miner2.key -p password.file -c 202599
./build/congress-cli send -f voteProposal_signed.json -l https://testnet-rpc.juchain.org

# miner3 赞成票（如果需要）
./build/congress-cli vote_proposal \
  -s 0x578c39eaf09a4e1abf428c423970b59bb8baf42e \
  -i b2be7f3cc702c7a24962df6aed188edbcfdebe20fd55f5670efaedace0e4bcdf \
  -a \
  -c 202599 \
  -l https://testnet-rpc.juchain.org
./build/congress-cli sign -f voteProposal.json -k miner3.key -p password.file -c 202599
./build/congress-cli send -f voteProposal_signed.json -l https://testnet-rpc.juchain.org
```

#### 步骤 5: 等待 7 天注册期限

提案通过后，验证者必须在 **7 天内**完成注册质押，否则资格失效。

#### 步骤 6: 验证者注册并质押

```bash
./build/congress-cli staking register-validator \
  --proposer 0x029DAB47e268575D4AC167De64052FB228B5fA41 \
  --stake-amount 10000 \
  --commission-rate 500 \
  -c 202599 \
  -l https://testnet-rpc.juchain.org

./build/congress-cli sign \
  -f registerValidator.json \
  -k new_validator.key \
  -p password.file \
  -c 202599

./build/congress-cli send \
  -f registerValidator_signed.json \
  -l https://testnet-rpc.juchain.org
```

#### 步骤 7: 等待下一个 Epoch 更新

注册后需要等待下一个 Epoch（约 24 小时）才能开始出块。

#### 步骤 8: 验证验证者是否进入验证者集合

```bash
./build/congress-cli staking list-top-validators -l https://testnet-rpc.juchain.org
./build/congress-cli miners -l https://testnet-rpc.juchain.org
```

### 5.2 更新系统配置完整流程

#### 步骤 1: 创建配置更新提案

```bash
./build/congress-cli create_config_proposal \
  -p 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  -i 5 \
  -v 833000000000000000 \
  -c 202599 \
  -l https://testnet-rpc.juchain.org
```

#### 步骤 2: 签名并发送提案

```bash
./build/congress-cli sign \
  -f createUpdateConfigProposal.json \
  -k miner1.key \
  -p password.file \
  -c 202599

./build/congress-cli send \
  -f createUpdateConfigProposal_signed.json \
  -l https://testnet-rpc.juchain.org
```

#### 步骤 3: 验证者投票（同添加验证者流程）

#### 步骤 4: 验证配置更新

提案通过后，配置会自动更新。

---

## 六、配置文件

### 生成的交易文件

工具会在当前目录生成以下JSON文件：

**提案相关：**
- `createProposal.json` - 创建提案的原始交易
- `createProposal_signed.json` - 已签名的创建提案交易
- `createUpdateConfigProposal.json` - 创建配置更新提案的原始交易
- `createUpdateConfigProposal_signed.json` - 已签名的配置更新提案交易
- `voteProposal.json` - 投票的原始交易
- `voteProposal_signed.json` - 已签名的投票交易

**验证者相关：**
- `withdrawProfits.json` - 提取收益的原始交易
- `withdrawProfits_signed.json` - 已签名的提取收益交易
- `editValidator.json` - 编辑验证者信息的原始交易
- `editValidator_signed.json` - 已签名的编辑验证者信息交易

**质押相关：**
- `registerValidator.json` - 注册验证者的原始交易
- `registerValidator_signed.json` - 已签名的注册验证者交易
- `delegate.json` - 委托的原始交易
- `delegate_signed.json` - 已签名的委托交易
- `undelegate.json` - 解除委托的原始交易
- `undelegate_signed.json` - 已签名的解除委托交易
- `claimRewards.json` - 提取奖励的原始交易
- `claimRewards_signed.json` - 已签名的提取奖励交易

### Keystore 文件格式

使用标准的 Ethereum keystore 格式，例如：

```
UTC--202599-08-06T08-30-51.139143000Z--f39fd6e51aad88f6f4ce6ab8827279cffFb92266
```

---

## 七、故障排除

### 常见错误

**1. EIP-155 错误**

```
send tx error only replay-protected (EIP-155) transactions allowed over RPC
```

解决方案：签名时必须指定正确的链ID：

```bash
./build/congress-cli sign -f transaction.json -k keystore -p password -c 202599
```

**2. 提取收益失败**

```
gas estimation failed: execution reverted: You must wait enough blocks to withdraw your profits
```

解决方案：需要等待足够的区块数才能提取收益，这是正常的安全机制。

**3. 注册验证者失败**

```
execution reverted: Proposal expired, must repropose
```

解决方案：提案通过后必须在 7 天内完成注册，否则需要重新提案。

**4. 连接RPC失败**

确保：
- RPC端点地址正确
- 区块链节点正在运行
- 网络连接正常

### 调试技巧

1. 使用 `--help` 参数查看命令详细用法
2. 检查生成的JSON文件内容
3. 验证keystore文件路径和密码文件
4. 确认链ID和RPC地址配置正确
5. 使用 `config list` 查看当前配置

---

## 八、技术架构

### 项目结构

```
congress-cli/
├── cmd/                    # 命令实现
│   ├── proposal.go        # 提案相关命令
│   ├── validator.go       # 验证者相关命令
│   ├── staking.go         # 质押相关命令
│   ├── config_cmd.go      # 配置管理命令
│   ├── tools.go           # 工具函数
│   └── utils.go           # 工具函数
├── contracts/             # 合约绑定（符号链接）
│   └── generated/         # 自动生成的Go绑定
├── build/                 # 编译输出
│   └── congress-cli      # 可执行文件
├── Makefile              # 构建配置
├── go.mod                # Go模块定义
└── README.md             # 本文档
```

### 依赖项

- `github.com/ethereum/go-ethereum` - 以太坊客户端库
- `github.com/spf13/cobra` - CLI框架
- `golang.org/x/crypto` - 加密库

---

## 九、贡献指南

1. Fork 项目
2. 创建功能分支
3. 提交更改
4. 创建 Pull Request

---

## 十、许可证

本项目采用 MIT 许可证。

---

## 十一、版本历史

- **v1.2.0** - 添加完整的质押管理功能、配置管理、提案查询增强
- **v1.1.0** - 改进的输入验证、更好的错误处理、增强的用户体验
- **v1.0.0** - 初始版本，支持基本的治理功能

---

**文档版本**: v1.2.0  
**最后更新**: 2025-01-21  
**维护者**: POSA 开发团队
