# Congress POA 共识管理工具使用指南

## 工具编译

首先需要编译 congress-cli 工具：

```shell
cd sys-contract/congress-cli
make build
# 生成的可执行文件位于 build/congress-cli
```

## 1.创建提案

```shell
# 测试网通用参数
./build/congress-cli --chainId 202599 --rpc_laddr https://testnet-rpc.juchain.org 

# 主网通用参数
./build/congress-cli --chainId 210000 --rpc_laddr https://rpc.juchain.org 
```

### 1.1. 创建原始交易

```shell
./build/congress-cli create_proposal -p 提案矿工地址 -t 新矿工地址 -o add --rpc_laddr https://rpc.juchain.org 

# 测试样例命令
./build/congress-cli create_proposal -p 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b -t 0x029DAB47e268575D4AC167De64052FB228B5fA41 -o add --rpc_laddr https://testnet-rpc.juchain.org 
```

### 1.2. 签名交易

```shell
./build/congress-cli sign -f createProposal.json -k 钱包文件 -p 钱包密码文件 --chainId 210000 

# 测试样例文件
./build/congress-cli sign -f createProposal.json -k miner1.key -p password.file --chainId 202599 
```

### 1.3. 发送交易

```shell
./build/congress-cli send -f createProposal_signed.json --rpc_laddr https://rpc.juchain.org 

# 测试样例文件
./build/congress-cli send -f createProposal_signed.json --rpc_laddr https://testnet-rpc.juchain.org 
```

> 执行成功后会生成提案信息，如下(后面投票会用到Proposal ID)：

```text
Wait for tx to be finished executing with hash 0xb72b3e4f2f4411fd467dcf3a4af16f12e5772a59ec91535ad18283c9a2e32ddf
tx confirmed in block 12535222
read sender from signed tx is 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
--------CreateProposal----------
Proposal ID: b2be7f3cc702c7a24962df6aed188edbcfdebe20fd55f5670efaedace0e4bcdf
Proposer: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
Destination: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8
Flag: true
Time: 1754909524
Block: 12535222
-----
send tx success!
```

## 2.提案投票

### 2.1. 创建原始交易

```shell
./build/congress-cli vote_proposal -s 签名矿工地址 -i 提案ID -a 是否通过 --rpc_laddr https://rpc.juchain.org 

# 测试样例命令（使用上面生成的提案ID）
./build/congress-cli vote_proposal -s 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b -i b2be7f3cc702c7a24962df6aed188edbcfdebe20fd55f5670efaedace0e4bcdf -a true --rpc_laddr https://testnet-rpc.juchain.org 
```

### 2.2. 签名交易

```shell
./build/congress-cli sign -f voteProposal.json -k 钱包文件 -p 钱包密码文件 --chainId 210000 

# 测试样例文件
./build/congress-cli sign -f voteProposal.json -k miner1.key -p password.file --chainId 202599 
```

### 2.3. 发送交易

```shell
./build/congress-cli send -f voteProposal_signed.json --rpc_laddr https://rpc.juchain.org 

# 测试样例文件
./build/congress-cli send -f voteProposal_signed.json --rpc_laddr https://testnet-rpc.juchain.org 
```

## 3. 查询操作

### 3.1 查询所有活动矿工

```shell
./build/congress-cli miners --rpc_laddr https://rpc.juchain.org 

# 测试网示例
./build/congress-cli miners --rpc_laddr https://testnet-rpc.juchain.org
```

### 3.2 查询单个矿工

```shell
./build/congress-cli miner --rpc_laddr https://rpc.juchain.org -a 0x311B37f01c04B84d1f94645BfBd58D82fc03F709

# 测试网示例
./build/congress-cli miner --rpc_laddr https://testnet-rpc.juchain.org -a 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b
```

> 输出示例：

```text
矿工  0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b 奖励地址 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b 活动状态 1 累计奖励 5035784561530401884829 罚没奖励 5323260025816819260865 
上次提取奖励区块 1206974
```

**状态说明：**

- 活动状态 1 = Active (活跃)
- 活动状态 2 = Inactive (异常)

## 4.修改参数配置

### 4.1. 创建原始交易

```shell
# 配置项ID对应的配置项信息
# 0 proposalLastingPeriod, 1 punishThreshold, 2 removeThreshold, 3 decreaseRate, 4 withdrawProfitPeriod
./build/congress-cli create_config_proposal -p 提案矿工地址 -i 配置项ID -v 配置项取值 --rpc_laddr https://rpc.juchain.org 

# 测试样例命令（注意：使用 -i 参数，不是 -c）
./build/congress-cli create_config_proposal -p 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b -i 0 -v 86400 --rpc_laddr https://testnet-rpc.juchain.org 
```

> 交易的发送和签名，以及后面的投票流程和前面一致，不再重复

## 5.矿工收益提取

### 5.1. 创建原始交易

```shell
./build/congress-cli withdraw_profits -a 矿工地址 --rpc_laddr https://rpc.juchain.org 

# 测试样例命令
./build/congress-cli withdraw_profits -a 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b --rpc_laddr https://testnet-rpc.juchain.org 
```

### 5.2. 签名和发送

```shell
# 签名交易
./build/congress-cli sign -f withdrawProfits.json -k miner1.key -p password.file --chainId 202599

# 发送交易
./build/congress-cli send -f withdrawProfits_signed.json --rpc_laddr https://testnet-rpc.juchain.org
```

> **注意：** 收益提取不需要投票流程，矿工可以直接提取自己的收益
>
## 6.完整测试交易示例

> miner1 创建提案，新增 0x029DAB47e268575D4AC167De64052FB228B5fA41 作为新的矿工，创建完提案后，miner1,miner2,miner3 投票通过

```shell
# step1 创建提案交易，并签名发送
./build/congress-cli create_proposal -p 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b -t 0x029DAB47e268575D4AC167De64052FB228B5fA41 -o add --rpc_laddr https://testnet-rpc.juchain.org 
./build/congress-cli sign -f createProposal.json -k miner1.key -p password.file --chainId 202599 
./build/congress-cli send -f createProposal_signed.json --rpc_laddr https://testnet-rpc.juchain.org 
# 这条命令执行后可以获取到提案ID，例如：b2be7f3cc702c7a24962df6aed188edbcfdebe20fd55f5670efaedace0e4bcdf

# step2 3个矿工对提案进行投票（请将 PROPOSAL_ID 替换为上一步获得的实际提案ID）
# miner1
./build/congress-cli vote_proposal -s 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b -i PROPOSAL_ID -a true --rpc_laddr https://testnet-rpc.juchain.org 
./build/congress-cli sign -f voteProposal.json -k miner1.key -p password.file --chainId 202599 
./build/congress-cli send -f voteProposal_signed.json --rpc_laddr https://testnet-rpc.juchain.org 

# miner2
./build/congress-cli vote_proposal -s 0x81f7a79a51edba249efa812eb2d5478f696f7558 -i PROPOSAL_ID -a true --rpc_laddr https://testnet-rpc.juchain.org 
./build/congress-cli sign -f voteProposal.json -k miner2.key -p password.file --chainId 202599 
./build/congress-cli send -f voteProposal_signed.json --rpc_laddr https://testnet-rpc.juchain.org 

# miner3
./build/congress-cli vote_proposal -s 0x578c39eaf09a4e1abf428c423970b59bb8baf42e -i PROPOSAL_ID -a true --rpc_laddr https://testnet-rpc.juchain.org 
./build/congress-cli sign -f voteProposal.json -k miner3.key -p password.file --chainId 202599 
./build/congress-cli send -f voteProposal_signed.json --rpc_laddr https://testnet-rpc.juchain.org 

# step3 查看新增矿工的信息
./build/congress-cli miner --rpc_laddr https://testnet-rpc.juchain.org -a 0x029DAB47e268575D4AC167De64052FB228B5fA41
./build/congress-cli miners --rpc_laddr https://testnet-rpc.juchain.org
```

## 7. 主网恢复矿工身份操作
>
> 现有5个矿工地址，miner1-miner5地址如下：

```
0xccafa71c31bc11ba24d526fd27ba57d743152807
0xd5da2b33c1f620a94bf2039b9cb540853e7928d7
0x311b37f01c04b84d1f94645bfbd58d82fc03f709
0xde0e48c5337db3ca7b3710c27e9728e68bf220b3
0x4d432df142823ca25b21bc3f9744ed21a275bdea
```

> 其中miner5状态异常，可以通过如下命令查询该矿工状态：

```shell
./build/congress-cli miner --rpc_laddr https://rpc.juchain.org -a 0x4d432df142823ca25b21bc3f9744ed21a275bdea

# 输出信息中"活动状态 2"，表示异常状态，1 为正常
```

### 7.1 创建提案

> 下面的操作通过 miner1 创建提案，重新提案 miner5 作为活动矿工，创建完提案后，miner1,miner2,miner3 投票通过提案，即可让miner5恢复活动状态

> ⚠️ **重要提醒** ⚠️ 在对新矿工投票之前，确保新矿工节点已经同步到最新状态，否则投票通过后没及时出块，会被再次踢出矿工列表！

```shell
# step1 创建提案交易，并签名发送，其中 -p 参数为创建提案的矿工，-t 参数为新增的矿工
./build/congress-cli create_proposal -p 0xccafa71c31bc11ba24d526fd27ba57d743152807 -t 0x4d432df142823ca25b21bc3f9744ed21a275bdea -o add --rpc_laddr https://rpc.juchain.org 

./build/congress-cli sign -f createProposal.json -k miner1.key -p password.file --chainId 210000

./build/congress-cli send -f createProposal_signed.json --rpc_laddr https://rpc.juchain.org 
# 这条命令执行后可以获取到提案ID，记录提案ID用于后续投票
```

### 7.2 验证者投票

```shell
# step2 3个矿工对提案进行投票（注意！！！ 先将命令中的 PROPOSAL_ID 替换为上一步生成的提案ID）
# miner1 投票通过， -s 参数为投票签名矿工, -i参数为上一步生成的提案id，-a true表示通过提案
./build/congress-cli vote_proposal -s 0xccafa71c31bc11ba24d526fd27ba57d743152807 -i PROPOSAL_ID -a true --rpc_laddr https://rpc.juchain.org 

./build/congress-cli sign -f voteProposal.json -k miner1.key -p password.file --chainId 210000 

./build/congress-cli send -f voteProposal_signed.json --rpc_laddr https://rpc.juchain.org 

# miner2
./build/congress-cli vote_proposal -s 0xd5da2b33c1f620a94bf2039b9cb540853e7928d7 -i PROPOSAL_ID -a true --rpc_laddr https://rpc.juchain.org 

./build/congress-cli sign -f voteProposal.json -k miner2.key -p password.file --chainId 210000

./build/congress-cli send -f voteProposal_signed.json --rpc_laddr https://rpc.juchain.org 

# miner3
./build/congress-cli vote_proposal -s 0x311b37f01c04b84d1f94645bfbd58d82fc03f709 -i PROPOSAL_ID -a true --rpc_laddr https://rpc.juchain.org 

./build/congress-cli sign -f voteProposal.json -k miner3.key -p password.file --chainId 210000 

./build/congress-cli send -f voteProposal_signed.json --rpc_laddr https://rpc.juchain.org 

# step3 查看新增矿工的信息 (活动状态为 1 表示Active)
./build/congress-cli miner --rpc_laddr https://rpc.juchain.org -a 0x4d432df142823ca25b21bc3f9744ed21a275bdea
# 查看当前所有矿工信息
./build/congress-cli miners --rpc_laddr https://rpc.juchain.org
```

## 8. 主网修改配置

### 8.1 创建提案

```shell
# 配置项ID对应的配置项信息
# 0 proposalLastingPeriod, 1 punishThreshold, 2 removeThreshold, 3 decreaseRate, 4 withdrawProfitPeriod
./build/congress-cli create_config_proposal -p 提案矿工地址 -i 配置项ID -v 配置项取值 --rpc_laddr https://rpc.juchain.org 

# 示例：修改 proposalLastingPeriod 为 86400 秒（注意：使用 -i 参数，不是 -c）
./build/congress-cli create_config_proposal -p 0xccafa71c31bc11ba24d526fd27ba57d743152807 -i 0 -v 86400 --rpc_laddr https://rpc.juchain.org 

./build/congress-cli sign -f createUpdateConfigProposal.json -k miner1.key -p password.file --chainId 210000

./build/congress-cli send -f createUpdateConfigProposal_signed.json --rpc_laddr https://rpc.juchain.org 
# 这条命令执行后可以获取到提案ID，记录提案ID用于后续投票
```

### 8.2 验证者投票

> 投票步骤与之前相同，请参考 [7.2 验证者投票](#72-验证者投票)

## 9. 主网矿工收益提取

```shell
# step1 创建原始交易
./build/congress-cli withdraw_profits -a 矿工地址 --rpc_laddr https://rpc.juchain.org 

# step2 交易签名
./build/congress-cli sign -f withdrawProfits.json -k miner.key -p password.file --chainId 210000

# step3 发送交易
./build/congress-cli send -f withdrawProfits_signed.json --rpc_laddr https://rpc.juchain.org
```

## 10. 工具信息

### 10.1 版本查看

```shell
./build/congress-cli version
```

### 10.2 帮助信息

```shell
./build/congress-cli help
./build/congress-cli [command] --help  # 查看特定命令的帮助
```

### 10.3 测试脚本

项目包含测试脚本，可以快速验证系统状态：

```shell
cd sys-contract/congress-cli
chmod +x test_congress.sh
./test_congress.sh
```

## 11. 注意事项

### 11.1 重要提醒

- ⚠️ **验证者要求**: 只有当前有效的验证者才能创建提案和投票
- ⚠️ **网络同步**: 在恢复矿工身份前，确保节点已完全同步到最新状态
- ⚠️ **提案ID**: 每次操作都会生成新的提案ID，务必使用正确的ID
- ⚠️ **密钥安全**: 妥善保管密钥文件和密码文件

### 11.2 常见错误

1. **"Validator only"**: 当前账户不是有效验证者
2. **"You can't vote for a proposal twice"**: 该验证者已对此提案投过票
3. **"gas estimation failed"**: 交易参数错误或网络问题

### 11.3 系统合约地址

- **Validators**: `0x000000000000000000000000000000000000f000`
- **Punish**: `0x000000000000000000000000000000000000f001`
- **Proposal**: `0x000000000000000000000000000000000000f002`

### 11.4 网络信息

- **测试网**: `https://testnet-rpc.juchain.org` (Chain ID: 202599)
- **主网**: `https://rpc.juchain.org` (Chain ID: 210000)
