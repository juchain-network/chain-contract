# Ju Congress POSA 区块链部署指南

## 📋 概述

Ju 是基于以太坊技术栈的 Congress POSA (Proof of Stake Authority) 共识区块链。本文档提供了完整的部署、配置和管理指南。

## 🏗️ 系统架构

### 核心组件

- **Geth 客户端**: 基于 go-ethereum 的区块链节点
- **Congress 共识**: 验证者投票的 POA 共识机制
- **系统合约**: 验证者管理、提案治理、惩罚机制
- **CLI 工具**: congress-cli 命令行管理工具

### 系统合约地址

- **Validators**: `0x000000000000000000000000000000000000f000`
- **Punish**: `0x000000000000000000000000000000000000f001`
- **Proposal**: `0x000000000000000000000000000000000000f002`

## ⚙️ 配置参数

### Congress 共识参数

在创世块文件中配置：

```json
"congress": {
    "period": 3,        // 出块时间间隔 (秒)
    "epoch": 200        // 验证者轮换周期 (块数)
}
```

### 系统合约参数

这些参数在合约编译时设置，修改后需重新编译：

| 参数 | 默认值 | 说明 |
|------|--------|------|
| `punishThreshold` | 24 | 没收收益阈值 (连续错过块数) |
| `removeThreshold` | 48 | 移除验证者阈值 (连续错过块数) |
| `decreaseRate` | 24 | 削减比例 |
| `withdrawProfitPeriod` | 28800 | 收益提取间隔 (块数) |
| `proposalLastingPeriod` | 86400 | 提案有效期 (秒) |
| `increasePeriod` | 100 | 增发周期 |

> **重要**: 修改合约参数需要重新编译系统合约并更新创世块文件。

## 📄 创世块配置

### 基本配置

创世块文件必须包含以下关键配置：

#### 1. 共识算法设置

```json
{
  "config": {
    "congress": {
      "period": 3,
      "epoch": 200
    }
  }
}
```

#### 2. 系统合约部署

在 `alloc` 部分预部署系统合约：

```json
{
  "alloc": {
    "000000000000000000000000000000000000f000": {
      "balance": "0x0",
      "code": "0x608060405234801561001057600080fd5b50..."
    },
    "000000000000000000000000000000000000f001": {
      "balance": "0x0", 
      "code": "0x608060405234801561001057600080fd5b50..."
    },
    "000000000000000000000000000000000000f002": {
      "balance": "0x0",
      "code": "0x608060405234801561001057600080fd5b50..."
    }
  }
}
```

#### 3. 初始验证者设置

`extraData` 字段包含初始验证者列表：

```json
{
  "extraData": "0x0000000000000000000000000000000000000000000000000000000000000000[VALIDATOR_ADDRESSES][SIGNATURE]"
}
```

> **注意**: 验证者数量建议为奇数(目前为5)，以避免投票平局。

### 获取合约字节码

系统合约字节码可通过以下方式获取：

```bash
# 1. 编译合约
cd sys-contract
forge build

# 2. 生成并更新创世块
npm run generate
npm run init-genesis
```

## 🚀 编译和部署

### 环境准备

#### 必需软件

- **Go 1.23+**: 编译 Geth 客户端
- **GCC/G++**: C++ 编译器
- **Node.js 20+**: 运行合约脚本
- **Git**: 版本控制

#### 获取源码

```bash
# 克隆主项目
git clone [repository-url]
cd ju-chain-work

# 包含两个主要目录:
# - chain/: Geth 客户端源码
# - sys-contract/: 系统合约源码
```

### 编译步骤

#### 1. 编译区块链客户端

```bash
cd chain
make all
# 或者只编译 geth
make geth

# 编译完成后在 build/bin/ 目录下找到可执行文件
ls build/bin/
```

#### 2. 编译系统合约

```bash
cd sys-contract

# 安装依赖
npm install

# 编译合约
forge build

# 生成合约代码
npm run generate

# 更新创世块文件
npm run init-genesis
```

#### 3. 编译 Congress CLI 工具

```bash
cd sys-contract/congress-cli
make build

# 测试 CLI 工具
./build/congress-cli help
```

### 节点部署

#### 1. 准备创世块文件

使用 `sys-contract/genesis.json` 或参考以太坊私链搭建文档 ([GitHub](https://github.com/ethereum/go-ethereum#operating-a-private-network)) 创建自定义创世块。

#### 2. 初始化节点

```bash
# 创建数据目录
mkdir -p node1/data

# 初始化创世块
./chain/build/bin/geth --datadir node1/data init sys-contract/genesis.json
```

#### 3. 准备验证者密钥

```bash
# 创建验证者账户
./chain/build/bin/geth --datadir node1/data account new

# 或导入现有私钥
./chain/build/bin/geth --datadir node1/data account import private.key
```

#### 4. 启动验证者节点

```bash
./chain/build/bin/geth \
  --datadir node1/data \
  --port 30303 \
  --rpc \
  --rpcport 8545 \
  --rpcapi eth,net,web3,congress \
  --allow-insecure-unlock \
  --mine \
  --unlock 0x[VALIDATOR_ADDRESS] \
  --password password.txt \
  --console
```

#### 5. 连接多个节点

**使用 static-nodes.json:**

```bash
# 在数据目录下创建 static-nodes.json
cat > node1/data/static-nodes.json << EOF
[
  "enode://[NODE1_PUBKEY]@127.0.0.1:30303",
  "enode://[NODE2_PUBKEY]@127.0.0.1:30304",
  "enode://[NODE3_PUBKEY]@127.0.0.1:30305"
]
EOF
```

**手动连接:**

```javascript
// 在 geth console 中执行
admin.addPeer("enode://[NODE_PUBKEY]@[IP]:[PORT]")
```

#### 6. 验证网络状态

```javascript
// 检查连接的节点
admin.peers

// 检查最新块
eth.blockNumber

// 检查是否在挖矿
eth.mining

// 检查验证者列表
congress.getValidators()
```

## 👥 验证者管理

### 添加新验证者

新验证者加入网络需要通过提案和投票流程：

#### 1. 准备新验证者

```bash
# 准备新验证者账户和节点
./chain/build/bin/geth --datadir newnode/data account new

# 记录新验证者地址
NEW_VALIDATOR_ADDR="0x[NEW_ADDRESS]"
```

#### 2. 创建提案

```bash
cd sys-contract/congress-cli

# 由现有验证者创建添加提案
./build/congress-cli create_proposal \
  -p $PROPOSER_ADDR \
  -t $NEW_VALIDATOR_ADDR \
  -o add \
  --rpc_laddr http://localhost:8545

# 签名交易
./build/congress-cli sign \
  -f createProposal.json \
  -k proposer.key \
  -p password.txt \
  --chainId 202599

# 发送交易
./build/congress-cli send \
  -f createProposal_signed.json \
  --rpc_laddr http://localhost:8545
```

#### 3. 验证者投票

```bash
# 其他验证者对提案投票
./build/congress-cli vote_proposal \
  -s $VOTER_ADDR \
  -i $PROPOSAL_ID \
  -a true \
  --rpc_laddr http://localhost:8545

# 签名并发送投票
./build/congress-cli sign -f voteProposal.json -k voter.key -p password.txt --chainId 202599
./build/congress-cli send -f voteProposal_signed.json --rpc_laddr http://localhost:8545
```

#### 4. 启动新验证者节点

提案通过后，启动新验证者节点：

```bash
./chain/build/bin/geth \
  --datadir newnode/data \
  --port 30306 \
  --rpc \
  --rpcport 8546 \
  --allow-insecure-unlock \
  --mine \
  --unlock $NEW_VALIDATOR_ADDR \
  --password password.txt
```

### 查询验证者信息

```bash
# 查询所有验证者
./build/congress-cli miners --rpc_laddr http://localhost:8545

# 查询特定验证者
./build/congress-cli miner --rpc_laddr http://localhost:8545 -a $VALIDATOR_ADDR
```

## 🔧 配置管理

### 更新系统配置

系统配置更新需要通过提案和投票流程：

#### 1. 创建配置更新提案

```bash
# 配置项说明：
# 0: proposalLastingPeriod (提案有效期)
# 1: punishThreshold (惩罚阈值)  
# 2: removeThreshold (移除阈值)
# 3: decreaseRate (削减比例)
# 4: withdrawProfitPeriod (收益提取间隔)

# 示例：修改提案有效期为 86400 秒
./build/congress-cli create_config_proposal \
  -p $PROPOSER_ADDR \
  -i 0 \
  -v 86400 \
  --rpc_laddr http://localhost:8545

# 签名并发送
./build/congress-cli sign -f createUpdateConfigProposal.json -k proposer.key -p password.txt --chainId 202599
./build/congress-cli send -f createUpdateConfigProposal_signed.json --rpc_laddr http://localhost:8545
```

#### 2. 验证者投票

配置更新提案的投票流程与新增验证者相同，需要超过半数验证者同意。

### 验证者收益管理

#### 收益提取

验证者可以提取累积的手续费收益：

```bash
# 创建收益提取交易
./build/congress-cli withdraw_profits \
  -a $VALIDATOR_ADDR \
  --rpc_laddr http://localhost:8545

# 签名并发送
./build/congress-cli sign -f withdrawProfits.json -k validator.key -p password.txt --chainId 202599
./build/congress-cli send -f withdrawProfits_signed.json --rpc_laddr http://localhost:8545
```

#### 提取限制

- 两次提取之间需间隔 `withdrawProfitPeriod` 个区块
- 只有设置的 `feeAddr` 可以提取对应验证者的收益
- 被惩罚的验证者无法提取收益

## 📚 合约接口

### Proposal 合约

#### 创建提案

```solidity
function createProposal(
    address dst,      // 目标验证者地址
    bool flag,        // true: 添加, false: 移除
    string calldata details  // 提案描述
) external
```

#### 投票

```solidity
function voteProposal(
    bytes32 id,       // 提案ID
    bool auth         // true: 同意, false: 反对
) external
```

### Validators 合约

#### 创建验证者信息

```solidity
function createOrEditValidator(
    address payable feeAddr,    // 收益地址
    string calldata moniker,    // 验证者名称
    string calldata identity,   // 身份标识
    string calldata website,    // 网站
    string calldata email,      // 邮箱
    string calldata details     // 详细描述
) external
```

#### 提取收益

```solidity
function withdrawProfits(address validator) external
```

## 🔍 监控与管理

### 系统状态查询

#### 查看当前验证者

```bash
# 获取当前验证者列表
./build/congress-cli get_validators --rpc_laddr http://localhost:8545

# 查看指定验证者信息
./build/congress-cli get_validator_info \
  -a $VALIDATOR_ADDR \
  --rpc_laddr http://localhost:8545
```

#### 查看系统参数

```bash
# 查看所有系统参数
./build/congress-cli get_params --rpc_laddr http://localhost:8545

# 查看具体参数
curl -X POST http://localhost:8545 \
  -H "Content-Type: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "method": "eth_call",
    "params": [{
      "to": "0x000000000000000000000000000000000000f001",
      "data": "0x5c19a95c"
    }, "latest"],
    "id": 1
  }'
```

### 事件监听

#### 监听验证者变更

```javascript
// 监听验证者添加事件
web3.eth.subscribe('logs', {
  address: '0x000000000000000000000000000000000000f000',
  topics: ['0x...'] // LogCreateValidator
})
.on('data', log => {
  console.log('新验证者添加:', log);
});

// 监听提案事件
web3.eth.subscribe('logs', {
  address: '0x000000000000000000000000000000000000f002',
  topics: ['0x...'] // LogCreateProposal
})
.on('data', log => {
  console.log('新提案创建:', log);
});
```

### 健康检查

```bash
# 检查节点同步状态
curl -X POST http://localhost:8545 \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"eth_syncing","params":[],"id":1}'

# 检查最新区块
curl -X POST http://localhost:8545 \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}'
```

## 🛠️ 故障排除

### 常见问题

#### 提案创建失败

**问题**: 创建提案时交易失败

**解决方案**:

1. 确认发起者是当前验证者
2. 检查gas费用设置
3. 确认网络连接正常

```bash
# 检查账户是否为验证者
./build/congress-cli is_validator \
  -a $PROPOSER_ADDR \
  --rpc_laddr http://localhost:8545
```

#### 验证者无法出块

**问题**: 验证者节点无法产生区块

**解决方案**:

1. 检查节点同步状态
2. 确认验证者是否被惩罚
3. 检查网络连接

#### 收益无法提取

**问题**: 调用 withdrawProfits 失败

**解决方案**:

1. 确认提取间隔是否足够
2. 检查是否使用正确的 feeAddr
3. 确认验证者未被惩罚

## 📋 总结

Ju 区块链部署涉及以下关键组件：

1. **consensus**: Congress POA 共识机制
2. **系统合约**: Validators、Punish、Proposal 三个核心合约
3. **congress-cli**: 命令行管理工具
4. **监控系统**: 实时监控验证者状态和系统参数

通过本指南，您应该能够：

- 成功部署 Ju 区块链网络
- 管理验证者的加入和移除
- 监控系统运行状态
- 处理常见故障

### 参考资源

- [Congress CLI 使用指南](./congress.md)
- [系统合约 API 文档](../contracts/)
- [Foundry 框架文档](https://book.getfoundry.sh/)
- [Go-Ethereum 文档](https://geth.ethereum.org/docs/)
