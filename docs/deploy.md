# 部署 aac

## 部署流程

`aac`可认为是以太坊`fork`的私链，部署流程总体与以太坊私链部署流程一致，可基本参照以太坊私链部署流程。

### 共识相关参数

创世块文件可直接配置：

1. period: 出块时间间隔
2. epoch: 块周期

合约参数：

1. punishThreshold = 24: 没收收益阈值。
2. removeThreshold = 48: 移除收益阈值
3. decreaseRate = 24: 削减比例，具体参考下面的共识算法介绍
4. withdrawProfitPeriod = 28800: 赎回收益间隔区块数量
5. increasePeriod = 100: 增发周期
6. receiverAddr: 默认未设置，增发`aac`接收地址
7. proposalLastingPeriod: 提案有效周期

修改合约参数需重新编译生成系统合约代码，并在创世块文件中更新相应的代码。

### 创世块文件重要配置说明

共识算法必须设置为`congress`:

```json
"congress": {
    "period": 3,
    "epoch": 200
}
```

1. validators 合约地址需设置对应合约代码：0x000000000000000000000000000000000000f000
2. punish 合约地址需设置对应合约代码：0x000000000000000000000000000000000000f001
3. proposal 合约地址需设置对应合约代码：0x000000000000000000000000000000000000f002
4. extraData 中包含了初始挖矿节点列表，设置不同的初始矿工需要修改此部分的内容。

```json
"extraData": "0x00000000000000000000000000000000000000000000000000000000000000000899c97c1245d88cfd6a33690a1b081e7b0478ff16577d8170cfd7f05a46454023570441397ad86816691afe40826a1bd34a2ec903009edc1756103630222ebda509ff15fc12c6c6ec0a13704acd6bff3165ad3ec91f35bf8a1cfb326237c5c57718b8af31cba086283d1543213d63884435d2453471039f3948e6dac060cd6d89acfe969361d77e89848a143a35dcd454af071d7f0d06850ead58dfb0e00bbe5e22e24722653222b13eae47405c55f41439bca76bb08e9fc04c6a07a041831375ac8eac4fddc6156ef32883c7bb2d991edcb43caff8b4f0f0f69d93761604606c155d538606a10310259c7529f26bcdb3cb1478db7e0b7e17c326af113fb15ed2d71be9b7359e99aa3e0a107256b2d5a26389ae6b886be6b835e3b08f488d345da066fd0d5c33f1f589050ac224979ebd7902a1e4760c7013ca828c8ff58f4fe6474211c6c7f3fd7fdfb149b9f6ec5debf2c1fee90d933d875b13efb237d209088c0e2ade0dc45eea4771671bcc59b7fa031dc1048c8e04cb6aaee4ee01977e6091f848ee214765302d1c756b8f248df0ce480ab907f1c0eeec30afb2b438984f973f230000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
"000000000000000000000000000000000000f000": {
      "balance": "0x0",
      "code": ""
},
"000000000000000000000000000000000000f001": {
      "balance": "0x0",
      "code": ""
},
"000000000000000000000000000000000000f002": {
      "balance": "0x0",
      "code": ""
},
```

上述文件中的`code`需替换为合约实际的代码，实际代码可通过`truffle`或`hardhat`编译合约得到:

`truffle compile` 或 `npx hardhat compile`，编译生成的文件在`build`或`artifacts/contracts`中，在其中找到合约生成对应的`json`文件，`deployedBytecode`字段的内容即为需要设置的`code`的具体内容。

### 编译部署

部署流程可参考以太坊私链搭建文档(https://github.com/ethereum/go-ethereum#operating-a-private-network)，本文只对主要步骤进行说明。

1. 准备好`golang`及其他相关环境(gcc,c++等)，进入项目代码根目录执行`make all`或者`make geth`编译代码，生成的命令在`build/bin`目录之下。

2. 准备好创世块文件。创世块文件可通过`puppeth`先生成样式文件（选择`clique`共识，设定好初始矿工节点列表，数量最好为奇数），后手动修改创世块配置文件，将共识算法改为`congress`，并将系统合约代码在文件中设置好(测试机器`1`中`/home/deploy/network`下有`test-genesis.json`文件可直接使用)。合约代码可由合约项目编译生成，也可使用合约项目中已编译好的代码（如果对合约参数进行了修改，需重新编译设置代码）。

3. 初始化创世块，并启动各个节点。

   1. 使用命令`./geth init test-genesis.json`初始化创世块，可根据实际需求更改命令内容。
   2. 启动节点，并开始挖矿操作，需要准备好矿工密钥文件，密码文件，节点`p2p`密钥等等，命名示例：`./geth --datadir data --port 21302 -nodekey meta/node2.key --allow-insecure-unlock --mine --unlock 0x9e3A1865A8bd5C89719c55Ba28dAbc91A1210e86 --password meta/pw`,命令中的参数信息请根据实际的情况进行修改。

4. 链接各个节点。节点启动后需要将节点连接到一起，可以使用命令行的方式，或者使用`static-nodes.json`的方式。以`static-nodes.json`为例，需要准备好各个节点的地址及端口（可通过`bootnode`生成`node key`，并获取相应的地址信息）。当文件准备好后，放入节点数据存入目录（可以在节点启动之前就放入目录中）。测试机器`1`中已经设置了相应的文件，对其手动编辑即可。

5. 检查网上是否正常出块。无法正常出块则检查各个节点之间的连接情况。

### 新矿工加入

新矿工加入现有网络需要经过创建提案、节点投票`2`个流程。具体内容可以参考[添加新节点脚本](./scripts/add_new_node.js)

测试机器`1`新节点添加流程:

1. 准备好新节点的私钥，将其在[脚本](./scripts/add_new_node.js)中设置好相应的`toAdd`地址。

2. 进入本项目根目录，安装库文件`npm install`

3. 运行添加新节点脚本`npx hardhat run scripts/add_new_node.js --network aacTest`

4. 等待新周期切换时，新节点即可进行挖矿操作，具体启动方式参见`编译部署`。

具体节点加入说明可以参考下面的共识算法说明。

### 配置更新

配置更新需要先提交提案，之后当前矿工对该提案进行投票，通过后配置更新成新设定的值，具体内容可参考[配置更新脚本](./scripts/update_config.js)

测试机器`1`更新配置流程:

1. 进入本项目根目录，确保库已安装`npm install`

2. 运行配置更新脚本`npx hardhat run scripts/update_config.js --network aacTest`

其他配置的更新与本方法一致，可以更改相应参数测试。

### 矿工准入

新加入矿工需首先发送交易至`Proposal`合约发起提案。

```solidity
createProposal(address dst, bool flag string calldata details)
```

提案创建成功后会提交`LogCreateProposal`日志，可通过该日志得到提案的 id。现有矿工可对该提案进行投票

```solidity
voteProposal(bytes32 id, bool auth)
```

当提案同意人数超半数时，该提案通过。提案通过后矿工可发送交易至 Validators 合约创建基本信息:

```solidity
createOrEditValidator(
    address payable feeAddr,
    string calldata moniker,
    string calldata identity,
    string calldata website,
    string calldata email,
    string calldata details
)
```

### 矿工收益赎回

矿工设置的`feeAddr`可赎回手续费奖励`validator`，两次赎回之间间隔需大于`WithdrawProfitPeriod`

```solidity
withdrawProfits(address validator)
```

## 共识算法介绍

`posa`算法基于`clique`及`dpos`算法修改而来，增加了系统合约实现了`Validator`准入的相关功能，可以对 `Validator`进行激励(均分获得手续费)和惩罚(没收当前收益、移除出`validator`列表)。

系统功能(共识代码功能):

1. 在第一个块时，初始化系统合约

2. 在块周期结束时(`number%Epoch==0`)时:

   - 通过系统合约`getTopValidators`获取当前的排名靠前的`TopValidators`，并将其填入`extraData`字段
   - 调用系统合约`updateActiveValidatorSet`更新合约当前激活的`Validators`列表
   - 调用系统合约`decreaseMissedBlocksCounter`尝试削减`validator`的出错次数，避免因出错次数一直累加而被意外移除出`validator`列表

3. 生产块周期第一个块时更换为新的`validator`列表，此时新的`validator`可以出块，合约中`validator`列表变化必须在下一块周期才会生效。

4. 当有`out of turn`的块出现时，且本应出块的`validator`最近未出块，则`validator`调用系统合约`punish`接口对 `validator`进行惩罚。如果`validator`出错次数达到`punishThreshold`(默认 10)，则会没收当前收益。当达到 `removeThreshold`(默认 30)时，则踢出`validator`列表，状态设置为`Jailed`。

5. 从合约中获取增发周期和增发`acc`接收地址，等经过设定的区块后，增发固定数量`acc`至指定地址。

合约功能:

1. `Validator`合约：

   - 通过了`proposal`的用户成为`validator`
   - 自动计算更新矿工出块收益
   - `validator`设置的`feeAddr`可以赎回收益

2. `Proposal`合约：

   - 用户创建提案，申请成为`validator`
   - 用户创建更新配置的提案
   - `validator`对提案进行投票

3. `Punish`合约: 主要由系统进行调用，惩罚`miss block`的`validator`

## 用户成为 validator 流程

1. 调用`Proposal`合约`createProposal`接口创建提案，提案`ID`在事件中可以得到

2. `Validator`对该合约进行投票

3. 投票通过后，用户调用`Validators`合约`createOrEditValidator`接口，主要提供收益地址获取矿工收益。

## 参数配置

共识参数主要分为两个部分:

1. 出块时间、块周期等内容，可在创世块中配置。

2. 合约参数配置，需要修改合约源代码，并其在创世块中设置对应系统合约的源代码。

### 创世块配置

以下内容必须在创世块中进行配置:

- Period: 出块时间间隔，设置为 0 则表示只有在有交易时才出块

- Epoch: 更新`validators`的块数间隔，系统在块周期第一个块更新`validators`列表

- 账户设置，请在创世块中设置系统合约相应的代码(**deployedCode**):
  - Validators(0x000000000000000000000000000000000000f000): 验证者合约
  - Punish(0x000000000000000000000000000000000000f001): 惩罚合约
  - Proposal(0x000000000000000000000000000000000000f002): 提案合约

### 系统合约参数

通用参数：

- punishThreshold(24): 没收当前收益出错阀值；当`validator`掉线导致未出块次数达到该阀值后，将会没收该 `validator`的当前收益，均分给其他的`validators`

- removeThreshold(48): 移除`validator`列表出错阀值；当`validator`掉线导致未出块次数达到该阀值后，将会没收其当前收益，且将其移除出`validator`列表

- decreaseRate(24): `validator`出错清除比例。每经过一个`epoch`时，系统会自动发送系统交易，对`validator`的出错数量按一定比例进行削减(防止出错比例很小，但值一直累加从而产生惩罚)。削减规则: 如果`validator`出错次未超过 removeThreshold / decreaseRate，不进行任何操作；否则将其出错次数削减 removeThreshold / decreaseRate

- withdrawProfitPeriod(28800): 赎回收益间隔区块数量，矿工两次赎回收益之间的区块间隔必须要大于该值。主要用于限制矿工赎回收益的频率，最大程度保证矿工在因离线被踢出时有收益可以被没收。

- increasePeriod(100): 增发周期，每经过固定数量的区块发送一定数量的`aac`至指定地址

- receiverAddr: 默认未设置，增发`aac`接收地址

- proposalLastingPeriod(7 days): 提案存在时间，当超出该时间段后，提案没有通过的话则作废

## 系统合约用户接口

本部分介绍了可供外部调用的接口.

### Proposal

用户如果想成为`validator`，必须由自己或者其他人创建提案，申请成为`validator`。当前激活的`validator`可以对该提案进行投票，当同意的人数超过了`1半`时，则提案通过后成为`validator`，在系统开始新一轮周期时加入出块列表。

#### createProposal

任意用户创建提案(提案默认持续`7天`，`7天`内提案没有通过则需要重新创建提案)。

```solidity
# dst: 成为validator候选人的地址
# flag: 增加或移除矿工，true: 增加；false: 移除
# details: validator候选人的详细说明(可选, 长度应不大于3000)
createProposal(address dst, string calldata details)


# 交易产生的日志
# id: 提案id，可用于投票
# proposer: 提案人
# dst: validator候选地址
# flag: 添加/移除
# time: 提案时间
event LogCreateProposal(
    bytes32 indexed id,
    address indexed proposer,
    address indexed dst,
    bool flag,
    uint256 time
);
```

#### createUpdateConfigProposal

创建配置更新提案，投票方式同矿工准入一致

```solidity
# cid: 配置对应id
# 0: proposalLastingPeriod
# 1: punishThreshold
# 2: removeThreshold
# 3: decreaseRate
# 4: withdrawProfitPeriod
# 5: increasePeriod
# 6: receiverAddr
# newValue: 配置的新值
createUpdateConfigProposal(uint256 cid, uint256 newValue)
```

#### voteProposal

当前`validator`对提案进行投票。当同意票数超过半数时，则提案通过

```solidity
# id: 提案id
# auth: 是否同意该提案
voteProposal(bytes32 id, bool auth)


# 交易产生日志
# 提案通过日志
# id: 提案id
# dst: validator候选地址
# time: 通过时间
event LogPassProposal(
    bytes32 indexed id,
    address indexed dst,
    uint256 time
);
# 提案未通过日志(超过半数不同意)
# id: 提案id
# dst: validator候选地址
# time: 提案未通过时间
event LogRejectProposal(
    bytes32 indexed id,
    address indexed dst,
    uint256 time
);
```

### Validators

`validator`调用该合约进行赎回收益等相关操作.

管理员调用接口:

validator 调用接口:

```solidity
# 编辑validator信息
# feeAddr: 受益地址
# moniker: 名称，长度不大于70
# identity: 身份信息，长度不大于3000
# website: 网站信息，长度不大于140
# email: 邮件信息，长度不大于140
# details: 详细信息，长度不大于280
createOrEditValidator(
    address payable feeAddr,
    string calldata moniker,
    string calldata identity,
    string calldata website,
    string calldata email,
    string calldata details
)

# 编辑validator信息
# val: validator地址
# fee: 更新后的受益地址
# time: 更新时间
event LogEditValidator(
    address indexed val,
    address indexed fee,
    uint256 time
);



# 赎回出块收益，该方法由收益地址调用，调用时需传入赎回的是哪个validator的收益
# validator: validator地址
withdrawProfits(address validator)
# 交易事件
# val: validator的地址
# fee: 收益人的地址
# aac: aac收益金额
# time: 交易时间
event LogWithdrawProfits(
    address indexed val,
    address indexed fee,
    uint256 aac,
    uint256 time
);



# 查看当前激活的validator列表，当前可以生产块的validator列表
getActiveValidators() returns (address[] memory)
# 查看当前top validator列表，当前质押金额最高的validator列表，下一周期会被激活
getTopValidators() returns (address[] memory)
# 返回validator基本信息
getValidatorDescription(address val) returns (
    string memory moniker,
    string memory identity,
    string memory website,
    string memory email,
    string memory details
)
# 返回validator详细信息
getValidatorInfo(address val)return (
    feeAddr, // 收益地址
    status, // 当前状态
    aacIncoming, // aac收益
    totalJailedAAc, // 被没收的aac
    lastWithdrawProfitsBlock, // 上次赎回收益的块号
);
```
