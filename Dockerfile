# Use the official Golang image as the builder stage
FROM golang:1.24 AS builder

RUN apt-get update

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files first to leverage Docker cache
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application (statically compiled for Alpine)
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags="-extldflags=-static" -o main .

# Use a minimal base image for the final stage
FROM alpine:latest

# Install dependency to run go binary
RUN apk add --no-cache libc6-compat

# Set the working directory inside the container
WORKDIR /root

# Copy the compiled binary and database file from the builder stage
COPY --from=builder /app/main .

# Expose the necessary port
EXPOSE $PORT

# Run the application
CMD ["./main"]