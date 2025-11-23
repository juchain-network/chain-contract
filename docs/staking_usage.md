# JuChain Staking 质押系统使用指南

## 📋 概述

JuChain 采用创新的 JPoSA (JuChain Proof of Stake Authority) 混合共识机制，结合了 PoA 的快速确认和 PoS 的经济激励。本指南详细介绍如何使用 `congress-cli` 工具进行质押操作。

### 🎯 质押系统特性

- **🏛️ 双合约架构**: Validators + Staking 合约分工协作
- **💰 经济激励**: 验证者和委托者共享收益
- **🔒 安全机制**: 7天解绑期和作恶惩罚
- **🎖️ 治理参与**: 质押者参与网络治理决策
- **📈 动态调整**: 佣金率和质押量实时可调
- **🛡️ 安全增强**: ReentrancyGuard 保护，配置参数验证
- **⚡ 性能优化**: 移除 SafeMath，使用 Solidity 0.8+ 内置运算符

# 签名交易
./build/congress-cli sign
  --file registerValidator.json
  --key ./keystore/UTC--2024-...
  --password ./password.txt
  --chainId 210000

## 🏗️ 系统架构

```
┌─────────────────────────────────────────────────────────────┐
│                  JuChain Staking System                     │
├─────────────────┬─────────────────┬─────────────────────────┤
│   Validators    │   Delegators    │     Staking Pool        │
│                 │                 │                         │
│ • Self-stake    │ • Delegate      │ • Reward Distribution   │
│ • Commission    │ • Undelegate    │ • Slashing Protection   │
│ • Block Rewards │ • Claim Rewards │ • Governance Voting     │
│ • Fee Address   │ • Multiple Val. │ • Unbonding Queue       │
└─────────────────┴─────────────────┴─────────────────────────┘
```

## ⚙️ 环境准备

### 软件要求

| 组件 | 最低版本 | 推荐版本 | 说明 |
|------|----------|----------|------|
| **Go** | 1.23+ | 1.24+ | 编译Congress CLI |
| **Git** | 2.30+ | 最新版 | 版本控制 |
| **curl** | 7.0+ | 最新版 | RPC调用测试 |

### 构建CLI工具

```bash
# 📁 进入项目目录
cd sys-contract/congress-cli

# 🔧 编译工具
go build -o build/congress-cli

# ✅ 验证安装
./build/congress-cli version
./build/congress-cli staking --help
```

### 网络配置

#### RPC端点配置

| 网络环境 | RPC地址 | Chain ID | 说明 |
|----------|---------|----------|------|
| **主网** | `https://rpc.juchain.org` | 210000 | 生产环境 |
| **测试网** | `https://testnet-rpc.juchain.org` | 202599 | 测试环境 |
| **本地网** | `http://localhost:8545` | 202599 | 开发环境（默认） |

#### 密钥管理

```bash
# 🔑 创建密钥存储
mkdir -p keystore
chmod 700 keystore/

# 🔐 创建密码文件
echo "你的安全密码" > password.txt
chmod 600 password.txt

# 📄 准备密钥文件路径
KEYSTORE_FILE="./keystore/UTC--2024-..."
PASSWORD_FILE="./password.txt"
```

## 🎯 核心命令概览

JuChain Staking 系统提供以下核心命令：

### 📊 命令分类

#### 🔐 验证者操作

- `register-validator`: 注册为验证者并自质押
- `claim-rewards`: 提取验证者奖励

#### 💰 委托操作  

- `delegate`: 委托代币给验证者
- `undelegate`: 开始解绑委托(7天周期)
- `claim-rewards`: 提取委托奖励

#### 🔍 查询操作

- `query-validator`: 查询验证者详情
- `query-delegation`: 查询委托详情  
- `list-top-validators`: 查询顶级验证者列表

### 📈 质押参数

| 参数名称 | 最小值 | 最大值 | 说明 |
|----------|--------|--------|------|
| **验证者自质押** | 10,000 JU | 无限制 | 成为验证者的最低要求 |
| **委托金额** | 1 JU | 无限制 | 单次委托最低金额 |
| **佣金率** | 0% | 100% | 验证者可设置的佣金率 |
| **解绑期** | 7天 | 7天 | 解除委托的等待时间 |
| **最大验证者** | - | 21个 | 网络同时活跃的验证者数量 |

## 🚀 验证者操作

### 注册验证者

成为网络验证者需要完成以下步骤：

#### 步骤 1: 创建提案（由现有验证者发起）

验证者必须先通过治理提案才能注册。现有验证者需要创建添加验证者的提案：

```bash
# 📝 创建验证者添加提案
./build/congress-cli create_proposal \
  --rpc_laddr http://localhost:8545 \
  --chainId 202599 \
  --proposer 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  --target 0x新验证者地址 \
  --operation add

# 签名并发送提案
./build/congress-cli sign -f createProposal.json -k keystore -p password --chainId 202599
./build/congress-cli send -f createProposal_signed.json
```

#### 步骤 2: 验证者投票

现有验证者需要对提案进行投票（需要多数同意）：

```bash
# 🗳️ 验证者投票（赞成）
./build/congress-cli vote_proposal \
  --rpc_laddr http://localhost:8545 \
  --chainId 202599 \
  --signer 0x验证者地址 \
  --proposalId 提案ID \
  --approve

# 签名并发送投票
./build/congress-cli sign -f voteProposal.json -k keystore -p password --chainId 202599
./build/congress-cli send -f voteProposal_signed.json
```

#### 步骤 3: 等待 7 天注册期限

提案通过后，验证者必须在 **7 天内**完成注册质押，否则资格失效。

#### 步骤 4: 注册并质押

```bash
# 📝 创建验证者注册交易（必须在提案通过后7天内）
./build/congress-cli staking register-validator \
  --rpc_laddr http://localhost:8545 \
  --chainId 202599 \
  --proposer 0x新验证者地址 \
  --stake-amount 10000 \
  --commission-rate 500

# 参数说明:
# --proposer: 验证者账户地址 (必需，必须是提案通过的目标地址)
# --stake-amount: 质押金额 (最低10,000 JU)
# --commission-rate: 佣金率，以基点计算 (500 = 5%)

# 签名并发送
./build/congress-cli sign -f registerValidator.json -k keystore -p password --chainId 202599
./build/congress-cli send -f registerValidator_signed.json
```

**输出文件**: `registerValidator.json`

**重要提示**：
- ⚠️ 必须在提案通过后 **7 天内**完成注册，否则需要重新提案
- ⚠️ 注册时账户必须有足够的余额（至少 10,000 JU + Gas 费用）
- ⚠️ 注册后需要等待下一个 Epoch（约 24 小时）才能开始出块

#### 佣金率设置指南

| 佣金率 | 基点值 | 适用场景 | 竞争力 |
|--------|--------|----------|--------|
| **0-2%** | 0-200 | 新验证者吸引委托 | ⭐⭐⭐⭐⭐ |
| **3-5%** | 300-500 | 平衡收益与竞争 | ⭐⭐⭐⭐ |
| **6-10%** | 600-1000 | 成熟验证者 | ⭐⭐⭐ |
| **10%+** | 1000+ | 高质量服务 | ⭐⭐ |

### 验证者奖励提取

```bash
# 💰 提取验证者奖励
./build/congress-cli staking claim-rewards \
  --rpc_laddr http://localhost:8545 \
  --claimer 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266

# 参数说明:
# --claimer: 提取账户地址 (通常与验证者地址相同)
# --validator: 验证者地址
```

**输出文件**: `claimRewards.json`

## 💎 委托操作

### 委托代币

将代币委托给信任的验证者以获得质押奖励：

```bash
# 🤝 委托代币给验证者
./build/congress-cli staking delegate \
  --rpc_laddr http://localhost:8545 \
  --delegator 0x970e8128ab834e3eac664312d6e30df9e93cb357 \
  --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  --amount 1000

# 参数说明:
# --delegator: 委托者账户地址 (必需)
# --validator: 目标验证者地址 (必需)
# --amount: 委托金额 (最低1 JU)
```

**输出文件**: `delegate.json`

#### 委托策略建议

```bash
# 🎯 多样化委托策略
echo "=== 推荐委托策略 ==="
echo "1. 分散风险: 委托给3-5个不同验证者"
echo "2. 性能优先: 选择高在线率验证者"
echo "3. 佣金平衡: 考虑佣金率与服务质量"
echo "4. 治理参与: 支持积极参与治理的验证者"
```

### 解除委托

开始7天解绑周期，期间代币无法转移：

```bash
# 📤 解除委托 (开始7天解绑期)
./build/congress-cli staking undelegate \
  --rpc_laddr http://localhost:8545 \
  --delegator 0x970e8128ab834e3eac664312d6e30df9e93cb357 \
  --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  --amount 500

# 参数说明:
# --delegator: 委托者账户地址 (必需)
# --validator: 目标验证者地址 (必需)  
# --amount: 解绑金额
```

**输出文件**: `undelegate.json`

#### 解绑时间计算

```bash
# ⏰ 解绑期说明
echo "解绑周期: 7天 (604,800个区块)"
echo "区块时间: 1秒/块"
echo "解绑开始: 交易确认后立即开始"
echo "资金可用: 解绑期结束后可提取（使用 withdrawUnbonded）"
echo ""
echo "注意: 解绑期间代币仍计入验证者总质押，但无法转移"
```

### 委托者奖励提取

```bash
# 🎁 提取委托奖励
./build/congress-cli staking claim-rewards \
  --rpc_laddr http://localhost:8545 \
  --claimer 0x970e8128ab834e3eac664312d6e30df9e93cb357 \
  --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266

# 注意: 每个委托关系需要单独提取奖励
```

## 🔍 查询命令

### 验证者信息查询

获取验证者的详细质押和状态信息：

```bash
# 📊 查询验证者详情
./build/congress-cli staking query-validator \
  --rpc_laddr http://localhost:8545 \
  --address 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
```

#### 验证者信息输出格式

```text
✅ 验证者信息
地址: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
自质押: 10,000 JU
总委托: 50,000 JU  
总质押: 60,000 JU
佣金率: 500 基点 (5%)
是否监禁: false
监禁至区块: 0
在线率: 99.8%
最后出块: #123456
```

### 委托信息查询

查询特定委托者与验证者之间的委托详情：

```bash
# 🔍 查询委托详情
./build/congress-cli staking query-delegation \
  --rpc_laddr http://localhost:8545 \
  --delegator 0x970e8128ab834e3eac664312d6e30df9e93cb357 \
  --validator 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
```

#### 委托信息输出格式

```text
✅ 委托信息
委托者: 0x970e8128ab834e3eac664312d6e30df9e93cb357
验证者: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
委托金额: 1,000 JU
待领奖励: 25 JU
解绑金额: 0 JU
解绑完成区块: 0
委托时间: 2024-08-27 10:30:00
年化收益率: 8.5%
```

### 顶级验证者查询

获取按总质押量排序的验证者列表：

```bash
# 🏆 查询顶级验证者
./build/congress-cli staking list-top-validators \
  --rpc_laddr http://localhost:8545 \
  --limit 21
```

#### 顶级验证者输出格式

```text
✅ 顶级验证者 (按质押量排序)
总数: 21个活跃验证者

排名 | 验证者地址                                     | 总质押    | 佣金率 | 状态
-----|-----------------------------------------------|-----------|--------|------
1    | 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266     | 60,000 JU | 5%     | 活跃
2    | 0x970e8128ab834e3eac664312d6e30df9e93cb357     | 55,000 JU | 3%     | 活跃  
3    | 0x6e30df9e93cb3578ec64c67c554dddd8d1da2c25     | 50,000 JU | 7%     | 活跃
...
21   | 0x3d968443d9b72bcef4409b3a2d5e31031390fc82     | 15,000 JU | 10%    | 活跃
```

## 🔄 交易执行流程

所有质押交易都遵循标准的三步流程：

### 第一步：创建交易

```bash
# 📝 创建未签名交易
./build/congress-cli staking register-validator \
  --proposer 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  --stake-amount 10000 \
  --commission-rate 500

echo "✅ 已生成交易文件: registerValidator.json"
```

### 第二步：签名交易

```bash
# ✍️ 使用私钥签名交易
./build/congress-cli sign \
  --file registerValidator.json \
  --key ./keystore/UTC--2024-... \
  --password ./password.txt \
  --chainId 202599

echo "✅ 已生成签名文件: registerValidator_signed.json"
```

### 第三步：广播交易

```bash
# 📡 广播交易到网络
./build/congress-cli send \
  --file registerValidator_signed.json \
  --rpc_laddr http://localhost:8545

echo "✅ 交易已广播，等待区块确认..."
```

### 交易状态验证

```bash
# 🔍 验证交易结果
echo "检查交易哈希: 0x..."
echo "验证质押状态:"
./build/congress-cli staking query-validator \
  --address 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  --rpc_laddr http://localhost:8545
```

## ⚙️ 高级配置

### Gas费用优化

```bash
# 💰 自定义Gas设置 (修改JSON文件)
echo "默认Gas配置:"
echo "  gasLimit: 自动估算 + 20%缓冲"
echo "  gasPrice: 20 Gwei"
echo ""
echo "高优先级交易:"
echo "  gasPrice: 50 Gwei"
echo "  gasLimit: 手动设置更高值"
```

### 批量操作

```bash
# 🔄 批量委托示例
VALIDATORS=(
  "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
  "0x970e8128ab834e3eac664312d6e30df9e93cb357"
  "0x6e30df9e93cb3578ec64c67c554dddd8d1da2c25"
)

DELEGATOR="0x3858ffca201b0a7d75fd23bb302c12332c5e4000"
AMOUNT=1000

for validator in "${VALIDATORS[@]}"; do
  echo "委托 $AMOUNT JU 给验证者: $validator"
  
  # 创建委托交易
  ./build/congress-cli staking delegate \
    --delegator $DELEGATOR \
    --validator $validator \
    --amount $AMOUNT \
    --rpc_laddr http://localhost:8545
  
  # 重命名文件避免覆盖
  mv delegate.json delegate_${validator:2:8}.json
  
  echo "已创建交易文件: delegate_${validator:2:8}.json"
done

echo "✅ 所有委托交易已创建，请依次签名和广播"
```

### 自动化脚本

```bash
#!/bin/bash
# auto-stake.sh - 自动化质押脚本

set -e

# 配置参数
RPC_URL="http://localhost:8545"
KEYSTORE="./keystore/UTC--2024-..."
PASSWORD="./password.txt"
CHAIN_ID="202599"  # 测试网，主网使用 210000
DELEGATOR="0x3858ffca201b0a7d75fd23bb302c12332c5e4000"

# 函数：执行完整质押流程
execute_staking() {
    local operation=$1
    local validator=$2
    local amount=$3
    
    echo "🚀 执行 $operation 操作"
    echo "验证者: $validator"
    echo "金额: $amount JU"
    
    # 创建交易
    ./build/congress-cli staking $operation \
        --delegator $DELEGATOR \
        --validator $validator \
        --amount $amount \
        --rpc_laddr $RPC_URL
    
    # 签名交易  
    ./build/congress-cli sign \
        --file $operation.json \
        --key $KEYSTORE \
        --password $PASSWORD \
        --chainId $CHAIN_ID
    
    # 广播交易
    ./build/congress-cli send \
        --file ${operation}_signed.json \
        --rpc_laddr $RPC_URL
    
    echo "✅ $operation 操作完成"
    echo ""
}

# 示例：委托给多个验证者
echo "=== 开始批量委托 ==="
execute_staking "delegate" "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266" 1000
execute_staking "delegate" "0x970e8128ab834e3eac664312d6e30df9e93cb357" 1000

echo "=== 批量委托完成 ==="
```

## Configuration

### RPC Endpoints

Common JuChain RPC endpoints:

- **Mainnet**: `https://rpc.juchain.org` (Chain ID: 210000)
- **Testnet**: `https://testnet-rpc.juchain.org` (Chain ID: 202599)
- **Local**: `http://localhost:8545` (Chain ID: 202599)

### Chain IDs

- **Mainnet**: `210000`
- **Testnet**: `202599`
- **Local**: `202599` (默认，可自定义)

### Gas Configuration

The CLI automatically estimates gas with a 20% buffer. For custom gas settings, modify the transaction JSON before signing.

## Best Practices

### Security

1. **Keystore Safety**: Store keystore files securely and never share passwords
2. **Amount Verification**: Double-check stake amounts before signing
3. **Address Validation**: Verify all addresses are correct before transactions
4. **Reentrancy Protection**: All critical functions are protected by `ReentrancyGuard`
5. **Parameter Validation**: All configuration parameters have range validation to prevent errors

### Staking Strategy

1. **Validator Selection**: Research validators' performance and commission rates
2. **Diversification**: Consider delegating to multiple validators
3. **Reward Timing**: Claim rewards regularly to compound earnings
4. **Unbonding Period**: Plan for the 7-day unbonding period when undelegating

### Monitoring

1. **Regular Queries**: Monitor validator and delegation status regularly
2. **Reward Tracking**: Keep track of accumulated rewards
3. **Network Health**: Monitor network performance and validator uptime

## 🚨 故障排除

### 常见错误处理

解决质押操作中的典型问题和错误代码。

#### RPC连接失败

```text
❌ 错误: invalid RPC URL format: localhost:8545
原因: 节点未运行或RPC端口错误
```

**解决方案:**

```bash
# 检查节点状态
ps aux | grep geth
netstat -tulpn | grep :8545

# 启动节点
./build/bin/geth --config config-validator1.toml --mine

# 验证连接
curl -X POST -H "Content-Type: application/json" \
  --data '{"jsonrpc":"2.0","method":"net_version","params":[],"id":1}' \
  http://localhost:8545
```

#### 地址格式无效

```text
❌ 错误: invalid address format
原因: 地址不符合以太坊格式规范
```

**解决方案:**

```bash
# 验证地址格式 (必须是42字符，以0x开头)
echo "正确格式: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
echo "错误格式: f39Fd6e51aad88F6F4ce6aB8827279cffFb92266"

# 检查地址校验和
./build/congress-cli utils checksum-address \
  --address 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
```

#### 质押金额不足

```text
❌ 错误: stake amount must be at least 10000 JU
原因: 质押金额低于最低要求
```

**解决方案:**

```bash
echo "最低质押要求:"
echo "  验证者注册: 10,000 JU"
echo "  委托质押: 1 JU"
echo "  增加质押: 1 JU"

# 检查账户余额
./build/congress-cli utils get-balance \
  --address 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  --rpc_laddr http://localhost:8545
```

#### 佣金率超出范围

```text
❌ 错误: commission rate must be between 0 and 10000 (100%)
原因: 佣金率超过10000基点(100%)
```

**解决方案:**

```bash
echo "佣金率设置指南:"
echo "  最小值: 0 基点 (0%)"
echo "  最大值: 10000 基点 (100%)"
echo "  推荐值: 100-1000 基点 (1%-10%)"
echo "  计算公式: 百分比 × 100 = 基点"
echo ""
echo "示例:"
echo "  1% = 100 基点"
echo "  5% = 500 基点"
echo "  10% = 1000 基点"
```

#### 交易Gas费用不足

```text
❌ 错误: intrinsic gas too low
原因: Gas费用设置过低
```

**解决方案:**

```bash
# 查看当前Gas价格
echo "推荐Gas设置:"
echo "  gasPrice: 20 Gwei (正常)"
echo "  gasPrice: 50 Gwei (快速)"
echo "  gasLimit: 由系统自动估算"

# 手动设置Gas (修改JSON文件)
echo "在交易JSON中添加:"
echo '{
  "gasPrice": "0x12a05f200",
  "gasLimit": "0x5208"
}'
```

### 🔧 诊断工具

#### 网络状态检查

```bash
#!/bin/bash
# network-health.sh

echo "=== JuChain网络健康检查 ==="

# 检查RPC连接
check_rpc() {
    echo "🔍 检查RPC连接..."
    response=$(curl -s -X POST -H "Content-Type: application/json" \
        --data '{"jsonrpc":"2.0","method":"net_version","params":[],"id":1}' \
        http://localhost:8545)
    
    if [[ $response == *"210000"* ]]; then
        echo "✅ RPC连接正常 - 主网"
    elif [[ $response == *"202599"* ]]; then
        echo "✅ RPC连接正常 - 测试网/本地网"
    else
        echo "❌ RPC连接异常"
    fi
}

# 检查同步状态
check_sync() {
    echo "🔄 检查同步状态..."
    sync_status=$(curl -s -X POST -H "Content-Type: application/json" \
        --data '{"jsonrpc":"2.0","method":"eth_syncing","params":[],"id":1}' \
        http://localhost:8545)
    
    if [[ $sync_status == *"false"* ]]; then
        echo "✅ 节点已完全同步"
    else
        echo "⏳ 节点正在同步中..."
    fi
}

# 检查验证者数量
check_validators() {
    echo "👥 检查活跃验证者..."
    ./build/congress-cli staking list-top-validators \
        --limit 100 \
        --rpc_laddr http://localhost:8545 | \
        grep "Count:" || echo "❌ 无法获取验证者信息"
}

# 执行检查
check_rpc
check_sync
check_validators
echo "=== 检查完成 ==="
```

#### 质押状态监控

```bash
#!/bin/bash
# stake-monitor.sh

VALIDATOR="0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
RPC_URL="http://localhost:8545"

echo "=== 质押状态监控 ==="
echo "验证者: $VALIDATOR"
echo ""

while true; do
    echo "$(date): 检查质押状态..."
    
    # 查询验证者信息
    ./build/congress-cli staking query-validator \
        --address $VALIDATOR \
        --rpc_laddr $RPC_URL | \
        grep -E "(Total Stake|Is Jailed|Commission Rate)"
    
    echo "---"
sleep 30
done
```

## 📚 附录

### 系统合约地址

| 合约名称 | 地址 | 功能描述 |
|---------|------|----------|
| Validators | `0x000000000000000000000000000000000000f000` | 验证者管理和治理 |
| Punish | `0x000000000000000000000000000000000000f001` | 惩罚机制和监禁 |
| Proposal | `0x000000000000000000000000000000000000f002` | 治理提案投票 |
| Staking | `0x000000000000000000000000000000000000f003` | 质押和委托管理 |

### 网络参数

| 参数名称 | 主网值 | 测试网值 | 说明 |
|---------|--------|----------|------|
| Chain ID | 210000 | 202599 | 网络标识符 |
| RPC端点 | `https://rpc.juchain.org` | `https://testnet-rpc.juchain.org` | 网络接入点 |
| 出块时间 | 1秒 | 1秒 | 区块生成间隔 |
| Epoch周期 | 86400区块 | 86400区块 | ~24小时轮换 |
| 最低质押 | 10,000 JU | 10,000 JU | 验证者注册要求 |
| 最低委托 | 1 JU | 1 JU | 委托最小金额 |
| 解绑期间 | 604800区块 | 604800区块 | 7天锁定期 |
| 注册期限 | 7天 | 7天 | 提案通过后必须在此期限内注册 |

### 有用链接

- **官方文档**: [https://juchain.org/docs](https://juchain.org/docs)
- **区块浏览器**: [https://juchain.org/explorer](https://juchain.org/explorer)
- **GitHub仓库**: [https://github.com/JuChain/go-juchain](https://github.com/JuChain/go-juchain)
- **社区论坛**: [https://forum.juchain.org](https://forum.juchain.org)
- **技术支持**: [support@juchain.org](mailto:support@juchain.org)

---

**版本**: v1.2.0  
**更新时间**: 2025年1月21日  
**适用范围**: JuChain主网和测试网

**更新内容（v1.2.0）：**
- 更新安全机制说明：添加 ReentrancyGuard 保护说明
- 更新配置参数：移除增发相关配置（cid 5 和 6）
- 更新技术细节：所有合约使用 Solidity 0.8+ 内置运算符（已移除 SafeMath）

**更新内容（v1.1.0）：**
- 修正合约地址格式（使用正确的 `0x0000...f000` 格式）
- 统一 Chain ID（主网 210000，测试网 202599）
- 更新验证者注册流程，添加提案前置步骤说明
- 添加 7 天注册期限的重要提示
- 修正解绑期说明（604800 块 = 7 天）
- 更新 RPC 端点地址

*本文档会持续更新，请关注最新版本以获取准确信息。*

### File Not Found Errors

If transaction files are missing, ensure:

1. The transaction creation command completed successfully
2. You're in the correct directory
3. The file wasn't moved or deleted

### Network Issues

For network connectivity problems:

1. Verify RPC endpoint is accessible
2. Check firewall settings
3. Ensure correct chain ID
4. Confirm network is operational

## Advanced Usage

### Batch Operations

Create multiple transactions and sign them together:

```bash
# Create multiple delegation transactions
./congress-cli staking delegate --delegator 0x... --validator 0x...1 --amount 1000
./congress-cli staking delegate --delegator 0x... --validator 0x...2 --amount 1000

# Sign all transactions
./congress-cli sign --file delegate.json --key keystore.json --password password.txt
./congress-cli sign --file delegate2.json --key keystore.json --password password.txt

# Broadcast sequentially
./congress-cli send --file delegate_signed.json 
./congress-cli send --file delegate2_signed.json
```

### Scripting

Automate staking operations with shell scripts:

```bash
#!/bin/bash
RPC="http://localhost:8545"
KEYSTORE="./keystore.json"
PASSWORD="./password.txt"
CHAIN_ID="202599"

# Function to stake with a validator
stake_with_validator() {
    local validator=$1
    local amount=$2
    
    echo "Delegating $amount JU to $validator"
    
    # Create transaction
    ./congress-cli staking delegate \
        --rpc_laddr $RPC \
        --delegator 0x1234567890123456789012345678901234567890 \
        --validator $validator \
        --amount $amount
    
    # Sign transaction
    ./congress-cli sign \
        --file delegate.json \
        --key $KEYSTORE \
        --password $PASSWORD \
        --chainId $CHAIN_ID
    
    # Broadcast transaction
    ./congress-cli send \
        --file delegate_signed.json \
        --rpc_laddr $RPC
}

# Delegate to multiple validators
stake_with_validator "0x0987654321098765432109876543210987654321" 1000
stake_with_validator "0x1111111111111111111111111111111111111111" 1000
```

## Support

For additional help:

1. Check the main README in `contracts/README.md`
2. Review the command help: `./congress-cli staking [command] --help`
3. Examine transaction files for debugging
4. Verify network status and RPC connectivity
