#!/bin/bash

# Loop through each subdirectory in the parent directory
for file in "service"/*; do
  if [ -f "$file" ]; then
      # Get the name of the subdirectory
    filename=$(basename "$file")
    file="${filename%.*}"
    capitalized_name="$(tr '[:lower:]' '[:upper:]' <<< ${file:0:1})${file:1}"
    mockery --name "${capitalized_name}Client" --dir ./service --output ./mocks --outpkg mocks --filename "${file}.go"
  fi
done