# Congress CLI v1.1.0

Juchain 区块链治理命令行工具，用于验证者管理、提案投票和网络治理。

## 🚀 新功能 (v1.1.0)

- ✅ **改进的输入验证**：完整的参数验证和错误提示
- ✅ **更好的错误处理**：结构化错误消息和详细错误信息
- ✅ **增强的用户体验**：彩色输出和清晰的状态指示
- ✅ **全局参数验证**：自动验证 RPC 地址和链 ID
- ✅ **配置集中管理**：统一的常量和配置管理
- ✅ **示例命令**：内置使用示例和帮助文档
- ✅ **改进的投票系统**：简化的投票语法
- ✅ **增强的版本信息**：详细的构建和版本信息

## 功能概述

Congress CLI 是一个用于 Juchain 区块链治理的命令行工具，提供了完整的验证者管理和提案投票功能。

### 核心功能

- **提案管理**：创建验证者添加/移除提案和配置更新提案
- **投票系统**：对提案进行投票表决（支持简化的投票语法）
- **验证者管理**：查询验证者信息和管理收益
- **交易处理**：签名和发送交易到区块链网络
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

### 快速开始

1. **查看帮助和示例**：

```bash
./build/congress-cli --help
./build/congress-cli examples
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

### 命令详解

#### 1. 查询验证者信息

**查询所有验证者：**

```bash
# 测试网
./build/congress-cli miners -c 202599 -l https://testnet-rpc.juchain.org

# 主网
./build/congress-cli miners -c 210000 -l https://rpc.juchain.org
```

**查询特定验证者：**

```bash
# 测试网示例
./build/congress-cli miner -c 202599 -l https://testnet-rpc.juchain.org \
  -a 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b

# 主网示例  
./build/congress-cli miner -c 210000 -l https://rpc.juchain.org \
  -a 0x311B37f01c04B84d1f94645BfBd58D82fc03F709
```

输出示例：

```text
Address: 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b
Fee Address: 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b
Status: 1
Accumulated Rewards: 5035784561530401884829
Penalized Rewards: 5323260025816819260865
Last Withdraw Block: 1206974
```

**状态说明：**
- Status 1 = Active (活跃)
- Status 2 = Inactive (异常)

#### 2. 创建提案

**创建验证者添加/移除提案：**

```bash
# 添加验证者 (测试网示例)
./build/congress-cli create_proposal -c 202599 -l https://testnet-rpc.juchain.org \
  -p 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  -t 0x029DAB47e268575D4AC167De64052FB228B5fA41 \
  -o add

# 移除验证者 (主网示例)
./build/congress-cli create_proposal -c 210000 -l https://rpc.juchain.org \
  -p 0xccafa71c31bc11ba24d526fd27ba57d743152807 \
  -t 0x4d432df142823ca25b21bc3f9744ed21a275bdea \
  -o remove
```

参数说明：

- `-p, --proposer` - 提案者地址（必须是有效验证者）
- `-t, --target` - 目标地址（要添加或移除的验证者）
- `-o, --operation` - 操作类型（add 或 remove）

**创建配置更新提案：**

```bash
# 测试网示例：修改 proposalLastingPeriod 为 86400
./build/congress-cli create_config_proposal -c 202599 -l https://testnet-rpc.juchain.org \
  -p 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  -i 0 \
  -v 86400

# 主网示例
./build/congress-cli create_config_proposal -c 210000 -l https://rpc.juchain.org \
  -p 0xccafa71c31bc11ba24d526fd27ba57d743152807 \
  -i 0 \
  -v 86400
```

参数说明：

- `-p, --proposer` - 提案者地址
- `-i, --cid` - 配置ID：
  - 0: proposalLastingPeriod（提案持续期）
  - 1: punishThreshold（惩罚阈值）
  - 2: removeThreshold（移除阈值）
  - 3: decreaseRate（减少率）
  - 4: withdrawProfitPeriod（提取收益周期）
- `-v, --value` - 新的配置值

#### 3. 投票提案

⚠️ **重要**：投票语法已优化！使用 `-a` 标志表示赞成，省略表示反对。

```bash
# 赞成票 (测试网示例)
./build/congress-cli vote_proposal -c 202599 -l https://testnet-rpc.juchain.org \
  -s 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  -i b2be7f3cc702c7a24962df6aed188edbcfdebe20fd55f5670efaedace0e4bcdf \
  -a

# 反对票 (省略 -a 参数)
./build/congress-cli vote_proposal -c 202599 -l https://testnet-rpc.juchain.org \
  -s 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  -i b2be7f3cc702c7a24962df6aed188edbcfdebe20fd55f5670efaedace0e4bcdf

# 主网示例 (赞成票)
./build/congress-cli vote_proposal -c 210000 -l https://rpc.juchain.org \
  -s 0xccafa71c31bc11ba24d526fd27ba57d743152807 \
  -i PROPOSAL_ID \
  -a
```

参数说明：

- `-s, --signer` - 签名者地址（必须是有效验证者）
- `-i, --proposalId` - 提案ID（从创建提案的输出中获取，64位十六进制字符串）
- `-a, --approve` - 赞成票标志（使用 `-a` 表示赞成，省略表示反对）

#### 4. 签名和发送交易

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

#### 5. 提取验证者收益

```bash
# 测试网示例
./build/congress-cli withdraw_profits -c 202599 -l https://testnet-rpc.juchain.org \
  -a 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b

# 主网示例
./build/congress-cli withdraw_profits -c 210000 -l https://rpc.juchain.org \
  -a 0xccafa71c31bc11ba24d526fd27ba57d743152807
```

⚠️ **注意**：
- 提取收益有最小等待块数限制
- 收益提取不需要投票流程，验证者可以直接提取自己的收益
- 需要等待足够的区块数才能提取（withdrawProfitPeriod 配置项控制）

## 完整工作流程示例

以下是一个完整的提案创建和投票流程（测试网示例）：

### 1. 查询当前验证者状态

```bash
./build/congress-cli miners -c 202599 -l https://testnet-rpc.juchain.org
```

### 2. 创建验证者添加提案

```bash
./build/congress-cli create_proposal -c 202599 -l https://testnet-rpc.juchain.org \
  -p 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  -t 0x029DAB47e268575D4AC167De64052FB228B5fA41 \
  -o add
```

### 3. 签名交易

```bash
./build/congress-cli sign \
  -f createProposal.json \
  -k miner1.key \
  -p password.file \
  -c 202599
```

### 4. 发送交易

```bash
./build/congress-cli send -f createProposal_signed.json -l https://testnet-rpc.juchain.org
```

**输出示例（记录提案ID）：**

```text
✅ Transaction broadcast successfully!
ℹ️  Transaction hash: 0x484662b140a0e98ffd629cee763e12c5f79e7dfd312adbe8cd53b49a99e89c06
✅ Transaction confirmed in block 24805
--------CreateProposal----------
Proposal ID: b2be7f3cc702c7a24962df6aed188edbcfdebe20fd55f5670efaedace0e4bcdf
Proposer: 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b
Destination: 0x029DAB47e268575D4AC167De64052FB228B5fA41
Flag: true
Time: 1754905540
Block: 24805
-----
```

### 5. 多个验证者投票

使用上面获取的提案ID进行投票：

```bash
# miner1 赞成票
./build/congress-cli vote_proposal -c 202599 -l https://testnet-rpc.juchain.org \
  -s 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  -i b2be7f3cc702c7a24962df6aed188edbcfdebe20fd55f5670efaedace0e4bcdf \
  -a
./build/congress-cli sign -f voteProposal.json -k miner1.key -p password.file -c 202599
./build/congress-cli send -f voteProposal_signed.json -l https://testnet-rpc.juchain.org

# miner2 赞成票
./build/congress-cli vote_proposal -c 202599 -l https://testnet-rpc.juchain.org \
  -s 0x81f7a79a51edba249efa812eb2d5478f696f7558 \
  -i b2be7f3cc702c7a24962df6aed188edbcfdebe20fd55f5670efaedace0e4bcdf \
  -a
./build/congress-cli sign -f voteProposal.json -k miner2.key -p password.file -c 202599
./build/congress-cli send -f voteProposal_signed.json -l https://testnet-rpc.juchain.org

# miner3 赞成票
./build/congress-cli vote_proposal -c 202599 -l https://testnet-rpc.juchain.org \
  -s 0x578c39eaf09a4e1abf428c423970b59bb8baf42e \
  -i b2be7f3cc702c7a24962df6aed188edbcfdebe20fd55f5670efaedace0e4bcdf \
  -a
./build/congress-cli sign -f voteProposal.json -k miner3.key -p password.file -c 202599
./build/congress-cli send -f voteProposal_signed.json -l https://testnet-rpc.juchain.org
```

### 6. 验证结果

```bash
# 查看新验证者信息
./build/congress-cli miner -c 202599 -l https://testnet-rpc.juchain.org \
  -a 0x029DAB47e268575D4AC167De64052FB228B5fA41

# 查看所有验证者
./build/congress-cli miners -c 202599 -l https://testnet-rpc.juchain.org
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
