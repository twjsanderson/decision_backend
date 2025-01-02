#!/bin/bash

# Check if the current directory is a Go project
if [ ! -f "go.mod" ]; then
    echo "Error: No go.mod file found. Please run this script in the root of a Go project."
    exit 1
fi

# Recursively format all Go files in the project
echo "Formatting all Go files in the project..."
find . -type f -name "*.go" -exec go fmt {} \;

echo "Formatting complete!"
