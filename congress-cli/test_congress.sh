#!/bin/bash

# Congress POA 测试脚本

echo "🧪 测试 Congress POA 共识"
echo "========================"

# 等待节点启动
echo "⏳ 等待节点启动..."
sleep 5

# 测试基本连接
echo "🔗 测试 RPC 连接..."
response=$(curl -s -X POST -H "Content-Type: application/json" \
  --data '{"jsonrpc":"2.0","method":"net_version","params":[],"id":1}' \
  http://localhost:8545)

if [[ $response == *"202599"* ]]; then
    echo "✅ RPC 连接成功，Chain ID: 202599"
else
    echo "❌ RPC 连接失败"
    exit 1
fi

# 测试区块高度
echo "📊 检查区块高度..."
block_response=$(curl -s -X POST -H "Content-Type: application/json" \
  --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' \
  http://localhost:8545)

echo "当前区块: $block_response"

# 测试 Congress 特定的 API
echo "🏛️ 测试 Congress API..."
validators_response=$(curl -s -X POST -H "Content-Type: application/json" \
  --data '{"jsonrpc":"2.0","method":"congress_getCurrentValidators","params":[],"id":1}' \
  http://localhost:8545)

echo "当前验证者: $validators_response"

# 检查系统合约
echo "📋 检查系统合约..."
validators_code=$(curl -s -X POST -H "Content-Type: application/json" \
  --data '{"jsonrpc":"2.0","method":"eth_getCode","params":["0x000000000000000000000000000000000000f000", "latest"],"id":1}' \
  http://localhost:8545)

if [[ $validators_code == *"0x"* ]] && [[ ${#validators_code} -gt 10 ]]; then
    echo "✅ Validators 系统合约已部署"
else
    echo "❌ Validators 系统合约未找到"
fi

echo ""
echo "🎯 Congress POA 状态检查完成"
echo "验证者账户: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
echo "系统合约地址:"
echo "  - Validators: 0x000000000000000000000000000000000000f000"
echo "  - Punish: 0x000000000000000000000000000000000000f001" 
echo "  - Proposal: 0x000000000000000000000000000000000000f002"
