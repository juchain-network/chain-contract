## 1.创建提案

```
./congress --chainId 202599 --rpc_laddr https://testnet-rpc.juchain.org 
./congress --chainId 210000 --rpc_laddr https://rpc.juchain.org 

```

### 1.1. 创建原始交易
```shell
./congress create_proposal -p 提案矿工地址 -t 新矿工地址 -o add  --rpc_laddr https://rpc.juchain.org 

# 测试样例命令
./congress create_proposal -p 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b -t 0x029DAB47e268575D4AC167De64052FB228B5fA41 -o add  --rpc_laddr https://testnet-rpc.juchain.org 

```


### 1.2. 签名交易
```shell
./congress sign -f createProposal.json -k 钱包文件 -p 钱包密码  --chainId 210000 

# 测试样例文件
./congress sign -f createProposal.json -k miner1.key -p juchain  --chainId 202599 
```

### 1.3. 发送交易
```shell
./congress send  -f createProposal_signed.json -p 提案矿工地址  --rpc_laddr https://rpc.juchain.org 

# 测试样例文件
./congress send  -f createProposal_signed.json -p 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b  --rpc_laddr https://testnet-rpc.juchain.org 
```
>  执行成功后会生成提案信息，如下：
```
Wait for tx to be finished executing with hash 0x50a0b73f5dc2f8f2f4acb7611d07ca1fc3e6f1f6bc51f9b30bd02d79ad7d186d
tx confirmed in block 780439
Proposal ID: 80bae77feed9dbc69d162ed81160b1b32fa56a1e91b724ef0d846cb83780b26d
Proposer: 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b
Destination: 0x029DAB47e268575D4AC167De64052FB228B5fA41
Flag: true
Time: 1743154658
```

## 2.提案投票

### 2.1. 创建原始交易
```shell
./congress vote_proposal -s 签名矿工地址 -i 提案ID -a 是否通过  --rpc_laddr https://rpc.juchain.org 

# 测试样例命令
./congress vote_proposal -s 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b -i  80bae77feed9dbc69d162ed81160b1b32fa56a1e91b724ef0d846cb83780b26d -a true  --rpc_laddr https://testnet-rpc.juchain.org 

```

### 2.2. 签名交易
```shell
./congress sign -f voteProposal.json -k 钱包文件 -p 钱包密码  --chainId 210000 

# 测试样例文件
./congress sign -f voteProposal.json -k miner1.key -p juchain  --chainId 202599 
```

### 2.3. 发送交易
```shell
./congress send  -f voteProposal_signed.json -p 提案矿工地址  --rpc_laddr https://rpc.juchain.org 

# 测试样例文件
./congress send  -f voteProposal_signed.json -p 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b  --rpc_laddr https://testnet-rpc.juchain.org 
```

## 3. 查询操作
> 查询所有活动矿工
```shell
./congress miners  --rpc_laddr https://rpc.juchain.org 
```

> 查询单个矿工
```shell
./congress miner  --rpc_laddr https://rpc.juchain.org -a 0x311B37f01c04B84d1f94645BfBd58D82fc03F709
```

> 测试交易
> minner1 创建提案，新增 0x029DAB47e268575D4AC167De64052FB228B5fA41 作为新的矿工， 创建完提案后，miner1,miner2,miner3 投票通过
>
```shell
# step1 创建提案交易，并签名发送
./congress create_proposal -p 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b -t 0x029DAB47e268575D4AC167De64052FB228B5fA41 -o add  --rpc_laddr https://testnet-rpc.juchain.org 
./congress sign -f createProposal.json -k miner1.key -p juchain  --chainId 202599 
./congress send  -f createProposal_signed.json -p 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b  --rpc_laddr https://testnet-rpc.juchain.org 
# 这条命令执行后可以获取到提案ID

# step2 3个矿工对提案进行投票
# miner1
./congress vote_proposal -s 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b -i  2ff96b6f050e8bd407f502b1dc2a39ee66b0291cbfab0e46f2e4bae46f8d40fc -a true  --rpc_laddr https://testnet-rpc.juchain.org 
./congress sign -f voteProposal.json -k miner1.key -p juchain  --chainId 202599 
./congress send  -f voteProposal_signed.json -p 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b  --rpc_laddr https://testnet-rpc.juchain.org 

# miner2
./congress vote_proposal -s 0x81f7a79a51edba249efa812eb2d5478f696f7558 -i  2ff96b6f050e8bd407f502b1dc2a39ee66b0291cbfab0e46f2e4bae46f8d40fc -a true  --rpc_laddr https://testnet-rpc.juchain.org 
./congress sign -f voteProposal.json -k miner2.key -p juchain  --chainId 202599 
./congress send  -f voteProposal_signed.json -p 0x81f7a79a51edba249efa812eb2d5478f696f7558  --rpc_laddr https://testnet-rpc.juchain.org 

# miner3
./congress vote_proposal -s 0x578c39eaf09a4e1abf428c423970b59bb8baf42e -i  2ff96b6f050e8bd407f502b1dc2a39ee66b0291cbfab0e46f2e4bae46f8d40fc -a true  --rpc_laddr https://testnet-rpc.juchain.org 
./congress sign -f voteProposal.json -k miner3.key -p juchain  --chainId 202599 
./congress send  -f voteProposal_signed.json -p 0x578c39eaf09a4e1abf428c423970b59bb8baf42e  --rpc_laddr https://testnet-rpc.juchain.org 


# step3 查看新增矿工的信息
./congress miner  --rpc_laddr https://rpc.juchain.org -a 0x029DAB47e268575D4AC167De64052FB228B5fA41
```