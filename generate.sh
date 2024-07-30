#!/bin/bash

# Check if the parent directory argument is provided
if [ -z "$1" ]; then
  echo "Usage: $0 /path/to/aws-sdk-go-v2"
  exit 1
fi

# Define the parent directory from the first argument
parent_directory="$1"

# Loop through each subdirectory in the parent directory
for subdirectory in "$parent_directory"/*/; do
  if [ -d "$subdirectory" ]; then
      # Get the name of the subdirectory
    subdirectory_name=$(basename "$subdirectory")
    capitalized_name="$(tr '[:lower:]' '[:upper:]' <<< ${subdirectory_name:0:1})${subdirectory_name:1}"
    ifacemaker -f "$subdirectory/*.go" -s Client -i "${capitalized_name}Client" -d false -p service -o "./service/$subdirectory_name.go"
  fi
done