#!/bin/bash

# Parse command line arguments, separating private key (-k) and --rpc parameter
PRIVATE_KEY=""
RPC_URL=""
OTHER_ARGS=()

# Use Bash array indexing
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

# Check if private key is provided
if [[ -z $PRIVATE_KEY ]]; then
    echo "Error: Private key (-k) is required"
    exit 1
fi

# Step 1: Generating raw transaction...
echo "Step 1: Generating raw transaction..."
go run main.go "${OTHER_ARGS[@]}" --rpc "$RPC_URL"| tee step1_output.txt

# Extract transaction file name from step 1 output
TX_FILE=$(grep -oP 'Transaction file: \K[^\s]+' step1_output.txt)

if [[ -z $TX_FILE ]]; then
    echo "Error: Failed to extract transaction file name from step 1 output"
    exit 1
fi

echo "Extracted transaction file: $TX_FILE"
echo ""

# Step 2: Signing transaction...
echo "Step 2: Signing transaction..."
SIGNED_TX_FILE="${TX_FILE%.*}_signed.json"
go run main.go misc sign -f "$TX_FILE" -k "$PRIVATE_KEY" | tee step2_output.txt
echo ""

# Step 3: Sending transaction...
echo "Step 3: Sending transaction..."
SEND_CMD="go run main.go misc send -f $SIGNED_TX_FILE"
if [[ -n $RPC_URL ]]; then
    SEND_CMD="$SEND_CMD --rpc $RPC_URL"
fi
echo "Executing: $SEND_CMD"
$SEND_CMD | tee step3_output.txt

# Clean up temporary files
rm -f step1_output.txt step2_output.txt step3_output.txt

echo "All steps completed successfully!"
