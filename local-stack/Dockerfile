# Use the official Python 3.11 slim image
FROM python:3.11-slim

# Set the working directory
WORKDIR /app

# Install dependencies: zip, curl, unzip
RUN apt-get update && apt-get install -y zip curl unzip \
    && rm -rf /var/lib/apt/lists/*

# Install AWS CLI v2
RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-aarch64.zip" -o "awscliv2.zip" \
    && unzip awscliv2.zip \
    && ./aws/install \
    && rm -rf awscliv2.zip aws

# Upgrade pip and install awscli-local
RUN pip install --upgrade pip && pip install awscli-local

# Copy tests and entrypoint script
COPY . .

# Make the entrypoint script executable
RUN chmod +x /app/provision-stack.sh

# Set the entrypoint
ENTRYPOINT ["./provision-stack.sh"]