#!/bin/bash

# 解析命令行参数，分离出私钥(-k)和--rpc参数
PRIVATE_KEY=""
RPC_URL=""
OTHER_ARGS=()

# 使用Bash数组索引（从1开始）
args=($@)
for ((i=0; i<${#args[@]}; i++)); do
    arg=${args[$i]}
    if [[ $arg == -k ]]; then
        PRIVATE_KEY=${args[$i+1]}
    elif [[ $arg == --rpc ]]; then
        RPC_URL=${args[$i+1]}
    else
        OTHER_ARGS+=("$arg")
    fi
done

# 检查是否提供了私钥
if [[ -z $PRIVATE_KEY ]]; then
    echo "Error: Private key (-k) is required"
    exit 1
fi

# 执行第一步：生成原始交易
echo "Step 1: Generating raw transaction..."
go run main.go "${OTHER_ARGS[@]}" --rpc "$RPC_URL"| tee step1_output.txt

# 从第一步输出中提取交易文件名
TX_FILE=$(grep -oP 'Transaction file: \K[^\s]+' step1_output.txt)

if [[ -z $TX_FILE ]]; then
    echo "Error: Failed to extract transaction file name from step 1 output"
    exit 1
fi

echo "Extracted transaction file: $TX_FILE"
echo ""

# 执行第二步：签名交易
echo "Step 2: Signing transaction..."
SIGNED_TX_FILE="${TX_FILE%.*}_signed.json"
go run main.go misc sign -f "$TX_FILE" -k "$PRIVATE_KEY" | tee step2_output.txt
echo ""

# 执行第三步：发送交易
echo "Step 3: Sending transaction..."
SEND_CMD="go run main.go misc send -f $SIGNED_TX_FILE"
if [[ -n $RPC_URL ]]; then
    SEND_CMD="$SEND_CMD --rpc $RPC_URL"
fi
echo "Executing: $SEND_CMD"
$SEND_CMD | tee step3_output.txt

# 清理临时文件
rm -f step1_output.txt step2_output.txt step3_output.txt

echo "All steps completed successfully!"
