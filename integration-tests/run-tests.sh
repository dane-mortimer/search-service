#!/usr/bin/env bash

# Run pytest and save the exit code
pytest tests/main.py 
TEST_RESULT=$?

# Write the result to a file
if [ $TEST_RESULT -eq 0 ]; then
  echo "healthy" > /tmp/healthstatus
else
  echo "unhealthy" > /tmp/healthstatus
fi

# Exit with the pytest exit code
exit $TEST_RESULT