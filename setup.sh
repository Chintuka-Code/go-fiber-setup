#!/bin/bash

# Check if Task is installed
if ! command -v task &> /dev/null
then
    echo "Task is not installed. Installing globally..."

    # Install Task globally using Go
    if ! command -v go &> /dev/null
    then
        echo "Go is not installed. Please install Go first."
        exit 1
    fi

    go install github.com/go-task/task/v3/cmd/task@latest

    # Ensure the Go bin directory is in the PATH
    export PATH=$PATH:$(go env GOPATH)/bin

    echo "Task has been installed."
else
    echo "Task is already installed."
fi

# Run the setup:app task
echo "Running 'task setup:app'..."
task setup:app
