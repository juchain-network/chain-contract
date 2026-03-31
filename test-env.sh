#!/bin/bash

# load environment variables from .env file
if [ -f ".env" ]; then
    export $(grep -v '^#' ".env" | xargs)
    echo "✅ Environment variables loaded from .env file"
    echo "VALIDATOR_COUNT: $VALIDATOR_COUNT"
    echo "DELEGATOR_COUNT: $DELEGATOR_COUNT"
    echo "EPOCH_DURATION: $EPOCH_DURATION"
else
    echo "⚠️  .env file not found"
fi