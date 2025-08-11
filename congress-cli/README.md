# Congress CLI

Juchain 区块链治理命令行工具，用于验证者管理、提案投票和网络治理。

## 功能概述

Congress CLI 是一个用于 Juchain 区块链治理的命令行工具，提供了完整的验证者管理和提案投票功能。

### 核心功能

- **提案管理**：创建验证者添加/移除提案和配置更新提案
- **投票系统**：对提案进行投票表决
- **验证者管理**：查询验证者信息和管理收益
- **交易处理**：签名和发送交易到区块链网络

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

- `-c, --chainId int` - 指定链ID（例如：202599）
- `-l, --rpc_laddr string` - 指定RPC端点地址（例如：http://localhost:8545）

### 命令详解

#### 1. 查看版本信息

```bash
./build/congress-cli version
```

#### 2. 查询验证者信息

**查询所有验证者：**

```bash
./build/congress-cli miners -l http://localhost:8545 -c 202599
```

**查询特定验证者：**

```bash
./build/congress-cli miner -l http://localhost:8545 -c 202599 -a 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
```

输出示例：

```
矿工  0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 奖励地址 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 活动状态 1 累计奖励 345564000000000 罚没奖励 0 上次提取奖励区块 0
```

#### 3. 创建提案

**创建验证者添加/移除提案：**

```bash
# 添加验证者
./build/congress-cli create_proposal -l http://localhost:8545 -c 202599 \
  -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  -t 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 \
  -o add

# 移除验证者
./build/congress-cli create_proposal -l http://localhost:8545 -c 202599 \
  -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  -t 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 \
  -o remove
```

参数说明：

- `-p` - 提案者地址（必须是有效验证者）
- `-t` - 目标地址（要添加或移除的验证者）
- `-o` - 操作类型（add 或 remove）

**创建配置更新提案：**

```bash
./build/congress-cli create_config_proposal -l http://localhost:8545 -c 202599 \
  -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  -i 0 \
  -v 1000
```

参数说明：

- `-p` - 提案者地址
- `-i` - 配置ID：
  - 0: proposalLastingPeriod（提案持续期）
  - 1: punishThreshold（惩罚阈值）
  - 2: removeThreshold（移除阈值）
  - 3: decreaseRate（减少率）
  - 4: withdrawProfitPeriod（提取收益周期）
- `-v` - 新的配置值

#### 4. 投票提案

```bash
# 赞成票
./build/congress-cli vote_proposal -l http://localhost:8545 -c 202599 \
  -s 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  -i 4881cb992217b8d050a87384de6998efb1f051c0cb62c10bdd393f68336618b6 \
  -a

# 反对票（移除 -a 参数）
./build/congress-cli vote_proposal -l http://localhost:8545 -c 202599 \
  -s 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  -i 4881cb992217b8d050a87384de6998efb1f051c0cb62c10bdd393f68336618b6
```

参数说明：

- `-s` - 签名者地址（必须是有效验证者）
- `-i` - 提案ID（从创建提案的输出中获取）
- `-a` - 赞成票（可选，不加此参数表示反对票）

#### 5. 签名和发送交易

**签名交易：**

```bash
./build/congress-cli sign \
  -f createProposal.json \
  -k /path/to/keystore/UTC--xxx \
  -p /path/to/password.txt \
  -c 202599
```

**发送已签名的交易：**

```bash
./build/congress-cli send \
  -f createProposal_signed.json \
  -l http://localhost:8545
```

#### 6. 提取验证者收益

```bash
./build/congress-cli withdraw_profits -l http://localhost:8545 -c 202599 \
  -a 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
```

注意：提取收益有最小等待块数限制。

## 完整工作流程示例

以下是一个完整的提案创建和投票流程：

### 1. 查询当前验证者状态

```bash
./build/congress-cli miners -l http://localhost:8545 -c 202599
```

### 2. 创建配置更新提案

```bash
./build/congress-cli create_config_proposal -l http://localhost:8545 -c 202599 \
  -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  -i 0 \
  -v 1000
```

### 3. 签名交易

```bash
./build/congress-cli sign \
  -f createUpdateConfigProposal.json \
  -k /Users/enty/ju-chain-work/chain/private-chain/data/keystore/UTC--2025-08-06T08-30-51.139143000Z--f39fd6e51aad88f6f4ce6ab8827279cfffb92266 \
  -p /Users/enty/ju-chain-work/chain/private-chain/data/password.txt \
  -c 202599
```

### 4. 发送交易

```bash
./build/congress-cli send -f createUpdateConfigProposal_signed.json -l http://localhost:8545
```

输出示例：

```
Wait for tx to be finished executing with hash 0x484662b140a0e98ffd629cee763e12c5f79e7dfd312adbe8cd53b49a99e89c06
tx confirmed in block 24805
read sender from signed tx is 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
--------CreateConfigProposal----------
Proposal ID: 4881cb992217b8d050a87384de6998efb1f051c0cb62c10bdd393f68336618b6
Proposer: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
CID: 0
NewValue: 1000
Time: 1754905540
Block: 24805
-----
send tx success!
```

### 5. 对提案投票

```bash
./build/congress-cli vote_proposal -l http://localhost:8545 -c 202599 \
  -s 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  -i 4881cb992217b8d050a87384de6998efb1f051c0cb62c10bdd393f68336618b6 \
  -a
```

### 6. 签名并发送投票交易

```bash
./build/congress-cli sign \
  -f voteProposal.json \
  -k /Users/enty/ju-chain-work/chain/private-chain/data/keystore/UTC--2025-08-06T08-30-51.139143000Z--f39fd6e51aad88f6f4ce6ab8827279cfffb92266 \
  -p /Users/enty/ju-chain-work/chain/private-chain/data/password.txt \
  -c 202599

./build/congress-cli send -f voteProposal_signed.json -l http://localhost:8545
```

## 配置文件

### 生成的交易文件

工具会在当前目录生成以下JSON文件：

- `createProposal.json` - 创建提案的原始交易
- `createProposal_signed.json` - 已签名的创建提案交易
- `createUpdateConfigProposal.json` - 创建配置更新提案的原始交易
- `createUpdateConfigProposal_signed.json` - 已签名的配置更新提案交易
- `voteProposal.json` - 投票的原始交易
- `voteProposal_signed.json` - 已签名的投票交易

### Keystore 文件格式

使用标准的 Ethereum keystore 格式，例如：

```
UTC--2025-08-06T08-30-51.139143000Z--f39fd6e51aad88f6f4ce6ab8827279cffFb92266
```

## 故障排除

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

**3. 连接RPC失败**
确保：

- RPC端点地址正确
- 区块链节点正在运行
- 网络连接正常

### 调试技巧

1. 使用 `--help` 参数查看命令详细用法
2. 检查生成的JSON文件内容
3. 验证keystore文件路径和密码文件
4. 确认链ID和RPC地址配置正确

## 技术架构

### 项目结构

```
congress-cli/
├── cmd/                    # 命令实现
│   ├── proposal.go        # 提案相关命令
│   ├── tools.go          # 工具函数
│   └── validator.go      # 验证者相关命令
├── contracts/            # 合约绑定（符号链接）
│   └── generated/       # 自动生成的Go绑定
├── build/               # 编译输出
│   └── congress-cli    # 可执行文件
├── Makefile            # 构建配置
├── go.mod              # Go模块定义
└── README.md           # 本文档
```

### 依赖项

- `github.com/ethereum/go-ethereum` - 以太坊客户端库
- `github.com/spf13/cobra` - CLI框架
- `golang.org/x/crypto` - 加密库

## 贡献指南

1. Fork 项目
2. 创建功能分支
3. 提交更改
4. 创建 Pull Request

## 许可证

本项目采用 MIT 许可证。

## 版本历史

- v1.0.0 - 初始版本，支持基本的治理功能
