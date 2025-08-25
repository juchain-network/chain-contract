# Congress POSA 共识管理工具使用指南

## 测试环境账户配置

为了演示完整的投票流程，本文档使用以下预配置的测试账户：

```bash
# 验证者1账户配置
VALIDATOR1_ADDRESS=0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
VALIDATOR1_PRIVATE_KEY=ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
VALIDATOR1_PASSWORD=123456

# 验证者2账户配置
VALIDATOR2_ADDRESS=0x70997970C51812dc3A010C7d01b50e0d17dc79C8
VALIDATOR2_PRIVATE_KEY=59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d
VALIDATOR2_PASSWORD=123456

# 验证者3账户配置
VALIDATOR3_ADDRESS=0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC
VALIDATOR3_PRIVATE_KEY=5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a
VALIDATOR3_PASSWORD=123456

# 验证者4账户配置
VALIDATOR4_ADDRESS=0x90F79bf6EB2c4f870365E785982E1f101E93b906
VALIDATOR4_PRIVATE_KEY=7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6
VALIDATOR4_PASSWORD=123456

# 验证者5账户配置
VALIDATOR5_ADDRESS=0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65
VALIDATOR5_PRIVATE_KEY=47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a
VALIDATOR5_PASSWORD=123456

# 验证者6账户配置（待添加为新验证者）
VALIDATOR6_ADDRESS=0x340D92a853Ae20A6E7a5b86272fa47AFf83a8F7A
VALIDATOR6_PRIVATE_KEY=8b3a350cf5c34c9194ca85829a2df0ec3153be0318b5e2d3348e872092edffba
VALIDATOR6_PASSWORD=123456
```

> **⚠️ 重要**: 以上私钥仅用于测试环境，切勿在生产环境使用！

## 工具编译

首先需要编译 congress-cli 工具：

```shell
cd sys-contract/congress-cli
make build
# 生成的可执行文件位于 build/congress-cli
```

## 版本信息

本文档已更新至 **Congress CLI v1.2.0**，构建日期：2025-08-25。

```shell
./build/congress-cli version
# 输出: Congress CLI Version: 1.2.0, Build Date: 2025-08-25
```

> **💡 语法说明**：
>
> - 本地测试：默认连接 `http://127.0.0.1:8545`，链ID `202599`
> - 测试网：`--chainId 202599 --rpc_laddr https://testnet-rpc.juchain.org`
> - 主网：`--chainId 210000 --rpc_laddr https://rpc.juchain.org`

## 配置管理命令

### 查看配置帮助

Congress CLI v1.2.0 提供了便捷的配置查询和设置功能。

```shell
# 查看config子命令帮助
./build/congress-cli config --help
```

### 查询系统配置

查询当前系统配置参数：

```shell
# 查询所有配置
./build/congress-cli config get

# 查询RPC端点
./build/congress-cli config get --rpc

# 查询链ID
./build/congress-cli config get --chain-id
```

**示例**：

```shell
# 查询所有配置
./build/congress-cli config get
# 输出：
# RPC endpoint: https://rpc.juchain.org
# Chain ID: 210000

# 只查询RPC端点
./build/congress-cli config get --rpc
# 输出：RPC endpoint: https://rpc.juchain.org

# 只查询链ID
./build/congress-cli config get --chain-id
# 输出：Chain ID: 210000
```

### 修改系统配置

设置RPC端点和链ID配置：

```shell
# 设置RPC端点
./build/congress-cli config set --rpc <RPC_URL>

# 设置链ID
./build/congress-cli config set --chain-id <CHAIN_ID>

# 同时设置RPC和链ID
./build/congress-cli config set --rpc <RPC_URL> --chain-id <CHAIN_ID>
```

**示例**：

```shell
# 设置本地测试环境
./build/congress-cli config set --rpc http://127.0.0.1:8545 --chain-id 202599

# 设置测试网环境
./build/congress-cli config set --rpc https://testnet-rpc.juchain.org --chain-id 202599

# 设置主网环境
./build/congress-cli config set --rpc https://rpc.juchain.org --chain-id 210000

# 只设置RPC端点
./build/congress-cli config set --rpc https://rpc.juchain.org

# 只设置链ID
./build/congress-cli config set --chain-id 210000
```

> **说明**: 这里的config命令是用于设置congress-cli工具本身的配置（如RPC端点、链ID），不是修改区块链系统参数。区块链系统参数修改请参考第4章。

## 1. 创建提案

### 1.1. 创建原始交易

```shell
# 基本语法
./build/congress-cli create_proposal -p 提案矿工地址 -t 新矿地址 -o add

# 完整示例：验证者1为验证者6创建添加提案
./build/congress-cli create_proposal -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 -t 0x340D92a853Ae20A6E7a5b86272fa47AFf83a8F7A -o add

# 其他示例：移除验证者
./build/congress-cli create_proposal -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 -t 0x340D92a853Ae20A6E7a5b86272fa47AFf83a8F7A -o remove
```

**参数说明**：

- `-p, --proposer`: 提案者地址（必须是有效验证者）
- `-t, --target`: 目标地址（要添加或移除的验证者）
- `-o, --operation`: 操作类型（add 或 remove）

> 执行成功后会生成 `createProposal.json` 文件

### 1.2. 签名交易

```shell
./build/congress-cli sign -f createProposal.json -k 钱包文件 -p 钱包密码文件

# 验证者1签名示例
./build/congress-cli sign -f createProposal.json -k validator1.key -p password.file
```

> 执行成功后会生成 `createProposal_signed.json` 文件

### 1.3. 发送交易

```shell
./build/congress-cli send -f createProposal_signed.json
```

> 执行成功后会输出提案信息，包含重要的提案ID：

```text
✅ Transaction confirmed in block 8758
read sender from signed tx is 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
--------CreateProposal----------
Proposal ID: 0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c
Proposer: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
Destination: 0x340D92a853Ae20A6E7a5b86272fa47AFf83a8F7A
Flag: true
Time: 1756110615
Block: 8758
-----
✅ Transaction broadcast successfully!
```

> **⚠️ 重要**: 记录提案ID，投票时需要使用！

## 2. 提案投票

### 2.1. 创建投票交易

现在多个验证者需要对提案进行投票。假设提案ID为 `0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c`

```shell
# 基本语法
./build/congress-cli vote_proposal -s 签名矿工地址 -i 提案ID -a  # 赞成票
./build/congress-cli vote_proposal -s 签名矿工地址 -i 提案ID     # 反对票

# 验证者1投票（提案者自己也要投票）
./build/congress-cli vote_proposal -s 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 -i 0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c -a

# 验证者2投票
./build/congress-cli vote_proposal -s 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 -i 0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c -a

# 验证者3投票
./build/congress-cli vote_proposal -s 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC -i 0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c -a

# 验证者4投票
./build/congress-cli vote_proposal -s 0x90F79bf6EB2c4f870365E785982E1f101E93b906 -i 0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c -a

# 验证者5投票
./build/congress-cli vote_proposal -s 0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65 -i 0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c -a
```

**参数说明**：

- `-s, --signer`: 签名者地址（必须是有效验证者）
- `-i, --proposalId`: 提案ID（64字符的十六进制字符串）
- `-a, --approve`: 赞成票标志（使用-a表示YES，省略表示NO）

> 执行成功后会生成 `voteProposal.json` 文件

### 2.2. 签名交易

每次投票都需要对应验证者的私钥进行签名：

```shell
# 验证者1签名
./build/congress-cli sign -f voteProposal.json -k validator1.key -p password.file

# 验证者2签名  
./build/congress-cli sign -f voteProposal.json -k validator2.key -p password.file

# 验证者3签名
./build/congress-cli sign -f voteProposal.json -k validator3.key -p password.file

# 验证者4签名
./build/congress-cli sign -f voteProposal.json -k validator4.key -p password.file

# 验证者5签名
./build/congress-cli sign -f voteProposal.json -k validator5.key -p password.file
```

> 执行成功后会生成 `voteProposal_signed.json` 文件

### 2.3. 发送投票交易

每个验证者都需要发送自己的投票交易：

```shell
# 依次发送每个验证者的投票交易
./build/congress-cli send -f voteProposal_signed.json
```

> 执行成功后会输出确认信息：

```text
✅ Transaction confirmed in block 8830
read sender from signed tx is 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
✅ Transaction broadcast successfully!
```

### 2.4. 完整投票流程示例

以下是为验证者6添加为新验证者的完整投票流程：

```shell
# 步骤1: 验证者1投票
./build/congress-cli vote_proposal -s 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 -i 提案ID -a
./build/congress-cli sign -f voteProposal.json -k validator1.key -p password.file
./build/congress-cli send -f voteProposal_signed.json

# 步骤2: 验证者2投票
./build/congress-cli vote_proposal -s 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 -i 提案ID -a
./build/congress-cli sign -f voteProposal.json -k validator2.key -p password.file
./build/congress-cli send -f voteProposal_signed.json

# 步骤3: 验证者3投票
./build/congress-cli vote_proposal -s 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC -i 提案ID -a
./build/congress-cli sign -f voteProposal.json -k validator3.key -p password.file
./build/congress-cli send -f voteProposal_signed.json

# 步骤4: 验证者4投票
./build/congress-cli vote_proposal -s 0x90F79bf6EB2c4f870365E785982E1f101E93b906 -i 提案ID -a
./build/congress-cli sign -f voteProposal.json -k validator4.key -p password.file
./build/congress-cli send -f voteProposal_signed.json

# 步骤5: 验证者5投票
./build/congress-cli vote_proposal -s 0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65 -i 提案ID -a
./build/congress-cli sign -f voteProposal.json -k validator5.key -p password.file
./build/congress-cli send -f voteProposal_signed.json
```

> **注意**:
>
> - 每个验证者只能对同一提案投票一次
> - 需要足够多的验证者投赞成票，提案才能通过
> - 请将上述命令中的"提案ID"替换为实际的提案ID

## 3. 查询操作

### 3.1 查询所有活动矿工

```shell
./build/congress-cli miners
```

> 输出示例：

```text
ℹ️  Fetching validator information...
ℹ️  Found 5 validators:

--- Validator 1 ---
Address: 0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65
Fee Address: 0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65
Status: Active ✅
Accumulated Rewards: 54174000000000
Penalized Rewards: 0
Last Withdraw Block: 0

--- Validator 2 ---
Address: 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC
Fee Address: 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC
Status: Active ✅
Accumulated Rewards: 54174000000000
Penalized Rewards: 0
Last Withdraw Block: 0
...
```

### 3.2 查询单个矿工

```shell
./build/congress-cli miner -a <验证者地址>

# 示例
./build/congress-cli miner -a 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
```

> 输出示例：

```text
ℹ️  Querying validator information for: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
Address: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
Fee Address: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
Status: Active ✅
Accumulated Rewards: 54174000000000
Penalized Rewards: 0
Last Withdraw Block: 0
```

**状态说明：**

- Status: Active ✅ = 活跃验证者
- Status: Inactive ❌ = 异常状态

### 3.3 查询所有提案

```shell
./build/congress-cli proposals
```

> 输出示例：

```text
ℹ️  Fetching all proposals...
ℹ️  Found 1 proposal(s):

--- Proposal 1 ---
ID: 0x0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c
Subject: Test Proposal Subject
Content: Test Proposal Content
Type: 4
Status: Voting
Block Number: 8829
Content Hash: 0xed2e9ba8a0b3ca2b9b7a2c4b8f9a7b5c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0
Contract Address: 0x0000000000000000000000000000000000001000
Current: 1
Start Time: 2025-01-22 20:20:21 +0000 UTC
End Time: 2025-01-23 20:20:21 +0000 UTC
```

### 3.4 查询单个提案

```shell
./build/congress-cli proposal -i <提案ID>

# 示例
./build/congress-cli proposal -i 0x0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c
```

> 输出示例：

```text
ℹ️  Fetching proposal details...
ID: 0x0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c
Subject: Test Proposal Subject
Content: Test Proposal Content
Type: 4
Status: Voting
Block Number: 8829
Content Hash: 0xed2e9ba8a0b3ca2b9b7a2c4b8f9a7b5c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0
Contract Address: 0x0000000000000000000000000000000000001000
Current: 1
Start Time: 2025-01-22 20:20:21 +0000 UTC
End Time: 2025-01-23 20:20:21 +0000 UTC
```

**提案状态说明：**

- Status: Voting = 投票中
- Status: Passed = 已通过
- Status: Failed = 已拒绝
- Status: Executed = 已执行

## 4. 修改参数配置

### 4.1 创建配置修改提案

```shell
./build/congress-cli create_config_proposal -p <提案者地址> -i <配置项ID> -v <配置值>

# 示例：修改提案持续时间为86400秒
./build/congress-cli create_config_proposal -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 -i 0 -v 86400
```

**配置项ID对应表：**

- 0: proposalLastingPeriod（提案持续时间）
- 1: punishThreshold（惩罚阈值）
- 2: removeThreshold（移除阈值）
- 3: decreaseRate（减少率）
- 4: withdrawProfitPeriod（提取收益周期）

**参数说明：**

- `-p, --proposer`: 提案者地址（必须是有效验证者）
- `-i, --configId`: 配置项ID（0-4）
- `-v, --value`: 配置项的新值

> 执行成功后会生成 `createConfigProposal.json` 文件

### 4.2 签名和发送交易

配置修改提案的签名和发送流程与普通提案相同：

```shell
# 签名交易
./build/congress-cli sign -f createConfigProposal.json -k miner1.key -p password.file

# 发送交易
./build/congress-cli send -f createConfigProposal_signed.json
```

> **注意：** 配置修改提案同样需要足够多的验证者投票才能通过并执行

## 5. 矿工收益提取

### 5.1 创建提取交易

```shell
./build/congress-cli withdraw_profits -a <矿工地址>

# 示例
./build/congress-cli withdraw_profits -a 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
```

**参数说明：**

- `-a, --address`: 要提取收益的矿工地址

### 5.2 签名和发送交易

```shell
# 签名交易
./build/congress-cli sign -f withdrawProfits.json -k miner1.key -p password.file

# 发送交易
./build/congress-cli send -f withdrawProfits_signed.json
```

> **注意：** 收益提取不需要投票流程，矿工可以直接提取自己的收益

## 6. Staking 操作

### 6.1 Staking 命令概览

Congress CLI v1.2.0 新增了完整的staking功能模块，支持验证者注册、委托、查询等操作。

```shell
# 查看staking子命令帮助
./build/congress-cli staking --help
```

**可用的Staking子命令：**

1. `register-validator` - 注册验证者
2. `delegate` - 委托质押给验证者
3. `undelegate` - 取消委托
4. `query-validator` - 查询指定验证者信息
5. `list-top-validators` - 列出顶级验证者
6. `unjail` - 解除验证者监禁状态
7. `withdraw` - 提取委托收益

### 6.2 注册验证者

注册成为新的验证者需要满足最低质押要求。

```shell
./build/congress-cli staking register-validator \
  --proposer <提案者地址> \
  --stake-amount <质押数量> \
  --commission-rate <佣金率>
```

**参数说明**：

- `--proposer`: 提案者地址（必需）
- `--stake-amount`: 质押的JU数量（必需，最少10000 JU）
- `--commission-rate`: 佣金率基点（0-10000，例如500表示5%）

**示例**：

```shell
# 本地测试环境
./build/congress-cli staking register-validator \
  --proposer 0x3F9DDeBE20b24B0DEC1d0B5A3c6e8Cb8D3eCF6A5 \
  --stake-amount 10000 \
  --commission-rate 500

# 主网环境
./build/congress-cli staking register-validator \
  --proposer 0x3F9DDeBE20b24B0DEC1d0B5A3c6e8Cb8D3eCF6A5 \
  --stake-amount 10000 \
  --commission-rate 500 \
  --rpc https://rpc.juchain.org --chainId 210000
```

### 6.3 委托质押

将JU代币委托给验证者以获得奖励。

```shell
./build/congress-cli staking delegate \
  --validator <验证者地址> \
  --amount <委托数量>
```

**示例**：

```shell
# 委托1000 JU给验证者
./build/congress-cli staking delegate \
  --validator 0x3F9DDeBE20b24B0DEC1d0B5A3c6e8Cb8D3eCF6A5 \
  --amount 1000
```

### 6.4 取消委托

取消之前的委托，需要等待解绑期。

```shell
./build/congress-cli staking undelegate \
  --validator <验证者地址> \
  --amount <取消委托数量>
```

**示例**：

```shell
# 取消委托500 JU
./build/congress-cli staking undelegate \
  --validator 0x3F9DDeBE20b24B0DEC1d0B5A3c6e8Cb8D3eCF6A5 \
  --amount 500
```

### 6.5 查询验证者信息

查询指定验证者的详细信息，包括质押量、佣金率、状态等。

```shell
./build/congress-cli staking query-validator --address <验证者地址>
```

**示例**：

```shell
# 查询验证者信息
./build/congress-cli staking query-validator \
  --address 0x3F9DDeBE20b24B0DEC1d0B5A3c6e8Cb8D3eCF6A5
```

> **注意**: 当前版本某些查询命令可能遇到ABI解析错误，这是已知问题，正在修复中。

### 6.6 列出顶级验证者

查看当前活跃的顶级验证者列表。

```shell
./build/congress-cli staking list-top-validators [数量]
```

**示例**：

```shell
# 查看前15名验证者（默认）
./build/congress-cli staking list-top-validators

# 查看前10名验证者
./build/congress-cli staking list-top-validators 10
```

### 6.7 解除监禁

验证者因违规被监禁后，可以申请解除监禁状态。

```shell
./build/congress-cli staking unjail --validator <验证者地址>
```

### 6.8 提取收益

提取委托产生的收益。

```shell
./build/congress-cli staking withdraw --validator <验证者地址>
```

## 7. 完整端到端流程：添加验证者6

本章节演示从头开始的完整流程，让验证者1-5为验证者6投票，使其成为新的验证者。

### 7.1 前置条件

确保以下账户已准备好对应的密钥文件：

- `validator1.key` - 验证者1的私钥文件
- `validator2.key` - 验证者2的私钥文件  
- `validator3.key` - 验证者3的私钥文件
- `validator4.key` - 验证者4的私钥文件
- `validator5.key` - 验证者5的私钥文件
- `password.file` - 包含密码"123456"的文件

### 7.2 第一步：验证者1创建提案

```shell
# 创建添加验证者6的提案
./build/congress-cli create_proposal \
  -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
  -t 0x340D92a853Ae20A6E7a5b86272fa47AFf83a8F7A \
  -o add

# 验证者1签名
./build/congress-cli sign -f createProposal.json -k validator1.key -p password.file

# 发送交易
./build/congress-cli send -f createProposal_signed.json
```

> **重要**: 记录输出中的提案ID，例如：`0943f0c9c31b9042ab6fc0891a216343324ce85e04ee83a9e39352cbedfd7a4c`

### 7.3 第二步：5个验证者投票

将下面命令中的 `PROPOSAL_ID` 替换为第一步获得的实际提案ID。

```shell
# 验证者1投票
./build/congress-cli vote_proposal -s 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k validator1.key -p password.file
./build/congress-cli send -f voteProposal_signed.json

# 验证者2投票
./build/congress-cli vote_proposal -s 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k validator2.key -p password.file
./build/congress-cli send -f voteProposal_signed.json

# 验证者3投票
./build/congress-cli vote_proposal -s 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k validator3.key -p password.file
./build/congress-cli send -f voteProposal_signed.json

# 验证者4投票
./build/congress-cli vote_proposal -s 0x90F79bf6EB2c4f870365E785982E1f101E93b906 -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k validator4.key -p password.file
./build/congress-cli send -f voteProposal_signed.json

# 验证者5投票
./build/congress-cli vote_proposal -s 0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65 -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k validator5.key -p password.file
./build/congress-cli send -f voteProposal_signed.json
```

### 7.4 第三步：验证结果

```shell
# 查询提案状态
./build/congress-cli proposal -i PROPOSAL_ID

# 查询验证者6的状态
./build/congress-cli miner -a 0x340D92a853Ae20A6E7a5b86272fa47AFf83a8F7A

# 查询所有验证者列表
./build/congress-cli miners
```

> **预期结果**:
>
> - 提案状态应显示为"已通过"或"已执行"
> - 验证者6应出现在活跃验证者列表中
> - 验证者总数应该从5个变为6个

### 7.5 自动化脚本示例

您也可以创建一个脚本来自动化整个流程：

```bash
#!/bin/bash

# 设置提案目标地址
TARGET_ADDRESS="0x340D92a853Ae20A6E7a5b86272fa47AFf83a8F7A"
PROPOSER_ADDRESS="0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"

# 验证者地址数组
VALIDATORS=(
    "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
    "0x70997970C51812dc3A010C7d01b50e0d17dc79C8"
    "0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC"
    "0x90F79bf6EB2c4f870365E785982E1f101E93b906"
    "0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65"
)

# 密钥文件数组
KEYS=(
    "validator1.key"
    "validator2.key"
    "validator3.key"
    "validator4.key"
    "validator5.key"
)

echo "=== 步骤1: 创建提案 ==="
./build/congress-cli create_proposal -p $PROPOSER_ADDRESS -t $TARGET_ADDRESS -o add
./build/congress-cli sign -f createProposal.json -k validator1.key -p password.file
./build/congress-cli send -f createProposal_signed.json

echo "请输入提案ID:"
read PROPOSAL_ID

echo "=== 步骤2: 验证者投票 ==="
for i in "${!VALIDATORS[@]}"; do
    echo "验证者 $((i+1)) 投票中..."
    ./build/congress-cli vote_proposal -s ${VALIDATORS[$i]} -i $PROPOSAL_ID -a
    ./build/congress-cli sign -f voteProposal.json -k ${KEYS[$i]} -p password.file
    ./build/congress-cli send -f voteProposal_signed.json
    echo "验证者 $((i+1)) 投票完成"
done

echo "=== 步骤3: 验证结果 ==="
./build/congress-cli proposal -i $PROPOSAL_ID
./build/congress-cli miner -a $TARGET_ADDRESS
./build/congress-cli miners
```

## 8. 完整测试交易示例

> **以下是一个完整的测试流程示例：**

```shell
# 1. 创建提案
./build/congress-cli create_proposal 
  --proposal "test proposal" 
  --action 0 
  --value 1000 
  --proposer 0x3F9DDeBE20b24B0DEC1d0B5A3c6e8Cb8D3eCF6A5

# 2. 投票（4个验证者投票同意）
./build/congress-cli vote_proposal --proposer 0x3F9DDeBE20b24B0DEC1d0B5A3c6e8Cb8D3eCF6A5 --id 1 --vote 1
./build/congress-cli vote_proposal --proposer 0x4B1E2D4D7C8F5A9B6E3F8A2C7D9E1F4B8C5E6A9 --id 1 --vote 1  
./build/congress-cli vote_proposal --proposer 0x5C2F3E5E8D9A6B7C4E5F9B3D8E2A5C9D6F7A8E --id 1 --vote 1
./build/congress-cli vote_proposal --proposer 0x6D3A4F6F9E8B7C5D6A7C8F4E9A3B6D7E8F9C1A --id 1 --vote 1

# 3. 查询提案状态
./build/congress-cli query_proposal --id 1

# 4. 委托质押
./build/congress-cli staking delegate 
  --validator 0x3F9DDeBE20b24B0DEC1d0B5A3c6e8Cb8D3eCF6A5 
  --amount 1000
```

## 9. 主网恢复矿工身份操作

> miner1 创建提案，新增 0x029DAB47e268575D4AC167De64052FB228B5fA41 作为新的矿工，创建完提案后，miner1,miner2,miner3 投票通过

```shell
# step1 创建提案交易，并签名发送
./build/congress-cli create_proposal -p 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b -t 0x029DAB47e268575D4AC167De64052FB228B5fA41 -o add
./build/congress-cli sign -f createProposal.json -k miner1.key -p password.file
./build/congress-cli send -f createProposal_signed.json
# 这条命令执行后可以获取到提案ID，例如：b2be7f3cc702c7a24962df6aed188edbcfdebe20fd55f5670efaedace0e4bcdf

# step2 3个矿工对提案进行投票（请将 PROPOSAL_ID 替换为上一步获得的实际提案ID）
# miner1
./build/congress-cli vote_proposal -s 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k miner1.key -p password.file
./build/congress-cli send -f voteProposal_signed.json

# miner2
./build/congress-cli vote_proposal -s 0x81f7a79a51edba249efa812eb2d5478f696f7558 -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k miner2.key -p password.file
./build/congress-cli send -f voteProposal_signed.json

# miner3
./build/congress-cli vote_proposal -s 0x578c39eaf09a4e1abf428c423970b59bb8baf42e -i PROPOSAL_ID -a
./build/congress-cli sign -f voteProposal.json -k miner3.key -p password.file
./build/congress-cli send -f voteProposal_signed.json

# step3 查看新增矿工的信息
./build/congress-cli miner -a 0x029DAB47e268575D4AC167De64052FB228B5fA41
./build/congress-cli miners
```

## 10. 主网修改配置

### 10.1 创建提案

```shell
# 配置项ID对应的配置项信息
# 0 proposalLastingPeriod, 1 punishThreshold, 2 removeThreshold, 3 decreaseRate, 4 withdrawProfitPeriod
./build/congress-cli create_config_proposal -p 提案矿工地址 -i 配置项ID -v 配置项取值

# 示例：修改 proposalLastingPeriod 为 86400 秒（注意：使用 -i 参数，不是 -c）
./build/congress-cli create_config_proposal -p 0xccafa71c31bc11ba24d526fd27ba57d743152807 -i 0 -v 86400

./build/congress-cli sign -f createUpdateConfigProposal.json -k miner1.key -p password.file

./build/congress-cli send -f createUpdateConfigProposal_signed.json
# 这条命令执行后可以获取到提案ID，记录提案ID用于后续投票
```

### 10.2 验证者投票

> 投票步骤与主网恢复矿工操作相同，请参考上面第9章的相关步骤

## 11. 主网矿工收益提取

```shell
# step1 创建原始交易
./build/congress-cli withdraw_profits -a 矿工地址

# step2 交易签名
./build/congress-cli sign -f withdrawProfits.json -k miner.key -p password.file

# step3 发送交易
./build/congress-cli send -f withdrawProfits_signed.json
```

## 12. 工具信息

### 12.1 版本查看

```shell
./build/congress-cli version
```

### 12.2 帮助信息

```shell
./build/congress-cli help
./build/congress-cli [command] --help  # 查看特定命令的帮助
```

### 12.3 测试脚本

项目包含测试脚本，可以快速验证系统状态：

```shell
cd sys-contract/congress-cli
chmod +x test_congress.sh
./test_congress.sh
```

## 13. 注意事项

### 13.1 重要提醒

- ⚠️ **验证者要求**: 只有当前有效的验证者才能创建提案和投票
- ⚠️ **网络同步**: 在恢复矿工身份前，确保节点已完全同步到最新状态
- ⚠️ **提案ID**: 每次操作都会生成新的提案ID，务必使用正确的ID
- ⚠️ **密钥安全**: 妥善保管密钥文件和密码文件

### 13.2 常见错误

1. **"Validator only"**: 当前账户不是有效验证者
2. **"You can't vote for a proposal twice"**: 该验证者已对此提案投过票
3. **"gas estimation failed"**: 交易参数错误或网络问题

### 13.3 系统合约地址

- **Validators**: `0x000000000000000000000000000000000000f000`
- **Punish**: `0x000000000000000000000000000000000000f001`
- **Proposal**: `0x000000000000000000000000000000000000f002`
- **Stake**: `0x000000000000000000000000000000000000f003`

### 13.4 网络信息

- **测试网**: `https://testnet-rpc.juchain.org` (Chain ID: 202599)
- **主网**: `https://rpc.juchain.org` (Chain ID: 210000)
