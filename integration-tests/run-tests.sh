#!/usr/bin/env bash

# Run the tests
pytest tests/main.py

# Capture the exit code of the pytest command
TEST_EXIT_CODE=$?

# If the tests failed, stop all Docker Compose services
if [ $TEST_EXIT_CODE -ne 0 ]; then
  echo "Tests failed with exit code $TEST_EXIT_CODE"
  exit $TEST_EXIT_CODE
fi

# If the tests passed, exit successfully
exit 0