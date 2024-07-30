#!/bin/bash

# Check if the parent directory argument is provided
if [ -z "$1" ]; then
  echo "Usage: $0 /path/to/parent_directory"
  exit 1
fi

# Define the parent directory from the first argument
parent_directory="$1"

# Loop through each subdirectory in the parent directory
for subdirectory in "$parent_directory"/*/; do
  if [ -d "$subdirectory" ]; then
      # Get the name of the subdirectory
    subdirectory_name=$(basename "$subdirectory")
    go get "github.com/aws/aws-sdk-go-v2/service/$subdirectory_name@latest"
  fi
done