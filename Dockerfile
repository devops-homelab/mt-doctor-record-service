# Step 1: Build Stage
FROM golang:1.23-alpine AS builder

# Set environment variables for cross-compilation
ENV GOOS=linux GOARCH=amd64

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache for dependency resolution
COPY go.mod go.sum ./

# Download dependencies (this will be cached if go.mod or go.sum doesn't change)
RUN go mod tidy

# Now copy the rest of the source code
COPY . .

# Build the Go app (your executable binary)
RUN go build -o doctor-record-service .

# Step 2: Run Stage (minimal image for running the app and tests)
FROM alpine:3.18

# Install necessary libraries (e.g., CA certificates for HTTPS support)
RUN apk --no-cache add ca-certificates

# Set the working directory for the application
WORKDIR /root/

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/doctor-record-service .

# Copy the source files for unit tests (assuming they're required in the `/tests` directory)
COPY --from=builder /app/tests /root/tests

# Copy go.mod file for testing
COPY --from=builder /app/go.mod /root/go.mod

# Expose port 8084 for the service
EXPOSE 8084

# Command to run the executable
CMD ["./doctor-record-service"]
