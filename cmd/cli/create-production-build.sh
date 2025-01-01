#!/bin/bash

BUILD_OUTPUT="tmp/main"
MAIN_FILE="cmd/server/main.go"

# Step 1: Remove the 'replace' directive from go.mod
sed -i '/replace github.com\/twjsanderson\/decision_backend => \./d' ./go.mod

# Step 2: Tidy up dependencies
go mod tidy

# Step 3: Build the production binary
go build -o tmp/main cmd/server/main.go

echo "Production build created at tmp/main"

# Step 4: Test build
echo "Testing the production binary..."
if [ -f "$BUILD_OUTPUT" ]; then
    echo "Running $BUILD_OUTPUT..."
    
    # Run the binary in the background
    $BUILD_OUTPUT &

    # Capture the process ID of the background job
    PID=$!
    
    # Wait for the process to complete or run for a timeout
    sleep 2

    # Kill the background process after the timeout (if it's still running)
    kill $PID
    
    # Check the exit status of the process
    EXIT_CODE=$?
    if [ $EXIT_CODE -eq 0 ]; then
        echo "Production binary ran successfully."
        exit 0
    else
        echo "Production binary failed with exit code $EXIT_CODE."
        exit $EXIT_CODE
    fi
else
    echo "Error: Binary not found at $BUILD_OUTPUT."
    exit 1
fi
