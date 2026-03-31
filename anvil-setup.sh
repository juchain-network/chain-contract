#!/bin/bash
# Default values
DEFAULT_PORT=8545
DEFAULT_CHAIN_ID=1337
DEFAULT_BLOCK_TIME=1
DEFAULT_BASE_FEE=0

# Load environment variables from .env file if exists
if [ -f "$(dirname "$0")/../../.env" ]; then
    echo "Loading environment variables from .env file..."
    export $(grep -v '^#' "$(dirname "$0")/../../.env" | xargs)
elif [ -f ".env" ]; then
    echo "Loading environment variables from .env file..."
    export $(grep -v '^#' .env | xargs)
else
    echo "No .env file found, using default values"
fi

# Use environment variables or default values
PORT=${ANVIL_PORT:-$DEFAULT_PORT}
CHAIN_ID=${ANVIL_CHAIN_ID:-$DEFAULT_CHAIN_ID}
BLOCK_TIME=${ANVIL_BLOCK_TIME:-$DEFAULT_BLOCK_TIME}
BASE_FEE=${ANVIL_BASE_FEE:-$DEFAULT_BASE_FEE}

# Function to display help
show_help() {
    echo "Usage: $0 [OPTIONS]"
    echo ""
    echo "Options:"
    echo "  --start     Start the Anvil node"
    echo "  --stop      Stop the Anvil node"
    echo "  --status    Check the status of the Anvil node"
    echo "  --clean     Clean up temporary files"
    echo "  -h, --help  Show this help message"
    echo ""
    echo "Environment Variables (can be set in .env file):"
    echo "  ANVIL_PORT         Port to run Anvil on (default: $DEFAULT_PORT)"
    echo "  ANVIL_CHAIN_ID     Chain ID for Anvil (default: $DEFAULT_CHAIN_ID)"
    echo "  ANVIL_BLOCK_TIME   Block time in seconds (default: $DEFAULT_BLOCK_TIME)"
    echo "  ANVIL_BASE_FEE     Base fee in wei (default: $DEFAULT_BASE_FEE)"
}

# Function to start Anvil
start_anvil() {
    echo "Starting Anvil node..."
    echo "Port: $PORT"
    echo "Chain ID: $CHAIN_ID"
    echo "Block Time: $BLOCK_TIME seconds"
    echo "Base Fee: $BASE_FEE wei"
    
    # Create log directory if not exists
    mkdir -p ./logs
    
    # Start Anvil in background with high initial balance for default account
    anvil --port $PORT --chain-id $CHAIN_ID --block-time $BLOCK_TIME --base-fee $BASE_FEE --balance $INITIAL_FUNDS > ./logs/anvil.log 2>&1 &
    
    # Save PID
    ANVIL_PID=$!
    echo $ANVIL_PID > ./logs/anvil.pid
    
    echo "Anvil node started with PID $ANVIL_PID"
    echo "Logs: ./logs/anvil.log"
    
    # Wait for Anvil to start
    sleep 2
    echo "Anvil node is running and ready!"
    echo "RPC URL: http://localhost:$PORT"
}

# Function to stop Anvil
stop_anvil() {
    if [ -f ./logs/anvil.pid ]; then
        ANVIL_PID=$(cat ./logs/anvil.pid)
        echo "Stopping Anvil node with PID $ANVIL_PID..."
        kill $ANVIL_PID 2>/dev/null || true
        rm -f ./logs/anvil.pid
        echo "Anvil node stopped"
    else
        echo "Anvil node is not running"
    fi
}

# Function to check Anvil status
check_status() {
    if [ -f ./logs/anvil.pid ]; then
        ANVIL_PID=$(cat ./logs/anvil.pid)
        if ps -p $ANVIL_PID > /dev/null; then
            echo "Anvil node is running with PID $ANVIL_PID"
            echo "Port: $PORT"
            echo "Chain ID: $CHAIN_ID"
            echo "RPC URL: http://localhost:$PORT"
        else
            echo "Anvil node PID file exists but process is not running"
            rm -f ./logs/anvil.pid
        fi
    else
        echo "Anvil node is not running"
    fi
}

# Function to clean up
clean_up() {
    echo "Cleaning up temporary files..."
    stop_anvil
    rm -rf ./logs
    echo "Clean up completed"
}

# Main script logic
case "$1" in
    --start)
        start_anvil
        ;;
    --stop)
        stop_anvil
        ;;
    --status)
        check_status
        ;;
    --clean)
        clean_up
        ;;
    -h|--help)
        show_help
        ;;
    *)
        show_help
        exit 1
        ;;
esac
