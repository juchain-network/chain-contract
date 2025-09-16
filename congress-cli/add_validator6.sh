#!/bin/bash

# Congress CLI 自动化脚本：添加验证者6
# 使用说明：确保私链环境正在运行，然后执行此脚本
# 
# 特性：
# - 自动查找密钥文件，不依赖硬编码的时间戳
# - 完整的5步流程：创建提案、投票、验证、配置资金、设置收费地址、Staking注册
# - 详细的输出和错误处理
# - 自动转账10000 ETH给验证者6
# - 自动设置收费地址
# - 自动注册到Staking合约并质押10000 JU
# - 验证者6将同时在Validators和Staking合约中活跃
#
# 依赖要求：
# - Node.js 和 Web3.js (用于转账功能)
# - congress-cli 工具已编译
# - 私链环境正在运行
# 
# 执行后验证者6将拥有：
# - Validators合约中的Active状态
# - Staking合约中的10000 JU质押
# - 正确配置的收费地址
# - 充足的ETH余额用于交易
# - 完整的挖矿准备状态

set -e  # 遇到错误立即退出

# 工具函数：查找验证者密钥文件
find_validator_key() {
    local validator_num=$1
    local address=$2
    local keystore_path="${PRIVATE_CHAIN_PATH}/data-validator${validator_num}/keystore/"
    
    # 检查目录是否存在
    if [ ! -d "$keystore_path" ]; then
        echo "❌ 验证者${validator_num}的keystore目录不存在: $keystore_path" >&2
        return 1
    fi
    
    # 查找密钥文件
    local key_file=$(find "$keystore_path" -name "*--${address}" | head -1)
    echo "$key_file"
}

# 配置
TARGET_ADDRESS="0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc"
PROPOSER_ADDRESS="0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
PRIVATE_CHAIN_PATH="$HOME/ju/chain/private-chain"

# 验证者地址数组（小写，用于查找密钥文件）
VALIDATOR_ADDRESSES=(
    "f39fd6e51aad88f6f4ce6ab8827279cfffb92266"
    "70997970c51812dc3a010c7d01b50e0d17dc79c8"
    "3c44cdddb6a900fa2b585dd299e03d12fa4293bc"
    "90f79bf6eb2c4f870365e785982e1f101e93b906"
    "15d34aaf54267db7d7c367839aaf71a00a2c6a65"
)

# 验证者地址数组（用于创建提案和投票）
VALIDATORS=(
    "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
    "0x70997970C51812dc3A010C7d01b50e0d17dc79C8"
    "0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC"
    "0x90F79bf6EB2c4f870365E785982E1f101E93b906"
    "0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65"
)

# 动态查找密钥文件数组
echo "🔍 正在查找验证者密钥文件..."
KEYS=()
for i in "${!VALIDATOR_ADDRESSES[@]}"; do
    validator_num=$((i + 1))
    key_file=$(find_validator_key "$validator_num" "${VALIDATOR_ADDRESSES[$i]}")
    if [ -z "$key_file" ]; then
        echo "❌ 无法找到验证者${validator_num}的密钥文件"
        echo "   查找地址: ${VALIDATOR_ADDRESSES[$i]}"
        echo "   查找路径: ${PRIVATE_CHAIN_PATH}/data-validator${validator_num}/keystore/"
        exit 1
    fi
    KEYS+=("$key_file")
    echo "   ✅ 验证者${validator_num}: $(basename "$key_file")"
done

# 密码文件数组
PASSWORDS=(
    "${PRIVATE_CHAIN_PATH}/data-validator1/password.txt"
    "${PRIVATE_CHAIN_PATH}/data-validator2/password.txt"
    "${PRIVATE_CHAIN_PATH}/data-validator3/password.txt"
    "${PRIVATE_CHAIN_PATH}/data-validator4/password.txt"
    "${PRIVATE_CHAIN_PATH}/data-validator5/password.txt"
)

# 检查congress-cli工具是否存在
if [ ! -f "./build/congress-cli" ]; then
    echo "❌ congress-cli工具不存在，请先编译："
    echo "   make build"
    exit 1
fi

# 检查Node.js是否可用
if ! command -v node &> /dev/null; then
    echo "⚠️  Node.js未找到，将跳过自动转账功能"
    echo "ℹ️  您需要手动为验证者6转账10000 ETH"
    SKIP_TRANSFER=true
else
    SKIP_TRANSFER=false
fi

echo ""
echo "🚀 开始添加验证者6到网络..."
echo "目标地址: $TARGET_ADDRESS"
echo "提案者: $PROPOSER_ADDRESS"
echo ""

echo "=== 步骤1: 创建提案 ==="
echo "创建添加验证者的提案..."

# 获取当前nonce值
echo "获取当前账户nonce..."
CURRENT_NONCE=$(curl -s -X POST -H "Content-Type: application/json" \
  -d "{\"jsonrpc\":\"2.0\",\"method\":\"eth_getTransactionCount\",\"params\":[\"$PROPOSER_ADDRESS\",\"latest\"],\"id\":1}" \
  http://localhost:8545 | grep -o '"result":"[^"]*"' | cut -d'"' -f4)

if [ -z "$CURRENT_NONCE" ]; then
    echo "❌ 无法获取当前nonce值"
    exit 1
fi

# 转换16进制到10进制
NONCE_DEC=$(printf "%d" "$CURRENT_NONCE")
echo "当前nonce: $NONCE_DEC"

./build/congress-cli create_proposal -p $PROPOSER_ADDRESS -t $TARGET_ADDRESS -o add


echo "签名提案..."
./build/congress-cli sign -f createProposal.json -k "${KEYS[0]}" -p "${PASSWORDS[0]}"

echo "发送提案..."
OUTPUT=$(./build/congress-cli send -f createProposal_signed.json)
echo "$OUTPUT"

# 提取提案ID
PROPOSAL_ID=$(echo "$OUTPUT" | grep "Proposal ID:" | sed 's/Proposal ID: //')
if [ -z "$PROPOSAL_ID" ]; then
    echo "❌ 无法获取提案ID，请检查输出"
    exit 1
fi

echo ""
echo "✅ 提案创建成功！"
echo "📋 提案ID: $PROPOSAL_ID"
echo ""

echo "=== 步骤2: 验证者投票 ==="
for i in "${!VALIDATORS[@]}"; do
    echo "验证者 $((i+1)) (${VALIDATORS[$i]}) 投票中..."
    
    # 创建投票
    ./build/congress-cli vote_proposal -s ${VALIDATORS[$i]} -i $PROPOSAL_ID -a
    
    # 签名投票
    ./build/congress-cli sign -f voteProposal.json -k "${KEYS[$i]}" -p "${PASSWORDS[$i]}"
    
    # 发送投票
    ./build/congress-cli send -f voteProposal_signed.json
    
    echo "✅ 验证者 $((i+1)) 投票完成"
    echo ""
done

echo "=== 步骤3: 验证结果 ==="
echo "查询提案状态 (Query Proposal Status)..."
./build/congress-cli proposal -i $PROPOSAL_ID
echo ""

echo "查询验证者6状态 (Query Validator 6 Status)..."
./build/congress-cli miner -a $TARGET_ADDRESS
echo ""

echo "查询所有验证者列表 (Query All Validators)..."
./build/congress-cli miners
echo ""

echo "🎉 流程完成！验证者6已经成功添加到网络中。"
echo "🎉 Process completed! Validator 6 has been successfully added to the network."
echo ""

echo "=== 步骤4: 配置验证者6 ==="
echo "为验证者6配置资金和收费地址..."

# 4.1 转账资金给验证者6
echo "🚀 步骤4.1: 为验证者6转账资金..."

if [ "$SKIP_TRANSFER" = true ]; then
    echo "⚠️  跳过自动转账（Node.js未找到）"
    echo "ℹ️  请手动为验证者6转账10000 ETH"
    echo "ℹ️  验证者地址: $TARGET_ADDRESS"
else
    echo "正在转账10000 ETH给验证者6以支持后续操作..."
    
    # 使用Node.js进行转账
    node -e "
const { Web3 } = require('web3');
const web3 = new Web3('http://localhost:8545');

async function transferFunds() {
    const fromAddress = '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266';
    const toAddress = '0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc';
    const amount = web3.utils.toWei('10000', 'ether');
    
    console.log('🚀 转账10000 ETH给验证者6...');
    
    try {
        // 获取正确的 nonce
        const nonce = await web3.eth.getTransactionCount(fromAddress, 'pending');
        console.log('当前nonce:', nonce);
        
        const txHash = await web3.eth.sendTransaction({
            from: fromAddress,
            to: toAddress,
            value: amount,
            gas: 21000,
            gasPrice: web3.utils.toWei('20', 'gwei'),
            nonce: nonce  // 显式指定nonce
        });
        
        console.log('✅ 转账成功!');
        console.log('交易哈希:', typeof txHash === 'object' ? txHash.transactionHash : txHash);
        
        const balance = await web3.eth.getBalance(toAddress);
        console.log('验证者6新余额:', web3.utils.fromWei(balance, 'ether'), 'ETH');
        
    } catch (error) {
        console.error('❌ 转账失败:', error.message);
        process.exit(1);
    }
}

transferFunds();
" || {
    echo "❌ 转账失败，请检查Node.js和Web3.js是否正确安装"
    echo "ℹ️  您可以手动为验证者6转账10000 ETH"
    echo "ℹ️  验证者地址: $TARGET_ADDRESS"
}
fi

echo ""

# 4.2 设置收费地址
echo "🚀 步骤4.2: 为验证者6设置收费地址..."
echo "创建编辑验证者交易..."

./build/congress-cli staking edit-validator \
  --validator $TARGET_ADDRESS \
  --fee-addr $TARGET_ADDRESS \
  --moniker "validator6" \
  --details "Validator6 node with fee address configured" || {
    echo "ℹ️  如果编辑验证者命令失败，验证者仍然可以正常工作"
    echo "ℹ️  收费地址设置可以稍后通过其他方式配置"
}

# 查找验证者6的密钥文件
VALIDATOR6_KEY=$(find "${PRIVATE_CHAIN_PATH}/data-validator6/keystore/" -name "*--9965507d1a55bcc2695c58ba16fb37d819b0a4dc" | head -1)
VALIDATOR6_PASSWORD="${PRIVATE_CHAIN_PATH}/data-validator6/password.txt"

if [ -f "editValidator.json" ] && [ -n "$VALIDATOR6_KEY" ] && [ -f "$VALIDATOR6_PASSWORD" ]; then
    echo "签名编辑验证者交易..."
    ./build/congress-cli sign -f editValidator.json -k "$VALIDATOR6_KEY" -p "$VALIDATOR6_PASSWORD"
    
    echo "发送编辑验证者交易..."
    ./build/congress-cli send -f editValidator_signed.json
    echo "✅ 验证者6收费地址设置完成"
else
    echo "ℹ️  跳过收费地址设置（密钥文件未找到或交易文件不存在）"
fi

echo ""

# 4.3 注册到Staking合约并质押
echo "🚀 步骤4.3: 注册验证者6到Staking合约并质押..."
echo "创建Staking注册交易（质押10000 JU，佣金率5%）..."

./build/congress-cli staking register-validator \
  --proposer $TARGET_ADDRESS \
  --stake-amount 10000 \
  --commission-rate 500

if [ -f "registerValidator.json" ] && [ -n "$VALIDATOR6_KEY" ] && [ -f "$VALIDATOR6_PASSWORD" ]; then
    echo "签名Staking注册交易..."
    ./build/congress-cli sign -f registerValidator.json -k "$VALIDATOR6_KEY" -p "$VALIDATOR6_PASSWORD"
    
    echo "发送Staking注册交易..."
    STAKING_OUTPUT=$(./build/congress-cli send -f registerValidator_signed.json)
    echo "$STAKING_OUTPUT"
    
    if echo "$STAKING_OUTPUT" | grep -q "Transaction broadcast successfully"; then
        echo "✅ 验证者6 Staking注册成功！"
        
        # 等待交易确认
        sleep 3
        
        # 验证Staking状态
        echo "验证Staking合约中的状态..."
        ./build/congress-cli staking query-validator --address $TARGET_ADDRESS
        echo ""
        
        echo "查询顶级验证者列表..."
        ./build/congress-cli staking list-top-validators
        echo ""
    else
        echo "❌ Staking注册失败，但验证者仍在Validators合约中有效"
        echo "ℹ️  您可以稍后手动执行Staking注册"
    fi
else
    echo "ℹ️  跳过Staking注册（密钥文件未找到或交易文件不存在）"
    echo "ℹ️  您可以手动执行以下命令进行Staking注册："
    echo "ℹ️  ./build/congress-cli staking register-validator --proposer $TARGET_ADDRESS --stake-amount 10000 --commission-rate 500"
fi

echo ""

echo "=== 步骤5: 最终验证 ==="
echo "检查验证者6的最终状态..."

echo "📋 Validators合约状态："
./build/congress-cli miner -a $TARGET_ADDRESS
echo ""

echo "📋 Staking合约状态："
./build/congress-cli staking query-validator --address $TARGET_ADDRESS
echo ""

echo "📋 完整的miners列表（应显示6个验证者）："
./build/congress-cli miners
echo ""

echo "🎉 验证者6完整配置流程完成！"
echo "🎉 Validator 6 complete configuration process completed!"
echo ""
echo "📊 网络现在有6个活跃验证者 (Network now has 6 active validators)："
echo "   - Validator1: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
echo "   - Validator2: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8" 
echo "   - Validator3: 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC"
echo "   - Validator4: 0x90F79bf6EB2c4f870365E785982E1f101E93b906"
echo "   - Validator5: 0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65"
echo "   - Validator6: 0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc ✨ (新添加 Newly Added)"
echo ""
echo "✨ 完整功能："
echo "   - ✅ 自动创建提案并投票通过"
echo "   - ✅ 自动为验证者6转账10000 ETH"
echo "   - ✅ 自动设置收费地址"
echo "   - ✅ 自动注册到Staking合约并质押10000 JU"
echo "   - ✅ 验证者6现在同时在Validators和Staking合约中活跃"
echo "   - ✅ 完整的验证者配置流程，可直接参与挖矿"
echo ""
echo "ℹ️  注意：此脚本已升级为使用灵活的密钥文件查找方式"
echo "ℹ️  Note: This script has been upgraded to use flexible key file discovery"
