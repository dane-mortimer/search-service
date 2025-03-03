FROM python:3.11-slim

WORKDIR /app

# Install dependencies
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Copy tests and entrypoint script
COPY . /app/tests

# Make the entrypoint script executable
RUN chmod +x /app/tests/run-tests.sh

# Set the entrypoint
ENTRYPOINT ["./tests/run-tests.sh"]