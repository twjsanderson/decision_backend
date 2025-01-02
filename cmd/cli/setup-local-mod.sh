#!/bin/bash

GO_MOD_FILE="go.mod"
REPLACE_DIRECTIVE='replace github.com/twjsanderson/decision_backend => ./'

# Check if go.mod file exists
if [ ! -f "$GO_MOD_FILE" ]; then
    echo "Error: $GO_MOD_FILE not found!"
    exit 1
fi

# Check if the replace directive already exists
if grep -Fxq "$REPLACE_DIRECTIVE" "$GO_MOD_FILE"; then
    echo "Replace directive already exists in $GO_MOD_FILE."
else
    # Append the replace directive to the end of the file
    echo "$REPLACE_DIRECTIVE" >> "$GO_MOD_FILE"
    echo "Replace directive added to $GO_MOD_FILE."
fi
