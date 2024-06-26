#!/bin/bash

TEST_DIRS=(
    "./"
    "./test"
)

# 運行測試
for dir in "${TEST_DIRS[@]}"; do
    if [ -d "$dir" ]; then
        echo "Running tests in $dir..."
        go test "$dir" -v
    else
        echo "Directory $dir not found."
    fi
done
