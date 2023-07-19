#!/bin/bash

# Check if the test name was provided
if [ -z "$1" ]; then
  echo "Error: The test name wasn't provided."
  exit 1
fi

test_name="$1"

# Check if the test's directory exists
if [ ! -d "./tests/${test_name}_test/" ]; then
  echo "Error: The `${test_name}` test doesn't exist."
  exit 1
fi

cd "./tests/${test_name}_test/" || exit 1
go run .
