#!/usr/bin/env bash

# Print a message to indicate the start of the build
echo "Starting Docker build..."

# Build the Docker image
docker build --no-cache -t search-service ./search-service

# Check if the build was successful
if [ $? -eq 0 ]; then
  echo "Docker build completed successfully."
else
  echo "Docker build failed."
  exit 1
fi

