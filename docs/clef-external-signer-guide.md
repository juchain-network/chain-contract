# Clef 外部签名器 账户管理

Clef 作为外部账户管理和签名工具，甚至可以在专用的安全外部硬件USB上运行

以太坊账户的加密存储文件（keystore V3 格式）

用户输入的 **口令 (password)** 通过 scrypt 派生出对称密钥，再用 AES-128-CTR 加密后得到的 ciphertext。

```json
{
  "address": "1fe9327d22584e2a8eec4539c541cb0ad897f698",  // 账户地址（小写 hex）
  "crypto": {
    "cipher": "aes-128-ctr",                          // 加密算法：AES-128-CTR
    "ciphertext": "811d...",                          // 加密后的私钥内容
    "cipherparams": {
      "iv": "202cbc5aa0448179538a468b6a26bc55"        // 初始化向量
    },
    "kdf": "scrypt",                                  // 密钥派生函数 (Key Derivation Function)
    "kdfparams": {
      "dklen": 32,
      "n": 262144,                                    // scrypt 参数，越大越安全但耗时
      "p": 1,
      "r": 8,
      "salt": "280c..."                               // 随机盐值
    },
    "mac": "c1d7..."                                  // 用来校验解密正确性的哈希
  },
  "id": "1b124e30-73d3-413d-9bd0-4e1a15f6d285",       // UUID
  "version": 3
}

```

用户需要手动审核所有涉及敏感数据的操作，签名是在 Clef 本地完成的

创建和列出账户，或离线签名数据

clef init --configdir clefdata

# 新建账户

clef newaccount --keystore  <path-to-keystore>

# 导入原始私钥

clef importraw <hexkey>

# **列出账户**

clef list-accounts --keystore <path-to-keystore>

clef list-wallets --keystore <path-to-keystore>

# 启动签名器

clef --keystore keys --configdir clefdata --chainid 202599  --http

WARNING!

Clef is an account management tool. It may, like any software, contain bugs.

Please take care to

- backup your keystore files,
- verify that the keystore(s) can be opened with your password.

Clef is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY;

without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR

PURPOSE. See the GNU General Public License for more details.

Enter 'ok' to proceed:

> ok

INFO [08-22|14:34:14.621] Using CLI as UI-channel

INFO [08-22|14:34:14.734] Loaded 4byte database                    embeds=268,621 locals=0 local=./4byte-custom.json

WARN [08-22|14:34:14.734] Failed to open master, rules disabled    err="failed stat on clefdata/masterseed.json: stat clefdata/masterseed.json: no such file or directory"

INFO [08-22|14:34:14.734] Starting signer                          chainid=202,599 keystore=keys light-kdf=false advanced=false

INFO [08-22|14:34:14.744] Audit logs configured                    file=audit.log

INFO [08-22|14:34:14.745] HTTP endpoint opened                     url=<http://127.0.0.1:8550/>

INFO [08-22|14:34:14.745] IPC endpoint opened                      url=clefdata/clef.ipc

- ------ Signer info -------
- intapi_version : 7.0.1
- extapi_version : 6.1.0
- extapi_http : <http://127.0.0.1:8550/>
- extapi_ipc : clefdata/clef.ipc
- ------ Available accounts -------

0. 0x3858FfcA201b0A7D75fd23BB302C12332c5e4000 at keystore:///Users/enty/ju-chain-work/chain/build/bin/keys/UTC--2025-08-22T06-33-36.854192000Z--3858ffca201b0a7d75fd23bb302c12332c5e4000

1. 0x3d968443D9B72bCeF4409B3A2D5e31031390FC82 at keystore:///Users/enty/ju-chain-work/chain/build/bin/keys/UTC--2025-08-22T06-33-49.461865000Z--3d968443d9b72bcef4409b3a2d5e31031390fc82

WARN [08-22|14:34:56.398] Served account_signTransaction           conn=127.0.0.1:58911 reqid=1 duration="247.202µs" err="invalid argument 0: json: cannot unmarshal non-string into Go struct field SendTxArgs.chainId of type *hexutil.Big"

ERROR[08-22|14:35:32.655] Signing request with wrong chain id      requested=202,391 configured=202,599

WARN [08-22|14:35:32.656] Served account_signTransaction           conn=127.0.0.1:59774 reqid=1 duration=1.252915ms  err="requested chainid 202391 does not match the configuration of the signer"

ERROR[08-22|14:36:26.262] Signing request with wrong chain id      requested=202,583 configured=202,599

WARN [08-22|14:36:26.262] Served account_signTransaction           conn=127.0.0.1:61145 reqid=1 duration="94.004µs"  err="requested chainid 202583 does not match the configuration of the signer"

ERROR[08-22|14:36:57.379] Signing request with wrong chain id      requested=202,583 configured=202,599

WARN [08-22|14:36:57.379] Served account_signTransaction           conn=127.0.0.1:61854 reqid=1 duration="98.968µs"  err="requested chainid 202583 does not match the configuration of the signer"

ERROR[08-22|14:38:40.566] Signing request with wrong chain id      requested=202,583 configured=202,599

WARN [08-22|14:38:40.566] Served account_signTransaction           conn=127.0.0.1:64264 reqid=1 duration="94.902µs"  err="requested chainid 202583 does not match the configuration of the signer"

- -------- Transaction request-------------

to:    0x1234567890123456789012345678901234567890

from:               0x3858ffca201b0a7d75fd23bb302c12332c5e4000 [chksum INVALID]

value:              1000000000000000000 wei

gas:                0x5208 (21000)

gasprice: 20000000000 wei

nonce:    0x0 (0)

chainid:  0x31767

Request context:

127.0.0.1:49531 -> http -> 127.0.0.1:8550

Additional HTTP header data, provided by the external caller:

User-Agent: "curl/8.1.2"

Origin: ""

- ------------------------------------------

Approve? [y/N]:

> > y

## Account password

Please enter the password for account 0x3858FfcA201b0A7D75fd23BB302C12332c5e4000

>

- ----------------------

Transaction signed:

{

"type": "0x0",

"chainId": "0x31767",

"nonce": "0x0",

"to": "0x1234567890123456789012345678901234567890",

"gas": "0x5208",

"gasPrice": "0x4a817c800",

"maxPriorityFeePerGas": null,

"maxFeePerGas": null,

"value": "0xde0b6b3a7640000",

"input": "0x",

"v": "0x62ef2",

"r": "0xfd80a31a67b54d3d9ffc0c4f27db69f369a22b00792b91102663252d35a89da3",

"s": "0x383db7f2a92dea4322793c2a393b760615ded4e99cfb7489a99bd7667d0bdeff",

"hash": "0xde20680041d236d200aa7dfbe41d2eea5c0d78a7e9fe64dc34ecd2a33c1f766e"

}

## 签名交易

```bash

curl -X POST \

-H "Content-Type: application/json" \

--data '{"jsonrpc":"2.0","method":"account_signTransaction","params":[{"from": "0x3858ffca201b0a7d75fd23bb302c12332c5e4000", "to": "0x1234567890123456789012345678901234567890", "value": "0xde0b6b3a7640000", "gas": "0x5208", "gasPrice": "0x4a817c800", "nonce": "0x0", "chainId": "0x31767"}], "id":1}' \

http://127.0.0.1:8550

{"jsonrpc":"2.0","id":1,"result":{"raw":"0xf86f808504a817c800825208941234567890123456789012345678901234567890880de0b6b3a76400008083062ef2a0fd80a31a67b54d3d9ffc0c4f27db69f369a22b00792b91102663252d35a89da3a0383db7f2a92dea4322793c2a393b760615ded4e99cfb7489a99bd7667d0bdeff","tx":{"type":"0x0","chainId":"0x31767","nonce":"0x0","to":"0x1234567890123456789012345678901234567890","gas":"0x5208","gasPrice":"0x4a817c800","maxPriorityFeePerGas":null,"maxFeePerGas":null,"value":"0xde0b6b3a7640000","input":"0x","v":"0x62ef2","r":"0xfd80a31a67b54d3d9ffc0c4f27db69f369a22b00792b91102663252d35a89da3","s":"0x383db7f2a92dea4322793c2a393b760615ded4e99cfb7489a99bd7667d0bdeff","hash":"0xde20680041d236d200aa7dfbe41d2eea5c0d78a7e9fe64dc34ecd2a33c1f766e"}}}

```

## 发送交易

···bash

curl -X POST \

-H "Content-Type: application/json" \

--data '{"jsonrpc":"2.0","method":"eth_sendRawTransaction","params":["0xf86f808504a817c800825208941234567890123456789012345678901234567890880de0b6b3a76400008083062ef2a0fd80a31a67b54d3d9ffc0c4f27db69f369a22b00792b91102663252d35a89da3a0383db7f2a92dea4322793c2a393b760615ded4e99cfb7489a99bd7667d0bdeff"],"id":1}' \

<http://127.0.0.1:8556>

{"jsonrpc":"2.0","id":1,"result":"0xde20680041d236d200aa7dfbe41d2eea5c0d78a7e9fe64dc34ecd2a33c1f766e"}%

···

## 查询交易状态收据

```

curl -X POST \

-H "Content-Type: application/json" \

--data '{"jsonrpc":"2.0","method":"eth_getTransactionReceipt","params":["0xde20680041d236d200aa7dfbe41d2eea5c0d78a7e9fe64dc34ecd2a33c1f766e"],"id":1}' \

http://127.0.0.1:8556

```

## 4. 使用Geth控制台与JuChain系统合约交互

### 4.1 准备工作：检查网络状态

首先在Geth控制台中检查当前节点状态：

```jsx
// 检查网络基本信息
eth.blockNumber           // 当前区块高度
eth.mining               // 是否正在挖矿
net.version              // 网络ID（应为202599）
admin.peers.length       // 连接的节点数量

// 检查预设账户余额
var account1 = "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
var account2 = "0x3858ffca201b0a7d75fd23bb302c12332c5e4000"

console.log("账户1余额:", web3.fromWei(eth.getBalance(account1), "ether"), "JU")
console.log("账户2余额:", web3.fromWei(eth.getBalance(account2), "ether"), "JU")
```

### 4.2 账户管理与解锁

```jsx
// 查看所有可用账户
eth.accounts

// 解锁主账户用于交易（根据实际密码修改）
personal.unlockAccount(account1, "password123", 300)  // 解锁300秒

// 设置挖矿账户并启动挖矿（如果还未启动）
miner.setEtherbase(account1)
if (!eth.mining) {
    miner.start(1)
    console.log("挖矿已启动")
}
```

### 4.3 发送基础交易测试

```jsx
// 发送简单的转账交易
var txHash = eth.sendTransaction({
    from: account1,
    to: account2,
    value: web3.toWei(1, "ether"),
    gas: 21000,
    gasPrice: web3.toWei(20, "gwei")
})

console.log("交易哈希:", txHash)

// 等待交易确认并查看收据
setTimeout(function() {
    var receipt = eth.getTransactionReceipt(txHash)
    console.log("交易状态:", receipt ? "已确认" : "待确认")
    if (receipt) {
        console.log("Gas使用量:", receipt.gasUsed)
        console.log("区块号:", receipt.blockNumber)
    }
}, 3000)
```

## 5. 与JuChain系统合约交互

### 5.1 初始化系统合约连接

JuChain的系统合约在创世区块时预部署在固定地址，我们可以直接与它们交互：

```jsx
// 定义系统合约地址
var CONTRACT_ADDRESSES = {
    validators: "0x000000000000000000000000000000000000f000",
    punish: "0x000000000000000000000000000000000000f001", 
    proposal: "0x000000000000000000000000000000000000f002",
    staking: "0x000000000000000000000000000000000000f003"
}

// Validators合约ABI（核心验证者管理）
var validatorsABI = [
    {"inputs":[],"name":"getActiveValidators","outputs":[{"internalType":"address[]","name":"","type":"address[]"}],"stateMutability":"view","type":"function"},
    {"inputs":[],"name":"getTopValidators","outputs":[{"internalType":"address[]","name":"","type":"address[]"}],"stateMutability":"view","type":"function"},
    {"inputs":[{"internalType":"address","name":"val","type":"address"}],"name":"getValidatorInfo","outputs":[{"internalType":"address","name":"feeAddr","type":"address"},{"internalType":"uint256","name":"status","type":"uint256"},{"internalType":"uint256","name":"accumulatedRewards","type":"uint256"},{"internalType":"uint256","name":"totalJailedHB","type":"uint256"},{"internalType":"uint256","name":"lastWithdrawProfitsBlock","type":"uint256"}],"stateMutability":"view","type":"function"}
]

// Staking合约ABI（质押管理）
var stakingABI = [
    {"inputs":[{"internalType":"uint256","name":"commissionRate","type":"uint256"}],"name":"register","outputs":[],"stateMutability":"payable","type":"function"},
    {"inputs":[{"internalType":"address","name":"validator","type":"address"}],"name":"getValidatorInfo","outputs":[{"internalType":"uint256","name":"selfStake","type":"uint256"},{"internalType":"uint256","name":"totalDelegated","type":"uint256"},{"internalType":"uint256","name":"totalStake","type":"uint256"},{"internalType":"uint256","name":"commissionRate","type":"uint256"},{"internalType":"bool","name":"isJailed","type":"bool"},{"internalType":"uint256","name":"jailUntilBlock","type":"uint256"}],"stateMutability":"view","type":"function"},
    {"inputs":[{"internalType":"uint256","name":"limit","type":"uint256"}],"name":"getTopValidators","outputs":[{"internalType":"address[]","name":"","type":"address[]"}],"stateMutability":"view","type":"function"}
]

// 创建合约实例
var validatorsContract = eth.contract(validatorsABI).at(CONTRACT_ADDRESSES.validators)
var stakingContract = eth.contract(stakingABI).at(CONTRACT_ADDRESSES.staking)

console.log("✅ 系统合约已连接")
console.log("📊 Validators合约:", CONTRACT_ADDRESSES.validators)
console.log("💰 Staking合约:", CONTRACT_ADDRESSES.staking)
```

### 5.2 查询验证者信息

```jsx
// 查询当前活跃验证者
console.log("\n=== 📋 活跃验证者列表 ===")
var activeValidators = validatorsContract.getActiveValidators()
console.log("活跃验证者数量:", activeValidators.length)
activeValidators.forEach(function(addr, index) {
    console.log((index + 1) + ".", addr)
})

// 查询顶级验证者（从Staking合约）
console.log("\n=== 🏆 顶级验证者（按质押排序） ===")
var topValidators = stakingContract.getTopValidators(21)
console.log("顶级验证者数量:", topValidators.length)

// 查询特定验证者的详细信息
console.log("\n=== 🔍 验证者详细信息 ===")
var targetValidator = activeValidators[0]  // 使用第一个活跃验证者
console.log("查询验证者:", targetValidator)

// 从Validators合约查询
var validatorInfo = validatorsContract.getValidatorInfo(targetValidator)
console.log("\n📊 Validators合约信息:")
console.log("  收费地址:", validatorInfo[0])
console.log("  状态码:", validatorInfo[1].toString())
console.log("  累积奖励:", web3.fromWei(validatorInfo[2], "ether"), "JU")
console.log("  监禁次数:", validatorInfo[3].toString())
console.log("  最后提取区块:", validatorInfo[4].toString())

// 从Staking合约查询
var stakingInfo = stakingContract.getValidatorInfo(targetValidator)
console.log("\n💰 Staking合约信息:")
console.log("  自质押:", web3.fromWei(stakingInfo[0], "ether"), "JU")
console.log("  总委托:", web3.fromWei(stakingInfo[1], "ether"), "JU") 
console.log("  总质押:", web3.fromWei(stakingInfo[2], "ether"), "JU")
console.log("  佣金率:", (stakingInfo[3].toNumber() / 100).toFixed(2) + "%")
console.log("  是否监禁:", stakingInfo[4])
console.log("  监禁至区块:", stakingInfo[5].toString())
```

### 5.3 网络状态总览

```jsx
// 综合网络状态查询
console.log("\n=== 🌐 JuChain网络状态总览 ===")
console.log("当前区块高度:", eth.blockNumber)
console.log("网络ID:", net.version)
console.log("是否挖矿:", eth.mining)
console.log("连接节点数:", admin.peers.length)

var latestBlock = eth.getBlock("latest")
console.log("最新区块信息:")
console.log("  区块哈希:", latestBlock.hash)
console.log("  矿工地址:", latestBlock.miner)
console.log("  交易数量:", latestBlock.transactions.length)
console.log("  Gas使用:", latestBlock.gasUsed.toString())
console.log("  时间戳:", new Date(latestBlock.timestamp * 1000))

// 检查是否为PoSA共识
var isValidator = activeValidators.indexOf(latestBlock.miner) !== -1
### 5.4 注册新验证者（高级操作）

```jsx
// ⚠️  注意：这是一个需要大量质押的真实操作示例
// 只有在你确实想要注册为验证者时才执行

console.log("\n=== 💎 验证者注册流程 ===")

// 检查账户余额（需要至少10000 JU）
var registrationAccount = account1  // 使用之前定义的账户
var currentBalance = web3.fromWei(eth.getBalance(registrationAccount), "ether")
var requiredStake = 10000  // 最低质押要求

console.log("注册账户:", registrationAccount)
console.log("当前余额:", currentBalance, "JU")
console.log("最低质押:", requiredStake, "JU")

if (parseFloat(currentBalance) >= requiredStake) {
    console.log("✅ 余额充足，可以进行注册")
    
    // 设置佣金率（以基点为单位，500 = 5%）
    var commissionRate = 500  // 5%佣金
    
    console.log("\n📝 准备注册参数:")
    console.log("  质押金额:", requiredStake, "JU")
    console.log("  佣金率:", (commissionRate/100).toFixed(2) + "%")
    
    // 执行注册（取消注释以实际执行）
    /*
    console.log("\n🚀 正在注册验证者...")
    var registerTx = stakingContract.register(commissionRate, {
        from: registrationAccount,
        value: web3.toWei(requiredStake, "ether"),
        gas: 500000,
        gasPrice: web3.toWei(20, "gwei")
    })
    
    console.log("注册交易哈希:", registerTx)
    console.log("请等待交易确认...")
    
    // 检查交易状态
    setTimeout(function() {
        var receipt = eth.getTransactionReceipt(registerTx)
        if (receipt) {
            console.log("✅ 注册交易已确认")
            console.log("Gas使用量:", receipt.gasUsed)
            console.log("交易状态:", receipt.status === "0x1" ? "成功" : "失败")
            
            // 验证注册结果
            var newStakingInfo = stakingContract.getValidatorInfo(registrationAccount)
            console.log("\n🎉 注册后信息:")
            console.log("  自质押:", web3.fromWei(newStakingInfo[0], "ether"), "JU")
            console.log("  佣金率:", (newStakingInfo[3].toNumber() / 100).toFixed(2) + "%")
        } else {
            console.log("⏳ 交易仍在确认中...")
        }
    }, 5000)
    */
    
} else {
    console.log("❌ 余额不足，无法注册验证者")
    console.log("需要额外:", (requiredStake - parseFloat(currentBalance)).toFixed(2), "JU")
}
```

### 5.5 完整操作流程总结

```jsx
// 🎯 完整的验证者查询和管理流程
console.log("\n=== 🎯 JuChain验证者管理总览 ===")

function displayValidatorSummary() {
    var activeVals = validatorsContract.getActiveValidators()
    var topVals = stakingContract.getTopValidators(21)
    
    console.log("📊 验证者统计:")
    console.log("  活跃验证者:", activeVals.length)
    console.log("  Top验证者:", topVals.length)
    
    console.log("\n🏆 前3名验证者详情:")
    for (var i = 0; i < Math.min(3, activeVals.length); i++) {
        var addr = activeVals[i]
        var stakingInfo = stakingContract.getValidatorInfo(addr)
        console.log((i+1) + ". " + addr)
        console.log("   质押:", web3.fromWei(stakingInfo[2], "ether"), "JU")
        console.log("   佣金:", (stakingInfo[3].toNumber() / 100).toFixed(2) + "%")
        console.log("   状态:", stakingInfo[4] ? "监禁" : "正常")
    }
    
    console.log("\n📈 网络健康状况:")
    console.log("  当前区块:", eth.blockNumber)
    console.log("  最新区块矿工:", eth.getBlock("latest").miner)
    console.log("  PoSA共识运行正常:", activeVals.length > 0 ? "✅" : "❌")
}

// 执行总览
displayValidatorSummary()
```

**执行说明：**

1. 🚀 **分步执行**：将上述代码分段复制到Geth控制台中执行，每段执行后查看结果
2. ⏱️ **等待确认**：由于你的节点正在挖矿，交易通常在几秒内确认
3. 🔍 **实时监控**：可以随时使用 `eth.blockNumber` 检查当前区块高度
4. 📋 **状态验证**：每次操作后都会显示详细的结果和状态信息

## 6. 使用Congress-CLI工具

除了Geth控制台，JuChain还提供了专门的命令行工具来管理验证者和治理：

### 6.1 验证者查询命令

```bash
# 🔍 查看所有验证者（Validators合约）
./build/congress-cli miners

# 👤 查询特定验证者信息
./build/congress-cli miner -a 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266

# 🏆 查看Staking合约中的顶级验证者
./build/congress-cli staking list-top-validators

# 💰 查询Staking合约中特定验证者信息
./build/congress-cli staking query-validator --address 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
```

### 6.2 治理提案管理

```bash
# 📝 创建添加验证者的提案
./build/congress-cli create_proposal \
    -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
    -t 0x新验证者地址 \
    -o add

# 🗳️ 投票支持提案
./build/congress-cli vote_proposal \
    -s 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \
    -i 提案ID \
    -a

# 📊 查询提案状态
./build/congress-cli proposal -i 提案ID

# 📋 列出所有活跃提案
./build/congress-cli list_proposals
```

### 6.3 自动化脚本

```bash
# 🤖 使用自动化脚本添加验证者（包含完整流程）
./sys-contract/congress-cli/add_validator6.sh

# 这个脚本会自动执行：
# 1. 创建添加验证者提案
# 2. 收集足够的投票
# 3. 执行提案
# 4. 在Staking合约中注册验证者
# 5. 验证所有步骤的结果
```

## 7. 系统合约地址参考

JuChain的系统合约部署在以下固定地址：

- **Validators合约**: `0x000000000000000000000000000000000000f000` - 管理验证者状态和奖励
- **Punish合约**: `0x000000000000000000000000000000000000f001` - 处理验证者惩罚
- **Proposal合约**: `0x000000000000000000000000000000000000f002` - 管理治理提案
- **Staking合约**: `0x000000000000000000000000000000000000f003` - 管理质押和委托

这些合约在创世区块时自动初始化，无需手动部署。
