#!/bin/bash

# Congress CLI 自动化脚本：添加验证者6
# 使用说明：确保私链环境正在运行，然后执行此脚本
# 
# 特性：
# - 自动查找密钥文件，不依赖硬编码的时间戳
# - 完整的3步流程：创建提案、投票、验证
# - 详细的输出和错误处理

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
PRIVATE_CHAIN_PATH="$HOME/ju-chain-work/chain/private-chain"

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

echo ""
echo "🚀 开始添加验证者6到网络..."
echo "目标地址: $TARGET_ADDRESS"
echo "提案者: $PROPOSER_ADDRESS"
echo ""

echo "=== 步骤1: 创建提案 ==="
echo "创建添加验证者的提案..."
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
echo "📊 网络现在有6个活跃验证者 (Network now has 6 active validators)："
echo "   - Validator1: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
echo "   - Validator2: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8" 
echo "   - Validator3: 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC"
echo "   - Validator4: 0x90F79bf6EB2c4f870365E785982E1f101E93b906"
echo "   - Validator5: 0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65"
echo "   - Validator6: 0x9965507d1a55bcc2695c58ba16fb37d819b0a4dc ✨ (新添加 Newly Added)"
echo ""
echo "ℹ️  注意：此脚本已升级为使用灵活的密钥文件查找方式"
echo "ℹ️  Note: This script has been upgraded to use flexible key file discovery"
